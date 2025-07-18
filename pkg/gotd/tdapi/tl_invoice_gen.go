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

// Invoice represents TL type `invoice#67dc0e89`.
type Invoice struct {
	// ISO 4217 currency code
	Currency string
	// A list of objects used to calculate the total price of the product
	PriceParts []LabeledPricePart
	// The number of seconds between consecutive Telegram Star debiting for subscription
	// invoices; 0 if the invoice doesn't create subscription
	SubscriptionPeriod int32
	// The maximum allowed amount of tip in the smallest units of the currency
	MaxTipAmount int64
	// Suggested amounts of tip in the smallest units of the currency
	SuggestedTipAmounts []int64
	// An HTTP URL with terms of service for recurring payments. If non-empty, the invoice
	// payment will result in recurring payments and the user must accept the terms of
	// service before allowed to pay
	RecurringPaymentTermsOfServiceURL string
	// An HTTP URL with terms of service for non-recurring payments. If non-empty, then the
	// user must accept the terms of service before allowed to pay
	TermsOfServiceURL string
	// True, if the payment is a test payment
	IsTest bool
	// True, if the user's name is needed for payment
	NeedName bool
	// True, if the user's phone number is needed for payment
	NeedPhoneNumber bool
	// True, if the user's email address is needed for payment
	NeedEmailAddress bool
	// True, if the user's shipping address is needed for payment
	NeedShippingAddress bool
	// True, if the user's phone number will be sent to the provider
	SendPhoneNumberToProvider bool
	// True, if the user's email address will be sent to the provider
	SendEmailAddressToProvider bool
	// True, if the total price depends on the shipping method
	IsFlexible bool
}

// InvoiceTypeID is TL type id of Invoice.
const InvoiceTypeID = 0x67dc0e89

// Ensuring interfaces in compile-time for Invoice.
var (
	_ bin.Encoder     = &Invoice{}
	_ bin.Decoder     = &Invoice{}
	_ bin.BareEncoder = &Invoice{}
	_ bin.BareDecoder = &Invoice{}
)

func (i *Invoice) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Currency == "") {
		return false
	}
	if !(i.PriceParts == nil) {
		return false
	}
	if !(i.SubscriptionPeriod == 0) {
		return false
	}
	if !(i.MaxTipAmount == 0) {
		return false
	}
	if !(i.SuggestedTipAmounts == nil) {
		return false
	}
	if !(i.RecurringPaymentTermsOfServiceURL == "") {
		return false
	}
	if !(i.TermsOfServiceURL == "") {
		return false
	}
	if !(i.IsTest == false) {
		return false
	}
	if !(i.NeedName == false) {
		return false
	}
	if !(i.NeedPhoneNumber == false) {
		return false
	}
	if !(i.NeedEmailAddress == false) {
		return false
	}
	if !(i.NeedShippingAddress == false) {
		return false
	}
	if !(i.SendPhoneNumberToProvider == false) {
		return false
	}
	if !(i.SendEmailAddressToProvider == false) {
		return false
	}
	if !(i.IsFlexible == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *Invoice) String() string {
	if i == nil {
		return "Invoice(nil)"
	}
	type Alias Invoice
	return fmt.Sprintf("Invoice%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Invoice) TypeID() uint32 {
	return InvoiceTypeID
}

// TypeName returns name of type in TL schema.
func (*Invoice) TypeName() string {
	return "invoice"
}

// TypeInfo returns info about TL type.
func (i *Invoice) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "invoice",
		ID:   InvoiceTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Currency",
			SchemaName: "currency",
		},
		{
			Name:       "PriceParts",
			SchemaName: "price_parts",
		},
		{
			Name:       "SubscriptionPeriod",
			SchemaName: "subscription_period",
		},
		{
			Name:       "MaxTipAmount",
			SchemaName: "max_tip_amount",
		},
		{
			Name:       "SuggestedTipAmounts",
			SchemaName: "suggested_tip_amounts",
		},
		{
			Name:       "RecurringPaymentTermsOfServiceURL",
			SchemaName: "recurring_payment_terms_of_service_url",
		},
		{
			Name:       "TermsOfServiceURL",
			SchemaName: "terms_of_service_url",
		},
		{
			Name:       "IsTest",
			SchemaName: "is_test",
		},
		{
			Name:       "NeedName",
			SchemaName: "need_name",
		},
		{
			Name:       "NeedPhoneNumber",
			SchemaName: "need_phone_number",
		},
		{
			Name:       "NeedEmailAddress",
			SchemaName: "need_email_address",
		},
		{
			Name:       "NeedShippingAddress",
			SchemaName: "need_shipping_address",
		},
		{
			Name:       "SendPhoneNumberToProvider",
			SchemaName: "send_phone_number_to_provider",
		},
		{
			Name:       "SendEmailAddressToProvider",
			SchemaName: "send_email_address_to_provider",
		},
		{
			Name:       "IsFlexible",
			SchemaName: "is_flexible",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *Invoice) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invoice#67dc0e89 as nil")
	}
	b.PutID(InvoiceTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *Invoice) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invoice#67dc0e89 as nil")
	}
	b.PutString(i.Currency)
	b.PutInt(len(i.PriceParts))
	for idx, v := range i.PriceParts {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare invoice#67dc0e89: field price_parts element with index %d: %w", idx, err)
		}
	}
	b.PutInt32(i.SubscriptionPeriod)
	b.PutInt53(i.MaxTipAmount)
	b.PutInt(len(i.SuggestedTipAmounts))
	for _, v := range i.SuggestedTipAmounts {
		b.PutInt53(v)
	}
	b.PutString(i.RecurringPaymentTermsOfServiceURL)
	b.PutString(i.TermsOfServiceURL)
	b.PutBool(i.IsTest)
	b.PutBool(i.NeedName)
	b.PutBool(i.NeedPhoneNumber)
	b.PutBool(i.NeedEmailAddress)
	b.PutBool(i.NeedShippingAddress)
	b.PutBool(i.SendPhoneNumberToProvider)
	b.PutBool(i.SendEmailAddressToProvider)
	b.PutBool(i.IsFlexible)
	return nil
}

