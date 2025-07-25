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

// GetPasswordStateRequest represents TL type `getPasswordState#f5957b78`.
type GetPasswordStateRequest struct {
}

// GetPasswordStateRequestTypeID is TL type id of GetPasswordStateRequest.
const GetPasswordStateRequestTypeID = 0xf5957b78

// Ensuring interfaces in compile-time for GetPasswordStateRequest.
var (
	_ bin.Encoder     = &GetPasswordStateRequest{}
	_ bin.Decoder     = &GetPasswordStateRequest{}
	_ bin.BareEncoder = &GetPasswordStateRequest{}
	_ bin.BareDecoder = &GetPasswordStateRequest{}
)

func (g *GetPasswordStateRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetPasswordStateRequest) String() string {
	if g == nil {
		return "GetPasswordStateRequest(nil)"
	}
	type Alias GetPasswordStateRequest
	return fmt.Sprintf("GetPasswordStateRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetPasswordStateRequest) TypeID() uint32 {
	return GetPasswordStateRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetPasswordStateRequest) TypeName() string {
	return "getPasswordState"
}

// TypeInfo returns info about TL type.
func (g *GetPasswordStateRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getPasswordState",
		ID:   GetPasswordStateRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetPasswordStateRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPasswordState#f5957b78 as nil")
	}
	b.PutID(GetPasswordStateRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetPasswordStateRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPasswordState#f5957b78 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetPasswordStateRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPasswordState#f5957b78 to nil")
	}
	if err := b.ConsumeID(GetPasswordStateRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getPasswordState#f5957b78: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetPasswordStateRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPasswordState#f5957b78 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetPasswordStateRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getPasswordState#f5957b78 as nil")
	}
	b.ObjStart()
	b.PutID("getPasswordState")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetPasswordStateRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getPasswordState#f5957b78 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getPasswordState"); err != nil {
				return fmt.Errorf("unable to decode getPasswordState#f5957b78: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPasswordState invokes method getPasswordState#f5957b78 returning error if any.
func (c *Client) GetPasswordState(ctx context.Context) (*PasswordState, error) {
	var result PasswordState

	request := &GetPasswordStateRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
