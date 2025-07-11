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

// AccountConnectedBots represents TL type `account.connectedBots#17d7f87b`.
// Info about currently connected business bots¹.
//
// Links:
//  1. https://core.telegram.org/api/business#connected-bots
//
// See https://core.telegram.org/constructor/account.connectedBots for reference.
type AccountConnectedBots struct {
	// Info about the connected bots
	ConnectedBots []ConnectedBot
	// Bot information
	Users []UserClass
}

// AccountConnectedBotsTypeID is TL type id of AccountConnectedBots.
const AccountConnectedBotsTypeID = 0x17d7f87b

// Ensuring interfaces in compile-time for AccountConnectedBots.
var (
	_ bin.Encoder     = &AccountConnectedBots{}
	_ bin.Decoder     = &AccountConnectedBots{}
	_ bin.BareEncoder = &AccountConnectedBots{}
	_ bin.BareDecoder = &AccountConnectedBots{}
)

func (c *AccountConnectedBots) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ConnectedBots == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AccountConnectedBots) String() string {
	if c == nil {
		return "AccountConnectedBots(nil)"
	}
	type Alias AccountConnectedBots
	return fmt.Sprintf("AccountConnectedBots%+v", Alias(*c))
}

// FillFrom fills AccountConnectedBots from given interface.
func (c *AccountConnectedBots) FillFrom(from interface {
	GetConnectedBots() (value []ConnectedBot)
	GetUsers() (value []UserClass)
}) {
	c.ConnectedBots = from.GetConnectedBots()
	c.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountConnectedBots) TypeID() uint32 {
	return AccountConnectedBotsTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountConnectedBots) TypeName() string {
	return "account.connectedBots"
}

// TypeInfo returns info about TL type.
func (c *AccountConnectedBots) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.connectedBots",
		ID:   AccountConnectedBotsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ConnectedBots",
			SchemaName: "connected_bots",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *AccountConnectedBots) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.connectedBots#17d7f87b as nil")
	}
	b.PutID(AccountConnectedBotsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *AccountConnectedBots) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.connectedBots#17d7f87b as nil")
	}
	b.PutVectorHeader(len(c.ConnectedBots))
	for idx, v := range c.ConnectedBots {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.connectedBots#17d7f87b: field connected_bots element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode account.connectedBots#17d7f87b: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.connectedBots#17d7f87b: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *AccountConnectedBots) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.connectedBots#17d7f87b to nil")
	}
	if err := b.ConsumeID(AccountConnectedBotsTypeID); err != nil {
		return fmt.Errorf("unable to decode account.connectedBots#17d7f87b: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *AccountConnectedBots) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.connectedBots#17d7f87b to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.connectedBots#17d7f87b: field connected_bots: %w", err)
		}

		if headerLen > 0 {
			c.ConnectedBots = make([]ConnectedBot, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ConnectedBot
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode account.connectedBots#17d7f87b: field connected_bots: %w", err)
			}
			c.ConnectedBots = append(c.ConnectedBots, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.connectedBots#17d7f87b: field users: %w", err)
		}

		if headerLen > 0 {
			c.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.connectedBots#17d7f87b: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// GetConnectedBots returns value of ConnectedBots field.
func (c *AccountConnectedBots) GetConnectedBots() (value []ConnectedBot) {
	if c == nil {
		return
	}
	return c.ConnectedBots
}

// GetUsers returns value of Users field.
func (c *AccountConnectedBots) GetUsers() (value []UserClass) {
	if c == nil {
		return
	}
	return c.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (c *AccountConnectedBots) MapUsers() (value UserClassArray) {
	return UserClassArray(c.Users)
}
