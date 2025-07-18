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

// AccountGetContentSettingsRequest represents TL type `account.getContentSettings#8b9b4dae`.
// Get sensitive content settings
//
// See https://core.telegram.org/method/account.getContentSettings for reference.
type AccountGetContentSettingsRequest struct {
}

// AccountGetContentSettingsRequestTypeID is TL type id of AccountGetContentSettingsRequest.
const AccountGetContentSettingsRequestTypeID = 0x8b9b4dae

// Ensuring interfaces in compile-time for AccountGetContentSettingsRequest.
var (
	_ bin.Encoder     = &AccountGetContentSettingsRequest{}
	_ bin.Decoder     = &AccountGetContentSettingsRequest{}
	_ bin.BareEncoder = &AccountGetContentSettingsRequest{}
	_ bin.BareDecoder = &AccountGetContentSettingsRequest{}
)

func (g *AccountGetContentSettingsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetContentSettingsRequest) String() string {
	if g == nil {
		return "AccountGetContentSettingsRequest(nil)"
	}
	type Alias AccountGetContentSettingsRequest
	return fmt.Sprintf("AccountGetContentSettingsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetContentSettingsRequest) TypeID() uint32 {
	return AccountGetContentSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetContentSettingsRequest) TypeName() string {
	return "account.getContentSettings"
}

// TypeInfo returns info about TL type.
func (g *AccountGetContentSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getContentSettings",
		ID:   AccountGetContentSettingsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetContentSettingsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getContentSettings#8b9b4dae as nil")
	}
	b.PutID(AccountGetContentSettingsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetContentSettingsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getContentSettings#8b9b4dae as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetContentSettingsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getContentSettings#8b9b4dae to nil")
	}
	if err := b.ConsumeID(AccountGetContentSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getContentSettings#8b9b4dae: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetContentSettingsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getContentSettings#8b9b4dae to nil")
	}
	return nil
}

// AccountGetContentSettings invokes method account.getContentSettings#8b9b4dae returning error if any.
// Get sensitive content settings
//
// See https://core.telegram.org/method/account.getContentSettings for reference.
func (c *Client) AccountGetContentSettings(ctx context.Context) (*AccountContentSettings, error) {
	var result AccountContentSettings

	request := &AccountGetContentSettingsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
