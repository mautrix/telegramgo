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

// InlineQueryResultsButtonTypeStartBot represents TL type `inlineQueryResultsButtonTypeStartBot#fe9af0d5`.
type InlineQueryResultsButtonTypeStartBot struct {
	// The parameter for the bot start message
	Parameter string
}

// InlineQueryResultsButtonTypeStartBotTypeID is TL type id of InlineQueryResultsButtonTypeStartBot.
const InlineQueryResultsButtonTypeStartBotTypeID = 0xfe9af0d5

// construct implements constructor of InlineQueryResultsButtonTypeClass.
func (i InlineQueryResultsButtonTypeStartBot) construct() InlineQueryResultsButtonTypeClass {
	return &i
}

// Ensuring interfaces in compile-time for InlineQueryResultsButtonTypeStartBot.
var (
	_ bin.Encoder     = &InlineQueryResultsButtonTypeStartBot{}
	_ bin.Decoder     = &InlineQueryResultsButtonTypeStartBot{}
	_ bin.BareEncoder = &InlineQueryResultsButtonTypeStartBot{}
	_ bin.BareDecoder = &InlineQueryResultsButtonTypeStartBot{}

	_ InlineQueryResultsButtonTypeClass = &InlineQueryResultsButtonTypeStartBot{}
)

func (i *InlineQueryResultsButtonTypeStartBot) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Parameter == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineQueryResultsButtonTypeStartBot) String() string {
	if i == nil {
		return "InlineQueryResultsButtonTypeStartBot(nil)"
	}
	type Alias InlineQueryResultsButtonTypeStartBot
	return fmt.Sprintf("InlineQueryResultsButtonTypeStartBot%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineQueryResultsButtonTypeStartBot) TypeID() uint32 {
	return InlineQueryResultsButtonTypeStartBotTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineQueryResultsButtonTypeStartBot) TypeName() string {
	return "inlineQueryResultsButtonTypeStartBot"
}

// TypeInfo returns info about TL type.
func (i *InlineQueryResultsButtonTypeStartBot) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineQueryResultsButtonTypeStartBot",
		ID:   InlineQueryResultsButtonTypeStartBotTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Parameter",
			SchemaName: "parameter",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineQueryResultsButtonTypeStartBot) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryResultsButtonTypeStartBot#fe9af0d5 as nil")
	}
	b.PutID(InlineQueryResultsButtonTypeStartBotTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineQueryResultsButtonTypeStartBot) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryResultsButtonTypeStartBot#fe9af0d5 as nil")
	}
	b.PutString(i.Parameter)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineQueryResultsButtonTypeStartBot) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryResultsButtonTypeStartBot#fe9af0d5 to nil")
	}
	if err := b.ConsumeID(InlineQueryResultsButtonTypeStartBotTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineQueryResultsButtonTypeStartBot#fe9af0d5: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineQueryResultsButtonTypeStartBot) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryResultsButtonTypeStartBot#fe9af0d5 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineQueryResultsButtonTypeStartBot#fe9af0d5: field parameter: %w", err)
		}
		i.Parameter = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InlineQueryResultsButtonTypeStartBot) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryResultsButtonTypeStartBot#fe9af0d5 as nil")
	}
	b.ObjStart()
	b.PutID("inlineQueryResultsButtonTypeStartBot")
	b.Comma()
	b.FieldStart("parameter")
	b.PutString(i.Parameter)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InlineQueryResultsButtonTypeStartBot) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryResultsButtonTypeStartBot#fe9af0d5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inlineQueryResultsButtonTypeStartBot"); err != nil {
				return fmt.Errorf("unable to decode inlineQueryResultsButtonTypeStartBot#fe9af0d5: %w", err)
			}
		case "parameter":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inlineQueryResultsButtonTypeStartBot#fe9af0d5: field parameter: %w", err)
			}
			i.Parameter = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetParameter returns value of Parameter field.
