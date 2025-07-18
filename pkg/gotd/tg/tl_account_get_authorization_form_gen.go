// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"go.mau.fi/mautrix-telegram/pkg/gotd/bin"
	"go.mau.fi/mautrix-telegram/pkg/gotd/tdjson"
	"go.mau.fi/mautrix-telegram/pkg/gotd/tdp"
	"go.mau.fi/mautrix-telegram/pkg/gotd/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// AccountGetAuthorizationFormRequest represents TL type `account.getAuthorizationForm#a929597a`.
// Returns a Telegram Passport authorization form for sharing data with a service
//
// See https://core.telegram.org/method/account.getAuthorizationForm for reference.
type AccountGetAuthorizationFormRequest struct {
	// User identifier of the service's bot
	BotID int64
	// Telegram Passport element types requested by the service
	Scope string
	// Service's public key
	PublicKey string
}

// AccountGetAuthorizationFormRequestTypeID is TL type id of AccountGetAuthorizationFormRequest.
const AccountGetAuthorizationFormRequestTypeID = 0xa929597a

// Ensuring interfaces in compile-time for AccountGetAuthorizationFormRequest.
var (
	_ bin.Encoder     = &AccountGetAuthorizationFormRequest{}
	_ bin.Decoder     = &AccountGetAuthorizationFormRequest{}
	_ bin.BareEncoder = &AccountGetAuthorizationFormRequest{}
	_ bin.BareDecoder = &AccountGetAuthorizationFormRequest{}
)

func (g *AccountGetAuthorizationFormRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.BotID == 0) {
		return false
	}
	if !(g.Scope == "") {
		return false
	}
	if !(g.PublicKey == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetAuthorizationFormRequest) String() string {
	if g == nil {
		return "AccountGetAuthorizationFormRequest(nil)"
	}
	type Alias AccountGetAuthorizationFormRequest
	return fmt.Sprintf("AccountGetAuthorizationFormRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetAuthorizationFormRequest from given interface.
func (g *AccountGetAuthorizationFormRequest) FillFrom(from interface {
	GetBotID() (value int64)
	GetScope() (value string)
	GetPublicKey() (value string)
}) {
	g.BotID = from.GetBotID()
	g.Scope = from.GetScope()
	g.PublicKey = from.GetPublicKey()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetAuthorizationFormRequest) TypeID() uint32 {
	return AccountGetAuthorizationFormRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetAuthorizationFormRequest) TypeName() string {
	return "account.getAuthorizationForm"
}

// TypeInfo returns info about TL type.
func (g *AccountGetAuthorizationFormRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getAuthorizationForm",
		ID:   AccountGetAuthorizationFormRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotID",
			SchemaName: "bot_id",
		},
		{
			Name:       "Scope",
			SchemaName: "scope",
		},
		{
			Name:       "PublicKey",
			SchemaName: "public_key",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetAuthorizationFormRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getAuthorizationForm#a929597a as nil")
	}
	b.PutID(AccountGetAuthorizationFormRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetAuthorizationFormRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getAuthorizationForm#a929597a as nil")
	}
	b.PutLong(g.BotID)
	b.PutString(g.Scope)
	b.PutString(g.PublicKey)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetAuthorizationFormRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getAuthorizationForm#a929597a to nil")
	}
	if err := b.ConsumeID(AccountGetAuthorizationFormRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getAuthorizationForm#a929597a: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetAuthorizationFormRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getAuthorizationForm#a929597a to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.getAuthorizationForm#a929597a: field bot_id: %w", err)
		}
		g.BotID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.getAuthorizationForm#a929597a: field scope: %w", err)
		}
		g.Scope = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.getAuthorizationForm#a929597a: field public_key: %w", err)
		}
		g.PublicKey = value
	}
	return nil
}

// GetBotID returns value of BotID field.
func (g *AccountGetAuthorizationFormRequest) GetBotID() (value int64) {
	if g == nil {
		return
	}
	return g.BotID
}

// GetScope returns value of Scope field.
func (g *AccountGetAuthorizationFormRequest) GetScope() (value string) {
	if g == nil {
		return
	}
	return g.Scope
}

// GetPublicKey returns value of PublicKey field.
func (g *AccountGetAuthorizationFormRequest) GetPublicKey() (value string) {
	if g == nil {
		return
	}
	return g.PublicKey
}

// AccountGetAuthorizationForm invokes method account.getAuthorizationForm#a929597a returning error if any.
// Returns a Telegram Passport authorization form for sharing data with a service
//
// Possible errors:
//
//	400 BOT_INVALID: This is not a valid bot.
//	400 PUBLIC_KEY_REQUIRED: A public key is required.
//
// See https://core.telegram.org/method/account.getAuthorizationForm for reference.
func (c *Client) AccountGetAuthorizationForm(ctx context.Context, request *AccountGetAuthorizationFormRequest) (*AccountAuthorizationForm, error) {
	var result AccountAuthorizationForm

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
