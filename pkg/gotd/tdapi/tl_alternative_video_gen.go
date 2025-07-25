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

// AlternativeVideo represents TL type `alternativeVideo#1ccfc90e`.
type AlternativeVideo struct {
	// Unique identifier of the alternative video, which is used in the HLS file
	ID int64
	// Video width
	Width int32
	// Video height
	Height int32
	// Codec used for video file encoding, for example, "h264", "h265", or "av1"
	Codec string
	// HLS file describing the video
	HlsFile File
	// File containing the video
	Video File
}

// AlternativeVideoTypeID is TL type id of AlternativeVideo.
const AlternativeVideoTypeID = 0x1ccfc90e

// Ensuring interfaces in compile-time for AlternativeVideo.
var (
	_ bin.Encoder     = &AlternativeVideo{}
	_ bin.Decoder     = &AlternativeVideo{}
	_ bin.BareEncoder = &AlternativeVideo{}
	_ bin.BareDecoder = &AlternativeVideo{}
)

func (a *AlternativeVideo) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ID == 0) {
		return false
	}
	if !(a.Width == 0) {
		return false
	}
	if !(a.Height == 0) {
		return false
	}
	if !(a.Codec == "") {
		return false
	}
	if !(a.HlsFile.Zero()) {
		return false
	}
	if !(a.Video.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AlternativeVideo) String() string {
	if a == nil {
		return "AlternativeVideo(nil)"
	}
	type Alias AlternativeVideo
	return fmt.Sprintf("AlternativeVideo%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AlternativeVideo) TypeID() uint32 {
	return AlternativeVideoTypeID
}

// TypeName returns name of type in TL schema.
func (*AlternativeVideo) TypeName() string {
	return "alternativeVideo"
}

// TypeInfo returns info about TL type.
func (a *AlternativeVideo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "alternativeVideo",
		ID:   AlternativeVideoTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
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
			Name:       "Codec",
			SchemaName: "codec",
		},
		{
			Name:       "HlsFile",
			SchemaName: "hls_file",
		},
		{
			Name:       "Video",
			SchemaName: "video",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AlternativeVideo) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode alternativeVideo#1ccfc90e as nil")
	}
	b.PutID(AlternativeVideoTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AlternativeVideo) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode alternativeVideo#1ccfc90e as nil")
	}
	b.PutLong(a.ID)
	b.PutInt32(a.Width)
	b.PutInt32(a.Height)
	b.PutString(a.Codec)
	if err := a.HlsFile.Encode(b); err != nil {
		return fmt.Errorf("unable to encode alternativeVideo#1ccfc90e: field hls_file: %w", err)
	}
	if err := a.Video.Encode(b); err != nil {
		return fmt.Errorf("unable to encode alternativeVideo#1ccfc90e: field video: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AlternativeVideo) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode alternativeVideo#1ccfc90e to nil")
	}
	if err := b.ConsumeID(AlternativeVideoTypeID); err != nil {
		return fmt.Errorf("unable to decode alternativeVideo#1ccfc90e: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AlternativeVideo) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode alternativeVideo#1ccfc90e to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode alternativeVideo#1ccfc90e: field id: %w", err)
		}
		a.ID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode alternativeVideo#1ccfc90e: field width: %w", err)
		}
		a.Width = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode alternativeVideo#1ccfc90e: field height: %w", err)
		}
		a.Height = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode alternativeVideo#1ccfc90e: field codec: %w", err)
		}
		a.Codec = value
	}
	{
		if err := a.HlsFile.Decode(b); err != nil {
			return fmt.Errorf("unable to decode alternativeVideo#1ccfc90e: field hls_file: %w", err)
		}
	}
	{
		if err := a.Video.Decode(b); err != nil {
			return fmt.Errorf("unable to decode alternativeVideo#1ccfc90e: field video: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AlternativeVideo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode alternativeVideo#1ccfc90e as nil")
	}
	b.ObjStart()
	b.PutID("alternativeVideo")
	b.Comma()
	b.FieldStart("id")
	b.PutLong(a.ID)
	b.Comma()
	b.FieldStart("width")
	b.PutInt32(a.Width)
	b.Comma()
	b.FieldStart("height")
	b.PutInt32(a.Height)
	b.Comma()
	b.FieldStart("codec")
	b.PutString(a.Codec)
	b.Comma()
	b.FieldStart("hls_file")
	if err := a.HlsFile.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode alternativeVideo#1ccfc90e: field hls_file: %w", err)
	}
	b.Comma()
	b.FieldStart("video")
	if err := a.Video.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode alternativeVideo#1ccfc90e: field video: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AlternativeVideo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode alternativeVideo#1ccfc90e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("alternativeVideo"); err != nil {
				return fmt.Errorf("unable to decode alternativeVideo#1ccfc90e: %w", err)
			}
		case "id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode alternativeVideo#1ccfc90e: field id: %w", err)
			}
			a.ID = value
		case "width":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode alternativeVideo#1ccfc90e: field width: %w", err)
			}
			a.Width = value
		case "height":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode alternativeVideo#1ccfc90e: field height: %w", err)
			}
			a.Height = value
		case "codec":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode alternativeVideo#1ccfc90e: field codec: %w", err)
			}
			a.Codec = value
		case "hls_file":
			if err := a.HlsFile.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode alternativeVideo#1ccfc90e: field hls_file: %w", err)
			}
		case "video":
			if err := a.Video.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode alternativeVideo#1ccfc90e: field video: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (a *AlternativeVideo) GetID() (value int64) {
	if a == nil {
		return
	}
	return a.ID
}

// GetWidth returns value of Width field.
func (a *AlternativeVideo) GetWidth() (value int32) {
	if a == nil {
		return
	}
	return a.Width
}

// GetHeight returns value of Height field.
func (a *AlternativeVideo) GetHeight() (value int32) {
	if a == nil {
		return
	}
	return a.Height
}

// GetCodec returns value of Codec field.
func (a *AlternativeVideo) GetCodec() (value string) {
	if a == nil {
		return
	}
	return a.Codec
}

// GetHlsFile returns value of HlsFile field.
func (a *AlternativeVideo) GetHlsFile() (value File) {
	if a == nil {
		return
	}
	return a.HlsFile
}

// GetVideo returns value of Video field.
func (a *AlternativeVideo) GetVideo() (value File) {
	if a == nil {
		return
	}
	return a.Video
}
