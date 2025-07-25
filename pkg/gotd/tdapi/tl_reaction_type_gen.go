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

// ReactionTypeEmoji represents TL type `reactionTypeEmoji#8c3e22c8`.
type ReactionTypeEmoji struct {
	// Text representation of the reaction
	Emoji string
}

// ReactionTypeEmojiTypeID is TL type id of ReactionTypeEmoji.
const ReactionTypeEmojiTypeID = 0x8c3e22c8

// construct implements constructor of ReactionTypeClass.
func (r ReactionTypeEmoji) construct() ReactionTypeClass { return &r }

// Ensuring interfaces in compile-time for ReactionTypeEmoji.
var (
	_ bin.Encoder     = &ReactionTypeEmoji{}
	_ bin.Decoder     = &ReactionTypeEmoji{}
	_ bin.BareEncoder = &ReactionTypeEmoji{}
	_ bin.BareDecoder = &ReactionTypeEmoji{}

	_ ReactionTypeClass = &ReactionTypeEmoji{}
)

func (r *ReactionTypeEmoji) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Emoji == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReactionTypeEmoji) String() string {
	if r == nil {
		return "ReactionTypeEmoji(nil)"
	}
	type Alias ReactionTypeEmoji
	return fmt.Sprintf("ReactionTypeEmoji%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReactionTypeEmoji) TypeID() uint32 {
	return ReactionTypeEmojiTypeID
}

// TypeName returns name of type in TL schema.
func (*ReactionTypeEmoji) TypeName() string {
	return "reactionTypeEmoji"
}

// TypeInfo returns info about TL type.
func (r *ReactionTypeEmoji) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reactionTypeEmoji",
		ID:   ReactionTypeEmojiTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Emoji",
			SchemaName: "emoji",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReactionTypeEmoji) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionTypeEmoji#8c3e22c8 as nil")
	}
	b.PutID(ReactionTypeEmojiTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReactionTypeEmoji) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionTypeEmoji#8c3e22c8 as nil")
	}
	b.PutString(r.Emoji)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReactionTypeEmoji) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionTypeEmoji#8c3e22c8 to nil")
	}
	if err := b.ConsumeID(ReactionTypeEmojiTypeID); err != nil {
		return fmt.Errorf("unable to decode reactionTypeEmoji#8c3e22c8: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReactionTypeEmoji) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionTypeEmoji#8c3e22c8 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode reactionTypeEmoji#8c3e22c8: field emoji: %w", err)
		}
		r.Emoji = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReactionTypeEmoji) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionTypeEmoji#8c3e22c8 as nil")
	}
	b.ObjStart()
	b.PutID("reactionTypeEmoji")
	b.Comma()
	b.FieldStart("emoji")
	b.PutString(r.Emoji)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReactionTypeEmoji) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionTypeEmoji#8c3e22c8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reactionTypeEmoji"); err != nil {
				return fmt.Errorf("unable to decode reactionTypeEmoji#8c3e22c8: %w", err)
			}
		case "emoji":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode reactionTypeEmoji#8c3e22c8: field emoji: %w", err)
			}
			r.Emoji = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetEmoji returns value of Emoji field.
func (r *ReactionTypeEmoji) GetEmoji() (value string) {
	if r == nil {
		return
	}
	return r.Emoji
}

// ReactionTypeCustomEmoji represents TL type `reactionTypeCustomEmoji#c50b42f3`.
type ReactionTypeCustomEmoji struct {
	// Unique identifier of the custom emoji
	CustomEmojiID int64
}

// ReactionTypeCustomEmojiTypeID is TL type id of ReactionTypeCustomEmoji.
const ReactionTypeCustomEmojiTypeID = 0xc50b42f3

// construct implements constructor of ReactionTypeClass.
func (r ReactionTypeCustomEmoji) construct() ReactionTypeClass { return &r }

// Ensuring interfaces in compile-time for ReactionTypeCustomEmoji.
var (
	_ bin.Encoder     = &ReactionTypeCustomEmoji{}
	_ bin.Decoder     = &ReactionTypeCustomEmoji{}
	_ bin.BareEncoder = &ReactionTypeCustomEmoji{}
	_ bin.BareDecoder = &ReactionTypeCustomEmoji{}

	_ ReactionTypeClass = &ReactionTypeCustomEmoji{}
)

