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

// DocumentEmpty represents TL type `documentEmpty#36f8c871`.
// Empty constructor, document doesn't exist.
//
// See https://core.telegram.org/constructor/documentEmpty for reference.
type DocumentEmpty struct {
	// Document ID or 0
	ID int64
}

// DocumentEmptyTypeID is TL type id of DocumentEmpty.
const DocumentEmptyTypeID = 0x36f8c871

// construct implements constructor of DocumentClass.
func (d DocumentEmpty) construct() DocumentClass { return &d }

// Ensuring interfaces in compile-time for DocumentEmpty.
var (
	_ bin.Encoder     = &DocumentEmpty{}
	_ bin.Decoder     = &DocumentEmpty{}
	_ bin.BareEncoder = &DocumentEmpty{}
	_ bin.BareDecoder = &DocumentEmpty{}

	_ DocumentClass = &DocumentEmpty{}
)

func (d *DocumentEmpty) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DocumentEmpty) String() string {
	if d == nil {
		return "DocumentEmpty(nil)"
	}
	type Alias DocumentEmpty
	return fmt.Sprintf("DocumentEmpty%+v", Alias(*d))
}

// FillFrom fills DocumentEmpty from given interface.
func (d *DocumentEmpty) FillFrom(from interface {
	GetID() (value int64)
}) {
	d.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DocumentEmpty) TypeID() uint32 {
	return DocumentEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*DocumentEmpty) TypeName() string {
	return "documentEmpty"
}

// TypeInfo returns info about TL type.
func (d *DocumentEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "documentEmpty",
		ID:   DocumentEmptyTypeID,
	}
	if d == nil {
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
func (d *DocumentEmpty) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentEmpty#36f8c871 as nil")
	}
	b.PutID(DocumentEmptyTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DocumentEmpty) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentEmpty#36f8c871 as nil")
	}
	b.PutLong(d.ID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DocumentEmpty) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentEmpty#36f8c871 to nil")
	}
	if err := b.ConsumeID(DocumentEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode documentEmpty#36f8c871: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DocumentEmpty) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentEmpty#36f8c871 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode documentEmpty#36f8c871: field id: %w", err)
		}
		d.ID = value
	}
	return nil
}

// GetID returns value of ID field.
func (d *DocumentEmpty) GetID() (value int64) {
	if d == nil {
		return
	}
	return d.ID
}

// Document represents TL type `document#8fd4c4d8`.
// Document
//
// See https://core.telegram.org/constructor/document for reference.
type Document struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Document ID
	ID int64
	// Check sum, dependent on document ID
	AccessHash int64
	// File reference¹
	//
	// Links:
	//  1) https://core.telegram.org/api/file_reference
	FileReference []byte
	// Creation date
	Date int
	// MIME type
	MimeType string
	// Size
	Size int64
	// Thumbnails
	//
	// Use SetThumbs and GetThumbs helpers.
	Thumbs []PhotoSizeClass
	// Video thumbnails
	//
	// Use SetVideoThumbs and GetVideoThumbs helpers.
	VideoThumbs []VideoSizeClass
	// DC ID
	DCID int
	// Attributes
	Attributes []DocumentAttributeClass
}

// DocumentTypeID is TL type id of Document.
const DocumentTypeID = 0x8fd4c4d8

// construct implements constructor of DocumentClass.
func (d Document) construct() DocumentClass { return &d }

// Ensuring interfaces in compile-time for Document.
var (
	_ bin.Encoder     = &Document{}
	_ bin.Decoder     = &Document{}
	_ bin.BareEncoder = &Document{}
	_ bin.BareDecoder = &Document{}

	_ DocumentClass = &Document{}
)

func (d *Document) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.ID == 0) {
		return false
	}
	if !(d.AccessHash == 0) {
		return false
	}
	if !(d.FileReference == nil) {
		return false
	}
	if !(d.Date == 0) {
		return false
	}
	if !(d.MimeType == "") {
		return false
	}
	if !(d.Size == 0) {
		return false
	}
	if !(d.Thumbs == nil) {
		return false
	}
	if !(d.VideoThumbs == nil) {
		return false
	}
	if !(d.DCID == 0) {
		return false
	}
	if !(d.Attributes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *Document) String() string {
	if d == nil {
		return "Document(nil)"
	}
	type Alias Document
	return fmt.Sprintf("Document%+v", Alias(*d))
}

// FillFrom fills Document from given interface.
func (d *Document) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetFileReference() (value []byte)
	GetDate() (value int)
	GetMimeType() (value string)
	GetSize() (value int64)
	GetThumbs() (value []PhotoSizeClass, ok bool)
	GetVideoThumbs() (value []VideoSizeClass, ok bool)
	GetDCID() (value int)
	GetAttributes() (value []DocumentAttributeClass)
}) {
	d.ID = from.GetID()
	d.AccessHash = from.GetAccessHash()
	d.FileReference = from.GetFileReference()
	d.Date = from.GetDate()
	d.MimeType = from.GetMimeType()
	d.Size = from.GetSize()
	if val, ok := from.GetThumbs(); ok {
		d.Thumbs = val
	}

	if val, ok := from.GetVideoThumbs(); ok {
		d.VideoThumbs = val
	}

	d.DCID = from.GetDCID()
	d.Attributes = from.GetAttributes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Document) TypeID() uint32 {
	return DocumentTypeID
}

