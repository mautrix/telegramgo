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

// GetPremiumStickerExamplesRequest represents TL type `getPremiumStickerExamples#5369cb98`.
type GetPremiumStickerExamplesRequest struct {
}

// GetPremiumStickerExamplesRequestTypeID is TL type id of GetPremiumStickerExamplesRequest.
const GetPremiumStickerExamplesRequestTypeID = 0x5369cb98

// Ensuring interfaces in compile-time for GetPremiumStickerExamplesRequest.
var (
	_ bin.Encoder     = &GetPremiumStickerExamplesRequest{}
	_ bin.Decoder     = &GetPremiumStickerExamplesRequest{}
	_ bin.BareEncoder = &GetPremiumStickerExamplesRequest{}
	_ bin.BareDecoder = &GetPremiumStickerExamplesRequest{}
)

func (g *GetPremiumStickerExamplesRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetPremiumStickerExamplesRequest) String() string {
	if g == nil {
		return "GetPremiumStickerExamplesRequest(nil)"
	}
	type Alias GetPremiumStickerExamplesRequest
	return fmt.Sprintf("GetPremiumStickerExamplesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetPremiumStickerExamplesRequest) TypeID() uint32 {
	return GetPremiumStickerExamplesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetPremiumStickerExamplesRequest) TypeName() string {
	return "getPremiumStickerExamples"
}

// TypeInfo returns info about TL type.
func (g *GetPremiumStickerExamplesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getPremiumStickerExamples",
		ID:   GetPremiumStickerExamplesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetPremiumStickerExamplesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPremiumStickerExamples#5369cb98 as nil")
	}
	b.PutID(GetPremiumStickerExamplesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetPremiumStickerExamplesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPremiumStickerExamples#5369cb98 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetPremiumStickerExamplesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPremiumStickerExamples#5369cb98 to nil")
	}
	if err := b.ConsumeID(GetPremiumStickerExamplesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getPremiumStickerExamples#5369cb98: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetPremiumStickerExamplesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPremiumStickerExamples#5369cb98 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetPremiumStickerExamplesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getPremiumStickerExamples#5369cb98 as nil")
	}
	b.ObjStart()
	b.PutID("getPremiumStickerExamples")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetPremiumStickerExamplesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getPremiumStickerExamples#5369cb98 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getPremiumStickerExamples"); err != nil {
				return fmt.Errorf("unable to decode getPremiumStickerExamples#5369cb98: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPremiumStickerExamples invokes method getPremiumStickerExamples#5369cb98 returning error if any.
func (c *Client) GetPremiumStickerExamples(ctx context.Context) (*Stickers, error) {
	var result Stickers

	request := &GetPremiumStickerExamplesRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
