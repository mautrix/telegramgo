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

// DisableAllSupergroupUsernamesRequest represents TL type `disableAllSupergroupUsernames#3246f5b0`.
type DisableAllSupergroupUsernamesRequest struct {
	// Identifier of the supergroup or channel
	SupergroupID int64
}

// DisableAllSupergroupUsernamesRequestTypeID is TL type id of DisableAllSupergroupUsernamesRequest.
const DisableAllSupergroupUsernamesRequestTypeID = 0x3246f5b0

// Ensuring interfaces in compile-time for DisableAllSupergroupUsernamesRequest.
var (
	_ bin.Encoder     = &DisableAllSupergroupUsernamesRequest{}
	_ bin.Decoder     = &DisableAllSupergroupUsernamesRequest{}
	_ bin.BareEncoder = &DisableAllSupergroupUsernamesRequest{}
	_ bin.BareDecoder = &DisableAllSupergroupUsernamesRequest{}
)

func (d *DisableAllSupergroupUsernamesRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.SupergroupID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DisableAllSupergroupUsernamesRequest) String() string {
	if d == nil {
		return "DisableAllSupergroupUsernamesRequest(nil)"
	}
	type Alias DisableAllSupergroupUsernamesRequest
	return fmt.Sprintf("DisableAllSupergroupUsernamesRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DisableAllSupergroupUsernamesRequest) TypeID() uint32 {
	return DisableAllSupergroupUsernamesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DisableAllSupergroupUsernamesRequest) TypeName() string {
	return "disableAllSupergroupUsernames"
}

// TypeInfo returns info about TL type.
func (d *DisableAllSupergroupUsernamesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "disableAllSupergroupUsernames",
		ID:   DisableAllSupergroupUsernamesRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DisableAllSupergroupUsernamesRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode disableAllSupergroupUsernames#3246f5b0 as nil")
	}
	b.PutID(DisableAllSupergroupUsernamesRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DisableAllSupergroupUsernamesRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode disableAllSupergroupUsernames#3246f5b0 as nil")
	}
	b.PutInt53(d.SupergroupID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DisableAllSupergroupUsernamesRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode disableAllSupergroupUsernames#3246f5b0 to nil")
	}
	if err := b.ConsumeID(DisableAllSupergroupUsernamesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode disableAllSupergroupUsernames#3246f5b0: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DisableAllSupergroupUsernamesRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode disableAllSupergroupUsernames#3246f5b0 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode disableAllSupergroupUsernames#3246f5b0: field supergroup_id: %w", err)
		}
		d.SupergroupID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DisableAllSupergroupUsernamesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode disableAllSupergroupUsernames#3246f5b0 as nil")
	}
	b.ObjStart()
	b.PutID("disableAllSupergroupUsernames")
	b.Comma()
	b.FieldStart("supergroup_id")
	b.PutInt53(d.SupergroupID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DisableAllSupergroupUsernamesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode disableAllSupergroupUsernames#3246f5b0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("disableAllSupergroupUsernames"); err != nil {
				return fmt.Errorf("unable to decode disableAllSupergroupUsernames#3246f5b0: %w", err)
			}
		case "supergroup_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode disableAllSupergroupUsernames#3246f5b0: field supergroup_id: %w", err)
			}
			d.SupergroupID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (d *DisableAllSupergroupUsernamesRequest) GetSupergroupID() (value int64) {
	if d == nil {
		return
	}
	return d.SupergroupID
}

// DisableAllSupergroupUsernames invokes method disableAllSupergroupUsernames#3246f5b0 returning error if any.
func (c *Client) DisableAllSupergroupUsernames(ctx context.Context, supergroupid int64) error {
	var ok Ok

	request := &DisableAllSupergroupUsernamesRequest{
		SupergroupID: supergroupid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
