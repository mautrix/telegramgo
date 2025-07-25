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

// MessagesSetDefaultReactionRequest represents TL type `messages.setDefaultReaction#4f47a016`.
// Change default emoji reaction to use in the quick reaction menu: the value is synced
// across devices and can be fetched using help.getConfig, reactions_default field¹.
//
// Links:
//  1. https://core.telegram.org/method/help.getConfig
//
// See https://core.telegram.org/method/messages.setDefaultReaction for reference.
type MessagesSetDefaultReactionRequest struct {
	// New emoji reaction
	Reaction ReactionClass
}

// MessagesSetDefaultReactionRequestTypeID is TL type id of MessagesSetDefaultReactionRequest.
const MessagesSetDefaultReactionRequestTypeID = 0x4f47a016

// Ensuring interfaces in compile-time for MessagesSetDefaultReactionRequest.
var (
	_ bin.Encoder     = &MessagesSetDefaultReactionRequest{}
	_ bin.Decoder     = &MessagesSetDefaultReactionRequest{}
	_ bin.BareEncoder = &MessagesSetDefaultReactionRequest{}
	_ bin.BareDecoder = &MessagesSetDefaultReactionRequest{}
)

func (s *MessagesSetDefaultReactionRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Reaction == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetDefaultReactionRequest) String() string {
	if s == nil {
		return "MessagesSetDefaultReactionRequest(nil)"
	}
	type Alias MessagesSetDefaultReactionRequest
	return fmt.Sprintf("MessagesSetDefaultReactionRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSetDefaultReactionRequest from given interface.
func (s *MessagesSetDefaultReactionRequest) FillFrom(from interface {
	GetReaction() (value ReactionClass)
}) {
	s.Reaction = from.GetReaction()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetDefaultReactionRequest) TypeID() uint32 {
	return MessagesSetDefaultReactionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetDefaultReactionRequest) TypeName() string {
	return "messages.setDefaultReaction"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetDefaultReactionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setDefaultReaction",
		ID:   MessagesSetDefaultReactionRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Reaction",
			SchemaName: "reaction",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSetDefaultReactionRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setDefaultReaction#4f47a016 as nil")
	}
	b.PutID(MessagesSetDefaultReactionRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSetDefaultReactionRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setDefaultReaction#4f47a016 as nil")
	}
	if s.Reaction == nil {
		return fmt.Errorf("unable to encode messages.setDefaultReaction#4f47a016: field reaction is nil")
	}
	if err := s.Reaction.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setDefaultReaction#4f47a016: field reaction: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSetDefaultReactionRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setDefaultReaction#4f47a016 to nil")
	}
	if err := b.ConsumeID(MessagesSetDefaultReactionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setDefaultReaction#4f47a016: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSetDefaultReactionRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setDefaultReaction#4f47a016 to nil")
	}
	{
		value, err := DecodeReaction(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setDefaultReaction#4f47a016: field reaction: %w", err)
		}
		s.Reaction = value
	}
	return nil
}

// GetReaction returns value of Reaction field.
func (s *MessagesSetDefaultReactionRequest) GetReaction() (value ReactionClass) {
	if s == nil {
		return
	}
	return s.Reaction
}

// MessagesSetDefaultReaction invokes method messages.setDefaultReaction#4f47a016 returning error if any.
// Change default emoji reaction to use in the quick reaction menu: the value is synced
// across devices and can be fetched using help.getConfig, reactions_default field¹.
//
// Links:
//  1. https://core.telegram.org/method/help.getConfig
//
// Possible errors:
//
//	400 REACTION_INVALID: The specified reaction is invalid.
//
// See https://core.telegram.org/method/messages.setDefaultReaction for reference.
func (c *Client) MessagesSetDefaultReaction(ctx context.Context, reaction ReactionClass) (bool, error) {
	var result BoolBox

	request := &MessagesSetDefaultReactionRequest{
		Reaction: reaction,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
