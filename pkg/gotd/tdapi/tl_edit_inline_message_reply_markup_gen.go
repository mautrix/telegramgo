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

// EditInlineMessageReplyMarkupRequest represents TL type `editInlineMessageReplyMarkup#fbf906de`.
type EditInlineMessageReplyMarkupRequest struct {
	// Inline message identifier
	InlineMessageID string
	// The new message reply markup; pass null if none
	ReplyMarkup ReplyMarkupClass
}

// EditInlineMessageReplyMarkupRequestTypeID is TL type id of EditInlineMessageReplyMarkupRequest.
const EditInlineMessageReplyMarkupRequestTypeID = 0xfbf906de

// Ensuring interfaces in compile-time for EditInlineMessageReplyMarkupRequest.
var (
	_ bin.Encoder     = &EditInlineMessageReplyMarkupRequest{}
	_ bin.Decoder     = &EditInlineMessageReplyMarkupRequest{}
	_ bin.BareEncoder = &EditInlineMessageReplyMarkupRequest{}
	_ bin.BareDecoder = &EditInlineMessageReplyMarkupRequest{}
)

func (e *EditInlineMessageReplyMarkupRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.InlineMessageID == "") {
		return false
	}
	if !(e.ReplyMarkup == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditInlineMessageReplyMarkupRequest) String() string {
	if e == nil {
		return "EditInlineMessageReplyMarkupRequest(nil)"
	}
	type Alias EditInlineMessageReplyMarkupRequest
	return fmt.Sprintf("EditInlineMessageReplyMarkupRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditInlineMessageReplyMarkupRequest) TypeID() uint32 {
	return EditInlineMessageReplyMarkupRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditInlineMessageReplyMarkupRequest) TypeName() string {
	return "editInlineMessageReplyMarkup"
}

// TypeInfo returns info about TL type.
func (e *EditInlineMessageReplyMarkupRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editInlineMessageReplyMarkup",
		ID:   EditInlineMessageReplyMarkupRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InlineMessageID",
			SchemaName: "inline_message_id",
		},
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditInlineMessageReplyMarkupRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageReplyMarkup#fbf906de as nil")
	}
	b.PutID(EditInlineMessageReplyMarkupRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditInlineMessageReplyMarkupRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageReplyMarkup#fbf906de as nil")
	}
	b.PutString(e.InlineMessageID)
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editInlineMessageReplyMarkup#fbf906de: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageReplyMarkup#fbf906de: field reply_markup: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EditInlineMessageReplyMarkupRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageReplyMarkup#fbf906de to nil")
	}
	if err := b.ConsumeID(EditInlineMessageReplyMarkupRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editInlineMessageReplyMarkup#fbf906de: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditInlineMessageReplyMarkupRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageReplyMarkup#fbf906de to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode editInlineMessageReplyMarkup#fbf906de: field inline_message_id: %w", err)
		}
		e.InlineMessageID = value
	}
	{
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode editInlineMessageReplyMarkup#fbf906de: field reply_markup: %w", err)
		}
		e.ReplyMarkup = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditInlineMessageReplyMarkupRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageReplyMarkup#fbf906de as nil")
	}
	b.ObjStart()
	b.PutID("editInlineMessageReplyMarkup")
	b.Comma()
	b.FieldStart("inline_message_id")
	b.PutString(e.InlineMessageID)
	b.Comma()
	b.FieldStart("reply_markup")
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editInlineMessageReplyMarkup#fbf906de: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageReplyMarkup#fbf906de: field reply_markup: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditInlineMessageReplyMarkupRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageReplyMarkup#fbf906de to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editInlineMessageReplyMarkup"); err != nil {
				return fmt.Errorf("unable to decode editInlineMessageReplyMarkup#fbf906de: %w", err)
			}
		case "inline_message_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode editInlineMessageReplyMarkup#fbf906de: field inline_message_id: %w", err)
			}
			e.InlineMessageID = value
		case "reply_markup":
			value, err := DecodeTDLibJSONReplyMarkup(b)
			if err != nil {
				return fmt.Errorf("unable to decode editInlineMessageReplyMarkup#fbf906de: field reply_markup: %w", err)
			}
			e.ReplyMarkup = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInlineMessageID returns value of InlineMessageID field.
func (e *EditInlineMessageReplyMarkupRequest) GetInlineMessageID() (value string) {
	if e == nil {
		return
	}
	return e.InlineMessageID
}

// GetReplyMarkup returns value of ReplyMarkup field.
func (e *EditInlineMessageReplyMarkupRequest) GetReplyMarkup() (value ReplyMarkupClass) {
	if e == nil {
		return
	}
	return e.ReplyMarkup
}

// EditInlineMessageReplyMarkup invokes method editInlineMessageReplyMarkup#fbf906de returning error if any.
func (c *Client) EditInlineMessageReplyMarkup(ctx context.Context, request *EditInlineMessageReplyMarkupRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
