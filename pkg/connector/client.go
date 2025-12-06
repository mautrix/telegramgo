// mautrix-telegram - A Matrix-Telegram puppeting bridge.
// Copyright (C) 2025 Sumner Evans
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package connector

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"go.mau.fi/util/exsync"
	"go.mau.fi/zerozap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"maunium.net/go/mautrix/bridgev2"
	"maunium.net/go/mautrix/bridgev2/database"
	"maunium.net/go/mautrix/bridgev2/networkid"
	"maunium.net/go/mautrix/bridgev2/simplevent"
	"maunium.net/go/mautrix/bridgev2/status"
	"maunium.net/go/mautrix/id"

	"go.mau.fi/mautrix-telegram/pkg/connector/humanise"
	"go.mau.fi/mautrix-telegram/pkg/connector/ids"
	"go.mau.fi/mautrix-telegram/pkg/connector/matrixfmt"
	"go.mau.fi/mautrix-telegram/pkg/connector/store"
	"go.mau.fi/mautrix-telegram/pkg/connector/telegramfmt"
	"go.mau.fi/mautrix-telegram/pkg/connector/util"
	"go.mau.fi/mautrix-telegram/pkg/gotd/telegram"
	"go.mau.fi/mautrix-telegram/pkg/gotd/telegram/auth"
	"go.mau.fi/mautrix-telegram/pkg/gotd/telegram/updates"
	"go.mau.fi/mautrix-telegram/pkg/gotd/tg"
)

var (
	ErrNoAuthKey        = errors.New("user does not have auth key")
	ErrFailToQueueEvent = errors.New("failed to queue event")
)

func resultToError(res bridgev2.EventHandlingResult) error {
	if !res.Success {
		if res.Error != nil {
			return fmt.Errorf("%w: %w", ErrFailToQueueEvent, res.Error)
		}
		return ErrFailToQueueEvent
	}
	return nil
}

type TelegramClient struct {
	main              *TelegramConnector
	ScopedStore       *store.ScopedStore
	telegramUserID    int64
	loginID           networkid.UserLoginID
	userID            networkid.UserID
	userLogin         *bridgev2.UserLogin
	metadata          *UserLoginMetadata
	client            *telegram.Client
	updatesManager    *updates.Manager
	clientCtx         context.Context
	clientCancel      context.CancelFunc
	clientDone        *Future[error]
	clientInitialized *exsync.Event
	mu                sync.Mutex

	appConfigLock sync.Mutex
	appConfig     map[string]any
	appConfigHash int

	availableReactionsLock    sync.Mutex
	availableReactions        map[string]struct{}
	availableReactionsHash    int
	availableReactionsFetched time.Time
	availableReactionsList    []string
	isPremiumCache            atomic.Bool

	telegramFmtParams *telegramfmt.FormatParams
	matrixParser      *matrixfmt.HTMLParser

	cachedContacts     *tg.ContactsContacts
	cachedContactsHash int64

	takeoutLock        sync.Mutex
	takeoutAccepted    *exsync.Event
	stopTakeoutTimer   *time.Timer
	takeoutDialogsOnce sync.Once

	prevReactionPoll     map[networkid.PortalKey]time.Time
	prevReactionPollLock sync.Mutex
}

var _ bridgev2.NetworkAPI = (*TelegramClient)(nil)

type UpdateDispatcher struct {
	tg.UpdateDispatcher
	EntityHandler func(context.Context, tg.Entities) error
}

func (u UpdateDispatcher) Handle(ctx context.Context, updates tg.UpdatesClass) error {
	var e tg.Entities
	switch u := updates.(type) {
	case *tg.Updates:
		e.Users = u.MapUsers().NotEmptyToMap()
		chats := u.MapChats()
		e.Chats = chats.ChatToMap()
		e.Channels = chats.ChannelToMap()
	case *tg.UpdatesCombined:
		e.Users = u.MapUsers().NotEmptyToMap()
		chats := u.MapChats()
		e.Chats = chats.ChatToMap()
		e.Channels = chats.ChannelToMap()
	}
	if u.EntityHandler != nil {
		u.EntityHandler(ctx, e)
	}

	return u.UpdateDispatcher.Handle(ctx, updates)
}

