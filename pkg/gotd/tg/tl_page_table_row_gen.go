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

// PageTableRow represents TL type `pageTableRow#e0c0c5e5`.
// Table row
//
// See https://core.telegram.org/constructor/pageTableRow for reference.
type PageTableRow struct {
	// Table cells
	Cells []PageTableCell
}

// PageTableRowTypeID is TL type id of PageTableRow.
const PageTableRowTypeID = 0xe0c0c5e5

// Ensuring interfaces in compile-time for PageTableRow.
var (
	_ bin.Encoder     = &PageTableRow{}
	_ bin.Decoder     = &PageTableRow{}
	_ bin.BareEncoder = &PageTableRow{}
	_ bin.BareDecoder = &PageTableRow{}
)

func (p *PageTableRow) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Cells == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PageTableRow) String() string {
	if p == nil {
		return "PageTableRow(nil)"
	}
	type Alias PageTableRow
	return fmt.Sprintf("PageTableRow%+v", Alias(*p))
}

// FillFrom fills PageTableRow from given interface.
func (p *PageTableRow) FillFrom(from interface {
	GetCells() (value []PageTableCell)
}) {
	p.Cells = from.GetCells()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PageTableRow) TypeID() uint32 {
	return PageTableRowTypeID
}

// TypeName returns name of type in TL schema.
func (*PageTableRow) TypeName() string {
	return "pageTableRow"
}

// TypeInfo returns info about TL type.
func (p *PageTableRow) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pageTableRow",
		ID:   PageTableRowTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Cells",
			SchemaName: "cells",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PageTableRow) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageTableRow#e0c0c5e5 as nil")
	}
	b.PutID(PageTableRowTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PageTableRow) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageTableRow#e0c0c5e5 as nil")
	}
	b.PutVectorHeader(len(p.Cells))
	for idx, v := range p.Cells {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode pageTableRow#e0c0c5e5: field cells element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageTableRow) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageTableRow#e0c0c5e5 to nil")
	}
	if err := b.ConsumeID(PageTableRowTypeID); err != nil {
		return fmt.Errorf("unable to decode pageTableRow#e0c0c5e5: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PageTableRow) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageTableRow#e0c0c5e5 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode pageTableRow#e0c0c5e5: field cells: %w", err)
		}

		if headerLen > 0 {
			p.Cells = make([]PageTableCell, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PageTableCell
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode pageTableRow#e0c0c5e5: field cells: %w", err)
			}
			p.Cells = append(p.Cells, value)
		}
	}
	return nil
}

// GetCells returns value of Cells field.
func (p *PageTableRow) GetCells() (value []PageTableCell) {
	if p == nil {
		return
	}
	return p.Cells
}
