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

// MissingInvitee represents TL type `missingInvitee#628c9224`.
// Info about why a specific user could not be invited »¹.
//
// Links:
//  1. https://core.telegram.org/api/invites#direct-invites
//
// See https://core.telegram.org/constructor/missingInvitee for reference.
type MissingInvitee struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, we could not add the user only because the current account needs to purchase a
	// Telegram Premium¹ subscription to complete the operation.
	//
	// Links:
	//  1) https://core.telegram.org/api/premium
	PremiumWouldAllowInvite bool
	// If set, we could not add the user because of their privacy settings, and additionally,
	// the current account needs to purchase a Telegram Premium¹ subscription to directly
	// share an invite link with the user via a private message.
	//
	// Links:
	//  1) https://core.telegram.org/api/premium
	PremiumRequiredForPm bool
	// ID of the user. If neither of the flags below are set, we could not add the user
	// because of their privacy settings, and we can create and directly share an invite
	// link¹ with them using a normal message, instead.
	//
	// Links:
	//  1) https://core.telegram.org/api/invites#invite-links
	UserID int64
}

// MissingInviteeTypeID is TL type id of MissingInvitee.
const MissingInviteeTypeID = 0x628c9224

// Ensuring interfaces in compile-time for MissingInvitee.
var (
	_ bin.Encoder     = &MissingInvitee{}
	_ bin.Decoder     = &MissingInvitee{}
	_ bin.BareEncoder = &MissingInvitee{}
	_ bin.BareDecoder = &MissingInvitee{}
)

func (m *MissingInvitee) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.PremiumWouldAllowInvite == false) {
		return false
	}
	if !(m.PremiumRequiredForPm == false) {
		return false
	}
	if !(m.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MissingInvitee) String() string {
	if m == nil {
		return "MissingInvitee(nil)"
	}
	type Alias MissingInvitee
	return fmt.Sprintf("MissingInvitee%+v", Alias(*m))
}

// FillFrom fills MissingInvitee from given interface.
func (m *MissingInvitee) FillFrom(from interface {
	GetPremiumWouldAllowInvite() (value bool)
	GetPremiumRequiredForPm() (value bool)
	GetUserID() (value int64)
}) {
	m.PremiumWouldAllowInvite = from.GetPremiumWouldAllowInvite()
	m.PremiumRequiredForPm = from.GetPremiumRequiredForPm()
	m.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MissingInvitee) TypeID() uint32 {
	return MissingInviteeTypeID
}

// TypeName returns name of type in TL schema.
func (*MissingInvitee) TypeName() string {
	return "missingInvitee"
}

// TypeInfo returns info about TL type.
func (m *MissingInvitee) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "missingInvitee",
		ID:   MissingInviteeTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PremiumWouldAllowInvite",
			SchemaName: "premium_would_allow_invite",
			Null:       !m.Flags.Has(0),
		},
		{
			Name:       "PremiumRequiredForPm",
			SchemaName: "premium_required_for_pm",
			Null:       !m.Flags.Has(1),
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (m *MissingInvitee) SetFlags() {
	if !(m.PremiumWouldAllowInvite == false) {
		m.Flags.Set(0)
	}
	if !(m.PremiumRequiredForPm == false) {
		m.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (m *MissingInvitee) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode missingInvitee#628c9224 as nil")
	}
	b.PutID(MissingInviteeTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MissingInvitee) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode missingInvitee#628c9224 as nil")
	}
	m.SetFlags()
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode missingInvitee#628c9224: field flags: %w", err)
	}
	b.PutLong(m.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (m *MissingInvitee) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode missingInvitee#628c9224 to nil")
	}
	if err := b.ConsumeID(MissingInviteeTypeID); err != nil {
		return fmt.Errorf("unable to decode missingInvitee#628c9224: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MissingInvitee) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode missingInvitee#628c9224 to nil")
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode missingInvitee#628c9224: field flags: %w", err)
		}
	}
	m.PremiumWouldAllowInvite = m.Flags.Has(0)
	m.PremiumRequiredForPm = m.Flags.Has(1)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode missingInvitee#628c9224: field user_id: %w", err)
		}
		m.UserID = value
	}
	return nil
}

// SetPremiumWouldAllowInvite sets value of PremiumWouldAllowInvite conditional field.
func (m *MissingInvitee) SetPremiumWouldAllowInvite(value bool) {
	if value {
		m.Flags.Set(0)
		m.PremiumWouldAllowInvite = true
	} else {
		m.Flags.Unset(0)
		m.PremiumWouldAllowInvite = false
	}
}

// GetPremiumWouldAllowInvite returns value of PremiumWouldAllowInvite conditional field.
func (m *MissingInvitee) GetPremiumWouldAllowInvite() (value bool) {
	if m == nil {
		return
	}
	return m.Flags.Has(0)
}

// SetPremiumRequiredForPm sets value of PremiumRequiredForPm conditional field.
func (m *MissingInvitee) SetPremiumRequiredForPm(value bool) {
	if value {
		m.Flags.Set(1)
		m.PremiumRequiredForPm = true
	} else {
		m.Flags.Unset(1)
		m.PremiumRequiredForPm = false
	}
}

// GetPremiumRequiredForPm returns value of PremiumRequiredForPm conditional field.
func (m *MissingInvitee) GetPremiumRequiredForPm() (value bool) {
	if m == nil {
		return
	}
	return m.Flags.Has(1)
}

// GetUserID returns value of UserID field.
func (m *MissingInvitee) GetUserID() (value int64) {
	if m == nil {
		return
	}
	return m.UserID
}
