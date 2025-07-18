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

// ToggleAllDownloadsArePausedRequest represents TL type `toggleAllDownloadsArePaused#4a989002`.
type ToggleAllDownloadsArePausedRequest struct {
	// Pass true to pause all downloads; pass false to unpause them
	ArePaused bool
}

// ToggleAllDownloadsArePausedRequestTypeID is TL type id of ToggleAllDownloadsArePausedRequest.
const ToggleAllDownloadsArePausedRequestTypeID = 0x4a989002

// Ensuring interfaces in compile-time for ToggleAllDownloadsArePausedRequest.
var (
	_ bin.Encoder     = &ToggleAllDownloadsArePausedRequest{}
	_ bin.Decoder     = &ToggleAllDownloadsArePausedRequest{}
	_ bin.BareEncoder = &ToggleAllDownloadsArePausedRequest{}
	_ bin.BareDecoder = &ToggleAllDownloadsArePausedRequest{}
)

func (t *ToggleAllDownloadsArePausedRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.ArePaused == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleAllDownloadsArePausedRequest) String() string {
	if t == nil {
		return "ToggleAllDownloadsArePausedRequest(nil)"
	}
	type Alias ToggleAllDownloadsArePausedRequest
	return fmt.Sprintf("ToggleAllDownloadsArePausedRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleAllDownloadsArePausedRequest) TypeID() uint32 {
	return ToggleAllDownloadsArePausedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleAllDownloadsArePausedRequest) TypeName() string {
	return "toggleAllDownloadsArePaused"
}

// TypeInfo returns info about TL type.
func (t *ToggleAllDownloadsArePausedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleAllDownloadsArePaused",
		ID:   ToggleAllDownloadsArePausedRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ArePaused",
			SchemaName: "are_paused",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleAllDownloadsArePausedRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleAllDownloadsArePaused#4a989002 as nil")
	}
	b.PutID(ToggleAllDownloadsArePausedRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleAllDownloadsArePausedRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleAllDownloadsArePaused#4a989002 as nil")
	}
	b.PutBool(t.ArePaused)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleAllDownloadsArePausedRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleAllDownloadsArePaused#4a989002 to nil")
	}
	if err := b.ConsumeID(ToggleAllDownloadsArePausedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleAllDownloadsArePaused#4a989002: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleAllDownloadsArePausedRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleAllDownloadsArePaused#4a989002 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleAllDownloadsArePaused#4a989002: field are_paused: %w", err)
		}
		t.ArePaused = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleAllDownloadsArePausedRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleAllDownloadsArePaused#4a989002 as nil")
	}
	b.ObjStart()
	b.PutID("toggleAllDownloadsArePaused")
	b.Comma()
	b.FieldStart("are_paused")
	b.PutBool(t.ArePaused)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleAllDownloadsArePausedRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleAllDownloadsArePaused#4a989002 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleAllDownloadsArePaused"); err != nil {
				return fmt.Errorf("unable to decode toggleAllDownloadsArePaused#4a989002: %w", err)
			}
		case "are_paused":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleAllDownloadsArePaused#4a989002: field are_paused: %w", err)
			}
			t.ArePaused = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetArePaused returns value of ArePaused field.
func (t *ToggleAllDownloadsArePausedRequest) GetArePaused() (value bool) {
	if t == nil {
		return
	}
	return t.ArePaused
}

// ToggleAllDownloadsArePaused invokes method toggleAllDownloadsArePaused#4a989002 returning error if any.
func (c *Client) ToggleAllDownloadsArePaused(ctx context.Context, arepaused bool) error {
	var ok Ok

	request := &ToggleAllDownloadsArePausedRequest{
		ArePaused: arepaused,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
