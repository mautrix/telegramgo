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

// PageBlockTableCell represents TL type `pageBlockTableCell#547fbf66`.
type PageBlockTableCell struct {
	// Cell text; may be null. If the text is null, then the cell must be invisible
	Text RichTextClass
	// True, if it is a header cell
	IsHeader bool
	// The number of columns the cell spans
	Colspan int32
	// The number of rows the cell spans
	Rowspan int32
	// Horizontal cell content alignment
	Align PageBlockHorizontalAlignmentClass
	// Vertical cell content alignment
	Valign PageBlockVerticalAlignmentClass
}

// PageBlockTableCellTypeID is TL type id of PageBlockTableCell.
const PageBlockTableCellTypeID = 0x547fbf66

// Ensuring interfaces in compile-time for PageBlockTableCell.
var (
	_ bin.Encoder     = &PageBlockTableCell{}
	_ bin.Decoder     = &PageBlockTableCell{}
	_ bin.BareEncoder = &PageBlockTableCell{}
	_ bin.BareDecoder = &PageBlockTableCell{}
)

func (p *PageBlockTableCell) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Text == nil) {
		return false
	}
	if !(p.IsHeader == false) {
		return false
	}
	if !(p.Colspan == 0) {
		return false
	}
	if !(p.Rowspan == 0) {
		return false
	}
	if !(p.Align == nil) {
		return false
	}
	if !(p.Valign == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PageBlockTableCell) String() string {
	if p == nil {
		return "PageBlockTableCell(nil)"
	}
	type Alias PageBlockTableCell
	return fmt.Sprintf("PageBlockTableCell%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PageBlockTableCell) TypeID() uint32 {
	return PageBlockTableCellTypeID
}

// TypeName returns name of type in TL schema.
func (*PageBlockTableCell) TypeName() string {
	return "pageBlockTableCell"
}

// TypeInfo returns info about TL type.
func (p *PageBlockTableCell) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pageBlockTableCell",
		ID:   PageBlockTableCellTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "IsHeader",
			SchemaName: "is_header",
		},
		{
			Name:       "Colspan",
			SchemaName: "colspan",
		},
		{
			Name:       "Rowspan",
			SchemaName: "rowspan",
		},
		{
			Name:       "Align",
			SchemaName: "align",
		},
		{
			Name:       "Valign",
			SchemaName: "valign",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PageBlockTableCell) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockTableCell#547fbf66 as nil")
	}
	b.PutID(PageBlockTableCellTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PageBlockTableCell) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockTableCell#547fbf66 as nil")
	}
	if p.Text == nil {
		return fmt.Errorf("unable to encode pageBlockTableCell#547fbf66: field text is nil")
	}
	if err := p.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pageBlockTableCell#547fbf66: field text: %w", err)
	}
	b.PutBool(p.IsHeader)
	b.PutInt32(p.Colspan)
	b.PutInt32(p.Rowspan)
	if p.Align == nil {
		return fmt.Errorf("unable to encode pageBlockTableCell#547fbf66: field align is nil")
	}
	if err := p.Align.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pageBlockTableCell#547fbf66: field align: %w", err)
	}
	if p.Valign == nil {
		return fmt.Errorf("unable to encode pageBlockTableCell#547fbf66: field valign is nil")
	}
	if err := p.Valign.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pageBlockTableCell#547fbf66: field valign: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageBlockTableCell) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockTableCell#547fbf66 to nil")
	}
	if err := b.ConsumeID(PageBlockTableCellTypeID); err != nil {
		return fmt.Errorf("unable to decode pageBlockTableCell#547fbf66: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PageBlockTableCell) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockTableCell#547fbf66 to nil")
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockTableCell#547fbf66: field text: %w", err)
		}
		p.Text = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockTableCell#547fbf66: field is_header: %w", err)
		}
		p.IsHeader = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockTableCell#547fbf66: field colspan: %w", err)
		}
		p.Colspan = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockTableCell#547fbf66: field rowspan: %w", err)
		}
		p.Rowspan = value
	}
	{
		value, err := DecodePageBlockHorizontalAlignment(b)
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockTableCell#547fbf66: field align: %w", err)
		}
		p.Align = value
	}
	{
		value, err := DecodePageBlockVerticalAlignment(b)
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockTableCell#547fbf66: field valign: %w", err)
		}
		p.Valign = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PageBlockTableCell) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockTableCell#547fbf66 as nil")
	}
	b.ObjStart()
	b.PutID("pageBlockTableCell")
	b.Comma()
	b.FieldStart("text")
	if p.Text == nil {
		return fmt.Errorf("unable to encode pageBlockTableCell#547fbf66: field text is nil")
	}
	if err := p.Text.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode pageBlockTableCell#547fbf66: field text: %w", err)
	}
	b.Comma()
	b.FieldStart("is_header")
	b.PutBool(p.IsHeader)
	b.Comma()
	b.FieldStart("colspan")
	b.PutInt32(p.Colspan)
	b.Comma()
	b.FieldStart("rowspan")
	b.PutInt32(p.Rowspan)
	b.Comma()
	b.FieldStart("align")
	if p.Align == nil {
		return fmt.Errorf("unable to encode pageBlockTableCell#547fbf66: field align is nil")
	}
	if err := p.Align.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode pageBlockTableCell#547fbf66: field align: %w", err)
	}
	b.Comma()
	b.FieldStart("valign")
	if p.Valign == nil {
		return fmt.Errorf("unable to encode pageBlockTableCell#547fbf66: field valign is nil")
	}
	if err := p.Valign.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode pageBlockTableCell#547fbf66: field valign: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PageBlockTableCell) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockTableCell#547fbf66 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("pageBlockTableCell"); err != nil {
				return fmt.Errorf("unable to decode pageBlockTableCell#547fbf66: %w", err)
			}
		case "text":
			value, err := DecodeTDLibJSONRichText(b)
			if err != nil {
				return fmt.Errorf("unable to decode pageBlockTableCell#547fbf66: field text: %w", err)
			}
			p.Text = value
		case "is_header":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode pageBlockTableCell#547fbf66: field is_header: %w", err)
			}
			p.IsHeader = value
		case "colspan":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode pageBlockTableCell#547fbf66: field colspan: %w", err)
			}
			p.Colspan = value
		case "rowspan":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode pageBlockTableCell#547fbf66: field rowspan: %w", err)
			}
			p.Rowspan = value
		case "align":
			value, err := DecodeTDLibJSONPageBlockHorizontalAlignment(b)
			if err != nil {
				return fmt.Errorf("unable to decode pageBlockTableCell#547fbf66: field align: %w", err)
			}
			p.Align = value
		case "valign":
			value, err := DecodeTDLibJSONPageBlockVerticalAlignment(b)
			if err != nil {
				return fmt.Errorf("unable to decode pageBlockTableCell#547fbf66: field valign: %w", err)
			}
			p.Valign = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (p *PageBlockTableCell) GetText() (value RichTextClass) {
	if p == nil {
		return
	}
	return p.Text
}

// GetIsHeader returns value of IsHeader field.
func (p *PageBlockTableCell) GetIsHeader() (value bool) {
	if p == nil {
		return
	}
	return p.IsHeader
}

// GetColspan returns value of Colspan field.
func (p *PageBlockTableCell) GetColspan() (value int32) {
	if p == nil {
		return
	}
	return p.Colspan
}

// GetRowspan returns value of Rowspan field.
func (p *PageBlockTableCell) GetRowspan() (value int32) {
	if p == nil {
		return
	}
	return p.Rowspan
}

// GetAlign returns value of Align field.
func (p *PageBlockTableCell) GetAlign() (value PageBlockHorizontalAlignmentClass) {
	if p == nil {
		return
	}
	return p.Align
}

// GetValign returns value of Valign field.
func (p *PageBlockTableCell) GetValign() (value PageBlockVerticalAlignmentClass) {
	if p == nil {
		return
	}
	return p.Valign
}
