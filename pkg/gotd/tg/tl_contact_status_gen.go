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

// ContactStatus represents TL type `contactStatus#16d9703b`.
// Contact status: online / offline.
//
// See https://core.telegram.org/constructor/contactStatus for reference.
type ContactStatus struct {
	// User identifier
	UserID int64
	// Online status
	Status UserStatusClass
}

// ContactStatusTypeID is TL type id of ContactStatus.
const ContactStatusTypeID = 0x16d9703b

// Ensuring interfaces in compile-time for ContactStatus.
var (
	_ bin.Encoder     = &ContactStatus{}
	_ bin.Decoder     = &ContactStatus{}
	_ bin.BareEncoder = &ContactStatus{}
	_ bin.BareDecoder = &ContactStatus{}
)

func (c *ContactStatus) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.Status == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ContactStatus) String() string {
	if c == nil {
		return "ContactStatus(nil)"
	}
	type Alias ContactStatus
	return fmt.Sprintf("ContactStatus%+v", Alias(*c))
}

// FillFrom fills ContactStatus from given interface.
func (c *ContactStatus) FillFrom(from interface {
	GetUserID() (value int64)
	GetStatus() (value UserStatusClass)
}) {
	c.UserID = from.GetUserID()
	c.Status = from.GetStatus()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactStatus) TypeID() uint32 {
	return ContactStatusTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactStatus) TypeName() string {
	return "contactStatus"
}

// TypeInfo returns info about TL type.
func (c *ContactStatus) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contactStatus",
		ID:   ContactStatusTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Status",
			SchemaName: "status",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ContactStatus) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode contactStatus#16d9703b as nil")
	}
	b.PutID(ContactStatusTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ContactStatus) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode contactStatus#16d9703b as nil")
	}
	b.PutLong(c.UserID)
	if c.Status == nil {
		return fmt.Errorf("unable to encode contactStatus#16d9703b: field status is nil")
	}
	if err := c.Status.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contactStatus#16d9703b: field status: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ContactStatus) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode contactStatus#16d9703b to nil")
	}
	if err := b.ConsumeID(ContactStatusTypeID); err != nil {
		return fmt.Errorf("unable to decode contactStatus#16d9703b: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ContactStatus) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode contactStatus#16d9703b to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode contactStatus#16d9703b: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := DecodeUserStatus(b)
		if err != nil {
			return fmt.Errorf("unable to decode contactStatus#16d9703b: field status: %w", err)
		}
		c.Status = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (c *ContactStatus) GetUserID() (value int64) {
	if c == nil {
		return
	}
	return c.UserID
}

// GetStatus returns value of Status field.
func (c *ContactStatus) GetStatus() (value UserStatusClass) {
	if c == nil {
		return
	}
	return c.Status
}
