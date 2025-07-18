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

// AccountSavedRingtone represents TL type `account.savedRingtone#b7263f6d`.
// The notification sound was already in MP3 format and was saved without any
// modification
//
// See https://core.telegram.org/constructor/account.savedRingtone for reference.
type AccountSavedRingtone struct {
}

// AccountSavedRingtoneTypeID is TL type id of AccountSavedRingtone.
const AccountSavedRingtoneTypeID = 0xb7263f6d

// construct implements constructor of AccountSavedRingtoneClass.
func (s AccountSavedRingtone) construct() AccountSavedRingtoneClass { return &s }

// Ensuring interfaces in compile-time for AccountSavedRingtone.
var (
	_ bin.Encoder     = &AccountSavedRingtone{}
	_ bin.Decoder     = &AccountSavedRingtone{}
	_ bin.BareEncoder = &AccountSavedRingtone{}
	_ bin.BareDecoder = &AccountSavedRingtone{}

	_ AccountSavedRingtoneClass = &AccountSavedRingtone{}
)

func (s *AccountSavedRingtone) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSavedRingtone) String() string {
	if s == nil {
		return "AccountSavedRingtone(nil)"
	}
	type Alias AccountSavedRingtone
	return fmt.Sprintf("AccountSavedRingtone%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSavedRingtone) TypeID() uint32 {
	return AccountSavedRingtoneTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSavedRingtone) TypeName() string {
	return "account.savedRingtone"
}

// TypeInfo returns info about TL type.
func (s *AccountSavedRingtone) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.savedRingtone",
		ID:   AccountSavedRingtoneTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSavedRingtone) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.savedRingtone#b7263f6d as nil")
	}
	b.PutID(AccountSavedRingtoneTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSavedRingtone) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.savedRingtone#b7263f6d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSavedRingtone) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.savedRingtone#b7263f6d to nil")
	}
	if err := b.ConsumeID(AccountSavedRingtoneTypeID); err != nil {
		return fmt.Errorf("unable to decode account.savedRingtone#b7263f6d: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSavedRingtone) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.savedRingtone#b7263f6d to nil")
	}
	return nil
}

// AccountSavedRingtoneConverted represents TL type `account.savedRingtoneConverted#1f307eb7`.
// The notification sound was not in MP3 format and was successfully converted and saved,
// use the returned Document¹ to refer to the notification sound from now on
//
// Links:
//  1. https://core.telegram.org/type/Document
//
// See https://core.telegram.org/constructor/account.savedRingtoneConverted for reference.
type AccountSavedRingtoneConverted struct {
	// The converted notification sound
	Document DocumentClass
}

// AccountSavedRingtoneConvertedTypeID is TL type id of AccountSavedRingtoneConverted.
const AccountSavedRingtoneConvertedTypeID = 0x1f307eb7

// construct implements constructor of AccountSavedRingtoneClass.
func (s AccountSavedRingtoneConverted) construct() AccountSavedRingtoneClass { return &s }

// Ensuring interfaces in compile-time for AccountSavedRingtoneConverted.
var (
	_ bin.Encoder     = &AccountSavedRingtoneConverted{}
	_ bin.Decoder     = &AccountSavedRingtoneConverted{}
	_ bin.BareEncoder = &AccountSavedRingtoneConverted{}
	_ bin.BareDecoder = &AccountSavedRingtoneConverted{}

	_ AccountSavedRingtoneClass = &AccountSavedRingtoneConverted{}
)

func (s *AccountSavedRingtoneConverted) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Document == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSavedRingtoneConverted) String() string {
	if s == nil {
		return "AccountSavedRingtoneConverted(nil)"
	}
	type Alias AccountSavedRingtoneConverted
	return fmt.Sprintf("AccountSavedRingtoneConverted%+v", Alias(*s))
}

