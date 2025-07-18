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

// HelpCountry represents TL type `help.country#c3878e23`.
// Name, ISO code, localized name and phone codes/patterns of a specific country
//
// See https://core.telegram.org/constructor/help.country for reference.
type HelpCountry struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this country should not be shown in the list
	Hidden bool
	// ISO code of country
	ISO2 string
	// Name of the country in the country's language
	DefaultName string
	// Name of the country in the user's language, if different from the original name
	//
	// Use SetName and GetName helpers.
	Name string
	// Phone codes/patterns
	CountryCodes []HelpCountryCode
}

// HelpCountryTypeID is TL type id of HelpCountry.
const HelpCountryTypeID = 0xc3878e23

// Ensuring interfaces in compile-time for HelpCountry.
var (
	_ bin.Encoder     = &HelpCountry{}
	_ bin.Decoder     = &HelpCountry{}
	_ bin.BareEncoder = &HelpCountry{}
	_ bin.BareDecoder = &HelpCountry{}
)

func (c *HelpCountry) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Hidden == false) {
		return false
	}
	if !(c.ISO2 == "") {
		return false
	}
	if !(c.DefaultName == "") {
		return false
	}
	if !(c.Name == "") {
		return false
	}
	if !(c.CountryCodes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *HelpCountry) String() string {
	if c == nil {
		return "HelpCountry(nil)"
	}
	type Alias HelpCountry
	return fmt.Sprintf("HelpCountry%+v", Alias(*c))
}

// FillFrom fills HelpCountry from given interface.
func (c *HelpCountry) FillFrom(from interface {
	GetHidden() (value bool)
	GetISO2() (value string)
	GetDefaultName() (value string)
	GetName() (value string, ok bool)
	GetCountryCodes() (value []HelpCountryCode)
}) {
	c.Hidden = from.GetHidden()
	c.ISO2 = from.GetISO2()
	c.DefaultName = from.GetDefaultName()
	if val, ok := from.GetName(); ok {
		c.Name = val
	}

	c.CountryCodes = from.GetCountryCodes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpCountry) TypeID() uint32 {
	return HelpCountryTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpCountry) TypeName() string {
	return "help.country"
}

// TypeInfo returns info about TL type.
func (c *HelpCountry) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.country",
		ID:   HelpCountryTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hidden",
			SchemaName: "hidden",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "ISO2",
			SchemaName: "iso2",
		},
		{
			Name:       "DefaultName",
			SchemaName: "default_name",
		},
		{
			Name:       "Name",
			SchemaName: "name",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "CountryCodes",
			SchemaName: "country_codes",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *HelpCountry) SetFlags() {
	if !(c.Hidden == false) {
		c.Flags.Set(0)
	}
	if !(c.Name == "") {
		c.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (c *HelpCountry) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode help.country#c3878e23 as nil")
	}
	b.PutID(HelpCountryTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *HelpCountry) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode help.country#c3878e23 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.country#c3878e23: field flags: %w", err)
	}
	b.PutString(c.ISO2)
	b.PutString(c.DefaultName)
	if c.Flags.Has(1) {
		b.PutString(c.Name)
	}
	b.PutVectorHeader(len(c.CountryCodes))
	for idx, v := range c.CountryCodes {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.country#c3878e23: field country_codes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *HelpCountry) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode help.country#c3878e23 to nil")
	}
	if err := b.ConsumeID(HelpCountryTypeID); err != nil {
		return fmt.Errorf("unable to decode help.country#c3878e23: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *HelpCountry) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode help.country#c3878e23 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.country#c3878e23: field flags: %w", err)
		}
	}
	c.Hidden = c.Flags.Has(0)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.country#c3878e23: field iso2: %w", err)
		}
		c.ISO2 = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.country#c3878e23: field default_name: %w", err)
		}
		c.DefaultName = value
	}
	if c.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.country#c3878e23: field name: %w", err)
		}
		c.Name = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.country#c3878e23: field country_codes: %w", err)
		}

		if headerLen > 0 {
			c.CountryCodes = make([]HelpCountryCode, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value HelpCountryCode
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode help.country#c3878e23: field country_codes: %w", err)
			}
			c.CountryCodes = append(c.CountryCodes, value)
		}
	}
	return nil
}

// SetHidden sets value of Hidden conditional field.
func (c *HelpCountry) SetHidden(value bool) {
	if value {
		c.Flags.Set(0)
		c.Hidden = true
	} else {
		c.Flags.Unset(0)
		c.Hidden = false
	}
}

// GetHidden returns value of Hidden conditional field.
func (c *HelpCountry) GetHidden() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(0)
}

// GetISO2 returns value of ISO2 field.
func (c *HelpCountry) GetISO2() (value string) {
	if c == nil {
		return
	}
	return c.ISO2
}

// GetDefaultName returns value of DefaultName field.
func (c *HelpCountry) GetDefaultName() (value string) {
	if c == nil {
		return
	}
	return c.DefaultName
}

// SetName sets value of Name conditional field.
func (c *HelpCountry) SetName(value string) {
	c.Flags.Set(1)
	c.Name = value
}

// GetName returns value of Name conditional field and
// boolean which is true if field was set.
func (c *HelpCountry) GetName() (value string, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.Name, true
}

// GetCountryCodes returns value of CountryCodes field.
func (c *HelpCountry) GetCountryCodes() (value []HelpCountryCode) {
	if c == nil {
		return
	}
	return c.CountryCodes
}
