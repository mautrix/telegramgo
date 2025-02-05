package connector

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/crypto"
	"github.com/gotd/td/tg"
	"github.com/rs/zerolog"
	"github.com/tidwall/gjson"
	"go.mau.fi/util/random"
	"maunium.net/go/mautrix/bridgev2"
	"maunium.net/go/mautrix/bridgev2/networkid"

	"go.mau.fi/mautrix-telegram/pkg/connector/ids"
)

var (
	_ bridgev2.PushableNetworkAPI = (*TelegramClient)(nil)
	_ bridgev2.PushParsingNetwork = (*TelegramConnector)(nil)
)

var PushAppSandbox = false

type PushCustomData struct {
	MessageID int `json:"msg_id,string"`

	ChannelID int64 `json:"channel_id,string"`
	ChatID    int64 `json:"chat_id,string"`
	FromID    int64 `json:"from_id,string"`

	ChatFromBroadcastID int64 `json:"chat_from_broadcast_id,string"`
	ChatFromGroupID     int64 `json:"chat_from_group_id,string"`
	ChatFromID          int64 `json:"chat_from_id,string"`
}

type PushNotificationData struct {
	LocKey  string         `json:"loc_key"`
	LocArgs []string       `json:"loc_args"`
	Custom  PushCustomData `json:"custom"`
	Sound   string         `json:"sound"`
	UserID  int            `json:"user_id,string"`
}

type APNSPayload struct {
	Aps       Aps   `json:"aps"`
	MsgID     int   `json:"msg_id,string"`
	FromID    int64 `json:"from_id,string"`
	ChannelID int64 `json:"channel_id,string"`

	ChatFromBroadcastID int64 `json:"chat_from_broadcast_id,string"`
	ChatFromGroupID     int64 `json:"chat_from_group_id,string"`
	ChatFromID          int64 `json:"chat_from_id,string"`
	ChatID              int64 `json:"chat_id,string"`

	LocKey *string `json:"loc-key"`
}

type Aps struct {
	Alert    Alert  `json:"alert"`
	ThreadID string `json:"thread-id"`
}

type Alert struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var FullSyncOnConnectBackground = false

func (t *TelegramClient) ConnectBackground(ctx context.Context, params *bridgev2.ConnectBackgroundParams) error {
	data, _ := params.ExtraData.(*APNSPayload)
	var relatedPortal *bridgev2.Portal
	var sender *bridgev2.Ghost
	var messageID networkid.MessageID
	var messageText, notificationText string
	if data != nil {
		if data.LocKey == nil {
			messageText = data.Aps.Alert.Body
			notificationText = data.Aps.Alert.Body
		}

		var err error
		if data.ChannelID != 0 {
			relatedPortal, err = t.main.Bridge.GetPortalByKey(ctx, t.makePortalKeyFromID(ids.PeerTypeChannel, data.ChannelID))
		} else if data.ChatID != 0 {
			relatedPortal, err = t.main.Bridge.GetPortalByKey(ctx, t.makePortalKeyFromID(ids.PeerTypeChat, data.ChatID))
		} else if data.FromID != 0 {
			relatedPortal, err = t.main.Bridge.GetPortalByKey(ctx, t.makePortalKeyFromID(ids.PeerTypeUser, data.FromID))
		}
		if err != nil {
			return fmt.Errorf("failed to get related portal: %w", err)
		}

		if data.ChatFromBroadcastID != 0 {
			sender, err = t.main.Bridge.GetGhostByID(ctx, ids.MakeChannelUserID(data.ChatFromBroadcastID))
		} else if data.ChatFromGroupID != 0 {
			sender, err = t.main.Bridge.GetGhostByID(ctx, ids.MakeChannelUserID(data.ChatFromGroupID))
		} else if data.FromID != 0 {
			sender, err = t.main.Bridge.GetGhostByID(ctx, ids.MakeUserID(data.FromID))
		}
		if err != nil {
			return fmt.Errorf("failed to get sender: %w", err)
		}
		if relatedPortal != nil && data.MsgID != 0 {
			messageID = ids.MakeMessageID(relatedPortal.PortalKey, data.MsgID)
		}
	}
	notifs, ok := t.main.Bridge.Matrix.(bridgev2.MatrixConnectorWithNotifications)
	if ok {
		notifs.DisplayNotification(ctx, &bridgev2.DirectNotificationData{
			Portal:    relatedPortal,
			Sender:    sender,
			Message:   messageText,
			MessageID: messageID,

			FormattedNotification: notificationText,
		})
	}
	if FullSyncOnConnectBackground {
		t.Connect(ctx)
		defer t.Disconnect()
		// TODO is it possible to safely only sync one chat?
		select {
		case <-time.After(20 * time.Second):
		case <-ctx.Done():
		}
	}
	return nil
}

