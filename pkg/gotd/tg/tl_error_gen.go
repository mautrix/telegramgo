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

// Error represents TL type `error#c4b9f9bb`.
// Error.
//
// See https://core.telegram.org/constructor/error for reference.
type Error struct {
	// Error code
	Code int
	// Message
	Text string
}

// ErrorTypeID is TL type id of Error.
const ErrorTypeID = 0xc4b9f9bb

// Ensuring interfaces in compile-time for Error.
var (
	_ bin.Encoder     = &Error{}
	_ bin.Decoder     = &Error{}
	_ bin.BareEncoder = &Error{}
	_ bin.BareDecoder = &Error{}
)

func (e *Error) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Code == 0) {
		return false
	}
	if !(e.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *Error) String() string {
	if e == nil {
		return "Error(nil)"
	}
	type Alias Error
	return fmt.Sprintf("Error%+v", Alias(*e))
}

// FillFrom fills Error from given interface.
func (e *Error) FillFrom(from interface {
	GetCode() (value int)
	GetText() (value string)
}) {
	e.Code = from.GetCode()
	e.Text = from.GetText()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Error) TypeID() uint32 {
	return ErrorTypeID
}

// TypeName returns name of type in TL schema.
func (*Error) TypeName() string {
	return "error"
}

// TypeInfo returns info about TL type.
func (e *Error) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "error",
		ID:   ErrorTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Code",
			SchemaName: "code",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *Error) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode error#c4b9f9bb as nil")
	}
	b.PutID(ErrorTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *Error) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode error#c4b9f9bb as nil")
	}
	b.PutInt(e.Code)
	b.PutString(e.Text)
	return nil
}

// Decode implements bin.Decoder.
func (e *Error) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode error#c4b9f9bb to nil")
	}
	if err := b.ConsumeID(ErrorTypeID); err != nil {
		return fmt.Errorf("unable to decode error#c4b9f9bb: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *Error) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode error#c4b9f9bb to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode error#c4b9f9bb: field code: %w", err)
		}
		e.Code = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode error#c4b9f9bb: field text: %w", err)
		}
		e.Text = value
	}
	return nil
}

// GetCode returns value of Code field.
func (e *Error) GetCode() (value int) {
	if e == nil {
		return
	}
	return e.Code
}

// GetText returns value of Text field.
func (e *Error) GetText() (value string) {
	if e == nil {
		return
	}
	return e.Text
}
