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

// BotsGetBotMenuButtonRequest represents TL type `bots.getBotMenuButton#9c60eb28`.
// Gets the menu button action for a given user or for all users, previously set using
// bots.setBotMenuButton¹; users can see this information in the botInfo² constructor.
//
// Links:
//  1. https://core.telegram.org/method/bots.setBotMenuButton
//  2. https://core.telegram.org/constructor/botInfo
//
// See https://core.telegram.org/method/bots.getBotMenuButton for reference.
type BotsGetBotMenuButtonRequest struct {
	// User ID or empty for the default menu button.
	UserID InputUserClass
}

// BotsGetBotMenuButtonRequestTypeID is TL type id of BotsGetBotMenuButtonRequest.
const BotsGetBotMenuButtonRequestTypeID = 0x9c60eb28

// Ensuring interfaces in compile-time for BotsGetBotMenuButtonRequest.
var (
	_ bin.Encoder     = &BotsGetBotMenuButtonRequest{}
	_ bin.Decoder     = &BotsGetBotMenuButtonRequest{}
	_ bin.BareEncoder = &BotsGetBotMenuButtonRequest{}
	_ bin.BareDecoder = &BotsGetBotMenuButtonRequest{}
)

func (g *BotsGetBotMenuButtonRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.UserID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *BotsGetBotMenuButtonRequest) String() string {
	if g == nil {
		return "BotsGetBotMenuButtonRequest(nil)"
	}
	type Alias BotsGetBotMenuButtonRequest
	return fmt.Sprintf("BotsGetBotMenuButtonRequest%+v", Alias(*g))
}

// FillFrom fills BotsGetBotMenuButtonRequest from given interface.
func (g *BotsGetBotMenuButtonRequest) FillFrom(from interface {
	GetUserID() (value InputUserClass)
}) {
	g.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsGetBotMenuButtonRequest) TypeID() uint32 {
	return BotsGetBotMenuButtonRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsGetBotMenuButtonRequest) TypeName() string {
	return "bots.getBotMenuButton"
}

// TypeInfo returns info about TL type.
func (g *BotsGetBotMenuButtonRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.getBotMenuButton",
		ID:   BotsGetBotMenuButtonRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *BotsGetBotMenuButtonRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode bots.getBotMenuButton#9c60eb28 as nil")
	}
	b.PutID(BotsGetBotMenuButtonRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *BotsGetBotMenuButtonRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode bots.getBotMenuButton#9c60eb28 as nil")
	}
	if g.UserID == nil {
		return fmt.Errorf("unable to encode bots.getBotMenuButton#9c60eb28: field user_id is nil")
	}
	if err := g.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.getBotMenuButton#9c60eb28: field user_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *BotsGetBotMenuButtonRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode bots.getBotMenuButton#9c60eb28 to nil")
	}
	if err := b.ConsumeID(BotsGetBotMenuButtonRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.getBotMenuButton#9c60eb28: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *BotsGetBotMenuButtonRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode bots.getBotMenuButton#9c60eb28 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.getBotMenuButton#9c60eb28: field user_id: %w", err)
		}
		g.UserID = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (g *BotsGetBotMenuButtonRequest) GetUserID() (value InputUserClass) {
	if g == nil {
		return
	}
	return g.UserID
}

// BotsGetBotMenuButton invokes method bots.getBotMenuButton#9c60eb28 returning error if any.
// Gets the menu button action for a given user or for all users, previously set using
// bots.setBotMenuButton¹; users can see this information in the botInfo² constructor.
//
// Links:
//  1. https://core.telegram.org/method/bots.setBotMenuButton
//  2. https://core.telegram.org/constructor/botInfo
//
// Possible errors:
//
//	400 USER_BOT_REQUIRED: This method can only be called by a bot.
//
// See https://core.telegram.org/method/bots.getBotMenuButton for reference.
// Can be used by bots.
func (c *Client) BotsGetBotMenuButton(ctx context.Context, userid InputUserClass) (BotMenuButtonClass, error) {
	var result BotMenuButtonBox

	request := &BotsGetBotMenuButtonRequest{
		UserID: userid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.BotMenuButton, nil
}
