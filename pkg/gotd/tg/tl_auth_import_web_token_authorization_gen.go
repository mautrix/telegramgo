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

// AuthImportWebTokenAuthorizationRequest represents TL type `auth.importWebTokenAuthorization#2db873a9`.
// Login by importing an authorization token
//
// See https://core.telegram.org/method/auth.importWebTokenAuthorization for reference.
type AuthImportWebTokenAuthorizationRequest struct {
	// API ID¹
	//
	// Links:
	//  1) https://core.telegram.org/api/obtaining_api_id
	APIID int
	// API hash¹
	//
	// Links:
	//  1) https://core.telegram.org/api/obtaining_api_id
	APIHash string
	// The authorization token
	WebAuthToken string
}

// AuthImportWebTokenAuthorizationRequestTypeID is TL type id of AuthImportWebTokenAuthorizationRequest.
const AuthImportWebTokenAuthorizationRequestTypeID = 0x2db873a9

// Ensuring interfaces in compile-time for AuthImportWebTokenAuthorizationRequest.
var (
	_ bin.Encoder     = &AuthImportWebTokenAuthorizationRequest{}
	_ bin.Decoder     = &AuthImportWebTokenAuthorizationRequest{}
	_ bin.BareEncoder = &AuthImportWebTokenAuthorizationRequest{}
	_ bin.BareDecoder = &AuthImportWebTokenAuthorizationRequest{}
)

func (i *AuthImportWebTokenAuthorizationRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.APIID == 0) {
		return false
	}
	if !(i.APIHash == "") {
		return false
	}
	if !(i.WebAuthToken == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *AuthImportWebTokenAuthorizationRequest) String() string {
	if i == nil {
		return "AuthImportWebTokenAuthorizationRequest(nil)"
	}
	type Alias AuthImportWebTokenAuthorizationRequest
	return fmt.Sprintf("AuthImportWebTokenAuthorizationRequest%+v", Alias(*i))
}

// FillFrom fills AuthImportWebTokenAuthorizationRequest from given interface.
func (i *AuthImportWebTokenAuthorizationRequest) FillFrom(from interface {
	GetAPIID() (value int)
	GetAPIHash() (value string)
	GetWebAuthToken() (value string)
}) {
	i.APIID = from.GetAPIID()
	i.APIHash = from.GetAPIHash()
	i.WebAuthToken = from.GetWebAuthToken()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthImportWebTokenAuthorizationRequest) TypeID() uint32 {
	return AuthImportWebTokenAuthorizationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthImportWebTokenAuthorizationRequest) TypeName() string {
	return "auth.importWebTokenAuthorization"
}

// TypeInfo returns info about TL type.
func (i *AuthImportWebTokenAuthorizationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.importWebTokenAuthorization",
		ID:   AuthImportWebTokenAuthorizationRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "APIID",
			SchemaName: "api_id",
		},
		{
			Name:       "APIHash",
			SchemaName: "api_hash",
		},
		{
			Name:       "WebAuthToken",
			SchemaName: "web_auth_token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *AuthImportWebTokenAuthorizationRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode auth.importWebTokenAuthorization#2db873a9 as nil")
	}
	b.PutID(AuthImportWebTokenAuthorizationRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *AuthImportWebTokenAuthorizationRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode auth.importWebTokenAuthorization#2db873a9 as nil")
	}
	b.PutInt(i.APIID)
	b.PutString(i.APIHash)
	b.PutString(i.WebAuthToken)
	return nil
}

// Decode implements bin.Decoder.
func (i *AuthImportWebTokenAuthorizationRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode auth.importWebTokenAuthorization#2db873a9 to nil")
	}
	if err := b.ConsumeID(AuthImportWebTokenAuthorizationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.importWebTokenAuthorization#2db873a9: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *AuthImportWebTokenAuthorizationRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode auth.importWebTokenAuthorization#2db873a9 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importWebTokenAuthorization#2db873a9: field api_id: %w", err)
		}
		i.APIID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importWebTokenAuthorization#2db873a9: field api_hash: %w", err)
		}
		i.APIHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importWebTokenAuthorization#2db873a9: field web_auth_token: %w", err)
		}
		i.WebAuthToken = value
	}
	return nil
}

// GetAPIID returns value of APIID field.
func (i *AuthImportWebTokenAuthorizationRequest) GetAPIID() (value int) {
	if i == nil {
		return
	}
	return i.APIID
}

// GetAPIHash returns value of APIHash field.
func (i *AuthImportWebTokenAuthorizationRequest) GetAPIHash() (value string) {
	if i == nil {
		return
	}
	return i.APIHash
}

// GetWebAuthToken returns value of WebAuthToken field.
func (i *AuthImportWebTokenAuthorizationRequest) GetWebAuthToken() (value string) {
	if i == nil {
		return
	}
	return i.WebAuthToken
}

// AuthImportWebTokenAuthorization invokes method auth.importWebTokenAuthorization#2db873a9 returning error if any.
// Login by importing an authorization token
//
// Possible errors:
//
//	400 API_ID_INVALID: API ID invalid.
//
// See https://core.telegram.org/method/auth.importWebTokenAuthorization for reference.
func (c *Client) AuthImportWebTokenAuthorization(ctx context.Context, request *AuthImportWebTokenAuthorizationRequest) (AuthAuthorizationClass, error) {
	var result AuthAuthorizationBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Authorization, nil
}
