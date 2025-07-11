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

// ContactsImportContactsRequest represents TL type `contacts.importContacts#2c800be5`.
// Imports contacts: saves a full list on the server, adds already registered contacts to
// the contact list, returns added contacts and their info.
// Use contacts.addContact¹ to add Telegram contacts without actually using their phone
// number.
//
// Links:
//  1. https://core.telegram.org/method/contacts.addContact
//
// See https://core.telegram.org/method/contacts.importContacts for reference.
type ContactsImportContactsRequest struct {
	// List of contacts to import
	Contacts []InputPhoneContact
}

// ContactsImportContactsRequestTypeID is TL type id of ContactsImportContactsRequest.
const ContactsImportContactsRequestTypeID = 0x2c800be5

// Ensuring interfaces in compile-time for ContactsImportContactsRequest.
var (
	_ bin.Encoder     = &ContactsImportContactsRequest{}
	_ bin.Decoder     = &ContactsImportContactsRequest{}
	_ bin.BareEncoder = &ContactsImportContactsRequest{}
	_ bin.BareDecoder = &ContactsImportContactsRequest{}
)

func (i *ContactsImportContactsRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Contacts == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *ContactsImportContactsRequest) String() string {
	if i == nil {
		return "ContactsImportContactsRequest(nil)"
	}
	type Alias ContactsImportContactsRequest
	return fmt.Sprintf("ContactsImportContactsRequest%+v", Alias(*i))
}

// FillFrom fills ContactsImportContactsRequest from given interface.
func (i *ContactsImportContactsRequest) FillFrom(from interface {
	GetContacts() (value []InputPhoneContact)
}) {
	i.Contacts = from.GetContacts()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsImportContactsRequest) TypeID() uint32 {
	return ContactsImportContactsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsImportContactsRequest) TypeName() string {
	return "contacts.importContacts"
}

// TypeInfo returns info about TL type.
func (i *ContactsImportContactsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.importContacts",
		ID:   ContactsImportContactsRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Contacts",
			SchemaName: "contacts",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *ContactsImportContactsRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode contacts.importContacts#2c800be5 as nil")
	}
	b.PutID(ContactsImportContactsRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *ContactsImportContactsRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode contacts.importContacts#2c800be5 as nil")
	}
	b.PutVectorHeader(len(i.Contacts))
	for idx, v := range i.Contacts {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.importContacts#2c800be5: field contacts element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *ContactsImportContactsRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode contacts.importContacts#2c800be5 to nil")
	}
	if err := b.ConsumeID(ContactsImportContactsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.importContacts#2c800be5: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *ContactsImportContactsRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode contacts.importContacts#2c800be5 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.importContacts#2c800be5: field contacts: %w", err)
		}

		if headerLen > 0 {
			i.Contacts = make([]InputPhoneContact, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value InputPhoneContact
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode contacts.importContacts#2c800be5: field contacts: %w", err)
			}
			i.Contacts = append(i.Contacts, value)
		}
	}
	return nil
}

// GetContacts returns value of Contacts field.
func (i *ContactsImportContactsRequest) GetContacts() (value []InputPhoneContact) {
	if i == nil {
		return
	}
	return i.Contacts
}

// ContactsImportContacts invokes method contacts.importContacts#2c800be5 returning error if any.
// Imports contacts: saves a full list on the server, adds already registered contacts to
// the contact list, returns added contacts and their info.
// Use contacts.addContact¹ to add Telegram contacts without actually using their phone
// number.
//
// Links:
//  1. https://core.telegram.org/method/contacts.addContact
//
// See https://core.telegram.org/method/contacts.importContacts for reference.
func (c *Client) ContactsImportContacts(ctx context.Context, contacts []InputPhoneContact) (*ContactsImportedContacts, error) {
	var result ContactsImportedContacts

	request := &ContactsImportContactsRequest{
		Contacts: contacts,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
