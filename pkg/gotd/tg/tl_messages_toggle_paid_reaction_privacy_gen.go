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

// MessagesTogglePaidReactionPrivacyRequest represents TL type `messages.togglePaidReactionPrivacy#435885b5`.
// Changes the privacy of already sent paid reactions¹ on a specific message.
//
// Links:
//  1. https://core.telegram.org/api/reactions#paid-reactions
//
// See https://core.telegram.org/method/messages.togglePaidReactionPrivacy for reference.
type MessagesTogglePaidReactionPrivacyRequest struct {
	// The channel
	Peer InputPeerClass
	// The ID of the message to which we sent the paid reactions
	MsgID int
	// If true, makes the current anonymous in the top sender leaderboard for this message;
	// otherwise, does the opposite.
	Private PaidReactionPrivacyClass
}

// MessagesTogglePaidReactionPrivacyRequestTypeID is TL type id of MessagesTogglePaidReactionPrivacyRequest.
const MessagesTogglePaidReactionPrivacyRequestTypeID = 0x435885b5

// Ensuring interfaces in compile-time for MessagesTogglePaidReactionPrivacyRequest.
var (
	_ bin.Encoder     = &MessagesTogglePaidReactionPrivacyRequest{}
	_ bin.Decoder     = &MessagesTogglePaidReactionPrivacyRequest{}
	_ bin.BareEncoder = &MessagesTogglePaidReactionPrivacyRequest{}
	_ bin.BareDecoder = &MessagesTogglePaidReactionPrivacyRequest{}
)

func (t *MessagesTogglePaidReactionPrivacyRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Peer == nil) {
		return false
	}
	if !(t.MsgID == 0) {
		return false
	}
	if !(t.Private == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *MessagesTogglePaidReactionPrivacyRequest) String() string {
	if t == nil {
		return "MessagesTogglePaidReactionPrivacyRequest(nil)"
	}
	type Alias MessagesTogglePaidReactionPrivacyRequest
	return fmt.Sprintf("MessagesTogglePaidReactionPrivacyRequest%+v", Alias(*t))
}

// FillFrom fills MessagesTogglePaidReactionPrivacyRequest from given interface.
func (t *MessagesTogglePaidReactionPrivacyRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetPrivate() (value PaidReactionPrivacyClass)
}) {
	t.Peer = from.GetPeer()
	t.MsgID = from.GetMsgID()
	t.Private = from.GetPrivate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesTogglePaidReactionPrivacyRequest) TypeID() uint32 {
	return MessagesTogglePaidReactionPrivacyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesTogglePaidReactionPrivacyRequest) TypeName() string {
	return "messages.togglePaidReactionPrivacy"
}

// TypeInfo returns info about TL type.
func (t *MessagesTogglePaidReactionPrivacyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.togglePaidReactionPrivacy",
		ID:   MessagesTogglePaidReactionPrivacyRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "Private",
			SchemaName: "private",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *MessagesTogglePaidReactionPrivacyRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.togglePaidReactionPrivacy#435885b5 as nil")
	}
	b.PutID(MessagesTogglePaidReactionPrivacyRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *MessagesTogglePaidReactionPrivacyRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.togglePaidReactionPrivacy#435885b5 as nil")
	}
	if t.Peer == nil {
		return fmt.Errorf("unable to encode messages.togglePaidReactionPrivacy#435885b5: field peer is nil")
	}
	if err := t.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.togglePaidReactionPrivacy#435885b5: field peer: %w", err)
	}
	b.PutInt(t.MsgID)
	if t.Private == nil {
		return fmt.Errorf("unable to encode messages.togglePaidReactionPrivacy#435885b5: field private is nil")
	}
	if err := t.Private.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.togglePaidReactionPrivacy#435885b5: field private: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *MessagesTogglePaidReactionPrivacyRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.togglePaidReactionPrivacy#435885b5 to nil")
	}
	if err := b.ConsumeID(MessagesTogglePaidReactionPrivacyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.togglePaidReactionPrivacy#435885b5: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *MessagesTogglePaidReactionPrivacyRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.togglePaidReactionPrivacy#435885b5 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.togglePaidReactionPrivacy#435885b5: field peer: %w", err)
		}
		t.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.togglePaidReactionPrivacy#435885b5: field msg_id: %w", err)
		}
		t.MsgID = value
	}
	{
		value, err := DecodePaidReactionPrivacy(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.togglePaidReactionPrivacy#435885b5: field private: %w", err)
		}
		t.Private = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (t *MessagesTogglePaidReactionPrivacyRequest) GetPeer() (value InputPeerClass) {
	if t == nil {
		return
	}
	return t.Peer
}

// GetMsgID returns value of MsgID field.
func (t *MessagesTogglePaidReactionPrivacyRequest) GetMsgID() (value int) {
	if t == nil {
		return
	}
	return t.MsgID
}

// GetPrivate returns value of Private field.
func (t *MessagesTogglePaidReactionPrivacyRequest) GetPrivate() (value PaidReactionPrivacyClass) {
	if t == nil {
		return
	}
	return t.Private
}

// MessagesTogglePaidReactionPrivacy invokes method messages.togglePaidReactionPrivacy#435885b5 returning error if any.
// Changes the privacy of already sent paid reactions¹ on a specific message.
//
// Links:
//  1. https://core.telegram.org/api/reactions#paid-reactions
//
// See https://core.telegram.org/method/messages.togglePaidReactionPrivacy for reference.
// Can be used by bots.
func (c *Client) MessagesTogglePaidReactionPrivacy(ctx context.Context, request *MessagesTogglePaidReactionPrivacyRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
