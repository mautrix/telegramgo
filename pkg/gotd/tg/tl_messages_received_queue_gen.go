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

// MessagesReceivedQueueRequest represents TL type `messages.receivedQueue#55a5bb66`.
// Confirms receipt of messages in a secret chat by client, cancels push notifications.
// The method returns a list of random_ids of messages for which push notifications were
// cancelled.
//
// See https://core.telegram.org/method/messages.receivedQueue for reference.
type MessagesReceivedQueueRequest struct {
	// Maximum qts value available at the client
	MaxQts int
}

// MessagesReceivedQueueRequestTypeID is TL type id of MessagesReceivedQueueRequest.
const MessagesReceivedQueueRequestTypeID = 0x55a5bb66

// Ensuring interfaces in compile-time for MessagesReceivedQueueRequest.
var (
	_ bin.Encoder     = &MessagesReceivedQueueRequest{}
	_ bin.Decoder     = &MessagesReceivedQueueRequest{}
	_ bin.BareEncoder = &MessagesReceivedQueueRequest{}
	_ bin.BareDecoder = &MessagesReceivedQueueRequest{}
)

func (r *MessagesReceivedQueueRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.MaxQts == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReceivedQueueRequest) String() string {
	if r == nil {
		return "MessagesReceivedQueueRequest(nil)"
	}
	type Alias MessagesReceivedQueueRequest
	return fmt.Sprintf("MessagesReceivedQueueRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReceivedQueueRequest from given interface.
func (r *MessagesReceivedQueueRequest) FillFrom(from interface {
	GetMaxQts() (value int)
}) {
	r.MaxQts = from.GetMaxQts()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReceivedQueueRequest) TypeID() uint32 {
	return MessagesReceivedQueueRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReceivedQueueRequest) TypeName() string {
	return "messages.receivedQueue"
}

// TypeInfo returns info about TL type.
func (r *MessagesReceivedQueueRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.receivedQueue",
		ID:   MessagesReceivedQueueRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MaxQts",
			SchemaName: "max_qts",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesReceivedQueueRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.receivedQueue#55a5bb66 as nil")
	}
	b.PutID(MessagesReceivedQueueRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReceivedQueueRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.receivedQueue#55a5bb66 as nil")
	}
	b.PutInt(r.MaxQts)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReceivedQueueRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.receivedQueue#55a5bb66 to nil")
	}
	if err := b.ConsumeID(MessagesReceivedQueueRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.receivedQueue#55a5bb66: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReceivedQueueRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.receivedQueue#55a5bb66 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.receivedQueue#55a5bb66: field max_qts: %w", err)
		}
		r.MaxQts = value
	}
	return nil
}

// GetMaxQts returns value of MaxQts field.
func (r *MessagesReceivedQueueRequest) GetMaxQts() (value int) {
	if r == nil {
		return
	}
	return r.MaxQts
}

// MessagesReceivedQueue invokes method messages.receivedQueue#55a5bb66 returning error if any.
// Confirms receipt of messages in a secret chat by client, cancels push notifications.
// The method returns a list of random_ids of messages for which push notifications were
// cancelled.
//
// Possible errors:
//
//	400 MAX_QTS_INVALID: The specified max_qts is invalid.
//	500 MSG_WAIT_FAILED: A waiting call returned an error.
//
// See https://core.telegram.org/method/messages.receivedQueue for reference.
func (c *Client) MessagesReceivedQueue(ctx context.Context, maxqts int) ([]int64, error) {
	var result LongVector

	request := &MessagesReceivedQueueRequest{
		MaxQts: maxqts,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []int64(result.Elems), nil
}
