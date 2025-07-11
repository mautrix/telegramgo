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

// Invoice represents TL type `invoice#49ee584`.
// Invoice
//
// See https://core.telegram.org/constructor/invoice for reference.
type Invoice struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Test invoice
	Test bool
	// Set this flag if you require the user's full name to complete the order
	NameRequested bool
	// Set this flag if you require the user's phone number to complete the order
	PhoneRequested bool
	// Set this flag if you require the user's email address to complete the order
	EmailRequested bool
	// Set this flag if you require the user's shipping address to complete the order
	ShippingAddressRequested bool
	// Set this flag if the final price depends on the shipping method
	Flexible bool
	// Set this flag if user's phone number should be sent to provider
	PhoneToProvider bool
	// Set this flag if user's email address should be sent to provider
	EmailToProvider bool
	// Whether this is a recurring payment
	Recurring bool
	// Three-letter ISO 4217 currency¹ code, or XTR for Telegram Stars².
	//
	// Links:
	//  1) https://core.telegram.org/bots/payments#supported-currencies
	//  2) https://core.telegram.org/api/stars
	Currency string
	// Price breakdown, a list of components (e.g. product price, tax, discount, delivery
	// cost, delivery tax, bonus, etc.)
	Prices []LabeledPrice
	// The maximum accepted amount for tips in the smallest units of the currency (integer,
	// not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp
	// parameter in currencies.json¹, it shows the number of digits past the decimal point
	// for each currency (2 for the majority of currencies).
	//
	// Links:
	//  1) https://core.telegram.org/bots/payments/currencies.json
	//
	// Use SetMaxTipAmount and GetMaxTipAmount helpers.
	MaxTipAmount int64
	// A vector of suggested amounts of tips in the smallest units of the currency (integer,
	// not float/double). At most 4 suggested tip amounts can be specified. The suggested tip
	// amounts must be positive, passed in a strictly increased order and must not exceed
	// max_tip_amount.
	//
	// Use SetSuggestedTipAmounts and GetSuggestedTipAmounts helpers.
	SuggestedTipAmounts []int64
	// Terms of service URL
	//
	// Use SetTermsURL and GetTermsURL helpers.
	TermsURL string
	// SubscriptionPeriod field of Invoice.
	//
	// Use SetSubscriptionPeriod and GetSubscriptionPeriod helpers.
	SubscriptionPeriod int
}

