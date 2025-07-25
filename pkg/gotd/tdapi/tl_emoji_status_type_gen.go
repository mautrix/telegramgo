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

// EmojiStatusTypeCustomEmoji represents TL type `emojiStatusTypeCustomEmoji#9ca6f0f5`.
type EmojiStatusTypeCustomEmoji struct {
	// Identifier of the custom emoji in stickerFormatTgs format
	CustomEmojiID int64
}

// EmojiStatusTypeCustomEmojiTypeID is TL type id of EmojiStatusTypeCustomEmoji.
const EmojiStatusTypeCustomEmojiTypeID = 0x9ca6f0f5

// construct implements constructor of EmojiStatusTypeClass.
func (e EmojiStatusTypeCustomEmoji) construct() EmojiStatusTypeClass { return &e }

// Ensuring interfaces in compile-time for EmojiStatusTypeCustomEmoji.
var (
	_ bin.Encoder     = &EmojiStatusTypeCustomEmoji{}
	_ bin.Decoder     = &EmojiStatusTypeCustomEmoji{}
	_ bin.BareEncoder = &EmojiStatusTypeCustomEmoji{}
	_ bin.BareDecoder = &EmojiStatusTypeCustomEmoji{}

	_ EmojiStatusTypeClass = &EmojiStatusTypeCustomEmoji{}
)

func (e *EmojiStatusTypeCustomEmoji) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.CustomEmojiID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmojiStatusTypeCustomEmoji) String() string {
	if e == nil {
		return "EmojiStatusTypeCustomEmoji(nil)"
	}
	type Alias EmojiStatusTypeCustomEmoji
	return fmt.Sprintf("EmojiStatusTypeCustomEmoji%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmojiStatusTypeCustomEmoji) TypeID() uint32 {
	return EmojiStatusTypeCustomEmojiTypeID
}

// TypeName returns name of type in TL schema.
func (*EmojiStatusTypeCustomEmoji) TypeName() string {
	return "emojiStatusTypeCustomEmoji"
}

// TypeInfo returns info about TL type.
func (e *EmojiStatusTypeCustomEmoji) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emojiStatusTypeCustomEmoji",
		ID:   EmojiStatusTypeCustomEmojiTypeID,
	}
	if e == nil {
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
func (e *EmojiStatusTypeCustomEmoji) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiStatusTypeCustomEmoji#9ca6f0f5 as nil")
	}
	b.PutID(EmojiStatusTypeCustomEmojiTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EmojiStatusTypeCustomEmoji) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiStatusTypeCustomEmoji#9ca6f0f5 as nil")
	}
	b.PutLong(e.CustomEmojiID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EmojiStatusTypeCustomEmoji) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiStatusTypeCustomEmoji#9ca6f0f5 to nil")
	}
	if err := b.ConsumeID(EmojiStatusTypeCustomEmojiTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiStatusTypeCustomEmoji#9ca6f0f5: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EmojiStatusTypeCustomEmoji) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiStatusTypeCustomEmoji#9ca6f0f5 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode emojiStatusTypeCustomEmoji#9ca6f0f5: field custom_emoji_id: %w", err)
		}
		e.CustomEmojiID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EmojiStatusTypeCustomEmoji) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiStatusTypeCustomEmoji#9ca6f0f5 as nil")
	}
	b.ObjStart()
	b.PutID("emojiStatusTypeCustomEmoji")
	b.Comma()
	b.FieldStart("custom_emoji_id")
	b.PutLong(e.CustomEmojiID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EmojiStatusTypeCustomEmoji) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiStatusTypeCustomEmoji#9ca6f0f5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("emojiStatusTypeCustomEmoji"); err != nil {
				return fmt.Errorf("unable to decode emojiStatusTypeCustomEmoji#9ca6f0f5: %w", err)
			}
		case "custom_emoji_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode emojiStatusTypeCustomEmoji#9ca6f0f5: field custom_emoji_id: %w", err)
			}
			e.CustomEmojiID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCustomEmojiID returns value of CustomEmojiID field.
func (e *EmojiStatusTypeCustomEmoji) GetCustomEmojiID() (value int64) {
	if e == nil {
		return
	}
	return e.CustomEmojiID
}

