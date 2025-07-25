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

// GetPremiumLimitRequest represents TL type `getPremiumLimit#4017fcea`.
type GetPremiumLimitRequest struct {
	// Type of the limit
	LimitType PremiumLimitTypeClass
}

// GetPremiumLimitRequestTypeID is TL type id of GetPremiumLimitRequest.
const GetPremiumLimitRequestTypeID = 0x4017fcea

// Ensuring interfaces in compile-time for GetPremiumLimitRequest.
var (
	_ bin.Encoder     = &GetPremiumLimitRequest{}
	_ bin.Decoder     = &GetPremiumLimitRequest{}
	_ bin.BareEncoder = &GetPremiumLimitRequest{}
	_ bin.BareDecoder = &GetPremiumLimitRequest{}
)

func (g *GetPremiumLimitRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LimitType == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetPremiumLimitRequest) String() string {
	if g == nil {
		return "GetPremiumLimitRequest(nil)"
	}
	type Alias GetPremiumLimitRequest
	return fmt.Sprintf("GetPremiumLimitRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetPremiumLimitRequest) TypeID() uint32 {
	return GetPremiumLimitRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetPremiumLimitRequest) TypeName() string {
	return "getPremiumLimit"
}

// TypeInfo returns info about TL type.
func (g *GetPremiumLimitRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getPremiumLimit",
		ID:   GetPremiumLimitRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LimitType",
			SchemaName: "limit_type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetPremiumLimitRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPremiumLimit#4017fcea as nil")
	}
	b.PutID(GetPremiumLimitRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetPremiumLimitRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPremiumLimit#4017fcea as nil")
	}
	if g.LimitType == nil {
		return fmt.Errorf("unable to encode getPremiumLimit#4017fcea: field limit_type is nil")
	}
	if err := g.LimitType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getPremiumLimit#4017fcea: field limit_type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetPremiumLimitRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPremiumLimit#4017fcea to nil")
	}
	if err := b.ConsumeID(GetPremiumLimitRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getPremiumLimit#4017fcea: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetPremiumLimitRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPremiumLimit#4017fcea to nil")
	}
	{
		value, err := DecodePremiumLimitType(b)
		if err != nil {
			return fmt.Errorf("unable to decode getPremiumLimit#4017fcea: field limit_type: %w", err)
		}
		g.LimitType = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetPremiumLimitRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getPremiumLimit#4017fcea as nil")
	}
	b.ObjStart()
	b.PutID("getPremiumLimit")
	b.Comma()
	b.FieldStart("limit_type")
	if g.LimitType == nil {
		return fmt.Errorf("unable to encode getPremiumLimit#4017fcea: field limit_type is nil")
	}
	if err := g.LimitType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getPremiumLimit#4017fcea: field limit_type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetPremiumLimitRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getPremiumLimit#4017fcea to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getPremiumLimit"); err != nil {
				return fmt.Errorf("unable to decode getPremiumLimit#4017fcea: %w", err)
			}
		case "limit_type":
			value, err := DecodeTDLibJSONPremiumLimitType(b)
			if err != nil {
				return fmt.Errorf("unable to decode getPremiumLimit#4017fcea: field limit_type: %w", err)
			}
			g.LimitType = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLimitType returns value of LimitType field.
func (g *GetPremiumLimitRequest) GetLimitType() (value PremiumLimitTypeClass) {
	if g == nil {
		return
	}
	return g.LimitType
}

// GetPremiumLimit invokes method getPremiumLimit#4017fcea returning error if any.
func (c *Client) GetPremiumLimit(ctx context.Context, limittype PremiumLimitTypeClass) (*PremiumLimit, error) {
	var result PremiumLimit

	request := &GetPremiumLimitRequest{
		LimitType: limittype,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
