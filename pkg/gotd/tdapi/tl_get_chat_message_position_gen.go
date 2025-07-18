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

// GetChatMessagePositionRequest represents TL type `getChatMessagePosition#a87d6f0f`.
type GetChatMessagePositionRequest struct {
	// Identifier of the chat in which to find message position
	ChatID int64
	// Pass topic identifier to get position among messages only in specific topic; pass null
	// to get position among all chat messages
	TopicID MessageTopicClass
	// Filter for message content; searchMessagesFilterEmpty,
	// searchMessagesFilterUnreadMention, searchMessagesFilterUnreadReaction, and
	// searchMessagesFilterFailedToSend are unsupported in this function
	Filter SearchMessagesFilterClass
	// Message identifier
	MessageID int64
}

// GetChatMessagePositionRequestTypeID is TL type id of GetChatMessagePositionRequest.
const GetChatMessagePositionRequestTypeID = 0xa87d6f0f

// Ensuring interfaces in compile-time for GetChatMessagePositionRequest.
var (
	_ bin.Encoder     = &GetChatMessagePositionRequest{}
	_ bin.Decoder     = &GetChatMessagePositionRequest{}
	_ bin.BareEncoder = &GetChatMessagePositionRequest{}
	_ bin.BareDecoder = &GetChatMessagePositionRequest{}
)

func (g *GetChatMessagePositionRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.TopicID == nil) {
		return false
	}
	if !(g.Filter == nil) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatMessagePositionRequest) String() string {
	if g == nil {
		return "GetChatMessagePositionRequest(nil)"
	}
	type Alias GetChatMessagePositionRequest
	return fmt.Sprintf("GetChatMessagePositionRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatMessagePositionRequest) TypeID() uint32 {
	return GetChatMessagePositionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatMessagePositionRequest) TypeName() string {
	return "getChatMessagePosition"
}

// TypeInfo returns info about TL type.
func (g *GetChatMessagePositionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatMessagePosition",
		ID:   GetChatMessagePositionRequestTypeID,
	}
	if g == nil {
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
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatMessagePositionRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessagePosition#a87d6f0f as nil")
	}
	b.PutID(GetChatMessagePositionRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatMessagePositionRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessagePosition#a87d6f0f as nil")
	}
	b.PutInt53(g.ChatID)
	if g.TopicID == nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#a87d6f0f: field topic_id is nil")
	}
	if err := g.TopicID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#a87d6f0f: field topic_id: %w", err)
	}
	if g.Filter == nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#a87d6f0f: field filter is nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#a87d6f0f: field filter: %w", err)
	}
	b.PutInt53(g.MessageID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatMessagePositionRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessagePosition#a87d6f0f to nil")
	}
	if err := b.ConsumeID(GetChatMessagePositionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatMessagePosition#a87d6f0f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatMessagePositionRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessagePosition#a87d6f0f to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessagePosition#a87d6f0f: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := DecodeMessageTopic(b)
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessagePosition#a87d6f0f: field topic_id: %w", err)
		}
		g.TopicID = value
	}
	{
		value, err := DecodeSearchMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessagePosition#a87d6f0f: field filter: %w", err)
		}
		g.Filter = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessagePosition#a87d6f0f: field message_id: %w", err)
		}
		g.MessageID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatMessagePositionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessagePosition#a87d6f0f as nil")
	}
	b.ObjStart()
	b.PutID("getChatMessagePosition")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("topic_id")
	if g.TopicID == nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#a87d6f0f: field topic_id is nil")
	}
	if err := g.TopicID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#a87d6f0f: field topic_id: %w", err)
	}
	b.Comma()
	b.FieldStart("filter")
	if g.Filter == nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#a87d6f0f: field filter is nil")
	}
	if err := g.Filter.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#a87d6f0f: field filter: %w", err)
	}
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(g.MessageID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatMessagePositionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessagePosition#a87d6f0f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatMessagePosition"); err != nil {
				return fmt.Errorf("unable to decode getChatMessagePosition#a87d6f0f: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessagePosition#a87d6f0f: field chat_id: %w", err)
			}
			g.ChatID = value
		case "topic_id":
			value, err := DecodeTDLibJSONMessageTopic(b)
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessagePosition#a87d6f0f: field topic_id: %w", err)
			}
			g.TopicID = value
		case "filter":
			value, err := DecodeTDLibJSONSearchMessagesFilter(b)
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessagePosition#a87d6f0f: field filter: %w", err)
			}
			g.Filter = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessagePosition#a87d6f0f: field message_id: %w", err)
			}
			g.MessageID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatMessagePositionRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetTopicID returns value of TopicID field.
func (g *GetChatMessagePositionRequest) GetTopicID() (value MessageTopicClass) {
	if g == nil {
		return
	}
	return g.TopicID
}

// GetFilter returns value of Filter field.
func (g *GetChatMessagePositionRequest) GetFilter() (value SearchMessagesFilterClass) {
	if g == nil {
		return
	}
	return g.Filter
}

// GetMessageID returns value of MessageID field.
func (g *GetChatMessagePositionRequest) GetMessageID() (value int64) {
	if g == nil {
		return
	}
	return g.MessageID
}

// GetChatMessagePosition invokes method getChatMessagePosition#a87d6f0f returning error if any.
func (c *Client) GetChatMessagePosition(ctx context.Context, request *GetChatMessagePositionRequest) (*Count, error) {
	var result Count

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
