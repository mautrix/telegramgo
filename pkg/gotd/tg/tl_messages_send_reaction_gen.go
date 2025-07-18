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

// MessagesSendReactionRequest represents TL type `messages.sendReaction#d30d78d4`.
// React to message.
// Starting from layer 159, the reaction will be sent from the peer specified using
// messages.saveDefaultSendAs¹.
//
// Links:
//  1. https://core.telegram.org/method/messages.saveDefaultSendAs
//
// See https://core.telegram.org/method/messages.sendReaction for reference.
type MessagesSendReactionRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether a bigger and longer reaction should be shown
	Big bool
	// Whether to add this reaction to the recent reactions list »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/reactions#recent-reactions
	AddToRecent bool
	// Peer
	Peer InputPeerClass
	// Message ID to react to
	MsgID int
	// A list of reactions
	//
	// Use SetReaction and GetReaction helpers.
	Reaction []ReactionClass
}

// MessagesSendReactionRequestTypeID is TL type id of MessagesSendReactionRequest.
const MessagesSendReactionRequestTypeID = 0xd30d78d4

// Ensuring interfaces in compile-time for MessagesSendReactionRequest.
var (
	_ bin.Encoder     = &MessagesSendReactionRequest{}
	_ bin.Decoder     = &MessagesSendReactionRequest{}
	_ bin.BareEncoder = &MessagesSendReactionRequest{}
	_ bin.BareDecoder = &MessagesSendReactionRequest{}
)

func (s *MessagesSendReactionRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Big == false) {
		return false
	}
	if !(s.AddToRecent == false) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.MsgID == 0) {
		return false
	}
	if !(s.Reaction == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendReactionRequest) String() string {
	if s == nil {
		return "MessagesSendReactionRequest(nil)"
	}
	type Alias MessagesSendReactionRequest
	return fmt.Sprintf("MessagesSendReactionRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSendReactionRequest from given interface.
func (s *MessagesSendReactionRequest) FillFrom(from interface {
	GetBig() (value bool)
	GetAddToRecent() (value bool)
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetReaction() (value []ReactionClass, ok bool)
}) {
	s.Big = from.GetBig()
	s.AddToRecent = from.GetAddToRecent()
	s.Peer = from.GetPeer()
	s.MsgID = from.GetMsgID()
	if val, ok := from.GetReaction(); ok {
		s.Reaction = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSendReactionRequest) TypeID() uint32 {
	return MessagesSendReactionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSendReactionRequest) TypeName() string {
	return "messages.sendReaction"
}

// TypeInfo returns info about TL type.
func (s *MessagesSendReactionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sendReaction",
		ID:   MessagesSendReactionRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Big",
			SchemaName: "big",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "AddToRecent",
			SchemaName: "add_to_recent",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "Reaction",
			SchemaName: "reaction",
			Null:       !s.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSendReactionRequest) SetFlags() {
	if !(s.Big == false) {
		s.Flags.Set(1)
	}
	if !(s.AddToRecent == false) {
		s.Flags.Set(2)
	}
	if !(s.Reaction == nil) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSendReactionRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendReaction#d30d78d4 as nil")
	}
	b.PutID(MessagesSendReactionRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSendReactionRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendReaction#d30d78d4 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendReaction#d30d78d4: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendReaction#d30d78d4: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendReaction#d30d78d4: field peer: %w", err)
	}
	b.PutInt(s.MsgID)
	if s.Flags.Has(0) {
		b.PutVectorHeader(len(s.Reaction))
		for idx, v := range s.Reaction {
			if v == nil {
				return fmt.Errorf("unable to encode messages.sendReaction#d30d78d4: field reaction element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.sendReaction#d30d78d4: field reaction element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendReactionRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendReaction#d30d78d4 to nil")
	}
	if err := b.ConsumeID(MessagesSendReactionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendReaction#d30d78d4: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSendReactionRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendReaction#d30d78d4 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendReaction#d30d78d4: field flags: %w", err)
		}
	}
	s.Big = s.Flags.Has(1)
	s.AddToRecent = s.Flags.Has(2)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendReaction#d30d78d4: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendReaction#d30d78d4: field msg_id: %w", err)
		}
		s.MsgID = value
	}
	if s.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendReaction#d30d78d4: field reaction: %w", err)
		}

		if headerLen > 0 {
			s.Reaction = make([]ReactionClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeReaction(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.sendReaction#d30d78d4: field reaction: %w", err)
			}
			s.Reaction = append(s.Reaction, value)
		}
	}
	return nil
}

// SetBig sets value of Big conditional field.
func (s *MessagesSendReactionRequest) SetBig(value bool) {
	if value {
		s.Flags.Set(1)
		s.Big = true
	} else {
		s.Flags.Unset(1)
		s.Big = false
	}
}

// GetBig returns value of Big conditional field.
func (s *MessagesSendReactionRequest) GetBig() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(1)
}

// SetAddToRecent sets value of AddToRecent conditional field.
func (s *MessagesSendReactionRequest) SetAddToRecent(value bool) {
	if value {
		s.Flags.Set(2)
		s.AddToRecent = true
	} else {
		s.Flags.Unset(2)
		s.AddToRecent = false
	}
}

// GetAddToRecent returns value of AddToRecent conditional field.
func (s *MessagesSendReactionRequest) GetAddToRecent() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(2)
}

// GetPeer returns value of Peer field.
func (s *MessagesSendReactionRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetMsgID returns value of MsgID field.
func (s *MessagesSendReactionRequest) GetMsgID() (value int) {
	if s == nil {
		return
	}
	return s.MsgID
}

// SetReaction sets value of Reaction conditional field.
func (s *MessagesSendReactionRequest) SetReaction(value []ReactionClass) {
	s.Flags.Set(0)
	s.Reaction = value
}

// GetReaction returns value of Reaction conditional field and
// boolean which is true if field was set.
func (s *MessagesSendReactionRequest) GetReaction() (value []ReactionClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.Reaction, true
}

// MapReaction returns field Reaction wrapped in ReactionClassArray helper.
func (s *MessagesSendReactionRequest) MapReaction() (value ReactionClassArray, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return ReactionClassArray(s.Reaction), true
}

// MessagesSendReaction invokes method messages.sendReaction#d30d78d4 returning error if any.
// React to message.
// Starting from layer 159, the reaction will be sent from the peer specified using
// messages.saveDefaultSendAs¹.
//
// Links:
//  1. https://core.telegram.org/method/messages.saveDefaultSendAs
//
// Possible errors:
//
//	403 ANONYMOUS_REACTIONS_DISABLED: Sorry, anonymous administrators cannot leave reactions or participate in polls.
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	400 CUSTOM_REACTIONS_TOO_MANY: Too many custom reactions were specified.
//	400 DOCUMENT_INVALID: The specified document is invalid.
//	400 MESSAGE_ID_INVALID: The provided message id is invalid.
//	400 MESSAGE_NOT_MODIFIED: The provided message data is identical to the previous message data, the message wasn't modified.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	403 PREMIUM_ACCOUNT_REQUIRED: A premium account is required to execute this action.
//	400 REACTIONS_TOO_MANY: The message already has exactly reactions_uniq_max reaction emojis, you can't react with a new emoji, see the docs for more info ».
//	400 REACTION_EMPTY: Empty reaction provided.
//	400 REACTION_INVALID: The specified reaction is invalid.
//	400 USER_BANNED_IN_CHANNEL: You're banned from sending messages in supergroups/channels.
//
// See https://core.telegram.org/method/messages.sendReaction for reference.
func (c *Client) MessagesSendReaction(ctx context.Context, request *MessagesSendReactionRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
