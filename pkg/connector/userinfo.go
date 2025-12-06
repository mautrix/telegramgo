package connector

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"maunium.net/go/mautrix/bridgev2"

	"go.mau.fi/mautrix-telegram/pkg/connector/ids"
	"go.mau.fi/mautrix-telegram/pkg/connector/util"
	"go.mau.fi/mautrix-telegram/pkg/gotd/tg"
)

func (t *TelegramClient) GetUserInfo(ctx context.Context, ghost *bridgev2.Ghost) (*bridgev2.UserInfo, error) {
	peerType, id, err := ids.ParseUserID(ghost.ID)
	if err != nil {
		return nil, err
	}
	switch peerType {
	case ids.PeerTypeUser:
		if user, err := t.getSingleUser(ctx, id); err != nil {
			return nil, fmt.Errorf("failed to get user %d: %w", id, err)
		} else {
			return t.wrapUserInfo(ctx, user)
		}
	case ids.PeerTypeChannel:
		if channel, err := t.getSingleChannel(ctx, id); err != nil {
			return nil, fmt.Errorf("failed to get channel %d: %w", id, err)
		} else {
			return t.wrapChannelGhostInfo(ctx, channel)
		}
	default:
		return nil, fmt.Errorf("unexpected peer type: %s", peerType)
	}
}

func (t *TelegramClient) wrapChannelGhostInfo(ctx context.Context, channel *tg.Channel) (*bridgev2.UserInfo, error) {
	var err error
	if accessHash, ok := channel.GetAccessHash(); ok && !channel.Min {
		if err = t.ScopedStore.SetAccessHash(ctx, ids.PeerTypeChannel, channel.ID, accessHash); err != nil {
			return nil, err
		}
	}

	if !channel.Broadcast {
		return nil, nil
	}

	var avatar *bridgev2.Avatar
	avatar, err = t.convertChatPhoto(channel.AsInputPeer(), channel.GetPhoto())
	if err != nil {
		return nil, err
	}

	var identifiers []string
	if username, set := channel.GetUsername(); set {
		err = t.main.Store.Username.Set(ctx, ids.PeerTypeChannel, channel.ID, username)
		if err != nil {
			return nil, err
		}
		identifiers = append(identifiers, fmt.Sprintf("telegram:%s", username))
	}

	return &bridgev2.UserInfo{
		Name:        &channel.Title,
		Avatar:      avatar,
		Identifiers: identifiers,
		ExtraUpdates: func(ctx context.Context, g *bridgev2.Ghost) bool {
			updated := !g.Metadata.(*GhostMetadata).IsChannel
			g.Metadata.(*GhostMetadata).IsChannel = true
			return updated
		},
	}, nil
}

func (t *TelegramClient) wrapUserInfo(ctx context.Context, u tg.UserClass) (*bridgev2.UserInfo, error) {
	user, ok := u.(*tg.User)
	if !ok {
		return nil, fmt.Errorf("user is %T not *tg.User", user)
	}
	var identifiers []string
	if !user.Min {
		if accessHash, ok := user.GetAccessHash(); ok {
			if err := t.ScopedStore.SetAccessHash(ctx, ids.PeerTypeUser, user.ID, accessHash); err != nil {
				return nil, err
			}
		}

		if err := t.main.Store.Username.Set(ctx, ids.PeerTypeUser, user.ID, user.Username); err != nil {
			return nil, err
		}

		if user.Username != "" {
			identifiers = append(identifiers, fmt.Sprintf("telegram:%s", user.Username))
		}
		for _, username := range user.Usernames {
			identifiers = append(identifiers, fmt.Sprintf("telegram:%s", username.Username))
		}
		if phone, ok := user.GetPhone(); ok {
			normalized := strings.TrimPrefix(phone, "+")
			identifiers = append(identifiers, fmt.Sprintf("tel:+%s", normalized))
			if err := t.main.Store.PhoneNumber.Set(ctx, user.ID, normalized); err != nil {
				return nil, err
			}
		}
	}
	slices.Sort(identifiers)
	identifiers = slices.Compact(identifiers)

	var avatar *bridgev2.Avatar
	if p, ok := user.GetPhoto(); ok && p.TypeID() == tg.UserProfilePhotoTypeID {
		photo := p.(*tg.UserProfilePhoto)
		var err error
		avatar, err = t.convertUserProfilePhoto(ctx, user.ID, photo)
		if err != nil {
			return nil, err
		}
	}

	name := util.FormatFullName(user.FirstName, user.LastName, user.Deleted, user.ID)
	return &bridgev2.UserInfo{
		IsBot:       &user.Bot,
		Name:        &name,
		Avatar:      avatar,
		Identifiers: identifiers,
		ExtraUpdates: func(ctx context.Context, ghost *bridgev2.Ghost) (changed bool) {
			meta := ghost.Metadata.(*GhostMetadata)
			if !user.Min {
				changed = changed || meta.IsPremium != user.Premium || meta.IsBot != user.Bot
				meta.IsPremium = user.Premium
				meta.IsBot = user.Bot
				meta.Deleted = user.Deleted
			}
			return changed
		},
	}, nil
}
