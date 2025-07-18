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

// ChannelParticipantsRecent represents TL type `channelParticipantsRecent#de3f3c79`.
// Fetch only recent participants
//
// See https://core.telegram.org/constructor/channelParticipantsRecent for reference.
type ChannelParticipantsRecent struct {
}

// ChannelParticipantsRecentTypeID is TL type id of ChannelParticipantsRecent.
const ChannelParticipantsRecentTypeID = 0xde3f3c79

// construct implements constructor of ChannelParticipantsFilterClass.
func (c ChannelParticipantsRecent) construct() ChannelParticipantsFilterClass { return &c }

// Ensuring interfaces in compile-time for ChannelParticipantsRecent.
var (
	_ bin.Encoder     = &ChannelParticipantsRecent{}
	_ bin.Decoder     = &ChannelParticipantsRecent{}
	_ bin.BareEncoder = &ChannelParticipantsRecent{}
	_ bin.BareDecoder = &ChannelParticipantsRecent{}

	_ ChannelParticipantsFilterClass = &ChannelParticipantsRecent{}
)

func (c *ChannelParticipantsRecent) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelParticipantsRecent) String() string {
	if c == nil {
		return "ChannelParticipantsRecent(nil)"
	}
	type Alias ChannelParticipantsRecent
	return fmt.Sprintf("ChannelParticipantsRecent%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelParticipantsRecent) TypeID() uint32 {
	return ChannelParticipantsRecentTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelParticipantsRecent) TypeName() string {
	return "channelParticipantsRecent"
}

// TypeInfo returns info about TL type.
func (c *ChannelParticipantsRecent) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channelParticipantsRecent",
		ID:   ChannelParticipantsRecentTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelParticipantsRecent) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsRecent#de3f3c79 as nil")
	}
	b.PutID(ChannelParticipantsRecentTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelParticipantsRecent) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsRecent#de3f3c79 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelParticipantsRecent) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsRecent#de3f3c79 to nil")
	}
	if err := b.ConsumeID(ChannelParticipantsRecentTypeID); err != nil {
		return fmt.Errorf("unable to decode channelParticipantsRecent#de3f3c79: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelParticipantsRecent) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsRecent#de3f3c79 to nil")
	}
	return nil
}

// ChannelParticipantsAdmins represents TL type `channelParticipantsAdmins#b4608969`.
// Fetch only admin participants
//
// See https://core.telegram.org/constructor/channelParticipantsAdmins for reference.
type ChannelParticipantsAdmins struct {
}

// ChannelParticipantsAdminsTypeID is TL type id of ChannelParticipantsAdmins.
const ChannelParticipantsAdminsTypeID = 0xb4608969

// construct implements constructor of ChannelParticipantsFilterClass.
func (c ChannelParticipantsAdmins) construct() ChannelParticipantsFilterClass { return &c }

// Ensuring interfaces in compile-time for ChannelParticipantsAdmins.
var (
	_ bin.Encoder     = &ChannelParticipantsAdmins{}
	_ bin.Decoder     = &ChannelParticipantsAdmins{}
	_ bin.BareEncoder = &ChannelParticipantsAdmins{}
	_ bin.BareDecoder = &ChannelParticipantsAdmins{}

	_ ChannelParticipantsFilterClass = &ChannelParticipantsAdmins{}
)

func (c *ChannelParticipantsAdmins) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelParticipantsAdmins) String() string {
	if c == nil {
		return "ChannelParticipantsAdmins(nil)"
	}
	type Alias ChannelParticipantsAdmins
	return fmt.Sprintf("ChannelParticipantsAdmins%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelParticipantsAdmins) TypeID() uint32 {
	return ChannelParticipantsAdminsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelParticipantsAdmins) TypeName() string {
	return "channelParticipantsAdmins"
}

// TypeInfo returns info about TL type.
func (c *ChannelParticipantsAdmins) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channelParticipantsAdmins",
		ID:   ChannelParticipantsAdminsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelParticipantsAdmins) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsAdmins#b4608969 as nil")
	}
	b.PutID(ChannelParticipantsAdminsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelParticipantsAdmins) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsAdmins#b4608969 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelParticipantsAdmins) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsAdmins#b4608969 to nil")
	}
	if err := b.ConsumeID(ChannelParticipantsAdminsTypeID); err != nil {
		return fmt.Errorf("unable to decode channelParticipantsAdmins#b4608969: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelParticipantsAdmins) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsAdmins#b4608969 to nil")
	}
	return nil
}

