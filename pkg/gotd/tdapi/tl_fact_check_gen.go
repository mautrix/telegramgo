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

// FactCheck represents TL type `factCheck#c185f918`.
type FactCheck struct {
	// Text of the fact-check
	Text FormattedText
	// A two-letter ISO 3166-1 alpha-2 country code of the country for which the fact-check
	// is shown
	CountryCode string
}

// FactCheckTypeID is TL type id of FactCheck.
const FactCheckTypeID = 0xc185f918

// Ensuring interfaces in compile-time for FactCheck.
var (
	_ bin.Encoder     = &FactCheck{}
	_ bin.Decoder     = &FactCheck{}
	_ bin.BareEncoder = &FactCheck{}
	_ bin.BareDecoder = &FactCheck{}
)

func (f *FactCheck) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Text.Zero()) {
		return false
	}
	if !(f.CountryCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FactCheck) String() string {
	if f == nil {
		return "FactCheck(nil)"
	}
	type Alias FactCheck
	return fmt.Sprintf("FactCheck%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FactCheck) TypeID() uint32 {
	return FactCheckTypeID
}

// TypeName returns name of type in TL schema.
func (*FactCheck) TypeName() string {
	return "factCheck"
}

// TypeInfo returns info about TL type.
func (f *FactCheck) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "factCheck",
		ID:   FactCheckTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "CountryCode",
			SchemaName: "country_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *FactCheck) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode factCheck#c185f918 as nil")
	}
	b.PutID(FactCheckTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FactCheck) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode factCheck#c185f918 as nil")
	}
	if err := f.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode factCheck#c185f918: field text: %w", err)
	}
	b.PutString(f.CountryCode)
	return nil
}

// Decode implements bin.Decoder.
func (f *FactCheck) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode factCheck#c185f918 to nil")
	}
	if err := b.ConsumeID(FactCheckTypeID); err != nil {
		return fmt.Errorf("unable to decode factCheck#c185f918: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FactCheck) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode factCheck#c185f918 to nil")
	}
	{
		if err := f.Text.Decode(b); err != nil {
			return fmt.Errorf("unable to decode factCheck#c185f918: field text: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode factCheck#c185f918: field country_code: %w", err)
		}
		f.CountryCode = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *FactCheck) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode factCheck#c185f918 as nil")
	}
	b.ObjStart()
	b.PutID("factCheck")
	b.Comma()
	b.FieldStart("text")
	if err := f.Text.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode factCheck#c185f918: field text: %w", err)
	}
	b.Comma()
	b.FieldStart("country_code")
	b.PutString(f.CountryCode)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (f *FactCheck) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode factCheck#c185f918 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("factCheck"); err != nil {
				return fmt.Errorf("unable to decode factCheck#c185f918: %w", err)
			}
		case "text":
			if err := f.Text.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode factCheck#c185f918: field text: %w", err)
			}
		case "country_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode factCheck#c185f918: field country_code: %w", err)
			}
			f.CountryCode = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (f *FactCheck) GetText() (value FormattedText) {
	if f == nil {
		return
	}
	return f.Text
}

// GetCountryCode returns value of CountryCode field.
func (f *FactCheck) GetCountryCode() (value string) {
	if f == nil {
		return
	}
	return f.CountryCode
}
