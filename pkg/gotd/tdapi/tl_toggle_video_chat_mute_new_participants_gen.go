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

// ToggleVideoChatMuteNewParticipantsRequest represents TL type `toggleVideoChatMuteNewParticipants#3ad4c98c`.
type ToggleVideoChatMuteNewParticipantsRequest struct {
	// Group call identifier
	GroupCallID int32
	// New value of the mute_new_participants setting
	MuteNewParticipants bool
}

// ToggleVideoChatMuteNewParticipantsRequestTypeID is TL type id of ToggleVideoChatMuteNewParticipantsRequest.
const ToggleVideoChatMuteNewParticipantsRequestTypeID = 0x3ad4c98c

// Ensuring interfaces in compile-time for ToggleVideoChatMuteNewParticipantsRequest.
var (
	_ bin.Encoder     = &ToggleVideoChatMuteNewParticipantsRequest{}
	_ bin.Decoder     = &ToggleVideoChatMuteNewParticipantsRequest{}
	_ bin.BareEncoder = &ToggleVideoChatMuteNewParticipantsRequest{}
	_ bin.BareDecoder = &ToggleVideoChatMuteNewParticipantsRequest{}
)

func (t *ToggleVideoChatMuteNewParticipantsRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.GroupCallID == 0) {
		return false
	}
	if !(t.MuteNewParticipants == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleVideoChatMuteNewParticipantsRequest) String() string {
	if t == nil {
		return "ToggleVideoChatMuteNewParticipantsRequest(nil)"
	}
	type Alias ToggleVideoChatMuteNewParticipantsRequest
	return fmt.Sprintf("ToggleVideoChatMuteNewParticipantsRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleVideoChatMuteNewParticipantsRequest) TypeID() uint32 {
	return ToggleVideoChatMuteNewParticipantsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleVideoChatMuteNewParticipantsRequest) TypeName() string {
	return "toggleVideoChatMuteNewParticipants"
}

// TypeInfo returns info about TL type.
func (t *ToggleVideoChatMuteNewParticipantsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleVideoChatMuteNewParticipants",
		ID:   ToggleVideoChatMuteNewParticipantsRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
		{
			Name:       "MuteNewParticipants",
			SchemaName: "mute_new_participants",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleVideoChatMuteNewParticipantsRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleVideoChatMuteNewParticipants#3ad4c98c as nil")
	}
	b.PutID(ToggleVideoChatMuteNewParticipantsRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleVideoChatMuteNewParticipantsRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleVideoChatMuteNewParticipants#3ad4c98c as nil")
	}
	b.PutInt32(t.GroupCallID)
	b.PutBool(t.MuteNewParticipants)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleVideoChatMuteNewParticipantsRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleVideoChatMuteNewParticipants#3ad4c98c to nil")
	}
	if err := b.ConsumeID(ToggleVideoChatMuteNewParticipantsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleVideoChatMuteNewParticipants#3ad4c98c: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleVideoChatMuteNewParticipantsRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleVideoChatMuteNewParticipants#3ad4c98c to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode toggleVideoChatMuteNewParticipants#3ad4c98c: field group_call_id: %w", err)
		}
		t.GroupCallID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleVideoChatMuteNewParticipants#3ad4c98c: field mute_new_participants: %w", err)
		}
		t.MuteNewParticipants = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleVideoChatMuteNewParticipantsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleVideoChatMuteNewParticipants#3ad4c98c as nil")
	}
	b.ObjStart()
	b.PutID("toggleVideoChatMuteNewParticipants")
	b.Comma()
	b.FieldStart("group_call_id")
	b.PutInt32(t.GroupCallID)
	b.Comma()
	b.FieldStart("mute_new_participants")
	b.PutBool(t.MuteNewParticipants)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleVideoChatMuteNewParticipantsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleVideoChatMuteNewParticipants#3ad4c98c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleVideoChatMuteNewParticipants"); err != nil {
				return fmt.Errorf("unable to decode toggleVideoChatMuteNewParticipants#3ad4c98c: %w", err)
			}
		case "group_call_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode toggleVideoChatMuteNewParticipants#3ad4c98c: field group_call_id: %w", err)
			}
			t.GroupCallID = value
		case "mute_new_participants":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleVideoChatMuteNewParticipants#3ad4c98c: field mute_new_participants: %w", err)
			}
			t.MuteNewParticipants = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGroupCallID returns value of GroupCallID field.
func (t *ToggleVideoChatMuteNewParticipantsRequest) GetGroupCallID() (value int32) {
	if t == nil {
		return
	}
	return t.GroupCallID
}

// GetMuteNewParticipants returns value of MuteNewParticipants field.
func (t *ToggleVideoChatMuteNewParticipantsRequest) GetMuteNewParticipants() (value bool) {
	if t == nil {
		return
	}
	return t.MuteNewParticipants
}

// ToggleVideoChatMuteNewParticipants invokes method toggleVideoChatMuteNewParticipants#3ad4c98c returning error if any.
func (c *Client) ToggleVideoChatMuteNewParticipants(ctx context.Context, request *ToggleVideoChatMuteNewParticipantsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
