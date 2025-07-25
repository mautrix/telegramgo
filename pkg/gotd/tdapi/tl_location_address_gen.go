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

// LocationAddress represents TL type `locationAddress#a3dad322`.
type LocationAddress struct {
	// A two-letter ISO 3166-1 alpha-2 country code
	CountryCode string
	// State, if applicable; empty if unknown
	State string
	// City; empty if unknown
	City string
	// The address; empty if unknown
	Street string
}

// LocationAddressTypeID is TL type id of LocationAddress.
const LocationAddressTypeID = 0xa3dad322

// Ensuring interfaces in compile-time for LocationAddress.
var (
	_ bin.Encoder     = &LocationAddress{}
	_ bin.Decoder     = &LocationAddress{}
	_ bin.BareEncoder = &LocationAddress{}
	_ bin.BareDecoder = &LocationAddress{}
)

func (l *LocationAddress) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.CountryCode == "") {
		return false
	}
	if !(l.State == "") {
		return false
	}
	if !(l.City == "") {
		return false
	}
	if !(l.Street == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LocationAddress) String() string {
	if l == nil {
		return "LocationAddress(nil)"
	}
	type Alias LocationAddress
	return fmt.Sprintf("LocationAddress%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LocationAddress) TypeID() uint32 {
	return LocationAddressTypeID
}

// TypeName returns name of type in TL schema.
func (*LocationAddress) TypeName() string {
	return "locationAddress"
}

// TypeInfo returns info about TL type.
func (l *LocationAddress) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "locationAddress",
		ID:   LocationAddressTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CountryCode",
			SchemaName: "country_code",
		},
		{
			Name:       "State",
			SchemaName: "state",
		},
		{
			Name:       "City",
			SchemaName: "city",
		},
		{
			Name:       "Street",
			SchemaName: "street",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *LocationAddress) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode locationAddress#a3dad322 as nil")
	}
	b.PutID(LocationAddressTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LocationAddress) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode locationAddress#a3dad322 as nil")
	}
	b.PutString(l.CountryCode)
	b.PutString(l.State)
	b.PutString(l.City)
	b.PutString(l.Street)
	return nil
}

// Decode implements bin.Decoder.
func (l *LocationAddress) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode locationAddress#a3dad322 to nil")
	}
	if err := b.ConsumeID(LocationAddressTypeID); err != nil {
		return fmt.Errorf("unable to decode locationAddress#a3dad322: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LocationAddress) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode locationAddress#a3dad322 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode locationAddress#a3dad322: field country_code: %w", err)
		}
		l.CountryCode = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode locationAddress#a3dad322: field state: %w", err)
		}
		l.State = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode locationAddress#a3dad322: field city: %w", err)
		}
		l.City = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode locationAddress#a3dad322: field street: %w", err)
		}
		l.Street = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (l *LocationAddress) EncodeTDLibJSON(b tdjson.Encoder) error {
	if l == nil {
		return fmt.Errorf("can't encode locationAddress#a3dad322 as nil")
	}
	b.ObjStart()
	b.PutID("locationAddress")
	b.Comma()
	b.FieldStart("country_code")
	b.PutString(l.CountryCode)
	b.Comma()
	b.FieldStart("state")
	b.PutString(l.State)
	b.Comma()
	b.FieldStart("city")
	b.PutString(l.City)
	b.Comma()
	b.FieldStart("street")
	b.PutString(l.Street)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (l *LocationAddress) DecodeTDLibJSON(b tdjson.Decoder) error {
	if l == nil {
		return fmt.Errorf("can't decode locationAddress#a3dad322 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("locationAddress"); err != nil {
				return fmt.Errorf("unable to decode locationAddress#a3dad322: %w", err)
			}
		case "country_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode locationAddress#a3dad322: field country_code: %w", err)
			}
			l.CountryCode = value
		case "state":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode locationAddress#a3dad322: field state: %w", err)
			}
			l.State = value
		case "city":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode locationAddress#a3dad322: field city: %w", err)
			}
			l.City = value
		case "street":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode locationAddress#a3dad322: field street: %w", err)
			}
			l.Street = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCountryCode returns value of CountryCode field.
func (l *LocationAddress) GetCountryCode() (value string) {
	if l == nil {
		return
	}
	return l.CountryCode
}

// GetState returns value of State field.
func (l *LocationAddress) GetState() (value string) {
	if l == nil {
		return
	}
	return l.State
}

// GetCity returns value of City field.
func (l *LocationAddress) GetCity() (value string) {
	if l == nil {
		return
	}
	return l.City
}

// GetStreet returns value of Street field.
func (l *LocationAddress) GetStreet() (value string) {
	if l == nil {
		return
	}
	return l.Street
}
