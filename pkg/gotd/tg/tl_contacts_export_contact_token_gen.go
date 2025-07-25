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

// ContactsExportContactTokenRequest represents TL type `contacts.exportContactToken#f8654027`.
// Generates a temporary profile link¹ for the currently logged-in user.
//
// Links:
//  1. https://core.telegram.org/api/links#temporary-profile-links
//
// See https://core.telegram.org/method/contacts.exportContactToken for reference.
type ContactsExportContactTokenRequest struct {
}

// ContactsExportContactTokenRequestTypeID is TL type id of ContactsExportContactTokenRequest.
const ContactsExportContactTokenRequestTypeID = 0xf8654027

// Ensuring interfaces in compile-time for ContactsExportContactTokenRequest.
var (
	_ bin.Encoder     = &ContactsExportContactTokenRequest{}
	_ bin.Decoder     = &ContactsExportContactTokenRequest{}
	_ bin.BareEncoder = &ContactsExportContactTokenRequest{}
	_ bin.BareDecoder = &ContactsExportContactTokenRequest{}
)

func (e *ContactsExportContactTokenRequest) Zero() bool {
	if e == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (e *ContactsExportContactTokenRequest) String() string {
	if e == nil {
		return "ContactsExportContactTokenRequest(nil)"
	}
	type Alias ContactsExportContactTokenRequest
	return fmt.Sprintf("ContactsExportContactTokenRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsExportContactTokenRequest) TypeID() uint32 {
	return ContactsExportContactTokenRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsExportContactTokenRequest) TypeName() string {
	return "contacts.exportContactToken"
}

// TypeInfo returns info about TL type.
func (e *ContactsExportContactTokenRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.exportContactToken",
		ID:   ContactsExportContactTokenRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (e *ContactsExportContactTokenRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode contacts.exportContactToken#f8654027 as nil")
	}
	b.PutID(ContactsExportContactTokenRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *ContactsExportContactTokenRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode contacts.exportContactToken#f8654027 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *ContactsExportContactTokenRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode contacts.exportContactToken#f8654027 to nil")
	}
	if err := b.ConsumeID(ContactsExportContactTokenRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.exportContactToken#f8654027: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *ContactsExportContactTokenRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode contacts.exportContactToken#f8654027 to nil")
	}
	return nil
}

// ContactsExportContactToken invokes method contacts.exportContactToken#f8654027 returning error if any.
// Generates a temporary profile link¹ for the currently logged-in user.
//
// Links:
//  1. https://core.telegram.org/api/links#temporary-profile-links
//
// See https://core.telegram.org/method/contacts.exportContactToken for reference.
func (c *Client) ContactsExportContactToken(ctx context.Context) (*ExportedContactToken, error) {
	var result ExportedContactToken

	request := &ContactsExportContactTokenRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