// InvoiceTypeID is TL type id of Invoice.
const InvoiceTypeID = 0x49ee584

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
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Test == false) {
		return false
	}
	if !(i.NameRequested == false) {
		return false
	}
	if !(i.PhoneRequested == false) {
		return false
	}
	if !(i.EmailRequested == false) {
		return false
	}
	if !(i.ShippingAddressRequested == false) {
		return false
	}
	if !(i.Flexible == false) {
		return false
	}
	if !(i.PhoneToProvider == false) {
		return false
	}
	if !(i.EmailToProvider == false) {
		return false
	}
	if !(i.Recurring == false) {
		return false
	}
	if !(i.Currency == "") {
		return false
	}
	if !(i.Prices == nil) {
		return false
	}
	if !(i.MaxTipAmount == 0) {
		return false
	}
	if !(i.SuggestedTipAmounts == nil) {
		return false
	}
	if !(i.TermsURL == "") {
		return false
	}
	if !(i.SubscriptionPeriod == 0) {
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

// FillFrom fills Invoice from given interface.
func (i *Invoice) FillFrom(from interface {
	GetTest() (value bool)
	GetNameRequested() (value bool)
	GetPhoneRequested() (value bool)
	GetEmailRequested() (value bool)
	GetShippingAddressRequested() (value bool)
	GetFlexible() (value bool)
	GetPhoneToProvider() (value bool)
	GetEmailToProvider() (value bool)
	GetRecurring() (value bool)
	GetCurrency() (value string)
	GetPrices() (value []LabeledPrice)
	GetMaxTipAmount() (value int64, ok bool)
	GetSuggestedTipAmounts() (value []int64, ok bool)
	GetTermsURL() (value string, ok bool)
	GetSubscriptionPeriod() (value int, ok bool)
}) {
	i.Test = from.GetTest()
	i.NameRequested = from.GetNameRequested()
	i.PhoneRequested = from.GetPhoneRequested()
	i.EmailRequested = from.GetEmailRequested()
	i.ShippingAddressRequested = from.GetShippingAddressRequested()
	i.Flexible = from.GetFlexible()
	i.PhoneToProvider = from.GetPhoneToProvider()
	i.EmailToProvider = from.GetEmailToProvider()
	i.Recurring = from.GetRecurring()
	i.Currency = from.GetCurrency()
	i.Prices = from.GetPrices()
	if val, ok := from.GetMaxTipAmount(); ok {
		i.MaxTipAmount = val
	}

	if val, ok := from.GetSuggestedTipAmounts(); ok {
		i.SuggestedTipAmounts = val
	}

	if val, ok := from.GetTermsURL(); ok {
		i.TermsURL = val
	}

	if val, ok := from.GetSubscriptionPeriod(); ok {
		i.SubscriptionPeriod = val
	}

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
			Name:       "Test",
			SchemaName: "test",
			Null:       !i.Flags.Has(0),
		},
		{
			Name:       "NameRequested",
			SchemaName: "name_requested",
			Null:       !i.Flags.Has(1),
		},
		{
			Name:       "PhoneRequested",
			SchemaName: "phone_requested",
			Null:       !i.Flags.Has(2),
		},
		{
			Name:       "EmailRequested",
			SchemaName: "email_requested",
			Null:       !i.Flags.Has(3),
		},
		{
			Name:       "ShippingAddressRequested",
			SchemaName: "shipping_address_requested",
			Null:       !i.Flags.Has(4),
		},
		{
			Name:       "Flexible",
			SchemaName: "flexible",
			Null:       !i.Flags.Has(5),
		},
		{
			Name:       "PhoneToProvider",
			SchemaName: "phone_to_provider",
			Null:       !i.Flags.Has(6),
		},
		{
			Name:       "EmailToProvider",
			SchemaName: "email_to_provider",
			Null:       !i.Flags.Has(7),
		},
		{
			Name:       "Recurring",
			SchemaName: "recurring",
			Null:       !i.Flags.Has(9),
		},
		{
			Name:       "Currency",
			SchemaName: "currency",
		},
		{
			Name:       "Prices",
			SchemaName: "prices",
		},
		{
			Name:       "MaxTipAmount",
			SchemaName: "max_tip_amount",
			Null:       !i.Flags.Has(8),
		},
		{
			Name:       "SuggestedTipAmounts",
			SchemaName: "suggested_tip_amounts",
			Null:       !i.Flags.Has(8),
		},
		{
			Name:       "TermsURL",
			SchemaName: "terms_url",
			Null:       !i.Flags.Has(10),
		},
		{
			Name:       "SubscriptionPeriod",
			SchemaName: "subscription_period",
			Null:       !i.Flags.Has(11),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (i *Invoice) SetFlags() {
	if !(i.Test == false) {
		i.Flags.Set(0)
	}
	if !(i.NameRequested == false) {
		i.Flags.Set(1)
	}
	if !(i.PhoneRequested == false) {
		i.Flags.Set(2)
	}
	if !(i.EmailRequested == false) {
		i.Flags.Set(3)
	}
	if !(i.ShippingAddressRequested == false) {
		i.Flags.Set(4)
	}
	if !(i.Flexible == false) {
		i.Flags.Set(5)
	}
	if !(i.PhoneToProvider == false) {
		i.Flags.Set(6)
	}
	if !(i.EmailToProvider == false) {
		i.Flags.Set(7)
	}
	if !(i.Recurring == false) {
		i.Flags.Set(9)
	}
	if !(i.MaxTipAmount == 0) {
		i.Flags.Set(8)
	}
	if !(i.SuggestedTipAmounts == nil) {
		i.Flags.Set(8)
	}
	if !(i.TermsURL == "") {
		i.Flags.Set(10)
	}
	if !(i.SubscriptionPeriod == 0) {
		i.Flags.Set(11)
	}
}

// Encode implements bin.Encoder.
func (i *Invoice) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invoice#49ee584 as nil")
	}
	b.PutID(InvoiceTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *Invoice) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invoice#49ee584 as nil")
	}
	i.SetFlags()
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode invoice#49ee584: field flags: %w", err)
	}
	b.PutString(i.Currency)
	b.PutVectorHeader(len(i.Prices))
	for idx, v := range i.Prices {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode invoice#49ee584: field prices element with index %d: %w", idx, err)
		}
	}
	if i.Flags.Has(8) {
		b.PutLong(i.MaxTipAmount)
	}
	if i.Flags.Has(8) {
		b.PutVectorHeader(len(i.SuggestedTipAmounts))
		for _, v := range i.SuggestedTipAmounts {
			b.PutLong(v)
		}
	}
	if i.Flags.Has(10) {
		b.PutString(i.TermsURL)
	}
	if i.Flags.Has(11) {
		b.PutInt(i.SubscriptionPeriod)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *Invoice) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invoice#49ee584 to nil")
	}
	if err := b.ConsumeID(InvoiceTypeID); err != nil {
		return fmt.Errorf("unable to decode invoice#49ee584: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *Invoice) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invoice#49ee584 to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode invoice#49ee584: field flags: %w", err)
		}
	}
	i.Test = i.Flags.Has(0)
	i.NameRequested = i.Flags.Has(1)
	i.PhoneRequested = i.Flags.Has(2)
	i.EmailRequested = i.Flags.Has(3)
	i.ShippingAddressRequested = i.Flags.Has(4)
	i.Flexible = i.Flags.Has(5)
	i.PhoneToProvider = i.Flags.Has(6)
	i.EmailToProvider = i.Flags.Has(7)
	i.Recurring = i.Flags.Has(9)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#49ee584: field currency: %w", err)
		}
		i.Currency = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#49ee584: field prices: %w", err)
		}

		if headerLen > 0 {
			i.Prices = make([]LabeledPrice, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value LabeledPrice
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode invoice#49ee584: field prices: %w", err)
			}
			i.Prices = append(i.Prices, value)
		}
	}
	if i.Flags.Has(8) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#49ee584: field max_tip_amount: %w", err)
		}
		i.MaxTipAmount = value
	}
	if i.Flags.Has(8) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#49ee584: field suggested_tip_amounts: %w", err)
		}

		if headerLen > 0 {
			i.SuggestedTipAmounts = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode invoice#49ee584: field suggested_tip_amounts: %w", err)
			}
			i.SuggestedTipAmounts = append(i.SuggestedTipAmounts, value)
		}
	}
	if i.Flags.Has(10) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#49ee584: field terms_url: %w", err)
		}
		i.TermsURL = value
	}
	if i.Flags.Has(11) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#49ee584: field subscription_period: %w", err)
		}
		i.SubscriptionPeriod = value
	}
	return nil
}

