// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// RemoveMessageReactionRequest represents TL type `removeMessageReaction#97474d7b`.
type RemoveMessageReactionRequest struct {
	// Identifier of the chat to which the message belongs
	ChatID int64
	// Identifier of the message
	MessageID int64
	// Type of the reaction to remove. The paid reaction can't be removed
	ReactionType ReactionTypeClass
}

// RemoveMessageReactionRequestTypeID is TL type id of RemoveMessageReactionRequest.
const RemoveMessageReactionRequestTypeID = 0x97474d7b

// Ensuring interfaces in compile-time for RemoveMessageReactionRequest.
var (
	_ bin.Encoder     = &RemoveMessageReactionRequest{}
	_ bin.Decoder     = &RemoveMessageReactionRequest{}
	_ bin.BareEncoder = &RemoveMessageReactionRequest{}
	_ bin.BareDecoder = &RemoveMessageReactionRequest{}
)

func (r *RemoveMessageReactionRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatID == 0) {
		return false
	}
	if !(r.MessageID == 0) {
		return false
	}
	if !(r.ReactionType == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RemoveMessageReactionRequest) String() string {
	if r == nil {
		return "RemoveMessageReactionRequest(nil)"
	}
	type Alias RemoveMessageReactionRequest
	return fmt.Sprintf("RemoveMessageReactionRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RemoveMessageReactionRequest) TypeID() uint32 {
	return RemoveMessageReactionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RemoveMessageReactionRequest) TypeName() string {
	return "removeMessageReaction"
}

// TypeInfo returns info about TL type.
func (r *RemoveMessageReactionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "removeMessageReaction",
		ID:   RemoveMessageReactionRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "ReactionType",
			SchemaName: "reaction_type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RemoveMessageReactionRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeMessageReaction#97474d7b as nil")
	}
	b.PutID(RemoveMessageReactionRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RemoveMessageReactionRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeMessageReaction#97474d7b as nil")
	}
	b.PutInt53(r.ChatID)
	b.PutInt53(r.MessageID)
	if r.ReactionType == nil {
		return fmt.Errorf("unable to encode removeMessageReaction#97474d7b: field reaction_type is nil")
	}
	if err := r.ReactionType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode removeMessageReaction#97474d7b: field reaction_type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RemoveMessageReactionRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeMessageReaction#97474d7b to nil")
	}
	if err := b.ConsumeID(RemoveMessageReactionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode removeMessageReaction#97474d7b: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RemoveMessageReactionRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeMessageReaction#97474d7b to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode removeMessageReaction#97474d7b: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode removeMessageReaction#97474d7b: field message_id: %w", err)
		}
		r.MessageID = value
	}
	{
		value, err := DecodeReactionType(b)
		if err != nil {
			return fmt.Errorf("unable to decode removeMessageReaction#97474d7b: field reaction_type: %w", err)
		}
		r.ReactionType = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RemoveMessageReactionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode removeMessageReaction#97474d7b as nil")
	}
	b.ObjStart()
	b.PutID("removeMessageReaction")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(r.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(r.MessageID)
	b.Comma()
	b.FieldStart("reaction_type")
	if r.ReactionType == nil {
		return fmt.Errorf("unable to encode removeMessageReaction#97474d7b: field reaction_type is nil")
	}
	if err := r.ReactionType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode removeMessageReaction#97474d7b: field reaction_type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RemoveMessageReactionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode removeMessageReaction#97474d7b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("removeMessageReaction"); err != nil {
				return fmt.Errorf("unable to decode removeMessageReaction#97474d7b: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode removeMessageReaction#97474d7b: field chat_id: %w", err)
			}
			r.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode removeMessageReaction#97474d7b: field message_id: %w", err)
			}
			r.MessageID = value
		case "reaction_type":
			value, err := DecodeTDLibJSONReactionType(b)
			if err != nil {
				return fmt.Errorf("unable to decode removeMessageReaction#97474d7b: field reaction_type: %w", err)
			}
			r.ReactionType = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (r *RemoveMessageReactionRequest) GetChatID() (value int64) {
	if r == nil {
		return
	}
	return r.ChatID
}

// GetMessageID returns value of MessageID field.
func (r *RemoveMessageReactionRequest) GetMessageID() (value int64) {
	if r == nil {
		return
	}
	return r.MessageID
}

// GetReactionType returns value of ReactionType field.
func (r *RemoveMessageReactionRequest) GetReactionType() (value ReactionTypeClass) {
	if r == nil {
		return
	}
	return r.ReactionType
}

// RemoveMessageReaction invokes method removeMessageReaction#97474d7b returning error if any.
func (c *Client) RemoveMessageReaction(ctx context.Context, request *RemoveMessageReactionRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
