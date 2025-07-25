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

// JSONValueNull represents TL type `jsonValueNull#fa76e0cd`.
type JSONValueNull struct {
}

// JSONValueNullTypeID is TL type id of JSONValueNull.
const JSONValueNullTypeID = 0xfa76e0cd

// construct implements constructor of JSONValueClass.
func (j JSONValueNull) construct() JSONValueClass { return &j }

// Ensuring interfaces in compile-time for JSONValueNull.
var (
	_ bin.Encoder     = &JSONValueNull{}
	_ bin.Decoder     = &JSONValueNull{}
	_ bin.BareEncoder = &JSONValueNull{}
	_ bin.BareDecoder = &JSONValueNull{}

	_ JSONValueClass = &JSONValueNull{}
)

func (j *JSONValueNull) Zero() bool {
	if j == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (j *JSONValueNull) String() string {
	if j == nil {
		return "JSONValueNull(nil)"
	}
	type Alias JSONValueNull
	return fmt.Sprintf("JSONValueNull%+v", Alias(*j))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*JSONValueNull) TypeID() uint32 {
	return JSONValueNullTypeID
}

// TypeName returns name of type in TL schema.
func (*JSONValueNull) TypeName() string {
	return "jsonValueNull"
}

// TypeInfo returns info about TL type.
func (j *JSONValueNull) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "jsonValueNull",
		ID:   JSONValueNullTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (j *JSONValueNull) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueNull#fa76e0cd as nil")
	}
	b.PutID(JSONValueNullTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *JSONValueNull) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueNull#fa76e0cd as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (j *JSONValueNull) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueNull#fa76e0cd to nil")
	}
	if err := b.ConsumeID(JSONValueNullTypeID); err != nil {
		return fmt.Errorf("unable to decode jsonValueNull#fa76e0cd: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *JSONValueNull) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueNull#fa76e0cd to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (j *JSONValueNull) EncodeTDLibJSON(b tdjson.Encoder) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueNull#fa76e0cd as nil")
	}
	b.ObjStart()
	b.PutID("jsonValueNull")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (j *JSONValueNull) DecodeTDLibJSON(b tdjson.Decoder) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueNull#fa76e0cd to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("jsonValueNull"); err != nil {
				return fmt.Errorf("unable to decode jsonValueNull#fa76e0cd: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// JSONValueBoolean represents TL type `jsonValueBoolean#8050d3b0`.
type JSONValueBoolean struct {
	// The value
	Value bool
}

// JSONValueBooleanTypeID is TL type id of JSONValueBoolean.
const JSONValueBooleanTypeID = 0x8050d3b0

// construct implements constructor of JSONValueClass.
func (j JSONValueBoolean) construct() JSONValueClass { return &j }

// Ensuring interfaces in compile-time for JSONValueBoolean.
var (
	_ bin.Encoder     = &JSONValueBoolean{}
	_ bin.Decoder     = &JSONValueBoolean{}
	_ bin.BareEncoder = &JSONValueBoolean{}
	_ bin.BareDecoder = &JSONValueBoolean{}

	_ JSONValueClass = &JSONValueBoolean{}
)

func (j *JSONValueBoolean) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.Value == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *JSONValueBoolean) String() string {
	if j == nil {
		return "JSONValueBoolean(nil)"
	}
	type Alias JSONValueBoolean
	return fmt.Sprintf("JSONValueBoolean%+v", Alias(*j))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*JSONValueBoolean) TypeID() uint32 {
	return JSONValueBooleanTypeID
}

// TypeName returns name of type in TL schema.
func (*JSONValueBoolean) TypeName() string {
	return "jsonValueBoolean"
}

// TypeInfo returns info about TL type.
func (j *JSONValueBoolean) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "jsonValueBoolean",
		ID:   JSONValueBooleanTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (j *JSONValueBoolean) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueBoolean#8050d3b0 as nil")
	}
	b.PutID(JSONValueBooleanTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *JSONValueBoolean) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueBoolean#8050d3b0 as nil")
	}
	b.PutBool(j.Value)
	return nil
}

