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

// GetDatabaseStatisticsRequest represents TL type `getDatabaseStatistics#8c33d4b9`.
type GetDatabaseStatisticsRequest struct {
}

// GetDatabaseStatisticsRequestTypeID is TL type id of GetDatabaseStatisticsRequest.
const GetDatabaseStatisticsRequestTypeID = 0x8c33d4b9

// Ensuring interfaces in compile-time for GetDatabaseStatisticsRequest.
var (
	_ bin.Encoder     = &GetDatabaseStatisticsRequest{}
	_ bin.Decoder     = &GetDatabaseStatisticsRequest{}
	_ bin.BareEncoder = &GetDatabaseStatisticsRequest{}
	_ bin.BareDecoder = &GetDatabaseStatisticsRequest{}
)

func (g *GetDatabaseStatisticsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetDatabaseStatisticsRequest) String() string {
	if g == nil {
		return "GetDatabaseStatisticsRequest(nil)"
	}
	type Alias GetDatabaseStatisticsRequest
	return fmt.Sprintf("GetDatabaseStatisticsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetDatabaseStatisticsRequest) TypeID() uint32 {
	return GetDatabaseStatisticsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetDatabaseStatisticsRequest) TypeName() string {
	return "getDatabaseStatistics"
}

// TypeInfo returns info about TL type.
func (g *GetDatabaseStatisticsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getDatabaseStatistics",
		ID:   GetDatabaseStatisticsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetDatabaseStatisticsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getDatabaseStatistics#8c33d4b9 as nil")
	}
	b.PutID(GetDatabaseStatisticsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetDatabaseStatisticsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getDatabaseStatistics#8c33d4b9 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetDatabaseStatisticsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getDatabaseStatistics#8c33d4b9 to nil")
	}
	if err := b.ConsumeID(GetDatabaseStatisticsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getDatabaseStatistics#8c33d4b9: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetDatabaseStatisticsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getDatabaseStatistics#8c33d4b9 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetDatabaseStatisticsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getDatabaseStatistics#8c33d4b9 as nil")
	}
	b.ObjStart()
	b.PutID("getDatabaseStatistics")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetDatabaseStatisticsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getDatabaseStatistics#8c33d4b9 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getDatabaseStatistics"); err != nil {
				return fmt.Errorf("unable to decode getDatabaseStatistics#8c33d4b9: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDatabaseStatistics invokes method getDatabaseStatistics#8c33d4b9 returning error if any.
func (c *Client) GetDatabaseStatistics(ctx context.Context) (*DatabaseStatistics, error) {
	var result DatabaseStatistics

	request := &GetDatabaseStatisticsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
