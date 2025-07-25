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

// EncryptedFileEmpty represents TL type `encryptedFileEmpty#c21f497e`.
// Empty constructor, non-existing file.
//
// See https://core.telegram.org/constructor/encryptedFileEmpty for reference.
type EncryptedFileEmpty struct {
}

// EncryptedFileEmptyTypeID is TL type id of EncryptedFileEmpty.
const EncryptedFileEmptyTypeID = 0xc21f497e

// construct implements constructor of EncryptedFileClass.
func (e EncryptedFileEmpty) construct() EncryptedFileClass { return &e }

// Ensuring interfaces in compile-time for EncryptedFileEmpty.
var (
	_ bin.Encoder     = &EncryptedFileEmpty{}
	_ bin.Decoder     = &EncryptedFileEmpty{}
	_ bin.BareEncoder = &EncryptedFileEmpty{}
	_ bin.BareDecoder = &EncryptedFileEmpty{}

	_ EncryptedFileClass = &EncryptedFileEmpty{}
)

func (e *EncryptedFileEmpty) Zero() bool {
	if e == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedFileEmpty) String() string {
	if e == nil {
		return "EncryptedFileEmpty(nil)"
	}
	type Alias EncryptedFileEmpty
	return fmt.Sprintf("EncryptedFileEmpty%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EncryptedFileEmpty) TypeID() uint32 {
	return EncryptedFileEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*EncryptedFileEmpty) TypeName() string {
	return "encryptedFileEmpty"
}

// TypeInfo returns info about TL type.
func (e *EncryptedFileEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "encryptedFileEmpty",
		ID:   EncryptedFileEmptyTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (e *EncryptedFileEmpty) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedFileEmpty#c21f497e as nil")
	}
	b.PutID(EncryptedFileEmptyTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EncryptedFileEmpty) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedFileEmpty#c21f497e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedFileEmpty) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedFileEmpty#c21f497e to nil")
	}
	if err := b.ConsumeID(EncryptedFileEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedFileEmpty#c21f497e: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EncryptedFileEmpty) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedFileEmpty#c21f497e to nil")
	}
	return nil
}

// EncryptedFile represents TL type `encryptedFile#a8008cd8`.
// Encrypted file.
//
// See https://core.telegram.org/constructor/encryptedFile for reference.
type EncryptedFile struct {
	// File ID
	ID int64
	// Checking sum depending on user ID
	AccessHash int64
	// File size in bytes
	Size int64
	// Number of data center
	DCID int
	// 32-bit fingerprint of key used for file encryption
	KeyFingerprint int
}

// EncryptedFileTypeID is TL type id of EncryptedFile.
const EncryptedFileTypeID = 0xa8008cd8

// construct implements constructor of EncryptedFileClass.
func (e EncryptedFile) construct() EncryptedFileClass { return &e }

// Ensuring interfaces in compile-time for EncryptedFile.
var (
	_ bin.Encoder     = &EncryptedFile{}
	_ bin.Decoder     = &EncryptedFile{}
	_ bin.BareEncoder = &EncryptedFile{}
	_ bin.BareDecoder = &EncryptedFile{}

	_ EncryptedFileClass = &EncryptedFile{}
)

func (e *EncryptedFile) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ID == 0) {
		return false
	}
	if !(e.AccessHash == 0) {
		return false
	}
	if !(e.Size == 0) {
		return false
	}
	if !(e.DCID == 0) {
		return false
	}
	if !(e.KeyFingerprint == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedFile) String() string {
	if e == nil {
		return "EncryptedFile(nil)"
	}
	type Alias EncryptedFile
	return fmt.Sprintf("EncryptedFile%+v", Alias(*e))
}

// FillFrom fills EncryptedFile from given interface.
func (e *EncryptedFile) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetSize() (value int64)
	GetDCID() (value int)
	GetKeyFingerprint() (value int)
}) {
	e.ID = from.GetID()
	e.AccessHash = from.GetAccessHash()
	e.Size = from.GetSize()
	e.DCID = from.GetDCID()
	e.KeyFingerprint = from.GetKeyFingerprint()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EncryptedFile) TypeID() uint32 {
	return EncryptedFileTypeID
}

// TypeName returns name of type in TL schema.
func (*EncryptedFile) TypeName() string {
	return "encryptedFile"
}

