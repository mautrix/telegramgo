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

// GetStoryAvailableReactionsRequest represents TL type `getStoryAvailableReactions#23854d3b`.
type GetStoryAvailableReactionsRequest struct {
	// Number of reaction per row, 5-25
	RowSize int32
}

// GetStoryAvailableReactionsRequestTypeID is TL type id of GetStoryAvailableReactionsRequest.
const GetStoryAvailableReactionsRequestTypeID = 0x23854d3b

// Ensuring interfaces in compile-time for GetStoryAvailableReactionsRequest.
var (
	_ bin.Encoder     = &GetStoryAvailableReactionsRequest{}
	_ bin.Decoder     = &GetStoryAvailableReactionsRequest{}
	_ bin.BareEncoder = &GetStoryAvailableReactionsRequest{}
	_ bin.BareDecoder = &GetStoryAvailableReactionsRequest{}
)

func (g *GetStoryAvailableReactionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.RowSize == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetStoryAvailableReactionsRequest) String() string {
	if g == nil {
		return "GetStoryAvailableReactionsRequest(nil)"
	}
	type Alias GetStoryAvailableReactionsRequest
	return fmt.Sprintf("GetStoryAvailableReactionsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetStoryAvailableReactionsRequest) TypeID() uint32 {
	return GetStoryAvailableReactionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetStoryAvailableReactionsRequest) TypeName() string {
	return "getStoryAvailableReactions"
}

// TypeInfo returns info about TL type.
func (g *GetStoryAvailableReactionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getStoryAvailableReactions",
		ID:   GetStoryAvailableReactionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RowSize",
			SchemaName: "row_size",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetStoryAvailableReactionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStoryAvailableReactions#23854d3b as nil")
	}
	b.PutID(GetStoryAvailableReactionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetStoryAvailableReactionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStoryAvailableReactions#23854d3b as nil")
	}
	b.PutInt32(g.RowSize)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetStoryAvailableReactionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStoryAvailableReactions#23854d3b to nil")
	}
	if err := b.ConsumeID(GetStoryAvailableReactionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getStoryAvailableReactions#23854d3b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetStoryAvailableReactionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStoryAvailableReactions#23854d3b to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getStoryAvailableReactions#23854d3b: field row_size: %w", err)
		}
		g.RowSize = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetStoryAvailableReactionsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getStoryAvailableReactions#23854d3b as nil")
	}
	b.ObjStart()
	b.PutID("getStoryAvailableReactions")
	b.Comma()
	b.FieldStart("row_size")
	b.PutInt32(g.RowSize)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetStoryAvailableReactionsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getStoryAvailableReactions#23854d3b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getStoryAvailableReactions"); err != nil {
				return fmt.Errorf("unable to decode getStoryAvailableReactions#23854d3b: %w", err)
			}
		case "row_size":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getStoryAvailableReactions#23854d3b: field row_size: %w", err)
			}
			g.RowSize = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetRowSize returns value of RowSize field.
func (g *GetStoryAvailableReactionsRequest) GetRowSize() (value int32) {
	if g == nil {
		return
	}
	return g.RowSize
}

// GetStoryAvailableReactions invokes method getStoryAvailableReactions#23854d3b returning error if any.
func (c *Client) GetStoryAvailableReactions(ctx context.Context, rowsize int32) (*AvailableReactions, error) {
	var result AvailableReactions

	request := &GetStoryAvailableReactionsRequest{
		RowSize: rowsize,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
