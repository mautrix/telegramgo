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

// AccountRegisterDeviceRequest represents TL type `account.registerDevice#ec86017a`.
// Register device to receive PUSH notifications¹
//
// Links:
//  1. https://core.telegram.org/api/push-updates
//
// See https://core.telegram.org/method/account.registerDevice for reference.
type AccountRegisterDeviceRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Avoid receiving (silent and invisible background) notifications. Useful to save
	// battery.
	NoMuted bool
	// Device token type, see PUSH updates¹ for the possible values.
	//
	// Links:
	//  1) https://core.telegram.org/api/push-updates#subscribing-to-notifications
	TokenType int
	// Device token, see PUSH updates¹ for the possible values.
	//
	// Links:
	//  1) https://core.telegram.org/api/push-updates#subscribing-to-notifications
	Token string
	// If (boolTrue)¹ is transmitted, a sandbox-certificate will be used during transmission.
	//
	// Links:
	//  1) https://core.telegram.org/constructor/boolTrue
	AppSandbox bool
	// For FCM and APNS VoIP, optional encryption key used to encrypt push notifications
	Secret []byte
	// List of user identifiers of other users currently using the client
	OtherUIDs []int64
}

// AccountRegisterDeviceRequestTypeID is TL type id of AccountRegisterDeviceRequest.
const AccountRegisterDeviceRequestTypeID = 0xec86017a

// Ensuring interfaces in compile-time for AccountRegisterDeviceRequest.
var (
	_ bin.Encoder     = &AccountRegisterDeviceRequest{}
	_ bin.Decoder     = &AccountRegisterDeviceRequest{}
	_ bin.BareEncoder = &AccountRegisterDeviceRequest{}
	_ bin.BareDecoder = &AccountRegisterDeviceRequest{}
)

func (r *AccountRegisterDeviceRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.NoMuted == false) {
		return false
	}
	if !(r.TokenType == 0) {
		return false
	}
	if !(r.Token == "") {
		return false
	}
	if !(r.AppSandbox == false) {
		return false
	}
	if !(r.Secret == nil) {
		return false
	}
	if !(r.OtherUIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountRegisterDeviceRequest) String() string {
	if r == nil {
		return "AccountRegisterDeviceRequest(nil)"
	}
	type Alias AccountRegisterDeviceRequest
	return fmt.Sprintf("AccountRegisterDeviceRequest%+v", Alias(*r))
}

// FillFrom fills AccountRegisterDeviceRequest from given interface.
func (r *AccountRegisterDeviceRequest) FillFrom(from interface {
	GetNoMuted() (value bool)
	GetTokenType() (value int)
	GetToken() (value string)
	GetAppSandbox() (value bool)
	GetSecret() (value []byte)
	GetOtherUIDs() (value []int64)
}) {
	r.NoMuted = from.GetNoMuted()
	r.TokenType = from.GetTokenType()
	r.Token = from.GetToken()
	r.AppSandbox = from.GetAppSandbox()
	r.Secret = from.GetSecret()
	r.OtherUIDs = from.GetOtherUIDs()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountRegisterDeviceRequest) TypeID() uint32 {
	return AccountRegisterDeviceRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountRegisterDeviceRequest) TypeName() string {
	return "account.registerDevice"
}

