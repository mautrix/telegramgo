package connector

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/updates"
	"github.com/gotd/td/tg"
	"github.com/rs/zerolog"
	"go.mau.fi/zerozap"
	"go.uber.org/zap"
	"maunium.net/go/mautrix/bridgev2"
	"maunium.net/go/mautrix/bridgev2/networkid"

	"go.mau.fi/mautrix-telegram/pkg/connector/ids"
	"go.mau.fi/mautrix-telegram/pkg/connector/media"
	"go.mau.fi/mautrix-telegram/pkg/connector/msgconv"
	"go.mau.fi/mautrix-telegram/pkg/connector/util"
)

type TelegramClient struct {
	main           *TelegramConnector
	telegramUserID int64
	loginID        networkid.UserLoginID
	userID         networkid.UserID
	userLogin      *bridgev2.UserLogin
	client         *telegram.Client
	clientCancel   context.CancelFunc
	msgConv        *msgconv.MessageConverter

	reactionMessageLocks map[int]*sync.Mutex

	appConfig     map[string]any
	appConfigHash int
}

var (
	_ bridgev2.NetworkAPI                    = (*TelegramClient)(nil)
	_ bridgev2.EditHandlingNetworkAPI        = (*TelegramClient)(nil)
	_ bridgev2.ReactionHandlingNetworkAPI    = (*TelegramClient)(nil)
	_ bridgev2.RedactionHandlingNetworkAPI   = (*TelegramClient)(nil)
	_ bridgev2.ReadReceiptHandlingNetworkAPI = (*TelegramClient)(nil)
	_ bridgev2.ReadReceiptHandlingNetworkAPI = (*TelegramClient)(nil)
	_ bridgev2.TypingHandlingNetworkAPI      = (*TelegramClient)(nil)
	// _ bridgev2.IdentifierResolvingNetworkAPI = (*TelegramClient)(nil)
	// _ bridgev2.GroupCreatingNetworkAPI       = (*TelegramClient)(nil)
	// _ bridgev2.ContactListingNetworkAPI      = (*TelegramClient)(nil)
)

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

func NewTelegramClient(ctx context.Context, tc *TelegramConnector, login *bridgev2.UserLogin) (*TelegramClient, error) {
	telegramUserID, err := ids.ParseUserLoginID(login.ID)
	if err != nil {
		return nil, err
	}

	log := zerolog.Ctx(ctx).With().
		Str("component", "telegram_client").
		Str("user_login_id", string(login.ID)).
		Logger()

	zaplog := zap.New(zerozap.New(log))

	client := TelegramClient{
		main:           tc,
		telegramUserID: telegramUserID,
		loginID:        login.ID,
		userID:         networkid.UserID(login.ID),
		userLogin:      login,
	}
	dispatcher := UpdateDispatcher{
		UpdateDispatcher: tg.NewUpdateDispatcher(),
		EntityHandler:    client.onEntityUpdate,
	}
	dispatcher.OnNewMessage(client.onUpdateNewMessage)
	dispatcher.OnNewChannelMessage(client.onUpdateNewChannelMessage)
	dispatcher.OnUserName(client.onUserName)
	dispatcher.OnDeleteMessages(client.onDeleteMessages)
	dispatcher.OnEditMessage(client.onMessageEdit)

	store := tc.Store.GetScopedStore(telegramUserID)

	updatesManager := updates.New(updates.Config{
		OnChannelTooLong: func(channelID int64) {
			log.Error().Int64("channel_id", channelID).Msg("OnChannelTooLong")
			panic("unimplemented channel too long")
		},
		Handler:      dispatcher,
		Logger:       zaplog.Named("gaps"),
		Storage:      store,
		AccessHasher: store,
	})

	client.client = telegram.NewClient(tc.Config.AppID, tc.Config.AppHash, telegram.Options{
		SessionStorage: store,
		Logger:         zaplog,
		UpdateHandler:  updatesManager,
	})
	client.msgConv = msgconv.NewMessageConverter(client.client, tc.Bridge.Matrix, tc.Store, tc.Config.AnimatedSticker, tc.useDirectMedia)
	client.clientCancel, err = connectTelegramClient(ctx, client.client)
	client.reactionMessageLocks = map[int]*sync.Mutex{}
	go func() {
		err = updatesManager.Run(ctx, client.client.API(), telegramUserID, updates.AuthOptions{})
		if err != nil {
			log.Err(err).Msg("updates manager error")
			client.clientCancel()
		}
	}()
	return &client, err
}

