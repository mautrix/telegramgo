// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// BotVerification represents TL type `botVerification#b160b6f2`.
type BotVerification struct {
	// Identifier of the bot that provided the verification
	BotUserID int64
	// Identifier of the custom emoji that is used as the verification sign
	IconCustomEmojiID int64
	// Custom description of verification reason set by the bot. Can contain only Mention,
	// Hashtag, Cashtag, PhoneNumber, BankCardNumber, Url, and EmailAddress entities
	CustomDescription FormattedText
}

// BotVerificationTypeID is TL type id of BotVerification.
const BotVerificationTypeID = 0xb160b6f2

// Ensuring interfaces in compile-time for BotVerification.
var (
	_ bin.Encoder     = &BotVerification{}
	_ bin.Decoder     = &BotVerification{}
	_ bin.BareEncoder = &BotVerification{}
	_ bin.BareDecoder = &BotVerification{}
)

func (b *BotVerification) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.BotUserID == 0) {
		return false
	}
	if !(b.IconCustomEmojiID == 0) {
		return false
	}
	if !(b.CustomDescription.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotVerification) String() string {
	if b == nil {
		return "BotVerification(nil)"
	}
	type Alias BotVerification
	return fmt.Sprintf("BotVerification%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotVerification) TypeID() uint32 {
	return BotVerificationTypeID
}

// TypeName returns name of type in TL schema.
func (*BotVerification) TypeName() string {
	return "botVerification"
}

// TypeInfo returns info about TL type.
func (b *BotVerification) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botVerification",
		ID:   BotVerificationTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "IconCustomEmojiID",
			SchemaName: "icon_custom_emoji_id",
		},
		{
			Name:       "CustomDescription",
			SchemaName: "custom_description",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotVerification) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botVerification#b160b6f2 as nil")
	}
	buf.PutID(BotVerificationTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotVerification) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botVerification#b160b6f2 as nil")
	}
	buf.PutInt53(b.BotUserID)
	buf.PutLong(b.IconCustomEmojiID)
	if err := b.CustomDescription.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botVerification#b160b6f2: field custom_description: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotVerification) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botVerification#b160b6f2 to nil")
	}
	if err := buf.ConsumeID(BotVerificationTypeID); err != nil {
		return fmt.Errorf("unable to decode botVerification#b160b6f2: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotVerification) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botVerification#b160b6f2 to nil")
	}
	{
		value, err := buf.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode botVerification#b160b6f2: field bot_user_id: %w", err)
		}
		b.BotUserID = value
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode botVerification#b160b6f2: field icon_custom_emoji_id: %w", err)
		}
		b.IconCustomEmojiID = value
	}
	{
		if err := b.CustomDescription.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botVerification#b160b6f2: field custom_description: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BotVerification) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode botVerification#b160b6f2 as nil")
	}
	buf.ObjStart()
	buf.PutID("botVerification")
	buf.Comma()
	buf.FieldStart("bot_user_id")
	buf.PutInt53(b.BotUserID)
	buf.Comma()
	buf.FieldStart("icon_custom_emoji_id")
	buf.PutLong(b.IconCustomEmojiID)
	buf.Comma()
	buf.FieldStart("custom_description")
	if err := b.CustomDescription.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode botVerification#b160b6f2: field custom_description: %w", err)
	}
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BotVerification) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode botVerification#b160b6f2 to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("botVerification"); err != nil {
				return fmt.Errorf("unable to decode botVerification#b160b6f2: %w", err)
			}
		case "bot_user_id":
			value, err := buf.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode botVerification#b160b6f2: field bot_user_id: %w", err)
			}
			b.BotUserID = value
		case "icon_custom_emoji_id":
			value, err := buf.Long()
			if err != nil {
				return fmt.Errorf("unable to decode botVerification#b160b6f2: field icon_custom_emoji_id: %w", err)
			}
			b.IconCustomEmojiID = value
		case "custom_description":
			if err := b.CustomDescription.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode botVerification#b160b6f2: field custom_description: %w", err)
			}
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetBotUserID returns value of BotUserID field.
func (b *BotVerification) GetBotUserID() (value int64) {
	if b == nil {
		return
	}
	return b.BotUserID
}

// GetIconCustomEmojiID returns value of IconCustomEmojiID field.
func (b *BotVerification) GetIconCustomEmojiID() (value int64) {
	if b == nil {
		return
	}
	return b.IconCustomEmojiID
}

// GetCustomDescription returns value of CustomDescription field.
func (b *BotVerification) GetCustomDescription() (value FormattedText) {
	if b == nil {
		return
	}
	return b.CustomDescription
}
