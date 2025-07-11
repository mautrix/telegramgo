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

// StoryListMain represents TL type `storyListMain#d7eeb3ff`.
type StoryListMain struct {
}

// StoryListMainTypeID is TL type id of StoryListMain.
const StoryListMainTypeID = 0xd7eeb3ff

// construct implements constructor of StoryListClass.
func (s StoryListMain) construct() StoryListClass { return &s }

// Ensuring interfaces in compile-time for StoryListMain.
var (
	_ bin.Encoder     = &StoryListMain{}
	_ bin.Decoder     = &StoryListMain{}
	_ bin.BareEncoder = &StoryListMain{}
	_ bin.BareDecoder = &StoryListMain{}

	_ StoryListClass = &StoryListMain{}
)

func (s *StoryListMain) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryListMain) String() string {
	if s == nil {
		return "StoryListMain(nil)"
	}
	type Alias StoryListMain
	return fmt.Sprintf("StoryListMain%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryListMain) TypeID() uint32 {
	return StoryListMainTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryListMain) TypeName() string {
	return "storyListMain"
}

// TypeInfo returns info about TL type.
func (s *StoryListMain) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyListMain",
		ID:   StoryListMainTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryListMain) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyListMain#d7eeb3ff as nil")
	}
	b.PutID(StoryListMainTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryListMain) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyListMain#d7eeb3ff as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryListMain) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyListMain#d7eeb3ff to nil")
	}
	if err := b.ConsumeID(StoryListMainTypeID); err != nil {
		return fmt.Errorf("unable to decode storyListMain#d7eeb3ff: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryListMain) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyListMain#d7eeb3ff to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryListMain) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyListMain#d7eeb3ff as nil")
	}
	b.ObjStart()
	b.PutID("storyListMain")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryListMain) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyListMain#d7eeb3ff to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyListMain"); err != nil {
				return fmt.Errorf("unable to decode storyListMain#d7eeb3ff: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// StoryListArchive represents TL type `storyListArchive#fd80a741`.
type StoryListArchive struct {
}

// StoryListArchiveTypeID is TL type id of StoryListArchive.
const StoryListArchiveTypeID = 0xfd80a741

// construct implements constructor of StoryListClass.
func (s StoryListArchive) construct() StoryListClass { return &s }

// Ensuring interfaces in compile-time for StoryListArchive.
var (
	_ bin.Encoder     = &StoryListArchive{}
	_ bin.Decoder     = &StoryListArchive{}
	_ bin.BareEncoder = &StoryListArchive{}
	_ bin.BareDecoder = &StoryListArchive{}

	_ StoryListClass = &StoryListArchive{}
)

func (s *StoryListArchive) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryListArchive) String() string {
	if s == nil {
		return "StoryListArchive(nil)"
	}
	type Alias StoryListArchive
	return fmt.Sprintf("StoryListArchive%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryListArchive) TypeID() uint32 {
	return StoryListArchiveTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryListArchive) TypeName() string {
	return "storyListArchive"
}

// TypeInfo returns info about TL type.
func (s *StoryListArchive) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyListArchive",
		ID:   StoryListArchiveTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryListArchive) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyListArchive#fd80a741 as nil")
	}
	b.PutID(StoryListArchiveTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryListArchive) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyListArchive#fd80a741 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryListArchive) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyListArchive#fd80a741 to nil")
	}
	if err := b.ConsumeID(StoryListArchiveTypeID); err != nil {
		return fmt.Errorf("unable to decode storyListArchive#fd80a741: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryListArchive) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyListArchive#fd80a741 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryListArchive) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyListArchive#fd80a741 as nil")
	}
	b.ObjStart()
	b.PutID("storyListArchive")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryListArchive) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyListArchive#fd80a741 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyListArchive"); err != nil {
				return fmt.Errorf("unable to decode storyListArchive#fd80a741: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// StoryListClassName is schema name of StoryListClass.
const StoryListClassName = "StoryList"

// StoryListClass represents StoryList generic type.
//
// Example:
//
//	g, err := tdapi.DecodeStoryList(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.StoryListMain: // storyListMain#d7eeb3ff
//	case *tdapi.StoryListArchive: // storyListArchive#fd80a741
//	default: panic(v)
//	}
type StoryListClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StoryListClass

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

// DecodeStoryList implements binary de-serialization for StoryListClass.
func DecodeStoryList(buf *bin.Buffer) (StoryListClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StoryListMainTypeID:
		// Decoding storyListMain#d7eeb3ff.
		v := StoryListMain{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryListClass: %w", err)
		}
		return &v, nil
	case StoryListArchiveTypeID:
		// Decoding storyListArchive#fd80a741.
		v := StoryListArchive{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryListClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StoryListClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONStoryList implements binary de-serialization for StoryListClass.
func DecodeTDLibJSONStoryList(buf tdjson.Decoder) (StoryListClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "storyListMain":
		// Decoding storyListMain#d7eeb3ff.
		v := StoryListMain{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryListClass: %w", err)
		}
		return &v, nil
	case "storyListArchive":
		// Decoding storyListArchive#fd80a741.
		v := StoryListArchive{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryListClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StoryListClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// StoryList boxes the StoryListClass providing a helper.
type StoryListBox struct {
	StoryList StoryListClass
}

// Decode implements bin.Decoder for StoryListBox.
func (b *StoryListBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StoryListBox to nil")
	}
	v, err := DecodeStoryList(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StoryList = v
	return nil
}

// Encode implements bin.Encode for StoryListBox.
func (b *StoryListBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StoryList == nil {
		return fmt.Errorf("unable to encode StoryListClass as nil")
	}
	return b.StoryList.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for StoryListBox.
func (b *StoryListBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode StoryListBox to nil")
	}
	v, err := DecodeTDLibJSONStoryList(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StoryList = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for StoryListBox.
func (b *StoryListBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.StoryList == nil {
		return fmt.Errorf("unable to encode StoryListClass as nil")
	}
	return b.StoryList.EncodeTDLibJSON(buf)
}
