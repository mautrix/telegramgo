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

// PaymentsExportedInvoice represents TL type `payments.exportedInvoice#aed0cbd9`.
// Exported invoice deep link¹
//
// Links:
//  1. https://core.telegram.org/api/links#invoice-links
//
// See https://core.telegram.org/constructor/payments.exportedInvoice for reference.
type PaymentsExportedInvoice struct {
	// Exported invoice deep link¹
	//
	// Links:
	//  1) https://core.telegram.org/api/links#invoice-links
	URL string
}

// PaymentsExportedInvoiceTypeID is TL type id of PaymentsExportedInvoice.
const PaymentsExportedInvoiceTypeID = 0xaed0cbd9

// Ensuring interfaces in compile-time for PaymentsExportedInvoice.
var (
	_ bin.Encoder     = &PaymentsExportedInvoice{}
	_ bin.Decoder     = &PaymentsExportedInvoice{}
	_ bin.BareEncoder = &PaymentsExportedInvoice{}
	_ bin.BareDecoder = &PaymentsExportedInvoice{}
)

func (e *PaymentsExportedInvoice) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *PaymentsExportedInvoice) String() string {
	if e == nil {
		return "PaymentsExportedInvoice(nil)"
	}
	type Alias PaymentsExportedInvoice
	return fmt.Sprintf("PaymentsExportedInvoice%+v", Alias(*e))
}

// FillFrom fills PaymentsExportedInvoice from given interface.
func (e *PaymentsExportedInvoice) FillFrom(from interface {
	GetURL() (value string)
}) {
	e.URL = from.GetURL()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsExportedInvoice) TypeID() uint32 {
	return PaymentsExportedInvoiceTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsExportedInvoice) TypeName() string {
	return "payments.exportedInvoice"
}

// TypeInfo returns info about TL type.
func (e *PaymentsExportedInvoice) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.exportedInvoice",
		ID:   PaymentsExportedInvoiceTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *PaymentsExportedInvoice) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode payments.exportedInvoice#aed0cbd9 as nil")
	}
	b.PutID(PaymentsExportedInvoiceTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *PaymentsExportedInvoice) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode payments.exportedInvoice#aed0cbd9 as nil")
	}
	b.PutString(e.URL)
	return nil
}

// Decode implements bin.Decoder.
func (e *PaymentsExportedInvoice) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode payments.exportedInvoice#aed0cbd9 to nil")
	}
	if err := b.ConsumeID(PaymentsExportedInvoiceTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.exportedInvoice#aed0cbd9: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *PaymentsExportedInvoice) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode payments.exportedInvoice#aed0cbd9 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.exportedInvoice#aed0cbd9: field url: %w", err)
		}
		e.URL = value
	}
	return nil
}

// GetURL returns value of URL field.
func (e *PaymentsExportedInvoice) GetURL() (value string) {
	if e == nil {
		return
	}
	return e.URL
}
