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

// StoryContentPhoto represents TL type `storyContentPhoto#d45f0050`.
type StoryContentPhoto struct {
	// The photo
	Photo Photo
}

// StoryContentPhotoTypeID is TL type id of StoryContentPhoto.
const StoryContentPhotoTypeID = 0xd45f0050

// construct implements constructor of StoryContentClass.
func (s StoryContentPhoto) construct() StoryContentClass { return &s }

// Ensuring interfaces in compile-time for StoryContentPhoto.
var (
	_ bin.Encoder     = &StoryContentPhoto{}
	_ bin.Decoder     = &StoryContentPhoto{}
	_ bin.BareEncoder = &StoryContentPhoto{}
	_ bin.BareDecoder = &StoryContentPhoto{}

	_ StoryContentClass = &StoryContentPhoto{}
)

func (s *StoryContentPhoto) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Photo.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryContentPhoto) String() string {
	if s == nil {
		return "StoryContentPhoto(nil)"
	}
	type Alias StoryContentPhoto
	return fmt.Sprintf("StoryContentPhoto%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryContentPhoto) TypeID() uint32 {
	return StoryContentPhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryContentPhoto) TypeName() string {
	return "storyContentPhoto"
}

// TypeInfo returns info about TL type.
func (s *StoryContentPhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyContentPhoto",
		ID:   StoryContentPhotoTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryContentPhoto) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyContentPhoto#d45f0050 as nil")
	}
	b.PutID(StoryContentPhotoTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryContentPhoto) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyContentPhoto#d45f0050 as nil")
	}
	if err := s.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyContentPhoto#d45f0050: field photo: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryContentPhoto) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyContentPhoto#d45f0050 to nil")
	}
	if err := b.ConsumeID(StoryContentPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode storyContentPhoto#d45f0050: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryContentPhoto) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyContentPhoto#d45f0050 to nil")
	}
	{
		if err := s.Photo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyContentPhoto#d45f0050: field photo: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryContentPhoto) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyContentPhoto#d45f0050 as nil")
	}
	b.ObjStart()
	b.PutID("storyContentPhoto")
	b.Comma()
	b.FieldStart("photo")
	if err := s.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode storyContentPhoto#d45f0050: field photo: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryContentPhoto) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyContentPhoto#d45f0050 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyContentPhoto"); err != nil {
				return fmt.Errorf("unable to decode storyContentPhoto#d45f0050: %w", err)
			}
		case "photo":
			if err := s.Photo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode storyContentPhoto#d45f0050: field photo: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPhoto returns value of Photo field.
func (s *StoryContentPhoto) GetPhoto() (value Photo) {
	if s == nil {
		return
	}
	return s.Photo
}

// StoryContentVideo represents TL type `storyContentVideo#b30162a6`.
type StoryContentVideo struct {
	// The video in MPEG4 format
	Video StoryVideo
	// Alternative version of the video in MPEG4 format, encoded with H.264 codec; may be
	// null
	AlternativeVideo StoryVideo
}

// StoryContentVideoTypeID is TL type id of StoryContentVideo.
const StoryContentVideoTypeID = 0xb30162a6

// construct implements constructor of StoryContentClass.
func (s StoryContentVideo) construct() StoryContentClass { return &s }

// Ensuring interfaces in compile-time for StoryContentVideo.
var (
	_ bin.Encoder     = &StoryContentVideo{}
	_ bin.Decoder     = &StoryContentVideo{}
	_ bin.BareEncoder = &StoryContentVideo{}
	_ bin.BareDecoder = &StoryContentVideo{}

	_ StoryContentClass = &StoryContentVideo{}
)

func (s *StoryContentVideo) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Video.Zero()) {
		return false
	}
	if !(s.AlternativeVideo.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryContentVideo) String() string {
	if s == nil {
		return "StoryContentVideo(nil)"
	}
	type Alias StoryContentVideo
	return fmt.Sprintf("StoryContentVideo%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryContentVideo) TypeID() uint32 {
	return StoryContentVideoTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryContentVideo) TypeName() string {
	return "storyContentVideo"
}

// TypeInfo returns info about TL type.
func (s *StoryContentVideo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyContentVideo",
		ID:   StoryContentVideoTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Video",
			SchemaName: "video",
		},
		{
			Name:       "AlternativeVideo",
			SchemaName: "alternative_video",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryContentVideo) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyContentVideo#b30162a6 as nil")
	}
	b.PutID(StoryContentVideoTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryContentVideo) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyContentVideo#b30162a6 as nil")
	}
	if err := s.Video.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyContentVideo#b30162a6: field video: %w", err)
	}
	if err := s.AlternativeVideo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyContentVideo#b30162a6: field alternative_video: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryContentVideo) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyContentVideo#b30162a6 to nil")
	}
	if err := b.ConsumeID(StoryContentVideoTypeID); err != nil {
		return fmt.Errorf("unable to decode storyContentVideo#b30162a6: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryContentVideo) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyContentVideo#b30162a6 to nil")
	}
	{
		if err := s.Video.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyContentVideo#b30162a6: field video: %w", err)
		}
	}
	{
		if err := s.AlternativeVideo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyContentVideo#b30162a6: field alternative_video: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryContentVideo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyContentVideo#b30162a6 as nil")
	}
	b.ObjStart()
	b.PutID("storyContentVideo")
	b.Comma()
	b.FieldStart("video")
	if err := s.Video.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode storyContentVideo#b30162a6: field video: %w", err)
	}
	b.Comma()
	b.FieldStart("alternative_video")
	if err := s.AlternativeVideo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode storyContentVideo#b30162a6: field alternative_video: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryContentVideo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyContentVideo#b30162a6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyContentVideo"); err != nil {
				return fmt.Errorf("unable to decode storyContentVideo#b30162a6: %w", err)
			}
		case "video":
			if err := s.Video.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode storyContentVideo#b30162a6: field video: %w", err)
			}
		case "alternative_video":
			if err := s.AlternativeVideo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode storyContentVideo#b30162a6: field alternative_video: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetVideo returns value of Video field.
