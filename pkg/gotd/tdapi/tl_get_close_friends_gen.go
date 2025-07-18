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

// GetCloseFriendsRequest represents TL type `getCloseFriends#a9d574ce`.
type GetCloseFriendsRequest struct {
}

// GetCloseFriendsRequestTypeID is TL type id of GetCloseFriendsRequest.
const GetCloseFriendsRequestTypeID = 0xa9d574ce

// Ensuring interfaces in compile-time for GetCloseFriendsRequest.
var (
	_ bin.Encoder     = &GetCloseFriendsRequest{}
	_ bin.Decoder     = &GetCloseFriendsRequest{}
	_ bin.BareEncoder = &GetCloseFriendsRequest{}
	_ bin.BareDecoder = &GetCloseFriendsRequest{}
)

func (g *GetCloseFriendsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetCloseFriendsRequest) String() string {
	if g == nil {
		return "GetCloseFriendsRequest(nil)"
	}
	type Alias GetCloseFriendsRequest
	return fmt.Sprintf("GetCloseFriendsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetCloseFriendsRequest) TypeID() uint32 {
	return GetCloseFriendsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetCloseFriendsRequest) TypeName() string {
	return "getCloseFriends"
}

// TypeInfo returns info about TL type.
func (g *GetCloseFriendsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getCloseFriends",
		ID:   GetCloseFriendsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetCloseFriendsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCloseFriends#a9d574ce as nil")
	}
	b.PutID(GetCloseFriendsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetCloseFriendsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCloseFriends#a9d574ce as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetCloseFriendsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCloseFriends#a9d574ce to nil")
	}
	if err := b.ConsumeID(GetCloseFriendsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getCloseFriends#a9d574ce: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetCloseFriendsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCloseFriends#a9d574ce to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetCloseFriendsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getCloseFriends#a9d574ce as nil")
	}
	b.ObjStart()
	b.PutID("getCloseFriends")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetCloseFriendsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getCloseFriends#a9d574ce to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getCloseFriends"); err != nil {
				return fmt.Errorf("unable to decode getCloseFriends#a9d574ce: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCloseFriends invokes method getCloseFriends#a9d574ce returning error if any.
func (c *Client) GetCloseFriends(ctx context.Context) (*Users, error) {
	var result Users

	request := &GetCloseFriendsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