// connectTelegramClient blocks until client is connected, calling Run
// internally.
// Technique from: https://github.com/gotd/contrib/blob/master/bg/connect.go
func connectTelegramClient(ctx context.Context, client *telegram.Client) (context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(ctx)

	errC := make(chan error, 1)
	initDone := make(chan struct{})
	go func() {
		defer close(errC)
		errC <- client.Run(ctx, func(ctx context.Context) error {
			close(initDone)
			<-ctx.Done()
			if errors.Is(ctx.Err(), context.Canceled) {
				return nil
			}
			return ctx.Err()
		})
	}()

	select {
	case <-ctx.Done(): // context canceled
		cancel()
		return func() {}, ctx.Err()
	case err := <-errC: // startup timeout
		cancel()
		return func() {}, err
	case <-initDone: // init done
	}

	return cancel, nil
}

func (t *TelegramClient) Connect(ctx context.Context) (err error) {
	t.clientCancel, err = connectTelegramClient(ctx, t.client)
	return
}

func (t *TelegramClient) Disconnect() {
	t.clientCancel()
}

func (t *TelegramClient) GetChatInfo(ctx context.Context, portal *bridgev2.Portal) (*bridgev2.ChatInfo, error) {
	fmt.Printf("%+v\n", portal)
	peerType, id, err := ids.ParsePortalID(portal.ID)
	if err != nil {
		return nil, err
	}
	var name string
	memberList := &bridgev2.ChatMemberList{
		IsFull: true, // TODO not true for channels
	}
	var isSpace, isDM bool
	var avatar *bridgev2.Avatar

	switch peerType {
	case ids.PeerTypeUser:
		users, err := t.client.API().UsersGetUsers(ctx, []tg.InputUserClass{&tg.InputUser{UserID: id}})
		if err != nil {
			return nil, err
		}
		if len(users) == 0 {
			return nil, fmt.Errorf("failed to get user info for user %d", id)
		}
		if user, ok := users[0].(*tg.User); !ok {
			return nil, fmt.Errorf("returned user is not *tg.User")
		} else {
			name = util.FormatFullName(user.FirstName, user.LastName) // TODO gate this behind a config?
			memberList.Members = []bridgev2.ChatMember{
				{
					EventSender: bridgev2.EventSender{
						SenderLogin: ids.MakeUserLoginID(id),
						Sender:      ids.MakeUserID(id),
					},
				},
				{
					EventSender: bridgev2.EventSender{
						IsFromMe:    true,
						SenderLogin: t.loginID,
						Sender:      t.userID,
					},
				},
			}
			isDM = true
		}
	case ids.PeerTypeChat:
		fullChat, err := t.client.API().MessagesGetFullChat(ctx, id)
		if err != nil {
			return nil, err
		}
		for _, c := range fullChat.Chats {
			if c.GetID() == id {
				name = c.(*tg.Chat).Title
				break
			}
		}

		chatFull, ok := fullChat.FullChat.(*tg.ChatFull)
		if !ok {
			return nil, fmt.Errorf("full chat is not %T", chatFull)
		}

		if photo, ok := chatFull.ChatPhoto.(*tg.Photo); ok {
			avatar = &bridgev2.Avatar{
				ID: ids.MakeAvatarID(photo.ID),
				Get: func(ctx context.Context) (data []byte, err error) {
					data, _, _, _, err = media.DownloadPhoto(ctx, t.client.API(), photo)
					return
				},
			}
		}

		for _, user := range fullChat.Users {
			memberList.Members = append(memberList.Members, bridgev2.ChatMember{
				EventSender: bridgev2.EventSender{
					IsFromMe:    user.GetID() == t.telegramUserID,
					SenderLogin: ids.MakeUserLoginID(user.GetID()),
					Sender:      ids.MakeUserID(user.GetID()),
				},
			})
		}
	default:
		fmt.Printf("%s %d\n", peerType, id)
		panic("unimplemented getchatinfo")
	}

	return &bridgev2.ChatInfo{
		Name:         &name,
		Avatar:       avatar,
		Members:      memberList,
		IsDirectChat: &isDM,
		IsSpace:      &isSpace,
	}, nil
}

