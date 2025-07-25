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

// ReadChatListRequest represents TL type `readChatList#bd6498aa`.
type ReadChatListRequest struct {
	// Chat list in which to mark all chats as read
	ChatList ChatListClass
}

// ReadChatListRequestTypeID is TL type id of ReadChatListRequest.
const ReadChatListRequestTypeID = 0xbd6498aa

// Ensuring interfaces in compile-time for ReadChatListRequest.
var (
	_ bin.Encoder     = &ReadChatListRequest{}
	_ bin.Decoder     = &ReadChatListRequest{}
	_ bin.BareEncoder = &ReadChatListRequest{}
	_ bin.BareDecoder = &ReadChatListRequest{}
)

func (r *ReadChatListRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatList == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReadChatListRequest) String() string {
	if r == nil {
		return "ReadChatListRequest(nil)"
	}
	type Alias ReadChatListRequest
	return fmt.Sprintf("ReadChatListRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReadChatListRequest) TypeID() uint32 {
	return ReadChatListRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReadChatListRequest) TypeName() string {
	return "readChatList"
}

// TypeInfo returns info about TL type.
func (r *ReadChatListRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "readChatList",
		ID:   ReadChatListRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatList",
			SchemaName: "chat_list",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReadChatListRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode readChatList#bd6498aa as nil")
	}
	b.PutID(ReadChatListRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReadChatListRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode readChatList#bd6498aa as nil")
	}
	if r.ChatList == nil {
		return fmt.Errorf("unable to encode readChatList#bd6498aa: field chat_list is nil")
	}
	if err := r.ChatList.Encode(b); err != nil {
		return fmt.Errorf("unable to encode readChatList#bd6498aa: field chat_list: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReadChatListRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode readChatList#bd6498aa to nil")
	}
	if err := b.ConsumeID(ReadChatListRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode readChatList#bd6498aa: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReadChatListRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode readChatList#bd6498aa to nil")
	}
	{
		value, err := DecodeChatList(b)
		if err != nil {
			return fmt.Errorf("unable to decode readChatList#bd6498aa: field chat_list: %w", err)
		}
		r.ChatList = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReadChatListRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode readChatList#bd6498aa as nil")
	}
	b.ObjStart()
	b.PutID("readChatList")
	b.Comma()
	b.FieldStart("chat_list")
	if r.ChatList == nil {
		return fmt.Errorf("unable to encode readChatList#bd6498aa: field chat_list is nil")
	}
	if err := r.ChatList.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode readChatList#bd6498aa: field chat_list: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReadChatListRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode readChatList#bd6498aa to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("readChatList"); err != nil {
				return fmt.Errorf("unable to decode readChatList#bd6498aa: %w", err)
			}
		case "chat_list":
			value, err := DecodeTDLibJSONChatList(b)
			if err != nil {
				return fmt.Errorf("unable to decode readChatList#bd6498aa: field chat_list: %w", err)
			}
			r.ChatList = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatList returns value of ChatList field.
func (r *ReadChatListRequest) GetChatList() (value ChatListClass) {
	if r == nil {
		return
	}
	return r.ChatList
}

// ReadChatList invokes method readChatList#bd6498aa returning error if any.
func (c *Client) ReadChatList(ctx context.Context, chatlist ChatListClass) error {
	var ok Ok

	request := &ReadChatListRequest{
		ChatList: chatlist,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