// ChannelParticipantsKicked represents TL type `channelParticipantsKicked#a3b54985`.
// Fetch only kicked participants
//
// See https://core.telegram.org/constructor/channelParticipantsKicked for reference.
type ChannelParticipantsKicked struct {
	// Optional filter for searching kicked participants by name (otherwise empty)
	Q string
}

// ChannelParticipantsKickedTypeID is TL type id of ChannelParticipantsKicked.
const ChannelParticipantsKickedTypeID = 0xa3b54985

// construct implements constructor of ChannelParticipantsFilterClass.
func (c ChannelParticipantsKicked) construct() ChannelParticipantsFilterClass { return &c }

// Ensuring interfaces in compile-time for ChannelParticipantsKicked.
var (
	_ bin.Encoder     = &ChannelParticipantsKicked{}
	_ bin.Decoder     = &ChannelParticipantsKicked{}
	_ bin.BareEncoder = &ChannelParticipantsKicked{}
	_ bin.BareDecoder = &ChannelParticipantsKicked{}

	_ ChannelParticipantsFilterClass = &ChannelParticipantsKicked{}
)

func (c *ChannelParticipantsKicked) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Q == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelParticipantsKicked) String() string {
	if c == nil {
		return "ChannelParticipantsKicked(nil)"
	}
	type Alias ChannelParticipantsKicked
	return fmt.Sprintf("ChannelParticipantsKicked%+v", Alias(*c))
}

// FillFrom fills ChannelParticipantsKicked from given interface.
func (c *ChannelParticipantsKicked) FillFrom(from interface {
	GetQ() (value string)
}) {
	c.Q = from.GetQ()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelParticipantsKicked) TypeID() uint32 {
	return ChannelParticipantsKickedTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelParticipantsKicked) TypeName() string {
	return "channelParticipantsKicked"
}

// TypeInfo returns info about TL type.
func (c *ChannelParticipantsKicked) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channelParticipantsKicked",
		ID:   ChannelParticipantsKickedTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Q",
			SchemaName: "q",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelParticipantsKicked) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsKicked#a3b54985 as nil")
	}
	b.PutID(ChannelParticipantsKickedTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelParticipantsKicked) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsKicked#a3b54985 as nil")
	}
	b.PutString(c.Q)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelParticipantsKicked) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsKicked#a3b54985 to nil")
	}
	if err := b.ConsumeID(ChannelParticipantsKickedTypeID); err != nil {
		return fmt.Errorf("unable to decode channelParticipantsKicked#a3b54985: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelParticipantsKicked) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsKicked#a3b54985 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channelParticipantsKicked#a3b54985: field q: %w", err)
		}
		c.Q = value
	}
	return nil
}

// GetQ returns value of Q field.
func (c *ChannelParticipantsKicked) GetQ() (value string) {
	if c == nil {
		return
	}
	return c.Q
}

// ChannelParticipantsBots represents TL type `channelParticipantsBots#b0d1865b`.
// Fetch only bot participants
//
// See https://core.telegram.org/constructor/channelParticipantsBots for reference.
type ChannelParticipantsBots struct {
}

// ChannelParticipantsBotsTypeID is TL type id of ChannelParticipantsBots.
const ChannelParticipantsBotsTypeID = 0xb0d1865b

// construct implements constructor of ChannelParticipantsFilterClass.
func (c ChannelParticipantsBots) construct() ChannelParticipantsFilterClass { return &c }

// Ensuring interfaces in compile-time for ChannelParticipantsBots.
var (
	_ bin.Encoder     = &ChannelParticipantsBots{}
	_ bin.Decoder     = &ChannelParticipantsBots{}
	_ bin.BareEncoder = &ChannelParticipantsBots{}
	_ bin.BareDecoder = &ChannelParticipantsBots{}

	_ ChannelParticipantsFilterClass = &ChannelParticipantsBots{}
)

func (c *ChannelParticipantsBots) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelParticipantsBots) String() string {
	if c == nil {
		return "ChannelParticipantsBots(nil)"
	}
	type Alias ChannelParticipantsBots
	return fmt.Sprintf("ChannelParticipantsBots%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelParticipantsBots) TypeID() uint32 {
	return ChannelParticipantsBotsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelParticipantsBots) TypeName() string {
	return "channelParticipantsBots"
}

// TypeInfo returns info about TL type.
func (c *ChannelParticipantsBots) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channelParticipantsBots",
		ID:   ChannelParticipantsBotsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelParticipantsBots) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsBots#b0d1865b as nil")
	}
	b.PutID(ChannelParticipantsBotsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelParticipantsBots) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsBots#b0d1865b as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelParticipantsBots) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsBots#b0d1865b to nil")
	}
	if err := b.ConsumeID(ChannelParticipantsBotsTypeID); err != nil {
		return fmt.Errorf("unable to decode channelParticipantsBots#b0d1865b: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelParticipantsBots) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsBots#b0d1865b to nil")
	}
	return nil
}

