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

// SetChatAvailableReactionsRequest represents TL type `setChatAvailableReactions#feb3e06`.
type SetChatAvailableReactionsRequest struct {
	// Identifier of the chat
	ChatID int64
	// Reactions available in the chat. All explicitly specified emoji reactions must be
	// active. In channel chats up to the chat's boost level custom emoji reactions can be
	// explicitly specified
	AvailableReactions ChatAvailableReactionsClass
}

// SetChatAvailableReactionsRequestTypeID is TL type id of SetChatAvailableReactionsRequest.
const SetChatAvailableReactionsRequestTypeID = 0xfeb3e06

// Ensuring interfaces in compile-time for SetChatAvailableReactionsRequest.
var (
	_ bin.Encoder     = &SetChatAvailableReactionsRequest{}
	_ bin.Decoder     = &SetChatAvailableReactionsRequest{}
	_ bin.BareEncoder = &SetChatAvailableReactionsRequest{}
	_ bin.BareDecoder = &SetChatAvailableReactionsRequest{}
)

func (s *SetChatAvailableReactionsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.AvailableReactions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetChatAvailableReactionsRequest) String() string {
	if s == nil {
		return "SetChatAvailableReactionsRequest(nil)"
	}
	type Alias SetChatAvailableReactionsRequest
	return fmt.Sprintf("SetChatAvailableReactionsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetChatAvailableReactionsRequest) TypeID() uint32 {
	return SetChatAvailableReactionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetChatAvailableReactionsRequest) TypeName() string {
	return "setChatAvailableReactions"
}

// TypeInfo returns info about TL type.
func (s *SetChatAvailableReactionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setChatAvailableReactions",
		ID:   SetChatAvailableReactionsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "AvailableReactions",
			SchemaName: "available_reactions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetChatAvailableReactionsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatAvailableReactions#feb3e06 as nil")
	}
	b.PutID(SetChatAvailableReactionsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetChatAvailableReactionsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatAvailableReactions#feb3e06 as nil")
	}
	b.PutInt53(s.ChatID)
	if s.AvailableReactions == nil {
		return fmt.Errorf("unable to encode setChatAvailableReactions#feb3e06: field available_reactions is nil")
	}
	if err := s.AvailableReactions.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setChatAvailableReactions#feb3e06: field available_reactions: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetChatAvailableReactionsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatAvailableReactions#feb3e06 to nil")
	}
	if err := b.ConsumeID(SetChatAvailableReactionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setChatAvailableReactions#feb3e06: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetChatAvailableReactionsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatAvailableReactions#feb3e06 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setChatAvailableReactions#feb3e06: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := DecodeChatAvailableReactions(b)
		if err != nil {
			return fmt.Errorf("unable to decode setChatAvailableReactions#feb3e06: field available_reactions: %w", err)
		}
		s.AvailableReactions = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetChatAvailableReactionsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatAvailableReactions#feb3e06 as nil")
	}
	b.ObjStart()
	b.PutID("setChatAvailableReactions")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("available_reactions")
	if s.AvailableReactions == nil {
		return fmt.Errorf("unable to encode setChatAvailableReactions#feb3e06: field available_reactions is nil")
	}
	if err := s.AvailableReactions.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setChatAvailableReactions#feb3e06: field available_reactions: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetChatAvailableReactionsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatAvailableReactions#feb3e06 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setChatAvailableReactions"); err != nil {
				return fmt.Errorf("unable to decode setChatAvailableReactions#feb3e06: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setChatAvailableReactions#feb3e06: field chat_id: %w", err)
			}
			s.ChatID = value
		case "available_reactions":
			value, err := DecodeTDLibJSONChatAvailableReactions(b)
			if err != nil {
				return fmt.Errorf("unable to decode setChatAvailableReactions#feb3e06: field available_reactions: %w", err)
			}
			s.AvailableReactions = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetChatAvailableReactionsRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetAvailableReactions returns value of AvailableReactions field.
func (s *SetChatAvailableReactionsRequest) GetAvailableReactions() (value ChatAvailableReactionsClass) {
	if s == nil {
		return
	}
	return s.AvailableReactions
}

// SetChatAvailableReactions invokes method setChatAvailableReactions#feb3e06 returning error if any.
func (c *Client) SetChatAvailableReactions(ctx context.Context, request *SetChatAvailableReactionsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
