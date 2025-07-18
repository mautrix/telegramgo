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

// CancelPasswordResetRequest represents TL type `cancelPasswordReset#38127462`.
type CancelPasswordResetRequest struct {
}

// CancelPasswordResetRequestTypeID is TL type id of CancelPasswordResetRequest.
const CancelPasswordResetRequestTypeID = 0x38127462

// Ensuring interfaces in compile-time for CancelPasswordResetRequest.
var (
	_ bin.Encoder     = &CancelPasswordResetRequest{}
	_ bin.Decoder     = &CancelPasswordResetRequest{}
	_ bin.BareEncoder = &CancelPasswordResetRequest{}
	_ bin.BareDecoder = &CancelPasswordResetRequest{}
)

func (c *CancelPasswordResetRequest) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *CancelPasswordResetRequest) String() string {
	if c == nil {
		return "CancelPasswordResetRequest(nil)"
	}
	type Alias CancelPasswordResetRequest
	return fmt.Sprintf("CancelPasswordResetRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CancelPasswordResetRequest) TypeID() uint32 {
	return CancelPasswordResetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CancelPasswordResetRequest) TypeName() string {
	return "cancelPasswordReset"
}

// TypeInfo returns info about TL type.
func (c *CancelPasswordResetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "cancelPasswordReset",
		ID:   CancelPasswordResetRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *CancelPasswordResetRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode cancelPasswordReset#38127462 as nil")
	}
	b.PutID(CancelPasswordResetRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CancelPasswordResetRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode cancelPasswordReset#38127462 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CancelPasswordResetRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode cancelPasswordReset#38127462 to nil")
	}
	if err := b.ConsumeID(CancelPasswordResetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode cancelPasswordReset#38127462: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CancelPasswordResetRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode cancelPasswordReset#38127462 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CancelPasswordResetRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode cancelPasswordReset#38127462 as nil")
	}
	b.ObjStart()
	b.PutID("cancelPasswordReset")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CancelPasswordResetRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode cancelPasswordReset#38127462 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("cancelPasswordReset"); err != nil {
				return fmt.Errorf("unable to decode cancelPasswordReset#38127462: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// CancelPasswordReset invokes method cancelPasswordReset#38127462 returning error if any.
func (c *Client) CancelPasswordReset(ctx context.Context) error {
	var ok Ok

	request := &CancelPasswordResetRequest{}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
