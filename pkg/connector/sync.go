package connector

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/gotd/td/tg"
	"github.com/rs/zerolog"
	"go.mau.fi/util/ptr"
	"maunium.net/go/mautrix/bridgev2"
	"maunium.net/go/mautrix/bridgev2/database"
	"maunium.net/go/mautrix/bridgev2/networkid"
	"maunium.net/go/mautrix/bridgev2/simplevent"
	"maunium.net/go/mautrix/event"

	"go.mau.fi/mautrix-telegram/pkg/connector/ids"
)

func (t *TelegramClient) SyncChats(ctx context.Context) error {
	limit := t.main.Config.Sync.UpdateLimit
	if limit <= 0 {
		limit = math.MaxInt32
	}

	dialogs, err := APICallWithUpdates(ctx, t, func() (tg.ModifiedMessagesDialogs, error) {
		d, err := t.client.API().MessagesGetDialogs(ctx, &tg.MessagesGetDialogsRequest{
			Limit:      limit,
			OffsetPeer: &tg.InputPeerEmpty{},
		})
		if err != nil {
			return nil, err
		} else if dialogs, ok := d.(tg.ModifiedMessagesDialogs); !ok {
			return nil, fmt.Errorf("unexpected dialogs type %T", d)
		} else {
			return dialogs, nil
		}
	})
	if err != nil {
		return err
	}

	if err := t.resetPinnedDialogs(ctx, dialogs.GetDialogs()); err != nil {
		return err
	}

	return t.handleDialogs(ctx, dialogs, t.main.Config.Sync.CreateLimit)
}

func (t *TelegramClient) resetPinnedDialogs(ctx context.Context, dialogs []tg.DialogClass) error {
	t.userLogin.Metadata.(*UserLoginMetadata).PinnedDialogs = nil
	for _, dialog := range dialogs {
		if dialog.GetPinned() {
			portalKey := t.makePortalKeyFromPeer(dialog.GetPeer())
			t.userLogin.Metadata.(*UserLoginMetadata).PinnedDialogs = append(t.userLogin.Metadata.(*UserLoginMetadata).PinnedDialogs, portalKey.ID)
		}
	}
	return t.userLogin.Save(ctx)
}

