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

// RevenueWithdrawalStatePending represents TL type `revenueWithdrawalStatePending#5d314fa5`.
type RevenueWithdrawalStatePending struct {
}

// RevenueWithdrawalStatePendingTypeID is TL type id of RevenueWithdrawalStatePending.
const RevenueWithdrawalStatePendingTypeID = 0x5d314fa5

// construct implements constructor of RevenueWithdrawalStateClass.
func (r RevenueWithdrawalStatePending) construct() RevenueWithdrawalStateClass { return &r }

// Ensuring interfaces in compile-time for RevenueWithdrawalStatePending.
var (
	_ bin.Encoder     = &RevenueWithdrawalStatePending{}
	_ bin.Decoder     = &RevenueWithdrawalStatePending{}
	_ bin.BareEncoder = &RevenueWithdrawalStatePending{}
	_ bin.BareDecoder = &RevenueWithdrawalStatePending{}

	_ RevenueWithdrawalStateClass = &RevenueWithdrawalStatePending{}
)

func (r *RevenueWithdrawalStatePending) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *RevenueWithdrawalStatePending) String() string {
	if r == nil {
		return "RevenueWithdrawalStatePending(nil)"
	}
	type Alias RevenueWithdrawalStatePending
	return fmt.Sprintf("RevenueWithdrawalStatePending%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RevenueWithdrawalStatePending) TypeID() uint32 {
	return RevenueWithdrawalStatePendingTypeID
}

// TypeName returns name of type in TL schema.
func (*RevenueWithdrawalStatePending) TypeName() string {
	return "revenueWithdrawalStatePending"
}

// TypeInfo returns info about TL type.
func (r *RevenueWithdrawalStatePending) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "revenueWithdrawalStatePending",
		ID:   RevenueWithdrawalStatePendingTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *RevenueWithdrawalStatePending) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode revenueWithdrawalStatePending#5d314fa5 as nil")
	}
	b.PutID(RevenueWithdrawalStatePendingTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RevenueWithdrawalStatePending) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode revenueWithdrawalStatePending#5d314fa5 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RevenueWithdrawalStatePending) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode revenueWithdrawalStatePending#5d314fa5 to nil")
	}
	if err := b.ConsumeID(RevenueWithdrawalStatePendingTypeID); err != nil {
		return fmt.Errorf("unable to decode revenueWithdrawalStatePending#5d314fa5: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RevenueWithdrawalStatePending) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode revenueWithdrawalStatePending#5d314fa5 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RevenueWithdrawalStatePending) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode revenueWithdrawalStatePending#5d314fa5 as nil")
	}
	b.ObjStart()
	b.PutID("revenueWithdrawalStatePending")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RevenueWithdrawalStatePending) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode revenueWithdrawalStatePending#5d314fa5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("revenueWithdrawalStatePending"); err != nil {
				return fmt.Errorf("unable to decode revenueWithdrawalStatePending#5d314fa5: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// RevenueWithdrawalStateSucceeded represents TL type `revenueWithdrawalStateSucceeded#fd14e0a`.
type RevenueWithdrawalStateSucceeded struct {
	// Point in time (Unix timestamp) when the withdrawal was completed
	Date int32
	// The URL where the withdrawal transaction can be viewed
	URL string
}

// RevenueWithdrawalStateSucceededTypeID is TL type id of RevenueWithdrawalStateSucceeded.
const RevenueWithdrawalStateSucceededTypeID = 0xfd14e0a

// construct implements constructor of RevenueWithdrawalStateClass.
func (r RevenueWithdrawalStateSucceeded) construct() RevenueWithdrawalStateClass { return &r }

// Ensuring interfaces in compile-time for RevenueWithdrawalStateSucceeded.
var (
	_ bin.Encoder     = &RevenueWithdrawalStateSucceeded{}
	_ bin.Decoder     = &RevenueWithdrawalStateSucceeded{}
	_ bin.BareEncoder = &RevenueWithdrawalStateSucceeded{}
	_ bin.BareDecoder = &RevenueWithdrawalStateSucceeded{}

	_ RevenueWithdrawalStateClass = &RevenueWithdrawalStateSucceeded{}
)

func (r *RevenueWithdrawalStateSucceeded) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Date == 0) {
		return false
	}
	if !(r.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RevenueWithdrawalStateSucceeded) String() string {
	if r == nil {
		return "RevenueWithdrawalStateSucceeded(nil)"
	}
	type Alias RevenueWithdrawalStateSucceeded
	return fmt.Sprintf("RevenueWithdrawalStateSucceeded%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RevenueWithdrawalStateSucceeded) TypeID() uint32 {
	return RevenueWithdrawalStateSucceededTypeID
}

// TypeName returns name of type in TL schema.
func (*RevenueWithdrawalStateSucceeded) TypeName() string {
	return "revenueWithdrawalStateSucceeded"
}

// TypeInfo returns info about TL type.
func (r *RevenueWithdrawalStateSucceeded) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "revenueWithdrawalStateSucceeded",
		ID:   RevenueWithdrawalStateSucceededTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RevenueWithdrawalStateSucceeded) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode revenueWithdrawalStateSucceeded#fd14e0a as nil")
	}
	b.PutID(RevenueWithdrawalStateSucceededTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RevenueWithdrawalStateSucceeded) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode revenueWithdrawalStateSucceeded#fd14e0a as nil")
	}
	b.PutInt32(r.Date)
	b.PutString(r.URL)
	return nil
}

