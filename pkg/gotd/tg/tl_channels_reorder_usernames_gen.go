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

// ChannelsReorderUsernamesRequest represents TL type `channels.reorderUsernames#b45ced1d`.
// Reorder active usernames
//
// See https://core.telegram.org/method/channels.reorderUsernames for reference.
type ChannelsReorderUsernamesRequest struct {
	// The supergroup or channel
	Channel InputChannelClass
	// The new order for active usernames. All active usernames must be specified.
	Order []string
}

// ChannelsReorderUsernamesRequestTypeID is TL type id of ChannelsReorderUsernamesRequest.
const ChannelsReorderUsernamesRequestTypeID = 0xb45ced1d

// Ensuring interfaces in compile-time for ChannelsReorderUsernamesRequest.
var (
	_ bin.Encoder     = &ChannelsReorderUsernamesRequest{}
	_ bin.Decoder     = &ChannelsReorderUsernamesRequest{}
	_ bin.BareEncoder = &ChannelsReorderUsernamesRequest{}
	_ bin.BareDecoder = &ChannelsReorderUsernamesRequest{}
)

func (r *ChannelsReorderUsernamesRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Channel == nil) {
		return false
	}
	if !(r.Order == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ChannelsReorderUsernamesRequest) String() string {
	if r == nil {
		return "ChannelsReorderUsernamesRequest(nil)"
	}
	type Alias ChannelsReorderUsernamesRequest
	return fmt.Sprintf("ChannelsReorderUsernamesRequest%+v", Alias(*r))
}

// FillFrom fills ChannelsReorderUsernamesRequest from given interface.
func (r *ChannelsReorderUsernamesRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetOrder() (value []string)
}) {
	r.Channel = from.GetChannel()
	r.Order = from.GetOrder()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsReorderUsernamesRequest) TypeID() uint32 {
	return ChannelsReorderUsernamesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsReorderUsernamesRequest) TypeName() string {
	return "channels.reorderUsernames"
}

// TypeInfo returns info about TL type.
func (r *ChannelsReorderUsernamesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.reorderUsernames",
		ID:   ChannelsReorderUsernamesRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Order",
			SchemaName: "order",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ChannelsReorderUsernamesRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode channels.reorderUsernames#b45ced1d as nil")
	}
	b.PutID(ChannelsReorderUsernamesRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ChannelsReorderUsernamesRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode channels.reorderUsernames#b45ced1d as nil")
	}
	if r.Channel == nil {
		return fmt.Errorf("unable to encode channels.reorderUsernames#b45ced1d: field channel is nil")
	}
	if err := r.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.reorderUsernames#b45ced1d: field channel: %w", err)
	}
	b.PutVectorHeader(len(r.Order))
	for _, v := range r.Order {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ChannelsReorderUsernamesRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode channels.reorderUsernames#b45ced1d to nil")
	}
	if err := b.ConsumeID(ChannelsReorderUsernamesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.reorderUsernames#b45ced1d: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ChannelsReorderUsernamesRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode channels.reorderUsernames#b45ced1d to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.reorderUsernames#b45ced1d: field channel: %w", err)
		}
		r.Channel = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.reorderUsernames#b45ced1d: field order: %w", err)
		}

		if headerLen > 0 {
			r.Order = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode channels.reorderUsernames#b45ced1d: field order: %w", err)
			}
			r.Order = append(r.Order, value)
		}
	}
	return nil
}

// GetChannel returns value of Channel field.
func (r *ChannelsReorderUsernamesRequest) GetChannel() (value InputChannelClass) {
	if r == nil {
		return
	}
	return r.Channel
}

// GetOrder returns value of Order field.
func (r *ChannelsReorderUsernamesRequest) GetOrder() (value []string) {
	if r == nil {
		return
	}
	return r.Order
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (r *ChannelsReorderUsernamesRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return r.Channel.AsNotEmpty()
}

// ChannelsReorderUsernames invokes method channels.reorderUsernames#b45ced1d returning error if any.
// Reorder active usernames
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHAT_NOT_MODIFIED: No changes were made to chat information because the new information you passed is identical to the current information.
//
// See https://core.telegram.org/method/channels.reorderUsernames for reference.
func (c *Client) ChannelsReorderUsernames(ctx context.Context, request *ChannelsReorderUsernamesRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
