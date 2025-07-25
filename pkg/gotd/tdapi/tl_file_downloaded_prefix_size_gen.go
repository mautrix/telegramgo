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

// FileDownloadedPrefixSize represents TL type `fileDownloadedPrefixSize#87e267fb`.
type FileDownloadedPrefixSize struct {
	// The prefix size, in bytes
	Size int64
}

// FileDownloadedPrefixSizeTypeID is TL type id of FileDownloadedPrefixSize.
const FileDownloadedPrefixSizeTypeID = 0x87e267fb

// Ensuring interfaces in compile-time for FileDownloadedPrefixSize.
var (
	_ bin.Encoder     = &FileDownloadedPrefixSize{}
	_ bin.Decoder     = &FileDownloadedPrefixSize{}
	_ bin.BareEncoder = &FileDownloadedPrefixSize{}
	_ bin.BareDecoder = &FileDownloadedPrefixSize{}
)

func (f *FileDownloadedPrefixSize) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Size == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FileDownloadedPrefixSize) String() string {
	if f == nil {
		return "FileDownloadedPrefixSize(nil)"
	}
	type Alias FileDownloadedPrefixSize
	return fmt.Sprintf("FileDownloadedPrefixSize%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FileDownloadedPrefixSize) TypeID() uint32 {
	return FileDownloadedPrefixSizeTypeID
}

// TypeName returns name of type in TL schema.
func (*FileDownloadedPrefixSize) TypeName() string {
	return "fileDownloadedPrefixSize"
}

// TypeInfo returns info about TL type.
func (f *FileDownloadedPrefixSize) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "fileDownloadedPrefixSize",
		ID:   FileDownloadedPrefixSizeTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Size",
			SchemaName: "size",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *FileDownloadedPrefixSize) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fileDownloadedPrefixSize#87e267fb as nil")
	}
	b.PutID(FileDownloadedPrefixSizeTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FileDownloadedPrefixSize) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fileDownloadedPrefixSize#87e267fb as nil")
	}
	b.PutInt53(f.Size)
	return nil
}

// Decode implements bin.Decoder.
func (f *FileDownloadedPrefixSize) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fileDownloadedPrefixSize#87e267fb to nil")
	}
	if err := b.ConsumeID(FileDownloadedPrefixSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode fileDownloadedPrefixSize#87e267fb: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FileDownloadedPrefixSize) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fileDownloadedPrefixSize#87e267fb to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode fileDownloadedPrefixSize#87e267fb: field size: %w", err)
		}
		f.Size = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *FileDownloadedPrefixSize) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode fileDownloadedPrefixSize#87e267fb as nil")
	}
	b.ObjStart()
	b.PutID("fileDownloadedPrefixSize")
	b.Comma()
	b.FieldStart("size")
	b.PutInt53(f.Size)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (f *FileDownloadedPrefixSize) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode fileDownloadedPrefixSize#87e267fb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("fileDownloadedPrefixSize"); err != nil {
				return fmt.Errorf("unable to decode fileDownloadedPrefixSize#87e267fb: %w", err)
			}
		case "size":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode fileDownloadedPrefixSize#87e267fb: field size: %w", err)
			}
			f.Size = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSize returns value of Size field.
func (f *FileDownloadedPrefixSize) GetSize() (value int64) {
	if f == nil {
		return
	}
	return f.Size
}
