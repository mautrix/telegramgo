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

// AccountEmojiStatusesNotModified represents TL type `account.emojiStatusesNotModified#d08ce645`.
// The server-side list of emoji statuses¹ hasn't changed
//
// Links:
//  1. https://core.telegram.org/api/emoji-status
//
// See https://core.telegram.org/constructor/account.emojiStatusesNotModified for reference.
type AccountEmojiStatusesNotModified struct {
}

// AccountEmojiStatusesNotModifiedTypeID is TL type id of AccountEmojiStatusesNotModified.
const AccountEmojiStatusesNotModifiedTypeID = 0xd08ce645

// construct implements constructor of AccountEmojiStatusesClass.
func (e AccountEmojiStatusesNotModified) construct() AccountEmojiStatusesClass { return &e }

// Ensuring interfaces in compile-time for AccountEmojiStatusesNotModified.
var (
	_ bin.Encoder     = &AccountEmojiStatusesNotModified{}
	_ bin.Decoder     = &AccountEmojiStatusesNotModified{}
	_ bin.BareEncoder = &AccountEmojiStatusesNotModified{}
	_ bin.BareDecoder = &AccountEmojiStatusesNotModified{}

	_ AccountEmojiStatusesClass = &AccountEmojiStatusesNotModified{}
)

func (e *AccountEmojiStatusesNotModified) Zero() bool {
	if e == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (e *AccountEmojiStatusesNotModified) String() string {
	if e == nil {
		return "AccountEmojiStatusesNotModified(nil)"
	}
	type Alias AccountEmojiStatusesNotModified
	return fmt.Sprintf("AccountEmojiStatusesNotModified%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountEmojiStatusesNotModified) TypeID() uint32 {
	return AccountEmojiStatusesNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountEmojiStatusesNotModified) TypeName() string {
	return "account.emojiStatusesNotModified"
}

// TypeInfo returns info about TL type.
func (e *AccountEmojiStatusesNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.emojiStatusesNotModified",
		ID:   AccountEmojiStatusesNotModifiedTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (e *AccountEmojiStatusesNotModified) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode account.emojiStatusesNotModified#d08ce645 as nil")
	}
	b.PutID(AccountEmojiStatusesNotModifiedTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *AccountEmojiStatusesNotModified) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode account.emojiStatusesNotModified#d08ce645 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *AccountEmojiStatusesNotModified) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode account.emojiStatusesNotModified#d08ce645 to nil")
	}
	if err := b.ConsumeID(AccountEmojiStatusesNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode account.emojiStatusesNotModified#d08ce645: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *AccountEmojiStatusesNotModified) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode account.emojiStatusesNotModified#d08ce645 to nil")
	}
	return nil
}

// AccountEmojiStatuses represents TL type `account.emojiStatuses#90c467d1`.
// A list of emoji statuses¹
//
// Links:
//  1. https://core.telegram.org/api/emoji-status
//
// See https://core.telegram.org/constructor/account.emojiStatuses for reference.
type AccountEmojiStatuses struct {
	// Hash used for caching, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
	// Emoji statuses¹
	//
	// Links:
	//  1) https://core.telegram.org/api/emoji-status
	Statuses []EmojiStatusClass
}

// AccountEmojiStatusesTypeID is TL type id of AccountEmojiStatuses.
const AccountEmojiStatusesTypeID = 0x90c467d1

// construct implements constructor of AccountEmojiStatusesClass.
func (e AccountEmojiStatuses) construct() AccountEmojiStatusesClass { return &e }

// Ensuring interfaces in compile-time for AccountEmojiStatuses.
var (
	_ bin.Encoder     = &AccountEmojiStatuses{}
	_ bin.Decoder     = &AccountEmojiStatuses{}
	_ bin.BareEncoder = &AccountEmojiStatuses{}
	_ bin.BareDecoder = &AccountEmojiStatuses{}

	_ AccountEmojiStatusesClass = &AccountEmojiStatuses{}
)

func (e *AccountEmojiStatuses) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Hash == 0) {
		return false
	}
	if !(e.Statuses == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *AccountEmojiStatuses) String() string {
	if e == nil {
		return "AccountEmojiStatuses(nil)"
	}
	type Alias AccountEmojiStatuses
	return fmt.Sprintf("AccountEmojiStatuses%+v", Alias(*e))
}

// FillFrom fills AccountEmojiStatuses from given interface.
func (e *AccountEmojiStatuses) FillFrom(from interface {
	GetHash() (value int64)
	GetStatuses() (value []EmojiStatusClass)
}) {
	e.Hash = from.GetHash()
	e.Statuses = from.GetStatuses()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountEmojiStatuses) TypeID() uint32 {
	return AccountEmojiStatusesTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountEmojiStatuses) TypeName() string {
	return "account.emojiStatuses"
}

