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

// PaymentsExportInvoiceRequest represents TL type `payments.exportInvoice#f91b065`.
// Generate an invoice deep link¹
//
// Links:
//  1. https://core.telegram.org/api/links#invoice-links
//
// See https://core.telegram.org/method/payments.exportInvoice for reference.
type PaymentsExportInvoiceRequest struct {
	// Invoice
	InvoiceMedia InputMediaClass
}

// PaymentsExportInvoiceRequestTypeID is TL type id of PaymentsExportInvoiceRequest.
const PaymentsExportInvoiceRequestTypeID = 0xf91b065

// Ensuring interfaces in compile-time for PaymentsExportInvoiceRequest.
var (
	_ bin.Encoder     = &PaymentsExportInvoiceRequest{}
	_ bin.Decoder     = &PaymentsExportInvoiceRequest{}
	_ bin.BareEncoder = &PaymentsExportInvoiceRequest{}
	_ bin.BareDecoder = &PaymentsExportInvoiceRequest{}
)

func (e *PaymentsExportInvoiceRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.InvoiceMedia == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *PaymentsExportInvoiceRequest) String() string {
	if e == nil {
		return "PaymentsExportInvoiceRequest(nil)"
	}
	type Alias PaymentsExportInvoiceRequest
	return fmt.Sprintf("PaymentsExportInvoiceRequest%+v", Alias(*e))
}

// FillFrom fills PaymentsExportInvoiceRequest from given interface.
func (e *PaymentsExportInvoiceRequest) FillFrom(from interface {
	GetInvoiceMedia() (value InputMediaClass)
}) {
	e.InvoiceMedia = from.GetInvoiceMedia()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsExportInvoiceRequest) TypeID() uint32 {
	return PaymentsExportInvoiceRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsExportInvoiceRequest) TypeName() string {
	return "payments.exportInvoice"
}

// TypeInfo returns info about TL type.
func (e *PaymentsExportInvoiceRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.exportInvoice",
		ID:   PaymentsExportInvoiceRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InvoiceMedia",
			SchemaName: "invoice_media",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *PaymentsExportInvoiceRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode payments.exportInvoice#f91b065 as nil")
	}
	b.PutID(PaymentsExportInvoiceRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *PaymentsExportInvoiceRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode payments.exportInvoice#f91b065 as nil")
	}
	if e.InvoiceMedia == nil {
		return fmt.Errorf("unable to encode payments.exportInvoice#f91b065: field invoice_media is nil")
	}
	if err := e.InvoiceMedia.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.exportInvoice#f91b065: field invoice_media: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *PaymentsExportInvoiceRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode payments.exportInvoice#f91b065 to nil")
	}
	if err := b.ConsumeID(PaymentsExportInvoiceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.exportInvoice#f91b065: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *PaymentsExportInvoiceRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode payments.exportInvoice#f91b065 to nil")
	}
	{
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.exportInvoice#f91b065: field invoice_media: %w", err)
		}
		e.InvoiceMedia = value
	}
	return nil
}

// GetInvoiceMedia returns value of InvoiceMedia field.
func (e *PaymentsExportInvoiceRequest) GetInvoiceMedia() (value InputMediaClass) {
	if e == nil {
		return
	}
	return e.InvoiceMedia
}

// PaymentsExportInvoice invokes method payments.exportInvoice#f91b065 returning error if any.
// Generate an invoice deep link¹
//
// Links:
//  1. https://core.telegram.org/api/links#invoice-links
//
// Possible errors:
//
//	400 CURRENCY_TOTAL_AMOUNT_INVALID: The total amount of all prices is invalid.
//	400 INVOICE_PAYLOAD_INVALID: The specified invoice payload is invalid.
//	400 MEDIA_INVALID: Media invalid.
//	400 PAYMENT_PROVIDER_INVALID: The specified payment provider is invalid.
//
// See https://core.telegram.org/method/payments.exportInvoice for reference.
// Can be used by bots.
func (c *Client) PaymentsExportInvoice(ctx context.Context, invoicemedia InputMediaClass) (*PaymentsExportedInvoice, error) {
	var result PaymentsExportedInvoice

	request := &PaymentsExportInvoiceRequest{
		InvoiceMedia: invoicemedia,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
