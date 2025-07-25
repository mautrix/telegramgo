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

// PasswordState represents TL type `passwordState#1cd63828`.
type PasswordState struct {
	// True, if a 2-step verification password is set
	HasPassword bool
	// Hint for the password; may be empty
	PasswordHint string
	// True, if a recovery email is set
	HasRecoveryEmailAddress bool
	// True, if some Telegram Passport elements were saved
	HasPassportData bool
	// Information about the recovery email address to which the confirmation email was sent;
	// may be null
	RecoveryEmailAddressCodeInfo EmailAddressAuthenticationCodeInfo
	// Pattern of the email address set up for logging in
	LoginEmailAddressPattern string
	// If not 0, point in time (Unix timestamp) after which the 2-step verification password
	// can be reset immediately using resetPassword
	PendingResetDate int32
}

// PasswordStateTypeID is TL type id of PasswordState.
const PasswordStateTypeID = 0x1cd63828

// Ensuring interfaces in compile-time for PasswordState.
var (
	_ bin.Encoder     = &PasswordState{}
	_ bin.Decoder     = &PasswordState{}
	_ bin.BareEncoder = &PasswordState{}
	_ bin.BareDecoder = &PasswordState{}
)

func (p *PasswordState) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.HasPassword == false) {
		return false
	}
	if !(p.PasswordHint == "") {
		return false
	}
	if !(p.HasRecoveryEmailAddress == false) {
		return false
	}
	if !(p.HasPassportData == false) {
		return false
	}
	if !(p.RecoveryEmailAddressCodeInfo.Zero()) {
		return false
	}
	if !(p.LoginEmailAddressPattern == "") {
		return false
	}
	if !(p.PendingResetDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PasswordState) String() string {
	if p == nil {
		return "PasswordState(nil)"
	}
	type Alias PasswordState
	return fmt.Sprintf("PasswordState%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PasswordState) TypeID() uint32 {
	return PasswordStateTypeID
}

// TypeName returns name of type in TL schema.
func (*PasswordState) TypeName() string {
	return "passwordState"
}

// TypeInfo returns info about TL type.
func (p *PasswordState) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "passwordState",
		ID:   PasswordStateTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "HasPassword",
			SchemaName: "has_password",
		},
		{
			Name:       "PasswordHint",
			SchemaName: "password_hint",
		},
		{
			Name:       "HasRecoveryEmailAddress",
			SchemaName: "has_recovery_email_address",
		},
		{
			Name:       "HasPassportData",
			SchemaName: "has_passport_data",
		},
		{
			Name:       "RecoveryEmailAddressCodeInfo",
			SchemaName: "recovery_email_address_code_info",
		},
		{
			Name:       "LoginEmailAddressPattern",
			SchemaName: "login_email_address_pattern",
		},
		{
			Name:       "PendingResetDate",
			SchemaName: "pending_reset_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PasswordState) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passwordState#1cd63828 as nil")
	}
	b.PutID(PasswordStateTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PasswordState) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passwordState#1cd63828 as nil")
	}
	b.PutBool(p.HasPassword)
	b.PutString(p.PasswordHint)
	b.PutBool(p.HasRecoveryEmailAddress)
	b.PutBool(p.HasPassportData)
	if err := p.RecoveryEmailAddressCodeInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode passwordState#1cd63828: field recovery_email_address_code_info: %w", err)
	}
	b.PutString(p.LoginEmailAddressPattern)
	b.PutInt32(p.PendingResetDate)
	return nil
}

