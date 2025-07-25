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

// GetUserSupportInfoRequest represents TL type `getUserSupportInfo#74a59305`.
type GetUserSupportInfoRequest struct {
	// User identifier
	UserID int64
}

// GetUserSupportInfoRequestTypeID is TL type id of GetUserSupportInfoRequest.
const GetUserSupportInfoRequestTypeID = 0x74a59305

// Ensuring interfaces in compile-time for GetUserSupportInfoRequest.
var (
	_ bin.Encoder     = &GetUserSupportInfoRequest{}
	_ bin.Decoder     = &GetUserSupportInfoRequest{}
	_ bin.BareEncoder = &GetUserSupportInfoRequest{}
	_ bin.BareDecoder = &GetUserSupportInfoRequest{}
)

func (g *GetUserSupportInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetUserSupportInfoRequest) String() string {
	if g == nil {
		return "GetUserSupportInfoRequest(nil)"
	}
	type Alias GetUserSupportInfoRequest
	return fmt.Sprintf("GetUserSupportInfoRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetUserSupportInfoRequest) TypeID() uint32 {
	return GetUserSupportInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetUserSupportInfoRequest) TypeName() string {
	return "getUserSupportInfo"
}

// TypeInfo returns info about TL type.
func (g *GetUserSupportInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getUserSupportInfo",
		ID:   GetUserSupportInfoRequestTypeID,
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
func (g *GetUserSupportInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserSupportInfo#74a59305 as nil")
	}
	b.PutID(GetUserSupportInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetUserSupportInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserSupportInfo#74a59305 as nil")
	}
	b.PutInt53(g.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetUserSupportInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getUserSupportInfo#74a59305 to nil")
	}
	if err := b.ConsumeID(GetUserSupportInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getUserSupportInfo#74a59305: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetUserSupportInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getUserSupportInfo#74a59305 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getUserSupportInfo#74a59305: field user_id: %w", err)
		}
		g.UserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetUserSupportInfoRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserSupportInfo#74a59305 as nil")
	}
	b.ObjStart()
	b.PutID("getUserSupportInfo")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(g.UserID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetUserSupportInfoRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getUserSupportInfo#74a59305 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getUserSupportInfo"); err != nil {
				return fmt.Errorf("unable to decode getUserSupportInfo#74a59305: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getUserSupportInfo#74a59305: field user_id: %w", err)
			}
			g.UserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (g *GetUserSupportInfoRequest) GetUserID() (value int64) {
	if g == nil {
		return
	}
	return g.UserID
}

// GetUserSupportInfo invokes method getUserSupportInfo#74a59305 returning error if any.
func (c *Client) GetUserSupportInfo(ctx context.Context, userid int64) (*UserSupportInfo, error) {
	var result UserSupportInfo

	request := &GetUserSupportInfoRequest{
		UserID: userid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
