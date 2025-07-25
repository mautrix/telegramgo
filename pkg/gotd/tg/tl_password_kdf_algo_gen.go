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

// PasswordKdfAlgoUnknown represents TL type `passwordKdfAlgoUnknown#d45ab096`.
// Unknown KDF (most likely, the client is outdated and does not support the specified
// KDF algorithm)
//
// See https://core.telegram.org/constructor/passwordKdfAlgoUnknown for reference.
type PasswordKdfAlgoUnknown struct {
}

// PasswordKdfAlgoUnknownTypeID is TL type id of PasswordKdfAlgoUnknown.
const PasswordKdfAlgoUnknownTypeID = 0xd45ab096

// construct implements constructor of PasswordKdfAlgoClass.
func (p PasswordKdfAlgoUnknown) construct() PasswordKdfAlgoClass { return &p }

// Ensuring interfaces in compile-time for PasswordKdfAlgoUnknown.
var (
	_ bin.Encoder     = &PasswordKdfAlgoUnknown{}
	_ bin.Decoder     = &PasswordKdfAlgoUnknown{}
	_ bin.BareEncoder = &PasswordKdfAlgoUnknown{}
	_ bin.BareDecoder = &PasswordKdfAlgoUnknown{}

	_ PasswordKdfAlgoClass = &PasswordKdfAlgoUnknown{}
)

func (p *PasswordKdfAlgoUnknown) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PasswordKdfAlgoUnknown) String() string {
	if p == nil {
		return "PasswordKdfAlgoUnknown(nil)"
	}
	type Alias PasswordKdfAlgoUnknown
	return fmt.Sprintf("PasswordKdfAlgoUnknown%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PasswordKdfAlgoUnknown) TypeID() uint32 {
	return PasswordKdfAlgoUnknownTypeID
}

// TypeName returns name of type in TL schema.
func (*PasswordKdfAlgoUnknown) TypeName() string {
	return "passwordKdfAlgoUnknown"
}

// TypeInfo returns info about TL type.
func (p *PasswordKdfAlgoUnknown) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "passwordKdfAlgoUnknown",
		ID:   PasswordKdfAlgoUnknownTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PasswordKdfAlgoUnknown) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passwordKdfAlgoUnknown#d45ab096 as nil")
	}
	b.PutID(PasswordKdfAlgoUnknownTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PasswordKdfAlgoUnknown) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passwordKdfAlgoUnknown#d45ab096 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PasswordKdfAlgoUnknown) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passwordKdfAlgoUnknown#d45ab096 to nil")
	}
	if err := b.ConsumeID(PasswordKdfAlgoUnknownTypeID); err != nil {
		return fmt.Errorf("unable to decode passwordKdfAlgoUnknown#d45ab096: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PasswordKdfAlgoUnknown) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passwordKdfAlgoUnknown#d45ab096 to nil")
	}
	return nil
}

// PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow represents TL type `passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a`.
// This key derivation algorithm defines that SRP 2FA login¹ must be used
//
// Links:
//  1. https://core.telegram.org/api/srp
//
// See https://core.telegram.org/constructor/passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow for reference.
type PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow struct {
	// One of two salts used by the derivation function (see SRP 2FA login¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	Salt1 []byte
	// One of two salts used by the derivation function (see SRP 2FA login¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	Salt2 []byte
	// Base (see SRP 2FA login¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	G int
	// 2048-bit modulus (see SRP 2FA login¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	P []byte
}

// PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowTypeID is TL type id of PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow.
const PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowTypeID = 0x3a912d4a

// construct implements constructor of PasswordKdfAlgoClass.
func (p PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) construct() PasswordKdfAlgoClass {
	return &p
}

// Ensuring interfaces in compile-time for PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow.
var (
	_ bin.Encoder     = &PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{}
	_ bin.Decoder     = &PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{}
	_ bin.BareEncoder = &PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{}
	_ bin.BareDecoder = &PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{}

	_ PasswordKdfAlgoClass = &PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{}
)

func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Salt1 == nil) {
		return false
	}
	if !(p.Salt2 == nil) {
		return false
	}
	if !(p.G == 0) {
		return false
	}
	if !(p.P == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) String() string {
	if p == nil {
		return "PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow(nil)"
	}
	type Alias PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow
	return fmt.Sprintf("PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow%+v", Alias(*p))
}