// Decode implements bin.Decoder.
func (r *RevenueWithdrawalStateSucceeded) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode revenueWithdrawalStateSucceeded#fd14e0a to nil")
	}
	if err := b.ConsumeID(RevenueWithdrawalStateSucceededTypeID); err != nil {
		return fmt.Errorf("unable to decode revenueWithdrawalStateSucceeded#fd14e0a: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RevenueWithdrawalStateSucceeded) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode revenueWithdrawalStateSucceeded#fd14e0a to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode revenueWithdrawalStateSucceeded#fd14e0a: field date: %w", err)
		}
		r.Date = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode revenueWithdrawalStateSucceeded#fd14e0a: field url: %w", err)
		}
		r.URL = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RevenueWithdrawalStateSucceeded) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode revenueWithdrawalStateSucceeded#fd14e0a as nil")
	}
	b.ObjStart()
	b.PutID("revenueWithdrawalStateSucceeded")
	b.Comma()
	b.FieldStart("date")
	b.PutInt32(r.Date)
	b.Comma()
	b.FieldStart("url")
	b.PutString(r.URL)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RevenueWithdrawalStateSucceeded) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode revenueWithdrawalStateSucceeded#fd14e0a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("revenueWithdrawalStateSucceeded"); err != nil {
				return fmt.Errorf("unable to decode revenueWithdrawalStateSucceeded#fd14e0a: %w", err)
			}
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode revenueWithdrawalStateSucceeded#fd14e0a: field date: %w", err)
			}
			r.Date = value
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode revenueWithdrawalStateSucceeded#fd14e0a: field url: %w", err)
			}
			r.URL = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDate returns value of Date field.
func (r *RevenueWithdrawalStateSucceeded) GetDate() (value int32) {
	if r == nil {
		return
	}
	return r.Date
}

// GetURL returns value of URL field.
func (r *RevenueWithdrawalStateSucceeded) GetURL() (value string) {
	if r == nil {
		return
	}
	return r.URL
}

// RevenueWithdrawalStateFailed represents TL type `revenueWithdrawalStateFailed#ff413089`.
type RevenueWithdrawalStateFailed struct {
}

// RevenueWithdrawalStateFailedTypeID is TL type id of RevenueWithdrawalStateFailed.
const RevenueWithdrawalStateFailedTypeID = 0xff413089

// construct implements constructor of RevenueWithdrawalStateClass.
func (r RevenueWithdrawalStateFailed) construct() RevenueWithdrawalStateClass { return &r }

// Ensuring interfaces in compile-time for RevenueWithdrawalStateFailed.
var (
	_ bin.Encoder     = &RevenueWithdrawalStateFailed{}
	_ bin.Decoder     = &RevenueWithdrawalStateFailed{}
	_ bin.BareEncoder = &RevenueWithdrawalStateFailed{}
	_ bin.BareDecoder = &RevenueWithdrawalStateFailed{}

	_ RevenueWithdrawalStateClass = &RevenueWithdrawalStateFailed{}
)

