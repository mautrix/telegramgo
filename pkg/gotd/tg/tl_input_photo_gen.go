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

// InputPhotoEmpty represents TL type `inputPhotoEmpty#1cd7bf0d`.
// Empty constructor.
//
// See https://core.telegram.org/constructor/inputPhotoEmpty for reference.
type InputPhotoEmpty struct {
}

// InputPhotoEmptyTypeID is TL type id of InputPhotoEmpty.
const InputPhotoEmptyTypeID = 0x1cd7bf0d

// construct implements constructor of InputPhotoClass.
func (i InputPhotoEmpty) construct() InputPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputPhotoEmpty.
var (
	_ bin.Encoder     = &InputPhotoEmpty{}
	_ bin.Decoder     = &InputPhotoEmpty{}
	_ bin.BareEncoder = &InputPhotoEmpty{}
	_ bin.BareDecoder = &InputPhotoEmpty{}

	_ InputPhotoClass = &InputPhotoEmpty{}
)

func (i *InputPhotoEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPhotoEmpty) String() string {
	if i == nil {
		return "InputPhotoEmpty(nil)"
	}
	type Alias InputPhotoEmpty
	return fmt.Sprintf("InputPhotoEmpty%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPhotoEmpty) TypeID() uint32 {
	return InputPhotoEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPhotoEmpty) TypeName() string {
	return "inputPhotoEmpty"
}

// TypeInfo returns info about TL type.
func (i *InputPhotoEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPhotoEmpty",
		ID:   InputPhotoEmptyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPhotoEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhotoEmpty#1cd7bf0d as nil")
	}
	b.PutID(InputPhotoEmptyTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPhotoEmpty) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhotoEmpty#1cd7bf0d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPhotoEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhotoEmpty#1cd7bf0d to nil")
	}
	if err := b.ConsumeID(InputPhotoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPhotoEmpty#1cd7bf0d: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPhotoEmpty) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhotoEmpty#1cd7bf0d to nil")
	}
	return nil
}

// InputPhoto represents TL type `inputPhoto#3bb3b94a`.
// Defines a photo for further interaction.
//
// See https://core.telegram.org/constructor/inputPhoto for reference.
type InputPhoto struct {
	// Photo identifier
	ID int64
	// access_hash value from the photo¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/constructor/photo
	AccessHash int64
	// File reference¹
	//
	// Links:
	//  1) https://core.telegram.org/api/file_reference
	FileReference []byte
}

// InputPhotoTypeID is TL type id of InputPhoto.
const InputPhotoTypeID = 0x3bb3b94a

// construct implements constructor of InputPhotoClass.
func (i InputPhoto) construct() InputPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputPhoto.
var (
	_ bin.Encoder     = &InputPhoto{}
	_ bin.Decoder     = &InputPhoto{}
	_ bin.BareEncoder = &InputPhoto{}
	_ bin.BareDecoder = &InputPhoto{}

	_ InputPhotoClass = &InputPhoto{}
)

func (i *InputPhoto) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}
	if !(i.FileReference == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPhoto) String() string {
	if i == nil {
		return "InputPhoto(nil)"
	}
	type Alias InputPhoto
	return fmt.Sprintf("InputPhoto%+v", Alias(*i))
}

// FillFrom fills InputPhoto from given interface.
func (i *InputPhoto) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetFileReference() (value []byte)
}) {
	i.ID = from.GetID()
	i.AccessHash = from.GetAccessHash()
	i.FileReference = from.GetFileReference()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPhoto) TypeID() uint32 {
	return InputPhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPhoto) TypeName() string {
	return "inputPhoto"
}

// TypeInfo returns info about TL type.
func (i *InputPhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPhoto",
		ID:   InputPhotoTypeID,
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
		{
			Name:       "FileReference",
			SchemaName: "file_reference",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhoto#3bb3b94a as nil")
	}
	b.PutID(InputPhotoTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPhoto) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhoto#3bb3b94a as nil")
	}
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	b.PutBytes(i.FileReference)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhoto#3bb3b94a to nil")
	}
	if err := b.ConsumeID(InputPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPhoto#3bb3b94a: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPhoto) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhoto#3bb3b94a to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoto#3bb3b94a: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoto#3bb3b94a: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoto#3bb3b94a: field file_reference: %w", err)
		}
		i.FileReference = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputPhoto) GetID() (value int64) {
	if i == nil {
		return
	}
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputPhoto) GetAccessHash() (value int64) {
	if i == nil {
		return
	}
	return i.AccessHash
}

// GetFileReference returns value of FileReference field.
func (i *InputPhoto) GetFileReference() (value []byte) {
	if i == nil {
		return
	}
	return i.FileReference
}

// InputPhotoClassName is schema name of InputPhotoClass.
const InputPhotoClassName = "InputPhoto"

// InputPhotoClass represents InputPhoto generic type.
//
// See https://core.telegram.org/type/InputPhoto for reference.
//
// Example:
//
//	g, err := tg.DecodeInputPhoto(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.InputPhotoEmpty: // inputPhotoEmpty#1cd7bf0d
//	case *tg.InputPhoto: // inputPhoto#3bb3b94a
//	default: panic(v)
//	}
type InputPhotoClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputPhotoClass

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

	// AsNotEmpty tries to map InputPhotoClass to InputPhoto.
	AsNotEmpty() (*InputPhoto, bool)
}

// AsNotEmpty tries to map InputPhotoEmpty to InputPhoto.
func (i *InputPhotoEmpty) AsNotEmpty() (*InputPhoto, bool) {
	return nil, false
}

// AsNotEmpty tries to map InputPhoto to InputPhoto.
func (i *InputPhoto) AsNotEmpty() (*InputPhoto, bool) {
	return i, true
}

// DecodeInputPhoto implements binary de-serialization for InputPhotoClass.
func DecodeInputPhoto(buf *bin.Buffer) (InputPhotoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPhotoEmptyTypeID:
		// Decoding inputPhotoEmpty#1cd7bf0d.
		v := InputPhotoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPhotoClass: %w", err)
		}
		return &v, nil
	case InputPhotoTypeID:
		// Decoding inputPhoto#3bb3b94a.
		v := InputPhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPhotoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPhotoClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputPhoto boxes the InputPhotoClass providing a helper.
type InputPhotoBox struct {
	InputPhoto InputPhotoClass
}

// Decode implements bin.Decoder for InputPhotoBox.
func (b *InputPhotoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPhotoBox to nil")
	}
	v, err := DecodeInputPhoto(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPhoto = v
	return nil
}

// Encode implements bin.Encode for InputPhotoBox.
func (b *InputPhotoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputPhoto == nil {
		return fmt.Errorf("unable to encode InputPhotoClass as nil")
	}
	return b.InputPhoto.Encode(buf)
}
