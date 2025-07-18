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

// EmojiKeywordsDifference represents TL type `emojiKeywordsDifference#5cc761bd`.
// Changes to emoji keywords
//
// See https://core.telegram.org/constructor/emojiKeywordsDifference for reference.
type EmojiKeywordsDifference struct {
	// Language code for keywords
	LangCode string
	// Previous emoji keyword list version
	FromVersion int
	// Current version of emoji keyword list
	Version int
	// Emojis associated to keywords
	Keywords []EmojiKeywordClass
}

// EmojiKeywordsDifferenceTypeID is TL type id of EmojiKeywordsDifference.
const EmojiKeywordsDifferenceTypeID = 0x5cc761bd

// Ensuring interfaces in compile-time for EmojiKeywordsDifference.
var (
	_ bin.Encoder     = &EmojiKeywordsDifference{}
	_ bin.Decoder     = &EmojiKeywordsDifference{}
	_ bin.BareEncoder = &EmojiKeywordsDifference{}
	_ bin.BareDecoder = &EmojiKeywordsDifference{}
)

func (e *EmojiKeywordsDifference) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.LangCode == "") {
		return false
	}
	if !(e.FromVersion == 0) {
		return false
	}
	if !(e.Version == 0) {
		return false
	}
	if !(e.Keywords == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmojiKeywordsDifference) String() string {
	if e == nil {
		return "EmojiKeywordsDifference(nil)"
	}
	type Alias EmojiKeywordsDifference
	return fmt.Sprintf("EmojiKeywordsDifference%+v", Alias(*e))
}

// FillFrom fills EmojiKeywordsDifference from given interface.
func (e *EmojiKeywordsDifference) FillFrom(from interface {
	GetLangCode() (value string)
	GetFromVersion() (value int)
	GetVersion() (value int)
	GetKeywords() (value []EmojiKeywordClass)
}) {
	e.LangCode = from.GetLangCode()
	e.FromVersion = from.GetFromVersion()
	e.Version = from.GetVersion()
	e.Keywords = from.GetKeywords()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmojiKeywordsDifference) TypeID() uint32 {
	return EmojiKeywordsDifferenceTypeID
}

// TypeName returns name of type in TL schema.
func (*EmojiKeywordsDifference) TypeName() string {
	return "emojiKeywordsDifference"
}

// TypeInfo returns info about TL type.
func (e *EmojiKeywordsDifference) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emojiKeywordsDifference",
		ID:   EmojiKeywordsDifferenceTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
		{
			Name:       "FromVersion",
			SchemaName: "from_version",
		},
		{
			Name:       "Version",
			SchemaName: "version",
		},
		{
			Name:       "Keywords",
			SchemaName: "keywords",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EmojiKeywordsDifference) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiKeywordsDifference#5cc761bd as nil")
	}
	b.PutID(EmojiKeywordsDifferenceTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EmojiKeywordsDifference) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiKeywordsDifference#5cc761bd as nil")
	}
	b.PutString(e.LangCode)
	b.PutInt(e.FromVersion)
	b.PutInt(e.Version)
	b.PutVectorHeader(len(e.Keywords))
	for idx, v := range e.Keywords {
		if v == nil {
			return fmt.Errorf("unable to encode emojiKeywordsDifference#5cc761bd: field keywords element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode emojiKeywordsDifference#5cc761bd: field keywords element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EmojiKeywordsDifference) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiKeywordsDifference#5cc761bd to nil")
	}
	if err := b.ConsumeID(EmojiKeywordsDifferenceTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiKeywordsDifference#5cc761bd: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EmojiKeywordsDifference) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiKeywordsDifference#5cc761bd to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeywordsDifference#5cc761bd: field lang_code: %w", err)
		}
		e.LangCode = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeywordsDifference#5cc761bd: field from_version: %w", err)
		}
		e.FromVersion = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeywordsDifference#5cc761bd: field version: %w", err)
		}
		e.Version = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeywordsDifference#5cc761bd: field keywords: %w", err)
		}

		if headerLen > 0 {
			e.Keywords = make([]EmojiKeywordClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeEmojiKeyword(b)
			if err != nil {
				return fmt.Errorf("unable to decode emojiKeywordsDifference#5cc761bd: field keywords: %w", err)
			}
			e.Keywords = append(e.Keywords, value)
		}
	}
	return nil
}

// GetLangCode returns value of LangCode field.
func (e *EmojiKeywordsDifference) GetLangCode() (value string) {
	if e == nil {
		return
	}
	return e.LangCode
}

// GetFromVersion returns value of FromVersion field.
func (e *EmojiKeywordsDifference) GetFromVersion() (value int) {
	if e == nil {
		return
	}
	return e.FromVersion
}

// GetVersion returns value of Version field.
func (e *EmojiKeywordsDifference) GetVersion() (value int) {
	if e == nil {
		return
	}
	return e.Version
}

// GetKeywords returns value of Keywords field.
func (e *EmojiKeywordsDifference) GetKeywords() (value []EmojiKeywordClass) {
	if e == nil {
		return
	}
	return e.Keywords
}

// MapKeywords returns field Keywords wrapped in EmojiKeywordClassArray helper.
func (e *EmojiKeywordsDifference) MapKeywords() (value EmojiKeywordClassArray) {
	return EmojiKeywordClassArray(e.Keywords)
}
