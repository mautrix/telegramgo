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

// PaymentsSendPaymentFormRequest represents TL type `payments.sendPaymentForm#2d03522f`.
// Send compiled payment form
//
// See https://core.telegram.org/method/payments.sendPaymentForm for reference.
type PaymentsSendPaymentFormRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Form ID
	FormID int64
	// Invoice
	Invoice InputInvoiceClass
	// ID of saved and validated order info¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/payments.validatedRequestedInfo
	//
	// Use SetRequestedInfoID and GetRequestedInfoID helpers.
	RequestedInfoID string
	// Chosen shipping option ID
	//
	// Use SetShippingOptionID and GetShippingOptionID helpers.
	ShippingOptionID string
	// Payment credentials
	Credentials InputPaymentCredentialsClass
	// Tip, in the smallest units of the currency (integer, not float/double). For example,
	// for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json¹,
	// it shows the number of digits past the decimal point for each currency (2 for the
	// majority of currencies).
	//
	// Links:
	//  1) https://core.telegram.org/bots/payments/currencies.json
	//
	// Use SetTipAmount and GetTipAmount helpers.
	TipAmount int64
}

// PaymentsSendPaymentFormRequestTypeID is TL type id of PaymentsSendPaymentFormRequest.
const PaymentsSendPaymentFormRequestTypeID = 0x2d03522f

// Ensuring interfaces in compile-time for PaymentsSendPaymentFormRequest.
var (
	_ bin.Encoder     = &PaymentsSendPaymentFormRequest{}
	_ bin.Decoder     = &PaymentsSendPaymentFormRequest{}
	_ bin.BareEncoder = &PaymentsSendPaymentFormRequest{}
	_ bin.BareDecoder = &PaymentsSendPaymentFormRequest{}
)

func (s *PaymentsSendPaymentFormRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.FormID == 0) {
		return false
	}
	if !(s.Invoice == nil) {
		return false
	}
	if !(s.RequestedInfoID == "") {
		return false
	}
	if !(s.ShippingOptionID == "") {
		return false
	}
	if !(s.Credentials == nil) {
		return false
	}
	if !(s.TipAmount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PaymentsSendPaymentFormRequest) String() string {
	if s == nil {
		return "PaymentsSendPaymentFormRequest(nil)"
	}
	type Alias PaymentsSendPaymentFormRequest
	return fmt.Sprintf("PaymentsSendPaymentFormRequest%+v", Alias(*s))
}

// FillFrom fills PaymentsSendPaymentFormRequest from given interface.
func (s *PaymentsSendPaymentFormRequest) FillFrom(from interface {
	GetFormID() (value int64)
	GetInvoice() (value InputInvoiceClass)
	GetRequestedInfoID() (value string, ok bool)
	GetShippingOptionID() (value string, ok bool)
	GetCredentials() (value InputPaymentCredentialsClass)
	GetTipAmount() (value int64, ok bool)
}) {
	s.FormID = from.GetFormID()
	s.Invoice = from.GetInvoice()
	if val, ok := from.GetRequestedInfoID(); ok {
		s.RequestedInfoID = val
	}

	if val, ok := from.GetShippingOptionID(); ok {
		s.ShippingOptionID = val
	}

	s.Credentials = from.GetCredentials()
	if val, ok := from.GetTipAmount(); ok {
		s.TipAmount = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsSendPaymentFormRequest) TypeID() uint32 {
	return PaymentsSendPaymentFormRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsSendPaymentFormRequest) TypeName() string {
	return "payments.sendPaymentForm"
}

// TypeInfo returns info about TL type.
func (s *PaymentsSendPaymentFormRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.sendPaymentForm",
		ID:   PaymentsSendPaymentFormRequestTypeID,
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
		{
			Name:       "RequestedInfoID",
			SchemaName: "requested_info_id",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "ShippingOptionID",
			SchemaName: "shipping_option_id",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Credentials",
			SchemaName: "credentials",
		},
		{
			Name:       "TipAmount",
			SchemaName: "tip_amount",
			Null:       !s.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *PaymentsSendPaymentFormRequest) SetFlags() {
	if !(s.RequestedInfoID == "") {
		s.Flags.Set(0)
	}
	if !(s.ShippingOptionID == "") {
		s.Flags.Set(1)
	}
	if !(s.TipAmount == 0) {
		s.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (s *PaymentsSendPaymentFormRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.sendPaymentForm#2d03522f as nil")
	}
	b.PutID(PaymentsSendPaymentFormRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *PaymentsSendPaymentFormRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.sendPaymentForm#2d03522f as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.sendPaymentForm#2d03522f: field flags: %w", err)
	}
	b.PutLong(s.FormID)
	if s.Invoice == nil {
		return fmt.Errorf("unable to encode payments.sendPaymentForm#2d03522f: field invoice is nil")
	}
	if err := s.Invoice.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.sendPaymentForm#2d03522f: field invoice: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutString(s.RequestedInfoID)
	}
	if s.Flags.Has(1) {
		b.PutString(s.ShippingOptionID)
	}
	if s.Credentials == nil {
		return fmt.Errorf("unable to encode payments.sendPaymentForm#2d03522f: field credentials is nil")
	}
	if err := s.Credentials.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.sendPaymentForm#2d03522f: field credentials: %w", err)
	}
	if s.Flags.Has(2) {
		b.PutLong(s.TipAmount)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *PaymentsSendPaymentFormRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.sendPaymentForm#2d03522f to nil")
	}
	if err := b.ConsumeID(PaymentsSendPaymentFormRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.sendPaymentForm#2d03522f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *PaymentsSendPaymentFormRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.sendPaymentForm#2d03522f to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2d03522f: field flags: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2d03522f: field form_id: %w", err)
		}
		s.FormID = value
	}
	{
		value, err := DecodeInputInvoice(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2d03522f: field invoice: %w", err)
		}
		s.Invoice = value
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2d03522f: field requested_info_id: %w", err)
		}
		s.RequestedInfoID = value
	}
	if s.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2d03522f: field shipping_option_id: %w", err)
		}
		s.ShippingOptionID = value
	}
	{
		value, err := DecodeInputPaymentCredentials(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2d03522f: field credentials: %w", err)
		}
		s.Credentials = value
	}
	if s.Flags.Has(2) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode payments.sendPaymentForm#2d03522f: field tip_amount: %w", err)
		}
		s.TipAmount = value
	}
	return nil
}