// Decode implements bin.Decoder.
func (j *JSONValueBoolean) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueBoolean#8050d3b0 to nil")
	}
	if err := b.ConsumeID(JSONValueBooleanTypeID); err != nil {
		return fmt.Errorf("unable to decode jsonValueBoolean#8050d3b0: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *JSONValueBoolean) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueBoolean#8050d3b0 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode jsonValueBoolean#8050d3b0: field value: %w", err)
		}
		j.Value = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (j *JSONValueBoolean) EncodeTDLibJSON(b tdjson.Encoder) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueBoolean#8050d3b0 as nil")
	}
	b.ObjStart()
	b.PutID("jsonValueBoolean")
	b.Comma()
	b.FieldStart("value")
	b.PutBool(j.Value)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (j *JSONValueBoolean) DecodeTDLibJSON(b tdjson.Decoder) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueBoolean#8050d3b0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("jsonValueBoolean"); err != nil {
				return fmt.Errorf("unable to decode jsonValueBoolean#8050d3b0: %w", err)
			}
		case "value":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode jsonValueBoolean#8050d3b0: field value: %w", err)
			}
			j.Value = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetValue returns value of Value field.
func (j *JSONValueBoolean) GetValue() (value bool) {
	if j == nil {
		return
	}
	return j.Value
}

// JSONValueNumber represents TL type `jsonValueNumber#c3c0146f`.
type JSONValueNumber struct {
	// The value
	Value float64
}

// JSONValueNumberTypeID is TL type id of JSONValueNumber.
const JSONValueNumberTypeID = 0xc3c0146f

// construct implements constructor of JSONValueClass.
func (j JSONValueNumber) construct() JSONValueClass { return &j }

// Ensuring interfaces in compile-time for JSONValueNumber.
var (
	_ bin.Encoder     = &JSONValueNumber{}
	_ bin.Decoder     = &JSONValueNumber{}
	_ bin.BareEncoder = &JSONValueNumber{}
	_ bin.BareDecoder = &JSONValueNumber{}

	_ JSONValueClass = &JSONValueNumber{}
)

func (j *JSONValueNumber) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.Value == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *JSONValueNumber) String() string {
	if j == nil {
		return "JSONValueNumber(nil)"
	}
	type Alias JSONValueNumber
	return fmt.Sprintf("JSONValueNumber%+v", Alias(*j))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*JSONValueNumber) TypeID() uint32 {
	return JSONValueNumberTypeID
}

// TypeName returns name of type in TL schema.
func (*JSONValueNumber) TypeName() string {
	return "jsonValueNumber"
}

// TypeInfo returns info about TL type.
func (j *JSONValueNumber) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "jsonValueNumber",
		ID:   JSONValueNumberTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (j *JSONValueNumber) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueNumber#c3c0146f as nil")
	}
	b.PutID(JSONValueNumberTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *JSONValueNumber) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueNumber#c3c0146f as nil")
	}
	b.PutDouble(j.Value)
	return nil
}

// Decode implements bin.Decoder.
func (j *JSONValueNumber) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueNumber#c3c0146f to nil")
	}
	if err := b.ConsumeID(JSONValueNumberTypeID); err != nil {
		return fmt.Errorf("unable to decode jsonValueNumber#c3c0146f: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *JSONValueNumber) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueNumber#c3c0146f to nil")
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode jsonValueNumber#c3c0146f: field value: %w", err)
		}
		j.Value = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (j *JSONValueNumber) EncodeTDLibJSON(b tdjson.Encoder) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueNumber#c3c0146f as nil")
	}
	b.ObjStart()
	b.PutID("jsonValueNumber")
	b.Comma()
	b.FieldStart("value")
	b.PutDouble(j.Value)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (j *JSONValueNumber) DecodeTDLibJSON(b tdjson.Decoder) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueNumber#c3c0146f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("jsonValueNumber"); err != nil {
				return fmt.Errorf("unable to decode jsonValueNumber#c3c0146f: %w", err)
			}
		case "value":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode jsonValueNumber#c3c0146f: field value: %w", err)
			}
			j.Value = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetValue returns value of Value field.
func (j *JSONValueNumber) GetValue() (value float64) {
	if j == nil {
		return
	}
	return j.Value
}

// JSONValueString represents TL type `jsonValueString#5f3ebdb1`.
type JSONValueString struct {
	// The value
	Value string
}

// JSONValueStringTypeID is TL type id of JSONValueString.
const JSONValueStringTypeID = 0x5f3ebdb1

// construct implements constructor of JSONValueClass.
func (j JSONValueString) construct() JSONValueClass { return &j }

