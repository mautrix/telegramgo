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

// GetGroupCallParticipantsRequest represents TL type `getGroupCallParticipants#766b3cc2`.
type GetGroupCallParticipantsRequest struct {
	// The group call which participants will be returned
	InputGroupCall InputGroupCallClass
	// The maximum number of participants to return; must be positive
	Limit int32
}

// GetGroupCallParticipantsRequestTypeID is TL type id of GetGroupCallParticipantsRequest.
const GetGroupCallParticipantsRequestTypeID = 0x766b3cc2

// Ensuring interfaces in compile-time for GetGroupCallParticipantsRequest.
var (
	_ bin.Encoder     = &GetGroupCallParticipantsRequest{}
	_ bin.Decoder     = &GetGroupCallParticipantsRequest{}
	_ bin.BareEncoder = &GetGroupCallParticipantsRequest{}
	_ bin.BareDecoder = &GetGroupCallParticipantsRequest{}
)

func (g *GetGroupCallParticipantsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.InputGroupCall == nil) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetGroupCallParticipantsRequest) String() string {
	if g == nil {
		return "GetGroupCallParticipantsRequest(nil)"
	}
	type Alias GetGroupCallParticipantsRequest
	return fmt.Sprintf("GetGroupCallParticipantsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetGroupCallParticipantsRequest) TypeID() uint32 {
	return GetGroupCallParticipantsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetGroupCallParticipantsRequest) TypeName() string {
	return "getGroupCallParticipants"
}

// TypeInfo returns info about TL type.
func (g *GetGroupCallParticipantsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getGroupCallParticipants",
		ID:   GetGroupCallParticipantsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InputGroupCall",
			SchemaName: "input_group_call",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetGroupCallParticipantsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupCallParticipants#766b3cc2 as nil")
	}
	b.PutID(GetGroupCallParticipantsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetGroupCallParticipantsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupCallParticipants#766b3cc2 as nil")
	}
	if g.InputGroupCall == nil {
		return fmt.Errorf("unable to encode getGroupCallParticipants#766b3cc2: field input_group_call is nil")
	}
	if err := g.InputGroupCall.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getGroupCallParticipants#766b3cc2: field input_group_call: %w", err)
	}
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetGroupCallParticipantsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupCallParticipants#766b3cc2 to nil")
	}
	if err := b.ConsumeID(GetGroupCallParticipantsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getGroupCallParticipants#766b3cc2: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetGroupCallParticipantsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupCallParticipants#766b3cc2 to nil")
	}
	{
		value, err := DecodeInputGroupCall(b)
		if err != nil {
			return fmt.Errorf("unable to decode getGroupCallParticipants#766b3cc2: field input_group_call: %w", err)
		}
		g.InputGroupCall = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getGroupCallParticipants#766b3cc2: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetGroupCallParticipantsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupCallParticipants#766b3cc2 as nil")
	}
	b.ObjStart()
	b.PutID("getGroupCallParticipants")
	b.Comma()
	b.FieldStart("input_group_call")
	if g.InputGroupCall == nil {
		return fmt.Errorf("unable to encode getGroupCallParticipants#766b3cc2: field input_group_call is nil")
	}
	if err := g.InputGroupCall.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getGroupCallParticipants#766b3cc2: field input_group_call: %w", err)
	}
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetGroupCallParticipantsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupCallParticipants#766b3cc2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getGroupCallParticipants"); err != nil {
				return fmt.Errorf("unable to decode getGroupCallParticipants#766b3cc2: %w", err)
			}
		case "input_group_call":
			value, err := DecodeTDLibJSONInputGroupCall(b)
			if err != nil {
				return fmt.Errorf("unable to decode getGroupCallParticipants#766b3cc2: field input_group_call: %w", err)
			}
			g.InputGroupCall = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getGroupCallParticipants#766b3cc2: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInputGroupCall returns value of InputGroupCall field.
func (g *GetGroupCallParticipantsRequest) GetInputGroupCall() (value InputGroupCallClass) {
	if g == nil {
		return
	}
	return g.InputGroupCall
}

// GetLimit returns value of Limit field.
func (g *GetGroupCallParticipantsRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetGroupCallParticipants invokes method getGroupCallParticipants#766b3cc2 returning error if any.
func (c *Client) GetGroupCallParticipants(ctx context.Context, request *GetGroupCallParticipantsRequest) (*GroupCallParticipants, error) {
	var result GroupCallParticipants

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
