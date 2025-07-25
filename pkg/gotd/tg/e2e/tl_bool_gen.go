// Code generated by gotdgen, DO NOT EDIT.

package e2e

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

// BoolFalse represents TL type `boolFalse#bc799737`.
// Constructor may be interpreted as a booleanfalse value.
//
// See https://core.telegram.org/constructor/boolFalse for reference.
type BoolFalse struct {
}

// BoolFalseTypeID is TL type id of BoolFalse.
const BoolFalseTypeID = 0xbc799737

// construct implements constructor of BoolClass.
func (b BoolFalse) construct() BoolClass { return &b }

// Ensuring interfaces in compile-time for BoolFalse.
var (
	_ bin.Encoder     = &BoolFalse{}
	_ bin.Decoder     = &BoolFalse{}
	_ bin.BareEncoder = &BoolFalse{}
	_ bin.BareDecoder = &BoolFalse{}

	_ BoolClass = &BoolFalse{}
)

func (b *BoolFalse) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BoolFalse) String() string {
	if b == nil {
		return "BoolFalse(nil)"
	}
	type Alias BoolFalse
	return fmt.Sprintf("BoolFalse%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BoolFalse) TypeID() uint32 {
	return BoolFalseTypeID
}

// TypeName returns name of type in TL schema.
func (*BoolFalse) TypeName() string {
	return "boolFalse"
}

// TypeInfo returns info about TL type.
func (b *BoolFalse) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "boolFalse",
		ID:   BoolFalseTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BoolFalse) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode boolFalse#bc799737 as nil")
	}
	buf.PutID(BoolFalseTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BoolFalse) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode boolFalse#bc799737 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BoolFalse) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode boolFalse#bc799737 to nil")
	}
	if err := buf.ConsumeID(BoolFalseTypeID); err != nil {
		return fmt.Errorf("unable to decode boolFalse#bc799737: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BoolFalse) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode boolFalse#bc799737 to nil")
	}
	return nil
}

// BoolTrue represents TL type `boolTrue#997275b5`.
// The constructor can be interpreted as a booleantrue value.
//
// See https://core.telegram.org/constructor/boolTrue for reference.
type BoolTrue struct {
}

// BoolTrueTypeID is TL type id of BoolTrue.
const BoolTrueTypeID = 0x997275b5

// construct implements constructor of BoolClass.
func (b BoolTrue) construct() BoolClass { return &b }

// Ensuring interfaces in compile-time for BoolTrue.
var (
	_ bin.Encoder     = &BoolTrue{}
	_ bin.Decoder     = &BoolTrue{}
	_ bin.BareEncoder = &BoolTrue{}
	_ bin.BareDecoder = &BoolTrue{}

	_ BoolClass = &BoolTrue{}
)

func (b *BoolTrue) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BoolTrue) String() string {
	if b == nil {
		return "BoolTrue(nil)"
	}
	type Alias BoolTrue
	return fmt.Sprintf("BoolTrue%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BoolTrue) TypeID() uint32 {
	return BoolTrueTypeID
}

// TypeName returns name of type in TL schema.
func (*BoolTrue) TypeName() string {
	return "boolTrue"
}

// TypeInfo returns info about TL type.
func (b *BoolTrue) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "boolTrue",
		ID:   BoolTrueTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BoolTrue) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode boolTrue#997275b5 as nil")
	}
	buf.PutID(BoolTrueTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BoolTrue) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode boolTrue#997275b5 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BoolTrue) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode boolTrue#997275b5 to nil")
	}
	if err := buf.ConsumeID(BoolTrueTypeID); err != nil {
		return fmt.Errorf("unable to decode boolTrue#997275b5: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BoolTrue) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode boolTrue#997275b5 to nil")
	}
	return nil
}

// BoolClassName is schema name of BoolClass.
const BoolClassName = "Bool"

// BoolClass represents Bool generic type.
//
// See https://core.telegram.org/type/Bool for reference.
//
// Example:
//
//	g, err := e2e.DecodeBool(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *e2e.BoolFalse: // boolFalse#bc799737
//	case *e2e.BoolTrue: // boolTrue#997275b5
//	default: panic(v)
//	}
type BoolClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() BoolClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeBool implements binary de-serialization for BoolClass.
func DecodeBool(buf *bin.Buffer) (BoolClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BoolFalseTypeID:
		// Decoding boolFalse#bc799737.
		v := BoolFalse{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BoolClass: %w", err)
		}
		return &v, nil
	case BoolTrueTypeID:
		// Decoding boolTrue#997275b5.
		v := BoolTrue{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BoolClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BoolClass: %w", bin.NewUnexpectedID(id))
	}
}

// Bool boxes the BoolClass providing a helper.
type BoolBox struct {
	Bool BoolClass
}

// Decode implements bin.Decoder for BoolBox.
func (b *BoolBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BoolBox to nil")
	}
	v, err := DecodeBool(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Bool = v
	return nil
}

// Encode implements bin.Encode for BoolBox.
func (b *BoolBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Bool == nil {
		return fmt.Errorf("unable to encode BoolClass as nil")
	}
	return b.Bool.Encode(buf)
}
