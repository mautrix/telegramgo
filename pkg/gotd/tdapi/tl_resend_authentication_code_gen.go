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

// ResendAuthenticationCodeRequest represents TL type `resendAuthenticationCode#a630bbb8`.
type ResendAuthenticationCodeRequest struct {
	// Reason of code resending; pass null if unknown
	Reason ResendCodeReasonClass
}

// ResendAuthenticationCodeRequestTypeID is TL type id of ResendAuthenticationCodeRequest.
const ResendAuthenticationCodeRequestTypeID = 0xa630bbb8

// Ensuring interfaces in compile-time for ResendAuthenticationCodeRequest.
var (
	_ bin.Encoder     = &ResendAuthenticationCodeRequest{}
	_ bin.Decoder     = &ResendAuthenticationCodeRequest{}
	_ bin.BareEncoder = &ResendAuthenticationCodeRequest{}
	_ bin.BareDecoder = &ResendAuthenticationCodeRequest{}
)

func (r *ResendAuthenticationCodeRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Reason == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ResendAuthenticationCodeRequest) String() string {
	if r == nil {
		return "ResendAuthenticationCodeRequest(nil)"
	}
	type Alias ResendAuthenticationCodeRequest
	return fmt.Sprintf("ResendAuthenticationCodeRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ResendAuthenticationCodeRequest) TypeID() uint32 {
	return ResendAuthenticationCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ResendAuthenticationCodeRequest) TypeName() string {
	return "resendAuthenticationCode"
}

// TypeInfo returns info about TL type.
func (r *ResendAuthenticationCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "resendAuthenticationCode",
		ID:   ResendAuthenticationCodeRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Reason",
			SchemaName: "reason",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ResendAuthenticationCodeRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resendAuthenticationCode#a630bbb8 as nil")
	}
	b.PutID(ResendAuthenticationCodeRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ResendAuthenticationCodeRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resendAuthenticationCode#a630bbb8 as nil")
	}
	if r.Reason == nil {
		return fmt.Errorf("unable to encode resendAuthenticationCode#a630bbb8: field reason is nil")
	}
	if err := r.Reason.Encode(b); err != nil {
		return fmt.Errorf("unable to encode resendAuthenticationCode#a630bbb8: field reason: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ResendAuthenticationCodeRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resendAuthenticationCode#a630bbb8 to nil")
	}
	if err := b.ConsumeID(ResendAuthenticationCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode resendAuthenticationCode#a630bbb8: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ResendAuthenticationCodeRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resendAuthenticationCode#a630bbb8 to nil")
	}
	{
		value, err := DecodeResendCodeReason(b)
		if err != nil {
			return fmt.Errorf("unable to decode resendAuthenticationCode#a630bbb8: field reason: %w", err)
		}
		r.Reason = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ResendAuthenticationCodeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode resendAuthenticationCode#a630bbb8 as nil")
	}
	b.ObjStart()
	b.PutID("resendAuthenticationCode")
	b.Comma()
	b.FieldStart("reason")
	if r.Reason == nil {
		return fmt.Errorf("unable to encode resendAuthenticationCode#a630bbb8: field reason is nil")
	}
	if err := r.Reason.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode resendAuthenticationCode#a630bbb8: field reason: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ResendAuthenticationCodeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode resendAuthenticationCode#a630bbb8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("resendAuthenticationCode"); err != nil {
				return fmt.Errorf("unable to decode resendAuthenticationCode#a630bbb8: %w", err)
			}
		case "reason":
			value, err := DecodeTDLibJSONResendCodeReason(b)
			if err != nil {
				return fmt.Errorf("unable to decode resendAuthenticationCode#a630bbb8: field reason: %w", err)
			}
			r.Reason = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetReason returns value of Reason field.
func (r *ResendAuthenticationCodeRequest) GetReason() (value ResendCodeReasonClass) {
	if r == nil {
		return
	}
	return r.Reason
}

// ResendAuthenticationCode invokes method resendAuthenticationCode#a630bbb8 returning error if any.
func (c *Client) ResendAuthenticationCode(ctx context.Context, reason ResendCodeReasonClass) error {
	var ok Ok

	request := &ResendAuthenticationCodeRequest{
		Reason: reason,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