func (s *StoryContentVideo) GetVideo() (value StoryVideo) {
	if s == nil {
		return
	}
	return s.Video
}

// GetAlternativeVideo returns value of AlternativeVideo field.
func (s *StoryContentVideo) GetAlternativeVideo() (value StoryVideo) {
	if s == nil {
		return
	}
	return s.AlternativeVideo
}

// StoryContentUnsupported represents TL type `storyContentUnsupported#86c7f56e`.
type StoryContentUnsupported struct {
}

// StoryContentUnsupportedTypeID is TL type id of StoryContentUnsupported.
const StoryContentUnsupportedTypeID = 0x86c7f56e

// construct implements constructor of StoryContentClass.
func (s StoryContentUnsupported) construct() StoryContentClass { return &s }

// Ensuring interfaces in compile-time for StoryContentUnsupported.
var (
	_ bin.Encoder     = &StoryContentUnsupported{}
	_ bin.Decoder     = &StoryContentUnsupported{}
	_ bin.BareEncoder = &StoryContentUnsupported{}
	_ bin.BareDecoder = &StoryContentUnsupported{}

	_ StoryContentClass = &StoryContentUnsupported{}
)

func (s *StoryContentUnsupported) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryContentUnsupported) String() string {
	if s == nil {
		return "StoryContentUnsupported(nil)"
	}
	type Alias StoryContentUnsupported
	return fmt.Sprintf("StoryContentUnsupported%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryContentUnsupported) TypeID() uint32 {
	return StoryContentUnsupportedTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryContentUnsupported) TypeName() string {
	return "storyContentUnsupported"
}

// TypeInfo returns info about TL type.
func (s *StoryContentUnsupported) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyContentUnsupported",
		ID:   StoryContentUnsupportedTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryContentUnsupported) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyContentUnsupported#86c7f56e as nil")
	}
	b.PutID(StoryContentUnsupportedTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryContentUnsupported) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyContentUnsupported#86c7f56e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryContentUnsupported) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyContentUnsupported#86c7f56e to nil")
	}
	if err := b.ConsumeID(StoryContentUnsupportedTypeID); err != nil {
		return fmt.Errorf("unable to decode storyContentUnsupported#86c7f56e: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryContentUnsupported) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyContentUnsupported#86c7f56e to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryContentUnsupported) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyContentUnsupported#86c7f56e as nil")
	}
	b.ObjStart()
	b.PutID("storyContentUnsupported")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryContentUnsupported) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyContentUnsupported#86c7f56e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyContentUnsupported"); err != nil {
				return fmt.Errorf("unable to decode storyContentUnsupported#86c7f56e: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// StoryContentClassName is schema name of StoryContentClass.
const StoryContentClassName = "StoryContent"

// StoryContentClass represents StoryContent generic type.
//
// Example:
//
//	g, err := tdapi.DecodeStoryContent(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.StoryContentPhoto: // storyContentPhoto#d45f0050
//	case *tdapi.StoryContentVideo: // storyContentVideo#b30162a6
//	case *tdapi.StoryContentUnsupported: // storyContentUnsupported#86c7f56e
//	default: panic(v)
//	}
type StoryContentClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StoryContentClass

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

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeStoryContent implements binary de-serialization for StoryContentClass.
func DecodeStoryContent(buf *bin.Buffer) (StoryContentClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StoryContentPhotoTypeID:
		// Decoding storyContentPhoto#d45f0050.
		v := StoryContentPhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryContentClass: %w", err)
		}
		return &v, nil
	case StoryContentVideoTypeID:
		// Decoding storyContentVideo#b30162a6.
		v := StoryContentVideo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryContentClass: %w", err)
		}
		return &v, nil
	case StoryContentUnsupportedTypeID:
		// Decoding storyContentUnsupported#86c7f56e.
		v := StoryContentUnsupported{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryContentClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StoryContentClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONStoryContent implements binary de-serialization for StoryContentClass.
func DecodeTDLibJSONStoryContent(buf tdjson.Decoder) (StoryContentClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "storyContentPhoto":
		// Decoding storyContentPhoto#d45f0050.
		v := StoryContentPhoto{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryContentClass: %w", err)
		}
		return &v, nil
	case "storyContentVideo":
		// Decoding storyContentVideo#b30162a6.
		v := StoryContentVideo{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryContentClass: %w", err)
		}
		return &v, nil
	case "storyContentUnsupported":
		// Decoding storyContentUnsupported#86c7f56e.
		v := StoryContentUnsupported{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryContentClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StoryContentClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// StoryContent boxes the StoryContentClass providing a helper.
type StoryContentBox struct {
	StoryContent StoryContentClass
}

// Decode implements bin.Decoder for StoryContentBox.
func (b *StoryContentBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StoryContentBox to nil")
	}
	v, err := DecodeStoryContent(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StoryContent = v
	return nil
}

// Encode implements bin.Encode for StoryContentBox.
func (b *StoryContentBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StoryContent == nil {
		return fmt.Errorf("unable to encode StoryContentClass as nil")
	}
	return b.StoryContent.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for StoryContentBox.
func (b *StoryContentBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode StoryContentBox to nil")
	}
	v, err := DecodeTDLibJSONStoryContent(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StoryContent = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for StoryContentBox.
func (b *StoryContentBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.StoryContent == nil {
		return fmt.Errorf("unable to encode StoryContentClass as nil")
	}
	return b.StoryContent.EncodeTDLibJSON(buf)
}
