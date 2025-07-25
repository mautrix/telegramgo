// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// GetFutureSaltsRequest represents TL type `get_future_salts#b921bd04`.
type GetFutureSaltsRequest struct {
	// Num field of GetFutureSaltsRequest.
	Num int
}

// GetFutureSaltsRequestTypeID is TL type id of GetFutureSaltsRequest.
const GetFutureSaltsRequestTypeID = 0xb921bd04

// Ensuring interfaces in compile-time for GetFutureSaltsRequest.
var (
	_ bin.Encoder     = &GetFutureSaltsRequest{}
	_ bin.Decoder     = &GetFutureSaltsRequest{}
	_ bin.BareEncoder = &GetFutureSaltsRequest{}
	_ bin.BareDecoder = &GetFutureSaltsRequest{}
)

func (g *GetFutureSaltsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Num == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetFutureSaltsRequest) String() string {
	if g == nil {
		return "GetFutureSaltsRequest(nil)"
	}
	type Alias GetFutureSaltsRequest
	return fmt.Sprintf("GetFutureSaltsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetFutureSaltsRequest) TypeID() uint32 {
	return GetFutureSaltsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetFutureSaltsRequest) TypeName() string {
	return "get_future_salts"
}

// TypeInfo returns info about TL type.
func (g *GetFutureSaltsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "get_future_salts",
		ID:   GetFutureSaltsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Num",
			SchemaName: "num",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetFutureSaltsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode get_future_salts#b921bd04 as nil")
	}
	b.PutID(GetFutureSaltsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetFutureSaltsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode get_future_salts#b921bd04 as nil")
	}
	b.PutInt(g.Num)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetFutureSaltsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode get_future_salts#b921bd04 to nil")
	}
	if err := b.ConsumeID(GetFutureSaltsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode get_future_salts#b921bd04: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetFutureSaltsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode get_future_salts#b921bd04 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode get_future_salts#b921bd04: field num: %w", err)
		}
		g.Num = value
	}
	return nil
}

// GetNum returns value of Num field.
func (g *GetFutureSaltsRequest) GetNum() (value int) {
	if g == nil {
		return
	}
	return g.Num
}