// ChannelParticipantsBanned represents TL type `channelParticipantsBanned#1427a5e1`.
// Fetch only banned participants
//
// See https://core.telegram.org/constructor/channelParticipantsBanned for reference.
type ChannelParticipantsBanned struct {
	// Optional filter for searching banned participants by name (otherwise empty)
	Q string
}

// ChannelParticipantsBannedTypeID is TL type id of ChannelParticipantsBanned.
const ChannelParticipantsBannedTypeID = 0x1427a5e1

// construct implements constructor of ChannelParticipantsFilterClass.
func (c ChannelParticipantsBanned) construct() ChannelParticipantsFilterClass { return &c }

// Ensuring interfaces in compile-time for ChannelParticipantsBanned.
var (
	_ bin.Encoder     = &ChannelParticipantsBanned{}
	_ bin.Decoder     = &ChannelParticipantsBanned{}
	_ bin.BareEncoder = &ChannelParticipantsBanned{}
	_ bin.BareDecoder = &ChannelParticipantsBanned{}

	_ ChannelParticipantsFilterClass = &ChannelParticipantsBanned{}
)

func (c *ChannelParticipantsBanned) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Q == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelParticipantsBanned) String() string {
	if c == nil {
		return "ChannelParticipantsBanned(nil)"
	}
	type Alias ChannelParticipantsBanned
	return fmt.Sprintf("ChannelParticipantsBanned%+v", Alias(*c))
}

// FillFrom fills ChannelParticipantsBanned from given interface.
func (c *ChannelParticipantsBanned) FillFrom(from interface {
	GetQ() (value string)
}) {
	c.Q = from.GetQ()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelParticipantsBanned) TypeID() uint32 {
	return ChannelParticipantsBannedTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelParticipantsBanned) TypeName() string {
	return "channelParticipantsBanned"
}

// TypeInfo returns info about TL type.
func (c *ChannelParticipantsBanned) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channelParticipantsBanned",
		ID:   ChannelParticipantsBannedTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Q",
			SchemaName: "q",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelParticipantsBanned) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsBanned#1427a5e1 as nil")
	}
	b.PutID(ChannelParticipantsBannedTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelParticipantsBanned) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsBanned#1427a5e1 as nil")
	}
	b.PutString(c.Q)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelParticipantsBanned) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsBanned#1427a5e1 to nil")
	}
	if err := b.ConsumeID(ChannelParticipantsBannedTypeID); err != nil {
		return fmt.Errorf("unable to decode channelParticipantsBanned#1427a5e1: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelParticipantsBanned) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsBanned#1427a5e1 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channelParticipantsBanned#1427a5e1: field q: %w", err)
		}
		c.Q = value
	}
	return nil
}

// GetQ returns value of Q field.
func (c *ChannelParticipantsBanned) GetQ() (value string) {
	if c == nil {
		return
	}
	return c.Q
}

// ChannelParticipantsSearch represents TL type `channelParticipantsSearch#656ac4b`.
// Query participants by name
//
// See https://core.telegram.org/constructor/channelParticipantsSearch for reference.
type ChannelParticipantsSearch struct {
	// Search query
	Q string
}

// ChannelParticipantsSearchTypeID is TL type id of ChannelParticipantsSearch.
const ChannelParticipantsSearchTypeID = 0x656ac4b

// construct implements constructor of ChannelParticipantsFilterClass.
func (c ChannelParticipantsSearch) construct() ChannelParticipantsFilterClass { return &c }

// Ensuring interfaces in compile-time for ChannelParticipantsSearch.
var (
	_ bin.Encoder     = &ChannelParticipantsSearch{}
	_ bin.Decoder     = &ChannelParticipantsSearch{}
	_ bin.BareEncoder = &ChannelParticipantsSearch{}
	_ bin.BareDecoder = &ChannelParticipantsSearch{}

	_ ChannelParticipantsFilterClass = &ChannelParticipantsSearch{}
)

