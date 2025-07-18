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

// InputStoryContentPhoto represents TL type `inputStoryContentPhoto#3286fbe0`.
type InputStoryContentPhoto struct {
	// Photo to send. The photo must be at most 10 MB in size. The photo size must be
	// 1080x1920
	Photo InputFileClass
	// File identifiers of the stickers added to the photo, if applicable
	AddedStickerFileIDs []int32
}

// InputStoryContentPhotoTypeID is TL type id of InputStoryContentPhoto.
const InputStoryContentPhotoTypeID = 0x3286fbe0

// construct implements constructor of InputStoryContentClass.
func (i InputStoryContentPhoto) construct() InputStoryContentClass { return &i }

// Ensuring interfaces in compile-time for InputStoryContentPhoto.
var (
	_ bin.Encoder     = &InputStoryContentPhoto{}
	_ bin.Decoder     = &InputStoryContentPhoto{}
	_ bin.BareEncoder = &InputStoryContentPhoto{}
	_ bin.BareDecoder = &InputStoryContentPhoto{}

	_ InputStoryContentClass = &InputStoryContentPhoto{}
)

func (i *InputStoryContentPhoto) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Photo == nil) {
		return false
	}
	if !(i.AddedStickerFileIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStoryContentPhoto) String() string {
	if i == nil {
		return "InputStoryContentPhoto(nil)"
	}
	type Alias InputStoryContentPhoto
	return fmt.Sprintf("InputStoryContentPhoto%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStoryContentPhoto) TypeID() uint32 {
	return InputStoryContentPhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStoryContentPhoto) TypeName() string {
	return "inputStoryContentPhoto"
}

// TypeInfo returns info about TL type.
func (i *InputStoryContentPhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStoryContentPhoto",
		ID:   InputStoryContentPhotoTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "AddedStickerFileIDs",
			SchemaName: "added_sticker_file_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStoryContentPhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryContentPhoto#3286fbe0 as nil")
	}
	b.PutID(InputStoryContentPhotoTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStoryContentPhoto) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryContentPhoto#3286fbe0 as nil")
	}
	if i.Photo == nil {
		return fmt.Errorf("unable to encode inputStoryContentPhoto#3286fbe0: field photo is nil")
	}
	if err := i.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStoryContentPhoto#3286fbe0: field photo: %w", err)
	}
	b.PutInt(len(i.AddedStickerFileIDs))
	for _, v := range i.AddedStickerFileIDs {
		b.PutInt32(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStoryContentPhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryContentPhoto#3286fbe0 to nil")
	}
	if err := b.ConsumeID(InputStoryContentPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStoryContentPhoto#3286fbe0: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStoryContentPhoto) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryContentPhoto#3286fbe0 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryContentPhoto#3286fbe0: field photo: %w", err)
		}
		i.Photo = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryContentPhoto#3286fbe0: field added_sticker_file_ids: %w", err)
		}

		if headerLen > 0 {
			i.AddedStickerFileIDs = make([]int32, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryContentPhoto#3286fbe0: field added_sticker_file_ids: %w", err)
			}
			i.AddedStickerFileIDs = append(i.AddedStickerFileIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputStoryContentPhoto) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryContentPhoto#3286fbe0 as nil")
	}
	b.ObjStart()
	b.PutID("inputStoryContentPhoto")
	b.Comma()
	b.FieldStart("photo")
	if i.Photo == nil {
		return fmt.Errorf("unable to encode inputStoryContentPhoto#3286fbe0: field photo is nil")
	}
	if err := i.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputStoryContentPhoto#3286fbe0: field photo: %w", err)
	}
	b.Comma()
	b.FieldStart("added_sticker_file_ids")
	b.ArrStart()
	for _, v := range i.AddedStickerFileIDs {
		b.PutInt32(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputStoryContentPhoto) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryContentPhoto#3286fbe0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputStoryContentPhoto"); err != nil {
				return fmt.Errorf("unable to decode inputStoryContentPhoto#3286fbe0: %w", err)
			}
		case "photo":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryContentPhoto#3286fbe0: field photo: %w", err)
			}
			i.Photo = value
		case "added_sticker_file_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int32()
				if err != nil {
					return fmt.Errorf("unable to decode inputStoryContentPhoto#3286fbe0: field added_sticker_file_ids: %w", err)
				}
				i.AddedStickerFileIDs = append(i.AddedStickerFileIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode inputStoryContentPhoto#3286fbe0: field added_sticker_file_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPhoto returns value of Photo field.
func (i *InputStoryContentPhoto) GetPhoto() (value InputFileClass) {
	if i == nil {
		return
	}
	return i.Photo
}

// GetAddedStickerFileIDs returns value of AddedStickerFileIDs field.
func (i *InputStoryContentPhoto) GetAddedStickerFileIDs() (value []int32) {
	if i == nil {
		return
	}
	return i.AddedStickerFileIDs
}

// InputStoryContentVideo represents TL type `inputStoryContentVideo#cc1e4239`.
type InputStoryContentVideo struct {
	// Video to be sent. The video size must be 720x1280. The video must be streamable and
	// stored in MPEG4 format, after encoding with H.265 codec and key frames added each
	// second
	Video InputFileClass
	// File identifiers of the stickers added to the video, if applicable
	AddedStickerFileIDs []int32
	// Precise duration of the video, in seconds; 0-60
	Duration float64
	// Timestamp of the frame, which will be used as video thumbnail
	CoverFrameTimestamp float64
	// True, if the video has no sound
	IsAnimation bool
}

// InputStoryContentVideoTypeID is TL type id of InputStoryContentVideo.
const InputStoryContentVideoTypeID = 0xcc1e4239

// construct implements constructor of InputStoryContentClass.
func (i InputStoryContentVideo) construct() InputStoryContentClass { return &i }

// Ensuring interfaces in compile-time for InputStoryContentVideo.
var (
	_ bin.Encoder     = &InputStoryContentVideo{}
	_ bin.Decoder     = &InputStoryContentVideo{}
	_ bin.BareEncoder = &InputStoryContentVideo{}
	_ bin.BareDecoder = &InputStoryContentVideo{}

	_ InputStoryContentClass = &InputStoryContentVideo{}
)

func (i *InputStoryContentVideo) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Video == nil) {
		return false
	}
	if !(i.AddedStickerFileIDs == nil) {
		return false
	}
	if !(i.Duration == 0) {
		return false
	}
	if !(i.CoverFrameTimestamp == 0) {
		return false
	}
	if !(i.IsAnimation == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStoryContentVideo) String() string {
	if i == nil {
		return "InputStoryContentVideo(nil)"
	}
	type Alias InputStoryContentVideo
	return fmt.Sprintf("InputStoryContentVideo%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStoryContentVideo) TypeID() uint32 {
	return InputStoryContentVideoTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStoryContentVideo) TypeName() string {
	return "inputStoryContentVideo"
}

// TypeInfo returns info about TL type.
func (i *InputStoryContentVideo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStoryContentVideo",
		ID:   InputStoryContentVideoTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Video",
			SchemaName: "video",
		},
		{
			Name:       "AddedStickerFileIDs",
			SchemaName: "added_sticker_file_ids",
		},
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
		{
			Name:       "CoverFrameTimestamp",
			SchemaName: "cover_frame_timestamp",
		},
		{
			Name:       "IsAnimation",
			SchemaName: "is_animation",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStoryContentVideo) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryContentVideo#cc1e4239 as nil")
	}
	b.PutID(InputStoryContentVideoTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStoryContentVideo) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryContentVideo#cc1e4239 as nil")
	}
	if i.Video == nil {
		return fmt.Errorf("unable to encode inputStoryContentVideo#cc1e4239: field video is nil")
	}
	if err := i.Video.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStoryContentVideo#cc1e4239: field video: %w", err)
	}
	b.PutInt(len(i.AddedStickerFileIDs))
	for _, v := range i.AddedStickerFileIDs {
		b.PutInt32(v)
	}
	b.PutDouble(i.Duration)
	b.PutDouble(i.CoverFrameTimestamp)
	b.PutBool(i.IsAnimation)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStoryContentVideo) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryContentVideo#cc1e4239 to nil")
	}
	if err := b.ConsumeID(InputStoryContentVideoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStoryContentVideo#cc1e4239: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStoryContentVideo) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryContentVideo#cc1e4239 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryContentVideo#cc1e4239: field video: %w", err)
		}
		i.Video = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryContentVideo#cc1e4239: field added_sticker_file_ids: %w", err)
		}

		if headerLen > 0 {
			i.AddedStickerFileIDs = make([]int32, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryContentVideo#cc1e4239: field added_sticker_file_ids: %w", err)
			}
			i.AddedStickerFileIDs = append(i.AddedStickerFileIDs, value)
		}
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryContentVideo#cc1e4239: field duration: %w", err)
		}
		i.Duration = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryContentVideo#cc1e4239: field cover_frame_timestamp: %w", err)
		}
		i.CoverFrameTimestamp = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryContentVideo#cc1e4239: field is_animation: %w", err)
		}
		i.IsAnimation = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputStoryContentVideo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryContentVideo#cc1e4239 as nil")
	}
	b.ObjStart()
	b.PutID("inputStoryContentVideo")
	b.Comma()
	b.FieldStart("video")
	if i.Video == nil {
		return fmt.Errorf("unable to encode inputStoryContentVideo#cc1e4239: field video is nil")
	}
	if err := i.Video.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputStoryContentVideo#cc1e4239: field video: %w", err)
	}
	b.Comma()
	b.FieldStart("added_sticker_file_ids")
	b.ArrStart()
	for _, v := range i.AddedStickerFileIDs {
		b.PutInt32(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("duration")
	b.PutDouble(i.Duration)
	b.Comma()
	b.FieldStart("cover_frame_timestamp")
	b.PutDouble(i.CoverFrameTimestamp)
	b.Comma()
	b.FieldStart("is_animation")
	b.PutBool(i.IsAnimation)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputStoryContentVideo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryContentVideo#cc1e4239 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputStoryContentVideo"); err != nil {
				return fmt.Errorf("unable to decode inputStoryContentVideo#cc1e4239: %w", err)
			}
		case "video":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryContentVideo#cc1e4239: field video: %w", err)
			}
			i.Video = value
		case "added_sticker_file_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int32()
				if err != nil {
					return fmt.Errorf("unable to decode inputStoryContentVideo#cc1e4239: field added_sticker_file_ids: %w", err)
				}
				i.AddedStickerFileIDs = append(i.AddedStickerFileIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode inputStoryContentVideo#cc1e4239: field added_sticker_file_ids: %w", err)
			}
		case "duration":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryContentVideo#cc1e4239: field duration: %w", err)
			}
			i.Duration = value
		case "cover_frame_timestamp":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryContentVideo#cc1e4239: field cover_frame_timestamp: %w", err)
			}
			i.CoverFrameTimestamp = value
		case "is_animation":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryContentVideo#cc1e4239: field is_animation: %w", err)
			}
			i.IsAnimation = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetVideo returns value of Video field.
