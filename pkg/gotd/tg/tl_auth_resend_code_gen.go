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

// AuthResendCodeRequest represents TL type `auth.resendCode#cae47523`.
// Resend the login code via another medium, the phone code type is determined by the
// return value of the previous auth.sendCode/auth.resendCode: see login¹ for more info.
//
// Links:
//  1. https://core.telegram.org/api/auth
//
// See https://core.telegram.org/method/auth.resendCode for reference.
type AuthResendCodeRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// The phone number
	PhoneNumber string
	// The phone code hash obtained from auth.sendCode¹
	//
	// Links:
	//  1) https://core.telegram.org/method/auth.sendCode
	PhoneCodeHash string
	// Official clients only, used if the device integrity verification failed, and no secret
	// could be obtained to invoke auth.requestFirebaseSms¹: in this case, the device
	// integrity verification failure reason must be passed here.
	//
	// Links:
	//  1) https://core.telegram.org/method/auth.requestFirebaseSms
	//
	// Use SetReason and GetReason helpers.
	Reason string
}

// AuthResendCodeRequestTypeID is TL type id of AuthResendCodeRequest.
const AuthResendCodeRequestTypeID = 0xcae47523

// Ensuring interfaces in compile-time for AuthResendCodeRequest.
var (
	_ bin.Encoder     = &AuthResendCodeRequest{}
	_ bin.Decoder     = &AuthResendCodeRequest{}
	_ bin.BareEncoder = &AuthResendCodeRequest{}
	_ bin.BareDecoder = &AuthResendCodeRequest{}
)

func (r *AuthResendCodeRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.PhoneNumber == "") {
		return false
	}
	if !(r.PhoneCodeHash == "") {
		return false
	}
	if !(r.Reason == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AuthResendCodeRequest) String() string {
	if r == nil {
		return "AuthResendCodeRequest(nil)"
	}
	type Alias AuthResendCodeRequest
	return fmt.Sprintf("AuthResendCodeRequest%+v", Alias(*r))
}

// FillFrom fills AuthResendCodeRequest from given interface.
func (r *AuthResendCodeRequest) FillFrom(from interface {
	GetPhoneNumber() (value string)
	GetPhoneCodeHash() (value string)
	GetReason() (value string, ok bool)
}) {
	r.PhoneNumber = from.GetPhoneNumber()
	r.PhoneCodeHash = from.GetPhoneCodeHash()
	if val, ok := from.GetReason(); ok {
		r.Reason = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthResendCodeRequest) TypeID() uint32 {
	return AuthResendCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthResendCodeRequest) TypeName() string {
	return "auth.resendCode"
}

// TypeInfo returns info about TL type.
func (r *AuthResendCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.resendCode",
		ID:   AuthResendCodeRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
		{
			Name:       "PhoneCodeHash",
			SchemaName: "phone_code_hash",
		},
		{
			Name:       "Reason",
			SchemaName: "reason",
			Null:       !r.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *AuthResendCodeRequest) SetFlags() {
	if !(r.Reason == "") {
		r.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (r *AuthResendCodeRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.resendCode#cae47523 as nil")
	}
	b.PutID(AuthResendCodeRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AuthResendCodeRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.resendCode#cae47523 as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.resendCode#cae47523: field flags: %w", err)
	}
	b.PutString(r.PhoneNumber)
	b.PutString(r.PhoneCodeHash)
	if r.Flags.Has(0) {
		b.PutString(r.Reason)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *AuthResendCodeRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.resendCode#cae47523 to nil")
	}
	if err := b.ConsumeID(AuthResendCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.resendCode#cae47523: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AuthResendCodeRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.resendCode#cae47523 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode auth.resendCode#cae47523: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.resendCode#cae47523: field phone_number: %w", err)
		}
		r.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.resendCode#cae47523: field phone_code_hash: %w", err)
		}
		r.PhoneCodeHash = value
	}
	if r.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.resendCode#cae47523: field reason: %w", err)
		}
		r.Reason = value
	}
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (r *AuthResendCodeRequest) GetPhoneNumber() (value string) {
	if r == nil {
		return
	}
	return r.PhoneNumber
}

// GetPhoneCodeHash returns value of PhoneCodeHash field.
func (r *AuthResendCodeRequest) GetPhoneCodeHash() (value string) {
	if r == nil {
		return
	}
	return r.PhoneCodeHash
}

// SetReason sets value of Reason conditional field.
func (r *AuthResendCodeRequest) SetReason(value string) {
	r.Flags.Set(0)
	r.Reason = value
}

// GetReason returns value of Reason conditional field and
// boolean which is true if field was set.
func (r *AuthResendCodeRequest) GetReason() (value string, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(0) {
		return value, false
	}
	return r.Reason, true
}

// AuthResendCode invokes method auth.resendCode#cae47523 returning error if any.
// Resend the login code via another medium, the phone code type is determined by the
// return value of the previous auth.sendCode/auth.resendCode: see login¹ for more info.
//
// Links:
//  1. https://core.telegram.org/api/auth
//
// Possible errors:
//
//	400 PHONE_CODE_EMPTY: phone_code is missing.
//	400 PHONE_CODE_EXPIRED: The phone code you provided has expired.
//	400 PHONE_CODE_HASH_EMPTY: phone_code_hash is missing.
//	406 PHONE_NUMBER_INVALID: The phone number is invalid.
//	406 SEND_CODE_UNAVAILABLE: Returned when all available options for this type of number were already used (e.g. flash-call, then SMS, then this error might be returned to trigger a second resend).
//
// See https://core.telegram.org/method/auth.resendCode for reference.
func (c *Client) AuthResendCode(ctx context.Context, request *AuthResendCodeRequest) (AuthSentCodeClass, error) {
	var result AuthSentCodeBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.SentCode, nil
}