func (t *TelegramClient) GetUserInfo(ctx context.Context, ghost *bridgev2.Ghost) (*bridgev2.UserInfo, error) {
	id, err := ids.ParseUserID(ghost.ID)
	if err != nil {
		return nil, err
	}
	users, err := t.client.API().UsersGetUsers(ctx, []tg.InputUserClass{&tg.InputUser{UserID: id}})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("failed to get user info for user %d", id)
	}
	if user, ok := users[0].(*tg.User); !ok {
		return nil, fmt.Errorf("returned user is not *tg.User")
	} else {
		return t.getUserInfoFromTelegramUser(user)
	}
}

func (t *TelegramClient) getUserInfoFromTelegramUser(user *tg.User) (*bridgev2.UserInfo, error) {
	var identifiers []string
	for _, username := range user.Usernames {
		identifiers = append(identifiers, fmt.Sprintf("telegram:%s", username.Username))
	}
	if phone, ok := user.GetPhone(); ok {
		identifiers = append(identifiers, fmt.Sprintf("tel:+%s", strings.TrimPrefix(phone, "+")))
	}

	var avatar *bridgev2.Avatar
	if p, ok := user.GetPhoto(); ok && p.TypeID() == tg.UserProfilePhotoTypeID {
		photo := p.(*tg.UserProfilePhoto)
		avatar = &bridgev2.Avatar{
			ID: ids.MakeAvatarID(photo.PhotoID),
			Get: func(ctx context.Context) (data []byte, err error) {
				data, _, err = media.DownloadFileLocation(ctx, t.client.API(), &tg.InputPeerPhotoFileLocation{
					Peer:    &tg.InputPeerUser{UserID: user.ID},
					PhotoID: photo.PhotoID,
					Big:     true,
				})
				return
			},
		}
	}

	name := util.FormatFullName(user.FirstName, user.LastName)
	return &bridgev2.UserInfo{
		IsBot:        &user.Bot,
		Name:         &name,
		Avatar:       avatar,
		Identifiers:  identifiers,
		ExtraUpdates: bridgev2.SimpleMetadataUpdater[*bridgev2.Ghost]("fi.mau.telegram.is_premium", user.Premium),
	}, nil
}

func (t *TelegramClient) IsLoggedIn() bool {
	_, err := t.client.Self(context.TODO())
	return err == nil
}

func (t *TelegramClient) LogoutRemote(ctx context.Context) {
	_, err := t.client.API().AuthLogOut(ctx)
	if err != nil {
		zerolog.Ctx(ctx).Err(err).Msg("failed to logout on Telegram")
	}
}

func (t *TelegramClient) IsThisUser(ctx context.Context, userID networkid.UserID) bool {
	return userID == networkid.UserID(t.userLogin.ID)
}

func (t *TelegramClient) GetCapabilities(ctx context.Context, portal *bridgev2.Portal) *bridgev2.NetworkRoomCapabilities {
	return &bridgev2.NetworkRoomCapabilities{
		FormattedText:    true,
		UserMentions:     true,
		RoomMentions:     true, // TODO?
		LocationMessages: true,
		Captions:         true,
		Threads:          false, // TODO
		Replies:          true,
		Edits:            true,
		Deletes:          true,
		ReadReceipts:     true,
		Reactions:        true,
	}
}
