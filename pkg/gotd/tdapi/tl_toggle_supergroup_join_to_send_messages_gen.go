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

// ToggleSupergroupJoinToSendMessagesRequest represents TL type `toggleSupergroupJoinToSendMessages#f5268e0e`.
type ToggleSupergroupJoinToSendMessagesRequest struct {
	// Identifier of the supergroup that isn't a broadcast group
	SupergroupID int64
	// New value of join_to_send_messages
	JoinToSendMessages bool
}

// ToggleSupergroupJoinToSendMessagesRequestTypeID is TL type id of ToggleSupergroupJoinToSendMessagesRequest.
const ToggleSupergroupJoinToSendMessagesRequestTypeID = 0xf5268e0e

// Ensuring interfaces in compile-time for ToggleSupergroupJoinToSendMessagesRequest.
var (
	_ bin.Encoder     = &ToggleSupergroupJoinToSendMessagesRequest{}
	_ bin.Decoder     = &ToggleSupergroupJoinToSendMessagesRequest{}
	_ bin.BareEncoder = &ToggleSupergroupJoinToSendMessagesRequest{}
	_ bin.BareDecoder = &ToggleSupergroupJoinToSendMessagesRequest{}
)

func (t *ToggleSupergroupJoinToSendMessagesRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.SupergroupID == 0) {
		return false
	}
	if !(t.JoinToSendMessages == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleSupergroupJoinToSendMessagesRequest) String() string {
	if t == nil {
		return "ToggleSupergroupJoinToSendMessagesRequest(nil)"
	}
	type Alias ToggleSupergroupJoinToSendMessagesRequest
	return fmt.Sprintf("ToggleSupergroupJoinToSendMessagesRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleSupergroupJoinToSendMessagesRequest) TypeID() uint32 {
	return ToggleSupergroupJoinToSendMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleSupergroupJoinToSendMessagesRequest) TypeName() string {
	return "toggleSupergroupJoinToSendMessages"
}

// TypeInfo returns info about TL type.
func (t *ToggleSupergroupJoinToSendMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleSupergroupJoinToSendMessages",
		ID:   ToggleSupergroupJoinToSendMessagesRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
		{
			Name:       "JoinToSendMessages",
			SchemaName: "join_to_send_messages",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleSupergroupJoinToSendMessagesRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupJoinToSendMessages#f5268e0e as nil")
	}
	b.PutID(ToggleSupergroupJoinToSendMessagesRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleSupergroupJoinToSendMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupJoinToSendMessages#f5268e0e as nil")
	}
	b.PutInt53(t.SupergroupID)
	b.PutBool(t.JoinToSendMessages)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleSupergroupJoinToSendMessagesRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupJoinToSendMessages#f5268e0e to nil")
	}
	if err := b.ConsumeID(ToggleSupergroupJoinToSendMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleSupergroupJoinToSendMessages#f5268e0e: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleSupergroupJoinToSendMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupJoinToSendMessages#f5268e0e to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode toggleSupergroupJoinToSendMessages#f5268e0e: field supergroup_id: %w", err)
		}
		t.SupergroupID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleSupergroupJoinToSendMessages#f5268e0e: field join_to_send_messages: %w", err)
		}
		t.JoinToSendMessages = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleSupergroupJoinToSendMessagesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupJoinToSendMessages#f5268e0e as nil")
	}
	b.ObjStart()
	b.PutID("toggleSupergroupJoinToSendMessages")
	b.Comma()
	b.FieldStart("supergroup_id")
	b.PutInt53(t.SupergroupID)
	b.Comma()
	b.FieldStart("join_to_send_messages")
	b.PutBool(t.JoinToSendMessages)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleSupergroupJoinToSendMessagesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupJoinToSendMessages#f5268e0e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleSupergroupJoinToSendMessages"); err != nil {
				return fmt.Errorf("unable to decode toggleSupergroupJoinToSendMessages#f5268e0e: %w", err)
			}
		case "supergroup_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode toggleSupergroupJoinToSendMessages#f5268e0e: field supergroup_id: %w", err)
			}
			t.SupergroupID = value
		case "join_to_send_messages":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleSupergroupJoinToSendMessages#f5268e0e: field join_to_send_messages: %w", err)
			}
			t.JoinToSendMessages = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (t *ToggleSupergroupJoinToSendMessagesRequest) GetSupergroupID() (value int64) {
	if t == nil {
		return
	}
	return t.SupergroupID
}

// GetJoinToSendMessages returns value of JoinToSendMessages field.
func (t *ToggleSupergroupJoinToSendMessagesRequest) GetJoinToSendMessages() (value bool) {
	if t == nil {
		return
	}
	return t.JoinToSendMessages
}

// ToggleSupergroupJoinToSendMessages invokes method toggleSupergroupJoinToSendMessages#f5268e0e returning error if any.
func (c *Client) ToggleSupergroupJoinToSendMessages(ctx context.Context, request *ToggleSupergroupJoinToSendMessagesRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
