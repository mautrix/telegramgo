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

// AccountGetGlobalPrivacySettingsRequest represents TL type `account.getGlobalPrivacySettings#eb2b4cf6`.
// Get global privacy settings
//
// See https://core.telegram.org/method/account.getGlobalPrivacySettings for reference.
type AccountGetGlobalPrivacySettingsRequest struct {
}

// AccountGetGlobalPrivacySettingsRequestTypeID is TL type id of AccountGetGlobalPrivacySettingsRequest.
const AccountGetGlobalPrivacySettingsRequestTypeID = 0xeb2b4cf6

// Ensuring interfaces in compile-time for AccountGetGlobalPrivacySettingsRequest.
var (
	_ bin.Encoder     = &AccountGetGlobalPrivacySettingsRequest{}
	_ bin.Decoder     = &AccountGetGlobalPrivacySettingsRequest{}
	_ bin.BareEncoder = &AccountGetGlobalPrivacySettingsRequest{}
	_ bin.BareDecoder = &AccountGetGlobalPrivacySettingsRequest{}
)

func (g *AccountGetGlobalPrivacySettingsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetGlobalPrivacySettingsRequest) String() string {
	if g == nil {
		return "AccountGetGlobalPrivacySettingsRequest(nil)"
	}
	type Alias AccountGetGlobalPrivacySettingsRequest
	return fmt.Sprintf("AccountGetGlobalPrivacySettingsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetGlobalPrivacySettingsRequest) TypeID() uint32 {
	return AccountGetGlobalPrivacySettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetGlobalPrivacySettingsRequest) TypeName() string {
	return "account.getGlobalPrivacySettings"
}

// TypeInfo returns info about TL type.
func (g *AccountGetGlobalPrivacySettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getGlobalPrivacySettings",
		ID:   AccountGetGlobalPrivacySettingsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetGlobalPrivacySettingsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getGlobalPrivacySettings#eb2b4cf6 as nil")
	}
	b.PutID(AccountGetGlobalPrivacySettingsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetGlobalPrivacySettingsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getGlobalPrivacySettings#eb2b4cf6 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetGlobalPrivacySettingsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getGlobalPrivacySettings#eb2b4cf6 to nil")
	}
	if err := b.ConsumeID(AccountGetGlobalPrivacySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getGlobalPrivacySettings#eb2b4cf6: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetGlobalPrivacySettingsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getGlobalPrivacySettings#eb2b4cf6 to nil")
	}
	return nil
}

// AccountGetGlobalPrivacySettings invokes method account.getGlobalPrivacySettings#eb2b4cf6 returning error if any.
// Get global privacy settings
//
// See https://core.telegram.org/method/account.getGlobalPrivacySettings for reference.
func (c *Client) AccountGetGlobalPrivacySettings(ctx context.Context) (*GlobalPrivacySettings, error) {
	var result GlobalPrivacySettings

	request := &AccountGetGlobalPrivacySettingsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