func (c *ChannelParticipantsSearch) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Q == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelParticipantsSearch) String() string {
	if c == nil {
		return "ChannelParticipantsSearch(nil)"
	}
	type Alias ChannelParticipantsSearch
	return fmt.Sprintf("ChannelParticipantsSearch%+v", Alias(*c))
}

// FillFrom fills ChannelParticipantsSearch from given interface.
func (c *ChannelParticipantsSearch) FillFrom(from interface {
	GetQ() (value string)
}) {
	c.Q = from.GetQ()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelParticipantsSearch) TypeID() uint32 {
	return ChannelParticipantsSearchTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelParticipantsSearch) TypeName() string {
	return "channelParticipantsSearch"
}

// TypeInfo returns info about TL type.
func (c *ChannelParticipantsSearch) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channelParticipantsSearch",
		ID:   ChannelParticipantsSearchTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Q",
			SchemaName: "q",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelParticipantsSearch) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsSearch#656ac4b as nil")
	}
	b.PutID(ChannelParticipantsSearchTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelParticipantsSearch) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsSearch#656ac4b as nil")
	}
	b.PutString(c.Q)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelParticipantsSearch) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsSearch#656ac4b to nil")
	}
	if err := b.ConsumeID(ChannelParticipantsSearchTypeID); err != nil {
		return fmt.Errorf("unable to decode channelParticipantsSearch#656ac4b: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelParticipantsSearch) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsSearch#656ac4b to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channelParticipantsSearch#656ac4b: field q: %w", err)
		}
		c.Q = value
	}
	return nil
}

// GetQ returns value of Q field.
func (c *ChannelParticipantsSearch) GetQ() (value string) {
	if c == nil {
		return
	}
	return c.Q
}

// ChannelParticipantsContacts represents TL type `channelParticipantsContacts#bb6ae88d`.
// Fetch only participants that are also contacts
//
// See https://core.telegram.org/constructor/channelParticipantsContacts for reference.
type ChannelParticipantsContacts struct {
	// Optional search query for searching contact participants by name
	Q string
}

// ChannelParticipantsContactsTypeID is TL type id of ChannelParticipantsContacts.
const ChannelParticipantsContactsTypeID = 0xbb6ae88d

// construct implements constructor of ChannelParticipantsFilterClass.
func (c ChannelParticipantsContacts) construct() ChannelParticipantsFilterClass { return &c }

// Ensuring interfaces in compile-time for ChannelParticipantsContacts.
var (
	_ bin.Encoder     = &ChannelParticipantsContacts{}
	_ bin.Decoder     = &ChannelParticipantsContacts{}
	_ bin.BareEncoder = &ChannelParticipantsContacts{}
	_ bin.BareDecoder = &ChannelParticipantsContacts{}

	_ ChannelParticipantsFilterClass = &ChannelParticipantsContacts{}
)

func (c *ChannelParticipantsContacts) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Q == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelParticipantsContacts) String() string {
	if c == nil {
		return "ChannelParticipantsContacts(nil)"
	}
	type Alias ChannelParticipantsContacts
	return fmt.Sprintf("ChannelParticipantsContacts%+v", Alias(*c))
}

// FillFrom fills ChannelParticipantsContacts from given interface.
func (c *ChannelParticipantsContacts) FillFrom(from interface {
	GetQ() (value string)
}) {
	c.Q = from.GetQ()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelParticipantsContacts) TypeID() uint32 {
	return ChannelParticipantsContactsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelParticipantsContacts) TypeName() string {
	return "channelParticipantsContacts"
}

// TypeInfo returns info about TL type.
func (c *ChannelParticipantsContacts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channelParticipantsContacts",
		ID:   ChannelParticipantsContactsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Q",
			SchemaName: "q",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelParticipantsContacts) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsContacts#bb6ae88d as nil")
	}
	b.PutID(ChannelParticipantsContactsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelParticipantsContacts) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsContacts#bb6ae88d as nil")
	}
	b.PutString(c.Q)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelParticipantsContacts) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsContacts#bb6ae88d to nil")
	}
	if err := b.ConsumeID(ChannelParticipantsContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode channelParticipantsContacts#bb6ae88d: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelParticipantsContacts) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsContacts#bb6ae88d to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channelParticipantsContacts#bb6ae88d: field q: %w", err)
		}
		c.Q = value
	}
	return nil
}

// GetQ returns value of Q field.
func (c *ChannelParticipantsContacts) GetQ() (value string) {
	if c == nil {
		return
	}
	return c.Q
}

