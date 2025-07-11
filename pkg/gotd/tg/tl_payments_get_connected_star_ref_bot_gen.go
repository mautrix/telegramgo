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

// PaymentsGetConnectedStarRefBotRequest represents TL type `payments.getConnectedStarRefBot#b7d998f0`.
//
// See https://core.telegram.org/method/payments.getConnectedStarRefBot for reference.
type PaymentsGetConnectedStarRefBotRequest struct {
	// Peer field of PaymentsGetConnectedStarRefBotRequest.
	Peer InputPeerClass
	// Bot field of PaymentsGetConnectedStarRefBotRequest.
	Bot InputUserClass
}

// PaymentsGetConnectedStarRefBotRequestTypeID is TL type id of PaymentsGetConnectedStarRefBotRequest.
const PaymentsGetConnectedStarRefBotRequestTypeID = 0xb7d998f0

// Ensuring interfaces in compile-time for PaymentsGetConnectedStarRefBotRequest.
var (
	_ bin.Encoder     = &PaymentsGetConnectedStarRefBotRequest{}
	_ bin.Decoder     = &PaymentsGetConnectedStarRefBotRequest{}
	_ bin.BareEncoder = &PaymentsGetConnectedStarRefBotRequest{}
	_ bin.BareDecoder = &PaymentsGetConnectedStarRefBotRequest{}
)

func (g *PaymentsGetConnectedStarRefBotRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.Bot == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PaymentsGetConnectedStarRefBotRequest) String() string {
	if g == nil {
		return "PaymentsGetConnectedStarRefBotRequest(nil)"
	}
	type Alias PaymentsGetConnectedStarRefBotRequest
	return fmt.Sprintf("PaymentsGetConnectedStarRefBotRequest%+v", Alias(*g))
}

// FillFrom fills PaymentsGetConnectedStarRefBotRequest from given interface.
func (g *PaymentsGetConnectedStarRefBotRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetBot() (value InputUserClass)
}) {
	g.Peer = from.GetPeer()
	g.Bot = from.GetBot()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsGetConnectedStarRefBotRequest) TypeID() uint32 {
	return PaymentsGetConnectedStarRefBotRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsGetConnectedStarRefBotRequest) TypeName() string {
	return "payments.getConnectedStarRefBot"
}

// TypeInfo returns info about TL type.
func (g *PaymentsGetConnectedStarRefBotRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.getConnectedStarRefBot",
		ID:   PaymentsGetConnectedStarRefBotRequestTypeID,
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
			Name:       "Bot",
			SchemaName: "bot",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PaymentsGetConnectedStarRefBotRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getConnectedStarRefBot#b7d998f0 as nil")
	}
	b.PutID(PaymentsGetConnectedStarRefBotRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PaymentsGetConnectedStarRefBotRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getConnectedStarRefBot#b7d998f0 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode payments.getConnectedStarRefBot#b7d998f0: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.getConnectedStarRefBot#b7d998f0: field peer: %w", err)
	}
	if g.Bot == nil {
		return fmt.Errorf("unable to encode payments.getConnectedStarRefBot#b7d998f0: field bot is nil")
	}
	if err := g.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.getConnectedStarRefBot#b7d998f0: field bot: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *PaymentsGetConnectedStarRefBotRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getConnectedStarRefBot#b7d998f0 to nil")
	}
	if err := b.ConsumeID(PaymentsGetConnectedStarRefBotRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getConnectedStarRefBot#b7d998f0: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PaymentsGetConnectedStarRefBotRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getConnectedStarRefBot#b7d998f0 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.getConnectedStarRefBot#b7d998f0: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.getConnectedStarRefBot#b7d998f0: field bot: %w", err)
		}
		g.Bot = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *PaymentsGetConnectedStarRefBotRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetBot returns value of Bot field.
func (g *PaymentsGetConnectedStarRefBotRequest) GetBot() (value InputUserClass) {
	if g == nil {
		return
	}
	return g.Bot
}

// PaymentsGetConnectedStarRefBot invokes method payments.getConnectedStarRefBot#b7d998f0 returning error if any.
//
// See https://core.telegram.org/method/payments.getConnectedStarRefBot for reference.
func (c *Client) PaymentsGetConnectedStarRefBot(ctx context.Context, request *PaymentsGetConnectedStarRefBotRequest) (*PaymentsConnectedStarRefBots, error) {
	var result PaymentsConnectedStarRefBots

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