func (i *InlineQueryResultsButtonTypeStartBot) GetParameter() (value string) {
	if i == nil {
		return
	}
	return i.Parameter
}

// InlineQueryResultsButtonTypeWebApp represents TL type `inlineQueryResultsButtonTypeWebApp#b8a16362`.
type InlineQueryResultsButtonTypeWebApp struct {
	// An HTTP URL to pass to getWebAppUrl
	URL string
}

// InlineQueryResultsButtonTypeWebAppTypeID is TL type id of InlineQueryResultsButtonTypeWebApp.
const InlineQueryResultsButtonTypeWebAppTypeID = 0xb8a16362

// construct implements constructor of InlineQueryResultsButtonTypeClass.
func (i InlineQueryResultsButtonTypeWebApp) construct() InlineQueryResultsButtonTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineQueryResultsButtonTypeWebApp.
var (
	_ bin.Encoder     = &InlineQueryResultsButtonTypeWebApp{}
	_ bin.Decoder     = &InlineQueryResultsButtonTypeWebApp{}
	_ bin.BareEncoder = &InlineQueryResultsButtonTypeWebApp{}
	_ bin.BareDecoder = &InlineQueryResultsButtonTypeWebApp{}

	_ InlineQueryResultsButtonTypeClass = &InlineQueryResultsButtonTypeWebApp{}
)

func (i *InlineQueryResultsButtonTypeWebApp) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineQueryResultsButtonTypeWebApp) String() string {
	if i == nil {
		return "InlineQueryResultsButtonTypeWebApp(nil)"
	}
	type Alias InlineQueryResultsButtonTypeWebApp
	return fmt.Sprintf("InlineQueryResultsButtonTypeWebApp%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineQueryResultsButtonTypeWebApp) TypeID() uint32 {
	return InlineQueryResultsButtonTypeWebAppTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineQueryResultsButtonTypeWebApp) TypeName() string {
	return "inlineQueryResultsButtonTypeWebApp"
}

// TypeInfo returns info about TL type.
func (i *InlineQueryResultsButtonTypeWebApp) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineQueryResultsButtonTypeWebApp",
		ID:   InlineQueryResultsButtonTypeWebAppTypeID,
	}
	if i == nil {
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
func (i *InlineQueryResultsButtonTypeWebApp) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryResultsButtonTypeWebApp#b8a16362 as nil")
	}
	b.PutID(InlineQueryResultsButtonTypeWebAppTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineQueryResultsButtonTypeWebApp) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryResultsButtonTypeWebApp#b8a16362 as nil")
	}
	b.PutString(i.URL)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineQueryResultsButtonTypeWebApp) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryResultsButtonTypeWebApp#b8a16362 to nil")
	}
	if err := b.ConsumeID(InlineQueryResultsButtonTypeWebAppTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineQueryResultsButtonTypeWebApp#b8a16362: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineQueryResultsButtonTypeWebApp) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryResultsButtonTypeWebApp#b8a16362 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineQueryResultsButtonTypeWebApp#b8a16362: field url: %w", err)
		}
		i.URL = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InlineQueryResultsButtonTypeWebApp) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryResultsButtonTypeWebApp#b8a16362 as nil")
	}
	b.ObjStart()
	b.PutID("inlineQueryResultsButtonTypeWebApp")
	b.Comma()
	b.FieldStart("url")
	b.PutString(i.URL)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InlineQueryResultsButtonTypeWebApp) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryResultsButtonTypeWebApp#b8a16362 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inlineQueryResultsButtonTypeWebApp"); err != nil {
				return fmt.Errorf("unable to decode inlineQueryResultsButtonTypeWebApp#b8a16362: %w", err)
			}
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inlineQueryResultsButtonTypeWebApp#b8a16362: field url: %w", err)
			}
			i.URL = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetURL returns value of URL field.
func (i *InlineQueryResultsButtonTypeWebApp) GetURL() (value string) {
	if i == nil {
		return
	}
	return i.URL
}

