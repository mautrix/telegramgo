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

// CustomRequestResult represents TL type `customRequestResult#88326ffc`.
type CustomRequestResult struct {
	// A JSON-serialized result
	Result string
}

// CustomRequestResultTypeID is TL type id of CustomRequestResult.
const CustomRequestResultTypeID = 0x88326ffc

// Ensuring interfaces in compile-time for CustomRequestResult.
var (
	_ bin.Encoder     = &CustomRequestResult{}
	_ bin.Decoder     = &CustomRequestResult{}
	_ bin.BareEncoder = &CustomRequestResult{}
	_ bin.BareDecoder = &CustomRequestResult{}
)

func (c *CustomRequestResult) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Result == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CustomRequestResult) String() string {
	if c == nil {
		return "CustomRequestResult(nil)"
	}
	type Alias CustomRequestResult
	return fmt.Sprintf("CustomRequestResult%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CustomRequestResult) TypeID() uint32 {
	return CustomRequestResultTypeID
}

// TypeName returns name of type in TL schema.
func (*CustomRequestResult) TypeName() string {
	return "customRequestResult"
}

// TypeInfo returns info about TL type.
func (c *CustomRequestResult) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "customRequestResult",
		ID:   CustomRequestResultTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Result",
			SchemaName: "result",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CustomRequestResult) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode customRequestResult#88326ffc as nil")
	}
	b.PutID(CustomRequestResultTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CustomRequestResult) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode customRequestResult#88326ffc as nil")
	}
	b.PutString(c.Result)
	return nil
}

// Decode implements bin.Decoder.
func (c *CustomRequestResult) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode customRequestResult#88326ffc to nil")
	}
	if err := b.ConsumeID(CustomRequestResultTypeID); err != nil {
		return fmt.Errorf("unable to decode customRequestResult#88326ffc: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CustomRequestResult) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode customRequestResult#88326ffc to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode customRequestResult#88326ffc: field result: %w", err)
		}
		c.Result = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CustomRequestResult) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode customRequestResult#88326ffc as nil")
	}
	b.ObjStart()
	b.PutID("customRequestResult")
	b.Comma()
	b.FieldStart("result")
	b.PutString(c.Result)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CustomRequestResult) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode customRequestResult#88326ffc to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("customRequestResult"); err != nil {
				return fmt.Errorf("unable to decode customRequestResult#88326ffc: %w", err)
			}
		case "result":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode customRequestResult#88326ffc: field result: %w", err)
			}
			c.Result = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetResult returns value of Result field.
func (c *CustomRequestResult) GetResult() (value string) {
	if c == nil {
		return
	}
	return c.Result
}
