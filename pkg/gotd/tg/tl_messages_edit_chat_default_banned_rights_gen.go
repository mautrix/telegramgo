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

// MessagesEditChatDefaultBannedRightsRequest represents TL type `messages.editChatDefaultBannedRights#a5866b41`.
// Edit the default banned rights of a channel/supergroup/group¹.
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/messages.editChatDefaultBannedRights for reference.
type MessagesEditChatDefaultBannedRightsRequest struct {
	// The peer
	Peer InputPeerClass
	// The new global rights
	BannedRights ChatBannedRights
}

// MessagesEditChatDefaultBannedRightsRequestTypeID is TL type id of MessagesEditChatDefaultBannedRightsRequest.
const MessagesEditChatDefaultBannedRightsRequestTypeID = 0xa5866b41

// Ensuring interfaces in compile-time for MessagesEditChatDefaultBannedRightsRequest.
var (
	_ bin.Encoder     = &MessagesEditChatDefaultBannedRightsRequest{}
	_ bin.Decoder     = &MessagesEditChatDefaultBannedRightsRequest{}
	_ bin.BareEncoder = &MessagesEditChatDefaultBannedRightsRequest{}
	_ bin.BareDecoder = &MessagesEditChatDefaultBannedRightsRequest{}
)

func (e *MessagesEditChatDefaultBannedRightsRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Peer == nil) {
		return false
	}
	if !(e.BannedRights.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesEditChatDefaultBannedRightsRequest) String() string {
	if e == nil {
		return "MessagesEditChatDefaultBannedRightsRequest(nil)"
	}
	type Alias MessagesEditChatDefaultBannedRightsRequest
	return fmt.Sprintf("MessagesEditChatDefaultBannedRightsRequest%+v", Alias(*e))
}

// FillFrom fills MessagesEditChatDefaultBannedRightsRequest from given interface.
func (e *MessagesEditChatDefaultBannedRightsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetBannedRights() (value ChatBannedRights)
}) {
	e.Peer = from.GetPeer()
	e.BannedRights = from.GetBannedRights()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesEditChatDefaultBannedRightsRequest) TypeID() uint32 {
	return MessagesEditChatDefaultBannedRightsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesEditChatDefaultBannedRightsRequest) TypeName() string {
	return "messages.editChatDefaultBannedRights"
}

// TypeInfo returns info about TL type.
func (e *MessagesEditChatDefaultBannedRightsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.editChatDefaultBannedRights",
		ID:   MessagesEditChatDefaultBannedRightsRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "BannedRights",
			SchemaName: "banned_rights",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *MessagesEditChatDefaultBannedRightsRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editChatDefaultBannedRights#a5866b41 as nil")
	}
	b.PutID(MessagesEditChatDefaultBannedRightsRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *MessagesEditChatDefaultBannedRightsRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editChatDefaultBannedRights#a5866b41 as nil")
	}
	if e.Peer == nil {
		return fmt.Errorf("unable to encode messages.editChatDefaultBannedRights#a5866b41: field peer is nil")
	}
	if err := e.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editChatDefaultBannedRights#a5866b41: field peer: %w", err)
	}
	if err := e.BannedRights.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editChatDefaultBannedRights#a5866b41: field banned_rights: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *MessagesEditChatDefaultBannedRightsRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editChatDefaultBannedRights#a5866b41 to nil")
	}
	if err := b.ConsumeID(MessagesEditChatDefaultBannedRightsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.editChatDefaultBannedRights#a5866b41: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *MessagesEditChatDefaultBannedRightsRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editChatDefaultBannedRights#a5866b41 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatDefaultBannedRights#a5866b41: field peer: %w", err)
		}
		e.Peer = value
	}
	{
		if err := e.BannedRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.editChatDefaultBannedRights#a5866b41: field banned_rights: %w", err)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (e *MessagesEditChatDefaultBannedRightsRequest) GetPeer() (value InputPeerClass) {
	if e == nil {
		return
	}
	return e.Peer
}

// GetBannedRights returns value of BannedRights field.
func (e *MessagesEditChatDefaultBannedRightsRequest) GetBannedRights() (value ChatBannedRights) {
	if e == nil {
		return
	}
	return e.BannedRights
}

// MessagesEditChatDefaultBannedRights invokes method messages.editChatDefaultBannedRights#a5866b41 returning error if any.
// Edit the default banned rights of a channel/supergroup/group¹.
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// Possible errors:
//
//	400 BANNED_RIGHTS_INVALID: You provided some invalid flags in the banned rights.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 CHAT_ID_INVALID: The provided chat id is invalid.
//	400 CHAT_NOT_MODIFIED: No changes were made to chat information because the new information you passed is identical to the current information.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 UNTIL_DATE_INVALID: Invalid until date provided.
//
// See https://core.telegram.org/method/messages.editChatDefaultBannedRights for reference.
// Can be used by bots.
func (c *Client) MessagesEditChatDefaultBannedRights(ctx context.Context, request *MessagesEditChatDefaultBannedRightsRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
