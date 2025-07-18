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

// GetRecentlyOpenedChatsRequest represents TL type `getRecentlyOpenedChats#8d4fb223`.
type GetRecentlyOpenedChatsRequest struct {
	// The maximum number of chats to be returned
	Limit int32
}

// GetRecentlyOpenedChatsRequestTypeID is TL type id of GetRecentlyOpenedChatsRequest.
const GetRecentlyOpenedChatsRequestTypeID = 0x8d4fb223

// Ensuring interfaces in compile-time for GetRecentlyOpenedChatsRequest.
var (
	_ bin.Encoder     = &GetRecentlyOpenedChatsRequest{}
	_ bin.Decoder     = &GetRecentlyOpenedChatsRequest{}
	_ bin.BareEncoder = &GetRecentlyOpenedChatsRequest{}
	_ bin.BareDecoder = &GetRecentlyOpenedChatsRequest{}
)

func (g *GetRecentlyOpenedChatsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetRecentlyOpenedChatsRequest) String() string {
	if g == nil {
		return "GetRecentlyOpenedChatsRequest(nil)"
	}
	type Alias GetRecentlyOpenedChatsRequest
	return fmt.Sprintf("GetRecentlyOpenedChatsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetRecentlyOpenedChatsRequest) TypeID() uint32 {
	return GetRecentlyOpenedChatsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetRecentlyOpenedChatsRequest) TypeName() string {
	return "getRecentlyOpenedChats"
}

// TypeInfo returns info about TL type.
func (g *GetRecentlyOpenedChatsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getRecentlyOpenedChats",
		ID:   GetRecentlyOpenedChatsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetRecentlyOpenedChatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getRecentlyOpenedChats#8d4fb223 as nil")
	}
	b.PutID(GetRecentlyOpenedChatsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetRecentlyOpenedChatsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getRecentlyOpenedChats#8d4fb223 as nil")
	}
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetRecentlyOpenedChatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getRecentlyOpenedChats#8d4fb223 to nil")
	}
	if err := b.ConsumeID(GetRecentlyOpenedChatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getRecentlyOpenedChats#8d4fb223: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetRecentlyOpenedChatsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getRecentlyOpenedChats#8d4fb223 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getRecentlyOpenedChats#8d4fb223: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetRecentlyOpenedChatsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getRecentlyOpenedChats#8d4fb223 as nil")
	}
	b.ObjStart()
	b.PutID("getRecentlyOpenedChats")
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetRecentlyOpenedChatsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getRecentlyOpenedChats#8d4fb223 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getRecentlyOpenedChats"); err != nil {
				return fmt.Errorf("unable to decode getRecentlyOpenedChats#8d4fb223: %w", err)
			}
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getRecentlyOpenedChats#8d4fb223: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLimit returns value of Limit field.
func (g *GetRecentlyOpenedChatsRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetRecentlyOpenedChats invokes method getRecentlyOpenedChats#8d4fb223 returning error if any.
func (c *Client) GetRecentlyOpenedChats(ctx context.Context, limit int32) (*Chats, error) {
	var result Chats

	request := &GetRecentlyOpenedChatsRequest{
		Limit: limit,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
