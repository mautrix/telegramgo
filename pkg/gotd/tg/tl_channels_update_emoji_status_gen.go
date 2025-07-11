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

// ChannelsUpdateEmojiStatusRequest represents TL type `channels.updateEmojiStatus#f0d3e6a8`.
// Set an emoji status¹ for a channel or supergroup.
//
// Links:
//  1. https://core.telegram.org/api/emoji-status
//
// See https://core.telegram.org/method/channels.updateEmojiStatus for reference.
type ChannelsUpdateEmojiStatusRequest struct {
	// The channel/supergroup, must have at least
	// channel_emoji_status_level_min¹/group_emoji_status_level_min² boosts.
	//
	// Links:
	//  1) https://core.telegram.org/api/config#channel-emoji-status-level-min
	//  2) https://core.telegram.org/api/config#group-emoji-status-level-min
	Channel InputChannelClass
	// Emoji status¹ to set
	//
	// Links:
	//  1) https://core.telegram.org/api/emoji-status
	EmojiStatus EmojiStatusClass
}

// ChannelsUpdateEmojiStatusRequestTypeID is TL type id of ChannelsUpdateEmojiStatusRequest.
const ChannelsUpdateEmojiStatusRequestTypeID = 0xf0d3e6a8

// Ensuring interfaces in compile-time for ChannelsUpdateEmojiStatusRequest.
var (
	_ bin.Encoder     = &ChannelsUpdateEmojiStatusRequest{}
	_ bin.Decoder     = &ChannelsUpdateEmojiStatusRequest{}
	_ bin.BareEncoder = &ChannelsUpdateEmojiStatusRequest{}
	_ bin.BareDecoder = &ChannelsUpdateEmojiStatusRequest{}
)

func (u *ChannelsUpdateEmojiStatusRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Channel == nil) {
		return false
	}
	if !(u.EmojiStatus == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *ChannelsUpdateEmojiStatusRequest) String() string {
	if u == nil {
		return "ChannelsUpdateEmojiStatusRequest(nil)"
	}
	type Alias ChannelsUpdateEmojiStatusRequest
	return fmt.Sprintf("ChannelsUpdateEmojiStatusRequest%+v", Alias(*u))
}

// FillFrom fills ChannelsUpdateEmojiStatusRequest from given interface.
func (u *ChannelsUpdateEmojiStatusRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetEmojiStatus() (value EmojiStatusClass)
}) {
	u.Channel = from.GetChannel()
	u.EmojiStatus = from.GetEmojiStatus()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsUpdateEmojiStatusRequest) TypeID() uint32 {
	return ChannelsUpdateEmojiStatusRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsUpdateEmojiStatusRequest) TypeName() string {
	return "channels.updateEmojiStatus"
}

// TypeInfo returns info about TL type.
func (u *ChannelsUpdateEmojiStatusRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.updateEmojiStatus",
		ID:   ChannelsUpdateEmojiStatusRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "EmojiStatus",
			SchemaName: "emoji_status",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *ChannelsUpdateEmojiStatusRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode channels.updateEmojiStatus#f0d3e6a8 as nil")
	}
	b.PutID(ChannelsUpdateEmojiStatusRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *ChannelsUpdateEmojiStatusRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode channels.updateEmojiStatus#f0d3e6a8 as nil")
	}
	if u.Channel == nil {
		return fmt.Errorf("unable to encode channels.updateEmojiStatus#f0d3e6a8: field channel is nil")
	}
	if err := u.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.updateEmojiStatus#f0d3e6a8: field channel: %w", err)
	}
	if u.EmojiStatus == nil {
		return fmt.Errorf("unable to encode channels.updateEmojiStatus#f0d3e6a8: field emoji_status is nil")
	}
	if err := u.EmojiStatus.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.updateEmojiStatus#f0d3e6a8: field emoji_status: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *ChannelsUpdateEmojiStatusRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode channels.updateEmojiStatus#f0d3e6a8 to nil")
	}
	if err := b.ConsumeID(ChannelsUpdateEmojiStatusRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.updateEmojiStatus#f0d3e6a8: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *ChannelsUpdateEmojiStatusRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode channels.updateEmojiStatus#f0d3e6a8 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.updateEmojiStatus#f0d3e6a8: field channel: %w", err)
		}
		u.Channel = value
	}
	{
		value, err := DecodeEmojiStatus(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.updateEmojiStatus#f0d3e6a8: field emoji_status: %w", err)
		}
		u.EmojiStatus = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (u *ChannelsUpdateEmojiStatusRequest) GetChannel() (value InputChannelClass) {
	if u == nil {
		return
	}
	return u.Channel
}

// GetEmojiStatus returns value of EmojiStatus field.
func (u *ChannelsUpdateEmojiStatusRequest) GetEmojiStatus() (value EmojiStatusClass) {
	if u == nil {
		return
	}
	return u.EmojiStatus
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (u *ChannelsUpdateEmojiStatusRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return u.Channel.AsNotEmpty()
}

// GetEmojiStatusAsNotEmpty returns mapped value of EmojiStatus field.
func (u *ChannelsUpdateEmojiStatusRequest) GetEmojiStatusAsNotEmpty() (NotEmptyEmojiStatus, bool) {
	return u.EmojiStatus.AsNotEmpty()
}

// ChannelsUpdateEmojiStatus invokes method channels.updateEmojiStatus#f0d3e6a8 returning error if any.
// Set an emoji status¹ for a channel or supergroup.
//
// Links:
//  1. https://core.telegram.org/api/emoji-status
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//
// See https://core.telegram.org/method/channels.updateEmojiStatus for reference.
func (c *Client) ChannelsUpdateEmojiStatus(ctx context.Context, request *ChannelsUpdateEmojiStatusRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
