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

// EditBusinessMessageReplyMarkupRequest represents TL type `editBusinessMessageReplyMarkup#29d46c17`.
type EditBusinessMessageReplyMarkupRequest struct {
	// Unique identifier of business connection on behalf of which the message was sent
	BusinessConnectionID string
	// The chat the message belongs to
	ChatID int64
	// Identifier of the message
	MessageID int64
	// The new message reply markup; pass null if none
	ReplyMarkup ReplyMarkupClass
}

// EditBusinessMessageReplyMarkupRequestTypeID is TL type id of EditBusinessMessageReplyMarkupRequest.
const EditBusinessMessageReplyMarkupRequestTypeID = 0x29d46c17

// Ensuring interfaces in compile-time for EditBusinessMessageReplyMarkupRequest.
var (
	_ bin.Encoder     = &EditBusinessMessageReplyMarkupRequest{}
	_ bin.Decoder     = &EditBusinessMessageReplyMarkupRequest{}
	_ bin.BareEncoder = &EditBusinessMessageReplyMarkupRequest{}
	_ bin.BareDecoder = &EditBusinessMessageReplyMarkupRequest{}
)

func (e *EditBusinessMessageReplyMarkupRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.BusinessConnectionID == "") {
		return false
	}
	if !(e.ChatID == 0) {
		return false
	}
	if !(e.MessageID == 0) {
		return false
	}
	if !(e.ReplyMarkup == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditBusinessMessageReplyMarkupRequest) String() string {
	if e == nil {
		return "EditBusinessMessageReplyMarkupRequest(nil)"
	}
	type Alias EditBusinessMessageReplyMarkupRequest
	return fmt.Sprintf("EditBusinessMessageReplyMarkupRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditBusinessMessageReplyMarkupRequest) TypeID() uint32 {
	return EditBusinessMessageReplyMarkupRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditBusinessMessageReplyMarkupRequest) TypeName() string {
	return "editBusinessMessageReplyMarkup"
}

// TypeInfo returns info about TL type.
func (e *EditBusinessMessageReplyMarkupRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editBusinessMessageReplyMarkup",
		ID:   EditBusinessMessageReplyMarkupRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BusinessConnectionID",
			SchemaName: "business_connection_id",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditBusinessMessageReplyMarkupRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editBusinessMessageReplyMarkup#29d46c17 as nil")
	}
	b.PutID(EditBusinessMessageReplyMarkupRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditBusinessMessageReplyMarkupRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editBusinessMessageReplyMarkup#29d46c17 as nil")
	}
	b.PutString(e.BusinessConnectionID)
	b.PutInt53(e.ChatID)
	b.PutInt53(e.MessageID)
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editBusinessMessageReplyMarkup#29d46c17: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editBusinessMessageReplyMarkup#29d46c17: field reply_markup: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EditBusinessMessageReplyMarkupRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editBusinessMessageReplyMarkup#29d46c17 to nil")
	}
	if err := b.ConsumeID(EditBusinessMessageReplyMarkupRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editBusinessMessageReplyMarkup#29d46c17: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditBusinessMessageReplyMarkupRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editBusinessMessageReplyMarkup#29d46c17 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode editBusinessMessageReplyMarkup#29d46c17: field business_connection_id: %w", err)
		}
		e.BusinessConnectionID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode editBusinessMessageReplyMarkup#29d46c17: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode editBusinessMessageReplyMarkup#29d46c17: field message_id: %w", err)
		}
		e.MessageID = value
	}
	{
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode editBusinessMessageReplyMarkup#29d46c17: field reply_markup: %w", err)
		}
		e.ReplyMarkup = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditBusinessMessageReplyMarkupRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editBusinessMessageReplyMarkup#29d46c17 as nil")
	}
	b.ObjStart()
	b.PutID("editBusinessMessageReplyMarkup")
	b.Comma()
	b.FieldStart("business_connection_id")
	b.PutString(e.BusinessConnectionID)
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(e.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(e.MessageID)
	b.Comma()
	b.FieldStart("reply_markup")
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editBusinessMessageReplyMarkup#29d46c17: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editBusinessMessageReplyMarkup#29d46c17: field reply_markup: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditBusinessMessageReplyMarkupRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editBusinessMessageReplyMarkup#29d46c17 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editBusinessMessageReplyMarkup"); err != nil {
				return fmt.Errorf("unable to decode editBusinessMessageReplyMarkup#29d46c17: %w", err)
			}
		case "business_connection_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode editBusinessMessageReplyMarkup#29d46c17: field business_connection_id: %w", err)
			}
			e.BusinessConnectionID = value
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode editBusinessMessageReplyMarkup#29d46c17: field chat_id: %w", err)
			}
			e.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode editBusinessMessageReplyMarkup#29d46c17: field message_id: %w", err)
			}
			e.MessageID = value
		case "reply_markup":
			value, err := DecodeTDLibJSONReplyMarkup(b)
			if err != nil {
				return fmt.Errorf("unable to decode editBusinessMessageReplyMarkup#29d46c17: field reply_markup: %w", err)
			}
			e.ReplyMarkup = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBusinessConnectionID returns value of BusinessConnectionID field.
func (e *EditBusinessMessageReplyMarkupRequest) GetBusinessConnectionID() (value string) {
	if e == nil {
		return
	}
	return e.BusinessConnectionID
}

// GetChatID returns value of ChatID field.
func (e *EditBusinessMessageReplyMarkupRequest) GetChatID() (value int64) {
	if e == nil {
		return
	}
	return e.ChatID
}

// GetMessageID returns value of MessageID field.
func (e *EditBusinessMessageReplyMarkupRequest) GetMessageID() (value int64) {
	if e == nil {
		return
	}
	return e.MessageID
}

// GetReplyMarkup returns value of ReplyMarkup field.
func (e *EditBusinessMessageReplyMarkupRequest) GetReplyMarkup() (value ReplyMarkupClass) {
	if e == nil {
		return
	}
	return e.ReplyMarkup
}

// EditBusinessMessageReplyMarkup invokes method editBusinessMessageReplyMarkup#29d46c17 returning error if any.
func (c *Client) EditBusinessMessageReplyMarkup(ctx context.Context, request *EditBusinessMessageReplyMarkupRequest) (*BusinessMessage, error) {
	var result BusinessMessage

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
