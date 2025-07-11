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

// BotsGetBotInfoRequest represents TL type `bots.getBotInfo#dcd914fd`.
// Get localized name, about text and description of a bot (or of the current account, if
// called by a bot).
//
// See https://core.telegram.org/method/bots.getBotInfo for reference.
type BotsGetBotInfoRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If called by a user, must contain the peer of a bot we own.
	//
	// Use SetBot and GetBot helpers.
	Bot InputUserClass
	// Language code, if left empty this method will return the fallback about text and
	// description.
	LangCode string
}

// BotsGetBotInfoRequestTypeID is TL type id of BotsGetBotInfoRequest.
const BotsGetBotInfoRequestTypeID = 0xdcd914fd

// Ensuring interfaces in compile-time for BotsGetBotInfoRequest.
var (
	_ bin.Encoder     = &BotsGetBotInfoRequest{}
	_ bin.Decoder     = &BotsGetBotInfoRequest{}
	_ bin.BareEncoder = &BotsGetBotInfoRequest{}
	_ bin.BareDecoder = &BotsGetBotInfoRequest{}
)

func (g *BotsGetBotInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Bot == nil) {
		return false
	}
	if !(g.LangCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *BotsGetBotInfoRequest) String() string {
	if g == nil {
		return "BotsGetBotInfoRequest(nil)"
	}
	type Alias BotsGetBotInfoRequest
	return fmt.Sprintf("BotsGetBotInfoRequest%+v", Alias(*g))
}

// FillFrom fills BotsGetBotInfoRequest from given interface.
func (g *BotsGetBotInfoRequest) FillFrom(from interface {
	GetBot() (value InputUserClass, ok bool)
	GetLangCode() (value string)
}) {
	if val, ok := from.GetBot(); ok {
		g.Bot = val
	}

	g.LangCode = from.GetLangCode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsGetBotInfoRequest) TypeID() uint32 {
	return BotsGetBotInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsGetBotInfoRequest) TypeName() string {
	return "bots.getBotInfo"
}

// TypeInfo returns info about TL type.
func (g *BotsGetBotInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.getBotInfo",
		ID:   BotsGetBotInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *BotsGetBotInfoRequest) SetFlags() {
	if !(g.Bot == nil) {
		g.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (g *BotsGetBotInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode bots.getBotInfo#dcd914fd as nil")
	}
	b.PutID(BotsGetBotInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *BotsGetBotInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode bots.getBotInfo#dcd914fd as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.getBotInfo#dcd914fd: field flags: %w", err)
	}
	if g.Flags.Has(0) {
		if g.Bot == nil {
			return fmt.Errorf("unable to encode bots.getBotInfo#dcd914fd: field bot is nil")
		}
		if err := g.Bot.Encode(b); err != nil {
			return fmt.Errorf("unable to encode bots.getBotInfo#dcd914fd: field bot: %w", err)
		}
	}
	b.PutString(g.LangCode)
	return nil
}

// Decode implements bin.Decoder.
func (g *BotsGetBotInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode bots.getBotInfo#dcd914fd to nil")
	}
	if err := b.ConsumeID(BotsGetBotInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.getBotInfo#dcd914fd: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *BotsGetBotInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode bots.getBotInfo#dcd914fd to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode bots.getBotInfo#dcd914fd: field flags: %w", err)
		}
	}
	if g.Flags.Has(0) {
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.getBotInfo#dcd914fd: field bot: %w", err)
		}
		g.Bot = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.getBotInfo#dcd914fd: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	return nil
}

// SetBot sets value of Bot conditional field.
func (g *BotsGetBotInfoRequest) SetBot(value InputUserClass) {
	g.Flags.Set(0)
	g.Bot = value
}

// GetBot returns value of Bot conditional field and
// boolean which is true if field was set.
func (g *BotsGetBotInfoRequest) GetBot() (value InputUserClass, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.Bot, true
}

// GetLangCode returns value of LangCode field.
func (g *BotsGetBotInfoRequest) GetLangCode() (value string) {
	if g == nil {
		return
	}
	return g.LangCode
}

// BotsGetBotInfo invokes method bots.getBotInfo#dcd914fd returning error if any.
// Get localized name, about text and description of a bot (or of the current account, if
// called by a bot).
//
// Possible errors:
//
//	400 BOT_INVALID: This is not a valid bot.
//	400 LANG_CODE_INVALID: The specified language code is invalid.
//	400 USER_BOT_INVALID: User accounts must provide the bot method parameter when calling this method. If there is no such method parameter, this method can only be invoked by bot accounts.
//
// See https://core.telegram.org/method/bots.getBotInfo for reference.
// Can be used by bots.
func (c *Client) BotsGetBotInfo(ctx context.Context, request *BotsGetBotInfoRequest) (*BotsBotInfo, error) {
	var result BotsBotInfo

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
