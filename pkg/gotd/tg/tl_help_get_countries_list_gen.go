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

// HelpGetCountriesListRequest represents TL type `help.getCountriesList#735787a8`.
// Get name, ISO code, localized name and phone codes/patterns of all available countries
//
// See https://core.telegram.org/method/help.getCountriesList for reference.
type HelpGetCountriesListRequest struct {
	// Language code of the current user
	LangCode string
	// Hash used for caching, for more info click here¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// HelpGetCountriesListRequestTypeID is TL type id of HelpGetCountriesListRequest.
const HelpGetCountriesListRequestTypeID = 0x735787a8

// Ensuring interfaces in compile-time for HelpGetCountriesListRequest.
var (
	_ bin.Encoder     = &HelpGetCountriesListRequest{}
	_ bin.Decoder     = &HelpGetCountriesListRequest{}
	_ bin.BareEncoder = &HelpGetCountriesListRequest{}
	_ bin.BareDecoder = &HelpGetCountriesListRequest{}
)

func (g *HelpGetCountriesListRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LangCode == "") {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetCountriesListRequest) String() string {
	if g == nil {
		return "HelpGetCountriesListRequest(nil)"
	}
	type Alias HelpGetCountriesListRequest
	return fmt.Sprintf("HelpGetCountriesListRequest%+v", Alias(*g))
}

// FillFrom fills HelpGetCountriesListRequest from given interface.
func (g *HelpGetCountriesListRequest) FillFrom(from interface {
	GetLangCode() (value string)
	GetHash() (value int)
}) {
	g.LangCode = from.GetLangCode()
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetCountriesListRequest) TypeID() uint32 {
	return HelpGetCountriesListRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetCountriesListRequest) TypeName() string {
	return "help.getCountriesList"
}

// TypeInfo returns info about TL type.
func (g *HelpGetCountriesListRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getCountriesList",
		ID:   HelpGetCountriesListRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetCountriesListRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getCountriesList#735787a8 as nil")
	}
	b.PutID(HelpGetCountriesListRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *HelpGetCountriesListRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getCountriesList#735787a8 as nil")
	}
	b.PutString(g.LangCode)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetCountriesListRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getCountriesList#735787a8 to nil")
	}
	if err := b.ConsumeID(HelpGetCountriesListRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getCountriesList#735787a8: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *HelpGetCountriesListRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getCountriesList#735787a8 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.getCountriesList#735787a8: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.getCountriesList#735787a8: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetLangCode returns value of LangCode field.
func (g *HelpGetCountriesListRequest) GetLangCode() (value string) {
	if g == nil {
		return
	}
	return g.LangCode
}

// GetHash returns value of Hash field.
func (g *HelpGetCountriesListRequest) GetHash() (value int) {
	if g == nil {
		return
	}
	return g.Hash
}

// HelpGetCountriesList invokes method help.getCountriesList#735787a8 returning error if any.
// Get name, ISO code, localized name and phone codes/patterns of all available countries
//
// See https://core.telegram.org/method/help.getCountriesList for reference.
func (c *Client) HelpGetCountriesList(ctx context.Context, request *HelpGetCountriesListRequest) (HelpCountriesListClass, error) {
	var result HelpCountriesListBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.CountriesList, nil
}
