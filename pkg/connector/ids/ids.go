package ids

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gotd/td/tg"
	"maunium.net/go/mautrix/bridgev2/networkid"
)

func MakeUserID(userID int64) networkid.UserID {
	return networkid.UserID(strconv.FormatInt(userID, 10))
}

func ParseUserID(userID networkid.UserID) (int64, error) {
	return strconv.ParseInt(string(userID), 10, 64)
}

func ParseUserLoginID(userID networkid.UserLoginID) (int64, error) {
	return strconv.ParseInt(string(userID), 10, 64)
}

func MakeUserLoginID(userID int64) networkid.UserLoginID {
	return networkid.UserLoginID(strconv.FormatInt(userID, 10))
}

// TODO: add space ID
func MakeMessageID(messageID int) networkid.MessageID {
	return networkid.MessageID(strconv.Itoa(messageID))
}

func ParseMessageID(messageID networkid.MessageID) (int, error) {
	return strconv.Atoi(string(messageID))
}

type PeerType string

const (
	PeerTypeUser    PeerType = "user"
	PeerTypeChat    PeerType = "chat"
	PeerTypeChannel PeerType = "channel"
)

func PeerTypeFromByte(pt byte) (PeerType, error) {
	switch pt {
	case 0x01:
		return PeerTypeUser, nil
	case 0x02:
		return PeerTypeChat, nil
	case 0x03:
		return PeerTypeChannel, nil
	default:
		return "", fmt.Errorf("unknown peer type %d", pt)
	}
}

func (pt PeerType) AsByte() byte {
	switch pt {
	case PeerTypeUser:
		return 0x01
	case PeerTypeChat:
		return 0x02
	case PeerTypeChannel:
		return 0x03
	default:
		panic(fmt.Errorf("unknown peer type %s", pt))
	}
}

func (pt PeerType) AsPortalKey(chatID int64, receiver networkid.UserLoginID) networkid.PortalKey {
	portalKey := networkid.PortalKey{
		ID: networkid.PortalID(fmt.Sprintf("%s:%d", pt, chatID)),
	}
	if pt == PeerTypeUser || pt == PeerTypeChat {
		portalKey.Receiver = receiver
	}
	return portalKey
}

func MakePortalKey(peer tg.PeerClass, receiver networkid.UserLoginID) networkid.PortalKey {
	switch v := peer.(type) {
	case *tg.PeerUser:
		return networkid.PortalKey{
			ID:       networkid.PortalID(fmt.Sprintf("%s:%d", PeerTypeUser, v.UserID)),
			Receiver: receiver,
		}
	case *tg.PeerChat:
		return networkid.PortalKey{
			ID:       networkid.PortalID(fmt.Sprintf("%s:%d", PeerTypeChat, v.ChatID)),
			Receiver: receiver,
		}
	case *tg.PeerChannel:
		return networkid.PortalKey{ID: networkid.PortalID(fmt.Sprintf("%s:%d", PeerTypeChannel, v.ChannelID))}
	default:
		panic(fmt.Errorf("unknown peer class type %T", v))
	}
}

func ParsePortalID(portalID networkid.PortalID) (pt PeerType, id int64, err error) {
	parts := strings.Split(string(portalID), ":")
	pt = PeerType(parts[0])
	id, err = strconv.ParseInt(parts[1], 10, 64)
	return
}

func MakeAvatarID(photoID int64) networkid.AvatarID {
	return networkid.AvatarID(strconv.FormatInt(photoID, 10))
}

func MakeEmojiIDFromDocumentID(documentID int64) networkid.EmojiID {
	return networkid.EmojiID(fmt.Sprintf("d%d", documentID))
}

func MakeEmojiIDFromEmoticon(emoticon string) networkid.EmojiID {
	return networkid.EmojiID(fmt.Sprintf("e%s", emoticon))
}

func ParseEmojiID(emojiID networkid.EmojiID) (documentID int64, emoticon string, err error) {
	switch emojiID[0] {
	case 'd':
		documentID, err = strconv.ParseInt(string(emojiID[1:]), 10, 64)
	case 'e':
		emoticon = string(emojiID[1:])
	default:
		err = fmt.Errorf("invalid emoji ID type %s", string(emojiID[0]))
	}
	return
}