// Ensuring interfaces in compile-time for JSONValueString.
var (
	_ bin.Encoder     = &JSONValueString{}
	_ bin.Decoder     = &JSONValueString{}
	_ bin.BareEncoder = &JSONValueString{}
	_ bin.BareDecoder = &JSONValueString{}

	_ JSONValueClass = &JSONValueString{}
)

func (j *JSONValueString) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.Value == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *JSONValueString) String() string {
	if j == nil {
		return "JSONValueString(nil)"
	}
	type Alias JSONValueString
	return fmt.Sprintf("JSONValueString%+v", Alias(*j))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*JSONValueString) TypeID() uint32 {
	return JSONValueStringTypeID
}

// TypeName returns name of type in TL schema.
func (*JSONValueString) TypeName() string {
	return "jsonValueString"
}

// TypeInfo returns info about TL type.
func (j *JSONValueString) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "jsonValueString",
		ID:   JSONValueStringTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (j *JSONValueString) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueString#5f3ebdb1 as nil")
	}
	b.PutID(JSONValueStringTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *JSONValueString) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueString#5f3ebdb1 as nil")
	}
	b.PutString(j.Value)
	return nil
}

// Decode implements bin.Decoder.
func (j *JSONValueString) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueString#5f3ebdb1 to nil")
	}
	if err := b.ConsumeID(JSONValueStringTypeID); err != nil {
		return fmt.Errorf("unable to decode jsonValueString#5f3ebdb1: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *JSONValueString) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueString#5f3ebdb1 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode jsonValueString#5f3ebdb1: field value: %w", err)
		}
		j.Value = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (j *JSONValueString) EncodeTDLibJSON(b tdjson.Encoder) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueString#5f3ebdb1 as nil")
	}
	b.ObjStart()
	b.PutID("jsonValueString")
	b.Comma()
	b.FieldStart("value")
	b.PutString(j.Value)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (j *JSONValueString) DecodeTDLibJSON(b tdjson.Decoder) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueString#5f3ebdb1 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("jsonValueString"); err != nil {
				return fmt.Errorf("unable to decode jsonValueString#5f3ebdb1: %w", err)
			}
		case "value":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode jsonValueString#5f3ebdb1: field value: %w", err)
			}
			j.Value = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetValue returns value of Value field.
func (j *JSONValueString) GetValue() (value string) {
	if j == nil {
		return
	}
	return j.Value
}

// JSONValueArray represents TL type `jsonValueArray#eccdb0d8`.
type JSONValueArray struct {
	// The list of array elements
	Values []JSONValueClass
}

// JSONValueArrayTypeID is TL type id of JSONValueArray.
const JSONValueArrayTypeID = 0xeccdb0d8

// construct implements constructor of JSONValueClass.
func (j JSONValueArray) construct() JSONValueClass { return &j }

// Ensuring interfaces in compile-time for JSONValueArray.
var (
	_ bin.Encoder     = &JSONValueArray{}
	_ bin.Decoder     = &JSONValueArray{}
	_ bin.BareEncoder = &JSONValueArray{}
	_ bin.BareDecoder = &JSONValueArray{}

	_ JSONValueClass = &JSONValueArray{}
)

func (j *JSONValueArray) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.Values == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *JSONValueArray) String() string {
	if j == nil {
		return "JSONValueArray(nil)"
	}
	type Alias JSONValueArray
	return fmt.Sprintf("JSONValueArray%+v", Alias(*j))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*JSONValueArray) TypeID() uint32 {
	return JSONValueArrayTypeID
}

// TypeName returns name of type in TL schema.
func (*JSONValueArray) TypeName() string {
	return "jsonValueArray"
}

