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

// GzipPacked represents TL type `gzip_packed#3072cfa1`.
type GzipPacked struct {
	// PackedData field of GzipPacked.
	PackedData []byte
}

// GzipPackedTypeID is TL type id of GzipPacked.
const GzipPackedTypeID = 0x3072cfa1

// Ensuring interfaces in compile-time for GzipPacked.
var (
	_ bin.Encoder     = &GzipPacked{}
	_ bin.Decoder     = &GzipPacked{}
	_ bin.BareEncoder = &GzipPacked{}
	_ bin.BareDecoder = &GzipPacked{}
)

func (g *GzipPacked) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.PackedData == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GzipPacked) String() string {
	if g == nil {
		return "GzipPacked(nil)"
	}
	type Alias GzipPacked
	return fmt.Sprintf("GzipPacked%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GzipPacked) TypeID() uint32 {
	return GzipPackedTypeID
}

// TypeName returns name of type in TL schema.
func (*GzipPacked) TypeName() string {
	return "gzip_packed"
}

// TypeInfo returns info about TL type.
func (g *GzipPacked) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "gzip_packed",
		ID:   GzipPackedTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PackedData",
			SchemaName: "packed_data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GzipPacked) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode gzip_packed#3072cfa1 as nil")
	}
	b.PutID(GzipPackedTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GzipPacked) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode gzip_packed#3072cfa1 as nil")
	}
	b.PutBytes(g.PackedData)
	return nil
}

// Decode implements bin.Decoder.
func (g *GzipPacked) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode gzip_packed#3072cfa1 to nil")
	}
	if err := b.ConsumeID(GzipPackedTypeID); err != nil {
		return fmt.Errorf("unable to decode gzip_packed#3072cfa1: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GzipPacked) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode gzip_packed#3072cfa1 to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode gzip_packed#3072cfa1: field packed_data: %w", err)
		}
		g.PackedData = value
	}
	return nil
}

// GetPackedData returns value of PackedData field.
func (g *GzipPacked) GetPackedData() (value []byte) {
	if g == nil {
		return
	}
	return g.PackedData
}
