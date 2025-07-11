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

// MessagesSaveRecentStickerRequest represents TL type `messages.saveRecentSticker#392718f8`.
// Add/remove sticker from recent stickers list
//
// See https://core.telegram.org/method/messages.saveRecentSticker for reference.
type MessagesSaveRecentStickerRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to add/remove stickers recently attached to photo or video files
	Attached bool
	// Sticker
	ID InputDocumentClass
	// Whether to save or unsave the sticker
	Unsave bool
}

// MessagesSaveRecentStickerRequestTypeID is TL type id of MessagesSaveRecentStickerRequest.
const MessagesSaveRecentStickerRequestTypeID = 0x392718f8

// Ensuring interfaces in compile-time for MessagesSaveRecentStickerRequest.
var (
	_ bin.Encoder     = &MessagesSaveRecentStickerRequest{}
	_ bin.Decoder     = &MessagesSaveRecentStickerRequest{}
	_ bin.BareEncoder = &MessagesSaveRecentStickerRequest{}
	_ bin.BareDecoder = &MessagesSaveRecentStickerRequest{}
)

func (s *MessagesSaveRecentStickerRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Attached == false) {
		return false
	}
	if !(s.ID == nil) {
		return false
	}
	if !(s.Unsave == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSaveRecentStickerRequest) String() string {
	if s == nil {
		return "MessagesSaveRecentStickerRequest(nil)"
	}
	type Alias MessagesSaveRecentStickerRequest
	return fmt.Sprintf("MessagesSaveRecentStickerRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSaveRecentStickerRequest from given interface.
func (s *MessagesSaveRecentStickerRequest) FillFrom(from interface {
	GetAttached() (value bool)
	GetID() (value InputDocumentClass)
	GetUnsave() (value bool)
}) {
	s.Attached = from.GetAttached()
	s.ID = from.GetID()
	s.Unsave = from.GetUnsave()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSaveRecentStickerRequest) TypeID() uint32 {
	return MessagesSaveRecentStickerRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSaveRecentStickerRequest) TypeName() string {
	return "messages.saveRecentSticker"
}

// TypeInfo returns info about TL type.
func (s *MessagesSaveRecentStickerRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.saveRecentSticker",
		ID:   MessagesSaveRecentStickerRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Attached",
			SchemaName: "attached",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Unsave",
			SchemaName: "unsave",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSaveRecentStickerRequest) SetFlags() {
	if !(s.Attached == false) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSaveRecentStickerRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.saveRecentSticker#392718f8 as nil")
	}
	b.PutID(MessagesSaveRecentStickerRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSaveRecentStickerRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.saveRecentSticker#392718f8 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.saveRecentSticker#392718f8: field flags: %w", err)
	}
	if s.ID == nil {
		return fmt.Errorf("unable to encode messages.saveRecentSticker#392718f8: field id is nil")
	}
	if err := s.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.saveRecentSticker#392718f8: field id: %w", err)
	}
	b.PutBool(s.Unsave)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSaveRecentStickerRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.saveRecentSticker#392718f8 to nil")
	}
	if err := b.ConsumeID(MessagesSaveRecentStickerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.saveRecentSticker#392718f8: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSaveRecentStickerRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.saveRecentSticker#392718f8 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.saveRecentSticker#392718f8: field flags: %w", err)
		}
	}
	s.Attached = s.Flags.Has(0)
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveRecentSticker#392718f8: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveRecentSticker#392718f8: field unsave: %w", err)
		}
		s.Unsave = value
	}
	return nil
}

// SetAttached sets value of Attached conditional field.
func (s *MessagesSaveRecentStickerRequest) SetAttached(value bool) {
	if value {
		s.Flags.Set(0)
		s.Attached = true
	} else {
		s.Flags.Unset(0)
		s.Attached = false
	}
}

// GetAttached returns value of Attached conditional field.
func (s *MessagesSaveRecentStickerRequest) GetAttached() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(0)
}

// GetID returns value of ID field.
func (s *MessagesSaveRecentStickerRequest) GetID() (value InputDocumentClass) {
	if s == nil {
		return
	}
	return s.ID
}

// GetUnsave returns value of Unsave field.
func (s *MessagesSaveRecentStickerRequest) GetUnsave() (value bool) {
	if s == nil {
		return
	}
	return s.Unsave
}

// GetIDAsNotEmpty returns mapped value of ID field.
func (s *MessagesSaveRecentStickerRequest) GetIDAsNotEmpty() (*InputDocument, bool) {
	return s.ID.AsNotEmpty()
}

// MessagesSaveRecentSticker invokes method messages.saveRecentSticker#392718f8 returning error if any.
// Add/remove sticker from recent stickers list
//
// Possible errors:
//
//	400 STICKER_ID_INVALID: The provided sticker ID is invalid.
//
// See https://core.telegram.org/method/messages.saveRecentSticker for reference.
func (c *Client) MessagesSaveRecentSticker(ctx context.Context, request *MessagesSaveRecentStickerRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
