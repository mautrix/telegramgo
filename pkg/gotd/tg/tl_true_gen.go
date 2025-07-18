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

// True represents TL type `true#3fedd339`.
// See predefined identifiers¹.
//
// Links:
//  1. https://core.telegram.org/mtproto/TL-formal#predefined-identifiers
//
// See https://core.telegram.org/constructor/true for reference.
type True struct {
}

// TrueTypeID is TL type id of True.
const TrueTypeID = 0x3fedd339

// Ensuring interfaces in compile-time for True.
var (
	_ bin.Encoder     = &True{}
	_ bin.Decoder     = &True{}
	_ bin.BareEncoder = &True{}
	_ bin.BareDecoder = &True{}
)

func (t *True) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *True) String() string {
	if t == nil {
		return "True(nil)"
	}
	type Alias True
	return fmt.Sprintf("True%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*True) TypeID() uint32 {
	return TrueTypeID
}

// TypeName returns name of type in TL schema.
func (*True) TypeName() string {
	return "true"
}

// TypeInfo returns info about TL type.
func (t *True) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "true",
		ID:   TrueTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *True) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode true#3fedd339 as nil")
	}
	b.PutID(TrueTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *True) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode true#3fedd339 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *True) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode true#3fedd339 to nil")
	}
	if err := b.ConsumeID(TrueTypeID); err != nil {
		return fmt.Errorf("unable to decode true#3fedd339: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *True) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode true#3fedd339 to nil")
	}
	return nil
}
