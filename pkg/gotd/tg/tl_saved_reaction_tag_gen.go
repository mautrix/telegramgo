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

// SavedReactionTag represents TL type `savedReactionTag#cb6ff828`.
// Info about a saved message reaction tag »¹.
//
// Links:
//  1. https://core.telegram.org/api/saved-messages#tags
//
// See https://core.telegram.org/constructor/savedReactionTag for reference.
type SavedReactionTag struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Reaction¹ associated to the tag.
	//
	// Links:
	//  1) https://core.telegram.org/api/reactions
	Reaction ReactionClass
	// Custom tag name assigned by the user (max 12 UTF-8 chars).
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// Number of messages tagged with this tag.
	Count int
}

// SavedReactionTagTypeID is TL type id of SavedReactionTag.
const SavedReactionTagTypeID = 0xcb6ff828

// Ensuring interfaces in compile-time for SavedReactionTag.
var (
	_ bin.Encoder     = &SavedReactionTag{}
	_ bin.Decoder     = &SavedReactionTag{}
	_ bin.BareEncoder = &SavedReactionTag{}
	_ bin.BareDecoder = &SavedReactionTag{}
)

func (s *SavedReactionTag) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Reaction == nil) {
		return false
	}
	if !(s.Title == "") {
		return false
	}
	if !(s.Count == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SavedReactionTag) String() string {
	if s == nil {
		return "SavedReactionTag(nil)"
	}
	type Alias SavedReactionTag
	return fmt.Sprintf("SavedReactionTag%+v", Alias(*s))
}

// FillFrom fills SavedReactionTag from given interface.
func (s *SavedReactionTag) FillFrom(from interface {
	GetReaction() (value ReactionClass)
	GetTitle() (value string, ok bool)
	GetCount() (value int)
}) {
	s.Reaction = from.GetReaction()
	if val, ok := from.GetTitle(); ok {
		s.Title = val
	}

	s.Count = from.GetCount()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SavedReactionTag) TypeID() uint32 {
	return SavedReactionTagTypeID
}

// TypeName returns name of type in TL schema.
func (*SavedReactionTag) TypeName() string {
	return "savedReactionTag"
}

// TypeInfo returns info about TL type.
func (s *SavedReactionTag) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "savedReactionTag",
		ID:   SavedReactionTagTypeID,
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
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *SavedReactionTag) SetFlags() {
	if !(s.Title == "") {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *SavedReactionTag) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode savedReactionTag#cb6ff828 as nil")
	}
	b.PutID(SavedReactionTagTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SavedReactionTag) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode savedReactionTag#cb6ff828 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode savedReactionTag#cb6ff828: field flags: %w", err)
	}
	if s.Reaction == nil {
		return fmt.Errorf("unable to encode savedReactionTag#cb6ff828: field reaction is nil")
	}
	if err := s.Reaction.Encode(b); err != nil {
		return fmt.Errorf("unable to encode savedReactionTag#cb6ff828: field reaction: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutString(s.Title)
	}
	b.PutInt(s.Count)
	return nil
}

// Decode implements bin.Decoder.
func (s *SavedReactionTag) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode savedReactionTag#cb6ff828 to nil")
	}
	if err := b.ConsumeID(SavedReactionTagTypeID); err != nil {
		return fmt.Errorf("unable to decode savedReactionTag#cb6ff828: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SavedReactionTag) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode savedReactionTag#cb6ff828 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode savedReactionTag#cb6ff828: field flags: %w", err)
		}
	}
	{
		value, err := DecodeReaction(b)
		if err != nil {
			return fmt.Errorf("unable to decode savedReactionTag#cb6ff828: field reaction: %w", err)
		}
		s.Reaction = value
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode savedReactionTag#cb6ff828: field title: %w", err)
		}
		s.Title = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode savedReactionTag#cb6ff828: field count: %w", err)
		}
		s.Count = value
	}
	return nil
}

// GetReaction returns value of Reaction field.
func (s *SavedReactionTag) GetReaction() (value ReactionClass) {
	if s == nil {
		return
	}
	return s.Reaction
}

// SetTitle sets value of Title conditional field.
func (s *SavedReactionTag) SetTitle(value string) {
	s.Flags.Set(0)
	s.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (s *SavedReactionTag) GetTitle() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.Title, true
}

// GetCount returns value of Count field.
func (s *SavedReactionTag) GetCount() (value int) {
	if s == nil {
		return
	}
	return s.Count
}