// Decode implements bin.Decoder.
func (i *Invoice) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invoice#67dc0e89 to nil")
	}
	if err := b.ConsumeID(InvoiceTypeID); err != nil {
		return fmt.Errorf("unable to decode invoice#67dc0e89: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *Invoice) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invoice#67dc0e89 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field currency: %w", err)
		}
		i.Currency = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field price_parts: %w", err)
		}

		if headerLen > 0 {
			i.PriceParts = make([]LabeledPricePart, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value LabeledPricePart
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare invoice#67dc0e89: field price_parts: %w", err)
			}
			i.PriceParts = append(i.PriceParts, value)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field subscription_period: %w", err)
		}
		i.SubscriptionPeriod = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field max_tip_amount: %w", err)
		}
		i.MaxTipAmount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field suggested_tip_amounts: %w", err)
		}

		if headerLen > 0 {
			i.SuggestedTipAmounts = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field suggested_tip_amounts: %w", err)
			}
			i.SuggestedTipAmounts = append(i.SuggestedTipAmounts, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field recurring_payment_terms_of_service_url: %w", err)
		}
		i.RecurringPaymentTermsOfServiceURL = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field terms_of_service_url: %w", err)
		}
		i.TermsOfServiceURL = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field is_test: %w", err)
		}
		i.IsTest = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field need_name: %w", err)
		}
		i.NeedName = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field need_phone_number: %w", err)
		}
		i.NeedPhoneNumber = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field need_email_address: %w", err)
		}
		i.NeedEmailAddress = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field need_shipping_address: %w", err)
		}
		i.NeedShippingAddress = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field send_phone_number_to_provider: %w", err)
		}
		i.SendPhoneNumberToProvider = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field send_email_address_to_provider: %w", err)
		}
		i.SendEmailAddressToProvider = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#67dc0e89: field is_flexible: %w", err)
		}
		i.IsFlexible = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *Invoice) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode invoice#67dc0e89 as nil")
	}
	b.ObjStart()
	b.PutID("invoice")
	b.Comma()
	b.FieldStart("currency")
	b.PutString(i.Currency)
	b.Comma()
	b.FieldStart("price_parts")
	b.ArrStart()
	for idx, v := range i.PriceParts {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode invoice#67dc0e89: field price_parts element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("subscription_period")
	b.PutInt32(i.SubscriptionPeriod)
	b.Comma()
	b.FieldStart("max_tip_amount")
	b.PutInt53(i.MaxTipAmount)
	b.Comma()
	b.FieldStart("suggested_tip_amounts")
	b.ArrStart()
	for _, v := range i.SuggestedTipAmounts {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("recurring_payment_terms_of_service_url")
	b.PutString(i.RecurringPaymentTermsOfServiceURL)
	b.Comma()
	b.FieldStart("terms_of_service_url")
	b.PutString(i.TermsOfServiceURL)
	b.Comma()
	b.FieldStart("is_test")
	b.PutBool(i.IsTest)
	b.Comma()
	b.FieldStart("need_name")
	b.PutBool(i.NeedName)
	b.Comma()
	b.FieldStart("need_phone_number")
	b.PutBool(i.NeedPhoneNumber)
	b.Comma()
	b.FieldStart("need_email_address")
	b.PutBool(i.NeedEmailAddress)
	b.Comma()
	b.FieldStart("need_shipping_address")
	b.PutBool(i.NeedShippingAddress)
	b.Comma()
	b.FieldStart("send_phone_number_to_provider")
	b.PutBool(i.SendPhoneNumberToProvider)
	b.Comma()
	b.FieldStart("send_email_address_to_provider")
	b.PutBool(i.SendEmailAddressToProvider)
	b.Comma()
	b.FieldStart("is_flexible")
	b.PutBool(i.IsFlexible)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *Invoice) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode invoice#67dc0e89 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("invoice"); err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: %w", err)
			}
		case "currency":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field currency: %w", err)
			}
			i.Currency = value
		case "price_parts":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value LabeledPricePart
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode invoice#67dc0e89: field price_parts: %w", err)
				}
				i.PriceParts = append(i.PriceParts, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field price_parts: %w", err)
			}
		case "subscription_period":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field subscription_period: %w", err)
			}
			i.SubscriptionPeriod = value
		case "max_tip_amount":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field max_tip_amount: %w", err)
			}
			i.MaxTipAmount = value
		case "suggested_tip_amounts":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode invoice#67dc0e89: field suggested_tip_amounts: %w", err)
				}
				i.SuggestedTipAmounts = append(i.SuggestedTipAmounts, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field suggested_tip_amounts: %w", err)
			}
		case "recurring_payment_terms_of_service_url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field recurring_payment_terms_of_service_url: %w", err)
			}
			i.RecurringPaymentTermsOfServiceURL = value
		case "terms_of_service_url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field terms_of_service_url: %w", err)
			}
			i.TermsOfServiceURL = value
		case "is_test":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field is_test: %w", err)
			}
			i.IsTest = value
		case "need_name":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field need_name: %w", err)
			}
			i.NeedName = value
		case "need_phone_number":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field need_phone_number: %w", err)
			}
			i.NeedPhoneNumber = value
		case "need_email_address":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field need_email_address: %w", err)
			}
			i.NeedEmailAddress = value
		case "need_shipping_address":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field need_shipping_address: %w", err)
			}
			i.NeedShippingAddress = value
		case "send_phone_number_to_provider":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field send_phone_number_to_provider: %w", err)
			}
			i.SendPhoneNumberToProvider = value
		case "send_email_address_to_provider":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field send_email_address_to_provider: %w", err)
			}
			i.SendEmailAddressToProvider = value
		case "is_flexible":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#67dc0e89: field is_flexible: %w", err)
			}
			i.IsFlexible = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCurrency returns value of Currency field.