func (r *ReactionTypeCustomEmoji) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.CustomEmojiID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReactionTypeCustomEmoji) String() string {
	if r == nil {
		return "ReactionTypeCustomEmoji(nil)"
	}
	type Alias ReactionTypeCustomEmoji
	return fmt.Sprintf("ReactionTypeCustomEmoji%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReactionTypeCustomEmoji) TypeID() uint32 {
	return ReactionTypeCustomEmojiTypeID
}

// TypeName returns name of type in TL schema.
func (*ReactionTypeCustomEmoji) TypeName() string {
	return "reactionTypeCustomEmoji"
}

// TypeInfo returns info about TL type.
func (r *ReactionTypeCustomEmoji) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reactionTypeCustomEmoji",
		ID:   ReactionTypeCustomEmojiTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CustomEmojiID",
			SchemaName: "custom_emoji_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReactionTypeCustomEmoji) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionTypeCustomEmoji#c50b42f3 as nil")
	}
	b.PutID(ReactionTypeCustomEmojiTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReactionTypeCustomEmoji) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionTypeCustomEmoji#c50b42f3 as nil")
	}
	b.PutLong(r.CustomEmojiID)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReactionTypeCustomEmoji) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionTypeCustomEmoji#c50b42f3 to nil")
	}
	if err := b.ConsumeID(ReactionTypeCustomEmojiTypeID); err != nil {
		return fmt.Errorf("unable to decode reactionTypeCustomEmoji#c50b42f3: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReactionTypeCustomEmoji) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionTypeCustomEmoji#c50b42f3 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode reactionTypeCustomEmoji#c50b42f3: field custom_emoji_id: %w", err)
		}
		r.CustomEmojiID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReactionTypeCustomEmoji) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionTypeCustomEmoji#c50b42f3 as nil")
	}
	b.ObjStart()
	b.PutID("reactionTypeCustomEmoji")
	b.Comma()
	b.FieldStart("custom_emoji_id")
	b.PutLong(r.CustomEmojiID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReactionTypeCustomEmoji) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionTypeCustomEmoji#c50b42f3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reactionTypeCustomEmoji"); err != nil {
				return fmt.Errorf("unable to decode reactionTypeCustomEmoji#c50b42f3: %w", err)
			}
		case "custom_emoji_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode reactionTypeCustomEmoji#c50b42f3: field custom_emoji_id: %w", err)
			}
			r.CustomEmojiID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCustomEmojiID returns value of CustomEmojiID field.
func (r *ReactionTypeCustomEmoji) GetCustomEmojiID() (value int64) {
	if r == nil {
		return
	}
	return r.CustomEmojiID
}

// ReactionTypePaid represents TL type `reactionTypePaid#1a0152ed`.
type ReactionTypePaid struct {
}

// ReactionTypePaidTypeID is TL type id of ReactionTypePaid.
const ReactionTypePaidTypeID = 0x1a0152ed

// construct implements constructor of ReactionTypeClass.
func (r ReactionTypePaid) construct() ReactionTypeClass { return &r }

// Ensuring interfaces in compile-time for ReactionTypePaid.
var (
	_ bin.Encoder     = &ReactionTypePaid{}
	_ bin.Decoder     = &ReactionTypePaid{}
	_ bin.BareEncoder = &ReactionTypePaid{}
	_ bin.BareDecoder = &ReactionTypePaid{}

	_ ReactionTypeClass = &ReactionTypePaid{}
)

