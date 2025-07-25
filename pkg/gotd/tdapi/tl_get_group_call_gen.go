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

// GetGroupCallRequest represents TL type `getGroupCall#5787668e`.
type GetGroupCallRequest struct {
	// Group call identifier
	GroupCallID int32
}

// GetGroupCallRequestTypeID is TL type id of GetGroupCallRequest.
const GetGroupCallRequestTypeID = 0x5787668e

// Ensuring interfaces in compile-time for GetGroupCallRequest.
var (
	_ bin.Encoder     = &GetGroupCallRequest{}
	_ bin.Decoder     = &GetGroupCallRequest{}
	_ bin.BareEncoder = &GetGroupCallRequest{}
	_ bin.BareDecoder = &GetGroupCallRequest{}
)

func (g *GetGroupCallRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.GroupCallID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetGroupCallRequest) String() string {
	if g == nil {
		return "GetGroupCallRequest(nil)"
	}
	type Alias GetGroupCallRequest
	return fmt.Sprintf("GetGroupCallRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetGroupCallRequest) TypeID() uint32 {
	return GetGroupCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetGroupCallRequest) TypeName() string {
	return "getGroupCall"
}

// TypeInfo returns info about TL type.
func (g *GetGroupCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getGroupCall",
		ID:   GetGroupCallRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetGroupCallRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupCall#5787668e as nil")
	}
	b.PutID(GetGroupCallRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetGroupCallRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupCall#5787668e as nil")
	}
	b.PutInt32(g.GroupCallID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetGroupCallRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupCall#5787668e to nil")
	}
	if err := b.ConsumeID(GetGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getGroupCall#5787668e: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetGroupCallRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupCall#5787668e to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getGroupCall#5787668e: field group_call_id: %w", err)
		}
		g.GroupCallID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetGroupCallRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupCall#5787668e as nil")
	}
	b.ObjStart()
	b.PutID("getGroupCall")
	b.Comma()
	b.FieldStart("group_call_id")
	b.PutInt32(g.GroupCallID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetGroupCallRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupCall#5787668e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getGroupCall"); err != nil {
				return fmt.Errorf("unable to decode getGroupCall#5787668e: %w", err)
			}
		case "group_call_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getGroupCall#5787668e: field group_call_id: %w", err)
			}
			g.GroupCallID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGroupCallID returns value of GroupCallID field.
func (g *GetGroupCallRequest) GetGroupCallID() (value int32) {
	if g == nil {
		return
	}
	return g.GroupCallID
}

// GetGroupCall invokes method getGroupCall#5787668e returning error if any.
func (c *Client) GetGroupCall(ctx context.Context, groupcallid int32) (*GroupCall, error) {
	var result GroupCall

	request := &GetGroupCallRequest{
		GroupCallID: groupcallid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