// FillFrom fills AccountSavedRingtoneConverted from given interface.
func (s *AccountSavedRingtoneConverted) FillFrom(from interface {
	GetDocument() (value DocumentClass)
}) {
	s.Document = from.GetDocument()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSavedRingtoneConverted) TypeID() uint32 {
	return AccountSavedRingtoneConvertedTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSavedRingtoneConverted) TypeName() string {
	return "account.savedRingtoneConverted"
}

// TypeInfo returns info about TL type.
func (s *AccountSavedRingtoneConverted) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.savedRingtoneConverted",
		ID:   AccountSavedRingtoneConvertedTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Document",
			SchemaName: "document",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSavedRingtoneConverted) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.savedRingtoneConverted#1f307eb7 as nil")
	}
	b.PutID(AccountSavedRingtoneConvertedTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSavedRingtoneConverted) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.savedRingtoneConverted#1f307eb7 as nil")
	}
	if s.Document == nil {
		return fmt.Errorf("unable to encode account.savedRingtoneConverted#1f307eb7: field document is nil")
	}
	if err := s.Document.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.savedRingtoneConverted#1f307eb7: field document: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSavedRingtoneConverted) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.savedRingtoneConverted#1f307eb7 to nil")
	}
	if err := b.ConsumeID(AccountSavedRingtoneConvertedTypeID); err != nil {
		return fmt.Errorf("unable to decode account.savedRingtoneConverted#1f307eb7: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSavedRingtoneConverted) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.savedRingtoneConverted#1f307eb7 to nil")
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.savedRingtoneConverted#1f307eb7: field document: %w", err)
		}
		s.Document = value
	}
	return nil
}

// GetDocument returns value of Document field.
func (s *AccountSavedRingtoneConverted) GetDocument() (value DocumentClass) {
	if s == nil {
		return
	}
	return s.Document
}

// AccountSavedRingtoneClassName is schema name of AccountSavedRingtoneClass.
const AccountSavedRingtoneClassName = "account.SavedRingtone"

// AccountSavedRingtoneClass represents account.SavedRingtone generic type.
//
// See https://core.telegram.org/type/account.SavedRingtone for reference.
//
// Example:
//
//	g, err := tg.DecodeAccountSavedRingtone(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.AccountSavedRingtone: // account.savedRingtone#b7263f6d
//	case *tg.AccountSavedRingtoneConverted: // account.savedRingtoneConverted#1f307eb7
//	default: panic(v)
//	}
type AccountSavedRingtoneClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AccountSavedRingtoneClass

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

// DecodeAccountSavedRingtone implements binary de-serialization for AccountSavedRingtoneClass.
func DecodeAccountSavedRingtone(buf *bin.Buffer) (AccountSavedRingtoneClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AccountSavedRingtoneTypeID:
		// Decoding account.savedRingtone#b7263f6d.
		v := AccountSavedRingtone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountSavedRingtoneClass: %w", err)
		}
		return &v, nil
	case AccountSavedRingtoneConvertedTypeID:
		// Decoding account.savedRingtoneConverted#1f307eb7.
		v := AccountSavedRingtoneConverted{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountSavedRingtoneClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AccountSavedRingtoneClass: %w", bin.NewUnexpectedID(id))
	}
}

// AccountSavedRingtone boxes the AccountSavedRingtoneClass providing a helper.
type AccountSavedRingtoneBox struct {
	SavedRingtone AccountSavedRingtoneClass
}

// Decode implements bin.Decoder for AccountSavedRingtoneBox.
func (b *AccountSavedRingtoneBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AccountSavedRingtoneBox to nil")
	}
	v, err := DecodeAccountSavedRingtone(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SavedRingtone = v
	return nil
}

// Encode implements bin.Encode for AccountSavedRingtoneBox.
func (b *AccountSavedRingtoneBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SavedRingtone == nil {
		return fmt.Errorf("unable to encode AccountSavedRingtoneClass as nil")
	}
	return b.SavedRingtone.Encode(buf)
}
