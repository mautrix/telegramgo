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

// LongVector is a box for Vector<long>
type LongVector struct {
	// Elements of Vector<long>
	Elems []int64
}

// LongVectorTypeID is TL type id of LongVector.
const LongVectorTypeID = bin.TypeVector

// Ensuring interfaces in compile-time for LongVector.
var (
	_ bin.Encoder     = &LongVector{}
	_ bin.Decoder     = &LongVector{}
	_ bin.BareEncoder = &LongVector{}
	_ bin.BareDecoder = &LongVector{}
)

func (vec *LongVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *LongVector) String() string {
	if vec == nil {
		return "LongVector(nil)"
	}
	type Alias LongVector
	return fmt.Sprintf("LongVector%+v", Alias(*vec))
}

// FillFrom fills LongVector from given interface.
func (vec *LongVector) FillFrom(from interface {
	GetElems() (value []int64)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LongVector) TypeID() uint32 {
	return LongVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*LongVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *LongVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   LongVectorTypeID,
	}
	if vec == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Elems",
			SchemaName: "Elems",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (vec *LongVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<long> as nil")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *LongVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<long> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for _, v := range vec.Elems {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *LongVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<long> to nil")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *LongVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<long> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<long>: field Elems: %w", err)
		}

		if headerLen > 0 {
			vec.Elems = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode Vector<long>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *LongVector) GetElems() (value []int64) {
	if vec == nil {
		return
	}
	return vec.Elems
}
