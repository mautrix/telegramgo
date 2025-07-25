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

// Sticker represents TL type `sticker#d96f5d3f`.
type Sticker struct {
	// Unique sticker identifier within the set; 0 if none
	ID int64
	// Identifier of the sticker set to which the sticker belongs; 0 if none
	SetID int64
	// Sticker width; as defined by the sender
	Width int32
	// Sticker height; as defined by the sender
	Height int32
	// Emoji corresponding to the sticker
	Emoji string
	// Sticker format
	Format StickerFormatClass
	// Sticker's full type
	FullType StickerFullTypeClass
	// Sticker thumbnail in WEBP or JPEG format; may be null
	Thumbnail Thumbnail
	// File containing the sticker
	Sticker File
}

// StickerTypeID is TL type id of Sticker.
const StickerTypeID = 0xd96f5d3f

// Ensuring interfaces in compile-time for Sticker.
var (
	_ bin.Encoder     = &Sticker{}
	_ bin.Decoder     = &Sticker{}
	_ bin.BareEncoder = &Sticker{}
	_ bin.BareDecoder = &Sticker{}
)

func (s *Sticker) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == 0) {
		return false
	}
	if !(s.SetID == 0) {
		return false
	}
	if !(s.Width == 0) {
		return false
	}
	if !(s.Height == 0) {
		return false
	}
	if !(s.Emoji == "") {
		return false
	}
	if !(s.Format == nil) {
		return false
	}
	if !(s.FullType == nil) {
		return false
	}
	if !(s.Thumbnail.Zero()) {
		return false
	}
	if !(s.Sticker.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *Sticker) String() string {
	if s == nil {
		return "Sticker(nil)"
	}
	type Alias Sticker
	return fmt.Sprintf("Sticker%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Sticker) TypeID() uint32 {
	return StickerTypeID
}

// TypeName returns name of type in TL schema.
func (*Sticker) TypeName() string {
	return "sticker"
}

// TypeInfo returns info about TL type.
func (s *Sticker) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sticker",
		ID:   StickerTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "SetID",
			SchemaName: "set_id",
		},
		{
			Name:       "Width",
			SchemaName: "width",
		},
		{
			Name:       "Height",
			SchemaName: "height",
		},
		{
			Name:       "Emoji",
			SchemaName: "emoji",
		},
		{
			Name:       "Format",
			SchemaName: "format",
		},
		{
			Name:       "FullType",
			SchemaName: "full_type",
		},
		{
			Name:       "Thumbnail",
			SchemaName: "thumbnail",
		},
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *Sticker) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sticker#d96f5d3f as nil")
	}
	b.PutID(StickerTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *Sticker) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sticker#d96f5d3f as nil")
	}
	b.PutLong(s.ID)
	b.PutLong(s.SetID)
	b.PutInt32(s.Width)
	b.PutInt32(s.Height)
	b.PutString(s.Emoji)
	if s.Format == nil {
		return fmt.Errorf("unable to encode sticker#d96f5d3f: field format is nil")
	}
	if err := s.Format.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sticker#d96f5d3f: field format: %w", err)
	}
	if s.FullType == nil {
		return fmt.Errorf("unable to encode sticker#d96f5d3f: field full_type is nil")
	}
	if err := s.FullType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sticker#d96f5d3f: field full_type: %w", err)
	}
	if err := s.Thumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sticker#d96f5d3f: field thumbnail: %w", err)
	}
	if err := s.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sticker#d96f5d3f: field sticker: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *Sticker) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sticker#d96f5d3f to nil")
	}
	if err := b.ConsumeID(StickerTypeID); err != nil {
		return fmt.Errorf("unable to decode sticker#d96f5d3f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *Sticker) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sticker#d96f5d3f to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode sticker#d96f5d3f: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode sticker#d96f5d3f: field set_id: %w", err)
		}
		s.SetID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode sticker#d96f5d3f: field width: %w", err)
		}
		s.Width = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode sticker#d96f5d3f: field height: %w", err)
		}
		s.Height = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sticker#d96f5d3f: field emoji: %w", err)
		}
		s.Emoji = value
	}
	{
		value, err := DecodeStickerFormat(b)
		if err != nil {
			return fmt.Errorf("unable to decode sticker#d96f5d3f: field format: %w", err)
		}
		s.Format = value
	}
	{
		value, err := DecodeStickerFullType(b)
		if err != nil {
			return fmt.Errorf("unable to decode sticker#d96f5d3f: field full_type: %w", err)
		}
		s.FullType = value
	}
	{
		if err := s.Thumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sticker#d96f5d3f: field thumbnail: %w", err)
		}
	}
	{
		if err := s.Sticker.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sticker#d96f5d3f: field sticker: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *Sticker) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sticker#d96f5d3f as nil")
	}
	b.ObjStart()
	b.PutID("sticker")
	b.Comma()
	b.FieldStart("id")
	b.PutLong(s.ID)
	b.Comma()
	b.FieldStart("set_id")
	b.PutLong(s.SetID)
	b.Comma()
	b.FieldStart("width")
	b.PutInt32(s.Width)
	b.Comma()
	b.FieldStart("height")
	b.PutInt32(s.Height)
	b.Comma()
	b.FieldStart("emoji")
	b.PutString(s.Emoji)
	b.Comma()
	b.FieldStart("format")
	if s.Format == nil {
		return fmt.Errorf("unable to encode sticker#d96f5d3f: field format is nil")
	}
	if err := s.Format.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sticker#d96f5d3f: field format: %w", err)
	}
	b.Comma()
	b.FieldStart("full_type")
	if s.FullType == nil {
		return fmt.Errorf("unable to encode sticker#d96f5d3f: field full_type is nil")
	}
	if err := s.FullType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sticker#d96f5d3f: field full_type: %w", err)
	}
	b.Comma()
	b.FieldStart("thumbnail")
	if err := s.Thumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sticker#d96f5d3f: field thumbnail: %w", err)
	}
	b.Comma()
	b.FieldStart("sticker")
	if err := s.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sticker#d96f5d3f: field sticker: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *Sticker) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sticker#d96f5d3f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sticker"); err != nil {
				return fmt.Errorf("unable to decode sticker#d96f5d3f: %w", err)
			}
		case "id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode sticker#d96f5d3f: field id: %w", err)
			}
			s.ID = value
		case "set_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode sticker#d96f5d3f: field set_id: %w", err)
			}
			s.SetID = value
		case "width":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode sticker#d96f5d3f: field width: %w", err)
			}
			s.Width = value
		case "height":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode sticker#d96f5d3f: field height: %w", err)
			}
			s.Height = value
		case "emoji":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sticker#d96f5d3f: field emoji: %w", err)
			}
			s.Emoji = value
		case "format":
			value, err := DecodeTDLibJSONStickerFormat(b)
			if err != nil {
				return fmt.Errorf("unable to decode sticker#d96f5d3f: field format: %w", err)
			}
			s.Format = value
		case "full_type":
			value, err := DecodeTDLibJSONStickerFullType(b)
			if err != nil {
				return fmt.Errorf("unable to decode sticker#d96f5d3f: field full_type: %w", err)
			}
			s.FullType = value
		case "thumbnail":
			if err := s.Thumbnail.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sticker#d96f5d3f: field thumbnail: %w", err)
			}
		case "sticker":
			if err := s.Sticker.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sticker#d96f5d3f: field sticker: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (s *Sticker) GetID() (value int64) {
	if s == nil {
		return
	}
	return s.ID
}

