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

// PaymentsStarsRevenueAdsAccountURL represents TL type `payments.starsRevenueAdsAccountUrl#394e7f21`.
// Contains a URL leading to a page where the user will be able to place ads for the
// channel/bot, paying using Telegram Stars¹.
//
// Links:
//  1. https://core.telegram.org/api/stars#paying-for-ads
//
// See https://core.telegram.org/constructor/payments.starsRevenueAdsAccountUrl for reference.
type PaymentsStarsRevenueAdsAccountURL struct {
	// URL to open.
	URL string
}

// PaymentsStarsRevenueAdsAccountURLTypeID is TL type id of PaymentsStarsRevenueAdsAccountURL.
const PaymentsStarsRevenueAdsAccountURLTypeID = 0x394e7f21

// Ensuring interfaces in compile-time for PaymentsStarsRevenueAdsAccountURL.
var (
	_ bin.Encoder     = &PaymentsStarsRevenueAdsAccountURL{}
	_ bin.Decoder     = &PaymentsStarsRevenueAdsAccountURL{}
	_ bin.BareEncoder = &PaymentsStarsRevenueAdsAccountURL{}
	_ bin.BareDecoder = &PaymentsStarsRevenueAdsAccountURL{}
)

func (s *PaymentsStarsRevenueAdsAccountURL) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PaymentsStarsRevenueAdsAccountURL) String() string {
	if s == nil {
		return "PaymentsStarsRevenueAdsAccountURL(nil)"
	}
	type Alias PaymentsStarsRevenueAdsAccountURL
	return fmt.Sprintf("PaymentsStarsRevenueAdsAccountURL%+v", Alias(*s))
}

// FillFrom fills PaymentsStarsRevenueAdsAccountURL from given interface.
func (s *PaymentsStarsRevenueAdsAccountURL) FillFrom(from interface {
	GetURL() (value string)
}) {
	s.URL = from.GetURL()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsStarsRevenueAdsAccountURL) TypeID() uint32 {
	return PaymentsStarsRevenueAdsAccountURLTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsStarsRevenueAdsAccountURL) TypeName() string {
	return "payments.starsRevenueAdsAccountUrl"
}

// TypeInfo returns info about TL type.
func (s *PaymentsStarsRevenueAdsAccountURL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.starsRevenueAdsAccountUrl",
		ID:   PaymentsStarsRevenueAdsAccountURLTypeID,
	}
	if s == nil {
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
func (s *PaymentsStarsRevenueAdsAccountURL) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.starsRevenueAdsAccountUrl#394e7f21 as nil")
	}
	b.PutID(PaymentsStarsRevenueAdsAccountURLTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *PaymentsStarsRevenueAdsAccountURL) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.starsRevenueAdsAccountUrl#394e7f21 as nil")
	}
	b.PutString(s.URL)
	return nil
}

// Decode implements bin.Decoder.
func (s *PaymentsStarsRevenueAdsAccountURL) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.starsRevenueAdsAccountUrl#394e7f21 to nil")
	}
	if err := b.ConsumeID(PaymentsStarsRevenueAdsAccountURLTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.starsRevenueAdsAccountUrl#394e7f21: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *PaymentsStarsRevenueAdsAccountURL) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.starsRevenueAdsAccountUrl#394e7f21 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.starsRevenueAdsAccountUrl#394e7f21: field url: %w", err)
		}
		s.URL = value
	}
	return nil
}

// GetURL returns value of URL field.
func (s *PaymentsStarsRevenueAdsAccountURL) GetURL() (value string) {
	if s == nil {
		return
	}
	return s.URL
}
