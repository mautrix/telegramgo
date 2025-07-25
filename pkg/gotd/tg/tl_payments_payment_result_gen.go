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

// PaymentsPaymentResult represents TL type `payments.paymentResult#4e5f810d`.
// Payment result
//
// See https://core.telegram.org/constructor/payments.paymentResult for reference.
type PaymentsPaymentResult struct {
	// Info about the payment
	Updates UpdatesClass
}

// PaymentsPaymentResultTypeID is TL type id of PaymentsPaymentResult.
const PaymentsPaymentResultTypeID = 0x4e5f810d

// construct implements constructor of PaymentsPaymentResultClass.
func (p PaymentsPaymentResult) construct() PaymentsPaymentResultClass { return &p }

// Ensuring interfaces in compile-time for PaymentsPaymentResult.
var (
	_ bin.Encoder     = &PaymentsPaymentResult{}
	_ bin.Decoder     = &PaymentsPaymentResult{}
	_ bin.BareEncoder = &PaymentsPaymentResult{}
	_ bin.BareDecoder = &PaymentsPaymentResult{}

	_ PaymentsPaymentResultClass = &PaymentsPaymentResult{}
)

func (p *PaymentsPaymentResult) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Updates == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentsPaymentResult) String() string {
	if p == nil {
		return "PaymentsPaymentResult(nil)"
	}
	type Alias PaymentsPaymentResult
	return fmt.Sprintf("PaymentsPaymentResult%+v", Alias(*p))
}

// FillFrom fills PaymentsPaymentResult from given interface.
func (p *PaymentsPaymentResult) FillFrom(from interface {
	GetUpdates() (value UpdatesClass)
}) {
	p.Updates = from.GetUpdates()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsPaymentResult) TypeID() uint32 {
	return PaymentsPaymentResultTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsPaymentResult) TypeName() string {
	return "payments.paymentResult"
}

// TypeInfo returns info about TL type.
func (p *PaymentsPaymentResult) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.paymentResult",
		ID:   PaymentsPaymentResultTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Updates",
			SchemaName: "updates",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PaymentsPaymentResult) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode payments.paymentResult#4e5f810d as nil")
	}
	b.PutID(PaymentsPaymentResultTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PaymentsPaymentResult) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode payments.paymentResult#4e5f810d as nil")
	}
	if p.Updates == nil {
		return fmt.Errorf("unable to encode payments.paymentResult#4e5f810d: field updates is nil")
	}
	if err := p.Updates.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.paymentResult#4e5f810d: field updates: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentsPaymentResult) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode payments.paymentResult#4e5f810d to nil")
	}
	if err := b.ConsumeID(PaymentsPaymentResultTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.paymentResult#4e5f810d: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PaymentsPaymentResult) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode payments.paymentResult#4e5f810d to nil")
	}
	{
		value, err := DecodeUpdates(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentResult#4e5f810d: field updates: %w", err)
		}
		p.Updates = value
	}
	return nil
}

// GetUpdates returns value of Updates field.
func (p *PaymentsPaymentResult) GetUpdates() (value UpdatesClass) {
	if p == nil {
		return
	}
	return p.Updates
}

// PaymentsPaymentVerificationNeeded represents TL type `payments.paymentVerificationNeeded#d8411139`.
// Payment was not successful, additional verification is needed
//
// See https://core.telegram.org/constructor/payments.paymentVerificationNeeded for reference.
type PaymentsPaymentVerificationNeeded struct {
	// URL for additional payment credentials verification
	URL string
}

// PaymentsPaymentVerificationNeededTypeID is TL type id of PaymentsPaymentVerificationNeeded.
const PaymentsPaymentVerificationNeededTypeID = 0xd8411139

// construct implements constructor of PaymentsPaymentResultClass.
func (p PaymentsPaymentVerificationNeeded) construct() PaymentsPaymentResultClass { return &p }

// Ensuring interfaces in compile-time for PaymentsPaymentVerificationNeeded.
var (
	_ bin.Encoder     = &PaymentsPaymentVerificationNeeded{}
	_ bin.Decoder     = &PaymentsPaymentVerificationNeeded{}
	_ bin.BareEncoder = &PaymentsPaymentVerificationNeeded{}
	_ bin.BareDecoder = &PaymentsPaymentVerificationNeeded{}

	_ PaymentsPaymentResultClass = &PaymentsPaymentVerificationNeeded{}
)

