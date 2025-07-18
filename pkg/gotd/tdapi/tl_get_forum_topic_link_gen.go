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

// GetForumTopicLinkRequest represents TL type `getForumTopicLink#c97b88cb`.
type GetForumTopicLinkRequest struct {
	// Identifier of the chat
	ChatID int64
	// Message thread identifier of the forum topic
	MessageThreadID int64
}

// GetForumTopicLinkRequestTypeID is TL type id of GetForumTopicLinkRequest.
const GetForumTopicLinkRequestTypeID = 0xc97b88cb

// Ensuring interfaces in compile-time for GetForumTopicLinkRequest.
var (
	_ bin.Encoder     = &GetForumTopicLinkRequest{}
	_ bin.Decoder     = &GetForumTopicLinkRequest{}
	_ bin.BareEncoder = &GetForumTopicLinkRequest{}
	_ bin.BareDecoder = &GetForumTopicLinkRequest{}
)

func (g *GetForumTopicLinkRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageThreadID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetForumTopicLinkRequest) String() string {
	if g == nil {
		return "GetForumTopicLinkRequest(nil)"
	}
	type Alias GetForumTopicLinkRequest
	return fmt.Sprintf("GetForumTopicLinkRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetForumTopicLinkRequest) TypeID() uint32 {
	return GetForumTopicLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetForumTopicLinkRequest) TypeName() string {
	return "getForumTopicLink"
}

// TypeInfo returns info about TL type.
func (g *GetForumTopicLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getForumTopicLink",
		ID:   GetForumTopicLinkRequestTypeID,
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
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetForumTopicLinkRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getForumTopicLink#c97b88cb as nil")
	}
	b.PutID(GetForumTopicLinkRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetForumTopicLinkRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getForumTopicLink#c97b88cb as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt53(g.MessageThreadID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetForumTopicLinkRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getForumTopicLink#c97b88cb to nil")
	}
	if err := b.ConsumeID(GetForumTopicLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getForumTopicLink#c97b88cb: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetForumTopicLinkRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getForumTopicLink#c97b88cb to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getForumTopicLink#c97b88cb: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getForumTopicLink#c97b88cb: field message_thread_id: %w", err)
		}
		g.MessageThreadID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetForumTopicLinkRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getForumTopicLink#c97b88cb as nil")
	}
	b.ObjStart()
	b.PutID("getForumTopicLink")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("message_thread_id")
	b.PutInt53(g.MessageThreadID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetForumTopicLinkRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getForumTopicLink#c97b88cb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getForumTopicLink"); err != nil {
				return fmt.Errorf("unable to decode getForumTopicLink#c97b88cb: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getForumTopicLink#c97b88cb: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_thread_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getForumTopicLink#c97b88cb: field message_thread_id: %w", err)
			}
			g.MessageThreadID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetForumTopicLinkRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetMessageThreadID returns value of MessageThreadID field.
func (g *GetForumTopicLinkRequest) GetMessageThreadID() (value int64) {
	if g == nil {
		return
	}
	return g.MessageThreadID
}

// GetForumTopicLink invokes method getForumTopicLink#c97b88cb returning error if any.
func (c *Client) GetForumTopicLink(ctx context.Context, request *GetForumTopicLinkRequest) (*MessageLink, error) {
	var result MessageLink

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