// InlineQueryResultsButtonTypeClassName is schema name of InlineQueryResultsButtonTypeClass.
const InlineQueryResultsButtonTypeClassName = "InlineQueryResultsButtonType"

// InlineQueryResultsButtonTypeClass represents InlineQueryResultsButtonType generic type.
//
// Example:
//
//	g, err := tdapi.DecodeInlineQueryResultsButtonType(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.InlineQueryResultsButtonTypeStartBot: // inlineQueryResultsButtonTypeStartBot#fe9af0d5
//	case *tdapi.InlineQueryResultsButtonTypeWebApp: // inlineQueryResultsButtonTypeWebApp#b8a16362
//	default: panic(v)
//	}
type InlineQueryResultsButtonTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InlineQueryResultsButtonTypeClass

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

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeInlineQueryResultsButtonType implements binary de-serialization for InlineQueryResultsButtonTypeClass.
func DecodeInlineQueryResultsButtonType(buf *bin.Buffer) (InlineQueryResultsButtonTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InlineQueryResultsButtonTypeStartBotTypeID:
		// Decoding inlineQueryResultsButtonTypeStartBot#fe9af0d5.
		v := InlineQueryResultsButtonTypeStartBot{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineQueryResultsButtonTypeClass: %w", err)
		}
		return &v, nil
	case InlineQueryResultsButtonTypeWebAppTypeID:
		// Decoding inlineQueryResultsButtonTypeWebApp#b8a16362.
		v := InlineQueryResultsButtonTypeWebApp{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineQueryResultsButtonTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InlineQueryResultsButtonTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONInlineQueryResultsButtonType implements binary de-serialization for InlineQueryResultsButtonTypeClass.
func DecodeTDLibJSONInlineQueryResultsButtonType(buf tdjson.Decoder) (InlineQueryResultsButtonTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "inlineQueryResultsButtonTypeStartBot":
		// Decoding inlineQueryResultsButtonTypeStartBot#fe9af0d5.
		v := InlineQueryResultsButtonTypeStartBot{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineQueryResultsButtonTypeClass: %w", err)
		}
		return &v, nil
	case "inlineQueryResultsButtonTypeWebApp":
		// Decoding inlineQueryResultsButtonTypeWebApp#b8a16362.
		v := InlineQueryResultsButtonTypeWebApp{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineQueryResultsButtonTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InlineQueryResultsButtonTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// InlineQueryResultsButtonType boxes the InlineQueryResultsButtonTypeClass providing a helper.
type InlineQueryResultsButtonTypeBox struct {
	InlineQueryResultsButtonType InlineQueryResultsButtonTypeClass
}

// Decode implements bin.Decoder for InlineQueryResultsButtonTypeBox.
func (b *InlineQueryResultsButtonTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InlineQueryResultsButtonTypeBox to nil")
	}
	v, err := DecodeInlineQueryResultsButtonType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InlineQueryResultsButtonType = v
	return nil
}

// Encode implements bin.Encode for InlineQueryResultsButtonTypeBox.
func (b *InlineQueryResultsButtonTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InlineQueryResultsButtonType == nil {
		return fmt.Errorf("unable to encode InlineQueryResultsButtonTypeClass as nil")
	}
	return b.InlineQueryResultsButtonType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for InlineQueryResultsButtonTypeBox.
func (b *InlineQueryResultsButtonTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode InlineQueryResultsButtonTypeBox to nil")
	}
	v, err := DecodeTDLibJSONInlineQueryResultsButtonType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InlineQueryResultsButtonType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for InlineQueryResultsButtonTypeBox.
func (b *InlineQueryResultsButtonTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.InlineQueryResultsButtonType == nil {
		return fmt.Errorf("unable to encode InlineQueryResultsButtonTypeClass as nil")
	}
	return b.InlineQueryResultsButtonType.EncodeTDLibJSON(buf)
}