func (i *InputStoryContentVideo) GetVideo() (value InputFileClass) {
	if i == nil {
		return
	}
	return i.Video
}

// GetAddedStickerFileIDs returns value of AddedStickerFileIDs field.
func (i *InputStoryContentVideo) GetAddedStickerFileIDs() (value []int32) {
	if i == nil {
		return
	}
	return i.AddedStickerFileIDs
}

// GetDuration returns value of Duration field.
func (i *InputStoryContentVideo) GetDuration() (value float64) {
	if i == nil {
		return
	}
	return i.Duration
}

// GetCoverFrameTimestamp returns value of CoverFrameTimestamp field.
func (i *InputStoryContentVideo) GetCoverFrameTimestamp() (value float64) {
	if i == nil {
		return
	}
	return i.CoverFrameTimestamp
}

// GetIsAnimation returns value of IsAnimation field.
func (i *InputStoryContentVideo) GetIsAnimation() (value bool) {
	if i == nil {
		return
	}
	return i.IsAnimation
}

// InputStoryContentClassName is schema name of InputStoryContentClass.
const InputStoryContentClassName = "InputStoryContent"

// InputStoryContentClass represents InputStoryContent generic type.
//
// Example:
//
//	g, err := tdapi.DecodeInputStoryContent(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.InputStoryContentPhoto: // inputStoryContentPhoto#3286fbe0
//	case *tdapi.InputStoryContentVideo: // inputStoryContentVideo#cc1e4239
//	default: panic(v)
//	}
type InputStoryContentClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputStoryContentClass

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

	// File identifiers of the stickers added to the photo, if applicable
	GetAddedStickerFileIDs() (value []int32)
}