// TypeInfo returns info about TL type.
func (e *AccountEmojiStatuses) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.emojiStatuses",
		ID:   AccountEmojiStatusesTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Statuses",
			SchemaName: "statuses",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *AccountEmojiStatuses) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode account.emojiStatuses#90c467d1 as nil")
	}
	b.PutID(AccountEmojiStatusesTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *AccountEmojiStatuses) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode account.emojiStatuses#90c467d1 as nil")
	}
	b.PutLong(e.Hash)
	b.PutVectorHeader(len(e.Statuses))
	for idx, v := range e.Statuses {
		if v == nil {
			return fmt.Errorf("unable to encode account.emojiStatuses#90c467d1: field statuses element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.emojiStatuses#90c467d1: field statuses element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *AccountEmojiStatuses) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode account.emojiStatuses#90c467d1 to nil")
	}
	if err := b.ConsumeID(AccountEmojiStatusesTypeID); err != nil {
		return fmt.Errorf("unable to decode account.emojiStatuses#90c467d1: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *AccountEmojiStatuses) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode account.emojiStatuses#90c467d1 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.emojiStatuses#90c467d1: field hash: %w", err)
		}
		e.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.emojiStatuses#90c467d1: field statuses: %w", err)
		}

		if headerLen > 0 {
			e.Statuses = make([]EmojiStatusClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeEmojiStatus(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.emojiStatuses#90c467d1: field statuses: %w", err)
			}
			e.Statuses = append(e.Statuses, value)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (e *AccountEmojiStatuses) GetHash() (value int64) {
	if e == nil {
		return
	}
	return e.Hash
}

// GetStatuses returns value of Statuses field.
func (e *AccountEmojiStatuses) GetStatuses() (value []EmojiStatusClass) {
	if e == nil {
		return
	}
	return e.Statuses
}

// MapStatuses returns field Statuses wrapped in EmojiStatusClassArray helper.
func (e *AccountEmojiStatuses) MapStatuses() (value EmojiStatusClassArray) {
	return EmojiStatusClassArray(e.Statuses)
}

// AccountEmojiStatusesClassName is schema name of AccountEmojiStatusesClass.
const AccountEmojiStatusesClassName = "account.EmojiStatuses"

// AccountEmojiStatusesClass represents account.EmojiStatuses generic type.
//
// See https://core.telegram.org/type/account.EmojiStatuses for reference.
//
// Example:
//
//	g, err := tg.DecodeAccountEmojiStatuses(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.AccountEmojiStatusesNotModified: // account.emojiStatusesNotModified#d08ce645
//	case *tg.AccountEmojiStatuses: // account.emojiStatuses#90c467d1
//	default: panic(v)
//	}
type AccountEmojiStatusesClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AccountEmojiStatusesClass

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

	// AsModified tries to map AccountEmojiStatusesClass to AccountEmojiStatuses.
	AsModified() (*AccountEmojiStatuses, bool)
}

// AsModified tries to map AccountEmojiStatusesNotModified to AccountEmojiStatuses.
func (e *AccountEmojiStatusesNotModified) AsModified() (*AccountEmojiStatuses, bool) {
	return nil, false
}

// AsModified tries to map AccountEmojiStatuses to AccountEmojiStatuses.
func (e *AccountEmojiStatuses) AsModified() (*AccountEmojiStatuses, bool) {
	return e, true
}

// DecodeAccountEmojiStatuses implements binary de-serialization for AccountEmojiStatusesClass.
func DecodeAccountEmojiStatuses(buf *bin.Buffer) (AccountEmojiStatusesClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AccountEmojiStatusesNotModifiedTypeID:
		// Decoding account.emojiStatusesNotModified#d08ce645.
		v := AccountEmojiStatusesNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountEmojiStatusesClass: %w", err)
		}
		return &v, nil
	case AccountEmojiStatusesTypeID:
		// Decoding account.emojiStatuses#90c467d1.
		v := AccountEmojiStatuses{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountEmojiStatusesClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AccountEmojiStatusesClass: %w", bin.NewUnexpectedID(id))
	}
}

// AccountEmojiStatuses boxes the AccountEmojiStatusesClass providing a helper.
type AccountEmojiStatusesBox struct {
	EmojiStatuses AccountEmojiStatusesClass
}

// Decode implements bin.Decoder for AccountEmojiStatusesBox.
func (b *AccountEmojiStatusesBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AccountEmojiStatusesBox to nil")
	}
	v, err := DecodeAccountEmojiStatuses(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EmojiStatuses = v
	return nil
}

// Encode implements bin.Encode for AccountEmojiStatusesBox.
func (b *AccountEmojiStatusesBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.EmojiStatuses == nil {
		return fmt.Errorf("unable to encode AccountEmojiStatusesClass as nil")
	}
	return b.EmojiStatuses.Encode(buf)
}