var messageLinkRegex = regexp.MustCompile(`^https?://t(?:elegram)?\.(?:me|dog)/([A-Za-z][A-Za-z0-9_]{3,31}[A-Za-z0-9]|[Cc]/[0-9]{1,20})/([0-9]{1,20})(?:/([0-9]{1,20}))?$`)

func (tg *TelegramConnector) deviceConfig() telegram.DeviceConfig {
	return telegram.DeviceConfig{
		DeviceModel:    tg.Config.DeviceInfo.DeviceModel,
		SystemVersion:  tg.Config.DeviceInfo.SystemVersion,
		AppVersion:     tg.Config.DeviceInfo.AppVersion,
		SystemLangCode: tg.Config.DeviceInfo.SystemLangCode,
		LangCode:       tg.Config.DeviceInfo.LangCode,
	}
}

var zapLevelMap = map[zapcore.Level]zerolog.Level{
	// shifted
	zapcore.DebugLevel: zerolog.TraceLevel,
	zapcore.InfoLevel:  zerolog.DebugLevel,

	// direct mapping
	zapcore.WarnLevel:   zerolog.WarnLevel,
	zapcore.ErrorLevel:  zerolog.ErrorLevel,
	zapcore.DPanicLevel: zerolog.PanicLevel,
	zapcore.PanicLevel:  zerolog.PanicLevel,
	zapcore.FatalLevel:  zerolog.FatalLevel,
}

