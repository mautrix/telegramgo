// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// ChatlistsLeaveChatlistRequest represents TL type `chatlists.leaveChatlist#74fae13a`.
// Delete a folder imported using a chat folder deep link »¹
//
// Links:
//  1. https://core.telegram.org/api/links#chat-folder-links
//
// See https://core.telegram.org/method/chatlists.leaveChatlist for reference.
type ChatlistsLeaveChatlistRequest struct {
	// Folder ID
	Chatlist InputChatlistDialogFilter
	// Also leave the specified channels and groups
	Peers []InputPeerClass
}

// ChatlistsLeaveChatlistRequestTypeID is TL type id of ChatlistsLeaveChatlistRequest.
const ChatlistsLeaveChatlistRequestTypeID = 0x74fae13a

// Ensuring interfaces in compile-time for ChatlistsLeaveChatlistRequest.
var (
	_ bin.Encoder     = &ChatlistsLeaveChatlistRequest{}
	_ bin.Decoder     = &ChatlistsLeaveChatlistRequest{}
	_ bin.BareEncoder = &ChatlistsLeaveChatlistRequest{}
	_ bin.BareDecoder = &ChatlistsLeaveChatlistRequest{}
)

func (l *ChatlistsLeaveChatlistRequest) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Chatlist.Zero()) {
		return false
	}
	if !(l.Peers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *ChatlistsLeaveChatlistRequest) String() string {
	if l == nil {
		return "ChatlistsLeaveChatlistRequest(nil)"
	}
	type Alias ChatlistsLeaveChatlistRequest
	return fmt.Sprintf("ChatlistsLeaveChatlistRequest%+v", Alias(*l))
}

// FillFrom fills ChatlistsLeaveChatlistRequest from given interface.
func (l *ChatlistsLeaveChatlistRequest) FillFrom(from interface {
	GetChatlist() (value InputChatlistDialogFilter)
	GetPeers() (value []InputPeerClass)
}) {
	l.Chatlist = from.GetChatlist()
	l.Peers = from.GetPeers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatlistsLeaveChatlistRequest) TypeID() uint32 {
	return ChatlistsLeaveChatlistRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatlistsLeaveChatlistRequest) TypeName() string {
	return "chatlists.leaveChatlist"
}

// TypeInfo returns info about TL type.
func (l *ChatlistsLeaveChatlistRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatlists.leaveChatlist",
		ID:   ChatlistsLeaveChatlistRequestTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Chatlist",
			SchemaName: "chatlist",
		},
		{
			Name:       "Peers",
			SchemaName: "peers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *ChatlistsLeaveChatlistRequest) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode chatlists.leaveChatlist#74fae13a as nil")
	}
	b.PutID(ChatlistsLeaveChatlistRequestTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *ChatlistsLeaveChatlistRequest) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode chatlists.leaveChatlist#74fae13a as nil")
	}
	if err := l.Chatlist.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatlists.leaveChatlist#74fae13a: field chatlist: %w", err)
	}
	b.PutVectorHeader(len(l.Peers))
	for idx, v := range l.Peers {
		if v == nil {
			return fmt.Errorf("unable to encode chatlists.leaveChatlist#74fae13a: field peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatlists.leaveChatlist#74fae13a: field peers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *ChatlistsLeaveChatlistRequest) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode chatlists.leaveChatlist#74fae13a to nil")
	}
	if err := b.ConsumeID(ChatlistsLeaveChatlistRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode chatlists.leaveChatlist#74fae13a: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *ChatlistsLeaveChatlistRequest) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode chatlists.leaveChatlist#74fae13a to nil")
	}
	{
		if err := l.Chatlist.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatlists.leaveChatlist#74fae13a: field chatlist: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.leaveChatlist#74fae13a: field peers: %w", err)
		}

		if headerLen > 0 {
			l.Peers = make([]InputPeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatlists.leaveChatlist#74fae13a: field peers: %w", err)
			}
			l.Peers = append(l.Peers, value)
		}
	}
	return nil
}

// GetChatlist returns value of Chatlist field.
func (l *ChatlistsLeaveChatlistRequest) GetChatlist() (value InputChatlistDialogFilter) {
	if l == nil {
		return
	}
	return l.Chatlist
}

// GetPeers returns value of Peers field.
func (l *ChatlistsLeaveChatlistRequest) GetPeers() (value []InputPeerClass) {
	if l == nil {
		return
	}
	return l.Peers
}

// MapPeers returns field Peers wrapped in InputPeerClassArray helper.
func (l *ChatlistsLeaveChatlistRequest) MapPeers() (value InputPeerClassArray) {
	return InputPeerClassArray(l.Peers)
}

// ChatlistsLeaveChatlist invokes method chatlists.leaveChatlist#74fae13a returning error if any.
// Delete a folder imported using a chat folder deep link »¹
//
// Links:
//  1. https://core.telegram.org/api/links#chat-folder-links
//
// Possible errors:
//
//	400 FILTER_ID_INVALID: The specified filter ID is invalid.
//
// See https://core.telegram.org/method/chatlists.leaveChatlist for reference.
func (c *Client) ChatlistsLeaveChatlist(ctx context.Context, request *ChatlistsLeaveChatlistRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
