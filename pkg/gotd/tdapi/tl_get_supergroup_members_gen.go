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

// GetSupergroupMembersRequest represents TL type `getSupergroupMembers#ddf821c8`.
type GetSupergroupMembersRequest struct {
	// Identifier of the supergroup or channel
	SupergroupID int64
	// The type of users to return; pass null to use supergroupMembersFilterRecent
	Filter SupergroupMembersFilterClass
	// Number of users to skip
	Offset int32
	// The maximum number of users to be returned; up to 200
	Limit int32
}

// GetSupergroupMembersRequestTypeID is TL type id of GetSupergroupMembersRequest.
const GetSupergroupMembersRequestTypeID = 0xddf821c8

// Ensuring interfaces in compile-time for GetSupergroupMembersRequest.
var (
	_ bin.Encoder     = &GetSupergroupMembersRequest{}
	_ bin.Decoder     = &GetSupergroupMembersRequest{}
	_ bin.BareEncoder = &GetSupergroupMembersRequest{}
	_ bin.BareDecoder = &GetSupergroupMembersRequest{}
)

func (g *GetSupergroupMembersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.SupergroupID == 0) {
		return false
	}
	if !(g.Filter == nil) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetSupergroupMembersRequest) String() string {
	if g == nil {
		return "GetSupergroupMembersRequest(nil)"
	}
	type Alias GetSupergroupMembersRequest
	return fmt.Sprintf("GetSupergroupMembersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetSupergroupMembersRequest) TypeID() uint32 {
	return GetSupergroupMembersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetSupergroupMembersRequest) TypeName() string {
	return "getSupergroupMembers"
}

// TypeInfo returns info about TL type.
func (g *GetSupergroupMembersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getSupergroupMembers",
		ID:   GetSupergroupMembersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetSupergroupMembersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSupergroupMembers#ddf821c8 as nil")
	}
	b.PutID(GetSupergroupMembersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetSupergroupMembersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSupergroupMembers#ddf821c8 as nil")
	}
	b.PutInt53(g.SupergroupID)
	if g.Filter == nil {
		return fmt.Errorf("unable to encode getSupergroupMembers#ddf821c8: field filter is nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getSupergroupMembers#ddf821c8: field filter: %w", err)
	}
	b.PutInt32(g.Offset)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetSupergroupMembersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSupergroupMembers#ddf821c8 to nil")
	}
	if err := b.ConsumeID(GetSupergroupMembersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getSupergroupMembers#ddf821c8: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetSupergroupMembersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSupergroupMembers#ddf821c8 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getSupergroupMembers#ddf821c8: field supergroup_id: %w", err)
		}
		g.SupergroupID = value
	}
	{
		value, err := DecodeSupergroupMembersFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode getSupergroupMembers#ddf821c8: field filter: %w", err)
		}
		g.Filter = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getSupergroupMembers#ddf821c8: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getSupergroupMembers#ddf821c8: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetSupergroupMembersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getSupergroupMembers#ddf821c8 as nil")
	}
	b.ObjStart()
	b.PutID("getSupergroupMembers")
	b.Comma()
	b.FieldStart("supergroup_id")
	b.PutInt53(g.SupergroupID)
	b.Comma()
	b.FieldStart("filter")
	if g.Filter == nil {
		return fmt.Errorf("unable to encode getSupergroupMembers#ddf821c8: field filter is nil")
	}
	if err := g.Filter.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getSupergroupMembers#ddf821c8: field filter: %w", err)
	}
	b.Comma()
	b.FieldStart("offset")
	b.PutInt32(g.Offset)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetSupergroupMembersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getSupergroupMembers#ddf821c8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getSupergroupMembers"); err != nil {
				return fmt.Errorf("unable to decode getSupergroupMembers#ddf821c8: %w", err)
			}
		case "supergroup_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getSupergroupMembers#ddf821c8: field supergroup_id: %w", err)
			}
			g.SupergroupID = value
		case "filter":
			value, err := DecodeTDLibJSONSupergroupMembersFilter(b)
			if err != nil {
				return fmt.Errorf("unable to decode getSupergroupMembers#ddf821c8: field filter: %w", err)
			}
			g.Filter = value
		case "offset":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getSupergroupMembers#ddf821c8: field offset: %w", err)
			}
			g.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getSupergroupMembers#ddf821c8: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (g *GetSupergroupMembersRequest) GetSupergroupID() (value int64) {
	if g == nil {
		return
	}
	return g.SupergroupID
}

// GetFilter returns value of Filter field.
func (g *GetSupergroupMembersRequest) GetFilter() (value SupergroupMembersFilterClass) {
	if g == nil {
		return
	}
	return g.Filter
}

// GetOffset returns value of Offset field.
func (g *GetSupergroupMembersRequest) GetOffset() (value int32) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *GetSupergroupMembersRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetSupergroupMembers invokes method getSupergroupMembers#ddf821c8 returning error if any.
func (c *Client) GetSupergroupMembers(ctx context.Context, request *GetSupergroupMembersRequest) (*ChatMembers, error) {
	var result ChatMembers

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