// TypeInfo returns info about TL type.
func (j *JSONValueArray) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "jsonValueArray",
		ID:   JSONValueArrayTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Values",
			SchemaName: "values",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (j *JSONValueArray) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueArray#eccdb0d8 as nil")
	}
	b.PutID(JSONValueArrayTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *JSONValueArray) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueArray#eccdb0d8 as nil")
	}
	b.PutInt(len(j.Values))
	for idx, v := range j.Values {
		if v == nil {
			return fmt.Errorf("unable to encode jsonValueArray#eccdb0d8: field values element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare jsonValueArray#eccdb0d8: field values element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (j *JSONValueArray) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueArray#eccdb0d8 to nil")
	}
	if err := b.ConsumeID(JSONValueArrayTypeID); err != nil {
		return fmt.Errorf("unable to decode jsonValueArray#eccdb0d8: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *JSONValueArray) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueArray#eccdb0d8 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode jsonValueArray#eccdb0d8: field values: %w", err)
		}

		if headerLen > 0 {
			j.Values = make([]JSONValueClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeJSONValue(b)
			if err != nil {
				return fmt.Errorf("unable to decode jsonValueArray#eccdb0d8: field values: %w", err)
			}
			j.Values = append(j.Values, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (j *JSONValueArray) EncodeTDLibJSON(b tdjson.Encoder) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueArray#eccdb0d8 as nil")
	}
	b.ObjStart()
	b.PutID("jsonValueArray")
	b.Comma()
	b.FieldStart("values")
	b.ArrStart()
	for idx, v := range j.Values {
		if v == nil {
			return fmt.Errorf("unable to encode jsonValueArray#eccdb0d8: field values element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode jsonValueArray#eccdb0d8: field values element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (j *JSONValueArray) DecodeTDLibJSON(b tdjson.Decoder) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueArray#eccdb0d8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("jsonValueArray"); err != nil {
				return fmt.Errorf("unable to decode jsonValueArray#eccdb0d8: %w", err)
			}
		case "values":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := DecodeTDLibJSONJSONValue(b)
				if err != nil {
					return fmt.Errorf("unable to decode jsonValueArray#eccdb0d8: field values: %w", err)
				}
				j.Values = append(j.Values, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode jsonValueArray#eccdb0d8: field values: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetValues returns value of Values field.
func (j *JSONValueArray) GetValues() (value []JSONValueClass) {
	if j == nil {
		return
	}
	return j.Values
}

// JSONValueObject represents TL type `jsonValueObject#c67bff40`.
type JSONValueObject struct {
	// The list of object members
	Members []JSONObjectMember
}

// JSONValueObjectTypeID is TL type id of JSONValueObject.
const JSONValueObjectTypeID = 0xc67bff40

// construct implements constructor of JSONValueClass.
func (j JSONValueObject) construct() JSONValueClass { return &j }

// Ensuring interfaces in compile-time for JSONValueObject.
var (
	_ bin.Encoder     = &JSONValueObject{}
	_ bin.Decoder     = &JSONValueObject{}
	_ bin.BareEncoder = &JSONValueObject{}
	_ bin.BareDecoder = &JSONValueObject{}

	_ JSONValueClass = &JSONValueObject{}
)

func (j *JSONValueObject) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.Members == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *JSONValueObject) String() string {
	if j == nil {
		return "JSONValueObject(nil)"
	}
	type Alias JSONValueObject
	return fmt.Sprintf("JSONValueObject%+v", Alias(*j))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*JSONValueObject) TypeID() uint32 {
	return JSONValueObjectTypeID
}

// TypeName returns name of type in TL schema.
func (*JSONValueObject) TypeName() string {
	return "jsonValueObject"
}

// TypeInfo returns info about TL type.
func (j *JSONValueObject) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "jsonValueObject",
		ID:   JSONValueObjectTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Members",
			SchemaName: "members",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (j *JSONValueObject) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueObject#c67bff40 as nil")
	}
	b.PutID(JSONValueObjectTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *JSONValueObject) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueObject#c67bff40 as nil")
	}
	b.PutInt(len(j.Members))
	for idx, v := range j.Members {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare jsonValueObject#c67bff40: field members element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (j *JSONValueObject) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueObject#c67bff40 to nil")
	}
	if err := b.ConsumeID(JSONValueObjectTypeID); err != nil {
		return fmt.Errorf("unable to decode jsonValueObject#c67bff40: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *JSONValueObject) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueObject#c67bff40 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode jsonValueObject#c67bff40: field members: %w", err)
		}

		if headerLen > 0 {
			j.Members = make([]JSONObjectMember, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value JSONObjectMember
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare jsonValueObject#c67bff40: field members: %w", err)
			}
			j.Members = append(j.Members, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (j *JSONValueObject) EncodeTDLibJSON(b tdjson.Encoder) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonValueObject#c67bff40 as nil")
	}
	b.ObjStart()
	b.PutID("jsonValueObject")
	b.Comma()
	b.FieldStart("members")
	b.ArrStart()
	for idx, v := range j.Members {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode jsonValueObject#c67bff40: field members element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (j *JSONValueObject) DecodeTDLibJSON(b tdjson.Decoder) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonValueObject#c67bff40 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("jsonValueObject"); err != nil {
				return fmt.Errorf("unable to decode jsonValueObject#c67bff40: %w", err)
			}
		case "members":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value JSONObjectMember
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode jsonValueObject#c67bff40: field members: %w", err)
				}
				j.Members = append(j.Members, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode jsonValueObject#c67bff40: field members: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetMembers returns value of Members field.
func (j *JSONValueObject) GetMembers() (value []JSONObjectMember) {
	if j == nil {
		return
	}
	return j.Members
}

// JSONValueClassName is schema name of JSONValueClass.
const JSONValueClassName = "JsonValue"

// JSONValueClass represents JsonValue generic type.
//
// Example:
//
//	g, err := tdapi.DecodeJSONValue(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.JSONValueNull: // jsonValueNull#fa76e0cd
//	case *tdapi.JSONValueBoolean: // jsonValueBoolean#8050d3b0
//	case *tdapi.JSONValueNumber: // jsonValueNumber#c3c0146f
//	case *tdapi.JSONValueString: // jsonValueString#5f3ebdb1
//	case *tdapi.JSONValueArray: // jsonValueArray#eccdb0d8
//	case *tdapi.JSONValueObject: // jsonValueObject#c67bff40
//	default: panic(v)
//	}
type JSONValueClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() JSONValueClass

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

// DecodeJSONValue implements binary de-serialization for JSONValueClass.
func DecodeJSONValue(buf *bin.Buffer) (JSONValueClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case JSONValueNullTypeID:
		// Decoding jsonValueNull#fa76e0cd.
		v := JSONValueNull{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case JSONValueBooleanTypeID:
		// Decoding jsonValueBoolean#8050d3b0.
		v := JSONValueBoolean{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case JSONValueNumberTypeID:
		// Decoding jsonValueNumber#c3c0146f.
		v := JSONValueNumber{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case JSONValueStringTypeID:
		// Decoding jsonValueString#5f3ebdb1.
		v := JSONValueString{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case JSONValueArrayTypeID:
		// Decoding jsonValueArray#eccdb0d8.
		v := JSONValueArray{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case JSONValueObjectTypeID:
		// Decoding jsonValueObject#c67bff40.
		v := JSONValueObject{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode JSONValueClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONJSONValue implements binary de-serialization for JSONValueClass.
func DecodeTDLibJSONJSONValue(buf tdjson.Decoder) (JSONValueClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "jsonValueNull":
		// Decoding jsonValueNull#fa76e0cd.
		v := JSONValueNull{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case "jsonValueBoolean":
		// Decoding jsonValueBoolean#8050d3b0.
		v := JSONValueBoolean{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case "jsonValueNumber":
		// Decoding jsonValueNumber#c3c0146f.
		v := JSONValueNumber{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case "jsonValueString":
		// Decoding jsonValueString#5f3ebdb1.
		v := JSONValueString{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case "jsonValueArray":
		// Decoding jsonValueArray#eccdb0d8.
		v := JSONValueArray{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case "jsonValueObject":
		// Decoding jsonValueObject#c67bff40.
		v := JSONValueObject{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode JSONValueClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// JSONValue boxes the JSONValueClass providing a helper.
type JSONValueBox struct {
	JsonValue JSONValueClass
}

// Decode implements bin.Decoder for JSONValueBox.
func (b *JSONValueBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode JSONValueBox to nil")
	}
	v, err := DecodeJSONValue(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.JsonValue = v
	return nil
}

// Encode implements bin.Encode for JSONValueBox.
func (b *JSONValueBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.JsonValue == nil {
		return fmt.Errorf("unable to encode JSONValueClass as nil")
	}
	return b.JsonValue.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for JSONValueBox.
func (b *JSONValueBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode JSONValueBox to nil")
	}
	v, err := DecodeTDLibJSONJSONValue(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.JsonValue = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for JSONValueBox.
func (b *JSONValueBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.JsonValue == nil {
		return fmt.Errorf("unable to encode JSONValueClass as nil")
	}
	return b.JsonValue.EncodeTDLibJSON(buf)
}