// TypeInfo returns info about TL type.
func (e *EncryptedFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "encryptedFile",
		ID:   EncryptedFileTypeID,
	}
	if e == nil {
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
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
		{
			Name:       "KeyFingerprint",
			SchemaName: "key_fingerprint",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EncryptedFile) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedFile#a8008cd8 as nil")
	}
	b.PutID(EncryptedFileTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EncryptedFile) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedFile#a8008cd8 as nil")
	}
	b.PutLong(e.ID)
	b.PutLong(e.AccessHash)
	b.PutLong(e.Size)
	b.PutInt(e.DCID)
	b.PutInt(e.KeyFingerprint)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedFile) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedFile#a8008cd8 to nil")
	}
	if err := b.ConsumeID(EncryptedFileTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedFile#a8008cd8: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EncryptedFile) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedFile#a8008cd8 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#a8008cd8: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#a8008cd8: field access_hash: %w", err)
		}
		e.AccessHash = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#a8008cd8: field size: %w", err)
		}
		e.Size = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#a8008cd8: field dc_id: %w", err)
		}
		e.DCID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#a8008cd8: field key_fingerprint: %w", err)
		}
		e.KeyFingerprint = value
	}
	return nil
}

// GetID returns value of ID field.
func (e *EncryptedFile) GetID() (value int64) {
	if e == nil {
		return
	}
	return e.ID
}

// GetAccessHash returns value of AccessHash field.
func (e *EncryptedFile) GetAccessHash() (value int64) {
	if e == nil {
		return
	}
	return e.AccessHash
}

// GetSize returns value of Size field.
func (e *EncryptedFile) GetSize() (value int64) {
	if e == nil {
		return
	}
	return e.Size
}

// GetDCID returns value of DCID field.
func (e *EncryptedFile) GetDCID() (value int) {
	if e == nil {
		return
	}
	return e.DCID
}

// GetKeyFingerprint returns value of KeyFingerprint field.
func (e *EncryptedFile) GetKeyFingerprint() (value int) {
	if e == nil {
		return
	}
	return e.KeyFingerprint
}

// EncryptedFileClassName is schema name of EncryptedFileClass.
const EncryptedFileClassName = "EncryptedFile"

// EncryptedFileClass represents EncryptedFile generic type.
//
// See https://core.telegram.org/type/EncryptedFile for reference.
//
// Example:
//
//	g, err := tg.DecodeEncryptedFile(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.EncryptedFileEmpty: // encryptedFileEmpty#c21f497e
//	case *tg.EncryptedFile: // encryptedFile#a8008cd8
//	default: panic(v)
//	}
type EncryptedFileClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() EncryptedFileClass

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

	// AsNotEmpty tries to map EncryptedFileClass to EncryptedFile.
	AsNotEmpty() (*EncryptedFile, bool)
}

// AsInputEncryptedFileLocation tries to map EncryptedFile to InputEncryptedFileLocation.
func (e *EncryptedFile) AsInputEncryptedFileLocation() *InputEncryptedFileLocation {
	value := new(InputEncryptedFileLocation)
	value.ID = e.GetID()
	value.AccessHash = e.GetAccessHash()

	return value
}

// AsInput tries to map EncryptedFile to InputEncryptedFile.
func (e *EncryptedFile) AsInput() *InputEncryptedFile {
	value := new(InputEncryptedFile)
	value.ID = e.GetID()
	value.AccessHash = e.GetAccessHash()

	return value
}

// AsNotEmpty tries to map EncryptedFileEmpty to EncryptedFile.
func (e *EncryptedFileEmpty) AsNotEmpty() (*EncryptedFile, bool) {
	return nil, false
}

// AsNotEmpty tries to map EncryptedFile to EncryptedFile.
func (e *EncryptedFile) AsNotEmpty() (*EncryptedFile, bool) {
	return e, true
}

// DecodeEncryptedFile implements binary de-serialization for EncryptedFileClass.
func DecodeEncryptedFile(buf *bin.Buffer) (EncryptedFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case EncryptedFileEmptyTypeID:
		// Decoding encryptedFileEmpty#c21f497e.
		v := EncryptedFileEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedFileClass: %w", err)
		}
		return &v, nil
	case EncryptedFileTypeID:
		// Decoding encryptedFile#a8008cd8.
		v := EncryptedFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EncryptedFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// EncryptedFile boxes the EncryptedFileClass providing a helper.
type EncryptedFileBox struct {
	EncryptedFile EncryptedFileClass
}

// Decode implements bin.Decoder for EncryptedFileBox.
func (b *EncryptedFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode EncryptedFileBox to nil")
	}
	v, err := DecodeEncryptedFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EncryptedFile = v
	return nil
}

// Encode implements bin.Encode for EncryptedFileBox.
func (b *EncryptedFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.EncryptedFile == nil {
		return fmt.Errorf("unable to encode EncryptedFileClass as nil")
	}
	return b.EncryptedFile.Encode(buf)
}