func NewTelegramClient(ctx context.Context, tc *TelegramConnector, login *bridgev2.UserLogin) (*TelegramClient, error) {
	telegramUserID, err := ids.ParseUserLoginID(login.ID)
	if err != nil {
		return nil, err
	}

	log := zerolog.Ctx(ctx).With().
		Str("component", "telegram_client").
		Str("user_login_id", string(login.ID)).
		Logger()

	zaplog := zap.New(zerozap.NewWithLevels(log, zapLevelMap))

	client := TelegramClient{
		ScopedStore: tc.Store.GetScopedStore(telegramUserID),

		main:           tc,
		telegramUserID: telegramUserID,
		loginID:        login.ID,
		userID:         networkid.UserID(login.ID),
		userLogin:      login,
		metadata:       login.Metadata.(*UserLoginMetadata),

		takeoutAccepted: exsync.NewEvent(),

		prevReactionPoll: map[networkid.PortalKey]time.Time{},

		clientInitialized: exsync.NewEvent(),
	}

	if !login.Metadata.(*UserLoginMetadata).Session.HasAuthKey() {
		return &client, nil
	}

	dispatcher := UpdateDispatcher{
		UpdateDispatcher: tg.NewUpdateDispatcher(),
		EntityHandler:    client.onEntityUpdate,
	}
	dispatcher.OnNewMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateNewMessage) error {
		return client.onUpdateNewMessage(ctx, e, update)
	})
	dispatcher.OnNewChannelMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateNewChannelMessage) error {
		return client.onUpdateNewMessage(ctx, e, update)
	})
	dispatcher.OnChannel(client.onUpdateChannel)
	dispatcher.OnUserName(client.onUserName)
	dispatcher.OnDeleteMessages(func(ctx context.Context, e tg.Entities, update *tg.UpdateDeleteMessages) error {
		return client.onDeleteMessages(ctx, 0, update)
	})
	dispatcher.OnDeleteChannelMessages(func(ctx context.Context, e tg.Entities, update *tg.UpdateDeleteChannelMessages) error {
		return client.onDeleteMessages(ctx, update.ChannelID, update)
	})
	dispatcher.OnEditMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateEditMessage) error {
		return client.onMessageEdit(ctx, update)
	})
	dispatcher.OnEditChannelMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateEditChannelMessage) error {
		return client.onMessageEdit(ctx, update)
	})
	dispatcher.OnUserTyping(func(ctx context.Context, e tg.Entities, update *tg.UpdateUserTyping) error {
		return client.handleTyping(client.makePortalKeyFromID(ids.PeerTypeUser, update.UserID, 0), client.senderForUserID(update.UserID), update.Action)
	})
	dispatcher.OnChatUserTyping(func(ctx context.Context, e tg.Entities, update *tg.UpdateChatUserTyping) error {
		if update.FromID.TypeID() != tg.PeerUserTypeID {
			log.Warn().Str("from_id_type", update.FromID.TypeName()).Msg("unsupported from_id type")
			return nil
		}
		return client.handleTyping(client.makePortalKeyFromID(ids.PeerTypeChat, update.ChatID, 0), client.getPeerSender(update.FromID), update.Action)
	})
	dispatcher.OnChannelUserTyping(func(ctx context.Context, e tg.Entities, update *tg.UpdateChannelUserTyping) error {
		return client.handleTyping(client.makePortalKeyFromID(ids.PeerTypeChannel, update.ChannelID, update.TopMsgID), client.getPeerSender(update.FromID), update.Action)
	})
	dispatcher.OnReadHistoryOutbox(client.updateReadReceipt)
	dispatcher.OnReadHistoryInbox(func(ctx context.Context, e tg.Entities, update *tg.UpdateReadHistoryInbox) error {
		return client.onOwnReadReceipt(client.makePortalKeyFromPeer(update.Peer, update.TopMsgID), update.MaxID)
	})
	dispatcher.OnReadChannelInbox(func(ctx context.Context, e tg.Entities, update *tg.UpdateReadChannelInbox) error {
		return client.onOwnReadReceipt(client.makePortalKeyFromID(ids.PeerTypeChannel, update.ChannelID, 0), update.MaxID)
	})
	dispatcher.OnNotifySettings(client.onNotifySettings)
	dispatcher.OnPinnedDialogs(client.onPinnedDialogs)
	dispatcher.OnChatDefaultBannedRights(client.onChatDefaultBannedRights)
	dispatcher.OnPeerBlocked(client.onPeerBlocked)
	dispatcher.OnChat(client.onChat)
	dispatcher.OnPhoneCall(client.onPhoneCall)

	client.updatesManager = updates.New(updates.Config{
		OnChannelTooLong: func(channelID int64) error {
			// TODO resync topics?
			res := tc.Bridge.QueueRemoteEvent(login, &simplevent.ChatResync{
				EventMeta: simplevent.EventMeta{
					Type: bridgev2.RemoteEventChatResync,
					LogContext: func(c zerolog.Context) zerolog.Context {
						return c.Str("update", "channel_too_long").Int64("channel_id", channelID)
					},
					PortalKey: client.makePortalKeyFromID(ids.PeerTypeChannel, channelID, 0),
				},
				CheckNeedsBackfillFunc: func(ctx context.Context, latestMessage *database.Message) (bool, error) {
					return true, nil
				},
			})

			return resultToError(res)
		},
		Handler:      dispatcher,
		Logger:       zaplog.Named("gaps"),
		Storage:      client.ScopedStore,
		AccessHasher: client.ScopedStore,
	})

	client.client = telegram.NewClient(tc.Config.APIID, tc.Config.APIHash, telegram.Options{
		CustomSessionStorage: &login.Metadata.(*UserLoginMetadata).Session,
		Logger:               zaplog,
		UpdateHandler:        client.updatesManager,
		OnDead:               client.onDead,
		OnSession:            client.onSession,
		OnConnected:          client.onConnected,
		PingCallback:         client.onPing,
		OnAuthError:          client.onAuthError,
		PingTimeout:          time.Duration(tc.Config.Ping.TimeoutSeconds) * time.Second,
		PingInterval:         time.Duration(tc.Config.Ping.IntervalSeconds) * time.Second,
		Device:               tc.deviceConfig(),
	})

	client.telegramFmtParams = &telegramfmt.FormatParams{
		GetUserInfoByID: func(ctx context.Context, id int64) (telegramfmt.UserInfo, error) {
			ghost, err := tc.Bridge.GetGhostByID(ctx, ids.MakeUserID(id))
			if err != nil {
				return telegramfmt.UserInfo{}, err
			}
			userInfo := telegramfmt.UserInfo{MXID: ghost.Intent.GetMXID(), Name: ghost.Name}
			// FIXME this should look for user logins by ID, not hardcode the current user
			if id == client.telegramUserID {
				userInfo.MXID = client.userLogin.UserMXID
			}
			return userInfo, nil
		},
		GetUserInfoByUsername: func(ctx context.Context, username string) (telegramfmt.UserInfo, error) {
			if peerType, userID, err := client.main.Store.Username.GetEntityID(ctx, username); err != nil {
				return telegramfmt.UserInfo{}, err
			} else if peerType != ids.PeerTypeUser {
				return telegramfmt.UserInfo{}, fmt.Errorf("unexpected peer type: %s", peerType)
			} else if ghost, err := tc.Bridge.GetGhostByID(ctx, ids.MakeUserID(userID)); err != nil {
				return telegramfmt.UserInfo{}, err
			} else {
				userInfo := telegramfmt.UserInfo{MXID: ghost.Intent.GetMXID(), Name: ghost.Name}
				if ghost.ID == client.userID {
					userInfo.MXID = client.userLogin.UserMXID
				}
				return userInfo, nil
			}
		},
		NormalizeURL: func(ctx context.Context, url string) string {
			log := zerolog.Ctx(ctx).With().
				Str("conversion_direction", "to_matrix").
				Str("entity_type", "url").
				Logger()

			if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "ftp://") && !strings.HasPrefix(url, "magnet://") {
				url = "http://" + url
			}

			submatches := messageLinkRegex.FindStringSubmatch(url)
			if len(submatches) == 0 {
				return url
			}
			group := submatches[1]
			msgID, err := strconv.Atoi(submatches[2])
			if err != nil {
				log.Err(err).Msg("error parsing message ID")
				return url
			}
			var topicID int
			if len(submatches) == 4 && submatches[3] != "" {
				lastID, err := strconv.Atoi(submatches[3])
				if err != nil {
					log.Err(err).Msg("error parsing actual message ID")
					return url
				}
				topicID = msgID
				msgID = lastID
			}
			log = log.With().Str("group", group).Int("topic_id", topicID).Int("msg_id", msgID).Logger()

			var portalKey networkid.PortalKey
			if strings.HasPrefix(group, "C/") || strings.HasPrefix(group, "c/") {
				chatID, err := strconv.ParseInt(submatches[1][2:], 10, 64)
				if err != nil {
					log.Err(err).Msg("error parsing channel ID")
					return url
				}
				portalKey = client.makePortalKeyFromID(ids.PeerTypeChannel, chatID, topicID)
			} else if submatches[1] == "premium" {
				portalKey = client.makePortalKeyFromID(ids.PeerTypeUser, 777000, 0)
			} else if userID, err := strconv.ParseInt(submatches[1], 10, 64); err == nil && userID > 0 {
				portalKey = client.makePortalKeyFromID(ids.PeerTypeUser, userID, 0)
			} else if peerType, peerID, err := client.main.Store.Username.GetEntityID(ctx, submatches[1]); err != nil {
				log.Err(err).Msg("Failed to get entity ID by username")
				return url
			} else if peerType != "" {
				portalKey = client.makePortalKeyFromID(peerType, peerID, topicID)
			} else {
				return url
			}

			portal, err := tc.Bridge.DB.Portal.GetByKey(ctx, portalKey)
			if err != nil {
				log.Err(err).Msg("error getting portal")
				return url
			} else if portal == nil {
				log.Warn().Msg("portal not found")
				return url
			}

			message, err := tc.Bridge.DB.Message.GetFirstPartByID(ctx, client.loginID, ids.MakeMessageID(portalKey, msgID))
			if err != nil {
				log.Err(err).Msg("error getting referenced message")
				return url
			} else if message == nil {
				log.Warn().Err(err).Msg("message not found")
				return url
			}

			return portal.MXID.EventURI(message.MXID, tc.Bridge.Matrix.ServerName()).MatrixToURL()
		},
	}
	client.matrixParser = &matrixfmt.HTMLParser{
		GetGhostDetails: func(ctx context.Context, ui id.UserID) (networkid.UserID, string, int64, bool) {
			if userID, ok := tc.Bridge.Matrix.ParseGhostMXID(ui); !ok {
				return "", "", 0, false
			} else if peerType, telegramUserID, err := ids.ParseUserID(userID); err != nil {
				return "", "", 0, false
			} else if accessHash, err := client.ScopedStore.GetAccessHash(ctx, peerType, telegramUserID); err != nil || accessHash == 0 {
				return "", "", 0, false
			} else if username, err := client.main.Store.Username.Get(ctx, peerType, telegramUserID); err != nil {
				return "", "", 0, false
			} else {
				return userID, username, accessHash, true
			}
		},
	}

	return &client, err
}

