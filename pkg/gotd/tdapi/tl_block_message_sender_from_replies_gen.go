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

// BlockMessageSenderFromRepliesRequest represents TL type `blockMessageSenderFromReplies#b79df58b`.
type BlockMessageSenderFromRepliesRequest struct {
	// The identifier of an incoming message in the Replies chat
	MessageID int64
	// Pass true to delete the message
	DeleteMessage bool
	// Pass true to delete all messages from the same sender
	DeleteAllMessages bool
	// Pass true to report the sender to the Telegram moderators
	ReportSpam bool
}

// BlockMessageSenderFromRepliesRequestTypeID is TL type id of BlockMessageSenderFromRepliesRequest.
const BlockMessageSenderFromRepliesRequestTypeID = 0xb79df58b

// Ensuring interfaces in compile-time for BlockMessageSenderFromRepliesRequest.
var (
	_ bin.Encoder     = &BlockMessageSenderFromRepliesRequest{}
	_ bin.Decoder     = &BlockMessageSenderFromRepliesRequest{}
	_ bin.BareEncoder = &BlockMessageSenderFromRepliesRequest{}
	_ bin.BareDecoder = &BlockMessageSenderFromRepliesRequest{}
)

func (b *BlockMessageSenderFromRepliesRequest) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.MessageID == 0) {
		return false
	}
	if !(b.DeleteMessage == false) {
		return false
	}
	if !(b.DeleteAllMessages == false) {
		return false
	}
	if !(b.ReportSpam == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BlockMessageSenderFromRepliesRequest) String() string {
	if b == nil {
		return "BlockMessageSenderFromRepliesRequest(nil)"
	}
	type Alias BlockMessageSenderFromRepliesRequest
	return fmt.Sprintf("BlockMessageSenderFromRepliesRequest%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BlockMessageSenderFromRepliesRequest) TypeID() uint32 {
	return BlockMessageSenderFromRepliesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BlockMessageSenderFromRepliesRequest) TypeName() string {
	return "blockMessageSenderFromReplies"
}

// TypeInfo returns info about TL type.
func (b *BlockMessageSenderFromRepliesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "blockMessageSenderFromReplies",
		ID:   BlockMessageSenderFromRepliesRequestTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "DeleteMessage",
			SchemaName: "delete_message",
		},
		{
			Name:       "DeleteAllMessages",
			SchemaName: "delete_all_messages",
		},
		{
			Name:       "ReportSpam",
			SchemaName: "report_spam",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BlockMessageSenderFromRepliesRequest) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode blockMessageSenderFromReplies#b79df58b as nil")
	}
	buf.PutID(BlockMessageSenderFromRepliesRequestTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BlockMessageSenderFromRepliesRequest) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode blockMessageSenderFromReplies#b79df58b as nil")
	}
	buf.PutInt53(b.MessageID)
	buf.PutBool(b.DeleteMessage)
	buf.PutBool(b.DeleteAllMessages)
	buf.PutBool(b.ReportSpam)
	return nil
}

// Decode implements bin.Decoder.
func (b *BlockMessageSenderFromRepliesRequest) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode blockMessageSenderFromReplies#b79df58b to nil")
	}
	if err := buf.ConsumeID(BlockMessageSenderFromRepliesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode blockMessageSenderFromReplies#b79df58b: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BlockMessageSenderFromRepliesRequest) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode blockMessageSenderFromReplies#b79df58b to nil")
	}
	{
		value, err := buf.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode blockMessageSenderFromReplies#b79df58b: field message_id: %w", err)
		}
		b.MessageID = value
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode blockMessageSenderFromReplies#b79df58b: field delete_message: %w", err)
		}
		b.DeleteMessage = value
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode blockMessageSenderFromReplies#b79df58b: field delete_all_messages: %w", err)
		}
		b.DeleteAllMessages = value
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode blockMessageSenderFromReplies#b79df58b: field report_spam: %w", err)
		}
		b.ReportSpam = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BlockMessageSenderFromRepliesRequest) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode blockMessageSenderFromReplies#b79df58b as nil")
	}
	buf.ObjStart()
	buf.PutID("blockMessageSenderFromReplies")
	buf.Comma()
	buf.FieldStart("message_id")
	buf.PutInt53(b.MessageID)
	buf.Comma()
	buf.FieldStart("delete_message")
	buf.PutBool(b.DeleteMessage)
	buf.Comma()
	buf.FieldStart("delete_all_messages")
	buf.PutBool(b.DeleteAllMessages)
	buf.Comma()
	buf.FieldStart("report_spam")
	buf.PutBool(b.ReportSpam)
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BlockMessageSenderFromRepliesRequest) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode blockMessageSenderFromReplies#b79df58b to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("blockMessageSenderFromReplies"); err != nil {
				return fmt.Errorf("unable to decode blockMessageSenderFromReplies#b79df58b: %w", err)
			}
		case "message_id":
			value, err := buf.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode blockMessageSenderFromReplies#b79df58b: field message_id: %w", err)
			}
			b.MessageID = value
		case "delete_message":
			value, err := buf.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode blockMessageSenderFromReplies#b79df58b: field delete_message: %w", err)
			}
			b.DeleteMessage = value
		case "delete_all_messages":
			value, err := buf.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode blockMessageSenderFromReplies#b79df58b: field delete_all_messages: %w", err)
			}
			b.DeleteAllMessages = value
		case "report_spam":
			value, err := buf.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode blockMessageSenderFromReplies#b79df58b: field report_spam: %w", err)
			}
			b.ReportSpam = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetMessageID returns value of MessageID field.
func (b *BlockMessageSenderFromRepliesRequest) GetMessageID() (value int64) {
	if b == nil {
		return
	}
	return b.MessageID
}

// GetDeleteMessage returns value of DeleteMessage field.
func (b *BlockMessageSenderFromRepliesRequest) GetDeleteMessage() (value bool) {
	if b == nil {
		return
	}
	return b.DeleteMessage
}

// GetDeleteAllMessages returns value of DeleteAllMessages field.
func (b *BlockMessageSenderFromRepliesRequest) GetDeleteAllMessages() (value bool) {
	if b == nil {
		return
	}
	return b.DeleteAllMessages
}

// GetReportSpam returns value of ReportSpam field.
func (b *BlockMessageSenderFromRepliesRequest) GetReportSpam() (value bool) {
	if b == nil {
		return
	}
	return b.ReportSpam
}

// BlockMessageSenderFromReplies invokes method blockMessageSenderFromReplies#b79df58b returning error if any.
func (c *Client) BlockMessageSenderFromReplies(ctx context.Context, request *BlockMessageSenderFromRepliesRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