// ChannelParticipantsMentions represents TL type `channelParticipantsMentions#e04b5ceb`.
// This filter is used when looking for supergroup members to mention.
// This filter will automatically remove anonymous admins, and return even
// non-participant users that replied to a specific thread¹ through the comment
// section² of a channel.
//
// Links:
//  1. https://core.telegram.org/api/threads
//  2. https://core.telegram.org/api/threads#channel-comments
//
// See https://core.telegram.org/constructor/channelParticipantsMentions for reference.
type ChannelParticipantsMentions struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Filter by user name or username
	//
	// Use SetQ and GetQ helpers.
	Q string
	// Look only for users that posted in this thread¹
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetTopMsgID and GetTopMsgID helpers.
	TopMsgID int
}

// ChannelParticipantsMentionsTypeID is TL type id of ChannelParticipantsMentions.
const ChannelParticipantsMentionsTypeID = 0xe04b5ceb

// construct implements constructor of ChannelParticipantsFilterClass.
func (c ChannelParticipantsMentions) construct() ChannelParticipantsFilterClass { return &c }

// Ensuring interfaces in compile-time for ChannelParticipantsMentions.
var (
	_ bin.Encoder     = &ChannelParticipantsMentions{}
	_ bin.Decoder     = &ChannelParticipantsMentions{}
	_ bin.BareEncoder = &ChannelParticipantsMentions{}
	_ bin.BareDecoder = &ChannelParticipantsMentions{}

	_ ChannelParticipantsFilterClass = &ChannelParticipantsMentions{}
)

func (c *ChannelParticipantsMentions) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Q == "") {
		return false
	}
	if !(c.TopMsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelParticipantsMentions) String() string {
	if c == nil {
		return "ChannelParticipantsMentions(nil)"
	}
	type Alias ChannelParticipantsMentions
	return fmt.Sprintf("ChannelParticipantsMentions%+v", Alias(*c))
}

// FillFrom fills ChannelParticipantsMentions from given interface.
func (c *ChannelParticipantsMentions) FillFrom(from interface {
	GetQ() (value string, ok bool)
	GetTopMsgID() (value int, ok bool)
}) {
	if val, ok := from.GetQ(); ok {
		c.Q = val
	}

	if val, ok := from.GetTopMsgID(); ok {
		c.TopMsgID = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelParticipantsMentions) TypeID() uint32 {
	return ChannelParticipantsMentionsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelParticipantsMentions) TypeName() string {
	return "channelParticipantsMentions"
}

// TypeInfo returns info about TL type.
func (c *ChannelParticipantsMentions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channelParticipantsMentions",
		ID:   ChannelParticipantsMentionsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Q",
			SchemaName: "q",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "TopMsgID",
			SchemaName: "top_msg_id",
			Null:       !c.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *ChannelParticipantsMentions) SetFlags() {
	if !(c.Q == "") {
		c.Flags.Set(0)
	}
	if !(c.TopMsgID == 0) {
		c.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (c *ChannelParticipantsMentions) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsMentions#e04b5ceb as nil")
	}
	b.PutID(ChannelParticipantsMentionsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelParticipantsMentions) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelParticipantsMentions#e04b5ceb as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channelParticipantsMentions#e04b5ceb: field flags: %w", err)
	}
	if c.Flags.Has(0) {
		b.PutString(c.Q)
	}
	if c.Flags.Has(1) {
		b.PutInt(c.TopMsgID)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelParticipantsMentions) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsMentions#e04b5ceb to nil")
	}
	if err := b.ConsumeID(ChannelParticipantsMentionsTypeID); err != nil {
		return fmt.Errorf("unable to decode channelParticipantsMentions#e04b5ceb: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelParticipantsMentions) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelParticipantsMentions#e04b5ceb to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channelParticipantsMentions#e04b5ceb: field flags: %w", err)
		}
	}
	if c.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channelParticipantsMentions#e04b5ceb: field q: %w", err)
		}
		c.Q = value
	}
	if c.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelParticipantsMentions#e04b5ceb: field top_msg_id: %w", err)
		}
		c.TopMsgID = value
	}
	return nil
}

// SetQ sets value of Q conditional field.
func (c *ChannelParticipantsMentions) SetQ(value string) {
	c.Flags.Set(0)
	c.Q = value
}

// GetQ returns value of Q conditional field and
// boolean which is true if field was set.
func (c *ChannelParticipantsMentions) GetQ() (value string, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(0) {
		return value, false
	}
	return c.Q, true
}