func (t *TelegramClient) onDead() {
	prevState := t.userLogin.BridgeState.GetPrev().StateEvent
	if slices.Contains([]status.BridgeStateEvent{
		status.StateTransientDisconnect,
		status.StateBadCredentials,
		status.StateLoggedOut,
		status.StateUnknownError,
	}, prevState) {
		t.userLogin.Log.Warn().
			Str("prev_state", string(prevState)).
			Msg("client is dead, not sending transient disconnect, because already in an error state")
		return
	}
	t.userLogin.BridgeState.Send(status.BridgeState{
		StateEvent: status.StateTransientDisconnect,
		Message:    "Telegram client disconnected",
	})
}

func (t *TelegramClient) sendBadCredentialsOrUnknownError(err error) {
	if auth.IsUnauthorized(err) || errors.Is(err, ErrNoAuthKey) {
		t.userLogin.BridgeState.Send(status.BridgeState{
			StateEvent: status.StateBadCredentials,
			Error:      "tg-no-auth",
			Message:    humanise.Error(err),
		})
	} else {
		t.userLogin.BridgeState.Send(status.BridgeState{
			StateEvent: status.StateUnknownError,
			Error:      "tg-unknown-error",
			Message:    humanise.Error(err),
		})
	}
}

func (t *TelegramClient) onPing() {
	if t.userLogin.BridgeState.GetPrev().StateEvent == status.StateConnected {
		return
	}
	ctx := t.userLogin.Log.WithContext(t.main.Bridge.BackgroundCtx)
	t.userLogin.Log.Debug().Msg("Got ping while not connected, checking auth")

	me, err := t.client.Self(ctx)
	if auth.IsUnauthorized(err) {
		t.onAuthError(fmt.Errorf("not logged in"))
	} else if errors.Is(err, syscall.EPIPE) {
		// This is a pipe error, try disconnecting which will force the
		// updatesManager to fail and cause the client to reconnect.
		t.userLogin.BridgeState.Send(status.BridgeState{
			StateEvent: status.StateTransientDisconnect,
			Error:      "pipe-error",
			Message:    humanise.Error(err),
		})
	} else if err != nil {
		t.sendBadCredentialsOrUnknownError(err)
	} else {
		t.onConnected(me)
	}
}

