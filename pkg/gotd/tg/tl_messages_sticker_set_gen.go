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

// MessagesStickerSet represents TL type `messages.stickerSet#6e153f16`.
// Stickerset and stickers inside it
//
// See https://core.telegram.org/constructor/messages.stickerSet for reference.
type MessagesStickerSet struct {
	// The stickerset
	Set StickerSet
	// Emoji info for stickers
	Packs []StickerPack
	// Keywords for some or every sticker in the stickerset.
	Keywords []StickerKeyword
	// Stickers in stickerset
	Documents []DocumentClass
}

// MessagesStickerSetTypeID is TL type id of MessagesStickerSet.
const MessagesStickerSetTypeID = 0x6e153f16

// construct implements constructor of MessagesStickerSetClass.
func (s MessagesStickerSet) construct() MessagesStickerSetClass { return &s }

// Ensuring interfaces in compile-time for MessagesStickerSet.
var (
	_ bin.Encoder     = &MessagesStickerSet{}
	_ bin.Decoder     = &MessagesStickerSet{}
	_ bin.BareEncoder = &MessagesStickerSet{}
	_ bin.BareDecoder = &MessagesStickerSet{}

	_ MessagesStickerSetClass = &MessagesStickerSet{}
)

func (s *MessagesStickerSet) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Set.Zero()) {
		return false
	}
	if !(s.Packs == nil) {
		return false
	}
	if !(s.Keywords == nil) {
		return false
	}
	if !(s.Documents == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesStickerSet) String() string {
	if s == nil {
		return "MessagesStickerSet(nil)"
	}
	type Alias MessagesStickerSet
	return fmt.Sprintf("MessagesStickerSet%+v", Alias(*s))
}

// FillFrom fills MessagesStickerSet from given interface.
func (s *MessagesStickerSet) FillFrom(from interface {
	GetSet() (value StickerSet)
	GetPacks() (value []StickerPack)
	GetKeywords() (value []StickerKeyword)
	GetDocuments() (value []DocumentClass)
}) {
	s.Set = from.GetSet()
	s.Packs = from.GetPacks()
	s.Keywords = from.GetKeywords()
	s.Documents = from.GetDocuments()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesStickerSet) TypeID() uint32 {
	return MessagesStickerSetTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesStickerSet) TypeName() string {
	return "messages.stickerSet"
}

// TypeInfo returns info about TL type.
func (s *MessagesStickerSet) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.stickerSet",
		ID:   MessagesStickerSetTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Set",
			SchemaName: "set",
		},
		{
			Name:       "Packs",
			SchemaName: "packs",
		},
		{
			Name:       "Keywords",
			SchemaName: "keywords",
		},
		{
			Name:       "Documents",
			SchemaName: "documents",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesStickerSet) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.stickerSet#6e153f16 as nil")
	}
	b.PutID(MessagesStickerSetTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesStickerSet) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.stickerSet#6e153f16 as nil")
	}
	if err := s.Set.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.stickerSet#6e153f16: field set: %w", err)
	}
	b.PutVectorHeader(len(s.Packs))
	for idx, v := range s.Packs {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.stickerSet#6e153f16: field packs element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(s.Keywords))
	for idx, v := range s.Keywords {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.stickerSet#6e153f16: field keywords element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(s.Documents))
	for idx, v := range s.Documents {
		if v == nil {
			return fmt.Errorf("unable to encode messages.stickerSet#6e153f16: field documents element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.stickerSet#6e153f16: field documents element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesStickerSet) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.stickerSet#6e153f16 to nil")
	}
	if err := b.ConsumeID(MessagesStickerSetTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.stickerSet#6e153f16: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesStickerSet) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.stickerSet#6e153f16 to nil")
	}
	{
		if err := s.Set.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.stickerSet#6e153f16: field set: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.stickerSet#6e153f16: field packs: %w", err)
		}

		if headerLen > 0 {
			s.Packs = make([]StickerPack, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StickerPack
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.stickerSet#6e153f16: field packs: %w", err)
			}
			s.Packs = append(s.Packs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.stickerSet#6e153f16: field keywords: %w", err)
		}

		if headerLen > 0 {
			s.Keywords = make([]StickerKeyword, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StickerKeyword
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.stickerSet#6e153f16: field keywords: %w", err)
			}
			s.Keywords = append(s.Keywords, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.stickerSet#6e153f16: field documents: %w", err)
		}

		if headerLen > 0 {
			s.Documents = make([]DocumentClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.stickerSet#6e153f16: field documents: %w", err)
			}
			s.Documents = append(s.Documents, value)
		}
	}
	return nil
}

// GetSet returns value of Set field.
func (s *MessagesStickerSet) GetSet() (value StickerSet) {
	if s == nil {
		return
	}
	return s.Set
}

// GetPacks returns value of Packs field.
func (s *MessagesStickerSet) GetPacks() (value []StickerPack) {
	if s == nil {
		return
	}
	return s.Packs
}

// GetKeywords returns value of Keywords field.
func (s *MessagesStickerSet) GetKeywords() (value []StickerKeyword) {
	if s == nil {
		return
	}
	return s.Keywords
}

// GetDocuments returns value of Documents field.
func (s *MessagesStickerSet) GetDocuments() (value []DocumentClass) {
	if s == nil {
		return
	}
	return s.Documents
}

// MapDocuments returns field Documents wrapped in DocumentClassArray helper.
func (s *MessagesStickerSet) MapDocuments() (value DocumentClassArray) {
	return DocumentClassArray(s.Documents)
}

// MessagesStickerSetNotModified represents TL type `messages.stickerSetNotModified#d3f924eb`.
// The stickerset hasn't changed
//
// See https://core.telegram.org/constructor/messages.stickerSetNotModified for reference.
type MessagesStickerSetNotModified struct {
}

// MessagesStickerSetNotModifiedTypeID is TL type id of MessagesStickerSetNotModified.
const MessagesStickerSetNotModifiedTypeID = 0xd3f924eb

// construct implements constructor of MessagesStickerSetClass.
func (s MessagesStickerSetNotModified) construct() MessagesStickerSetClass { return &s }

// Ensuring interfaces in compile-time for MessagesStickerSetNotModified.
var (
	_ bin.Encoder     = &MessagesStickerSetNotModified{}
	_ bin.Decoder     = &MessagesStickerSetNotModified{}
	_ bin.BareEncoder = &MessagesStickerSetNotModified{}
	_ bin.BareDecoder = &MessagesStickerSetNotModified{}

	_ MessagesStickerSetClass = &MessagesStickerSetNotModified{}
)

func (s *MessagesStickerSetNotModified) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesStickerSetNotModified) String() string {
	if s == nil {
		return "MessagesStickerSetNotModified(nil)"
	}
	type Alias MessagesStickerSetNotModified
	return fmt.Sprintf("MessagesStickerSetNotModified%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesStickerSetNotModified) TypeID() uint32 {
	return MessagesStickerSetNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesStickerSetNotModified) TypeName() string {
	return "messages.stickerSetNotModified"
}

// TypeInfo returns info about TL type.
func (s *MessagesStickerSetNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.stickerSetNotModified",
		ID:   MessagesStickerSetNotModifiedTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesStickerSetNotModified) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.stickerSetNotModified#d3f924eb as nil")
	}
	b.PutID(MessagesStickerSetNotModifiedTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesStickerSetNotModified) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.stickerSetNotModified#d3f924eb as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesStickerSetNotModified) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.stickerSetNotModified#d3f924eb to nil")
	}
	if err := b.ConsumeID(MessagesStickerSetNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.stickerSetNotModified#d3f924eb: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesStickerSetNotModified) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.stickerSetNotModified#d3f924eb to nil")
	}
	return nil
}

