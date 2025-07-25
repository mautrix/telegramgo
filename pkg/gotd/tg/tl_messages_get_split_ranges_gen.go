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

// MessagesGetSplitRangesRequest represents TL type `messages.getSplitRanges#1cff7e08`.
// Get message ranges for saving the user's chat history
//
// See https://core.telegram.org/method/messages.getSplitRanges for reference.
type MessagesGetSplitRangesRequest struct {
}

// MessagesGetSplitRangesRequestTypeID is TL type id of MessagesGetSplitRangesRequest.
const MessagesGetSplitRangesRequestTypeID = 0x1cff7e08

// Ensuring interfaces in compile-time for MessagesGetSplitRangesRequest.
var (
	_ bin.Encoder     = &MessagesGetSplitRangesRequest{}
	_ bin.Decoder     = &MessagesGetSplitRangesRequest{}
	_ bin.BareEncoder = &MessagesGetSplitRangesRequest{}
	_ bin.BareDecoder = &MessagesGetSplitRangesRequest{}
)

func (g *MessagesGetSplitRangesRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetSplitRangesRequest) String() string {
	if g == nil {
		return "MessagesGetSplitRangesRequest(nil)"
	}
	type Alias MessagesGetSplitRangesRequest
	return fmt.Sprintf("MessagesGetSplitRangesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetSplitRangesRequest) TypeID() uint32 {
	return MessagesGetSplitRangesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetSplitRangesRequest) TypeName() string {
	return "messages.getSplitRanges"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetSplitRangesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getSplitRanges",
		ID:   MessagesGetSplitRangesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetSplitRangesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getSplitRanges#1cff7e08 as nil")
	}
	b.PutID(MessagesGetSplitRangesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetSplitRangesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getSplitRanges#1cff7e08 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetSplitRangesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getSplitRanges#1cff7e08 to nil")
	}
	if err := b.ConsumeID(MessagesGetSplitRangesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getSplitRanges#1cff7e08: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetSplitRangesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getSplitRanges#1cff7e08 to nil")
	}
	return nil
}

// MessagesGetSplitRanges invokes method messages.getSplitRanges#1cff7e08 returning error if any.
// Get message ranges for saving the user's chat history
//
// See https://core.telegram.org/method/messages.getSplitRanges for reference.
func (c *Client) MessagesGetSplitRanges(ctx context.Context) ([]MessageRange, error) {
	var result MessageRangeVector

	request := &MessagesGetSplitRangesRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []MessageRange(result.Elems), nil
}