// FillFrom fills PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow from given interface.
func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) FillFrom(from interface {
	GetSalt1() (value []byte)
	GetSalt2() (value []byte)
	GetG() (value int)
	GetP() (value []byte)
}) {
	p.Salt1 = from.GetSalt1()
	p.Salt2 = from.GetSalt2()
	p.G = from.GetG()
	p.P = from.GetP()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) TypeID() uint32 {
	return PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowTypeID
}

// TypeName returns name of type in TL schema.
func (*PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) TypeName() string {
	return "passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow"
}

// TypeInfo returns info about TL type.
func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow",
		ID:   PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Salt1",
			SchemaName: "salt1",
		},
		{
			Name:       "Salt2",
			SchemaName: "salt2",
		},
		{
			Name:       "G",
			SchemaName: "g",
		},
		{
			Name:       "P",
			SchemaName: "p",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a as nil")
	}
	b.PutID(PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a as nil")
	}
	b.PutBytes(p.Salt1)
	b.PutBytes(p.Salt2)
	b.PutInt(p.G)
	b.PutBytes(p.P)
	return nil
}

// Decode implements bin.Decoder.
func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a to nil")
	}
	if err := b.ConsumeID(PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowTypeID); err != nil {
		return fmt.Errorf("unable to decode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a: field salt1: %w", err)
		}
		p.Salt1 = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a: field salt2: %w", err)
		}
		p.Salt2 = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a: field g: %w", err)
		}
		p.G = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a: field p: %w", err)
		}
		p.P = value
	}
	return nil
}

// GetSalt1 returns value of Salt1 field.
func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) GetSalt1() (value []byte) {
	if p == nil {
		return
	}
	return p.Salt1
}

// GetSalt2 returns value of Salt2 field.
func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) GetSalt2() (value []byte) {
	if p == nil {
		return
	}
	return p.Salt2
}

// GetG returns value of G field.
func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) GetG() (value int) {
	if p == nil {
		return
	}
	return p.G
}

// GetP returns value of P field.
func (p *PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) GetP() (value []byte) {
	if p == nil {
		return
	}
	return p.P
}

// PasswordKdfAlgoClassName is schema name of PasswordKdfAlgoClass.
const PasswordKdfAlgoClassName = "PasswordKdfAlgo"

// PasswordKdfAlgoClass represents PasswordKdfAlgo generic type.
//
// See https://core.telegram.org/type/PasswordKdfAlgo for reference.
//
// Example:
//
//	g, err := tg.DecodePasswordKdfAlgo(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.PasswordKdfAlgoUnknown: // passwordKdfAlgoUnknown#d45ab096
//	case *tg.PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow: // passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a
//	default: panic(v)
//	}
type PasswordKdfAlgoClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PasswordKdfAlgoClass

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
}

// DecodePasswordKdfAlgo implements binary de-serialization for PasswordKdfAlgoClass.
func DecodePasswordKdfAlgo(buf *bin.Buffer) (PasswordKdfAlgoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PasswordKdfAlgoUnknownTypeID:
		// Decoding passwordKdfAlgoUnknown#d45ab096.
		v := PasswordKdfAlgoUnknown{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PasswordKdfAlgoClass: %w", err)
		}
		return &v, nil
	case PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPowTypeID:
		// Decoding passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a.
		v := PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PasswordKdfAlgoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PasswordKdfAlgoClass: %w", bin.NewUnexpectedID(id))
	}
}

// PasswordKdfAlgo boxes the PasswordKdfAlgoClass providing a helper.
type PasswordKdfAlgoBox struct {
	PasswordKdfAlgo PasswordKdfAlgoClass
}

// Decode implements bin.Decoder for PasswordKdfAlgoBox.
func (b *PasswordKdfAlgoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PasswordKdfAlgoBox to nil")
	}
	v, err := DecodePasswordKdfAlgo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PasswordKdfAlgo = v
	return nil
}

// Encode implements bin.Encode for PasswordKdfAlgoBox.
func (b *PasswordKdfAlgoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PasswordKdfAlgo == nil {
		return fmt.Errorf("unable to encode PasswordKdfAlgoClass as nil")
	}
	return b.PasswordKdfAlgo.Encode(buf)
}
