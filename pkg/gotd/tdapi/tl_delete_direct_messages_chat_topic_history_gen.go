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

// DeleteDirectMessagesChatTopicHistoryRequest represents TL type `deleteDirectMessagesChatTopicHistory#8e63863c`.
type DeleteDirectMessagesChatTopicHistoryRequest struct {
	// Chat identifier of the channel direct messages chat
	ChatID int64
	// Identifier of the topic which messages will be deleted
	TopicID int64
}

// DeleteDirectMessagesChatTopicHistoryRequestTypeID is TL type id of DeleteDirectMessagesChatTopicHistoryRequest.
const DeleteDirectMessagesChatTopicHistoryRequestTypeID = 0x8e63863c

// Ensuring interfaces in compile-time for DeleteDirectMessagesChatTopicHistoryRequest.
var (
	_ bin.Encoder     = &DeleteDirectMessagesChatTopicHistoryRequest{}
	_ bin.Decoder     = &DeleteDirectMessagesChatTopicHistoryRequest{}
	_ bin.BareEncoder = &DeleteDirectMessagesChatTopicHistoryRequest{}
	_ bin.BareDecoder = &DeleteDirectMessagesChatTopicHistoryRequest{}
)

func (d *DeleteDirectMessagesChatTopicHistoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ChatID == 0) {
		return false
	}
	if !(d.TopicID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteDirectMessagesChatTopicHistoryRequest) String() string {
	if d == nil {
		return "DeleteDirectMessagesChatTopicHistoryRequest(nil)"
	}
	type Alias DeleteDirectMessagesChatTopicHistoryRequest
	return fmt.Sprintf("DeleteDirectMessagesChatTopicHistoryRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteDirectMessagesChatTopicHistoryRequest) TypeID() uint32 {
	return DeleteDirectMessagesChatTopicHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteDirectMessagesChatTopicHistoryRequest) TypeName() string {
	return "deleteDirectMessagesChatTopicHistory"
}

// TypeInfo returns info about TL type.
func (d *DeleteDirectMessagesChatTopicHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteDirectMessagesChatTopicHistory",
		ID:   DeleteDirectMessagesChatTopicHistoryRequestTypeID,
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
			Name:       "TopicID",
			SchemaName: "topic_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteDirectMessagesChatTopicHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteDirectMessagesChatTopicHistory#8e63863c as nil")
	}
	b.PutID(DeleteDirectMessagesChatTopicHistoryRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteDirectMessagesChatTopicHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteDirectMessagesChatTopicHistory#8e63863c as nil")
	}
	b.PutInt53(d.ChatID)
	b.PutInt53(d.TopicID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteDirectMessagesChatTopicHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteDirectMessagesChatTopicHistory#8e63863c to nil")
	}
	if err := b.ConsumeID(DeleteDirectMessagesChatTopicHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteDirectMessagesChatTopicHistory#8e63863c: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteDirectMessagesChatTopicHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteDirectMessagesChatTopicHistory#8e63863c to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode deleteDirectMessagesChatTopicHistory#8e63863c: field chat_id: %w", err)
		}
		d.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode deleteDirectMessagesChatTopicHistory#8e63863c: field topic_id: %w", err)
		}
		d.TopicID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteDirectMessagesChatTopicHistoryRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteDirectMessagesChatTopicHistory#8e63863c as nil")
	}
	b.ObjStart()
	b.PutID("deleteDirectMessagesChatTopicHistory")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(d.ChatID)
	b.Comma()
	b.FieldStart("topic_id")
	b.PutInt53(d.TopicID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteDirectMessagesChatTopicHistoryRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteDirectMessagesChatTopicHistory#8e63863c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteDirectMessagesChatTopicHistory"); err != nil {
				return fmt.Errorf("unable to decode deleteDirectMessagesChatTopicHistory#8e63863c: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode deleteDirectMessagesChatTopicHistory#8e63863c: field chat_id: %w", err)
			}
			d.ChatID = value
		case "topic_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode deleteDirectMessagesChatTopicHistory#8e63863c: field topic_id: %w", err)
			}
			d.TopicID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (d *DeleteDirectMessagesChatTopicHistoryRequest) GetChatID() (value int64) {
	if d == nil {
		return
	}
	return d.ChatID
}

// GetTopicID returns value of TopicID field.
func (d *DeleteDirectMessagesChatTopicHistoryRequest) GetTopicID() (value int64) {
	if d == nil {
		return
	}
	return d.TopicID
}

// DeleteDirectMessagesChatTopicHistory invokes method deleteDirectMessagesChatTopicHistory#8e63863c returning error if any.
func (c *Client) DeleteDirectMessagesChatTopicHistory(ctx context.Context, request *DeleteDirectMessagesChatTopicHistoryRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
