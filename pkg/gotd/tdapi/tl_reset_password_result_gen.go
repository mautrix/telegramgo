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

// ResetPasswordResultOk represents TL type `resetPasswordResultOk#acb763f9`.
type ResetPasswordResultOk struct {
}

// ResetPasswordResultOkTypeID is TL type id of ResetPasswordResultOk.
const ResetPasswordResultOkTypeID = 0xacb763f9

// construct implements constructor of ResetPasswordResultClass.
func (r ResetPasswordResultOk) construct() ResetPasswordResultClass { return &r }

// Ensuring interfaces in compile-time for ResetPasswordResultOk.
var (
	_ bin.Encoder     = &ResetPasswordResultOk{}
	_ bin.Decoder     = &ResetPasswordResultOk{}
	_ bin.BareEncoder = &ResetPasswordResultOk{}
	_ bin.BareDecoder = &ResetPasswordResultOk{}

	_ ResetPasswordResultClass = &ResetPasswordResultOk{}
)

func (r *ResetPasswordResultOk) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ResetPasswordResultOk) String() string {
	if r == nil {
		return "ResetPasswordResultOk(nil)"
	}
	type Alias ResetPasswordResultOk
	return fmt.Sprintf("ResetPasswordResultOk%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ResetPasswordResultOk) TypeID() uint32 {
	return ResetPasswordResultOkTypeID
}

// TypeName returns name of type in TL schema.
func (*ResetPasswordResultOk) TypeName() string {
	return "resetPasswordResultOk"
}

// TypeInfo returns info about TL type.
func (r *ResetPasswordResultOk) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "resetPasswordResultOk",
		ID:   ResetPasswordResultOkTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ResetPasswordResultOk) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resetPasswordResultOk#acb763f9 as nil")
	}
	b.PutID(ResetPasswordResultOkTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ResetPasswordResultOk) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resetPasswordResultOk#acb763f9 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ResetPasswordResultOk) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resetPasswordResultOk#acb763f9 to nil")
	}
	if err := b.ConsumeID(ResetPasswordResultOkTypeID); err != nil {
		return fmt.Errorf("unable to decode resetPasswordResultOk#acb763f9: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ResetPasswordResultOk) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resetPasswordResultOk#acb763f9 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ResetPasswordResultOk) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode resetPasswordResultOk#acb763f9 as nil")
	}
	b.ObjStart()
	b.PutID("resetPasswordResultOk")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ResetPasswordResultOk) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode resetPasswordResultOk#acb763f9 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("resetPasswordResultOk"); err != nil {
				return fmt.Errorf("unable to decode resetPasswordResultOk#acb763f9: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ResetPasswordResultPending represents TL type `resetPasswordResultPending#4729dc59`.
type ResetPasswordResultPending struct {
	// Point in time (Unix timestamp) after which the password can be reset immediately using
	// resetPassword
	PendingResetDate int32
}

// ResetPasswordResultPendingTypeID is TL type id of ResetPasswordResultPending.
const ResetPasswordResultPendingTypeID = 0x4729dc59

// construct implements constructor of ResetPasswordResultClass.
func (r ResetPasswordResultPending) construct() ResetPasswordResultClass { return &r }

// Ensuring interfaces in compile-time for ResetPasswordResultPending.
var (
	_ bin.Encoder     = &ResetPasswordResultPending{}
	_ bin.Decoder     = &ResetPasswordResultPending{}
	_ bin.BareEncoder = &ResetPasswordResultPending{}
	_ bin.BareDecoder = &ResetPasswordResultPending{}

	_ ResetPasswordResultClass = &ResetPasswordResultPending{}
)

func (r *ResetPasswordResultPending) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.PendingResetDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ResetPasswordResultPending) String() string {
	if r == nil {
		return "ResetPasswordResultPending(nil)"
	}
	type Alias ResetPasswordResultPending
	return fmt.Sprintf("ResetPasswordResultPending%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ResetPasswordResultPending) TypeID() uint32 {
	return ResetPasswordResultPendingTypeID
}

// TypeName returns name of type in TL schema.
func (*ResetPasswordResultPending) TypeName() string {
	return "resetPasswordResultPending"
}

// TypeInfo returns info about TL type.
func (r *ResetPasswordResultPending) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "resetPasswordResultPending",
		ID:   ResetPasswordResultPendingTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PendingResetDate",
			SchemaName: "pending_reset_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ResetPasswordResultPending) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resetPasswordResultPending#4729dc59 as nil")
	}
	b.PutID(ResetPasswordResultPendingTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ResetPasswordResultPending) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resetPasswordResultPending#4729dc59 as nil")
	}
	b.PutInt32(r.PendingResetDate)
	return nil
}

