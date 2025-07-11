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

// MessagesToggleDialogPinRequest represents TL type `messages.toggleDialogPin#a731e257`.
// Pin/unpin a dialog
//
// See https://core.telegram.org/method/messages.toggleDialogPin for reference.
type MessagesToggleDialogPinRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to pin or unpin the dialog
	Pinned bool
	// The dialog to pin
	Peer InputDialogPeerClass
}

// MessagesToggleDialogPinRequestTypeID is TL type id of MessagesToggleDialogPinRequest.
const MessagesToggleDialogPinRequestTypeID = 0xa731e257

// Ensuring interfaces in compile-time for MessagesToggleDialogPinRequest.
var (
	_ bin.Encoder     = &MessagesToggleDialogPinRequest{}
	_ bin.Decoder     = &MessagesToggleDialogPinRequest{}
	_ bin.BareEncoder = &MessagesToggleDialogPinRequest{}
	_ bin.BareDecoder = &MessagesToggleDialogPinRequest{}
)

func (t *MessagesToggleDialogPinRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.Pinned == false) {
		return false
	}
	if !(t.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *MessagesToggleDialogPinRequest) String() string {
	if t == nil {
		return "MessagesToggleDialogPinRequest(nil)"
	}
	type Alias MessagesToggleDialogPinRequest
	return fmt.Sprintf("MessagesToggleDialogPinRequest%+v", Alias(*t))
}

// FillFrom fills MessagesToggleDialogPinRequest from given interface.
func (t *MessagesToggleDialogPinRequest) FillFrom(from interface {
	GetPinned() (value bool)
	GetPeer() (value InputDialogPeerClass)
}) {
	t.Pinned = from.GetPinned()
	t.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesToggleDialogPinRequest) TypeID() uint32 {
	return MessagesToggleDialogPinRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesToggleDialogPinRequest) TypeName() string {
	return "messages.toggleDialogPin"
}

// TypeInfo returns info about TL type.
func (t *MessagesToggleDialogPinRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.toggleDialogPin",
		ID:   MessagesToggleDialogPinRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Pinned",
			SchemaName: "pinned",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (t *MessagesToggleDialogPinRequest) SetFlags() {
	if !(t.Pinned == false) {
		t.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (t *MessagesToggleDialogPinRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.toggleDialogPin#a731e257 as nil")
	}
	b.PutID(MessagesToggleDialogPinRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *MessagesToggleDialogPinRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.toggleDialogPin#a731e257 as nil")
	}
	t.SetFlags()
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.toggleDialogPin#a731e257: field flags: %w", err)
	}
	if t.Peer == nil {
		return fmt.Errorf("unable to encode messages.toggleDialogPin#a731e257: field peer is nil")
	}
	if err := t.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.toggleDialogPin#a731e257: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *MessagesToggleDialogPinRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.toggleDialogPin#a731e257 to nil")
	}
	if err := b.ConsumeID(MessagesToggleDialogPinRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.toggleDialogPin#a731e257: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *MessagesToggleDialogPinRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.toggleDialogPin#a731e257 to nil")
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.toggleDialogPin#a731e257: field flags: %w", err)
		}
	}
	t.Pinned = t.Flags.Has(0)
	{
		value, err := DecodeInputDialogPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.toggleDialogPin#a731e257: field peer: %w", err)
		}
		t.Peer = value
	}
	return nil
}

// SetPinned sets value of Pinned conditional field.
func (t *MessagesToggleDialogPinRequest) SetPinned(value bool) {
	if value {
		t.Flags.Set(0)
		t.Pinned = true
	} else {
		t.Flags.Unset(0)
		t.Pinned = false
	}
}

// GetPinned returns value of Pinned conditional field.
func (t *MessagesToggleDialogPinRequest) GetPinned() (value bool) {
	if t == nil {
		return
	}
	return t.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (t *MessagesToggleDialogPinRequest) GetPeer() (value InputDialogPeerClass) {
	if t == nil {
		return
	}
	return t.Peer
}

// MessagesToggleDialogPin invokes method messages.toggleDialogPin#a731e257 returning error if any.
// Pin/unpin a dialog
//
// Possible errors:
//
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 PEER_HISTORY_EMPTY: You can't pin an empty chat with a user.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 PINNED_DIALOGS_TOO_MUCH: Too many pinned dialogs.
//
// See https://core.telegram.org/method/messages.toggleDialogPin for reference.
func (c *Client) MessagesToggleDialogPin(ctx context.Context, request *MessagesToggleDialogPinRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