// SetTopMsgID sets value of TopMsgID conditional field.
func (c *ChannelParticipantsMentions) SetTopMsgID(value int) {
	c.Flags.Set(1)
	c.TopMsgID = value
}

// GetTopMsgID returns value of TopMsgID conditional field and
// boolean which is true if field was set.
func (c *ChannelParticipantsMentions) GetTopMsgID() (value int, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.TopMsgID, true
}

// ChannelParticipantsFilterClassName is schema name of ChannelParticipantsFilterClass.
const ChannelParticipantsFilterClassName = "ChannelParticipantsFilter"

// ChannelParticipantsFilterClass represents ChannelParticipantsFilter generic type.
//
// See https://core.telegram.org/type/ChannelParticipantsFilter for reference.
//
// Example:
//
//	g, err := tg.DecodeChannelParticipantsFilter(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.ChannelParticipantsRecent: // channelParticipantsRecent#de3f3c79
//	case *tg.ChannelParticipantsAdmins: // channelParticipantsAdmins#b4608969
//	case *tg.ChannelParticipantsKicked: // channelParticipantsKicked#a3b54985
//	case *tg.ChannelParticipantsBots: // channelParticipantsBots#b0d1865b
//	case *tg.ChannelParticipantsBanned: // channelParticipantsBanned#1427a5e1
//	case *tg.ChannelParticipantsSearch: // channelParticipantsSearch#656ac4b
//	case *tg.ChannelParticipantsContacts: // channelParticipantsContacts#bb6ae88d
//	case *tg.ChannelParticipantsMentions: // channelParticipantsMentions#e04b5ceb
//	default: panic(v)
//	}
type ChannelParticipantsFilterClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChannelParticipantsFilterClass

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

// DecodeChannelParticipantsFilter implements binary de-serialization for ChannelParticipantsFilterClass.
func DecodeChannelParticipantsFilter(buf *bin.Buffer) (ChannelParticipantsFilterClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChannelParticipantsRecentTypeID:
		// Decoding channelParticipantsRecent#de3f3c79.
		v := ChannelParticipantsRecent{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelParticipantsFilterClass: %w", err)
		}
		return &v, nil
	case ChannelParticipantsAdminsTypeID:
		// Decoding channelParticipantsAdmins#b4608969.
		v := ChannelParticipantsAdmins{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelParticipantsFilterClass: %w", err)
		}
		return &v, nil
	case ChannelParticipantsKickedTypeID:
		// Decoding channelParticipantsKicked#a3b54985.
		v := ChannelParticipantsKicked{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelParticipantsFilterClass: %w", err)
		}
		return &v, nil
	case ChannelParticipantsBotsTypeID:
		// Decoding channelParticipantsBots#b0d1865b.
		v := ChannelParticipantsBots{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelParticipantsFilterClass: %w", err)
		}
		return &v, nil
	case ChannelParticipantsBannedTypeID:
		// Decoding channelParticipantsBanned#1427a5e1.
		v := ChannelParticipantsBanned{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelParticipantsFilterClass: %w", err)
		}
		return &v, nil
	case ChannelParticipantsSearchTypeID:
		// Decoding channelParticipantsSearch#656ac4b.
		v := ChannelParticipantsSearch{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelParticipantsFilterClass: %w", err)
		}
		return &v, nil
	case ChannelParticipantsContactsTypeID:
		// Decoding channelParticipantsContacts#bb6ae88d.
		v := ChannelParticipantsContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelParticipantsFilterClass: %w", err)
		}
		return &v, nil
	case ChannelParticipantsMentionsTypeID:
		// Decoding channelParticipantsMentions#e04b5ceb.
		v := ChannelParticipantsMentions{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelParticipantsFilterClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChannelParticipantsFilterClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChannelParticipantsFilter boxes the ChannelParticipantsFilterClass providing a helper.
type ChannelParticipantsFilterBox struct {
	ChannelParticipantsFilter ChannelParticipantsFilterClass
}

// Decode implements bin.Decoder for ChannelParticipantsFilterBox.
func (b *ChannelParticipantsFilterBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChannelParticipantsFilterBox to nil")
	}
	v, err := DecodeChannelParticipantsFilter(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChannelParticipantsFilter = v
	return nil
}

// Encode implements bin.Encode for ChannelParticipantsFilterBox.
func (b *ChannelParticipantsFilterBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChannelParticipantsFilter == nil {
		return fmt.Errorf("unable to encode ChannelParticipantsFilterClass as nil")
	}
	return b.ChannelParticipantsFilter.Encode(buf)
}