// Decode implements bin.Decoder.
func (r *ResetPasswordResultPending) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resetPasswordResultPending#4729dc59 to nil")
	}
	if err := b.ConsumeID(ResetPasswordResultPendingTypeID); err != nil {
		return fmt.Errorf("unable to decode resetPasswordResultPending#4729dc59: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ResetPasswordResultPending) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resetPasswordResultPending#4729dc59 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode resetPasswordResultPending#4729dc59: field pending_reset_date: %w", err)
		}
		r.PendingResetDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ResetPasswordResultPending) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode resetPasswordResultPending#4729dc59 as nil")
	}
	b.ObjStart()
	b.PutID("resetPasswordResultPending")
	b.Comma()
	b.FieldStart("pending_reset_date")
	b.PutInt32(r.PendingResetDate)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ResetPasswordResultPending) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode resetPasswordResultPending#4729dc59 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("resetPasswordResultPending"); err != nil {
				return fmt.Errorf("unable to decode resetPasswordResultPending#4729dc59: %w", err)
			}
		case "pending_reset_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode resetPasswordResultPending#4729dc59: field pending_reset_date: %w", err)
			}
			r.PendingResetDate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPendingResetDate returns value of PendingResetDate field.
func (r *ResetPasswordResultPending) GetPendingResetDate() (value int32) {
	if r == nil {
		return
	}
	return r.PendingResetDate
}

// ResetPasswordResultDeclined represents TL type `resetPasswordResultDeclined#b857e0cb`.
type ResetPasswordResultDeclined struct {
	// Point in time (Unix timestamp) when the password reset can be retried
	RetryDate int32
}

// ResetPasswordResultDeclinedTypeID is TL type id of ResetPasswordResultDeclined.
const ResetPasswordResultDeclinedTypeID = 0xb857e0cb

// construct implements constructor of ResetPasswordResultClass.
func (r ResetPasswordResultDeclined) construct() ResetPasswordResultClass { return &r }

// Ensuring interfaces in compile-time for ResetPasswordResultDeclined.
var (
	_ bin.Encoder     = &ResetPasswordResultDeclined{}
	_ bin.Decoder     = &ResetPasswordResultDeclined{}
	_ bin.BareEncoder = &ResetPasswordResultDeclined{}
	_ bin.BareDecoder = &ResetPasswordResultDeclined{}

	_ ResetPasswordResultClass = &ResetPasswordResultDeclined{}
)

func (r *ResetPasswordResultDeclined) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.RetryDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ResetPasswordResultDeclined) String() string {
	if r == nil {
		return "ResetPasswordResultDeclined(nil)"
	}
	type Alias ResetPasswordResultDeclined
	return fmt.Sprintf("ResetPasswordResultDeclined%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ResetPasswordResultDeclined) TypeID() uint32 {
	return ResetPasswordResultDeclinedTypeID
}

// TypeName returns name of type in TL schema.
func (*ResetPasswordResultDeclined) TypeName() string {
	return "resetPasswordResultDeclined"
}

// TypeInfo returns info about TL type.
func (r *ResetPasswordResultDeclined) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "resetPasswordResultDeclined",
		ID:   ResetPasswordResultDeclinedTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RetryDate",
			SchemaName: "retry_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ResetPasswordResultDeclined) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resetPasswordResultDeclined#b857e0cb as nil")
	}
	b.PutID(ResetPasswordResultDeclinedTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ResetPasswordResultDeclined) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resetPasswordResultDeclined#b857e0cb as nil")
	}
	b.PutInt32(r.RetryDate)
	return nil
}

