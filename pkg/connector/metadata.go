package connector

import (
	"context"

	"github.com/gotd/td/crypto"
	"github.com/gotd/td/session"
	"maunium.net/go/mautrix/bridgev2/database"
	"maunium.net/go/mautrix/bridgev2/networkid"
	"maunium.net/go/mautrix/id"
)

func (tg *TelegramConnector) GetDBMetaTypes() database.MetaTypes {
	return database.MetaTypes{
		Ghost:     func() any { return &GhostMetadata{} },
		Portal:    func() any { return &PortalMetadata{} },
		Message:   func() any { return &MessageMetadata{} },
		Reaction:  nil,
		UserLogin: func() any { return &UserLoginMetadata{} },
	}
}

type GhostMetadata struct {
	IsPremium bool `json:"is_premium,omitempty"`
	IsBot     bool `json:"is_bot,omitempty"`
	IsChannel bool `json:"is_channel,omitempty"`
	Blocked   bool `json:"blocked,omitempty"`
	Deleted   bool `json:"deleted,omitempty"`
}

type PortalMetadata struct {
	IsSuperGroup bool `json:"is_supergroup,omitempty"`
	ReadUpTo     int  `json:"read_up_to,omitempty"`
	MessagesTTL  int  `json:"messages_ttl,omitempty"`
}

type MessageMetadata struct {
	ContentHash []byte              `json:"content_hash,omitempty"`
	ContentURI  id.ContentURIString `json:"content_uri,omitempty"`
}

type UserLoginMetadata struct {
	Phone     string           `json:"phone"`
	Session   UserLoginSession `json:"session"`
	TakeoutID int64            `json:"takeout_id,omitempty"`

	TakeoutDialogCrawlDone   bool               `json:"takeout_portal_crawl_done,omitempty"`
	TakeoutDialogCrawlCursor networkid.PortalID `json:"takeout_portal_crawl_cursor,omitempty"`

	PinnedDialogs []networkid.PortalID `json:"pinned_dialogs,omitempty"`

	PushEncryptionKey []byte `json:"push_encryption_key,omitempty"`
}

func (u *UserLoginMetadata) ResetOnLogout() {
	u.Session.AuthKey = nil
	u.TakeoutID = 0
	u.TakeoutDialogCrawlDone = false
	u.TakeoutDialogCrawlCursor = networkid.PortalID("")
	u.PushEncryptionKey = nil
}

type UserLoginSession struct {
	AuthKey       []byte `json:"auth_key,omitempty"`
	Datacenter    int    `json:"dc_id,omitempty"`
	ServerAddress string `json:"server_address,omitempty"`
	ServerPort    int    `json:"port,omitempty"`
	Salt          int64  `json:"salt,omitempty"`
}

func (u UserLoginSession) HasAuthKey() bool {
	return len(u.AuthKey) == 256
}

func (s *UserLoginSession) Load(_ context.Context) (*session.Data, error) {
	if !s.HasAuthKey() {
		return nil, session.ErrNotFound
	}
	keyID := crypto.Key(s.AuthKey).ID()
	return &session.Data{
		DC:        s.Datacenter,
		Addr:      s.ServerAddress,
		AuthKey:   s.AuthKey,
		AuthKeyID: keyID[:],
		Salt:      s.Salt,
	}, nil
}

func (s *UserLoginSession) Save(ctx context.Context, data *session.Data) error {
	s.Datacenter = data.DC
	s.ServerAddress = data.Addr
	s.AuthKey = data.AuthKey
	s.Salt = data.Salt
	// TODO save UserLogin to database?
	return nil
}
