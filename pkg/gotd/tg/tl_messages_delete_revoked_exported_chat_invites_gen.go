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

// MessagesDeleteRevokedExportedChatInvitesRequest represents TL type `messages.deleteRevokedExportedChatInvites#56987bd5`.
// Delete all revoked chat invites
//
// See https://core.telegram.org/method/messages.deleteRevokedExportedChatInvites for reference.
type MessagesDeleteRevokedExportedChatInvitesRequest struct {
	// Chat
	Peer InputPeerClass
	// ID of the admin that originally generated the revoked chat invites
	AdminID InputUserClass
}

// MessagesDeleteRevokedExportedChatInvitesRequestTypeID is TL type id of MessagesDeleteRevokedExportedChatInvitesRequest.
const MessagesDeleteRevokedExportedChatInvitesRequestTypeID = 0x56987bd5

// Ensuring interfaces in compile-time for MessagesDeleteRevokedExportedChatInvitesRequest.
var (
	_ bin.Encoder     = &MessagesDeleteRevokedExportedChatInvitesRequest{}
	_ bin.Decoder     = &MessagesDeleteRevokedExportedChatInvitesRequest{}
	_ bin.BareEncoder = &MessagesDeleteRevokedExportedChatInvitesRequest{}
	_ bin.BareDecoder = &MessagesDeleteRevokedExportedChatInvitesRequest{}
)

func (d *MessagesDeleteRevokedExportedChatInvitesRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Peer == nil) {
		return false
	}
	if !(d.AdminID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDeleteRevokedExportedChatInvitesRequest) String() string {
	if d == nil {
		return "MessagesDeleteRevokedExportedChatInvitesRequest(nil)"
	}
	type Alias MessagesDeleteRevokedExportedChatInvitesRequest
	return fmt.Sprintf("MessagesDeleteRevokedExportedChatInvitesRequest%+v", Alias(*d))
}

// FillFrom fills MessagesDeleteRevokedExportedChatInvitesRequest from given interface.
func (d *MessagesDeleteRevokedExportedChatInvitesRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetAdminID() (value InputUserClass)
}) {
	d.Peer = from.GetPeer()
	d.AdminID = from.GetAdminID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDeleteRevokedExportedChatInvitesRequest) TypeID() uint32 {
	return MessagesDeleteRevokedExportedChatInvitesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDeleteRevokedExportedChatInvitesRequest) TypeName() string {
	return "messages.deleteRevokedExportedChatInvites"
}

// TypeInfo returns info about TL type.
func (d *MessagesDeleteRevokedExportedChatInvitesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.deleteRevokedExportedChatInvites",
		ID:   MessagesDeleteRevokedExportedChatInvitesRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "AdminID",
			SchemaName: "admin_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *MessagesDeleteRevokedExportedChatInvitesRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteRevokedExportedChatInvites#56987bd5 as nil")
	}
	b.PutID(MessagesDeleteRevokedExportedChatInvitesRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDeleteRevokedExportedChatInvitesRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteRevokedExportedChatInvites#56987bd5 as nil")
	}
	if d.Peer == nil {
		return fmt.Errorf("unable to encode messages.deleteRevokedExportedChatInvites#56987bd5: field peer is nil")
	}
	if err := d.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteRevokedExportedChatInvites#56987bd5: field peer: %w", err)
	}
	if d.AdminID == nil {
		return fmt.Errorf("unable to encode messages.deleteRevokedExportedChatInvites#56987bd5: field admin_id is nil")
	}
	if err := d.AdminID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteRevokedExportedChatInvites#56987bd5: field admin_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDeleteRevokedExportedChatInvitesRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteRevokedExportedChatInvites#56987bd5 to nil")
	}
	if err := b.ConsumeID(MessagesDeleteRevokedExportedChatInvitesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.deleteRevokedExportedChatInvites#56987bd5: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDeleteRevokedExportedChatInvitesRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteRevokedExportedChatInvites#56987bd5 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteRevokedExportedChatInvites#56987bd5: field peer: %w", err)
		}
		d.Peer = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteRevokedExportedChatInvites#56987bd5: field admin_id: %w", err)
		}
		d.AdminID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (d *MessagesDeleteRevokedExportedChatInvitesRequest) GetPeer() (value InputPeerClass) {
	if d == nil {
		return
	}
	return d.Peer
}

// GetAdminID returns value of AdminID field.
func (d *MessagesDeleteRevokedExportedChatInvitesRequest) GetAdminID() (value InputUserClass) {
	if d == nil {
		return
	}
	return d.AdminID
}

// MessagesDeleteRevokedExportedChatInvites invokes method messages.deleteRevokedExportedChatInvites#56987bd5 returning error if any.
// Delete all revoked chat invites
//
// Possible errors:
//
//	400 ADMIN_ID_INVALID: The specified admin ID is invalid.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.deleteRevokedExportedChatInvites for reference.
func (c *Client) MessagesDeleteRevokedExportedChatInvites(ctx context.Context, request *MessagesDeleteRevokedExportedChatInvitesRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
