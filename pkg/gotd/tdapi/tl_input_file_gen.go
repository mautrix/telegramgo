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

// InputFileID represents TL type `inputFileId#6aa08b0d`.
type InputFileID struct {
	// Unique file identifier
	ID int32
}

// InputFileIDTypeID is TL type id of InputFileID.
const InputFileIDTypeID = 0x6aa08b0d

// construct implements constructor of InputFileClass.
func (i InputFileID) construct() InputFileClass { return &i }

// Ensuring interfaces in compile-time for InputFileID.
var (
	_ bin.Encoder     = &InputFileID{}
	_ bin.Decoder     = &InputFileID{}
	_ bin.BareEncoder = &InputFileID{}
	_ bin.BareDecoder = &InputFileID{}

	_ InputFileClass = &InputFileID{}
)

func (i *InputFileID) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputFileID) String() string {
	if i == nil {
		return "InputFileID(nil)"
	}
	type Alias InputFileID
	return fmt.Sprintf("InputFileID%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputFileID) TypeID() uint32 {
	return InputFileIDTypeID
}

// TypeName returns name of type in TL schema.
func (*InputFileID) TypeName() string {
	return "inputFileId"
}

// TypeInfo returns info about TL type.
func (i *InputFileID) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputFileId",
		ID:   InputFileIDTypeID,
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputFileID) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileId#6aa08b0d as nil")
	}
	b.PutID(InputFileIDTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputFileID) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileId#6aa08b0d as nil")
	}
	b.PutInt32(i.ID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputFileID) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileId#6aa08b0d to nil")
	}
	if err := b.ConsumeID(InputFileIDTypeID); err != nil {
		return fmt.Errorf("unable to decode inputFileId#6aa08b0d: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputFileID) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileId#6aa08b0d to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileId#6aa08b0d: field id: %w", err)
		}
		i.ID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputFileID) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileId#6aa08b0d as nil")
	}
	b.ObjStart()
	b.PutID("inputFileId")
	b.Comma()
	b.FieldStart("id")
	b.PutInt32(i.ID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputFileID) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileId#6aa08b0d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputFileId"); err != nil {
				return fmt.Errorf("unable to decode inputFileId#6aa08b0d: %w", err)
			}
		case "id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode inputFileId#6aa08b0d: field id: %w", err)
			}
			i.ID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (i *InputFileID) GetID() (value int32) {
	if i == nil {
		return
	}
	return i.ID
}

// InputFileRemote represents TL type `inputFileRemote#f9968b3e`.
type InputFileRemote struct {
	// Remote file identifier
	ID string
}

// InputFileRemoteTypeID is TL type id of InputFileRemote.
const InputFileRemoteTypeID = 0xf9968b3e

// construct implements constructor of InputFileClass.
func (i InputFileRemote) construct() InputFileClass { return &i }

// Ensuring interfaces in compile-time for InputFileRemote.
var (
	_ bin.Encoder     = &InputFileRemote{}
	_ bin.Decoder     = &InputFileRemote{}
	_ bin.BareEncoder = &InputFileRemote{}
	_ bin.BareDecoder = &InputFileRemote{}

	_ InputFileClass = &InputFileRemote{}
)

func (i *InputFileRemote) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputFileRemote) String() string {
	if i == nil {
		return "InputFileRemote(nil)"
	}
	type Alias InputFileRemote
	return fmt.Sprintf("InputFileRemote%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputFileRemote) TypeID() uint32 {
	return InputFileRemoteTypeID
}

// TypeName returns name of type in TL schema.
func (*InputFileRemote) TypeName() string {
	return "inputFileRemote"
}

// TypeInfo returns info about TL type.
func (i *InputFileRemote) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputFileRemote",
		ID:   InputFileRemoteTypeID,
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputFileRemote) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileRemote#f9968b3e as nil")
	}
	b.PutID(InputFileRemoteTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputFileRemote) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileRemote#f9968b3e as nil")
	}
	b.PutString(i.ID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputFileRemote) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileRemote#f9968b3e to nil")
	}
	if err := b.ConsumeID(InputFileRemoteTypeID); err != nil {
		return fmt.Errorf("unable to decode inputFileRemote#f9968b3e: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputFileRemote) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileRemote#f9968b3e to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileRemote#f9968b3e: field id: %w", err)
		}
		i.ID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputFileRemote) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileRemote#f9968b3e as nil")
	}
	b.ObjStart()
	b.PutID("inputFileRemote")
	b.Comma()
	b.FieldStart("id")
	b.PutString(i.ID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputFileRemote) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileRemote#f9968b3e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputFileRemote"); err != nil {
				return fmt.Errorf("unable to decode inputFileRemote#f9968b3e: %w", err)
			}
		case "id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputFileRemote#f9968b3e: field id: %w", err)
			}
			i.ID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (i *InputFileRemote) GetID() (value string) {
	if i == nil {
		return
	}
	return i.ID
}

// InputFileLocal represents TL type `inputFileLocal#7a8c8ac7`.
type InputFileLocal struct {
	// Local path to the file
	Path string
}

