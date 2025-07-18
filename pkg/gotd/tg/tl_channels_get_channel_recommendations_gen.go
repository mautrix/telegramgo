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

// ChannelsGetChannelRecommendationsRequest represents TL type `channels.getChannelRecommendations#25a71742`.
// Obtain a list of similarly themed public channels, selected based on similarities in
// their subscriber bases.
//
// See https://core.telegram.org/method/channels.getChannelRecommendations for reference.
type ChannelsGetChannelRecommendationsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// The method will return channels related to the passed channel. If not set, the method
	// will returns channels related to channels the user has joined.
	//
	// Use SetChannel and GetChannel helpers.
	Channel InputChannelClass
}

// ChannelsGetChannelRecommendationsRequestTypeID is TL type id of ChannelsGetChannelRecommendationsRequest.
const ChannelsGetChannelRecommendationsRequestTypeID = 0x25a71742

// Ensuring interfaces in compile-time for ChannelsGetChannelRecommendationsRequest.
var (
	_ bin.Encoder     = &ChannelsGetChannelRecommendationsRequest{}
	_ bin.Decoder     = &ChannelsGetChannelRecommendationsRequest{}
	_ bin.BareEncoder = &ChannelsGetChannelRecommendationsRequest{}
	_ bin.BareDecoder = &ChannelsGetChannelRecommendationsRequest{}
)

func (g *ChannelsGetChannelRecommendationsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Channel == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetChannelRecommendationsRequest) String() string {
	if g == nil {
		return "ChannelsGetChannelRecommendationsRequest(nil)"
	}
	type Alias ChannelsGetChannelRecommendationsRequest
	return fmt.Sprintf("ChannelsGetChannelRecommendationsRequest%+v", Alias(*g))
}

// FillFrom fills ChannelsGetChannelRecommendationsRequest from given interface.
func (g *ChannelsGetChannelRecommendationsRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass, ok bool)
}) {
	if val, ok := from.GetChannel(); ok {
		g.Channel = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsGetChannelRecommendationsRequest) TypeID() uint32 {
	return ChannelsGetChannelRecommendationsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsGetChannelRecommendationsRequest) TypeName() string {
	return "channels.getChannelRecommendations"
}

// TypeInfo returns info about TL type.
func (g *ChannelsGetChannelRecommendationsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.getChannelRecommendations",
		ID:   ChannelsGetChannelRecommendationsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
			Null:       !g.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *ChannelsGetChannelRecommendationsRequest) SetFlags() {
	if !(g.Channel == nil) {
		g.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (g *ChannelsGetChannelRecommendationsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getChannelRecommendations#25a71742 as nil")
	}
	b.PutID(ChannelsGetChannelRecommendationsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ChannelsGetChannelRecommendationsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getChannelRecommendations#25a71742 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getChannelRecommendations#25a71742: field flags: %w", err)
	}
	if g.Flags.Has(0) {
		if g.Channel == nil {
			return fmt.Errorf("unable to encode channels.getChannelRecommendations#25a71742: field channel is nil")
		}
		if err := g.Channel.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.getChannelRecommendations#25a71742: field channel: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetChannelRecommendationsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getChannelRecommendations#25a71742 to nil")
	}
	if err := b.ConsumeID(ChannelsGetChannelRecommendationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getChannelRecommendations#25a71742: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ChannelsGetChannelRecommendationsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getChannelRecommendations#25a71742 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.getChannelRecommendations#25a71742: field flags: %w", err)
		}
	}
	if g.Flags.Has(0) {
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getChannelRecommendations#25a71742: field channel: %w", err)
		}
		g.Channel = value
	}
	return nil
}

// SetChannel sets value of Channel conditional field.
func (g *ChannelsGetChannelRecommendationsRequest) SetChannel(value InputChannelClass) {
	g.Flags.Set(0)
	g.Channel = value
}

// GetChannel returns value of Channel conditional field and
// boolean which is true if field was set.
func (g *ChannelsGetChannelRecommendationsRequest) GetChannel() (value InputChannelClass, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.Channel, true
}

// GetChannelAsNotEmpty returns mapped value of Channel conditional field and
// boolean which is true if field was set.
func (g *ChannelsGetChannelRecommendationsRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	if value, ok := g.GetChannel(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// ChannelsGetChannelRecommendations invokes method channels.getChannelRecommendations#25a71742 returning error if any.
// Obtain a list of similarly themed public channels, selected based on similarities in
// their subscriber bases.
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//
// See https://core.telegram.org/method/channels.getChannelRecommendations for reference.
func (c *Client) ChannelsGetChannelRecommendations(ctx context.Context, request *ChannelsGetChannelRecommendationsRequest) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
