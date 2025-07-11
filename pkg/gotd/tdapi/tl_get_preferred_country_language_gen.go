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

// GetPreferredCountryLanguageRequest represents TL type `getPreferredCountryLanguage#c862cbd6`.
type GetPreferredCountryLanguageRequest struct {
	// A two-letter ISO 3166-1 alpha-2 country code
	CountryCode string
}

// GetPreferredCountryLanguageRequestTypeID is TL type id of GetPreferredCountryLanguageRequest.
const GetPreferredCountryLanguageRequestTypeID = 0xc862cbd6

// Ensuring interfaces in compile-time for GetPreferredCountryLanguageRequest.
var (
	_ bin.Encoder     = &GetPreferredCountryLanguageRequest{}
	_ bin.Decoder     = &GetPreferredCountryLanguageRequest{}
	_ bin.BareEncoder = &GetPreferredCountryLanguageRequest{}
	_ bin.BareDecoder = &GetPreferredCountryLanguageRequest{}
)

func (g *GetPreferredCountryLanguageRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.CountryCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetPreferredCountryLanguageRequest) String() string {
	if g == nil {
		return "GetPreferredCountryLanguageRequest(nil)"
	}
	type Alias GetPreferredCountryLanguageRequest
	return fmt.Sprintf("GetPreferredCountryLanguageRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetPreferredCountryLanguageRequest) TypeID() uint32 {
	return GetPreferredCountryLanguageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetPreferredCountryLanguageRequest) TypeName() string {
	return "getPreferredCountryLanguage"
}

// TypeInfo returns info about TL type.
func (g *GetPreferredCountryLanguageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getPreferredCountryLanguage",
		ID:   GetPreferredCountryLanguageRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CountryCode",
			SchemaName: "country_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetPreferredCountryLanguageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPreferredCountryLanguage#c862cbd6 as nil")
	}
	b.PutID(GetPreferredCountryLanguageRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetPreferredCountryLanguageRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPreferredCountryLanguage#c862cbd6 as nil")
	}
	b.PutString(g.CountryCode)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetPreferredCountryLanguageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPreferredCountryLanguage#c862cbd6 to nil")
	}
	if err := b.ConsumeID(GetPreferredCountryLanguageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getPreferredCountryLanguage#c862cbd6: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetPreferredCountryLanguageRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPreferredCountryLanguage#c862cbd6 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getPreferredCountryLanguage#c862cbd6: field country_code: %w", err)
		}
		g.CountryCode = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetPreferredCountryLanguageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getPreferredCountryLanguage#c862cbd6 as nil")
	}
	b.ObjStart()
	b.PutID("getPreferredCountryLanguage")
	b.Comma()
	b.FieldStart("country_code")
	b.PutString(g.CountryCode)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetPreferredCountryLanguageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getPreferredCountryLanguage#c862cbd6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getPreferredCountryLanguage"); err != nil {
				return fmt.Errorf("unable to decode getPreferredCountryLanguage#c862cbd6: %w", err)
			}
		case "country_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getPreferredCountryLanguage#c862cbd6: field country_code: %w", err)
			}
			g.CountryCode = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCountryCode returns value of CountryCode field.
func (g *GetPreferredCountryLanguageRequest) GetCountryCode() (value string) {
	if g == nil {
		return
	}
	return g.CountryCode
}

// GetPreferredCountryLanguage invokes method getPreferredCountryLanguage#c862cbd6 returning error if any.
func (c *Client) GetPreferredCountryLanguage(ctx context.Context, countrycode string) (*Text, error) {
	var result Text

	request := &GetPreferredCountryLanguageRequest{
		CountryCode: countrycode,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