func (tg *TelegramConnector) ParsePushNotification(ctx context.Context, data json.RawMessage) (networkid.UserLoginID, any, error) {
	val := gjson.GetBytes(data, "p")
	if val.Type != gjson.String {
		return "", nil, fmt.Errorf("missing or invalid p field")
	}
	valBytes, err := base64.RawURLEncoding.DecodeString(val.Str)
	if err != nil {
		return "", nil, fmt.Errorf("failed to base64 decode p field: %w", err)
	}
	var em crypto.EncryptedMessage
	err = em.DecodeWithoutCopy(&bin.Buffer{Buf: valBytes})
	if err != nil {
		return "", nil, fmt.Errorf("failed to decode auth key and message ID: %w", err)
	}
	userIDs, err := tg.Bridge.DB.UserLogin.GetAllUserIDsWithLogins(ctx)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get users with logins: %w", err)
	}
	var matchingAuthKey *crypto.AuthKey
	var userLoginID networkid.UserLoginID
UserLoop:
	for _, userID := range userIDs {
		user, err := tg.Bridge.GetExistingUserByMXID(ctx, userID)
		if err != nil {
			return "", nil, fmt.Errorf("failed to get user %s: %w", userID, err)
		}
		for _, login := range user.GetUserLogins() {
			key := login.Metadata.(*UserLoginMetadata).PushEncryptionKey
			if len(key) != 256 {
				continue
			}
			authKey := crypto.Key(key).WithID()
			if authKey.ID == em.AuthKeyID {
				matchingAuthKey = &authKey
				userLoginID = login.ID
				break UserLoop
			}
		}
	}
	if matchingAuthKey == nil {
		return "", nil, fmt.Errorf("no matching auth key found")
	}
	c := crypto.NewClientCipher(rand.Reader)
	plaintext, err := c.DecryptRaw(*matchingAuthKey, &em)
	if err != nil {
		return userLoginID, nil, fmt.Errorf("failed to decrypt payload: %w", err)
	} else if len(plaintext) < 4 {
		return userLoginID, nil, fmt.Errorf("decrypted payload too short (expected >4, got %d)", len(plaintext))
	}
	jsonLength := binary.LittleEndian.Uint32(plaintext[0:4])
	if len(plaintext) < int(jsonLength)+4 {
		return userLoginID, nil, fmt.Errorf("decrypted payload too short (expected 4+%d, got %d)", jsonLength, len(plaintext))
	}
	var pmd APNSPayload
	err = json.Unmarshal(plaintext[4:jsonLength+4], &pmd)
	if err != nil {
		zerolog.Ctx(ctx).Debug().Str("raw_data", base64.StdEncoding.EncodeToString(plaintext)).Msg("Decrypted non-JSON push data")
		return userLoginID, nil, fmt.Errorf("failed to unmarshal decrypted payload: %w", err)
	}
	return userLoginID, &pmd, nil
}

func (t *TelegramClient) RegisterPushNotifications(ctx context.Context, pushType bridgev2.PushType, token string) error {
	meta := t.userLogin.Metadata.(*UserLoginMetadata)
	if meta.PushEncryptionKey == nil {
		meta.PushEncryptionKey = random.Bytes(256)
		err := t.userLogin.Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to save push encryption key: %w", err)
		}
	}
	var tokenType int
	switch pushType {
	case bridgev2.PushTypeWeb:
		tokenType = 10
	case bridgev2.PushTypeFCM:
		tokenType = 2
	case bridgev2.PushTypeAPNs:
		tokenType = 1
	default:
		return fmt.Errorf("unsupported push type %s", pushType)
	}
	_, err := t.client.API().AccountRegisterDevice(ctx, &tg.AccountRegisterDeviceRequest{
		NoMuted:    false,
		TokenType:  tokenType,
		Token:      token,
		AppSandbox: PushAppSandbox,
		Secret:     meta.PushEncryptionKey,
		OtherUIDs:  nil, // TODO set properly
	})
	return err
}

func (t *TelegramClient) GetPushConfigs() *bridgev2.PushConfig {
	return &bridgev2.PushConfig{Native: true}
}
