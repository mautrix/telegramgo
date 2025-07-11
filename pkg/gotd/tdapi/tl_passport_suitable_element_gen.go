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

// PassportSuitableElement represents TL type `passportSuitableElement#d0f8831c`.
type PassportSuitableElement struct {
	// Type of the element
	Type PassportElementTypeClass
	// True, if a selfie is required with the identity document
	IsSelfieRequired bool
	// True, if a certified English translation is required with the document
	IsTranslationRequired bool
	// True, if personal details must include the user's name in the language of their
	// country of residence
	IsNativeNameRequired bool
}

// PassportSuitableElementTypeID is TL type id of PassportSuitableElement.
const PassportSuitableElementTypeID = 0xd0f8831c

// Ensuring interfaces in compile-time for PassportSuitableElement.
var (
	_ bin.Encoder     = &PassportSuitableElement{}
	_ bin.Decoder     = &PassportSuitableElement{}
	_ bin.BareEncoder = &PassportSuitableElement{}
	_ bin.BareDecoder = &PassportSuitableElement{}
)

func (p *PassportSuitableElement) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Type == nil) {
		return false
	}
	if !(p.IsSelfieRequired == false) {
		return false
	}
	if !(p.IsTranslationRequired == false) {
		return false
	}
	if !(p.IsNativeNameRequired == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PassportSuitableElement) String() string {
	if p == nil {
		return "PassportSuitableElement(nil)"
	}
	type Alias PassportSuitableElement
	return fmt.Sprintf("PassportSuitableElement%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PassportSuitableElement) TypeID() uint32 {
	return PassportSuitableElementTypeID
}

// TypeName returns name of type in TL schema.
func (*PassportSuitableElement) TypeName() string {
	return "passportSuitableElement"
}

// TypeInfo returns info about TL type.
func (p *PassportSuitableElement) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "passportSuitableElement",
		ID:   PassportSuitableElementTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "IsSelfieRequired",
			SchemaName: "is_selfie_required",
		},
		{
			Name:       "IsTranslationRequired",
			SchemaName: "is_translation_required",
		},
		{
			Name:       "IsNativeNameRequired",
			SchemaName: "is_native_name_required",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PassportSuitableElement) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passportSuitableElement#d0f8831c as nil")
	}
	b.PutID(PassportSuitableElementTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PassportSuitableElement) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passportSuitableElement#d0f8831c as nil")
	}
	if p.Type == nil {
		return fmt.Errorf("unable to encode passportSuitableElement#d0f8831c: field type is nil")
	}
	if err := p.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode passportSuitableElement#d0f8831c: field type: %w", err)
	}
	b.PutBool(p.IsSelfieRequired)
	b.PutBool(p.IsTranslationRequired)
	b.PutBool(p.IsNativeNameRequired)
	return nil
}

// Decode implements bin.Decoder.
func (p *PassportSuitableElement) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passportSuitableElement#d0f8831c to nil")
	}
	if err := b.ConsumeID(PassportSuitableElementTypeID); err != nil {
		return fmt.Errorf("unable to decode passportSuitableElement#d0f8831c: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PassportSuitableElement) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passportSuitableElement#d0f8831c to nil")
	}
	{
		value, err := DecodePassportElementType(b)
		if err != nil {
			return fmt.Errorf("unable to decode passportSuitableElement#d0f8831c: field type: %w", err)
		}
		p.Type = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode passportSuitableElement#d0f8831c: field is_selfie_required: %w", err)
		}
		p.IsSelfieRequired = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode passportSuitableElement#d0f8831c: field is_translation_required: %w", err)
		}
		p.IsTranslationRequired = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode passportSuitableElement#d0f8831c: field is_native_name_required: %w", err)
		}
		p.IsNativeNameRequired = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PassportSuitableElement) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode passportSuitableElement#d0f8831c as nil")
	}
	b.ObjStart()
	b.PutID("passportSuitableElement")
	b.Comma()
	b.FieldStart("type")
	if p.Type == nil {
		return fmt.Errorf("unable to encode passportSuitableElement#d0f8831c: field type is nil")
	}
	if err := p.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode passportSuitableElement#d0f8831c: field type: %w", err)
	}
	b.Comma()
	b.FieldStart("is_selfie_required")
	b.PutBool(p.IsSelfieRequired)
	b.Comma()
	b.FieldStart("is_translation_required")
	b.PutBool(p.IsTranslationRequired)
	b.Comma()
	b.FieldStart("is_native_name_required")
	b.PutBool(p.IsNativeNameRequired)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PassportSuitableElement) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode passportSuitableElement#d0f8831c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("passportSuitableElement"); err != nil {
				return fmt.Errorf("unable to decode passportSuitableElement#d0f8831c: %w", err)
			}
		case "type":
			value, err := DecodeTDLibJSONPassportElementType(b)
			if err != nil {
				return fmt.Errorf("unable to decode passportSuitableElement#d0f8831c: field type: %w", err)
			}
			p.Type = value
		case "is_selfie_required":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode passportSuitableElement#d0f8831c: field is_selfie_required: %w", err)
			}
			p.IsSelfieRequired = value
		case "is_translation_required":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode passportSuitableElement#d0f8831c: field is_translation_required: %w", err)
			}
			p.IsTranslationRequired = value
		case "is_native_name_required":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode passportSuitableElement#d0f8831c: field is_native_name_required: %w", err)
			}
			p.IsNativeNameRequired = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetType returns value of Type field.
func (p *PassportSuitableElement) GetType() (value PassportElementTypeClass) {
	if p == nil {
		return
	}
	return p.Type
}

// GetIsSelfieRequired returns value of IsSelfieRequired field.
func (p *PassportSuitableElement) GetIsSelfieRequired() (value bool) {
	if p == nil {
		return
	}
	return p.IsSelfieRequired
}

// GetIsTranslationRequired returns value of IsTranslationRequired field.
func (p *PassportSuitableElement) GetIsTranslationRequired() (value bool) {
	if p == nil {
		return
	}
	return p.IsTranslationRequired
}

// GetIsNativeNameRequired returns value of IsNativeNameRequired field.
func (p *PassportSuitableElement) GetIsNativeNameRequired() (value bool) {
	if p == nil {
		return
	}
	return p.IsNativeNameRequired
}