func (r *ReactionTypePaid) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReactionTypePaid) String() string {
	if r == nil {
		return "ReactionTypePaid(nil)"
	}
	type Alias ReactionTypePaid
	return fmt.Sprintf("ReactionTypePaid%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReactionTypePaid) TypeID() uint32 {
	return ReactionTypePaidTypeID
}

// TypeName returns name of type in TL schema.
func (*ReactionTypePaid) TypeName() string {
	return "reactionTypePaid"
}

// TypeInfo returns info about TL type.
func (r *ReactionTypePaid) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reactionTypePaid",
		ID:   ReactionTypePaidTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReactionTypePaid) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionTypePaid#1a0152ed as nil")
	}
	b.PutID(ReactionTypePaidTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReactionTypePaid) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionTypePaid#1a0152ed as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReactionTypePaid) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionTypePaid#1a0152ed to nil")
	}
	if err := b.ConsumeID(ReactionTypePaidTypeID); err != nil {
		return fmt.Errorf("unable to decode reactionTypePaid#1a0152ed: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReactionTypePaid) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionTypePaid#1a0152ed to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReactionTypePaid) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionTypePaid#1a0152ed as nil")
	}
	b.ObjStart()
	b.PutID("reactionTypePaid")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReactionTypePaid) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionTypePaid#1a0152ed to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reactionTypePaid"); err != nil {
				return fmt.Errorf("unable to decode reactionTypePaid#1a0152ed: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ReactionTypeClassName is schema name of ReactionTypeClass.
const ReactionTypeClassName = "ReactionType"

// ReactionTypeClass represents ReactionType generic type.
//
// Example:
//
//	g, err := tdapi.DecodeReactionType(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.ReactionTypeEmoji: // reactionTypeEmoji#8c3e22c8
//	case *tdapi.ReactionTypeCustomEmoji: // reactionTypeCustomEmoji#c50b42f3
//	case *tdapi.ReactionTypePaid: // reactionTypePaid#1a0152ed
//	default: panic(v)
//	}
type ReactionTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ReactionTypeClass

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

// DecodeReactionType implements binary de-serialization for ReactionTypeClass.
func DecodeReactionType(buf *bin.Buffer) (ReactionTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ReactionTypeEmojiTypeID:
		// Decoding reactionTypeEmoji#8c3e22c8.
		v := ReactionTypeEmoji{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReactionTypeClass: %w", err)
		}
		return &v, nil
	case ReactionTypeCustomEmojiTypeID:
		// Decoding reactionTypeCustomEmoji#c50b42f3.
		v := ReactionTypeCustomEmoji{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReactionTypeClass: %w", err)
		}
		return &v, nil
	case ReactionTypePaidTypeID:
		// Decoding reactionTypePaid#1a0152ed.
		v := ReactionTypePaid{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReactionTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ReactionTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONReactionType implements binary de-serialization for ReactionTypeClass.
func DecodeTDLibJSONReactionType(buf tdjson.Decoder) (ReactionTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "reactionTypeEmoji":
		// Decoding reactionTypeEmoji#8c3e22c8.
		v := ReactionTypeEmoji{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReactionTypeClass: %w", err)
		}
		return &v, nil
	case "reactionTypeCustomEmoji":
		// Decoding reactionTypeCustomEmoji#c50b42f3.
		v := ReactionTypeCustomEmoji{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReactionTypeClass: %w", err)
		}
		return &v, nil
	case "reactionTypePaid":
		// Decoding reactionTypePaid#1a0152ed.
		v := ReactionTypePaid{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReactionTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ReactionTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// ReactionType boxes the ReactionTypeClass providing a helper.
type ReactionTypeBox struct {
	ReactionType ReactionTypeClass
}

// Decode implements bin.Decoder for ReactionTypeBox.
func (b *ReactionTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ReactionTypeBox to nil")
	}
	v, err := DecodeReactionType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ReactionType = v
	return nil
}

// Encode implements bin.Encode for ReactionTypeBox.
func (b *ReactionTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ReactionType == nil {
		return fmt.Errorf("unable to encode ReactionTypeClass as nil")
	}
	return b.ReactionType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for ReactionTypeBox.
func (b *ReactionTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode ReactionTypeBox to nil")
	}
	v, err := DecodeTDLibJSONReactionType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ReactionType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for ReactionTypeBox.
func (b *ReactionTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.ReactionType == nil {
		return fmt.Errorf("unable to encode ReactionTypeClass as nil")
	}
	return b.ReactionType.EncodeTDLibJSON(buf)
}