func userToRemoteProfile(
	self *tg.User,
	ghost *bridgev2.Ghost,
	prevState *status.RemoteProfile,
) (profile status.RemoteProfile, name string) {
	profile.Name = util.FormatFullName(self.FirstName, self.LastName, self.Deleted, self.ID)
	if self.Phone != "" {
		profile.Phone = "+" + strings.TrimPrefix(self.Phone, "+")
	} else if prevState != nil {
		profile.Phone = prevState.Phone
	}
	profile.Username = self.Username
	if self.Username == "" && len(self.Usernames) > 0 {
		profile.Username = self.Usernames[0].Username
	}
	if ghost != nil {
		profile.Avatar = ghost.AvatarMXC
	} else if prevState != nil {
		profile.Avatar = prevState.Avatar
	}
	name = cmp.Or(profile.Username, profile.Phone, profile.Name)
	return
}

func (t *TelegramClient) updateRemoteProfile(ctx context.Context, self *tg.User, ghost *bridgev2.Ghost) bool {
	newProfile, newName := userToRemoteProfile(self, ghost, &t.userLogin.RemoteProfile)
	if t.userLogin.RemoteProfile != newProfile || t.userLogin.RemoteName != newName {
		t.userLogin.RemoteProfile = newProfile
		t.userLogin.RemoteName = newName
		err := t.userLogin.Save(ctx)
		if err != nil {
			t.userLogin.Log.Err(err).Msg("Failed to save user login after profile update")
		}
		return true
	}
	return false
}

