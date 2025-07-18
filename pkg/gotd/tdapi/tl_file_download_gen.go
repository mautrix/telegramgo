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

// FileDownload represents TL type `fileDownload#834d1354`.
type FileDownload struct {
	// File identifier
	FileID int32
	// The message with the file
	Message Message
	// Point in time (Unix timestamp) when the file was added to the download list
	AddDate int32
	// Point in time (Unix timestamp) when the file downloading was completed; 0 if the file
	// downloading isn't completed
	CompleteDate int32
	// True, if downloading of the file is paused
	IsPaused bool
}

// FileDownloadTypeID is TL type id of FileDownload.
const FileDownloadTypeID = 0x834d1354

// Ensuring interfaces in compile-time for FileDownload.
var (
	_ bin.Encoder     = &FileDownload{}
	_ bin.Decoder     = &FileDownload{}
	_ bin.BareEncoder = &FileDownload{}
	_ bin.BareDecoder = &FileDownload{}
)

func (f *FileDownload) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.FileID == 0) {
		return false
	}
	if !(f.Message.Zero()) {
		return false
	}
	if !(f.AddDate == 0) {
		return false
	}
	if !(f.CompleteDate == 0) {
		return false
	}
	if !(f.IsPaused == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FileDownload) String() string {
	if f == nil {
		return "FileDownload(nil)"
	}
	type Alias FileDownload
	return fmt.Sprintf("FileDownload%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FileDownload) TypeID() uint32 {
	return FileDownloadTypeID
}

// TypeName returns name of type in TL schema.
func (*FileDownload) TypeName() string {
	return "fileDownload"
}

// TypeInfo returns info about TL type.
func (f *FileDownload) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "fileDownload",
		ID:   FileDownloadTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileID",
			SchemaName: "file_id",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "AddDate",
			SchemaName: "add_date",
		},
		{
			Name:       "CompleteDate",
			SchemaName: "complete_date",
		},
		{
			Name:       "IsPaused",
			SchemaName: "is_paused",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *FileDownload) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fileDownload#834d1354 as nil")
	}
	b.PutID(FileDownloadTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FileDownload) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fileDownload#834d1354 as nil")
	}
	b.PutInt32(f.FileID)
	if err := f.Message.Encode(b); err != nil {
		return fmt.Errorf("unable to encode fileDownload#834d1354: field message: %w", err)
	}
	b.PutInt32(f.AddDate)
	b.PutInt32(f.CompleteDate)
	b.PutBool(f.IsPaused)
	return nil
}

// Decode implements bin.Decoder.
func (f *FileDownload) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fileDownload#834d1354 to nil")
	}
	if err := b.ConsumeID(FileDownloadTypeID); err != nil {
		return fmt.Errorf("unable to decode fileDownload#834d1354: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FileDownload) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fileDownload#834d1354 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode fileDownload#834d1354: field file_id: %w", err)
		}
		f.FileID = value
	}
	{
		if err := f.Message.Decode(b); err != nil {
			return fmt.Errorf("unable to decode fileDownload#834d1354: field message: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode fileDownload#834d1354: field add_date: %w", err)
		}
		f.AddDate = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode fileDownload#834d1354: field complete_date: %w", err)
		}
		f.CompleteDate = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode fileDownload#834d1354: field is_paused: %w", err)
		}
		f.IsPaused = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *FileDownload) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode fileDownload#834d1354 as nil")
	}
	b.ObjStart()
	b.PutID("fileDownload")
	b.Comma()
	b.FieldStart("file_id")
	b.PutInt32(f.FileID)
	b.Comma()
	b.FieldStart("message")
	if err := f.Message.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode fileDownload#834d1354: field message: %w", err)
	}
	b.Comma()
	b.FieldStart("add_date")
	b.PutInt32(f.AddDate)
	b.Comma()
	b.FieldStart("complete_date")
	b.PutInt32(f.CompleteDate)
	b.Comma()
	b.FieldStart("is_paused")
	b.PutBool(f.IsPaused)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (f *FileDownload) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode fileDownload#834d1354 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("fileDownload"); err != nil {
				return fmt.Errorf("unable to decode fileDownload#834d1354: %w", err)
			}
		case "file_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode fileDownload#834d1354: field file_id: %w", err)
			}
			f.FileID = value
		case "message":
			if err := f.Message.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode fileDownload#834d1354: field message: %w", err)
			}
		case "add_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode fileDownload#834d1354: field add_date: %w", err)
			}
			f.AddDate = value
		case "complete_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode fileDownload#834d1354: field complete_date: %w", err)
			}
			f.CompleteDate = value
		case "is_paused":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode fileDownload#834d1354: field is_paused: %w", err)
			}
			f.IsPaused = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFileID returns value of FileID field.
func (f *FileDownload) GetFileID() (value int32) {
	if f == nil {
		return
	}
	return f.FileID
}

// GetMessage returns value of Message field.
func (f *FileDownload) GetMessage() (value Message) {
	if f == nil {
		return
	}
	return f.Message
}

// GetAddDate returns value of AddDate field.
func (f *FileDownload) GetAddDate() (value int32) {
	if f == nil {
		return
	}
	return f.AddDate
}

// GetCompleteDate returns value of CompleteDate field.
func (f *FileDownload) GetCompleteDate() (value int32) {
	if f == nil {
		return
	}
	return f.CompleteDate
}

// GetIsPaused returns value of IsPaused field.
func (f *FileDownload) GetIsPaused() (value bool) {
	if f == nil {
		return
	}
	return f.IsPaused
}
