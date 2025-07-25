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

// GetStarGiftPaymentOptionsRequest represents TL type `getStarGiftPaymentOptions#e22760e3`.
type GetStarGiftPaymentOptionsRequest struct {
	// Identifier of the user that will receive Telegram Stars; pass 0 to get options for an
	// unspecified user
	UserID int64
}

// GetStarGiftPaymentOptionsRequestTypeID is TL type id of GetStarGiftPaymentOptionsRequest.
const GetStarGiftPaymentOptionsRequestTypeID = 0xe22760e3

// Ensuring interfaces in compile-time for GetStarGiftPaymentOptionsRequest.
var (
	_ bin.Encoder     = &GetStarGiftPaymentOptionsRequest{}
	_ bin.Decoder     = &GetStarGiftPaymentOptionsRequest{}
	_ bin.BareEncoder = &GetStarGiftPaymentOptionsRequest{}
	_ bin.BareDecoder = &GetStarGiftPaymentOptionsRequest{}
)

func (g *GetStarGiftPaymentOptionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetStarGiftPaymentOptionsRequest) String() string {
	if g == nil {
		return "GetStarGiftPaymentOptionsRequest(nil)"
	}
	type Alias GetStarGiftPaymentOptionsRequest
	return fmt.Sprintf("GetStarGiftPaymentOptionsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetStarGiftPaymentOptionsRequest) TypeID() uint32 {
	return GetStarGiftPaymentOptionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetStarGiftPaymentOptionsRequest) TypeName() string {
	return "getStarGiftPaymentOptions"
}

// TypeInfo returns info about TL type.
func (g *GetStarGiftPaymentOptionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getStarGiftPaymentOptions",
		ID:   GetStarGiftPaymentOptionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetStarGiftPaymentOptionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStarGiftPaymentOptions#e22760e3 as nil")
	}
	b.PutID(GetStarGiftPaymentOptionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetStarGiftPaymentOptionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStarGiftPaymentOptions#e22760e3 as nil")
	}
	b.PutInt53(g.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetStarGiftPaymentOptionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStarGiftPaymentOptions#e22760e3 to nil")
	}
	if err := b.ConsumeID(GetStarGiftPaymentOptionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getStarGiftPaymentOptions#e22760e3: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetStarGiftPaymentOptionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStarGiftPaymentOptions#e22760e3 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getStarGiftPaymentOptions#e22760e3: field user_id: %w", err)
		}
		g.UserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetStarGiftPaymentOptionsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getStarGiftPaymentOptions#e22760e3 as nil")
	}
	b.ObjStart()
	b.PutID("getStarGiftPaymentOptions")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(g.UserID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetStarGiftPaymentOptionsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getStarGiftPaymentOptions#e22760e3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getStarGiftPaymentOptions"); err != nil {
				return fmt.Errorf("unable to decode getStarGiftPaymentOptions#e22760e3: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getStarGiftPaymentOptions#e22760e3: field user_id: %w", err)
			}
			g.UserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (g *GetStarGiftPaymentOptionsRequest) GetUserID() (value int64) {
	if g == nil {
		return
	}
	return g.UserID
}

// GetStarGiftPaymentOptions invokes method getStarGiftPaymentOptions#e22760e3 returning error if any.
func (c *Client) GetStarGiftPaymentOptions(ctx context.Context, userid int64) (*StarPaymentOptions, error) {
	var result StarPaymentOptions

	request := &GetStarGiftPaymentOptionsRequest{
		UserID: userid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