func (t *TelegramClient) onConnected(self *tg.User) {
	log := t.userLogin.Log
	ctx := log.WithContext(t.main.Bridge.BackgroundCtx)
	ghost, err := t.main.Bridge.GetGhostByID(ctx, t.userID)
	if err != nil {
		log.Err(err).Msg("Failed to get own ghost")
	} else if wrapped, err := t.wrapUserInfo(ctx, self); err != nil {
		log.Err(err).Msg("Failed to wrap own user info")
	} else {
		ghost.UpdateInfo(ctx, wrapped)
	}

	t.updateRemoteProfile(ctx, self, ghost)
	t.userLogin.BridgeState.Send(status.BridgeState{StateEvent: status.StateConnected})
}

func (t *TelegramClient) onSession() {
	t.userLogin.Log.Debug().Msg("Got session created event")
}

func (t *TelegramClient) onAuthError(err error) {
	t.sendBadCredentialsOrUnknownError(err)
	t.metadata.ResetOnLogout()
	go func() {
		if err := t.userLogin.Save(context.Background()); err != nil {
			t.main.Bridge.Log.Err(err).Msg("failed to save user login")
		}
	}()
}

func (t *TelegramClient) Connect(_ context.Context) {
	t.mu.Lock()
	defer t.mu.Unlock()

	ctx := context.Background()

	log := zerolog.Ctx(context.Background()).With().Int64("user_id", t.telegramUserID).Logger()
	ctx = log.WithContext(ctx)

	if !t.metadata.Session.HasAuthKey() {
		log.Warn().Msg("user does not have an auth key, sending bad credentials state")
		t.sendBadCredentialsOrUnknownError(ErrNoAuthKey)
		return
	}

	t.userLogin.BridgeState.Send(status.BridgeState{StateEvent: status.StateConnecting})

	log.Info().Msg("Connecting client")

	// Add a cancellation layer we can use for explicit Disconnect

	ctx, cancel := context.WithCancel(ctx)
	t.clientCtx = ctx
	t.clientCancel = cancel
	t.clientDone = NewFuture[error]()
	t.clientInitialized.Clear()

	runTelegramClient(ctx, t.client, t.clientInitialized, t.clientDone, func(ctx context.Context) error {
		log.Info().Msg("Client running starting updates")
		return t.updatesManager.Run(ctx, t.client.API(), t.telegramUserID, updates.AuthOptions{})
	})
}

func runTelegramClient(ctx context.Context, client *telegram.Client, initialized *exsync.Event, done *Future[error], callback func(ctx context.Context) error) {
	go func() {
		err := client.Run(ctx, func(ctx context.Context) error {
			initialized.Set()
			return callback(ctx)
		})
		initialized.Set()
		done.Set(err)
	}()
}

func (t *TelegramClient) Disconnect() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.userLogin.Log.Info().Msg("Disconnecting client")

	if t.clientCancel != nil {
		t.clientCancel()
		t.userLogin.Log.Info().Msg("Waiting for client")
		err, _ := t.clientDone.Get(context.Background())
		t.userLogin.Log.Info().Err(err).Msg("Client done")
	}

	t.userLogin.Log.Info().Msg("Disconnect complete")
}

func (t *TelegramClient) getInputUser(ctx context.Context, id int64) (*tg.InputUser, error) {
	accessHash, err := t.ScopedStore.GetAccessHash(ctx, ids.PeerTypeUser, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get access hash for user %d: %w", id, err)
	}
	return &tg.InputUser{UserID: id, AccessHash: accessHash}, nil
}

