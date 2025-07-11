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

// PaymentsValidatedRequestedInfo represents TL type `payments.validatedRequestedInfo#d1451883`.
//
// See https://core.telegram.org/constructor/payments.validatedRequestedInfo for reference.
type PaymentsValidatedRequestedInfo struct {
	// Flags field of PaymentsValidatedRequestedInfo.
	Flags bin.Fields
	// ID field of PaymentsValidatedRequestedInfo.
	//
	// Use SetID and GetID helpers.
	ID string
	// ShippingOptions field of PaymentsValidatedRequestedInfo.
	//
	// Use SetShippingOptions and GetShippingOptions helpers.
	ShippingOptions []ShippingOption
}

// PaymentsValidatedRequestedInfoTypeID is TL type id of PaymentsValidatedRequestedInfo.
const PaymentsValidatedRequestedInfoTypeID = 0xd1451883

// Ensuring interfaces in compile-time for PaymentsValidatedRequestedInfo.
var (
	_ bin.Encoder     = &PaymentsValidatedRequestedInfo{}
	_ bin.Decoder     = &PaymentsValidatedRequestedInfo{}
	_ bin.BareEncoder = &PaymentsValidatedRequestedInfo{}
	_ bin.BareDecoder = &PaymentsValidatedRequestedInfo{}
)

func (v *PaymentsValidatedRequestedInfo) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Flags.Zero()) {
		return false
	}
	if !(v.ID == "") {
		return false
	}
	if !(v.ShippingOptions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *PaymentsValidatedRequestedInfo) String() string {
	if v == nil {
		return "PaymentsValidatedRequestedInfo(nil)"
	}
	type Alias PaymentsValidatedRequestedInfo
	return fmt.Sprintf("PaymentsValidatedRequestedInfo%+v", Alias(*v))
}

// FillFrom fills PaymentsValidatedRequestedInfo from given interface.
func (v *PaymentsValidatedRequestedInfo) FillFrom(from interface {
	GetID() (value string, ok bool)
	GetShippingOptions() (value []ShippingOption, ok bool)
}) {
	if val, ok := from.GetID(); ok {
		v.ID = val
	}

	if val, ok := from.GetShippingOptions(); ok {
		v.ShippingOptions = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsValidatedRequestedInfo) TypeID() uint32 {
	return PaymentsValidatedRequestedInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsValidatedRequestedInfo) TypeName() string {
	return "payments.validatedRequestedInfo"
}

// TypeInfo returns info about TL type.
func (v *PaymentsValidatedRequestedInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.validatedRequestedInfo",
		ID:   PaymentsValidatedRequestedInfoTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
			Null:       !v.Flags.Has(0),
		},
		{
			Name:       "ShippingOptions",
			SchemaName: "shipping_options",
			Null:       !v.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (v *PaymentsValidatedRequestedInfo) SetFlags() {
	if !(v.ID == "") {
		v.Flags.Set(0)
	}
	if !(v.ShippingOptions == nil) {
		v.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (v *PaymentsValidatedRequestedInfo) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode payments.validatedRequestedInfo#d1451883 as nil")
	}
	b.PutID(PaymentsValidatedRequestedInfoTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *PaymentsValidatedRequestedInfo) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode payments.validatedRequestedInfo#d1451883 as nil")
	}
	v.SetFlags()
	if err := v.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.validatedRequestedInfo#d1451883: field flags: %w", err)
	}
	if v.Flags.Has(0) {
		b.PutString(v.ID)
	}
	if v.Flags.Has(1) {
		b.PutVectorHeader(len(v.ShippingOptions))
		for idx, v := range v.ShippingOptions {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode payments.validatedRequestedInfo#d1451883: field shipping_options element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (v *PaymentsValidatedRequestedInfo) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode payments.validatedRequestedInfo#d1451883 to nil")
	}
	if err := b.ConsumeID(PaymentsValidatedRequestedInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.validatedRequestedInfo#d1451883: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *PaymentsValidatedRequestedInfo) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode payments.validatedRequestedInfo#d1451883 to nil")
	}
	{
		if err := v.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.validatedRequestedInfo#d1451883: field flags: %w", err)
		}
	}
	if v.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.validatedRequestedInfo#d1451883: field id: %w", err)
		}
		v.ID = value
	}
	if v.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode payments.validatedRequestedInfo#d1451883: field shipping_options: %w", err)
		}

		if headerLen > 0 {
			v.ShippingOptions = make([]ShippingOption, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ShippingOption
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode payments.validatedRequestedInfo#d1451883: field shipping_options: %w", err)
			}
			v.ShippingOptions = append(v.ShippingOptions, value)
		}
	}
	return nil
}

// SetID sets value of ID conditional field.
func (v *PaymentsValidatedRequestedInfo) SetID(value string) {
	v.Flags.Set(0)
	v.ID = value
}

// GetID returns value of ID conditional field and
// boolean which is true if field was set.
func (v *PaymentsValidatedRequestedInfo) GetID() (value string, ok bool) {
	if v == nil {
		return
	}
	if !v.Flags.Has(0) {
		return value, false
	}
	return v.ID, true
}

// SetShippingOptions sets value of ShippingOptions conditional field.
func (v *PaymentsValidatedRequestedInfo) SetShippingOptions(value []ShippingOption) {
	v.Flags.Set(1)
	v.ShippingOptions = value
}

// GetShippingOptions returns value of ShippingOptions conditional field and
// boolean which is true if field was set.
func (v *PaymentsValidatedRequestedInfo) GetShippingOptions() (value []ShippingOption, ok bool) {
	if v == nil {
		return
	}
	if !v.Flags.Has(1) {
		return value, false
	}
	return v.ShippingOptions, true
}
