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

// GetChatBoostFeaturesRequest represents TL type `getChatBoostFeatures#e8c128a0`.
type GetChatBoostFeaturesRequest struct {
	// Pass true to get the list of features for channels; pass false to get the list of
	// features for supergroups
	IsChannel bool
}

// GetChatBoostFeaturesRequestTypeID is TL type id of GetChatBoostFeaturesRequest.
const GetChatBoostFeaturesRequestTypeID = 0xe8c128a0

// Ensuring interfaces in compile-time for GetChatBoostFeaturesRequest.
var (
	_ bin.Encoder     = &GetChatBoostFeaturesRequest{}
	_ bin.Decoder     = &GetChatBoostFeaturesRequest{}
	_ bin.BareEncoder = &GetChatBoostFeaturesRequest{}
	_ bin.BareDecoder = &GetChatBoostFeaturesRequest{}
)

func (g *GetChatBoostFeaturesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.IsChannel == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatBoostFeaturesRequest) String() string {
	if g == nil {
		return "GetChatBoostFeaturesRequest(nil)"
	}
	type Alias GetChatBoostFeaturesRequest
	return fmt.Sprintf("GetChatBoostFeaturesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatBoostFeaturesRequest) TypeID() uint32 {
	return GetChatBoostFeaturesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatBoostFeaturesRequest) TypeName() string {
	return "getChatBoostFeatures"
}

// TypeInfo returns info about TL type.
func (g *GetChatBoostFeaturesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatBoostFeatures",
		ID:   GetChatBoostFeaturesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "IsChannel",
			SchemaName: "is_channel",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatBoostFeaturesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatBoostFeatures#e8c128a0 as nil")
	}
	b.PutID(GetChatBoostFeaturesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatBoostFeaturesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatBoostFeatures#e8c128a0 as nil")
	}
	b.PutBool(g.IsChannel)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatBoostFeaturesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatBoostFeatures#e8c128a0 to nil")
	}
	if err := b.ConsumeID(GetChatBoostFeaturesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatBoostFeatures#e8c128a0: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatBoostFeaturesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatBoostFeatures#e8c128a0 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getChatBoostFeatures#e8c128a0: field is_channel: %w", err)
		}
		g.IsChannel = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatBoostFeaturesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatBoostFeatures#e8c128a0 as nil")
	}
	b.ObjStart()
	b.PutID("getChatBoostFeatures")
	b.Comma()
	b.FieldStart("is_channel")
	b.PutBool(g.IsChannel)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatBoostFeaturesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatBoostFeatures#e8c128a0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatBoostFeatures"); err != nil {
				return fmt.Errorf("unable to decode getChatBoostFeatures#e8c128a0: %w", err)
			}
		case "is_channel":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getChatBoostFeatures#e8c128a0: field is_channel: %w", err)
			}
			g.IsChannel = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetIsChannel returns value of IsChannel field.
func (g *GetChatBoostFeaturesRequest) GetIsChannel() (value bool) {
	if g == nil {
		return
	}
	return g.IsChannel
}

// GetChatBoostFeatures invokes method getChatBoostFeatures#e8c128a0 returning error if any.
func (c *Client) GetChatBoostFeatures(ctx context.Context, ischannel bool) (*ChatBoostFeatures, error) {
	var result ChatBoostFeatures

	request := &GetChatBoostFeaturesRequest{
		IsChannel: ischannel,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
