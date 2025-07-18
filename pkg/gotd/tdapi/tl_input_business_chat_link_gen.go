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

// InputBusinessChatLink represents TL type `inputBusinessChatLink#e2d6df8`.
type InputBusinessChatLink struct {
	// Message draft text that will be added to the input field
	Text FormattedText
	// Link title
	Title string
}

// InputBusinessChatLinkTypeID is TL type id of InputBusinessChatLink.
const InputBusinessChatLinkTypeID = 0xe2d6df8

// Ensuring interfaces in compile-time for InputBusinessChatLink.
var (
	_ bin.Encoder     = &InputBusinessChatLink{}
	_ bin.Decoder     = &InputBusinessChatLink{}
	_ bin.BareEncoder = &InputBusinessChatLink{}
	_ bin.BareDecoder = &InputBusinessChatLink{}
)

func (i *InputBusinessChatLink) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Text.Zero()) {
		return false
	}
	if !(i.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputBusinessChatLink) String() string {
	if i == nil {
		return "InputBusinessChatLink(nil)"
	}
	type Alias InputBusinessChatLink
	return fmt.Sprintf("InputBusinessChatLink%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputBusinessChatLink) TypeID() uint32 {
	return InputBusinessChatLinkTypeID
}

// TypeName returns name of type in TL schema.
func (*InputBusinessChatLink) TypeName() string {
	return "inputBusinessChatLink"
}

// TypeInfo returns info about TL type.
func (i *InputBusinessChatLink) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputBusinessChatLink",
		ID:   InputBusinessChatLinkTypeID,
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
			Name:       "Title",
			SchemaName: "title",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputBusinessChatLink) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBusinessChatLink#e2d6df8 as nil")
	}
	b.PutID(InputBusinessChatLinkTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputBusinessChatLink) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBusinessChatLink#e2d6df8 as nil")
	}
	if err := i.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBusinessChatLink#e2d6df8: field text: %w", err)
	}
	b.PutString(i.Title)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputBusinessChatLink) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBusinessChatLink#e2d6df8 to nil")
	}
	if err := b.ConsumeID(InputBusinessChatLinkTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBusinessChatLink#e2d6df8: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputBusinessChatLink) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBusinessChatLink#e2d6df8 to nil")
	}
	{
		if err := i.Text.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputBusinessChatLink#e2d6df8: field text: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBusinessChatLink#e2d6df8: field title: %w", err)
		}
		i.Title = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputBusinessChatLink) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBusinessChatLink#e2d6df8 as nil")
	}
	b.ObjStart()
	b.PutID("inputBusinessChatLink")
	b.Comma()
	b.FieldStart("text")
	if err := i.Text.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputBusinessChatLink#e2d6df8: field text: %w", err)
	}
	b.Comma()
	b.FieldStart("title")
	b.PutString(i.Title)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputBusinessChatLink) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBusinessChatLink#e2d6df8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputBusinessChatLink"); err != nil {
				return fmt.Errorf("unable to decode inputBusinessChatLink#e2d6df8: %w", err)
			}
		case "text":
			if err := i.Text.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode inputBusinessChatLink#e2d6df8: field text: %w", err)
			}
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputBusinessChatLink#e2d6df8: field title: %w", err)
			}
			i.Title = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (i *InputBusinessChatLink) GetText() (value FormattedText) {
	if i == nil {
		return
	}
	return i.Text
}

// GetTitle returns value of Title field.
func (i *InputBusinessChatLink) GetTitle() (value string) {
	if i == nil {
		return
	}
	return i.Title
}