// InputFileLocalTypeID is TL type id of InputFileLocal.
const InputFileLocalTypeID = 0x7a8c8ac7

// construct implements constructor of InputFileClass.
func (i InputFileLocal) construct() InputFileClass { return &i }

// Ensuring interfaces in compile-time for InputFileLocal.
var (
	_ bin.Encoder     = &InputFileLocal{}
	_ bin.Decoder     = &InputFileLocal{}
	_ bin.BareEncoder = &InputFileLocal{}
	_ bin.BareDecoder = &InputFileLocal{}

	_ InputFileClass = &InputFileLocal{}
)

func (i *InputFileLocal) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Path == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputFileLocal) String() string {
	if i == nil {
		return "InputFileLocal(nil)"
	}
	type Alias InputFileLocal
	return fmt.Sprintf("InputFileLocal%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputFileLocal) TypeID() uint32 {
	return InputFileLocalTypeID
}

// TypeName returns name of type in TL schema.
func (*InputFileLocal) TypeName() string {
	return "inputFileLocal"
}

// TypeInfo returns info about TL type.
func (i *InputFileLocal) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputFileLocal",
		ID:   InputFileLocalTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Path",
			SchemaName: "path",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputFileLocal) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileLocal#7a8c8ac7 as nil")
	}
	b.PutID(InputFileLocalTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputFileLocal) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileLocal#7a8c8ac7 as nil")
	}
	b.PutString(i.Path)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputFileLocal) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileLocal#7a8c8ac7 to nil")
	}
	if err := b.ConsumeID(InputFileLocalTypeID); err != nil {
		return fmt.Errorf("unable to decode inputFileLocal#7a8c8ac7: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputFileLocal) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileLocal#7a8c8ac7 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileLocal#7a8c8ac7: field path: %w", err)
		}
		i.Path = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputFileLocal) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileLocal#7a8c8ac7 as nil")
	}
	b.ObjStart()
	b.PutID("inputFileLocal")
	b.Comma()
	b.FieldStart("path")
	b.PutString(i.Path)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputFileLocal) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileLocal#7a8c8ac7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputFileLocal"); err != nil {
				return fmt.Errorf("unable to decode inputFileLocal#7a8c8ac7: %w", err)
			}
		case "path":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputFileLocal#7a8c8ac7: field path: %w", err)
			}
			i.Path = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPath returns value of Path field.
func (i *InputFileLocal) GetPath() (value string) {
	if i == nil {
		return
	}
	return i.Path
}

// InputFileGenerated represents TL type `inputFileGenerated#b0862800`.
type InputFileGenerated struct {
	// Local path to a file from which the file is generated. The path doesn't have to be a
	// valid path and is used by TDLib only to detect name and MIME type of the generated
	// file
	OriginalPath string
	// String specifying the conversion applied to the original file; must be persistent
	// across application restarts. Conversions beginning with '#' are reserved for internal
	// TDLib usage
	Conversion string
	// Expected size of the generated file, in bytes; pass 0 if unknown
	ExpectedSize int64
}

// InputFileGeneratedTypeID is TL type id of InputFileGenerated.
const InputFileGeneratedTypeID = 0xb0862800

// construct implements constructor of InputFileClass.
func (i InputFileGenerated) construct() InputFileClass { return &i }

// Ensuring interfaces in compile-time for InputFileGenerated.
var (
	_ bin.Encoder     = &InputFileGenerated{}
	_ bin.Decoder     = &InputFileGenerated{}
	_ bin.BareEncoder = &InputFileGenerated{}
	_ bin.BareDecoder = &InputFileGenerated{}

	_ InputFileClass = &InputFileGenerated{}
)

