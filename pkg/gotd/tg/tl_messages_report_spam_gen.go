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

// MessagesReportSpamRequest represents TL type `messages.reportSpam#cf1592db`.
// Report a new incoming chat for spam, if the peer settings¹ of the chat allow us to do
// that
//
// Links:
//  1. https://core.telegram.org/constructor/peerSettings
//
// See https://core.telegram.org/method/messages.reportSpam for reference.
type MessagesReportSpamRequest struct {
	// Peer to report
	Peer InputPeerClass
}

// MessagesReportSpamRequestTypeID is TL type id of MessagesReportSpamRequest.
const MessagesReportSpamRequestTypeID = 0xcf1592db

// Ensuring interfaces in compile-time for MessagesReportSpamRequest.
var (
	_ bin.Encoder     = &MessagesReportSpamRequest{}
	_ bin.Decoder     = &MessagesReportSpamRequest{}
	_ bin.BareEncoder = &MessagesReportSpamRequest{}
	_ bin.BareDecoder = &MessagesReportSpamRequest{}
)

func (r *MessagesReportSpamRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReportSpamRequest) String() string {
	if r == nil {
		return "MessagesReportSpamRequest(nil)"
	}
	type Alias MessagesReportSpamRequest
	return fmt.Sprintf("MessagesReportSpamRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReportSpamRequest from given interface.
func (r *MessagesReportSpamRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
}) {
	r.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReportSpamRequest) TypeID() uint32 {
	return MessagesReportSpamRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReportSpamRequest) TypeName() string {
	return "messages.reportSpam"
}

// TypeInfo returns info about TL type.
func (r *MessagesReportSpamRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.reportSpam",
		ID:   MessagesReportSpamRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesReportSpamRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reportSpam#cf1592db as nil")
	}
	b.PutID(MessagesReportSpamRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReportSpamRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reportSpam#cf1592db as nil")
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.reportSpam#cf1592db: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.reportSpam#cf1592db: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReportSpamRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reportSpam#cf1592db to nil")
	}
	if err := b.ConsumeID(MessagesReportSpamRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.reportSpam#cf1592db: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReportSpamRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reportSpam#cf1592db to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.reportSpam#cf1592db: field peer: %w", err)
		}
		r.Peer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (r *MessagesReportSpamRequest) GetPeer() (value InputPeerClass) {
	if r == nil {
		return
	}
	return r.Peer
}

// MessagesReportSpam invokes method messages.reportSpam#cf1592db returning error if any.
// Report a new incoming chat for spam, if the peer settings¹ of the chat allow us to do
// that
//
// Links:
//  1. https://core.telegram.org/constructor/peerSettings
//
// Possible errors:
//
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.reportSpam for reference.
func (c *Client) MessagesReportSpam(ctx context.Context, peer InputPeerClass) (bool, error) {
	var result BoolBox

	request := &MessagesReportSpamRequest{
		Peer: peer,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
