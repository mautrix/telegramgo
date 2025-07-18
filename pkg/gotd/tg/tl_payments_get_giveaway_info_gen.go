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

// PaymentsGetGiveawayInfoRequest represents TL type `payments.getGiveawayInfo#f4239425`.
// Obtain information about a Telegram Premium giveaway »¹.
//
// Links:
//  1. https://core.telegram.org/api/giveaways
//
// See https://core.telegram.org/method/payments.getGiveawayInfo for reference.
type PaymentsGetGiveawayInfoRequest struct {
	// The peer where the giveaway was posted.
	Peer InputPeerClass
	// Message ID of the messageActionGiveawayLaunch¹ service message
	//
	// Links:
	//  1) https://core.telegram.org/constructor/messageActionGiveawayLaunch
	MsgID int
}

// PaymentsGetGiveawayInfoRequestTypeID is TL type id of PaymentsGetGiveawayInfoRequest.
const PaymentsGetGiveawayInfoRequestTypeID = 0xf4239425

// Ensuring interfaces in compile-time for PaymentsGetGiveawayInfoRequest.
var (
	_ bin.Encoder     = &PaymentsGetGiveawayInfoRequest{}
	_ bin.Decoder     = &PaymentsGetGiveawayInfoRequest{}
	_ bin.BareEncoder = &PaymentsGetGiveawayInfoRequest{}
	_ bin.BareDecoder = &PaymentsGetGiveawayInfoRequest{}
)

func (g *PaymentsGetGiveawayInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PaymentsGetGiveawayInfoRequest) String() string {
	if g == nil {
		return "PaymentsGetGiveawayInfoRequest(nil)"
	}
	type Alias PaymentsGetGiveawayInfoRequest
	return fmt.Sprintf("PaymentsGetGiveawayInfoRequest%+v", Alias(*g))
}

// FillFrom fills PaymentsGetGiveawayInfoRequest from given interface.
func (g *PaymentsGetGiveawayInfoRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
}) {
	g.Peer = from.GetPeer()
	g.MsgID = from.GetMsgID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsGetGiveawayInfoRequest) TypeID() uint32 {
	return PaymentsGetGiveawayInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsGetGiveawayInfoRequest) TypeName() string {
	return "payments.getGiveawayInfo"
}

// TypeInfo returns info about TL type.
func (g *PaymentsGetGiveawayInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.getGiveawayInfo",
		ID:   PaymentsGetGiveawayInfoRequestTypeID,
	}
	if g == nil {
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PaymentsGetGiveawayInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getGiveawayInfo#f4239425 as nil")
	}
	b.PutID(PaymentsGetGiveawayInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PaymentsGetGiveawayInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getGiveawayInfo#f4239425 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode payments.getGiveawayInfo#f4239425: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.getGiveawayInfo#f4239425: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	return nil
}

// Decode implements bin.Decoder.
func (g *PaymentsGetGiveawayInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getGiveawayInfo#f4239425 to nil")
	}
	if err := b.ConsumeID(PaymentsGetGiveawayInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getGiveawayInfo#f4239425: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PaymentsGetGiveawayInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getGiveawayInfo#f4239425 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.getGiveawayInfo#f4239425: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getGiveawayInfo#f4239425: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *PaymentsGetGiveawayInfoRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetMsgID returns value of MsgID field.
func (g *PaymentsGetGiveawayInfoRequest) GetMsgID() (value int) {
	if g == nil {
		return
	}
	return g.MsgID
}

// PaymentsGetGiveawayInfo invokes method payments.getGiveawayInfo#f4239425 returning error if any.
// Obtain information about a Telegram Premium giveaway »¹.
//
// Links:
//  1. https://core.telegram.org/api/giveaways
//
// Possible errors:
//
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/payments.getGiveawayInfo for reference.
func (c *Client) PaymentsGetGiveawayInfo(ctx context.Context, request *PaymentsGetGiveawayInfoRequest) (PaymentsGiveawayInfoClass, error) {
	var result PaymentsGiveawayInfoBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.GiveawayInfo, nil
}
