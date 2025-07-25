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

// GetNetworkStatisticsRequest represents TL type `getNetworkStatistics#c537581e`.
type GetNetworkStatisticsRequest struct {
	// Pass true to get statistics only for the current library launch
	OnlyCurrent bool
}

// GetNetworkStatisticsRequestTypeID is TL type id of GetNetworkStatisticsRequest.
const GetNetworkStatisticsRequestTypeID = 0xc537581e

// Ensuring interfaces in compile-time for GetNetworkStatisticsRequest.
var (
	_ bin.Encoder     = &GetNetworkStatisticsRequest{}
	_ bin.Decoder     = &GetNetworkStatisticsRequest{}
	_ bin.BareEncoder = &GetNetworkStatisticsRequest{}
	_ bin.BareDecoder = &GetNetworkStatisticsRequest{}
)

func (g *GetNetworkStatisticsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.OnlyCurrent == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetNetworkStatisticsRequest) String() string {
	if g == nil {
		return "GetNetworkStatisticsRequest(nil)"
	}
	type Alias GetNetworkStatisticsRequest
	return fmt.Sprintf("GetNetworkStatisticsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetNetworkStatisticsRequest) TypeID() uint32 {
	return GetNetworkStatisticsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetNetworkStatisticsRequest) TypeName() string {
	return "getNetworkStatistics"
}

// TypeInfo returns info about TL type.
func (g *GetNetworkStatisticsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getNetworkStatistics",
		ID:   GetNetworkStatisticsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "OnlyCurrent",
			SchemaName: "only_current",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetNetworkStatisticsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getNetworkStatistics#c537581e as nil")
	}
	b.PutID(GetNetworkStatisticsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetNetworkStatisticsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getNetworkStatistics#c537581e as nil")
	}
	b.PutBool(g.OnlyCurrent)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetNetworkStatisticsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getNetworkStatistics#c537581e to nil")
	}
	if err := b.ConsumeID(GetNetworkStatisticsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getNetworkStatistics#c537581e: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetNetworkStatisticsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getNetworkStatistics#c537581e to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getNetworkStatistics#c537581e: field only_current: %w", err)
		}
		g.OnlyCurrent = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetNetworkStatisticsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getNetworkStatistics#c537581e as nil")
	}
	b.ObjStart()
	b.PutID("getNetworkStatistics")
	b.Comma()
	b.FieldStart("only_current")
	b.PutBool(g.OnlyCurrent)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetNetworkStatisticsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getNetworkStatistics#c537581e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getNetworkStatistics"); err != nil {
				return fmt.Errorf("unable to decode getNetworkStatistics#c537581e: %w", err)
			}
		case "only_current":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getNetworkStatistics#c537581e: field only_current: %w", err)
			}
			g.OnlyCurrent = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetOnlyCurrent returns value of OnlyCurrent field.
func (g *GetNetworkStatisticsRequest) GetOnlyCurrent() (value bool) {
	if g == nil {
		return
	}
	return g.OnlyCurrent
}

// GetNetworkStatistics invokes method getNetworkStatistics#c537581e returning error if any.
func (c *Client) GetNetworkStatistics(ctx context.Context, onlycurrent bool) (*NetworkStatistics, error) {
	var result NetworkStatistics

	request := &GetNetworkStatisticsRequest{
		OnlyCurrent: onlycurrent,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