func (t *TelegramClient) handleDialogs(ctx context.Context, dialogs tg.ModifiedMessagesDialogs, createLimit int) error {
	log := zerolog.Ctx(ctx)

	users := map[networkid.UserID]tg.UserClass{}
	for _, user := range dialogs.GetUsers() {
		users[ids.MakeUserID(user.GetID())] = user
	}
	chats := map[int64]tg.ChatClass{}
	for _, chat := range dialogs.GetChats() {
		chats[chat.GetID()] = chat
	}
	messages := map[networkid.MessageID]tg.MessageClass{}
	for _, message := range dialogs.GetMessages() {
		messages[ids.GetMessageIDFromMessage(message)] = message
	}

	var created int
	for _, d := range dialogs.GetDialogs() {
		if d.TypeID() != tg.DialogTypeID {
			continue
		}
		dialog := d.(*tg.Dialog)

		log := log.With().
			Stringer("peer", dialog.Peer).
			Int("top_message", dialog.TopMessage).
			Logger()
		log.Debug().Msg("Syncing dialog")

		portalKey := t.makePortalKeyFromPeer(dialog.GetPeer())
		portal, err := t.main.Bridge.GetPortalByKey(ctx, portalKey)
		if err != nil {
			log.Err(err).Msg("Failed to get portal")
			continue
		}
		if dialog.UnreadCount == 0 && !dialog.UnreadMark {
			portal.Metadata.(*PortalMetadata).ReadUpTo = dialog.TopMessage
		}

		// If this is a DM, make sure that the user isn't deleted.
		if user, ok := dialog.Peer.(*tg.PeerUser); ok {
			if users[ids.MakeUserID(user.UserID)].(*tg.User).GetDeleted() {
				log.Debug().Msg("Not syncing portal because user is deleted")
				continue
			}
		}
		chatInfo := bridgev2.ChatInfo{
			CanBackfill: true,
			UserLocal:   &bridgev2.UserLocalPortalInfo{},
			Members: &bridgev2.ChatMemberList{
				MemberMap: map[networkid.UserID]bridgev2.ChatMember{
					t.userID: {
						EventSender: t.mySender(),
						Membership:  event.MembershipJoin,
					},
				},
			},
		}

		switch peer := dialog.Peer.(type) {
		case *tg.PeerUser:
			userID := ids.MakeUserID(peer.UserID)
			if users[userID].(*tg.User).GetDeleted() {
				log.Debug().Int64("user_id", peer.UserID).Msg("Not syncing portal because user is deleted")
				continue
			}
			chatInfo.Members.MemberMap[userID] = bridgev2.ChatMember{
				EventSender: t.senderForUserID(peer.UserID),
				Membership:  event.MembershipJoin,
			}
		case *tg.PeerChat:
			chat := chats[peer.ChatID]
			if chat.TypeID() == tg.ChatForbiddenTypeID {
				log.Debug().
					Int64("chat_id", peer.ChatID).
					Msg("Not syncing portal because chat is forbidden")
				continue
			} else if chat.TypeID() != tg.ChatTypeID {
				log.Debug().
					Int64("chat_id", peer.ChatID).
					Type("chat_type", chat).
					Msg("Not syncing portal because chat type is unsupported")
				continue
			}
			chatInfo.Members.PowerLevels = t.getGroupChatPowerLevels(ctx, chat)
			chatInfo.Name = &chat.(*tg.Chat).Title
		case *tg.PeerChannel:
			channel := chats[peer.ChannelID]
			if channel.TypeID() == tg.ChannelForbiddenTypeID {
				log.Debug().
					Int64("channel_id", peer.ChannelID).
					Msg("Not syncing portal because channel is forbidden")
				continue
			} else if channel.TypeID() != tg.ChannelTypeID {
				log.Debug().
					Int64("channel_id", peer.ChannelID).
					Type("channel_type", channel).
					Msg("Not syncing portal because channel type is unsupported")
				continue
			}
			chatInfo.Name = &channel.(*tg.Channel).Title
			chatInfo.Members.PowerLevels = t.getGroupChatPowerLevels(ctx, channel)
			if !portal.Metadata.(*PortalMetadata).IsSuperGroup {
				// Add the channel user
				sender := ids.MakeChannelUserID(peer.ChannelID)
				chatInfo.Members.MemberMap = map[networkid.UserID]bridgev2.ChatMember{
					sender: {
						EventSender: bridgev2.EventSender{Sender: sender},
						Membership:  event.MembershipJoin,
						PowerLevel:  superadminPowerLevel,
					},
				}
				chatInfo.Members.MemberMap = map[networkid.UserID]bridgev2.ChatMember{
					t.userID: {
						EventSender: t.mySender(),
						Membership:  event.MembershipJoin,
						PowerLevel:  adminRightsToPowerLevel(channel.(*tg.Channel).AdminRights),
					},
				}
			}

		}

		if portal == nil || portal.MXID == "" {
			// Check what the latest message is
			topMessage := messages[ids.MakeMessageID(dialog.Peer, dialog.TopMessage)]
			if topMessage.TypeID() == tg.MessageServiceTypeID {
				action := topMessage.(*tg.MessageService).Action
				if action.TypeID() == tg.MessageActionContactSignUpTypeID || action.TypeID() == tg.MessageActionHistoryClearTypeID {
					log.Debug().Str("action_type", action.TypeName()).Msg("Not syncing portal because it's a contact sign up or history clear")
					continue
				}
			}

			created++ // The portal will have to be created
			if createLimit >= 0 && created > createLimit {
				break
			}
		}

		if mu, ok := dialog.NotifySettings.GetMuteUntil(); ok {
			chatInfo.UserLocal.MutedUntil = ptr.Ptr(time.Unix(int64(mu), 0))
		} else {
			chatInfo.UserLocal.MutedUntil = &bridgev2.Unmuted
		}
		if dialog.Pinned {
			chatInfo.UserLocal.Tag = ptr.Ptr(event.RoomTagFavourite)
		}

		t.main.Bridge.QueueRemoteEvent(t.userLogin, &simplevent.ChatResync{
			ChatInfo: &chatInfo,
			EventMeta: simplevent.EventMeta{
				Type: bridgev2.RemoteEventChatResync,
				LogContext: func(c zerolog.Context) zerolog.Context {
					return c.Str("update", "sync")
				},
				PortalKey:    portalKey,
				CreatePortal: true,
			},
			CheckNeedsBackfillFunc: func(ctx context.Context, latestMessage *database.Message) (bool, error) {
				if latestMessage == nil {
					return true, nil
				}
				_, latestMessageID, err := ids.ParseMessageID(latestMessage.ID)
				if err != nil {
					panic(err)
				}
				return dialog.TopMessage > latestMessageID, nil
			},
		})
	}
	return nil
}