func (r *RevenueWithdrawalStateFailed) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *RevenueWithdrawalStateFailed) String() string {
	if r == nil {
		return "RevenueWithdrawalStateFailed(nil)"
	}
	type Alias RevenueWithdrawalStateFailed
	return fmt.Sprintf("RevenueWithdrawalStateFailed%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RevenueWithdrawalStateFailed) TypeID() uint32 {
	return RevenueWithdrawalStateFailedTypeID
}

// TypeName returns name of type in TL schema.
func (*RevenueWithdrawalStateFailed) TypeName() string {
	return "revenueWithdrawalStateFailed"
}

// TypeInfo returns info about TL type.
func (r *RevenueWithdrawalStateFailed) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "revenueWithdrawalStateFailed",
		ID:   RevenueWithdrawalStateFailedTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *RevenueWithdrawalStateFailed) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode revenueWithdrawalStateFailed#ff413089 as nil")
	}
	b.PutID(RevenueWithdrawalStateFailedTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RevenueWithdrawalStateFailed) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode revenueWithdrawalStateFailed#ff413089 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RevenueWithdrawalStateFailed) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode revenueWithdrawalStateFailed#ff413089 to nil")
	}
	if err := b.ConsumeID(RevenueWithdrawalStateFailedTypeID); err != nil {
		return fmt.Errorf("unable to decode revenueWithdrawalStateFailed#ff413089: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RevenueWithdrawalStateFailed) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode revenueWithdrawalStateFailed#ff413089 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RevenueWithdrawalStateFailed) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode revenueWithdrawalStateFailed#ff413089 as nil")
	}
	b.ObjStart()
	b.PutID("revenueWithdrawalStateFailed")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RevenueWithdrawalStateFailed) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode revenueWithdrawalStateFailed#ff413089 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("revenueWithdrawalStateFailed"); err != nil {
				return fmt.Errorf("unable to decode revenueWithdrawalStateFailed#ff413089: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// RevenueWithdrawalStateClassName is schema name of RevenueWithdrawalStateClass.
const RevenueWithdrawalStateClassName = "RevenueWithdrawalState"

// RevenueWithdrawalStateClass represents RevenueWithdrawalState generic type.
//
// Example:
//
//	g, err := tdapi.DecodeRevenueWithdrawalState(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.RevenueWithdrawalStatePending: // revenueWithdrawalStatePending#5d314fa5
//	case *tdapi.RevenueWithdrawalStateSucceeded: // revenueWithdrawalStateSucceeded#fd14e0a
//	case *tdapi.RevenueWithdrawalStateFailed: // revenueWithdrawalStateFailed#ff413089
//	default: panic(v)
//	}
type RevenueWithdrawalStateClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() RevenueWithdrawalStateClass

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

// DecodeRevenueWithdrawalState implements binary de-serialization for RevenueWithdrawalStateClass.
func DecodeRevenueWithdrawalState(buf *bin.Buffer) (RevenueWithdrawalStateClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case RevenueWithdrawalStatePendingTypeID:
		// Decoding revenueWithdrawalStatePending#5d314fa5.
		v := RevenueWithdrawalStatePending{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RevenueWithdrawalStateClass: %w", err)
		}
		return &v, nil
	case RevenueWithdrawalStateSucceededTypeID:
		// Decoding revenueWithdrawalStateSucceeded#fd14e0a.
		v := RevenueWithdrawalStateSucceeded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RevenueWithdrawalStateClass: %w", err)
		}
		return &v, nil
	case RevenueWithdrawalStateFailedTypeID:
		// Decoding revenueWithdrawalStateFailed#ff413089.
		v := RevenueWithdrawalStateFailed{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RevenueWithdrawalStateClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode RevenueWithdrawalStateClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONRevenueWithdrawalState implements binary de-serialization for RevenueWithdrawalStateClass.
func DecodeTDLibJSONRevenueWithdrawalState(buf tdjson.Decoder) (RevenueWithdrawalStateClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "revenueWithdrawalStatePending":
		// Decoding revenueWithdrawalStatePending#5d314fa5.
		v := RevenueWithdrawalStatePending{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RevenueWithdrawalStateClass: %w", err)
		}
		return &v, nil
	case "revenueWithdrawalStateSucceeded":
		// Decoding revenueWithdrawalStateSucceeded#fd14e0a.
		v := RevenueWithdrawalStateSucceeded{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RevenueWithdrawalStateClass: %w", err)
		}
		return &v, nil
	case "revenueWithdrawalStateFailed":
		// Decoding revenueWithdrawalStateFailed#ff413089.
		v := RevenueWithdrawalStateFailed{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RevenueWithdrawalStateClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode RevenueWithdrawalStateClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// RevenueWithdrawalState boxes the RevenueWithdrawalStateClass providing a helper.
type RevenueWithdrawalStateBox struct {
	RevenueWithdrawalState RevenueWithdrawalStateClass
}

// Decode implements bin.Decoder for RevenueWithdrawalStateBox.
func (b *RevenueWithdrawalStateBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode RevenueWithdrawalStateBox to nil")
	}
	v, err := DecodeRevenueWithdrawalState(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.RevenueWithdrawalState = v
	return nil
}

// Encode implements bin.Encode for RevenueWithdrawalStateBox.
func (b *RevenueWithdrawalStateBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.RevenueWithdrawalState == nil {
		return fmt.Errorf("unable to encode RevenueWithdrawalStateClass as nil")
	}
	return b.RevenueWithdrawalState.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for RevenueWithdrawalStateBox.
func (b *RevenueWithdrawalStateBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode RevenueWithdrawalStateBox to nil")
	}
	v, err := DecodeTDLibJSONRevenueWithdrawalState(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.RevenueWithdrawalState = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for RevenueWithdrawalStateBox.
func (b *RevenueWithdrawalStateBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.RevenueWithdrawalState == nil {
		return fmt.Errorf("unable to encode RevenueWithdrawalStateClass as nil")
	}
	return b.RevenueWithdrawalState.EncodeTDLibJSON(buf)
}