// SetTest sets value of Test conditional field.
func (i *Invoice) SetTest(value bool) {
	if value {
		i.Flags.Set(0)
		i.Test = true
	} else {
		i.Flags.Unset(0)
		i.Test = false
	}
}

// GetTest returns value of Test conditional field.
func (i *Invoice) GetTest() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(0)
}

// SetNameRequested sets value of NameRequested conditional field.
func (i *Invoice) SetNameRequested(value bool) {
	if value {
		i.Flags.Set(1)
		i.NameRequested = true
	} else {
		i.Flags.Unset(1)
		i.NameRequested = false
	}
}

// GetNameRequested returns value of NameRequested conditional field.
func (i *Invoice) GetNameRequested() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(1)
}

// SetPhoneRequested sets value of PhoneRequested conditional field.
func (i *Invoice) SetPhoneRequested(value bool) {
	if value {
		i.Flags.Set(2)
		i.PhoneRequested = true
	} else {
		i.Flags.Unset(2)
		i.PhoneRequested = false
	}
}

// GetPhoneRequested returns value of PhoneRequested conditional field.
func (i *Invoice) GetPhoneRequested() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(2)
}

// SetEmailRequested sets value of EmailRequested conditional field.
func (i *Invoice) SetEmailRequested(value bool) {
	if value {
		i.Flags.Set(3)
		i.EmailRequested = true
	} else {
		i.Flags.Unset(3)
		i.EmailRequested = false
	}
}

// GetEmailRequested returns value of EmailRequested conditional field.
func (i *Invoice) GetEmailRequested() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(3)
}

