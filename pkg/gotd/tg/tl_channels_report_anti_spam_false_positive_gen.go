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

// ChannelsReportAntiSpamFalsePositiveRequest represents TL type `channels.reportAntiSpamFalsePositive#a850a693`.
// Report a native antispam¹ false positive
//
// Links:
//  1. https://core.telegram.org/api/antispam
//
// See https://core.telegram.org/method/channels.reportAntiSpamFalsePositive for reference.
type ChannelsReportAntiSpamFalsePositiveRequest struct {
	// Supergroup ID
	Channel InputChannelClass
	// Message ID that was mistakenly deleted by the native antispam¹ system, taken from the
	// admin log²
	//
	// Links:
	//  1) https://core.telegram.org/api/antispam
	//  2) https://core.telegram.org/api/recent-actions
	MsgID int
}

// ChannelsReportAntiSpamFalsePositiveRequestTypeID is TL type id of ChannelsReportAntiSpamFalsePositiveRequest.
const ChannelsReportAntiSpamFalsePositiveRequestTypeID = 0xa850a693

// Ensuring interfaces in compile-time for ChannelsReportAntiSpamFalsePositiveRequest.
var (
	_ bin.Encoder     = &ChannelsReportAntiSpamFalsePositiveRequest{}
	_ bin.Decoder     = &ChannelsReportAntiSpamFalsePositiveRequest{}
	_ bin.BareEncoder = &ChannelsReportAntiSpamFalsePositiveRequest{}
	_ bin.BareDecoder = &ChannelsReportAntiSpamFalsePositiveRequest{}
)

func (r *ChannelsReportAntiSpamFalsePositiveRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Channel == nil) {
		return false
	}
	if !(r.MsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ChannelsReportAntiSpamFalsePositiveRequest) String() string {
	if r == nil {
		return "ChannelsReportAntiSpamFalsePositiveRequest(nil)"
	}
	type Alias ChannelsReportAntiSpamFalsePositiveRequest
	return fmt.Sprintf("ChannelsReportAntiSpamFalsePositiveRequest%+v", Alias(*r))
}

// FillFrom fills ChannelsReportAntiSpamFalsePositiveRequest from given interface.
func (r *ChannelsReportAntiSpamFalsePositiveRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetMsgID() (value int)
}) {
	r.Channel = from.GetChannel()
	r.MsgID = from.GetMsgID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsReportAntiSpamFalsePositiveRequest) TypeID() uint32 {
	return ChannelsReportAntiSpamFalsePositiveRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsReportAntiSpamFalsePositiveRequest) TypeName() string {
	return "channels.reportAntiSpamFalsePositive"
}

// TypeInfo returns info about TL type.
func (r *ChannelsReportAntiSpamFalsePositiveRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.reportAntiSpamFalsePositive",
		ID:   ChannelsReportAntiSpamFalsePositiveRequestTypeID,
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
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ChannelsReportAntiSpamFalsePositiveRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode channels.reportAntiSpamFalsePositive#a850a693 as nil")
	}
	b.PutID(ChannelsReportAntiSpamFalsePositiveRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ChannelsReportAntiSpamFalsePositiveRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode channels.reportAntiSpamFalsePositive#a850a693 as nil")
	}
	if r.Channel == nil {
		return fmt.Errorf("unable to encode channels.reportAntiSpamFalsePositive#a850a693: field channel is nil")
	}
	if err := r.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.reportAntiSpamFalsePositive#a850a693: field channel: %w", err)
	}
	b.PutInt(r.MsgID)
	return nil
}

// Decode implements bin.Decoder.
func (r *ChannelsReportAntiSpamFalsePositiveRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode channels.reportAntiSpamFalsePositive#a850a693 to nil")
	}
	if err := b.ConsumeID(ChannelsReportAntiSpamFalsePositiveRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.reportAntiSpamFalsePositive#a850a693: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ChannelsReportAntiSpamFalsePositiveRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode channels.reportAntiSpamFalsePositive#a850a693 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.reportAntiSpamFalsePositive#a850a693: field channel: %w", err)
		}
		r.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.reportAntiSpamFalsePositive#a850a693: field msg_id: %w", err)
		}
		r.MsgID = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (r *ChannelsReportAntiSpamFalsePositiveRequest) GetChannel() (value InputChannelClass) {
	if r == nil {
		return
	}
	return r.Channel
}

// GetMsgID returns value of MsgID field.
func (r *ChannelsReportAntiSpamFalsePositiveRequest) GetMsgID() (value int) {
	if r == nil {
		return
	}
	return r.MsgID
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (r *ChannelsReportAntiSpamFalsePositiveRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return r.Channel.AsNotEmpty()
}

// ChannelsReportAntiSpamFalsePositive invokes method channels.reportAntiSpamFalsePositive#a850a693 returning error if any.
// Report a native antispam¹ false positive
//
// Links:
//  1. https://core.telegram.org/api/antispam
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//
// See https://core.telegram.org/method/channels.reportAntiSpamFalsePositive for reference.
func (c *Client) ChannelsReportAntiSpamFalsePositive(ctx context.Context, request *ChannelsReportAntiSpamFalsePositiveRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