func (i *InputFileGenerated) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.OriginalPath == "") {
		return false
	}
	if !(i.Conversion == "") {
		return false
	}
	if !(i.ExpectedSize == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputFileGenerated) String() string {
	if i == nil {
		return "InputFileGenerated(nil)"
	}
	type Alias InputFileGenerated
	return fmt.Sprintf("InputFileGenerated%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputFileGenerated) TypeID() uint32 {
	return InputFileGeneratedTypeID
}

// TypeName returns name of type in TL schema.
func (*InputFileGenerated) TypeName() string {
	return "inputFileGenerated"
}

// TypeInfo returns info about TL type.
func (i *InputFileGenerated) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputFileGenerated",
		ID:   InputFileGeneratedTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "OriginalPath",
			SchemaName: "original_path",
		},
		{
			Name:       "Conversion",
			SchemaName: "conversion",
		},
		{
			Name:       "ExpectedSize",
			SchemaName: "expected_size",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputFileGenerated) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileGenerated#b0862800 as nil")
	}
	b.PutID(InputFileGeneratedTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputFileGenerated) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileGenerated#b0862800 as nil")
	}
	b.PutString(i.OriginalPath)
	b.PutString(i.Conversion)
	b.PutInt53(i.ExpectedSize)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputFileGenerated) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileGenerated#b0862800 to nil")
	}
	if err := b.ConsumeID(InputFileGeneratedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputFileGenerated#b0862800: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputFileGenerated) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileGenerated#b0862800 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileGenerated#b0862800: field original_path: %w", err)
		}
		i.OriginalPath = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileGenerated#b0862800: field conversion: %w", err)
		}
		i.Conversion = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileGenerated#b0862800: field expected_size: %w", err)
		}
		i.ExpectedSize = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputFileGenerated) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileGenerated#b0862800 as nil")
	}
	b.ObjStart()
	b.PutID("inputFileGenerated")
	b.Comma()
	b.FieldStart("original_path")
	b.PutString(i.OriginalPath)
	b.Comma()
	b.FieldStart("conversion")
	b.PutString(i.Conversion)
	b.Comma()
	b.FieldStart("expected_size")
	b.PutInt53(i.ExpectedSize)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputFileGenerated) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileGenerated#b0862800 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputFileGenerated"); err != nil {
				return fmt.Errorf("unable to decode inputFileGenerated#b0862800: %w", err)
			}
		case "original_path":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputFileGenerated#b0862800: field original_path: %w", err)
			}
			i.OriginalPath = value
		case "conversion":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputFileGenerated#b0862800: field conversion: %w", err)
			}
			i.Conversion = value
		case "expected_size":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode inputFileGenerated#b0862800: field expected_size: %w", err)
			}
			i.ExpectedSize = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetOriginalPath returns value of OriginalPath field.
func (i *InputFileGenerated) GetOriginalPath() (value string) {
	if i == nil {
		return
	}
	return i.OriginalPath
}

// GetConversion returns value of Conversion field.
func (i *InputFileGenerated) GetConversion() (value string) {
	if i == nil {
		return
	}
	return i.Conversion
}

// GetExpectedSize returns value of ExpectedSize field.
func (i *InputFileGenerated) GetExpectedSize() (value int64) {
	if i == nil {
		return
	}
	return i.ExpectedSize
}

// InputFileClassName is schema name of InputFileClass.
const InputFileClassName = "InputFile"

// InputFileClass represents InputFile generic type.
//
// Example:
//
//	g, err := tdapi.DecodeInputFile(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.InputFileID: // inputFileId#6aa08b0d
//	case *tdapi.InputFileRemote: // inputFileRemote#f9968b3e
//	case *tdapi.InputFileLocal: // inputFileLocal#7a8c8ac7
//	case *tdapi.InputFileGenerated: // inputFileGenerated#b0862800
//	default: panic(v)
//	}
type InputFileClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputFileClass

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

// DecodeInputFile implements binary de-serialization for InputFileClass.
func DecodeInputFile(buf *bin.Buffer) (InputFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputFileIDTypeID:
		// Decoding inputFileId#6aa08b0d.
		v := InputFileID{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileClass: %w", err)
		}
		return &v, nil
	case InputFileRemoteTypeID:
		// Decoding inputFileRemote#f9968b3e.
		v := InputFileRemote{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileClass: %w", err)
		}
		return &v, nil
	case InputFileLocalTypeID:
		// Decoding inputFileLocal#7a8c8ac7.
		v := InputFileLocal{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileClass: %w", err)
		}
		return &v, nil
	case InputFileGeneratedTypeID:
		// Decoding inputFileGenerated#b0862800.
		v := InputFileGenerated{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONInputFile implements binary de-serialization for InputFileClass.
func DecodeTDLibJSONInputFile(buf tdjson.Decoder) (InputFileClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "inputFileId":
		// Decoding inputFileId#6aa08b0d.
		v := InputFileID{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileClass: %w", err)
		}
		return &v, nil
	case "inputFileRemote":
		// Decoding inputFileRemote#f9968b3e.
		v := InputFileRemote{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileClass: %w", err)
		}
		return &v, nil
	case "inputFileLocal":
		// Decoding inputFileLocal#7a8c8ac7.
		v := InputFileLocal{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileClass: %w", err)
		}
		return &v, nil
	case "inputFileGenerated":
		// Decoding inputFileGenerated#b0862800.
		v := InputFileGenerated{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputFileClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// InputFile boxes the InputFileClass providing a helper.
type InputFileBox struct {
	InputFile InputFileClass
}

// Decode implements bin.Decoder for InputFileBox.
func (b *InputFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputFileBox to nil")
	}
	v, err := DecodeInputFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputFile = v
	return nil
}

// Encode implements bin.Encode for InputFileBox.
func (b *InputFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputFile == nil {
		return fmt.Errorf("unable to encode InputFileClass as nil")
	}
	return b.InputFile.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for InputFileBox.
func (b *InputFileBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputFileBox to nil")
	}
	v, err := DecodeTDLibJSONInputFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputFile = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for InputFileBox.
func (b *InputFileBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.InputFile == nil {
		return fmt.Errorf("unable to encode InputFileClass as nil")
	}
	return b.InputFile.EncodeTDLibJSON(buf)
}
