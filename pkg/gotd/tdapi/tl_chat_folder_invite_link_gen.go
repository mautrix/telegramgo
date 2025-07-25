// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// ChatFolderInviteLink represents TL type `chatFolderInviteLink#d33caf97`.
type ChatFolderInviteLink struct {
	// The chat folder invite link
	InviteLink string
	// Name of the link
	Name string
	// Identifiers of chats, included in the link
	ChatIDs []int64
}

// ChatFolderInviteLinkTypeID is TL type id of ChatFolderInviteLink.
const ChatFolderInviteLinkTypeID = 0xd33caf97

// Ensuring interfaces in compile-time for ChatFolderInviteLink.
var (
	_ bin.Encoder     = &ChatFolderInviteLink{}
	_ bin.Decoder     = &ChatFolderInviteLink{}
	_ bin.BareEncoder = &ChatFolderInviteLink{}
	_ bin.BareDecoder = &ChatFolderInviteLink{}
)

func (c *ChatFolderInviteLink) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.InviteLink == "") {
		return false
	}
	if !(c.Name == "") {
		return false
	}
	if !(c.ChatIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatFolderInviteLink) String() string {
	if c == nil {
		return "ChatFolderInviteLink(nil)"
	}
	type Alias ChatFolderInviteLink
	return fmt.Sprintf("ChatFolderInviteLink%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatFolderInviteLink) TypeID() uint32 {
	return ChatFolderInviteLinkTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatFolderInviteLink) TypeName() string {
	return "chatFolderInviteLink"
}

// TypeInfo returns info about TL type.
func (c *ChatFolderInviteLink) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatFolderInviteLink",
		ID:   ChatFolderInviteLinkTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InviteLink",
			SchemaName: "invite_link",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "ChatIDs",
			SchemaName: "chat_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatFolderInviteLink) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatFolderInviteLink#d33caf97 as nil")
	}
	b.PutID(ChatFolderInviteLinkTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatFolderInviteLink) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatFolderInviteLink#d33caf97 as nil")
	}
	b.PutString(c.InviteLink)
	b.PutString(c.Name)
	b.PutInt(len(c.ChatIDs))
	for _, v := range c.ChatIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatFolderInviteLink) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatFolderInviteLink#d33caf97 to nil")
	}
	if err := b.ConsumeID(ChatFolderInviteLinkTypeID); err != nil {
		return fmt.Errorf("unable to decode chatFolderInviteLink#d33caf97: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatFolderInviteLink) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatFolderInviteLink#d33caf97 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatFolderInviteLink#d33caf97: field invite_link: %w", err)
		}
		c.InviteLink = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatFolderInviteLink#d33caf97: field name: %w", err)
		}
		c.Name = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatFolderInviteLink#d33caf97: field chat_ids: %w", err)
		}

		if headerLen > 0 {
			c.ChatIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chatFolderInviteLink#d33caf97: field chat_ids: %w", err)
			}
			c.ChatIDs = append(c.ChatIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatFolderInviteLink) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatFolderInviteLink#d33caf97 as nil")
	}
	b.ObjStart()
	b.PutID("chatFolderInviteLink")
	b.Comma()
	b.FieldStart("invite_link")
	b.PutString(c.InviteLink)
	b.Comma()
	b.FieldStart("name")
	b.PutString(c.Name)
	b.Comma()
	b.FieldStart("chat_ids")
	b.ArrStart()
	for _, v := range c.ChatIDs {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatFolderInviteLink) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatFolderInviteLink#d33caf97 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatFolderInviteLink"); err != nil {
				return fmt.Errorf("unable to decode chatFolderInviteLink#d33caf97: %w", err)
			}
		case "invite_link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatFolderInviteLink#d33caf97: field invite_link: %w", err)
			}
			c.InviteLink = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatFolderInviteLink#d33caf97: field name: %w", err)
			}
			c.Name = value
		case "chat_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode chatFolderInviteLink#d33caf97: field chat_ids: %w", err)
				}
				c.ChatIDs = append(c.ChatIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode chatFolderInviteLink#d33caf97: field chat_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInviteLink returns value of InviteLink field.
func (c *ChatFolderInviteLink) GetInviteLink() (value string) {
	if c == nil {
		return
	}
	return c.InviteLink
}

// GetName returns value of Name field.
func (c *ChatFolderInviteLink) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

// GetChatIDs returns value of ChatIDs field.
func (c *ChatFolderInviteLink) GetChatIDs() (value []int64) {
	if c == nil {
		return
	}
	return c.ChatIDs
}
