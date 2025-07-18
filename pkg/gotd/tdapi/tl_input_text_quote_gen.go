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

// InputTextQuote represents TL type `inputTextQuote#b74a6d1c`.
type InputTextQuote struct {
	// Text of the quote; 0-getOption("message_reply_quote_length_max") characters. Only Bold
	// Italic, Underline, Strikethrough, Spoiler, and CustomEmoji entities are allowed to be
	// kept and must be kept in the quote
	Text FormattedText
	// Quote position in the original message in UTF-16 code units
	Position int32
}

// InputTextQuoteTypeID is TL type id of InputTextQuote.
const InputTextQuoteTypeID = 0xb74a6d1c

// Ensuring interfaces in compile-time for InputTextQuote.
var (
	_ bin.Encoder     = &InputTextQuote{}
	_ bin.Decoder     = &InputTextQuote{}
	_ bin.BareEncoder = &InputTextQuote{}
	_ bin.BareDecoder = &InputTextQuote{}
)

func (i *InputTextQuote) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Text.Zero()) {
		return false
	}
	if !(i.Position == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputTextQuote) String() string {
	if i == nil {
		return "InputTextQuote(nil)"
	}
	type Alias InputTextQuote
	return fmt.Sprintf("InputTextQuote%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputTextQuote) TypeID() uint32 {
	return InputTextQuoteTypeID
}

// TypeName returns name of type in TL schema.
func (*InputTextQuote) TypeName() string {
	return "inputTextQuote"
}

// TypeInfo returns info about TL type.
func (i *InputTextQuote) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputTextQuote",
		ID:   InputTextQuoteTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "Position",
			SchemaName: "position",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputTextQuote) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputTextQuote#b74a6d1c as nil")
	}
	b.PutID(InputTextQuoteTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputTextQuote) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputTextQuote#b74a6d1c as nil")
	}
	if err := i.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputTextQuote#b74a6d1c: field text: %w", err)
	}
	b.PutInt32(i.Position)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputTextQuote) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputTextQuote#b74a6d1c to nil")
	}
	if err := b.ConsumeID(InputTextQuoteTypeID); err != nil {
		return fmt.Errorf("unable to decode inputTextQuote#b74a6d1c: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputTextQuote) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputTextQuote#b74a6d1c to nil")
	}
	{
		if err := i.Text.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputTextQuote#b74a6d1c: field text: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inputTextQuote#b74a6d1c: field position: %w", err)
		}
		i.Position = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputTextQuote) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputTextQuote#b74a6d1c as nil")
	}
	b.ObjStart()
	b.PutID("inputTextQuote")
	b.Comma()
	b.FieldStart("text")
	if err := i.Text.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputTextQuote#b74a6d1c: field text: %w", err)
	}
	b.Comma()
	b.FieldStart("position")
	b.PutInt32(i.Position)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputTextQuote) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputTextQuote#b74a6d1c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputTextQuote"); err != nil {
				return fmt.Errorf("unable to decode inputTextQuote#b74a6d1c: %w", err)
			}
		case "text":
			if err := i.Text.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode inputTextQuote#b74a6d1c: field text: %w", err)
			}
		case "position":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode inputTextQuote#b74a6d1c: field position: %w", err)
			}
			i.Position = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (i *InputTextQuote) GetText() (value FormattedText) {
	if i == nil {
		return
	}
	return i.Text
}

// GetPosition returns value of Position field.
func (i *InputTextQuote) GetPosition() (value int32) {
	if i == nil {
		return
	}
	return i.Position
}
