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

// PaymentsSendStarsFormRequest represents TL type `payments.sendStarsForm#7998c914`.
// Make a payment using Telegram Stars, see here »¹ for more info.
//
// Links:
//  1. https://core.telegram.org/api/stars#using-stars
//
// See https://core.telegram.org/method/payments.sendStarsForm for reference.
type PaymentsSendStarsFormRequest struct {
	// Payment form ID
	FormID int64
	// Invoice
	Invoice InputInvoiceClass
}

// PaymentsSendStarsFormRequestTypeID is TL type id of PaymentsSendStarsFormRequest.
const PaymentsSendStarsFormRequestTypeID = 0x7998c914

// Ensuring interfaces in compile-time for PaymentsSendStarsFormRequest.
var (
	_ bin.Encoder     = &PaymentsSendStarsFormRequest{}
	_ bin.Decoder     = &PaymentsSendStarsFormRequest{}
	_ bin.BareEncoder = &PaymentsSendStarsFormRequest{}
	_ bin.BareDecoder = &PaymentsSendStarsFormRequest{}
)

func (s *PaymentsSendStarsFormRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.FormID == 0) {
		return false
	}
	if !(s.Invoice == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PaymentsSendStarsFormRequest) String() string {
	if s == nil {
		return "PaymentsSendStarsFormRequest(nil)"
	}
	type Alias PaymentsSendStarsFormRequest
	return fmt.Sprintf("PaymentsSendStarsFormRequest%+v", Alias(*s))
}

// FillFrom fills PaymentsSendStarsFormRequest from given interface.
func (s *PaymentsSendStarsFormRequest) FillFrom(from interface {
	GetFormID() (value int64)
	GetInvoice() (value InputInvoiceClass)
}) {
	s.FormID = from.GetFormID()
	s.Invoice = from.GetInvoice()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsSendStarsFormRequest) TypeID() uint32 {
	return PaymentsSendStarsFormRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsSendStarsFormRequest) TypeName() string {
	return "payments.sendStarsForm"
}

// TypeInfo returns info about TL type.
func (s *PaymentsSendStarsFormRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.sendStarsForm",
		ID:   PaymentsSendStarsFormRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FormID",
			SchemaName: "form_id",
		},
		{
			Name:       "Invoice",
			SchemaName: "invoice",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *PaymentsSendStarsFormRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.sendStarsForm#7998c914 as nil")
	}
	b.PutID(PaymentsSendStarsFormRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *PaymentsSendStarsFormRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.sendStarsForm#7998c914 as nil")
	}
	b.PutLong(s.FormID)
	if s.Invoice == nil {
		return fmt.Errorf("unable to encode payments.sendStarsForm#7998c914: field invoice is nil")
	}
	if err := s.Invoice.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.sendStarsForm#7998c914: field invoice: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *PaymentsSendStarsFormRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.sendStarsForm#7998c914 to nil")
	}
	if err := b.ConsumeID(PaymentsSendStarsFormRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.sendStarsForm#7998c914: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *PaymentsSendStarsFormRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.sendStarsForm#7998c914 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendStarsForm#7998c914: field form_id: %w", err)
		}
		s.FormID = value
	}
	{
		value, err := DecodeInputInvoice(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendStarsForm#7998c914: field invoice: %w", err)
		}
		s.Invoice = value
	}
	return nil
}

// GetFormID returns value of FormID field.
func (s *PaymentsSendStarsFormRequest) GetFormID() (value int64) {
	if s == nil {
		return
	}
	return s.FormID
}

// GetInvoice returns value of Invoice field.
func (s *PaymentsSendStarsFormRequest) GetInvoice() (value InputInvoiceClass) {
	if s == nil {
		return
	}
	return s.Invoice
}

// PaymentsSendStarsForm invokes method payments.sendStarsForm#7998c914 returning error if any.
// Make a payment using Telegram Stars, see here »¹ for more info.
//
// Links:
//  1. https://core.telegram.org/api/stars#using-stars
//
// Possible errors:
//
//	400 BALANCE_TOO_LOW: The transaction cannot be completed because the current Telegram Stars balance is too low.
//	400 FORM_EXPIRED: The form was generated more than 10 minutes ago and has expired, please re-generate it using payments.getPaymentForm and pass the new form_id.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/payments.sendStarsForm for reference.
func (c *Client) PaymentsSendStarsForm(ctx context.Context, request *PaymentsSendStarsFormRequest) (PaymentsPaymentResultClass, error) {
	var result PaymentsPaymentResultBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.PaymentResult, nil
}