// SetShippingAddressRequested sets value of ShippingAddressRequested conditional field.
func (i *Invoice) SetShippingAddressRequested(value bool) {
	if value {
		i.Flags.Set(4)
		i.ShippingAddressRequested = true
	} else {
		i.Flags.Unset(4)
		i.ShippingAddressRequested = false
	}
}

// GetShippingAddressRequested returns value of ShippingAddressRequested conditional field.
func (i *Invoice) GetShippingAddressRequested() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(4)
}

// SetFlexible sets value of Flexible conditional field.
func (i *Invoice) SetFlexible(value bool) {
	if value {
		i.Flags.Set(5)
		i.Flexible = true
	} else {
		i.Flags.Unset(5)
		i.Flexible = false
	}
}

// GetFlexible returns value of Flexible conditional field.
func (i *Invoice) GetFlexible() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(5)
}

// SetPhoneToProvider sets value of PhoneToProvider conditional field.
func (i *Invoice) SetPhoneToProvider(value bool) {
	if value {
		i.Flags.Set(6)
		i.PhoneToProvider = true
	} else {
		i.Flags.Unset(6)
		i.PhoneToProvider = false
	}
}

// GetPhoneToProvider returns value of PhoneToProvider conditional field.
func (i *Invoice) GetPhoneToProvider() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(6)
}

// SetEmailToProvider sets value of EmailToProvider conditional field.
func (i *Invoice) SetEmailToProvider(value bool) {
	if value {
		i.Flags.Set(7)
		i.EmailToProvider = true
	} else {
		i.Flags.Unset(7)
		i.EmailToProvider = false
	}
}

// GetEmailToProvider returns value of EmailToProvider conditional field.
func (i *Invoice) GetEmailToProvider() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(7)
}

// SetRecurring sets value of Recurring conditional field.
func (i *Invoice) SetRecurring(value bool) {
	if value {
		i.Flags.Set(9)
		i.Recurring = true
	} else {
		i.Flags.Unset(9)
		i.Recurring = false
	}
}

// GetRecurring returns value of Recurring conditional field.
func (i *Invoice) GetRecurring() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(9)
}

// GetCurrency returns value of Currency field.
func (i *Invoice) GetCurrency() (value string) {
	if i == nil {
		return
	}
	return i.Currency
}

// GetPrices returns value of Prices field.
func (i *Invoice) GetPrices() (value []LabeledPrice) {
	if i == nil {
		return
	}
	return i.Prices
}

// SetMaxTipAmount sets value of MaxTipAmount conditional field.
func (i *Invoice) SetMaxTipAmount(value int64) {
	i.Flags.Set(8)
	i.MaxTipAmount = value
}

// GetMaxTipAmount returns value of MaxTipAmount conditional field and
// boolean which is true if field was set.
func (i *Invoice) GetMaxTipAmount() (value int64, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(8) {
		return value, false
	}
	return i.MaxTipAmount, true
}

// SetSuggestedTipAmounts sets value of SuggestedTipAmounts conditional field.
func (i *Invoice) SetSuggestedTipAmounts(value []int64) {
	i.Flags.Set(8)
	i.SuggestedTipAmounts = value
}

// GetSuggestedTipAmounts returns value of SuggestedTipAmounts conditional field and
// boolean which is true if field was set.
func (i *Invoice) GetSuggestedTipAmounts() (value []int64, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(8) {
		return value, false
	}
	return i.SuggestedTipAmounts, true
}

// SetTermsURL sets value of TermsURL conditional field.
func (i *Invoice) SetTermsURL(value string) {
	i.Flags.Set(10)
	i.TermsURL = value
}

// GetTermsURL returns value of TermsURL conditional field and
// boolean which is true if field was set.
func (i *Invoice) GetTermsURL() (value string, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(10) {
		return value, false
	}
	return i.TermsURL, true
}

// SetSubscriptionPeriod sets value of SubscriptionPeriod conditional field.
func (i *Invoice) SetSubscriptionPeriod(value int) {
	i.Flags.Set(11)
	i.SubscriptionPeriod = value
}

// GetSubscriptionPeriod returns value of SubscriptionPeriod conditional field and
// boolean which is true if field was set.
func (i *Invoice) GetSubscriptionPeriod() (value int, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(11) {
		return value, false
	}
	return i.SubscriptionPeriod, true
}