// EmojiStatusTypeUpgradedGift represents TL type `emojiStatusTypeUpgradedGift#ce0e53f4`.
type EmojiStatusTypeUpgradedGift struct {
	// Identifier of the upgraded gift
	UpgradedGiftID int64
	// The title of the upgraded gift
	GiftTitle string
	// Unique name of the upgraded gift that can be used with internalLinkTypeUpgradedGift
	GiftName string
	// Custom emoji identifier of the model of the upgraded gift
	ModelCustomEmojiID int64
	// Custom emoji identifier of the symbol of the upgraded gift
	SymbolCustomEmojiID int64
	// Colors of the backdrop of the upgraded gift
	BackdropColors UpgradedGiftBackdropColors
}

// EmojiStatusTypeUpgradedGiftTypeID is TL type id of EmojiStatusTypeUpgradedGift.
const EmojiStatusTypeUpgradedGiftTypeID = 0xce0e53f4

// construct implements constructor of EmojiStatusTypeClass.
func (e EmojiStatusTypeUpgradedGift) construct() EmojiStatusTypeClass { return &e }

// Ensuring interfaces in compile-time for EmojiStatusTypeUpgradedGift.
var (
	_ bin.Encoder     = &EmojiStatusTypeUpgradedGift{}
	_ bin.Decoder     = &EmojiStatusTypeUpgradedGift{}
	_ bin.BareEncoder = &EmojiStatusTypeUpgradedGift{}
	_ bin.BareDecoder = &EmojiStatusTypeUpgradedGift{}

	_ EmojiStatusTypeClass = &EmojiStatusTypeUpgradedGift{}
)