// GetFormID returns value of FormID field.
func (s *PaymentsSendPaymentFormRequest) GetFormID() (value int64) {
	if s == nil {
		return
	}
	return s.FormID
}

// GetInvoice returns value of Invoice field.
func (s *PaymentsSendPaymentFormRequest) GetInvoice() (value InputInvoiceClass) {
	if s == nil {
		return
	}
	return s.Invoice
}

// SetRequestedInfoID sets value of RequestedInfoID conditional field.
func (s *PaymentsSendPaymentFormRequest) SetRequestedInfoID(value string) {
	s.Flags.Set(0)
	s.RequestedInfoID = value
}

// GetRequestedInfoID returns value of RequestedInfoID conditional field and
// boolean which is true if field was set.
func (s *PaymentsSendPaymentFormRequest) GetRequestedInfoID() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.RequestedInfoID, true
}

// SetShippingOptionID sets value of ShippingOptionID conditional field.
func (s *PaymentsSendPaymentFormRequest) SetShippingOptionID(value string) {
	s.Flags.Set(1)
	s.ShippingOptionID = value
}

// GetShippingOptionID returns value of ShippingOptionID conditional field and
// boolean which is true if field was set.
func (s *PaymentsSendPaymentFormRequest) GetShippingOptionID() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.ShippingOptionID, true
}

// GetCredentials returns value of Credentials field.
func (s *PaymentsSendPaymentFormRequest) GetCredentials() (value InputPaymentCredentialsClass) {
	if s == nil {
		return
	}
	return s.Credentials
}

// SetTipAmount sets value of TipAmount conditional field.
func (s *PaymentsSendPaymentFormRequest) SetTipAmount(value int64) {
	s.Flags.Set(2)
	s.TipAmount = value
}

// GetTipAmount returns value of TipAmount conditional field and
// boolean which is true if field was set.
func (s *PaymentsSendPaymentFormRequest) GetTipAmount() (value int64, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.TipAmount, true
}

// PaymentsSendPaymentForm invokes method payments.sendPaymentForm#2d03522f returning error if any.
// Send compiled payment form
//
// Possible errors:
//
//	400 MESSAGE_ID_INVALID: The provided message id is invalid.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/payments.sendPaymentForm for reference.
func (c *Client) PaymentsSendPaymentForm(ctx context.Context, request *PaymentsSendPaymentFormRequest) (PaymentsPaymentResultClass, error) {
	var result PaymentsPaymentResultBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.PaymentResult, nil
}