// DecodeInputStoryContent implements binary de-serialization for InputStoryContentClass.
func DecodeInputStoryContent(buf *bin.Buffer) (InputStoryContentClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputStoryContentPhotoTypeID:
		// Decoding inputStoryContentPhoto#3286fbe0.
		v := InputStoryContentPhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStoryContentClass: %w", err)
		}
		return &v, nil
	case InputStoryContentVideoTypeID:
		// Decoding inputStoryContentVideo#cc1e4239.
		v := InputStoryContentVideo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStoryContentClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputStoryContentClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONInputStoryContent implements binary de-serialization for InputStoryContentClass.
func DecodeTDLibJSONInputStoryContent(buf tdjson.Decoder) (InputStoryContentClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "inputStoryContentPhoto":
		// Decoding inputStoryContentPhoto#3286fbe0.
		v := InputStoryContentPhoto{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStoryContentClass: %w", err)
		}
		return &v, nil
	case "inputStoryContentVideo":
		// Decoding inputStoryContentVideo#cc1e4239.
		v := InputStoryContentVideo{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStoryContentClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputStoryContentClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// InputStoryContent boxes the InputStoryContentClass providing a helper.
type InputStoryContentBox struct {
	InputStoryContent InputStoryContentClass
}

// Decode implements bin.Decoder for InputStoryContentBox.
func (b *InputStoryContentBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputStoryContentBox to nil")
	}
	v, err := DecodeInputStoryContent(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputStoryContent = v
	return nil
}

// Encode implements bin.Encode for InputStoryContentBox.
func (b *InputStoryContentBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputStoryContent == nil {
		return fmt.Errorf("unable to encode InputStoryContentClass as nil")
	}
	return b.InputStoryContent.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for InputStoryContentBox.
func (b *InputStoryContentBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputStoryContentBox to nil")
	}
	v, err := DecodeTDLibJSONInputStoryContent(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputStoryContent = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for InputStoryContentBox.
func (b *InputStoryContentBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.InputStoryContent == nil {
		return fmt.Errorf("unable to encode InputStoryContentClass as nil")
	}
	return b.InputStoryContent.EncodeTDLibJSON(buf)
}
