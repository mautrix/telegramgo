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

// GetChatInviteLinkCountsRequest represents TL type `getChatInviteLinkCounts#3510e291`.
type GetChatInviteLinkCountsRequest struct {
	// Chat identifier
	ChatID int64
}

// GetChatInviteLinkCountsRequestTypeID is TL type id of GetChatInviteLinkCountsRequest.
const GetChatInviteLinkCountsRequestTypeID = 0x3510e291

// Ensuring interfaces in compile-time for GetChatInviteLinkCountsRequest.
var (
	_ bin.Encoder     = &GetChatInviteLinkCountsRequest{}
	_ bin.Decoder     = &GetChatInviteLinkCountsRequest{}
	_ bin.BareEncoder = &GetChatInviteLinkCountsRequest{}
	_ bin.BareDecoder = &GetChatInviteLinkCountsRequest{}
)

func (g *GetChatInviteLinkCountsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatInviteLinkCountsRequest) String() string {
	if g == nil {
		return "GetChatInviteLinkCountsRequest(nil)"
	}
	type Alias GetChatInviteLinkCountsRequest
	return fmt.Sprintf("GetChatInviteLinkCountsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatInviteLinkCountsRequest) TypeID() uint32 {
	return GetChatInviteLinkCountsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatInviteLinkCountsRequest) TypeName() string {
	return "getChatInviteLinkCounts"
}

// TypeInfo returns info about TL type.
func (g *GetChatInviteLinkCountsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatInviteLinkCounts",
		ID:   GetChatInviteLinkCountsRequestTypeID,
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
func (g *GetChatInviteLinkCountsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatInviteLinkCounts#3510e291 as nil")
	}
	b.PutID(GetChatInviteLinkCountsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatInviteLinkCountsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatInviteLinkCounts#3510e291 as nil")
	}
	b.PutInt53(g.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatInviteLinkCountsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatInviteLinkCounts#3510e291 to nil")
	}
	if err := b.ConsumeID(GetChatInviteLinkCountsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatInviteLinkCounts#3510e291: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatInviteLinkCountsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatInviteLinkCounts#3510e291 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinkCounts#3510e291: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatInviteLinkCountsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatInviteLinkCounts#3510e291 as nil")
	}
	b.ObjStart()
	b.PutID("getChatInviteLinkCounts")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatInviteLinkCountsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatInviteLinkCounts#3510e291 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatInviteLinkCounts"); err != nil {
				return fmt.Errorf("unable to decode getChatInviteLinkCounts#3510e291: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatInviteLinkCounts#3510e291: field chat_id: %w", err)
			}
			g.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatInviteLinkCountsRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetChatInviteLinkCounts invokes method getChatInviteLinkCounts#3510e291 returning error if any.
func (c *Client) GetChatInviteLinkCounts(ctx context.Context, chatid int64) (*ChatInviteLinkCounts, error) {
	var result ChatInviteLinkCounts

	request := &GetChatInviteLinkCountsRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
