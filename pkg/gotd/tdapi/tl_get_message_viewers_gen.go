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

// GetMessageViewersRequest represents TL type `getMessageViewers#a18f1ace`.
type GetMessageViewersRequest struct {
	// Chat identifier
	ChatID int64
	// Identifier of the message
	MessageID int64
}

// GetMessageViewersRequestTypeID is TL type id of GetMessageViewersRequest.
const GetMessageViewersRequestTypeID = 0xa18f1ace

// Ensuring interfaces in compile-time for GetMessageViewersRequest.
var (
	_ bin.Encoder     = &GetMessageViewersRequest{}
	_ bin.Decoder     = &GetMessageViewersRequest{}
	_ bin.BareEncoder = &GetMessageViewersRequest{}
	_ bin.BareDecoder = &GetMessageViewersRequest{}
)

func (g *GetMessageViewersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetMessageViewersRequest) String() string {
	if g == nil {
		return "GetMessageViewersRequest(nil)"
	}
	type Alias GetMessageViewersRequest
	return fmt.Sprintf("GetMessageViewersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMessageViewersRequest) TypeID() uint32 {
	return GetMessageViewersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMessageViewersRequest) TypeName() string {
	return "getMessageViewers"
}

// TypeInfo returns info about TL type.
func (g *GetMessageViewersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMessageViewers",
		ID:   GetMessageViewersRequestTypeID,
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
			Name:       "MessageID",
			SchemaName: "message_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetMessageViewersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageViewers#a18f1ace as nil")
	}
	b.PutID(GetMessageViewersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMessageViewersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageViewers#a18f1ace as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt53(g.MessageID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMessageViewersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageViewers#a18f1ace to nil")
	}
	if err := b.ConsumeID(GetMessageViewersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMessageViewers#a18f1ace: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMessageViewersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageViewers#a18f1ace to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageViewers#a18f1ace: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageViewers#a18f1ace: field message_id: %w", err)
		}
		g.MessageID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetMessageViewersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageViewers#a18f1ace as nil")
	}
	b.ObjStart()
	b.PutID("getMessageViewers")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(g.MessageID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetMessageViewersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageViewers#a18f1ace to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getMessageViewers"); err != nil {
				return fmt.Errorf("unable to decode getMessageViewers#a18f1ace: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageViewers#a18f1ace: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageViewers#a18f1ace: field message_id: %w", err)
			}
			g.MessageID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetMessageViewersRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetMessageViewersRequest) GetMessageID() (value int64) {
	if g == nil {
		return
	}
	return g.MessageID
}

// GetMessageViewers invokes method getMessageViewers#a18f1ace returning error if any.
func (c *Client) GetMessageViewers(ctx context.Context, request *GetMessageViewersRequest) (*MessageViewers, error) {
	var result MessageViewers

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