// MessagesStickerSetClassName is schema name of MessagesStickerSetClass.
const MessagesStickerSetClassName = "messages.StickerSet"

// MessagesStickerSetClass represents messages.StickerSet generic type.
//
// See https://core.telegram.org/type/messages.StickerSet for reference.
//
// Example:
//
//	g, err := tg.DecodeMessagesStickerSet(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.MessagesStickerSet: // messages.stickerSet#6e153f16
//	case *tg.MessagesStickerSetNotModified: // messages.stickerSetNotModified#d3f924eb
//	default: panic(v)
//	}
type MessagesStickerSetClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesStickerSetClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsModified tries to map MessagesStickerSetClass to MessagesStickerSet.
	AsModified() (*MessagesStickerSet, bool)
}

// AsModified tries to map MessagesStickerSet to MessagesStickerSet.
func (s *MessagesStickerSet) AsModified() (*MessagesStickerSet, bool) {
	return s, true
}

// AsModified tries to map MessagesStickerSetNotModified to MessagesStickerSet.
func (s *MessagesStickerSetNotModified) AsModified() (*MessagesStickerSet, bool) {
	return nil, false
}

// DecodeMessagesStickerSet implements binary de-serialization for MessagesStickerSetClass.
func DecodeMessagesStickerSet(buf *bin.Buffer) (MessagesStickerSetClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesStickerSetTypeID:
		// Decoding messages.stickerSet#6e153f16.
		v := MessagesStickerSet{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesStickerSetClass: %w", err)
		}
		return &v, nil
	case MessagesStickerSetNotModifiedTypeID:
		// Decoding messages.stickerSetNotModified#d3f924eb.
		v := MessagesStickerSetNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesStickerSetClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesStickerSetClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesStickerSet boxes the MessagesStickerSetClass providing a helper.
type MessagesStickerSetBox struct {
	StickerSet MessagesStickerSetClass
}

// Decode implements bin.Decoder for MessagesStickerSetBox.
func (b *MessagesStickerSetBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesStickerSetBox to nil")
	}
	v, err := DecodeMessagesStickerSet(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StickerSet = v
	return nil
}

// Encode implements bin.Encode for MessagesStickerSetBox.
func (b *MessagesStickerSetBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StickerSet == nil {
		return fmt.Errorf("unable to encode MessagesStickerSetClass as nil")
	}
	return b.StickerSet.Encode(buf)
}
