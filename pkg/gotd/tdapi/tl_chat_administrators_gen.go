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

// ChatAdministrators represents TL type `chatAdministrators#5141ca21`.
type ChatAdministrators struct {
	// A list of chat administrators
	Administrators []ChatAdministrator
}

// ChatAdministratorsTypeID is TL type id of ChatAdministrators.
const ChatAdministratorsTypeID = 0x5141ca21

// Ensuring interfaces in compile-time for ChatAdministrators.
var (
	_ bin.Encoder     = &ChatAdministrators{}
	_ bin.Decoder     = &ChatAdministrators{}
	_ bin.BareEncoder = &ChatAdministrators{}
	_ bin.BareDecoder = &ChatAdministrators{}
)

func (c *ChatAdministrators) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Administrators == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatAdministrators) String() string {
	if c == nil {
		return "ChatAdministrators(nil)"
	}
	type Alias ChatAdministrators
	return fmt.Sprintf("ChatAdministrators%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatAdministrators) TypeID() uint32 {
	return ChatAdministratorsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatAdministrators) TypeName() string {
	return "chatAdministrators"
}

// TypeInfo returns info about TL type.
func (c *ChatAdministrators) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatAdministrators",
		ID:   ChatAdministratorsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Administrators",
			SchemaName: "administrators",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatAdministrators) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdministrators#5141ca21 as nil")
	}
	b.PutID(ChatAdministratorsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatAdministrators) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdministrators#5141ca21 as nil")
	}
	b.PutInt(len(c.Administrators))
	for idx, v := range c.Administrators {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare chatAdministrators#5141ca21: field administrators element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatAdministrators) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdministrators#5141ca21 to nil")
	}
	if err := b.ConsumeID(ChatAdministratorsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatAdministrators#5141ca21: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatAdministrators) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdministrators#5141ca21 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministrators#5141ca21: field administrators: %w", err)
		}

		if headerLen > 0 {
			c.Administrators = make([]ChatAdministrator, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatAdministrator
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare chatAdministrators#5141ca21: field administrators: %w", err)
			}
			c.Administrators = append(c.Administrators, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatAdministrators) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdministrators#5141ca21 as nil")
	}
	b.ObjStart()
	b.PutID("chatAdministrators")
	b.Comma()
	b.FieldStart("administrators")
	b.ArrStart()
	for idx, v := range c.Administrators {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode chatAdministrators#5141ca21: field administrators element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatAdministrators) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdministrators#5141ca21 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatAdministrators"); err != nil {
				return fmt.Errorf("unable to decode chatAdministrators#5141ca21: %w", err)
			}
		case "administrators":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value ChatAdministrator
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode chatAdministrators#5141ca21: field administrators: %w", err)
				}
				c.Administrators = append(c.Administrators, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode chatAdministrators#5141ca21: field administrators: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAdministrators returns value of Administrators field.
func (c *ChatAdministrators) GetAdministrators() (value []ChatAdministrator) {
	if c == nil {
		return
	}
	return c.Administrators
}