func (e *EmojiStatusTypeUpgradedGift) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.UpgradedGiftID == 0) {
		return false
	}
	if !(e.GiftTitle == "") {
		return false
	}
	if !(e.GiftName == "") {
		return false
	}
	if !(e.ModelCustomEmojiID == 0) {
		return false
	}
	if !(e.SymbolCustomEmojiID == 0) {
		return false
	}
	if !(e.BackdropColors.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmojiStatusTypeUpgradedGift) String() string {
	if e == nil {
		return "EmojiStatusTypeUpgradedGift(nil)"
	}
	type Alias EmojiStatusTypeUpgradedGift
	return fmt.Sprintf("EmojiStatusTypeUpgradedGift%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmojiStatusTypeUpgradedGift) TypeID() uint32 {
	return EmojiStatusTypeUpgradedGiftTypeID
}

// TypeName returns name of type in TL schema.
func (*EmojiStatusTypeUpgradedGift) TypeName() string {
	return "emojiStatusTypeUpgradedGift"
}

// TypeInfo returns info about TL type.
func (e *EmojiStatusTypeUpgradedGift) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emojiStatusTypeUpgradedGift",
		ID:   EmojiStatusTypeUpgradedGiftTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UpgradedGiftID",
			SchemaName: "upgraded_gift_id",
		},
		{
			Name:       "GiftTitle",
			SchemaName: "gift_title",
		},
		{
			Name:       "GiftName",
			SchemaName: "gift_name",
		},
		{
			Name:       "ModelCustomEmojiID",
			SchemaName: "model_custom_emoji_id",
		},
		{
			Name:       "SymbolCustomEmojiID",
			SchemaName: "symbol_custom_emoji_id",
		},
		{
			Name:       "BackdropColors",
			SchemaName: "backdrop_colors",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EmojiStatusTypeUpgradedGift) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiStatusTypeUpgradedGift#ce0e53f4 as nil")
	}
	b.PutID(EmojiStatusTypeUpgradedGiftTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EmojiStatusTypeUpgradedGift) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiStatusTypeUpgradedGift#ce0e53f4 as nil")
	}
	b.PutLong(e.UpgradedGiftID)
	b.PutString(e.GiftTitle)
	b.PutString(e.GiftName)
	b.PutLong(e.ModelCustomEmojiID)
	b.PutLong(e.SymbolCustomEmojiID)
	if err := e.BackdropColors.Encode(b); err != nil {
		return fmt.Errorf("unable to encode emojiStatusTypeUpgradedGift#ce0e53f4: field backdrop_colors: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EmojiStatusTypeUpgradedGift) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiStatusTypeUpgradedGift#ce0e53f4 to nil")
	}
	if err := b.ConsumeID(EmojiStatusTypeUpgradedGiftTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiStatusTypeUpgradedGift#ce0e53f4: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EmojiStatusTypeUpgradedGift) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiStatusTypeUpgradedGift#ce0e53f4 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode emojiStatusTypeUpgradedGift#ce0e53f4: field upgraded_gift_id: %w", err)
		}
		e.UpgradedGiftID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emojiStatusTypeUpgradedGift#ce0e53f4: field gift_title: %w", err)
		}
		e.GiftTitle = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emojiStatusTypeUpgradedGift#ce0e53f4: field gift_name: %w", err)
		}
		e.GiftName = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode emojiStatusTypeUpgradedGift#ce0e53f4: field model_custom_emoji_id: %w", err)
		}
		e.ModelCustomEmojiID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode emojiStatusTypeUpgradedGift#ce0e53f4: field symbol_custom_emoji_id: %w", err)
		}
		e.SymbolCustomEmojiID = value
	}
	{
		if err := e.BackdropColors.Decode(b); err != nil {
			return fmt.Errorf("unable to decode emojiStatusTypeUpgradedGift#ce0e53f4: field backdrop_colors: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EmojiStatusTypeUpgradedGift) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiStatusTypeUpgradedGift#ce0e53f4 as nil")
	}
	b.ObjStart()
	b.PutID("emojiStatusTypeUpgradedGift")
	b.Comma()
	b.FieldStart("upgraded_gift_id")
	b.PutLong(e.UpgradedGiftID)
	b.Comma()
	b.FieldStart("gift_title")
	b.PutString(e.GiftTitle)
	b.Comma()
	b.FieldStart("gift_name")
	b.PutString(e.GiftName)
	b.Comma()
	b.FieldStart("model_custom_emoji_id")
	b.PutLong(e.ModelCustomEmojiID)
	b.Comma()
	b.FieldStart("symbol_custom_emoji_id")
	b.PutLong(e.SymbolCustomEmojiID)
	b.Comma()
	b.FieldStart("backdrop_colors")
	if err := e.BackdropColors.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode emojiStatusTypeUpgradedGift#ce0e53f4: field backdrop_colors: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EmojiStatusTypeUpgradedGift) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiStatusTypeUpgradedGift#ce0e53f4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("emojiStatusTypeUpgradedGift"); err != nil {
				return fmt.Errorf("unable to decode emojiStatusTypeUpgradedGift#ce0e53f4: %w", err)
			}
		case "upgraded_gift_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode emojiStatusTypeUpgradedGift#ce0e53f4: field upgraded_gift_id: %w", err)
			}
			e.UpgradedGiftID = value
		case "gift_title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode emojiStatusTypeUpgradedGift#ce0e53f4: field gift_title: %w", err)
			}
			e.GiftTitle = value
		case "gift_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode emojiStatusTypeUpgradedGift#ce0e53f4: field gift_name: %w", err)
			}
			e.GiftName = value
		case "model_custom_emoji_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode emojiStatusTypeUpgradedGift#ce0e53f4: field model_custom_emoji_id: %w", err)
			}
			e.ModelCustomEmojiID = value
		case "symbol_custom_emoji_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode emojiStatusTypeUpgradedGift#ce0e53f4: field symbol_custom_emoji_id: %w", err)
			}
			e.SymbolCustomEmojiID = value
		case "backdrop_colors":
			if err := e.BackdropColors.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode emojiStatusTypeUpgradedGift#ce0e53f4: field backdrop_colors: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUpgradedGiftID returns value of UpgradedGiftID field.
func (e *EmojiStatusTypeUpgradedGift) GetUpgradedGiftID() (value int64) {
	if e == nil {
		return
	}
	return e.UpgradedGiftID
}

// GetGiftTitle returns value of GiftTitle field.
func (e *EmojiStatusTypeUpgradedGift) GetGiftTitle() (value string) {
	if e == nil {
		return
	}
	return e.GiftTitle
}