// TypeName returns name of type in TL schema.
func (*Document) TypeName() string {
	return "document"
}

// TypeInfo returns info about TL type.
func (d *Document) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "document",
		ID:   DocumentTypeID,
	}
	if d == nil {
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
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "MimeType",
			SchemaName: "mime_type",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "Thumbs",
			SchemaName: "thumbs",
			Null:       !d.Flags.Has(0),
		},
		{
			Name:       "VideoThumbs",
			SchemaName: "video_thumbs",
			Null:       !d.Flags.Has(1),
		},
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
		{
			Name:       "Attributes",
			SchemaName: "attributes",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (d *Document) SetFlags() {
	if !(d.Thumbs == nil) {
		d.Flags.Set(0)
	}
	if !(d.VideoThumbs == nil) {
		d.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (d *Document) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode document#8fd4c4d8 as nil")
	}
	b.PutID(DocumentTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *Document) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode document#8fd4c4d8 as nil")
	}
	d.SetFlags()
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode document#8fd4c4d8: field flags: %w", err)
	}
	b.PutLong(d.ID)
	b.PutLong(d.AccessHash)
	b.PutBytes(d.FileReference)
	b.PutInt(d.Date)
	b.PutString(d.MimeType)
	b.PutLong(d.Size)
	if d.Flags.Has(0) {
		b.PutVectorHeader(len(d.Thumbs))
		for idx, v := range d.Thumbs {
			if v == nil {
				return fmt.Errorf("unable to encode document#8fd4c4d8: field thumbs element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode document#8fd4c4d8: field thumbs element with index %d: %w", idx, err)
			}
		}
	}
	if d.Flags.Has(1) {
		b.PutVectorHeader(len(d.VideoThumbs))
		for idx, v := range d.VideoThumbs {
			if v == nil {
				return fmt.Errorf("unable to encode document#8fd4c4d8: field video_thumbs element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode document#8fd4c4d8: field video_thumbs element with index %d: %w", idx, err)
			}
		}
	}
	b.PutInt(d.DCID)
	b.PutVectorHeader(len(d.Attributes))
	for idx, v := range d.Attributes {
		if v == nil {
			return fmt.Errorf("unable to encode document#8fd4c4d8: field attributes element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode document#8fd4c4d8: field attributes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *Document) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode document#8fd4c4d8 to nil")
	}
	if err := b.ConsumeID(DocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode document#8fd4c4d8: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *Document) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode document#8fd4c4d8 to nil")
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode document#8fd4c4d8: field flags: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode document#8fd4c4d8: field id: %w", err)
		}
		d.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode document#8fd4c4d8: field access_hash: %w", err)
		}
		d.AccessHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode document#8fd4c4d8: field file_reference: %w", err)
		}
		d.FileReference = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode document#8fd4c4d8: field date: %w", err)
		}
		d.Date = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode document#8fd4c4d8: field mime_type: %w", err)
		}
		d.MimeType = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode document#8fd4c4d8: field size: %w", err)
		}
		d.Size = value
	}
	if d.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode document#8fd4c4d8: field thumbs: %w", err)
		}

		if headerLen > 0 {
			d.Thumbs = make([]PhotoSizeClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePhotoSize(b)
			if err != nil {
				return fmt.Errorf("unable to decode document#8fd4c4d8: field thumbs: %w", err)
			}
			d.Thumbs = append(d.Thumbs, value)
		}
	}
	if d.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode document#8fd4c4d8: field video_thumbs: %w", err)
		}

		if headerLen > 0 {
			d.VideoThumbs = make([]VideoSizeClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeVideoSize(b)
			if err != nil {
				return fmt.Errorf("unable to decode document#8fd4c4d8: field video_thumbs: %w", err)
			}
			d.VideoThumbs = append(d.VideoThumbs, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode document#8fd4c4d8: field dc_id: %w", err)
		}
		d.DCID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode document#8fd4c4d8: field attributes: %w", err)
		}

		if headerLen > 0 {
			d.Attributes = make([]DocumentAttributeClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocumentAttribute(b)
			if err != nil {
				return fmt.Errorf("unable to decode document#8fd4c4d8: field attributes: %w", err)
			}
			d.Attributes = append(d.Attributes, value)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (d *Document) GetID() (value int64) {
	if d == nil {
		return
	}
	return d.ID
}

// GetAccessHash returns value of AccessHash field.
func (d *Document) GetAccessHash() (value int64) {
	if d == nil {
		return
	}
	return d.AccessHash
}

// GetFileReference returns value of FileReference field.
func (d *Document) GetFileReference() (value []byte) {
	if d == nil {
		return
	}
	return d.FileReference
}

// GetDate returns value of Date field.
func (d *Document) GetDate() (value int) {
	if d == nil {
		return
	}
	return d.Date
}

// GetMimeType returns value of MimeType field.
func (d *Document) GetMimeType() (value string) {
	if d == nil {
		return
	}
	return d.MimeType
}

// GetSize returns value of Size field.
func (d *Document) GetSize() (value int64) {
	if d == nil {
		return
	}
	return d.Size
}

// SetThumbs sets value of Thumbs conditional field.
func (d *Document) SetThumbs(value []PhotoSizeClass) {
	d.Flags.Set(0)
	d.Thumbs = value
}

// GetThumbs returns value of Thumbs conditional field and
// boolean which is true if field was set.
func (d *Document) GetThumbs() (value []PhotoSizeClass, ok bool) {
	if d == nil {
		return
	}
	if !d.Flags.Has(0) {
		return value, false
	}
	return d.Thumbs, true
}

// SetVideoThumbs sets value of VideoThumbs conditional field.
func (d *Document) SetVideoThumbs(value []VideoSizeClass) {
	d.Flags.Set(1)
	d.VideoThumbs = value
}

// GetVideoThumbs returns value of VideoThumbs conditional field and
// boolean which is true if field was set.
func (d *Document) GetVideoThumbs() (value []VideoSizeClass, ok bool) {
	if d == nil {
		return
	}
	if !d.Flags.Has(1) {
		return value, false
	}
	return d.VideoThumbs, true
}

// GetDCID returns value of DCID field.
func (d *Document) GetDCID() (value int) {
	if d == nil {
		return
	}
	return d.DCID
}

// GetAttributes returns value of Attributes field.
func (d *Document) GetAttributes() (value []DocumentAttributeClass) {
	if d == nil {
		return
	}
	return d.Attributes
}

// MapThumbs returns field Thumbs wrapped in PhotoSizeClassArray helper.
func (d *Document) MapThumbs() (value PhotoSizeClassArray, ok bool) {
	if !d.Flags.Has(0) {
		return value, false
	}
	return PhotoSizeClassArray(d.Thumbs), true
}

// MapVideoThumbs returns field VideoThumbs wrapped in VideoSizeClassArray helper.
func (d *Document) MapVideoThumbs() (value VideoSizeClassArray, ok bool) {
	if !d.Flags.Has(1) {
		return value, false
	}
	return VideoSizeClassArray(d.VideoThumbs), true
}

// MapAttributes returns field Attributes wrapped in DocumentAttributeClassArray helper.
func (d *Document) MapAttributes() (value DocumentAttributeClassArray) {
	return DocumentAttributeClassArray(d.Attributes)
}

// DocumentClassName is schema name of DocumentClass.
const DocumentClassName = "Document"

// DocumentClass represents Document generic type.
//
// See https://core.telegram.org/type/Document for reference.
//
// Example:
//
//	g, err := tg.DecodeDocument(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.DocumentEmpty: // documentEmpty#36f8c871
//	case *tg.Document: // document#8fd4c4d8
//	default: panic(v)
//	}
type DocumentClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() DocumentClass

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

	// Document ID or 0
	GetID() (value int64)

	// AsNotEmpty tries to map DocumentClass to Document.
	AsNotEmpty() (*Document, bool)
}

// AsInputDocumentFileLocation tries to map Document to InputDocumentFileLocation.
func (d *Document) AsInputDocumentFileLocation() *InputDocumentFileLocation {
	value := new(InputDocumentFileLocation)
	value.ID = d.GetID()
	value.AccessHash = d.GetAccessHash()
	value.FileReference = d.GetFileReference()

	return value
}

// AsInput tries to map Document to InputDocument.
func (d *Document) AsInput() *InputDocument {
	value := new(InputDocument)
	value.ID = d.GetID()
	value.AccessHash = d.GetAccessHash()
	value.FileReference = d.GetFileReference()

	return value
}

// AsNotEmpty tries to map DocumentEmpty to Document.
func (d *DocumentEmpty) AsNotEmpty() (*Document, bool) {
	return nil, false
}

// AsNotEmpty tries to map Document to Document.
func (d *Document) AsNotEmpty() (*Document, bool) {
	return d, true
}

// DecodeDocument implements binary de-serialization for DocumentClass.
func DecodeDocument(buf *bin.Buffer) (DocumentClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DocumentEmptyTypeID:
		// Decoding documentEmpty#36f8c871.
		v := DocumentEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentClass: %w", err)
		}
		return &v, nil
	case DocumentTypeID:
		// Decoding document#8fd4c4d8.
		v := Document{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DocumentClass: %w", bin.NewUnexpectedID(id))
	}
}

// Document boxes the DocumentClass providing a helper.
type DocumentBox struct {
	Document DocumentClass
}

// Decode implements bin.Decoder for DocumentBox.
func (b *DocumentBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode DocumentBox to nil")
	}
	v, err := DecodeDocument(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Document = v
	return nil
}

// Encode implements bin.Encode for DocumentBox.
func (b *DocumentBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Document == nil {
		return fmt.Errorf("unable to encode DocumentClass as nil")
	}
	return b.Document.Encode(buf)
}
