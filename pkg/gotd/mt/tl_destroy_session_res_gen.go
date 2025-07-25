// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// DestroySessionOk represents TL type `destroy_session_ok#e22045fc`.
type DestroySessionOk struct {
	// SessionID field of DestroySessionOk.
	SessionID int64
}

// DestroySessionOkTypeID is TL type id of DestroySessionOk.
const DestroySessionOkTypeID = 0xe22045fc

// construct implements constructor of DestroySessionResClass.
func (d DestroySessionOk) construct() DestroySessionResClass { return &d }

// Ensuring interfaces in compile-time for DestroySessionOk.
var (
	_ bin.Encoder     = &DestroySessionOk{}
	_ bin.Decoder     = &DestroySessionOk{}
	_ bin.BareEncoder = &DestroySessionOk{}
	_ bin.BareDecoder = &DestroySessionOk{}

	_ DestroySessionResClass = &DestroySessionOk{}
)

func (d *DestroySessionOk) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.SessionID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DestroySessionOk) String() string {
	if d == nil {
		return "DestroySessionOk(nil)"
	}
	type Alias DestroySessionOk
	return fmt.Sprintf("DestroySessionOk%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DestroySessionOk) TypeID() uint32 {
	return DestroySessionOkTypeID
}

// TypeName returns name of type in TL schema.
func (*DestroySessionOk) TypeName() string {
	return "destroy_session_ok"
}

// TypeInfo returns info about TL type.
func (d *DestroySessionOk) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "destroy_session_ok",
		ID:   DestroySessionOkTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SessionID",
			SchemaName: "session_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DestroySessionOk) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode destroy_session_ok#e22045fc as nil")
	}
	b.PutID(DestroySessionOkTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DestroySessionOk) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode destroy_session_ok#e22045fc as nil")
	}
	b.PutLong(d.SessionID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DestroySessionOk) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode destroy_session_ok#e22045fc to nil")
	}
	if err := b.ConsumeID(DestroySessionOkTypeID); err != nil {
		return fmt.Errorf("unable to decode destroy_session_ok#e22045fc: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DestroySessionOk) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode destroy_session_ok#e22045fc to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode destroy_session_ok#e22045fc: field session_id: %w", err)
		}
		d.SessionID = value
	}
	return nil
}

// GetSessionID returns value of SessionID field.
func (d *DestroySessionOk) GetSessionID() (value int64) {
	if d == nil {
		return
	}
	return d.SessionID
}

// DestroySessionNone represents TL type `destroy_session_none#62d350c9`.
type DestroySessionNone struct {
	// SessionID field of DestroySessionNone.
	SessionID int64
}

// DestroySessionNoneTypeID is TL type id of DestroySessionNone.
const DestroySessionNoneTypeID = 0x62d350c9

// construct implements constructor of DestroySessionResClass.
func (d DestroySessionNone) construct() DestroySessionResClass { return &d }

// Ensuring interfaces in compile-time for DestroySessionNone.
var (
	_ bin.Encoder     = &DestroySessionNone{}
	_ bin.Decoder     = &DestroySessionNone{}
	_ bin.BareEncoder = &DestroySessionNone{}
	_ bin.BareDecoder = &DestroySessionNone{}

	_ DestroySessionResClass = &DestroySessionNone{}
)

func (d *DestroySessionNone) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.SessionID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DestroySessionNone) String() string {
	if d == nil {
		return "DestroySessionNone(nil)"
	}
	type Alias DestroySessionNone
	return fmt.Sprintf("DestroySessionNone%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DestroySessionNone) TypeID() uint32 {
	return DestroySessionNoneTypeID
}

// TypeName returns name of type in TL schema.
func (*DestroySessionNone) TypeName() string {
	return "destroy_session_none"
}

// TypeInfo returns info about TL type.
func (d *DestroySessionNone) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "destroy_session_none",
		ID:   DestroySessionNoneTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SessionID",
			SchemaName: "session_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DestroySessionNone) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode destroy_session_none#62d350c9 as nil")
	}
	b.PutID(DestroySessionNoneTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DestroySessionNone) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode destroy_session_none#62d350c9 as nil")
	}
	b.PutLong(d.SessionID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DestroySessionNone) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode destroy_session_none#62d350c9 to nil")
	}
	if err := b.ConsumeID(DestroySessionNoneTypeID); err != nil {
		return fmt.Errorf("unable to decode destroy_session_none#62d350c9: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DestroySessionNone) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode destroy_session_none#62d350c9 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode destroy_session_none#62d350c9: field session_id: %w", err)
		}
		d.SessionID = value
	}
	return nil
}

// GetSessionID returns value of SessionID field.
func (d *DestroySessionNone) GetSessionID() (value int64) {
	if d == nil {
		return
	}
	return d.SessionID
}

// DestroySessionResClassName is schema name of DestroySessionResClass.
const DestroySessionResClassName = "DestroySessionRes"

// DestroySessionResClass represents DestroySessionRes generic type.
//
// Example:
//
//	g, err := mt.DecodeDestroySessionRes(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *mt.DestroySessionOk: // destroy_session_ok#e22045fc
//	case *mt.DestroySessionNone: // destroy_session_none#62d350c9
//	default: panic(v)
//	}
type DestroySessionResClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() DestroySessionResClass

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

	// SessionID field of DestroySessionOk.
	GetSessionID() (value int64)
}

// DecodeDestroySessionRes implements binary de-serialization for DestroySessionResClass.
func DecodeDestroySessionRes(buf *bin.Buffer) (DestroySessionResClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DestroySessionOkTypeID:
		// Decoding destroy_session_ok#e22045fc.
		v := DestroySessionOk{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DestroySessionResClass: %w", err)
		}
		return &v, nil
	case DestroySessionNoneTypeID:
		// Decoding destroy_session_none#62d350c9.
		v := DestroySessionNone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DestroySessionResClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DestroySessionResClass: %w", bin.NewUnexpectedID(id))
	}
}

// DestroySessionRes boxes the DestroySessionResClass providing a helper.
type DestroySessionResBox struct {
	DestroySessionRes DestroySessionResClass
}

// Decode implements bin.Decoder for DestroySessionResBox.
func (b *DestroySessionResBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode DestroySessionResBox to nil")
	}
	v, err := DecodeDestroySessionRes(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DestroySessionRes = v
	return nil
}

// Encode implements bin.Encode for DestroySessionResBox.
func (b *DestroySessionResBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.DestroySessionRes == nil {
		return fmt.Errorf("unable to encode DestroySessionResClass as nil")
	}
	return b.DestroySessionRes.Encode(buf)
}
