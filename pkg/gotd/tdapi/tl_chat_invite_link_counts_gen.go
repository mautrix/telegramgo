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

// ChatInviteLinkCounts represents TL type `chatInviteLinkCounts#c953d7f8`.
type ChatInviteLinkCounts struct {
	// List of invite link counts
	InviteLinkCounts []ChatInviteLinkCount
}

// ChatInviteLinkCountsTypeID is TL type id of ChatInviteLinkCounts.
const ChatInviteLinkCountsTypeID = 0xc953d7f8

// Ensuring interfaces in compile-time for ChatInviteLinkCounts.
var (
	_ bin.Encoder     = &ChatInviteLinkCounts{}
	_ bin.Decoder     = &ChatInviteLinkCounts{}
	_ bin.BareEncoder = &ChatInviteLinkCounts{}
	_ bin.BareDecoder = &ChatInviteLinkCounts{}
)

func (c *ChatInviteLinkCounts) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.InviteLinkCounts == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInviteLinkCounts) String() string {
	if c == nil {
		return "ChatInviteLinkCounts(nil)"
	}
	type Alias ChatInviteLinkCounts
	return fmt.Sprintf("ChatInviteLinkCounts%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInviteLinkCounts) TypeID() uint32 {
	return ChatInviteLinkCountsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInviteLinkCounts) TypeName() string {
	return "chatInviteLinkCounts"
}

// TypeInfo returns info about TL type.
func (c *ChatInviteLinkCounts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInviteLinkCounts",
		ID:   ChatInviteLinkCountsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InviteLinkCounts",
			SchemaName: "invite_link_counts",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatInviteLinkCounts) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinkCounts#c953d7f8 as nil")
	}
	b.PutID(ChatInviteLinkCountsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatInviteLinkCounts) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinkCounts#c953d7f8 as nil")
	}
	b.PutInt(len(c.InviteLinkCounts))
	for idx, v := range c.InviteLinkCounts {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare chatInviteLinkCounts#c953d7f8: field invite_link_counts element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInviteLinkCounts) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinkCounts#c953d7f8 to nil")
	}
	if err := b.ConsumeID(ChatInviteLinkCountsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteLinkCounts#c953d7f8: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatInviteLinkCounts) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinkCounts#c953d7f8 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkCounts#c953d7f8: field invite_link_counts: %w", err)
		}

		if headerLen > 0 {
			c.InviteLinkCounts = make([]ChatInviteLinkCount, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatInviteLinkCount
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare chatInviteLinkCounts#c953d7f8: field invite_link_counts: %w", err)
			}
			c.InviteLinkCounts = append(c.InviteLinkCounts, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatInviteLinkCounts) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinkCounts#c953d7f8 as nil")
	}
	b.ObjStart()
	b.PutID("chatInviteLinkCounts")
	b.Comma()
	b.FieldStart("invite_link_counts")
	b.ArrStart()
	for idx, v := range c.InviteLinkCounts {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode chatInviteLinkCounts#c953d7f8: field invite_link_counts element with index %d: %w", idx, err)
		}
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
func (c *ChatInviteLinkCounts) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinkCounts#c953d7f8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatInviteLinkCounts"); err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkCounts#c953d7f8: %w", err)
			}
		case "invite_link_counts":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value ChatInviteLinkCount
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode chatInviteLinkCounts#c953d7f8: field invite_link_counts: %w", err)
				}
				c.InviteLinkCounts = append(c.InviteLinkCounts, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkCounts#c953d7f8: field invite_link_counts: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInviteLinkCounts returns value of InviteLinkCounts field.
func (c *ChatInviteLinkCounts) GetInviteLinkCounts() (value []ChatInviteLinkCount) {
	if c == nil {
		return
	}
	return c.InviteLinkCounts
}