// GetGiftName returns value of GiftName field.
func (e *EmojiStatusTypeUpgradedGift) GetGiftName() (value string) {
	if e == nil {
		return
	}
	return e.GiftName
}

// GetModelCustomEmojiID returns value of ModelCustomEmojiID field.
func (e *EmojiStatusTypeUpgradedGift) GetModelCustomEmojiID() (value int64) {
	if e == nil {
		return
	}
	return e.ModelCustomEmojiID
}

// GetSymbolCustomEmojiID returns value of SymbolCustomEmojiID field.
func (e *EmojiStatusTypeUpgradedGift) GetSymbolCustomEmojiID() (value int64) {
	if e == nil {
		return
	}
	return e.SymbolCustomEmojiID
}

// GetBackdropColors returns value of BackdropColors field.
func (e *EmojiStatusTypeUpgradedGift) GetBackdropColors() (value UpgradedGiftBackdropColors) {
	if e == nil {
		return
	}
	return e.BackdropColors
}

// EmojiStatusTypeClassName is schema name of EmojiStatusTypeClass.
const EmojiStatusTypeClassName = "EmojiStatusType"

// EmojiStatusTypeClass represents EmojiStatusType generic type.
//
// Example:
//
//	g, err := tdapi.DecodeEmojiStatusType(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.EmojiStatusTypeCustomEmoji: // emojiStatusTypeCustomEmoji#9ca6f0f5
//	case *tdapi.EmojiStatusTypeUpgradedGift: // emojiStatusTypeUpgradedGift#ce0e53f4
//	default: panic(v)
//	}
type EmojiStatusTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() EmojiStatusTypeClass

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

// DecodeEmojiStatusType implements binary de-serialization for EmojiStatusTypeClass.
func DecodeEmojiStatusType(buf *bin.Buffer) (EmojiStatusTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case EmojiStatusTypeCustomEmojiTypeID:
		// Decoding emojiStatusTypeCustomEmoji#9ca6f0f5.
		v := EmojiStatusTypeCustomEmoji{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmojiStatusTypeClass: %w", err)
		}
		return &v, nil
	case EmojiStatusTypeUpgradedGiftTypeID:
		// Decoding emojiStatusTypeUpgradedGift#ce0e53f4.
		v := EmojiStatusTypeUpgradedGift{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmojiStatusTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EmojiStatusTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONEmojiStatusType implements binary de-serialization for EmojiStatusTypeClass.
func DecodeTDLibJSONEmojiStatusType(buf tdjson.Decoder) (EmojiStatusTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "emojiStatusTypeCustomEmoji":
		// Decoding emojiStatusTypeCustomEmoji#9ca6f0f5.
		v := EmojiStatusTypeCustomEmoji{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmojiStatusTypeClass: %w", err)
		}
		return &v, nil
	case "emojiStatusTypeUpgradedGift":
		// Decoding emojiStatusTypeUpgradedGift#ce0e53f4.
		v := EmojiStatusTypeUpgradedGift{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmojiStatusTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EmojiStatusTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// EmojiStatusType boxes the EmojiStatusTypeClass providing a helper.
type EmojiStatusTypeBox struct {
	EmojiStatusType EmojiStatusTypeClass
}

// Decode implements bin.Decoder for EmojiStatusTypeBox.
func (b *EmojiStatusTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode EmojiStatusTypeBox to nil")
	}
	v, err := DecodeEmojiStatusType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EmojiStatusType = v
	return nil
}

// Encode implements bin.Encode for EmojiStatusTypeBox.
func (b *EmojiStatusTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.EmojiStatusType == nil {
		return fmt.Errorf("unable to encode EmojiStatusTypeClass as nil")
	}
	return b.EmojiStatusType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for EmojiStatusTypeBox.
func (b *EmojiStatusTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode EmojiStatusTypeBox to nil")
	}
	v, err := DecodeTDLibJSONEmojiStatusType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EmojiStatusType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for EmojiStatusTypeBox.
func (b *EmojiStatusTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.EmojiStatusType == nil {
		return fmt.Errorf("unable to encode EmojiStatusTypeClass as nil")
	}
	return b.EmojiStatusType.EncodeTDLibJSON(buf)
}
