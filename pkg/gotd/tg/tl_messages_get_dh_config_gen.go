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

// MessagesGetDhConfigRequest represents TL type `messages.getDhConfig#26cf8950`.
// Returns configuration parameters for Diffie-Hellman key generation. Can also return a
// random sequence of bytes of required length.
//
// See https://core.telegram.org/method/messages.getDhConfig for reference.
type MessagesGetDhConfigRequest struct {
	// Value of the version parameter from messages.dhConfig¹, available at the client
	//
	// Links:
	//  1) https://core.telegram.org/constructor/messages.dhConfig
	Version int
	// Length of the required random sequence
	RandomLength int
}

// MessagesGetDhConfigRequestTypeID is TL type id of MessagesGetDhConfigRequest.
const MessagesGetDhConfigRequestTypeID = 0x26cf8950

// Ensuring interfaces in compile-time for MessagesGetDhConfigRequest.
var (
	_ bin.Encoder     = &MessagesGetDhConfigRequest{}
	_ bin.Decoder     = &MessagesGetDhConfigRequest{}
	_ bin.BareEncoder = &MessagesGetDhConfigRequest{}
	_ bin.BareDecoder = &MessagesGetDhConfigRequest{}
)

func (g *MessagesGetDhConfigRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Version == 0) {
		return false
	}
	if !(g.RandomLength == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetDhConfigRequest) String() string {
	if g == nil {
		return "MessagesGetDhConfigRequest(nil)"
	}
	type Alias MessagesGetDhConfigRequest
	return fmt.Sprintf("MessagesGetDhConfigRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetDhConfigRequest from given interface.
func (g *MessagesGetDhConfigRequest) FillFrom(from interface {
	GetVersion() (value int)
	GetRandomLength() (value int)
}) {
	g.Version = from.GetVersion()
	g.RandomLength = from.GetRandomLength()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetDhConfigRequest) TypeID() uint32 {
	return MessagesGetDhConfigRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetDhConfigRequest) TypeName() string {
	return "messages.getDhConfig"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetDhConfigRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getDhConfig",
		ID:   MessagesGetDhConfigRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Version",
			SchemaName: "version",
		},
		{
			Name:       "RandomLength",
			SchemaName: "random_length",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetDhConfigRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDhConfig#26cf8950 as nil")
	}
	b.PutID(MessagesGetDhConfigRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetDhConfigRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDhConfig#26cf8950 as nil")
	}
	b.PutInt(g.Version)
	b.PutInt(g.RandomLength)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetDhConfigRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDhConfig#26cf8950 to nil")
	}
	if err := b.ConsumeID(MessagesGetDhConfigRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDhConfig#26cf8950: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetDhConfigRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDhConfig#26cf8950 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDhConfig#26cf8950: field version: %w", err)
		}
		g.Version = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDhConfig#26cf8950: field random_length: %w", err)
		}
		g.RandomLength = value
	}
	return nil
}

// GetVersion returns value of Version field.
func (g *MessagesGetDhConfigRequest) GetVersion() (value int) {
	if g == nil {
		return
	}
	return g.Version
}

// GetRandomLength returns value of RandomLength field.
func (g *MessagesGetDhConfigRequest) GetRandomLength() (value int) {
	if g == nil {
		return
	}
	return g.RandomLength
}

// MessagesGetDhConfig invokes method messages.getDhConfig#26cf8950 returning error if any.
// Returns configuration parameters for Diffie-Hellman key generation. Can also return a
// random sequence of bytes of required length.
//
// Possible errors:
//
//	400 RANDOM_LENGTH_INVALID: Random length invalid.
//
// See https://core.telegram.org/method/messages.getDhConfig for reference.
func (c *Client) MessagesGetDhConfig(ctx context.Context, request *MessagesGetDhConfigRequest) (MessagesDhConfigClass, error) {
	var result MessagesDhConfigBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.DhConfig, nil
}
