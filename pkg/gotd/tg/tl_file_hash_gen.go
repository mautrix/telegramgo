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

// FileHash represents TL type `fileHash#f39b035c`.
//
// See https://core.telegram.org/constructor/fileHash for reference.
type FileHash struct {
	// Offset field of FileHash.
	Offset int64
	// Limit field of FileHash.
	Limit int
	// Hash field of FileHash.
	Hash []byte
}

// FileHashTypeID is TL type id of FileHash.
const FileHashTypeID = 0xf39b035c

// Ensuring interfaces in compile-time for FileHash.
var (
	_ bin.Encoder     = &FileHash{}
	_ bin.Decoder     = &FileHash{}
	_ bin.BareEncoder = &FileHash{}
	_ bin.BareDecoder = &FileHash{}
)

func (f *FileHash) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Offset == 0) {
		return false
	}
	if !(f.Limit == 0) {
		return false
	}
	if !(f.Hash == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FileHash) String() string {
	if f == nil {
		return "FileHash(nil)"
	}
	type Alias FileHash
	return fmt.Sprintf("FileHash%+v", Alias(*f))
}

// FillFrom fills FileHash from given interface.
func (f *FileHash) FillFrom(from interface {
	GetOffset() (value int64)
	GetLimit() (value int)
	GetHash() (value []byte)
}) {
	f.Offset = from.GetOffset()
	f.Limit = from.GetLimit()
	f.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FileHash) TypeID() uint32 {
	return FileHashTypeID
}

// TypeName returns name of type in TL schema.
func (*FileHash) TypeName() string {
	return "fileHash"
}

// TypeInfo returns info about TL type.
func (f *FileHash) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "fileHash",
		ID:   FileHashTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *FileHash) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fileHash#f39b035c as nil")
	}
	b.PutID(FileHashTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FileHash) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fileHash#f39b035c as nil")
	}
	b.PutLong(f.Offset)
	b.PutInt(f.Limit)
	b.PutBytes(f.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (f *FileHash) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fileHash#f39b035c to nil")
	}
	if err := b.ConsumeID(FileHashTypeID); err != nil {
		return fmt.Errorf("unable to decode fileHash#f39b035c: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FileHash) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fileHash#f39b035c to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode fileHash#f39b035c: field offset: %w", err)
		}
		f.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode fileHash#f39b035c: field limit: %w", err)
		}
		f.Limit = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode fileHash#f39b035c: field hash: %w", err)
		}
		f.Hash = value
	}
	return nil
}

// GetOffset returns value of Offset field.
func (f *FileHash) GetOffset() (value int64) {
	if f == nil {
		return
	}
	return f.Offset
}

// GetLimit returns value of Limit field.
func (f *FileHash) GetLimit() (value int) {
	if f == nil {
		return
	}
	return f.Limit
}

// GetHash returns value of Hash field.
func (f *FileHash) GetHash() (value []byte) {
	if f == nil {
		return
	}
	return f.Hash
}