// TypeInfo returns info about TL type.
func (r *AccountRegisterDeviceRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.registerDevice",
		ID:   AccountRegisterDeviceRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "NoMuted",
			SchemaName: "no_muted",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "TokenType",
			SchemaName: "token_type",
		},
		{
			Name:       "Token",
			SchemaName: "token",
		},
		{
			Name:       "AppSandbox",
			SchemaName: "app_sandbox",
		},
		{
			Name:       "Secret",
			SchemaName: "secret",
		},
		{
			Name:       "OtherUIDs",
			SchemaName: "other_uids",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *AccountRegisterDeviceRequest) SetFlags() {
	if !(r.NoMuted == false) {
		r.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (r *AccountRegisterDeviceRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.registerDevice#ec86017a as nil")
	}
	b.PutID(AccountRegisterDeviceRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AccountRegisterDeviceRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.registerDevice#ec86017a as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.registerDevice#ec86017a: field flags: %w", err)
	}
	b.PutInt(r.TokenType)
	b.PutString(r.Token)
	b.PutBool(r.AppSandbox)
	b.PutBytes(r.Secret)
	b.PutVectorHeader(len(r.OtherUIDs))
	for _, v := range r.OtherUIDs {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountRegisterDeviceRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.registerDevice#ec86017a to nil")
	}
	if err := b.ConsumeID(AccountRegisterDeviceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.registerDevice#ec86017a: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AccountRegisterDeviceRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.registerDevice#ec86017a to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.registerDevice#ec86017a: field flags: %w", err)
		}
	}
	r.NoMuted = r.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.registerDevice#ec86017a: field token_type: %w", err)
		}
		r.TokenType = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.registerDevice#ec86017a: field token: %w", err)
		}
		r.Token = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.registerDevice#ec86017a: field app_sandbox: %w", err)
		}
		r.AppSandbox = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode account.registerDevice#ec86017a: field secret: %w", err)
		}
		r.Secret = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.registerDevice#ec86017a: field other_uids: %w", err)
		}

		if headerLen > 0 {
			r.OtherUIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode account.registerDevice#ec86017a: field other_uids: %w", err)
			}
			r.OtherUIDs = append(r.OtherUIDs, value)
		}
	}
	return nil
}

// SetNoMuted sets value of NoMuted conditional field.
func (r *AccountRegisterDeviceRequest) SetNoMuted(value bool) {
	if value {
		r.Flags.Set(0)
		r.NoMuted = true
	} else {
		r.Flags.Unset(0)
		r.NoMuted = false
	}
}

// GetNoMuted returns value of NoMuted conditional field.
func (r *AccountRegisterDeviceRequest) GetNoMuted() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(0)
}

// GetTokenType returns value of TokenType field.
func (r *AccountRegisterDeviceRequest) GetTokenType() (value int) {
	if r == nil {
		return
	}
	return r.TokenType
}

// GetToken returns value of Token field.
func (r *AccountRegisterDeviceRequest) GetToken() (value string) {
	if r == nil {
		return
	}
	return r.Token
}

// GetAppSandbox returns value of AppSandbox field.
func (r *AccountRegisterDeviceRequest) GetAppSandbox() (value bool) {
	if r == nil {
		return
	}
	return r.AppSandbox
}

// GetSecret returns value of Secret field.
func (r *AccountRegisterDeviceRequest) GetSecret() (value []byte) {
	if r == nil {
		return
	}
	return r.Secret
}

// GetOtherUIDs returns value of OtherUIDs field.
func (r *AccountRegisterDeviceRequest) GetOtherUIDs() (value []int64) {
	if r == nil {
		return
	}
	return r.OtherUIDs
}

// AccountRegisterDevice invokes method account.registerDevice#ec86017a returning error if any.
// Register device to receive PUSH notifications¹
//
// Links:
//  1. https://core.telegram.org/api/push-updates
//
// Possible errors:
//
//	400 TOKEN_EMPTY: The specified token is empty.
//	400 TOKEN_INVALID: The provided token is invalid.
//	400 TOKEN_TYPE_INVALID: The specified token type is invalid.
//	400 WEBPUSH_AUTH_INVALID: The specified web push authentication secret is invalid.
//	400 WEBPUSH_KEY_INVALID: The specified web push elliptic curve Diffie-Hellman public key is invalid.
//	400 WEBPUSH_TOKEN_INVALID: The specified web push token is invalid.
//
// See https://core.telegram.org/method/account.registerDevice for reference.
func (c *Client) AccountRegisterDevice(ctx context.Context, request *AccountRegisterDeviceRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
