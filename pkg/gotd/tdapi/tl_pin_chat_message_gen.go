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

// PinChatMessageRequest represents TL type `pinChatMessage#79475baf`.
type PinChatMessageRequest struct {
	// Identifier of the chat
	ChatID int64
	// Identifier of the new pinned message
	MessageID int64
	// Pass true to disable notification about the pinned message. Notifications are always
	// disabled in channels and private chats
	DisableNotification bool
	// Pass true to pin the message only for self; private chats only
	OnlyForSelf bool
}

// PinChatMessageRequestTypeID is TL type id of PinChatMessageRequest.
const PinChatMessageRequestTypeID = 0x79475baf

// Ensuring interfaces in compile-time for PinChatMessageRequest.
var (
	_ bin.Encoder     = &PinChatMessageRequest{}
	_ bin.Decoder     = &PinChatMessageRequest{}
	_ bin.BareEncoder = &PinChatMessageRequest{}
	_ bin.BareDecoder = &PinChatMessageRequest{}
)

func (p *PinChatMessageRequest) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ChatID == 0) {
		return false
	}
	if !(p.MessageID == 0) {
		return false
	}
	if !(p.DisableNotification == false) {
		return false
	}
	if !(p.OnlyForSelf == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PinChatMessageRequest) String() string {
	if p == nil {
		return "PinChatMessageRequest(nil)"
	}
	type Alias PinChatMessageRequest
	return fmt.Sprintf("PinChatMessageRequest%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PinChatMessageRequest) TypeID() uint32 {
	return PinChatMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PinChatMessageRequest) TypeName() string {
	return "pinChatMessage"
}

// TypeInfo returns info about TL type.
func (p *PinChatMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pinChatMessage",
		ID:   PinChatMessageRequestTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "DisableNotification",
			SchemaName: "disable_notification",
		},
		{
			Name:       "OnlyForSelf",
			SchemaName: "only_for_self",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PinChatMessageRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pinChatMessage#79475baf as nil")
	}
	b.PutID(PinChatMessageRequestTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PinChatMessageRequest) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pinChatMessage#79475baf as nil")
	}
	b.PutInt53(p.ChatID)
	b.PutInt53(p.MessageID)
	b.PutBool(p.DisableNotification)
	b.PutBool(p.OnlyForSelf)
	return nil
}

// Decode implements bin.Decoder.
func (p *PinChatMessageRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pinChatMessage#79475baf to nil")
	}
	if err := b.ConsumeID(PinChatMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode pinChatMessage#79475baf: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PinChatMessageRequest) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pinChatMessage#79475baf to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode pinChatMessage#79475baf: field chat_id: %w", err)
		}
		p.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode pinChatMessage#79475baf: field message_id: %w", err)
		}
		p.MessageID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode pinChatMessage#79475baf: field disable_notification: %w", err)
		}
		p.DisableNotification = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode pinChatMessage#79475baf: field only_for_self: %w", err)
		}
		p.OnlyForSelf = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PinChatMessageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode pinChatMessage#79475baf as nil")
	}
	b.ObjStart()
	b.PutID("pinChatMessage")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(p.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(p.MessageID)
	b.Comma()
	b.FieldStart("disable_notification")
	b.PutBool(p.DisableNotification)
	b.Comma()
	b.FieldStart("only_for_self")
	b.PutBool(p.OnlyForSelf)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PinChatMessageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode pinChatMessage#79475baf to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("pinChatMessage"); err != nil {
				return fmt.Errorf("unable to decode pinChatMessage#79475baf: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode pinChatMessage#79475baf: field chat_id: %w", err)
			}
			p.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode pinChatMessage#79475baf: field message_id: %w", err)
			}
			p.MessageID = value
		case "disable_notification":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode pinChatMessage#79475baf: field disable_notification: %w", err)
			}
			p.DisableNotification = value
		case "only_for_self":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode pinChatMessage#79475baf: field only_for_self: %w", err)
			}
			p.OnlyForSelf = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (p *PinChatMessageRequest) GetChatID() (value int64) {
	if p == nil {
		return
	}
	return p.ChatID
}

// GetMessageID returns value of MessageID field.
func (p *PinChatMessageRequest) GetMessageID() (value int64) {
	if p == nil {
		return
	}
	return p.MessageID
}

// GetDisableNotification returns value of DisableNotification field.
func (p *PinChatMessageRequest) GetDisableNotification() (value bool) {
	if p == nil {
		return
	}
	return p.DisableNotification
}

// GetOnlyForSelf returns value of OnlyForSelf field.
func (p *PinChatMessageRequest) GetOnlyForSelf() (value bool) {
	if p == nil {
		return
	}
	return p.OnlyForSelf
}

// PinChatMessage invokes method pinChatMessage#79475baf returning error if any.
func (c *Client) PinChatMessage(ctx context.Context, request *PinChatMessageRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