// Decode implements bin.Decoder.
func (r *ResetPasswordResultDeclined) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resetPasswordResultDeclined#b857e0cb to nil")
	}
	if err := b.ConsumeID(ResetPasswordResultDeclinedTypeID); err != nil {
		return fmt.Errorf("unable to decode resetPasswordResultDeclined#b857e0cb: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ResetPasswordResultDeclined) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resetPasswordResultDeclined#b857e0cb to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode resetPasswordResultDeclined#b857e0cb: field retry_date: %w", err)
		}
		r.RetryDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ResetPasswordResultDeclined) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode resetPasswordResultDeclined#b857e0cb as nil")
	}
	b.ObjStart()
	b.PutID("resetPasswordResultDeclined")
	b.Comma()
	b.FieldStart("retry_date")
	b.PutInt32(r.RetryDate)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ResetPasswordResultDeclined) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode resetPasswordResultDeclined#b857e0cb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("resetPasswordResultDeclined"); err != nil {
				return fmt.Errorf("unable to decode resetPasswordResultDeclined#b857e0cb: %w", err)
			}
		case "retry_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode resetPasswordResultDeclined#b857e0cb: field retry_date: %w", err)
			}
			r.RetryDate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetRetryDate returns value of RetryDate field.
func (r *ResetPasswordResultDeclined) GetRetryDate() (value int32) {
	if r == nil {
		return
	}
	return r.RetryDate
}

// ResetPasswordResultClassName is schema name of ResetPasswordResultClass.
const ResetPasswordResultClassName = "ResetPasswordResult"

// ResetPasswordResultClass represents ResetPasswordResult generic type.
//
// Example:
//
//	g, err := tdapi.DecodeResetPasswordResult(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.ResetPasswordResultOk: // resetPasswordResultOk#acb763f9
//	case *tdapi.ResetPasswordResultPending: // resetPasswordResultPending#4729dc59
//	case *tdapi.ResetPasswordResultDeclined: // resetPasswordResultDeclined#b857e0cb
//	default: panic(v)
//	}
type ResetPasswordResultClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ResetPasswordResultClass

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

// DecodeResetPasswordResult implements binary de-serialization for ResetPasswordResultClass.
func DecodeResetPasswordResult(buf *bin.Buffer) (ResetPasswordResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ResetPasswordResultOkTypeID:
		// Decoding resetPasswordResultOk#acb763f9.
		v := ResetPasswordResultOk{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ResetPasswordResultClass: %w", err)
		}
		return &v, nil
	case ResetPasswordResultPendingTypeID:
		// Decoding resetPasswordResultPending#4729dc59.
		v := ResetPasswordResultPending{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ResetPasswordResultClass: %w", err)
		}
		return &v, nil
	case ResetPasswordResultDeclinedTypeID:
		// Decoding resetPasswordResultDeclined#b857e0cb.
		v := ResetPasswordResultDeclined{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ResetPasswordResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ResetPasswordResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONResetPasswordResult implements binary de-serialization for ResetPasswordResultClass.
func DecodeTDLibJSONResetPasswordResult(buf tdjson.Decoder) (ResetPasswordResultClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "resetPasswordResultOk":
		// Decoding resetPasswordResultOk#acb763f9.
		v := ResetPasswordResultOk{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ResetPasswordResultClass: %w", err)
		}
		return &v, nil
	case "resetPasswordResultPending":
		// Decoding resetPasswordResultPending#4729dc59.
		v := ResetPasswordResultPending{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ResetPasswordResultClass: %w", err)
		}
		return &v, nil
	case "resetPasswordResultDeclined":
		// Decoding resetPasswordResultDeclined#b857e0cb.
		v := ResetPasswordResultDeclined{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ResetPasswordResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ResetPasswordResultClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// ResetPasswordResult boxes the ResetPasswordResultClass providing a helper.
type ResetPasswordResultBox struct {
	ResetPasswordResult ResetPasswordResultClass
}

// Decode implements bin.Decoder for ResetPasswordResultBox.
func (b *ResetPasswordResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ResetPasswordResultBox to nil")
	}
	v, err := DecodeResetPasswordResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ResetPasswordResult = v
	return nil
}

// Encode implements bin.Encode for ResetPasswordResultBox.
func (b *ResetPasswordResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ResetPasswordResult == nil {
		return fmt.Errorf("unable to encode ResetPasswordResultClass as nil")
	}
	return b.ResetPasswordResult.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for ResetPasswordResultBox.
func (b *ResetPasswordResultBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode ResetPasswordResultBox to nil")
	}
	v, err := DecodeTDLibJSONResetPasswordResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ResetPasswordResult = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for ResetPasswordResultBox.
func (b *ResetPasswordResultBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.ResetPasswordResult == nil {
		return fmt.Errorf("unable to encode ResetPasswordResultClass as nil")
	}
	return b.ResetPasswordResult.EncodeTDLibJSON(buf)
}
