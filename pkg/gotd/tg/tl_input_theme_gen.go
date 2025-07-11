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

// InputTheme represents TL type `inputTheme#3c5693e9`.
// Theme
//
// See https://core.telegram.org/constructor/inputTheme for reference.
type InputTheme struct {
	// ID
	ID int64
	// Access hash
	AccessHash int64
}

// InputThemeTypeID is TL type id of InputTheme.
const InputThemeTypeID = 0x3c5693e9

// construct implements constructor of InputThemeClass.
func (i InputTheme) construct() InputThemeClass { return &i }

// Ensuring interfaces in compile-time for InputTheme.
var (
	_ bin.Encoder     = &InputTheme{}
	_ bin.Decoder     = &InputTheme{}
	_ bin.BareEncoder = &InputTheme{}
	_ bin.BareDecoder = &InputTheme{}

	_ InputThemeClass = &InputTheme{}
)

func (i *InputTheme) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputTheme) String() string {
	if i == nil {
		return "InputTheme(nil)"
	}
	type Alias InputTheme
	return fmt.Sprintf("InputTheme%+v", Alias(*i))
}

// FillFrom fills InputTheme from given interface.
func (i *InputTheme) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
}) {
	i.ID = from.GetID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputTheme) TypeID() uint32 {
	return InputThemeTypeID
}

// TypeName returns name of type in TL schema.
func (*InputTheme) TypeName() string {
	return "inputTheme"
}

// TypeInfo returns info about TL type.
func (i *InputTheme) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputTheme",
		ID:   InputThemeTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputTheme) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputTheme#3c5693e9 as nil")
	}
	b.PutID(InputThemeTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputTheme) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputTheme#3c5693e9 as nil")
	}
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputTheme) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputTheme#3c5693e9 to nil")
	}
	if err := b.ConsumeID(InputThemeTypeID); err != nil {
		return fmt.Errorf("unable to decode inputTheme#3c5693e9: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputTheme) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputTheme#3c5693e9 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputTheme#3c5693e9: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputTheme#3c5693e9: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputTheme) GetID() (value int64) {
	if i == nil {
		return
	}
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputTheme) GetAccessHash() (value int64) {
	if i == nil {
		return
	}
	return i.AccessHash
}

// InputThemeSlug represents TL type `inputThemeSlug#f5890df1`.
// Theme by theme ID
//
// See https://core.telegram.org/constructor/inputThemeSlug for reference.
type InputThemeSlug struct {
	// Unique theme ID obtained from a theme deep link »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/links#theme-links
	Slug string
}

// InputThemeSlugTypeID is TL type id of InputThemeSlug.
const InputThemeSlugTypeID = 0xf5890df1

// construct implements constructor of InputThemeClass.
func (i InputThemeSlug) construct() InputThemeClass { return &i }

// Ensuring interfaces in compile-time for InputThemeSlug.
var (
	_ bin.Encoder     = &InputThemeSlug{}
	_ bin.Decoder     = &InputThemeSlug{}
	_ bin.BareEncoder = &InputThemeSlug{}
	_ bin.BareDecoder = &InputThemeSlug{}

	_ InputThemeClass = &InputThemeSlug{}
)

func (i *InputThemeSlug) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Slug == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputThemeSlug) String() string {
	if i == nil {
		return "InputThemeSlug(nil)"
	}
	type Alias InputThemeSlug
	return fmt.Sprintf("InputThemeSlug%+v", Alias(*i))
}

// FillFrom fills InputThemeSlug from given interface.
func (i *InputThemeSlug) FillFrom(from interface {
	GetSlug() (value string)
}) {
	i.Slug = from.GetSlug()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputThemeSlug) TypeID() uint32 {
	return InputThemeSlugTypeID
}

// TypeName returns name of type in TL schema.
func (*InputThemeSlug) TypeName() string {
	return "inputThemeSlug"
}

// TypeInfo returns info about TL type.
func (i *InputThemeSlug) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputThemeSlug",
		ID:   InputThemeSlugTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Slug",
			SchemaName: "slug",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputThemeSlug) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputThemeSlug#f5890df1 as nil")
	}
	b.PutID(InputThemeSlugTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputThemeSlug) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputThemeSlug#f5890df1 as nil")
	}
	b.PutString(i.Slug)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputThemeSlug) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputThemeSlug#f5890df1 to nil")
	}
	if err := b.ConsumeID(InputThemeSlugTypeID); err != nil {
		return fmt.Errorf("unable to decode inputThemeSlug#f5890df1: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputThemeSlug) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputThemeSlug#f5890df1 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputThemeSlug#f5890df1: field slug: %w", err)
		}
		i.Slug = value
	}
	return nil
}

// GetSlug returns value of Slug field.
func (i *InputThemeSlug) GetSlug() (value string) {
	if i == nil {
		return
	}
	return i.Slug
}

// InputThemeClassName is schema name of InputThemeClass.
const InputThemeClassName = "InputTheme"

// InputThemeClass represents InputTheme generic type.
//
// See https://core.telegram.org/type/InputTheme for reference.
//
// Example:
//
//	g, err := tg.DecodeInputTheme(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.InputTheme: // inputTheme#3c5693e9
//	case *tg.InputThemeSlug: // inputThemeSlug#f5890df1
//	default: panic(v)
//	}
type InputThemeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputThemeClass

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
}

// DecodeInputTheme implements binary de-serialization for InputThemeClass.
func DecodeInputTheme(buf *bin.Buffer) (InputThemeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputThemeTypeID:
		// Decoding inputTheme#3c5693e9.
		v := InputTheme{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputThemeClass: %w", err)
		}
		return &v, nil
	case InputThemeSlugTypeID:
		// Decoding inputThemeSlug#f5890df1.
		v := InputThemeSlug{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputThemeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputThemeClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputTheme boxes the InputThemeClass providing a helper.
type InputThemeBox struct {
	InputTheme InputThemeClass
}

// Decode implements bin.Decoder for InputThemeBox.
func (b *InputThemeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputThemeBox to nil")
	}
	v, err := DecodeInputTheme(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputTheme = v
	return nil
}

// Encode implements bin.Encode for InputThemeBox.
func (b *InputThemeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputTheme == nil {
		return fmt.Errorf("unable to encode InputThemeClass as nil")
	}
	return b.InputTheme.Encode(buf)
}
