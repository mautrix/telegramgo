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

// ToggleChatIsPinnedRequest represents TL type `toggleChatIsPinned#a776263e`.
type ToggleChatIsPinnedRequest struct {
	// Chat list in which to change the pinned state of the chat
	ChatList ChatListClass
	// Chat identifier
	ChatID int64
	// Pass true to pin the chat; pass false to unpin it
	IsPinned bool
}

// ToggleChatIsPinnedRequestTypeID is TL type id of ToggleChatIsPinnedRequest.
const ToggleChatIsPinnedRequestTypeID = 0xa776263e

// Ensuring interfaces in compile-time for ToggleChatIsPinnedRequest.
var (
	_ bin.Encoder     = &ToggleChatIsPinnedRequest{}
	_ bin.Decoder     = &ToggleChatIsPinnedRequest{}
	_ bin.BareEncoder = &ToggleChatIsPinnedRequest{}
	_ bin.BareDecoder = &ToggleChatIsPinnedRequest{}
)

func (t *ToggleChatIsPinnedRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.ChatList == nil) {
		return false
	}
	if !(t.ChatID == 0) {
		return false
	}
	if !(t.IsPinned == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleChatIsPinnedRequest) String() string {
	if t == nil {
		return "ToggleChatIsPinnedRequest(nil)"
	}
	type Alias ToggleChatIsPinnedRequest
	return fmt.Sprintf("ToggleChatIsPinnedRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleChatIsPinnedRequest) TypeID() uint32 {
	return ToggleChatIsPinnedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleChatIsPinnedRequest) TypeName() string {
	return "toggleChatIsPinned"
}

// TypeInfo returns info about TL type.
func (t *ToggleChatIsPinnedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleChatIsPinned",
		ID:   ToggleChatIsPinnedRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatList",
			SchemaName: "chat_list",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "IsPinned",
			SchemaName: "is_pinned",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleChatIsPinnedRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatIsPinned#a776263e as nil")
	}
	b.PutID(ToggleChatIsPinnedRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleChatIsPinnedRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatIsPinned#a776263e as nil")
	}
	if t.ChatList == nil {
		return fmt.Errorf("unable to encode toggleChatIsPinned#a776263e: field chat_list is nil")
	}
	if err := t.ChatList.Encode(b); err != nil {
		return fmt.Errorf("unable to encode toggleChatIsPinned#a776263e: field chat_list: %w", err)
	}
	b.PutInt53(t.ChatID)
	b.PutBool(t.IsPinned)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleChatIsPinnedRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatIsPinned#a776263e to nil")
	}
	if err := b.ConsumeID(ToggleChatIsPinnedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleChatIsPinned#a776263e: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleChatIsPinnedRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatIsPinned#a776263e to nil")
	}
	{
		value, err := DecodeChatList(b)
		if err != nil {
			return fmt.Errorf("unable to decode toggleChatIsPinned#a776263e: field chat_list: %w", err)
		}
		t.ChatList = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode toggleChatIsPinned#a776263e: field chat_id: %w", err)
		}
		t.ChatID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleChatIsPinned#a776263e: field is_pinned: %w", err)
		}
		t.IsPinned = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleChatIsPinnedRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatIsPinned#a776263e as nil")
	}
	b.ObjStart()
	b.PutID("toggleChatIsPinned")
	b.Comma()
	b.FieldStart("chat_list")
	if t.ChatList == nil {
		return fmt.Errorf("unable to encode toggleChatIsPinned#a776263e: field chat_list is nil")
	}
	if err := t.ChatList.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode toggleChatIsPinned#a776263e: field chat_list: %w", err)
	}
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(t.ChatID)
	b.Comma()
	b.FieldStart("is_pinned")
	b.PutBool(t.IsPinned)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleChatIsPinnedRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatIsPinned#a776263e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleChatIsPinned"); err != nil {
				return fmt.Errorf("unable to decode toggleChatIsPinned#a776263e: %w", err)
			}
		case "chat_list":
			value, err := DecodeTDLibJSONChatList(b)
			if err != nil {
				return fmt.Errorf("unable to decode toggleChatIsPinned#a776263e: field chat_list: %w", err)
			}
			t.ChatList = value
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode toggleChatIsPinned#a776263e: field chat_id: %w", err)
			}
			t.ChatID = value
		case "is_pinned":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleChatIsPinned#a776263e: field is_pinned: %w", err)
			}
			t.IsPinned = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatList returns value of ChatList field.
func (t *ToggleChatIsPinnedRequest) GetChatList() (value ChatListClass) {
	if t == nil {
		return
	}
	return t.ChatList
}

// GetChatID returns value of ChatID field.
func (t *ToggleChatIsPinnedRequest) GetChatID() (value int64) {
	if t == nil {
		return
	}
	return t.ChatID
}

// GetIsPinned returns value of IsPinned field.
func (t *ToggleChatIsPinnedRequest) GetIsPinned() (value bool) {
	if t == nil {
		return
	}
	return t.IsPinned
}

// ToggleChatIsPinned invokes method toggleChatIsPinned#a776263e returning error if any.
func (c *Client) ToggleChatIsPinned(ctx context.Context, request *ToggleChatIsPinnedRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
