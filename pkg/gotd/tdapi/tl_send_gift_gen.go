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

// SendGiftRequest represents TL type `sendGift#b883472a`.
type SendGiftRequest struct {
	// Identifier of the gift to send
	GiftID int64
	// Identifier of the user or the channel chat that will receive the gift
	OwnerID MessageSenderClass
	// Text to show along with the gift; 0-getOption("gift_text_length_max") characters. Only
	// Bold, Italic, Underline, Strikethrough, Spoiler, and CustomEmoji entities are allowed.
	Text FormattedText
	// Pass true to show gift text and sender only to the gift receiver; otherwise, everyone
	// will be able to see them
	IsPrivate bool
	// Pass true to additionally pay for the gift upgrade and allow the receiver to upgrade
	// it for free
	PayForUpgrade bool
}

// SendGiftRequestTypeID is TL type id of SendGiftRequest.
const SendGiftRequestTypeID = 0xb883472a

// Ensuring interfaces in compile-time for SendGiftRequest.
var (
	_ bin.Encoder     = &SendGiftRequest{}
	_ bin.Decoder     = &SendGiftRequest{}
	_ bin.BareEncoder = &SendGiftRequest{}
	_ bin.BareDecoder = &SendGiftRequest{}
)

func (s *SendGiftRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.GiftID == 0) {
		return false
	}
	if !(s.OwnerID == nil) {
		return false
	}
	if !(s.Text.Zero()) {
		return false
	}
	if !(s.IsPrivate == false) {
		return false
	}
	if !(s.PayForUpgrade == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendGiftRequest) String() string {
	if s == nil {
		return "SendGiftRequest(nil)"
	}
	type Alias SendGiftRequest
	return fmt.Sprintf("SendGiftRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendGiftRequest) TypeID() uint32 {
	return SendGiftRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendGiftRequest) TypeName() string {
	return "sendGift"
}

// TypeInfo returns info about TL type.
func (s *SendGiftRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendGift",
		ID:   SendGiftRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GiftID",
			SchemaName: "gift_id",
		},
		{
			Name:       "OwnerID",
			SchemaName: "owner_id",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "IsPrivate",
			SchemaName: "is_private",
		},
		{
			Name:       "PayForUpgrade",
			SchemaName: "pay_for_upgrade",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendGiftRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendGift#b883472a as nil")
	}
	b.PutID(SendGiftRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendGiftRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendGift#b883472a as nil")
	}
	b.PutLong(s.GiftID)
	if s.OwnerID == nil {
		return fmt.Errorf("unable to encode sendGift#b883472a: field owner_id is nil")
	}
	if err := s.OwnerID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendGift#b883472a: field owner_id: %w", err)
	}
	if err := s.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendGift#b883472a: field text: %w", err)
	}
	b.PutBool(s.IsPrivate)
	b.PutBool(s.PayForUpgrade)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendGiftRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendGift#b883472a to nil")
	}
	if err := b.ConsumeID(SendGiftRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendGift#b883472a: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendGiftRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendGift#b883472a to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode sendGift#b883472a: field gift_id: %w", err)
		}
		s.GiftID = value
	}
	{
		value, err := DecodeMessageSender(b)
		if err != nil {
			return fmt.Errorf("unable to decode sendGift#b883472a: field owner_id: %w", err)
		}
		s.OwnerID = value
	}
	{
		if err := s.Text.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sendGift#b883472a: field text: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode sendGift#b883472a: field is_private: %w", err)
		}
		s.IsPrivate = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode sendGift#b883472a: field pay_for_upgrade: %w", err)
		}
		s.PayForUpgrade = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SendGiftRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sendGift#b883472a as nil")
	}
	b.ObjStart()
	b.PutID("sendGift")
	b.Comma()
	b.FieldStart("gift_id")
	b.PutLong(s.GiftID)
	b.Comma()
	b.FieldStart("owner_id")
	if s.OwnerID == nil {
		return fmt.Errorf("unable to encode sendGift#b883472a: field owner_id is nil")
	}
	if err := s.OwnerID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendGift#b883472a: field owner_id: %w", err)
	}
	b.Comma()
	b.FieldStart("text")
	if err := s.Text.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendGift#b883472a: field text: %w", err)
	}
	b.Comma()
	b.FieldStart("is_private")
	b.PutBool(s.IsPrivate)
	b.Comma()
	b.FieldStart("pay_for_upgrade")
	b.PutBool(s.PayForUpgrade)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SendGiftRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sendGift#b883472a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sendGift"); err != nil {
				return fmt.Errorf("unable to decode sendGift#b883472a: %w", err)
			}
		case "gift_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode sendGift#b883472a: field gift_id: %w", err)
			}
			s.GiftID = value
		case "owner_id":
			value, err := DecodeTDLibJSONMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode sendGift#b883472a: field owner_id: %w", err)
			}
			s.OwnerID = value
		case "text":
			if err := s.Text.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sendGift#b883472a: field text: %w", err)
			}
		case "is_private":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode sendGift#b883472a: field is_private: %w", err)
			}
			s.IsPrivate = value
		case "pay_for_upgrade":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode sendGift#b883472a: field pay_for_upgrade: %w", err)
			}
			s.PayForUpgrade = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGiftID returns value of GiftID field.
func (s *SendGiftRequest) GetGiftID() (value int64) {
	if s == nil {
		return
	}
	return s.GiftID
}

// GetOwnerID returns value of OwnerID field.
func (s *SendGiftRequest) GetOwnerID() (value MessageSenderClass) {
	if s == nil {
		return
	}
	return s.OwnerID
}

// GetText returns value of Text field.
func (s *SendGiftRequest) GetText() (value FormattedText) {
	if s == nil {
		return
	}
	return s.Text
}

// GetIsPrivate returns value of IsPrivate field.
func (s *SendGiftRequest) GetIsPrivate() (value bool) {
	if s == nil {
		return
	}
	return s.IsPrivate
}

// GetPayForUpgrade returns value of PayForUpgrade field.
func (s *SendGiftRequest) GetPayForUpgrade() (value bool) {
	if s == nil {
		return
	}
	return s.PayForUpgrade
}

// SendGift invokes method sendGift#b883472a returning error if any.
func (c *Client) SendGift(ctx context.Context, request *SendGiftRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
