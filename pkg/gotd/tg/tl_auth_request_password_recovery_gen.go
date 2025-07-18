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

// AuthRequestPasswordRecoveryRequest represents TL type `auth.requestPasswordRecovery#d897bc66`.
// Request recovery code of a 2FA password¹, only for accounts with a recovery email
// configured².
//
// Links:
//  1. https://core.telegram.org/api/srp
//  2. https://core.telegram.org/api/srp#email-verification
//
// See https://core.telegram.org/method/auth.requestPasswordRecovery for reference.
type AuthRequestPasswordRecoveryRequest struct {
}

// AuthRequestPasswordRecoveryRequestTypeID is TL type id of AuthRequestPasswordRecoveryRequest.
const AuthRequestPasswordRecoveryRequestTypeID = 0xd897bc66

// Ensuring interfaces in compile-time for AuthRequestPasswordRecoveryRequest.
var (
	_ bin.Encoder     = &AuthRequestPasswordRecoveryRequest{}
	_ bin.Decoder     = &AuthRequestPasswordRecoveryRequest{}
	_ bin.BareEncoder = &AuthRequestPasswordRecoveryRequest{}
	_ bin.BareDecoder = &AuthRequestPasswordRecoveryRequest{}
)

func (r *AuthRequestPasswordRecoveryRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *AuthRequestPasswordRecoveryRequest) String() string {
	if r == nil {
		return "AuthRequestPasswordRecoveryRequest(nil)"
	}
	type Alias AuthRequestPasswordRecoveryRequest
	return fmt.Sprintf("AuthRequestPasswordRecoveryRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthRequestPasswordRecoveryRequest) TypeID() uint32 {
	return AuthRequestPasswordRecoveryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthRequestPasswordRecoveryRequest) TypeName() string {
	return "auth.requestPasswordRecovery"
}

// TypeInfo returns info about TL type.
func (r *AuthRequestPasswordRecoveryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.requestPasswordRecovery",
		ID:   AuthRequestPasswordRecoveryRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *AuthRequestPasswordRecoveryRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.requestPasswordRecovery#d897bc66 as nil")
	}
	b.PutID(AuthRequestPasswordRecoveryRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AuthRequestPasswordRecoveryRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.requestPasswordRecovery#d897bc66 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *AuthRequestPasswordRecoveryRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.requestPasswordRecovery#d897bc66 to nil")
	}
	if err := b.ConsumeID(AuthRequestPasswordRecoveryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.requestPasswordRecovery#d897bc66: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AuthRequestPasswordRecoveryRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.requestPasswordRecovery#d897bc66 to nil")
	}
	return nil
}

// AuthRequestPasswordRecovery invokes method auth.requestPasswordRecovery#d897bc66 returning error if any.
// Request recovery code of a 2FA password¹, only for accounts with a recovery email
// configured².
//
// Links:
//  1. https://core.telegram.org/api/srp
//  2. https://core.telegram.org/api/srp#email-verification
//
// Possible errors:
//
//	400 PASSWORD_EMPTY: The provided password is empty.
//	400 PASSWORD_RECOVERY_NA: No email was set, can't recover password via email.
//
// See https://core.telegram.org/method/auth.requestPasswordRecovery for reference.
func (c *Client) AuthRequestPasswordRecovery(ctx context.Context) (*AuthPasswordRecovery, error) {
	var result AuthPasswordRecovery

	request := &AuthRequestPasswordRecoveryRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
