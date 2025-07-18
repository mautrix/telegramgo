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

// PreliminaryUploadFileRequest represents TL type `preliminaryUploadFile#70e7cb99`.
type PreliminaryUploadFileRequest struct {
	// File to upload
	File InputFileClass
	// File type; pass null if unknown
	FileType FileTypeClass
	// Priority of the upload (1-32). The higher the priority, the earlier the file will be
	// uploaded. If the priorities of two files are equal, then the first one for which
	// preliminaryUploadFile was called will be uploaded first
	Priority int32
}

// PreliminaryUploadFileRequestTypeID is TL type id of PreliminaryUploadFileRequest.
const PreliminaryUploadFileRequestTypeID = 0x70e7cb99

// Ensuring interfaces in compile-time for PreliminaryUploadFileRequest.
var (
	_ bin.Encoder     = &PreliminaryUploadFileRequest{}
	_ bin.Decoder     = &PreliminaryUploadFileRequest{}
	_ bin.BareEncoder = &PreliminaryUploadFileRequest{}
	_ bin.BareDecoder = &PreliminaryUploadFileRequest{}
)

func (p *PreliminaryUploadFileRequest) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.File == nil) {
		return false
	}
	if !(p.FileType == nil) {
		return false
	}
	if !(p.Priority == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PreliminaryUploadFileRequest) String() string {
	if p == nil {
		return "PreliminaryUploadFileRequest(nil)"
	}
	type Alias PreliminaryUploadFileRequest
	return fmt.Sprintf("PreliminaryUploadFileRequest%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PreliminaryUploadFileRequest) TypeID() uint32 {
	return PreliminaryUploadFileRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PreliminaryUploadFileRequest) TypeName() string {
	return "preliminaryUploadFile"
}

// TypeInfo returns info about TL type.
func (p *PreliminaryUploadFileRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "preliminaryUploadFile",
		ID:   PreliminaryUploadFileRequestTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "File",
			SchemaName: "file",
		},
		{
			Name:       "FileType",
			SchemaName: "file_type",
		},
		{
			Name:       "Priority",
			SchemaName: "priority",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PreliminaryUploadFileRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode preliminaryUploadFile#70e7cb99 as nil")
	}
	b.PutID(PreliminaryUploadFileRequestTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PreliminaryUploadFileRequest) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode preliminaryUploadFile#70e7cb99 as nil")
	}
	if p.File == nil {
		return fmt.Errorf("unable to encode preliminaryUploadFile#70e7cb99: field file is nil")
	}
	if err := p.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode preliminaryUploadFile#70e7cb99: field file: %w", err)
	}
	if p.FileType == nil {
		return fmt.Errorf("unable to encode preliminaryUploadFile#70e7cb99: field file_type is nil")
	}
	if err := p.FileType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode preliminaryUploadFile#70e7cb99: field file_type: %w", err)
	}
	b.PutInt32(p.Priority)
	return nil
}

// Decode implements bin.Decoder.
func (p *PreliminaryUploadFileRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode preliminaryUploadFile#70e7cb99 to nil")
	}
	if err := b.ConsumeID(PreliminaryUploadFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode preliminaryUploadFile#70e7cb99: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PreliminaryUploadFileRequest) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode preliminaryUploadFile#70e7cb99 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode preliminaryUploadFile#70e7cb99: field file: %w", err)
		}
		p.File = value
	}
	{
		value, err := DecodeFileType(b)
		if err != nil {
			return fmt.Errorf("unable to decode preliminaryUploadFile#70e7cb99: field file_type: %w", err)
		}
		p.FileType = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode preliminaryUploadFile#70e7cb99: field priority: %w", err)
		}
		p.Priority = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PreliminaryUploadFileRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode preliminaryUploadFile#70e7cb99 as nil")
	}
	b.ObjStart()
	b.PutID("preliminaryUploadFile")
	b.Comma()
	b.FieldStart("file")
	if p.File == nil {
		return fmt.Errorf("unable to encode preliminaryUploadFile#70e7cb99: field file is nil")
	}
	if err := p.File.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode preliminaryUploadFile#70e7cb99: field file: %w", err)
	}
	b.Comma()
	b.FieldStart("file_type")
	if p.FileType == nil {
		return fmt.Errorf("unable to encode preliminaryUploadFile#70e7cb99: field file_type is nil")
	}
	if err := p.FileType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode preliminaryUploadFile#70e7cb99: field file_type: %w", err)
	}
	b.Comma()
	b.FieldStart("priority")
	b.PutInt32(p.Priority)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PreliminaryUploadFileRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode preliminaryUploadFile#70e7cb99 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("preliminaryUploadFile"); err != nil {
				return fmt.Errorf("unable to decode preliminaryUploadFile#70e7cb99: %w", err)
			}
		case "file":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode preliminaryUploadFile#70e7cb99: field file: %w", err)
			}
			p.File = value
		case "file_type":
			value, err := DecodeTDLibJSONFileType(b)
			if err != nil {
				return fmt.Errorf("unable to decode preliminaryUploadFile#70e7cb99: field file_type: %w", err)
			}
			p.FileType = value
		case "priority":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode preliminaryUploadFile#70e7cb99: field priority: %w", err)
			}
			p.Priority = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFile returns value of File field.
func (p *PreliminaryUploadFileRequest) GetFile() (value InputFileClass) {
	if p == nil {
		return
	}
	return p.File
}

// GetFileType returns value of FileType field.
func (p *PreliminaryUploadFileRequest) GetFileType() (value FileTypeClass) {
	if p == nil {
		return
	}
	return p.FileType
}

// GetPriority returns value of Priority field.
func (p *PreliminaryUploadFileRequest) GetPriority() (value int32) {
	if p == nil {
		return
	}
	return p.Priority
}

// PreliminaryUploadFile invokes method preliminaryUploadFile#70e7cb99 returning error if any.
func (c *Client) PreliminaryUploadFile(ctx context.Context, request *PreliminaryUploadFileRequest) (*File, error) {
	var result File

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