func (t *TelegramClient) getSingleUser(ctx context.Context, id int64) (tg.UserClass, error) {
	if inputUser, err := t.getInputUser(ctx, id); err != nil {
		return nil, err
	} else if users, err := t.client.API().UsersGetUsers(ctx, []tg.InputUserClass{inputUser}); err != nil {
		return nil, err
	} else if len(users) == 0 {
		// TODO does this mean the user is deleted? Need to handle this a bit better
		return nil, fmt.Errorf("failed to get user info for user %d", id)
	} else {
		return users[0], nil
	}
}

func (t *TelegramClient) getSingleChannel(ctx context.Context, id int64) (*tg.Channel, error) {
	accessHash, err := t.ScopedStore.GetAccessHash(ctx, ids.PeerTypeChannel, id)
	if err != nil {
		return nil, err
	}
	chats, err := APICallWithOnlyChatUpdates(ctx, t, func() (tg.MessagesChatsClass, error) {
		return t.client.API().ChannelsGetChannels(ctx, []tg.InputChannelClass{
			&tg.InputChannel{ChannelID: id, AccessHash: accessHash},
		})
	})
	if err != nil {
		return nil, err
	} else if len(chats.GetChats()) == 0 {
		return nil, fmt.Errorf("failed to get channel info for channel %d", id)
	} else if channel, ok := chats.GetChats()[0].(*tg.Channel); !ok {
		return nil, fmt.Errorf("unexpected channel type %T", chats.GetChats()[id])
	} else {
		return channel, nil
	}
}

func (t *TelegramClient) IsLoggedIn() bool {
	// TODO use less hacky check than context cancellation
	return t != nil && t.clientCtx != nil && t.client != nil && t.clientCtx.Err() == nil &&
		t.metadata.Session.HasAuthKey()
}

func (t *TelegramClient) LogoutRemote(ctx context.Context) {
	log := zerolog.Ctx(ctx).With().
		Str("action", "logout_remote").
		Int64("user_id", t.telegramUserID).
		Logger()

	log.Info().Msg("Logging out and disconnecting")

	if t.metadata.Session.HasAuthKey() {
		log.Info().Msg("User has an auth key, logging out")

		// logging out is best effort, we want to logout even if we can't call the endpoint
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_, err := t.client.API().AuthLogOut(ctx)
		if err != nil {
			log.Err(err).Msg("failed to logout on Telegram")
		}
	}

	t.Disconnect()

	log.Info().Msg("Deleting user state")

	err := t.ScopedStore.DeleteUserState(ctx)
	if err != nil {
		log.Err(err).Msg("failed to delete user state")
	}

	err = t.ScopedStore.DeleteChannelStateForUser(ctx)
	if err != nil {
		log.Err(err).Msg("failed to delete channel state for user")
	}

	err = t.ScopedStore.DeleteAccessHashesForUser(ctx)
	if err != nil {
		log.Err(err).Msg("failed to delete access hashes for user")
	}

	log.Info().Msg("Logged out and deleted user state")
}

func (t *TelegramClient) IsThisUser(ctx context.Context, userID networkid.UserID) bool {
	return userID == networkid.UserID(t.userLogin.ID)
}

func (t *TelegramClient) mySender() bridgev2.EventSender {
	return bridgev2.EventSender{
		IsFromMe:    true,
		SenderLogin: t.loginID,
		Sender:      t.userID,
	}
}

func (t *TelegramClient) senderForUserID(userID int64) bridgev2.EventSender {
	return bridgev2.EventSender{
		IsFromMe:    userID == t.telegramUserID,
		SenderLogin: ids.MakeUserLoginID(userID),
		Sender:      ids.MakeUserID(userID),
	}
}

func (t *TelegramClient) FillBridgeState(state status.BridgeState) status.BridgeState {
	if state.Info == nil {
		state.Info = make(map[string]any)
	}
	state.Info["is_bot"] = t.metadata.IsBot
	state.Info["login_method"] = t.metadata.LoginMethod
	return state
}
