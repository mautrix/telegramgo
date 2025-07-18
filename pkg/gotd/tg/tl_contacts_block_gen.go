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

// ContactsBlockRequest represents TL type `contacts.block#2e2e8734`.
// Adds a peer to a blocklist, see here »¹ for more info.
//
// Links:
//  1. https://core.telegram.org/api/block
//
// See https://core.telegram.org/method/contacts.block for reference.
type ContactsBlockRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the peer should be added to the story blocklist; if not set, the peer will be
	// added to the main blocklist, see here »¹ for more info.
	//
	// Links:
	//  1) https://core.telegram.org/api/block
	MyStoriesFrom bool
	// Peer
	ID InputPeerClass
}

// ContactsBlockRequestTypeID is TL type id of ContactsBlockRequest.
const ContactsBlockRequestTypeID = 0x2e2e8734

// Ensuring interfaces in compile-time for ContactsBlockRequest.
var (
	_ bin.Encoder     = &ContactsBlockRequest{}
	_ bin.Decoder     = &ContactsBlockRequest{}
	_ bin.BareEncoder = &ContactsBlockRequest{}
	_ bin.BareDecoder = &ContactsBlockRequest{}
)

func (b *ContactsBlockRequest) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Flags.Zero()) {
		return false
	}
	if !(b.MyStoriesFrom == false) {
		return false
	}
	if !(b.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *ContactsBlockRequest) String() string {
	if b == nil {
		return "ContactsBlockRequest(nil)"
	}
	type Alias ContactsBlockRequest
	return fmt.Sprintf("ContactsBlockRequest%+v", Alias(*b))
}

// FillFrom fills ContactsBlockRequest from given interface.
func (b *ContactsBlockRequest) FillFrom(from interface {
	GetMyStoriesFrom() (value bool)
	GetID() (value InputPeerClass)
}) {
	b.MyStoriesFrom = from.GetMyStoriesFrom()
	b.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsBlockRequest) TypeID() uint32 {
	return ContactsBlockRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsBlockRequest) TypeName() string {
	return "contacts.block"
}

// TypeInfo returns info about TL type.
func (b *ContactsBlockRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.block",
		ID:   ContactsBlockRequestTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MyStoriesFrom",
			SchemaName: "my_stories_from",
			Null:       !b.Flags.Has(0),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (b *ContactsBlockRequest) SetFlags() {
	if !(b.MyStoriesFrom == false) {
		b.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (b *ContactsBlockRequest) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode contacts.block#2e2e8734 as nil")
	}
	buf.PutID(ContactsBlockRequestTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *ContactsBlockRequest) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode contacts.block#2e2e8734 as nil")
	}
	b.SetFlags()
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode contacts.block#2e2e8734: field flags: %w", err)
	}
	if b.ID == nil {
		return fmt.Errorf("unable to encode contacts.block#2e2e8734: field id is nil")
	}
	if err := b.ID.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode contacts.block#2e2e8734: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *ContactsBlockRequest) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode contacts.block#2e2e8734 to nil")
	}
	if err := buf.ConsumeID(ContactsBlockRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.block#2e2e8734: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *ContactsBlockRequest) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode contacts.block#2e2e8734 to nil")
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode contacts.block#2e2e8734: field flags: %w", err)
		}
	}
	b.MyStoriesFrom = b.Flags.Has(0)
	{
		value, err := DecodeInputPeer(buf)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.block#2e2e8734: field id: %w", err)
		}
		b.ID = value
	}
	return nil
}

// SetMyStoriesFrom sets value of MyStoriesFrom conditional field.
func (b *ContactsBlockRequest) SetMyStoriesFrom(value bool) {
	if value {
		b.Flags.Set(0)
		b.MyStoriesFrom = true
	} else {
		b.Flags.Unset(0)
		b.MyStoriesFrom = false
	}
}

// GetMyStoriesFrom returns value of MyStoriesFrom conditional field.
func (b *ContactsBlockRequest) GetMyStoriesFrom() (value bool) {
	if b == nil {
		return
	}
	return b.Flags.Has(0)
}

// GetID returns value of ID field.
func (b *ContactsBlockRequest) GetID() (value InputPeerClass) {
	if b == nil {
		return
	}
	return b.ID
}

// ContactsBlock invokes method contacts.block#2e2e8734 returning error if any.
// Adds a peer to a blocklist, see here »¹ for more info.
//
// Links:
//  1. https://core.telegram.org/api/block
//
// Possible errors:
//
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 CONTACT_ID_INVALID: The provided contact ID is invalid.
//	400 INPUT_USER_DEACTIVATED: The specified user was deleted.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/contacts.block for reference.
func (c *Client) ContactsBlock(ctx context.Context, request *ContactsBlockRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