// GetSetID returns value of SetID field.
func (s *Sticker) GetSetID() (value int64) {
	if s == nil {
		return
	}
	return s.SetID
}

// GetWidth returns value of Width field.
func (s *Sticker) GetWidth() (value int32) {
	if s == nil {
		return
	}
	return s.Width
}

// GetHeight returns value of Height field.
func (s *Sticker) GetHeight() (value int32) {
	if s == nil {
		return
	}
	return s.Height
}

// GetEmoji returns value of Emoji field.
func (s *Sticker) GetEmoji() (value string) {
	if s == nil {
		return
	}
	return s.Emoji
}

// GetFormat returns value of Format field.
func (s *Sticker) GetFormat() (value StickerFormatClass) {
	if s == nil {
		return
	}
	return s.Format
}

// GetFullType returns value of FullType field.
func (s *Sticker) GetFullType() (value StickerFullTypeClass) {
	if s == nil {
		return
	}
	return s.FullType
}

// GetThumbnail returns value of Thumbnail field.
func (s *Sticker) GetThumbnail() (value Thumbnail) {
	if s == nil {
		return
	}
	return s.Thumbnail
}

// GetSticker returns value of Sticker field.
func (s *Sticker) GetSticker() (value File) {
	if s == nil {
		return
	}
	return s.Sticker
}