// Decode implements bin.Decoder.
func (p *PasswordState) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passwordState#1cd63828 to nil")
	}
	if err := b.ConsumeID(PasswordStateTypeID); err != nil {
		return fmt.Errorf("unable to decode passwordState#1cd63828: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PasswordState) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passwordState#1cd63828 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode passwordState#1cd63828: field has_password: %w", err)
		}
		p.HasPassword = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode passwordState#1cd63828: field password_hint: %w", err)
		}
		p.PasswordHint = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode passwordState#1cd63828: field has_recovery_email_address: %w", err)
		}
		p.HasRecoveryEmailAddress = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode passwordState#1cd63828: field has_passport_data: %w", err)
		}
		p.HasPassportData = value
	}
	{
		if err := p.RecoveryEmailAddressCodeInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode passwordState#1cd63828: field recovery_email_address_code_info: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode passwordState#1cd63828: field login_email_address_pattern: %w", err)
		}
		p.LoginEmailAddressPattern = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode passwordState#1cd63828: field pending_reset_date: %w", err)
		}
		p.PendingResetDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PasswordState) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode passwordState#1cd63828 as nil")
	}
	b.ObjStart()
	b.PutID("passwordState")
	b.Comma()
	b.FieldStart("has_password")
	b.PutBool(p.HasPassword)
	b.Comma()
	b.FieldStart("password_hint")
	b.PutString(p.PasswordHint)
	b.Comma()
	b.FieldStart("has_recovery_email_address")
	b.PutBool(p.HasRecoveryEmailAddress)
	b.Comma()
	b.FieldStart("has_passport_data")
	b.PutBool(p.HasPassportData)
	b.Comma()
	b.FieldStart("recovery_email_address_code_info")
	if err := p.RecoveryEmailAddressCodeInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode passwordState#1cd63828: field recovery_email_address_code_info: %w", err)
	}
	b.Comma()
	b.FieldStart("login_email_address_pattern")
	b.PutString(p.LoginEmailAddressPattern)
	b.Comma()
	b.FieldStart("pending_reset_date")
	b.PutInt32(p.PendingResetDate)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PasswordState) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode passwordState#1cd63828 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("passwordState"); err != nil {
				return fmt.Errorf("unable to decode passwordState#1cd63828: %w", err)
			}
		case "has_password":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode passwordState#1cd63828: field has_password: %w", err)
			}
			p.HasPassword = value
		case "password_hint":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode passwordState#1cd63828: field password_hint: %w", err)
			}
			p.PasswordHint = value
		case "has_recovery_email_address":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode passwordState#1cd63828: field has_recovery_email_address: %w", err)
			}
			p.HasRecoveryEmailAddress = value
		case "has_passport_data":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode passwordState#1cd63828: field has_passport_data: %w", err)
			}
			p.HasPassportData = value
		case "recovery_email_address_code_info":
			if err := p.RecoveryEmailAddressCodeInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode passwordState#1cd63828: field recovery_email_address_code_info: %w", err)
			}
		case "login_email_address_pattern":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode passwordState#1cd63828: field login_email_address_pattern: %w", err)
			}
			p.LoginEmailAddressPattern = value
		case "pending_reset_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode passwordState#1cd63828: field pending_reset_date: %w", err)
			}
			p.PendingResetDate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetHasPassword returns value of HasPassword field.
func (p *PasswordState) GetHasPassword() (value bool) {
	if p == nil {
		return
	}
	return p.HasPassword
}

// GetPasswordHint returns value of PasswordHint field.
func (p *PasswordState) GetPasswordHint() (value string) {
	if p == nil {
		return
	}
	return p.PasswordHint
}

// GetHasRecoveryEmailAddress returns value of HasRecoveryEmailAddress field.
func (p *PasswordState) GetHasRecoveryEmailAddress() (value bool) {
	if p == nil {
		return
	}
	return p.HasRecoveryEmailAddress
}

// GetHasPassportData returns value of HasPassportData field.
func (p *PasswordState) GetHasPassportData() (value bool) {
	if p == nil {
		return
	}
	return p.HasPassportData
}

// GetRecoveryEmailAddressCodeInfo returns value of RecoveryEmailAddressCodeInfo field.
func (p *PasswordState) GetRecoveryEmailAddressCodeInfo() (value EmailAddressAuthenticationCodeInfo) {
	if p == nil {
		return
	}
	return p.RecoveryEmailAddressCodeInfo
}

// GetLoginEmailAddressPattern returns value of LoginEmailAddressPattern field.
func (p *PasswordState) GetLoginEmailAddressPattern() (value string) {
	if p == nil {
		return
	}
	return p.LoginEmailAddressPattern
}

// GetPendingResetDate returns value of PendingResetDate field.
func (p *PasswordState) GetPendingResetDate() (value int32) {
	if p == nil {
		return
	}
	return p.PendingResetDate
}
