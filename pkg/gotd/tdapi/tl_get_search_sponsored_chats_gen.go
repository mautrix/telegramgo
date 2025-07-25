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

// GetSearchSponsoredChatsRequest represents TL type `getSearchSponsoredChats#5d7d7d4c`.
type GetSearchSponsoredChatsRequest struct {
	// Query the user searches for
	Query string
}

// GetSearchSponsoredChatsRequestTypeID is TL type id of GetSearchSponsoredChatsRequest.
const GetSearchSponsoredChatsRequestTypeID = 0x5d7d7d4c

// Ensuring interfaces in compile-time for GetSearchSponsoredChatsRequest.
var (
	_ bin.Encoder     = &GetSearchSponsoredChatsRequest{}
	_ bin.Decoder     = &GetSearchSponsoredChatsRequest{}
	_ bin.BareEncoder = &GetSearchSponsoredChatsRequest{}
	_ bin.BareDecoder = &GetSearchSponsoredChatsRequest{}
)

func (g *GetSearchSponsoredChatsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Query == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetSearchSponsoredChatsRequest) String() string {
	if g == nil {
		return "GetSearchSponsoredChatsRequest(nil)"
	}
	type Alias GetSearchSponsoredChatsRequest
	return fmt.Sprintf("GetSearchSponsoredChatsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetSearchSponsoredChatsRequest) TypeID() uint32 {
	return GetSearchSponsoredChatsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetSearchSponsoredChatsRequest) TypeName() string {
	return "getSearchSponsoredChats"
}

// TypeInfo returns info about TL type.
func (g *GetSearchSponsoredChatsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getSearchSponsoredChats",
		ID:   GetSearchSponsoredChatsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Query",
			SchemaName: "query",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetSearchSponsoredChatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSearchSponsoredChats#5d7d7d4c as nil")
	}
	b.PutID(GetSearchSponsoredChatsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetSearchSponsoredChatsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSearchSponsoredChats#5d7d7d4c as nil")
	}
	b.PutString(g.Query)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetSearchSponsoredChatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSearchSponsoredChats#5d7d7d4c to nil")
	}
	if err := b.ConsumeID(GetSearchSponsoredChatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getSearchSponsoredChats#5d7d7d4c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetSearchSponsoredChatsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSearchSponsoredChats#5d7d7d4c to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getSearchSponsoredChats#5d7d7d4c: field query: %w", err)
		}
		g.Query = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetSearchSponsoredChatsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getSearchSponsoredChats#5d7d7d4c as nil")
	}
	b.ObjStart()
	b.PutID("getSearchSponsoredChats")
	b.Comma()
	b.FieldStart("query")
	b.PutString(g.Query)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetSearchSponsoredChatsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getSearchSponsoredChats#5d7d7d4c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getSearchSponsoredChats"); err != nil {
				return fmt.Errorf("unable to decode getSearchSponsoredChats#5d7d7d4c: %w", err)
			}
		case "query":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getSearchSponsoredChats#5d7d7d4c: field query: %w", err)
			}
			g.Query = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetQuery returns value of Query field.
func (g *GetSearchSponsoredChatsRequest) GetQuery() (value string) {
	if g == nil {
		return
	}
	return g.Query
}

// GetSearchSponsoredChats invokes method getSearchSponsoredChats#5d7d7d4c returning error if any.
func (c *Client) GetSearchSponsoredChats(ctx context.Context, query string) (*SponsoredChats, error) {
	var result SponsoredChats

	request := &GetSearchSponsoredChatsRequest{
		Query: query,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
