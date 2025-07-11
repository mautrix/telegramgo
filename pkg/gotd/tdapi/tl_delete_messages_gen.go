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

// DeleteMessagesRequest represents TL type `deleteMessages#c3ed9de2`.
type DeleteMessagesRequest struct {
	// Chat identifier
	ChatID int64
	// Identifiers of the messages to be deleted. Use messageProperties
	// can_be_deleted_only_for_self and messageProperties.can_be_deleted_for_all_users to get
	// suitable messages
	MessageIDs []int64
	// Pass true to delete messages for all chat members. Always true for supergroups,
	// channels and secret chats
	Revoke bool
}

// DeleteMessagesRequestTypeID is TL type id of DeleteMessagesRequest.
const DeleteMessagesRequestTypeID = 0xc3ed9de2

// Ensuring interfaces in compile-time for DeleteMessagesRequest.
var (
	_ bin.Encoder     = &DeleteMessagesRequest{}
	_ bin.Decoder     = &DeleteMessagesRequest{}
	_ bin.BareEncoder = &DeleteMessagesRequest{}
	_ bin.BareDecoder = &DeleteMessagesRequest{}
)

func (d *DeleteMessagesRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ChatID == 0) {
		return false
	}
	if !(d.MessageIDs == nil) {
		return false
	}
	if !(d.Revoke == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteMessagesRequest) String() string {
	if d == nil {
		return "DeleteMessagesRequest(nil)"
	}
	type Alias DeleteMessagesRequest
	return fmt.Sprintf("DeleteMessagesRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteMessagesRequest) TypeID() uint32 {
	return DeleteMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteMessagesRequest) TypeName() string {
	return "deleteMessages"
}

// TypeInfo returns info about TL type.
func (d *DeleteMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteMessages",
		ID:   DeleteMessagesRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageIDs",
			SchemaName: "message_ids",
		},
		{
			Name:       "Revoke",
			SchemaName: "revoke",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteMessagesRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteMessages#c3ed9de2 as nil")
	}
	b.PutID(DeleteMessagesRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteMessages#c3ed9de2 as nil")
	}
	b.PutInt53(d.ChatID)
	b.PutInt(len(d.MessageIDs))
	for _, v := range d.MessageIDs {
		b.PutInt53(v)
	}
	b.PutBool(d.Revoke)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteMessagesRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteMessages#c3ed9de2 to nil")
	}
	if err := b.ConsumeID(DeleteMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteMessages#c3ed9de2: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteMessages#c3ed9de2 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode deleteMessages#c3ed9de2: field chat_id: %w", err)
		}
		d.ChatID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode deleteMessages#c3ed9de2: field message_ids: %w", err)
		}

		if headerLen > 0 {
			d.MessageIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode deleteMessages#c3ed9de2: field message_ids: %w", err)
			}
			d.MessageIDs = append(d.MessageIDs, value)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode deleteMessages#c3ed9de2: field revoke: %w", err)
		}
		d.Revoke = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteMessagesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteMessages#c3ed9de2 as nil")
	}
	b.ObjStart()
	b.PutID("deleteMessages")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(d.ChatID)
	b.Comma()
	b.FieldStart("message_ids")
	b.ArrStart()
	for _, v := range d.MessageIDs {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("revoke")
	b.PutBool(d.Revoke)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteMessagesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteMessages#c3ed9de2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteMessages"); err != nil {
				return fmt.Errorf("unable to decode deleteMessages#c3ed9de2: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode deleteMessages#c3ed9de2: field chat_id: %w", err)
			}
			d.ChatID = value
		case "message_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode deleteMessages#c3ed9de2: field message_ids: %w", err)
				}
				d.MessageIDs = append(d.MessageIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode deleteMessages#c3ed9de2: field message_ids: %w", err)
			}
		case "revoke":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode deleteMessages#c3ed9de2: field revoke: %w", err)
			}
			d.Revoke = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (d *DeleteMessagesRequest) GetChatID() (value int64) {
	if d == nil {
		return
	}
	return d.ChatID
}

// GetMessageIDs returns value of MessageIDs field.
func (d *DeleteMessagesRequest) GetMessageIDs() (value []int64) {
	if d == nil {
		return
	}
	return d.MessageIDs
}

// GetRevoke returns value of Revoke field.
func (d *DeleteMessagesRequest) GetRevoke() (value bool) {
	if d == nil {
		return
	}
	return d.Revoke
}

// DeleteMessages invokes method deleteMessages#c3ed9de2 returning error if any.
func (c *Client) DeleteMessages(ctx context.Context, request *DeleteMessagesRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
