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

// AccountGetWebAuthorizationsRequest represents TL type `account.getWebAuthorizations#182e6d6f`.
// Get web login widget¹ authorizations
//
// Links:
//  1. https://core.telegram.org/widgets/login
//
// See https://core.telegram.org/method/account.getWebAuthorizations for reference.
type AccountGetWebAuthorizationsRequest struct {
}

// AccountGetWebAuthorizationsRequestTypeID is TL type id of AccountGetWebAuthorizationsRequest.
const AccountGetWebAuthorizationsRequestTypeID = 0x182e6d6f

// Ensuring interfaces in compile-time for AccountGetWebAuthorizationsRequest.
var (
	_ bin.Encoder     = &AccountGetWebAuthorizationsRequest{}
	_ bin.Decoder     = &AccountGetWebAuthorizationsRequest{}
	_ bin.BareEncoder = &AccountGetWebAuthorizationsRequest{}
	_ bin.BareDecoder = &AccountGetWebAuthorizationsRequest{}
)

func (g *AccountGetWebAuthorizationsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetWebAuthorizationsRequest) String() string {
	if g == nil {
		return "AccountGetWebAuthorizationsRequest(nil)"
	}
	type Alias AccountGetWebAuthorizationsRequest
	return fmt.Sprintf("AccountGetWebAuthorizationsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetWebAuthorizationsRequest) TypeID() uint32 {
	return AccountGetWebAuthorizationsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetWebAuthorizationsRequest) TypeName() string {
	return "account.getWebAuthorizations"
}

// TypeInfo returns info about TL type.
func (g *AccountGetWebAuthorizationsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getWebAuthorizations",
		ID:   AccountGetWebAuthorizationsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetWebAuthorizationsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getWebAuthorizations#182e6d6f as nil")
	}
	b.PutID(AccountGetWebAuthorizationsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetWebAuthorizationsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getWebAuthorizations#182e6d6f as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetWebAuthorizationsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getWebAuthorizations#182e6d6f to nil")
	}
	if err := b.ConsumeID(AccountGetWebAuthorizationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getWebAuthorizations#182e6d6f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetWebAuthorizationsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getWebAuthorizations#182e6d6f to nil")
	}
	return nil
}

// AccountGetWebAuthorizations invokes method account.getWebAuthorizations#182e6d6f returning error if any.
// Get web login widget¹ authorizations
//
// Links:
//  1. https://core.telegram.org/widgets/login
//
// See https://core.telegram.org/method/account.getWebAuthorizations for reference.
func (c *Client) AccountGetWebAuthorizations(ctx context.Context) (*AccountWebAuthorizations, error) {
	var result AccountWebAuthorizations

	request := &AccountGetWebAuthorizationsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
