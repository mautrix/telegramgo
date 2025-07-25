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

// StorageStatisticsFast represents TL type `storageStatisticsFast#cb412861`.
type StorageStatisticsFast struct {
	// Approximate total size of files, in bytes
	FilesSize int64
	// Approximate number of files
	FileCount int32
	// Size of the database
	DatabaseSize int64
	// Size of the language pack database
	LanguagePackDatabaseSize int64
	// Size of the TDLib internal log
	LogSize int64
}

// StorageStatisticsFastTypeID is TL type id of StorageStatisticsFast.
const StorageStatisticsFastTypeID = 0xcb412861

// Ensuring interfaces in compile-time for StorageStatisticsFast.
var (
	_ bin.Encoder     = &StorageStatisticsFast{}
	_ bin.Decoder     = &StorageStatisticsFast{}
	_ bin.BareEncoder = &StorageStatisticsFast{}
	_ bin.BareDecoder = &StorageStatisticsFast{}
)

func (s *StorageStatisticsFast) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.FilesSize == 0) {
		return false
	}
	if !(s.FileCount == 0) {
		return false
	}
	if !(s.DatabaseSize == 0) {
		return false
	}
	if !(s.LanguagePackDatabaseSize == 0) {
		return false
	}
	if !(s.LogSize == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StorageStatisticsFast) String() string {
	if s == nil {
		return "StorageStatisticsFast(nil)"
	}
	type Alias StorageStatisticsFast
	return fmt.Sprintf("StorageStatisticsFast%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StorageStatisticsFast) TypeID() uint32 {
	return StorageStatisticsFastTypeID
}

// TypeName returns name of type in TL schema.
func (*StorageStatisticsFast) TypeName() string {
	return "storageStatisticsFast"
}

// TypeInfo returns info about TL type.
func (s *StorageStatisticsFast) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storageStatisticsFast",
		ID:   StorageStatisticsFastTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FilesSize",
			SchemaName: "files_size",
		},
		{
			Name:       "FileCount",
			SchemaName: "file_count",
		},
		{
			Name:       "DatabaseSize",
			SchemaName: "database_size",
		},
		{
			Name:       "LanguagePackDatabaseSize",
			SchemaName: "language_pack_database_size",
		},
		{
			Name:       "LogSize",
			SchemaName: "log_size",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StorageStatisticsFast) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storageStatisticsFast#cb412861 as nil")
	}
	b.PutID(StorageStatisticsFastTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StorageStatisticsFast) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storageStatisticsFast#cb412861 as nil")
	}
	b.PutInt53(s.FilesSize)
	b.PutInt32(s.FileCount)
	b.PutInt53(s.DatabaseSize)
	b.PutInt53(s.LanguagePackDatabaseSize)
	b.PutInt53(s.LogSize)
	return nil
}

// Decode implements bin.Decoder.
func (s *StorageStatisticsFast) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storageStatisticsFast#cb412861 to nil")
	}
	if err := b.ConsumeID(StorageStatisticsFastTypeID); err != nil {
		return fmt.Errorf("unable to decode storageStatisticsFast#cb412861: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StorageStatisticsFast) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storageStatisticsFast#cb412861 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode storageStatisticsFast#cb412861: field files_size: %w", err)
		}
		s.FilesSize = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode storageStatisticsFast#cb412861: field file_count: %w", err)
		}
		s.FileCount = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode storageStatisticsFast#cb412861: field database_size: %w", err)
		}
		s.DatabaseSize = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode storageStatisticsFast#cb412861: field language_pack_database_size: %w", err)
		}
		s.LanguagePackDatabaseSize = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode storageStatisticsFast#cb412861: field log_size: %w", err)
		}
		s.LogSize = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StorageStatisticsFast) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storageStatisticsFast#cb412861 as nil")
	}
	b.ObjStart()
	b.PutID("storageStatisticsFast")
	b.Comma()
	b.FieldStart("files_size")
	b.PutInt53(s.FilesSize)
	b.Comma()
	b.FieldStart("file_count")
	b.PutInt32(s.FileCount)
	b.Comma()
	b.FieldStart("database_size")
	b.PutInt53(s.DatabaseSize)
	b.Comma()
	b.FieldStart("language_pack_database_size")
	b.PutInt53(s.LanguagePackDatabaseSize)
	b.Comma()
	b.FieldStart("log_size")
	b.PutInt53(s.LogSize)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StorageStatisticsFast) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storageStatisticsFast#cb412861 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storageStatisticsFast"); err != nil {
				return fmt.Errorf("unable to decode storageStatisticsFast#cb412861: %w", err)
			}
		case "files_size":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode storageStatisticsFast#cb412861: field files_size: %w", err)
			}
			s.FilesSize = value
		case "file_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode storageStatisticsFast#cb412861: field file_count: %w", err)
			}
			s.FileCount = value
		case "database_size":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode storageStatisticsFast#cb412861: field database_size: %w", err)
			}
			s.DatabaseSize = value
		case "language_pack_database_size":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode storageStatisticsFast#cb412861: field language_pack_database_size: %w", err)
			}
			s.LanguagePackDatabaseSize = value
		case "log_size":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode storageStatisticsFast#cb412861: field log_size: %w", err)
			}
			s.LogSize = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFilesSize returns value of FilesSize field.
func (s *StorageStatisticsFast) GetFilesSize() (value int64) {
	if s == nil {
		return
	}
	return s.FilesSize
}

// GetFileCount returns value of FileCount field.
func (s *StorageStatisticsFast) GetFileCount() (value int32) {
	if s == nil {
		return
	}
	return s.FileCount
}

// GetDatabaseSize returns value of DatabaseSize field.
func (s *StorageStatisticsFast) GetDatabaseSize() (value int64) {
	if s == nil {
		return
	}
	return s.DatabaseSize
}

// GetLanguagePackDatabaseSize returns value of LanguagePackDatabaseSize field.
func (s *StorageStatisticsFast) GetLanguagePackDatabaseSize() (value int64) {
	if s == nil {
		return
	}
	return s.LanguagePackDatabaseSize
}

// GetLogSize returns value of LogSize field.
func (s *StorageStatisticsFast) GetLogSize() (value int64) {
	if s == nil {
		return
	}
	return s.LogSize
}