func (i *Invoice) GetCurrency() (value string) {
	if i == nil {
		return
	}
	return i.Currency
}

// GetPriceParts returns value of PriceParts field.
func (i *Invoice) GetPriceParts() (value []LabeledPricePart) {
	if i == nil {
		return
	}
	return i.PriceParts
}

// GetSubscriptionPeriod returns value of SubscriptionPeriod field.
func (i *Invoice) GetSubscriptionPeriod() (value int32) {
	if i == nil {
		return
	}
	return i.SubscriptionPeriod
}

// GetMaxTipAmount returns value of MaxTipAmount field.
func (i *Invoice) GetMaxTipAmount() (value int64) {
	if i == nil {
		return
	}
	return i.MaxTipAmount
}

// GetSuggestedTipAmounts returns value of SuggestedTipAmounts field.
func (i *Invoice) GetSuggestedTipAmounts() (value []int64) {
	if i == nil {
		return
	}
	return i.SuggestedTipAmounts
}

// GetRecurringPaymentTermsOfServiceURL returns value of RecurringPaymentTermsOfServiceURL field.
func (i *Invoice) GetRecurringPaymentTermsOfServiceURL() (value string) {
	if i == nil {
		return
	}
	return i.RecurringPaymentTermsOfServiceURL
}

// GetTermsOfServiceURL returns value of TermsOfServiceURL field.
func (i *Invoice) GetTermsOfServiceURL() (value string) {
	if i == nil {
		return
	}
	return i.TermsOfServiceURL
}

// GetIsTest returns value of IsTest field.
func (i *Invoice) GetIsTest() (value bool) {
	if i == nil {
		return
	}
	return i.IsTest
}

// GetNeedName returns value of NeedName field.
func (i *Invoice) GetNeedName() (value bool) {
	if i == nil {
		return
	}
	return i.NeedName
}

// GetNeedPhoneNumber returns value of NeedPhoneNumber field.
func (i *Invoice) GetNeedPhoneNumber() (value bool) {
	if i == nil {
		return
	}
	return i.NeedPhoneNumber
}

// GetNeedEmailAddress returns value of NeedEmailAddress field.
func (i *Invoice) GetNeedEmailAddress() (value bool) {
	if i == nil {
		return
	}
	return i.NeedEmailAddress
}

// GetNeedShippingAddress returns value of NeedShippingAddress field.
func (i *Invoice) GetNeedShippingAddress() (value bool) {
	if i == nil {
		return
	}
	return i.NeedShippingAddress
}

// GetSendPhoneNumberToProvider returns value of SendPhoneNumberToProvider field.
func (i *Invoice) GetSendPhoneNumberToProvider() (value bool) {
	if i == nil {
		return
	}
	return i.SendPhoneNumberToProvider
}

// GetSendEmailAddressToProvider returns value of SendEmailAddressToProvider field.
func (i *Invoice) GetSendEmailAddressToProvider() (value bool) {
	if i == nil {
		return
	}
	return i.SendEmailAddressToProvider
}

// GetIsFlexible returns value of IsFlexible field.
func (i *Invoice) GetIsFlexible() (value bool) {
	if i == nil {
		return
	}
	return i.IsFlexible
}