func (p *PaymentsPaymentVerificationNeeded) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentsPaymentVerificationNeeded) String() string {
	if p == nil {
		return "PaymentsPaymentVerificationNeeded(nil)"
	}
	type Alias PaymentsPaymentVerificationNeeded
	return fmt.Sprintf("PaymentsPaymentVerificationNeeded%+v", Alias(*p))
}

// FillFrom fills PaymentsPaymentVerificationNeeded from given interface.
func (p *PaymentsPaymentVerificationNeeded) FillFrom(from interface {
	GetURL() (value string)
}) {
	p.URL = from.GetURL()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsPaymentVerificationNeeded) TypeID() uint32 {
	return PaymentsPaymentVerificationNeededTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsPaymentVerificationNeeded) TypeName() string {
	return "payments.paymentVerificationNeeded"
}

// TypeInfo returns info about TL type.
func (p *PaymentsPaymentVerificationNeeded) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.paymentVerificationNeeded",
		ID:   PaymentsPaymentVerificationNeededTypeID,
	}
	if p == nil {
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
func (p *PaymentsPaymentVerificationNeeded) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode payments.paymentVerificationNeeded#d8411139 as nil")
	}
	b.PutID(PaymentsPaymentVerificationNeededTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PaymentsPaymentVerificationNeeded) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode payments.paymentVerificationNeeded#d8411139 as nil")
	}
	b.PutString(p.URL)
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentsPaymentVerificationNeeded) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode payments.paymentVerificationNeeded#d8411139 to nil")
	}
	if err := b.ConsumeID(PaymentsPaymentVerificationNeededTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.paymentVerificationNeeded#d8411139: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PaymentsPaymentVerificationNeeded) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode payments.paymentVerificationNeeded#d8411139 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentVerificationNeeded#d8411139: field url: %w", err)
		}
		p.URL = value
	}
	return nil
}

// GetURL returns value of URL field.
func (p *PaymentsPaymentVerificationNeeded) GetURL() (value string) {
	if p == nil {
		return
	}
	return p.URL
}

// PaymentsPaymentResultClassName is schema name of PaymentsPaymentResultClass.
const PaymentsPaymentResultClassName = "payments.PaymentResult"

// PaymentsPaymentResultClass represents payments.PaymentResult generic type.
//
// See https://core.telegram.org/type/payments.PaymentResult for reference.
//
// Example:
//
//	g, err := tg.DecodePaymentsPaymentResult(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.PaymentsPaymentResult: // payments.paymentResult#4e5f810d
//	case *tg.PaymentsPaymentVerificationNeeded: // payments.paymentVerificationNeeded#d8411139
//	default: panic(v)
//	}
type PaymentsPaymentResultClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PaymentsPaymentResultClass

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
}

// DecodePaymentsPaymentResult implements binary de-serialization for PaymentsPaymentResultClass.
func DecodePaymentsPaymentResult(buf *bin.Buffer) (PaymentsPaymentResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PaymentsPaymentResultTypeID:
		// Decoding payments.paymentResult#4e5f810d.
		v := PaymentsPaymentResult{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PaymentsPaymentResultClass: %w", err)
		}
		return &v, nil
	case PaymentsPaymentVerificationNeededTypeID:
		// Decoding payments.paymentVerificationNeeded#d8411139.
		v := PaymentsPaymentVerificationNeeded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PaymentsPaymentResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PaymentsPaymentResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// PaymentsPaymentResult boxes the PaymentsPaymentResultClass providing a helper.
type PaymentsPaymentResultBox struct {
	PaymentResult PaymentsPaymentResultClass
}

// Decode implements bin.Decoder for PaymentsPaymentResultBox.
func (b *PaymentsPaymentResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PaymentsPaymentResultBox to nil")
	}
	v, err := DecodePaymentsPaymentResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PaymentResult = v
	return nil
}

// Encode implements bin.Encode for PaymentsPaymentResultBox.
func (b *PaymentsPaymentResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PaymentResult == nil {
		return fmt.Errorf("unable to encode PaymentsPaymentResultClass as nil")
	}
	return b.PaymentResult.Encode(buf)
}
