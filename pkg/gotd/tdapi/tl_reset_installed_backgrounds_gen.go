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

// ResetInstalledBackgroundsRequest represents TL type `resetInstalledBackgrounds#70540157`.
type ResetInstalledBackgroundsRequest struct {
}

// ResetInstalledBackgroundsRequestTypeID is TL type id of ResetInstalledBackgroundsRequest.
const ResetInstalledBackgroundsRequestTypeID = 0x70540157

// Ensuring interfaces in compile-time for ResetInstalledBackgroundsRequest.
var (
	_ bin.Encoder     = &ResetInstalledBackgroundsRequest{}
	_ bin.Decoder     = &ResetInstalledBackgroundsRequest{}
	_ bin.BareEncoder = &ResetInstalledBackgroundsRequest{}
	_ bin.BareDecoder = &ResetInstalledBackgroundsRequest{}
)

func (r *ResetInstalledBackgroundsRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ResetInstalledBackgroundsRequest) String() string {
	if r == nil {
		return "ResetInstalledBackgroundsRequest(nil)"
	}
	type Alias ResetInstalledBackgroundsRequest
	return fmt.Sprintf("ResetInstalledBackgroundsRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ResetInstalledBackgroundsRequest) TypeID() uint32 {
	return ResetInstalledBackgroundsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ResetInstalledBackgroundsRequest) TypeName() string {
	return "resetInstalledBackgrounds"
}

// TypeInfo returns info about TL type.
func (r *ResetInstalledBackgroundsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "resetInstalledBackgrounds",
		ID:   ResetInstalledBackgroundsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ResetInstalledBackgroundsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resetInstalledBackgrounds#70540157 as nil")
	}
	b.PutID(ResetInstalledBackgroundsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ResetInstalledBackgroundsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resetInstalledBackgrounds#70540157 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ResetInstalledBackgroundsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resetInstalledBackgrounds#70540157 to nil")
	}
	if err := b.ConsumeID(ResetInstalledBackgroundsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode resetInstalledBackgrounds#70540157: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ResetInstalledBackgroundsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resetInstalledBackgrounds#70540157 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ResetInstalledBackgroundsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode resetInstalledBackgrounds#70540157 as nil")
	}
	b.ObjStart()
	b.PutID("resetInstalledBackgrounds")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ResetInstalledBackgroundsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode resetInstalledBackgrounds#70540157 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("resetInstalledBackgrounds"); err != nil {
				return fmt.Errorf("unable to decode resetInstalledBackgrounds#70540157: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ResetInstalledBackgrounds invokes method resetInstalledBackgrounds#70540157 returning error if any.
func (c *Client) ResetInstalledBackgrounds(ctx context.Context) error {
	var ok Ok

	request := &ResetInstalledBackgroundsRequest{}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
