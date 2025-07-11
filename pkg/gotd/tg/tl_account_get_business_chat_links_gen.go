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

// AccountGetBusinessChatLinksRequest represents TL type `account.getBusinessChatLinks#6f70dde1`.
// List all created business chat deep links »¹.
//
// Links:
//  1. https://core.telegram.org/api/business#business-chat-links
//
// See https://core.telegram.org/method/account.getBusinessChatLinks for reference.
type AccountGetBusinessChatLinksRequest struct {
}

// AccountGetBusinessChatLinksRequestTypeID is TL type id of AccountGetBusinessChatLinksRequest.
const AccountGetBusinessChatLinksRequestTypeID = 0x6f70dde1

// Ensuring interfaces in compile-time for AccountGetBusinessChatLinksRequest.
var (
	_ bin.Encoder     = &AccountGetBusinessChatLinksRequest{}
	_ bin.Decoder     = &AccountGetBusinessChatLinksRequest{}
	_ bin.BareEncoder = &AccountGetBusinessChatLinksRequest{}
	_ bin.BareDecoder = &AccountGetBusinessChatLinksRequest{}
)

func (g *AccountGetBusinessChatLinksRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetBusinessChatLinksRequest) String() string {
	if g == nil {
		return "AccountGetBusinessChatLinksRequest(nil)"
	}
	type Alias AccountGetBusinessChatLinksRequest
	return fmt.Sprintf("AccountGetBusinessChatLinksRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetBusinessChatLinksRequest) TypeID() uint32 {
	return AccountGetBusinessChatLinksRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetBusinessChatLinksRequest) TypeName() string {
	return "account.getBusinessChatLinks"
}

// TypeInfo returns info about TL type.
func (g *AccountGetBusinessChatLinksRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getBusinessChatLinks",
		ID:   AccountGetBusinessChatLinksRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetBusinessChatLinksRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getBusinessChatLinks#6f70dde1 as nil")
	}
	b.PutID(AccountGetBusinessChatLinksRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetBusinessChatLinksRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getBusinessChatLinks#6f70dde1 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetBusinessChatLinksRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getBusinessChatLinks#6f70dde1 to nil")
	}
	if err := b.ConsumeID(AccountGetBusinessChatLinksRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getBusinessChatLinks#6f70dde1: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetBusinessChatLinksRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getBusinessChatLinks#6f70dde1 to nil")
	}
	return nil
}

// AccountGetBusinessChatLinks invokes method account.getBusinessChatLinks#6f70dde1 returning error if any.
// List all created business chat deep links »¹.
//
// Links:
//  1. https://core.telegram.org/api/business#business-chat-links
//
// See https://core.telegram.org/method/account.getBusinessChatLinks for reference.
func (c *Client) AccountGetBusinessChatLinks(ctx context.Context) (*AccountBusinessChatLinks, error) {
	var result AccountBusinessChatLinks

	request := &AccountGetBusinessChatLinksRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
