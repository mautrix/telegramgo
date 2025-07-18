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

// GetChatPinnedMessageRequest represents TL type `getChatPinnedMessage#15731ab0`.
type GetChatPinnedMessageRequest struct {
	// Identifier of the chat the message belongs to
	ChatID int64
}

// GetChatPinnedMessageRequestTypeID is TL type id of GetChatPinnedMessageRequest.
const GetChatPinnedMessageRequestTypeID = 0x15731ab0

// Ensuring interfaces in compile-time for GetChatPinnedMessageRequest.
var (
	_ bin.Encoder     = &GetChatPinnedMessageRequest{}
	_ bin.Decoder     = &GetChatPinnedMessageRequest{}
	_ bin.BareEncoder = &GetChatPinnedMessageRequest{}
	_ bin.BareDecoder = &GetChatPinnedMessageRequest{}
)

func (g *GetChatPinnedMessageRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatPinnedMessageRequest) String() string {
	if g == nil {
		return "GetChatPinnedMessageRequest(nil)"
	}
	type Alias GetChatPinnedMessageRequest
	return fmt.Sprintf("GetChatPinnedMessageRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatPinnedMessageRequest) TypeID() uint32 {
	return GetChatPinnedMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatPinnedMessageRequest) TypeName() string {
	return "getChatPinnedMessage"
}

// TypeInfo returns info about TL type.
func (g *GetChatPinnedMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatPinnedMessage",
		ID:   GetChatPinnedMessageRequestTypeID,
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatPinnedMessageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatPinnedMessage#15731ab0 as nil")
	}
	b.PutID(GetChatPinnedMessageRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatPinnedMessageRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatPinnedMessage#15731ab0 as nil")
	}
	b.PutInt53(g.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatPinnedMessageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatPinnedMessage#15731ab0 to nil")
	}
	if err := b.ConsumeID(GetChatPinnedMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatPinnedMessage#15731ab0: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatPinnedMessageRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatPinnedMessage#15731ab0 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatPinnedMessage#15731ab0: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatPinnedMessageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatPinnedMessage#15731ab0 as nil")
	}
	b.ObjStart()
	b.PutID("getChatPinnedMessage")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatPinnedMessageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatPinnedMessage#15731ab0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatPinnedMessage"); err != nil {
				return fmt.Errorf("unable to decode getChatPinnedMessage#15731ab0: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatPinnedMessage#15731ab0: field chat_id: %w", err)
			}
			g.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatPinnedMessageRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetChatPinnedMessage invokes method getChatPinnedMessage#15731ab0 returning error if any.
func (c *Client) GetChatPinnedMessage(ctx context.Context, chatid int64) (*Message, error) {
	var result Message

	request := &GetChatPinnedMessageRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
