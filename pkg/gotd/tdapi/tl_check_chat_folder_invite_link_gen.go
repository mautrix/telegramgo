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

// CheckChatFolderInviteLinkRequest represents TL type `checkChatFolderInviteLink#1f25999b`.
type CheckChatFolderInviteLinkRequest struct {
	// Invite link to be checked
	InviteLink string
}

// CheckChatFolderInviteLinkRequestTypeID is TL type id of CheckChatFolderInviteLinkRequest.
const CheckChatFolderInviteLinkRequestTypeID = 0x1f25999b

// Ensuring interfaces in compile-time for CheckChatFolderInviteLinkRequest.
var (
	_ bin.Encoder     = &CheckChatFolderInviteLinkRequest{}
	_ bin.Decoder     = &CheckChatFolderInviteLinkRequest{}
	_ bin.BareEncoder = &CheckChatFolderInviteLinkRequest{}
	_ bin.BareDecoder = &CheckChatFolderInviteLinkRequest{}
)

func (c *CheckChatFolderInviteLinkRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.InviteLink == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckChatFolderInviteLinkRequest) String() string {
	if c == nil {
		return "CheckChatFolderInviteLinkRequest(nil)"
	}
	type Alias CheckChatFolderInviteLinkRequest
	return fmt.Sprintf("CheckChatFolderInviteLinkRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckChatFolderInviteLinkRequest) TypeID() uint32 {
	return CheckChatFolderInviteLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckChatFolderInviteLinkRequest) TypeName() string {
	return "checkChatFolderInviteLink"
}

// TypeInfo returns info about TL type.
func (c *CheckChatFolderInviteLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkChatFolderInviteLink",
		ID:   CheckChatFolderInviteLinkRequestTypeID,
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckChatFolderInviteLinkRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatFolderInviteLink#1f25999b as nil")
	}
	b.PutID(CheckChatFolderInviteLinkRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckChatFolderInviteLinkRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatFolderInviteLink#1f25999b as nil")
	}
	b.PutString(c.InviteLink)
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckChatFolderInviteLinkRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatFolderInviteLink#1f25999b to nil")
	}
	if err := b.ConsumeID(CheckChatFolderInviteLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode checkChatFolderInviteLink#1f25999b: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckChatFolderInviteLinkRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatFolderInviteLink#1f25999b to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode checkChatFolderInviteLink#1f25999b: field invite_link: %w", err)
		}
		c.InviteLink = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckChatFolderInviteLinkRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatFolderInviteLink#1f25999b as nil")
	}
	b.ObjStart()
	b.PutID("checkChatFolderInviteLink")
	b.Comma()
	b.FieldStart("invite_link")
	b.PutString(c.InviteLink)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckChatFolderInviteLinkRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatFolderInviteLink#1f25999b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkChatFolderInviteLink"); err != nil {
				return fmt.Errorf("unable to decode checkChatFolderInviteLink#1f25999b: %w", err)
			}
		case "invite_link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode checkChatFolderInviteLink#1f25999b: field invite_link: %w", err)
			}
			c.InviteLink = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInviteLink returns value of InviteLink field.
func (c *CheckChatFolderInviteLinkRequest) GetInviteLink() (value string) {
	if c == nil {
		return
	}
	return c.InviteLink
}

// CheckChatFolderInviteLink invokes method checkChatFolderInviteLink#1f25999b returning error if any.
func (c *Client) CheckChatFolderInviteLink(ctx context.Context, invitelink string) (*ChatFolderInviteLinkInfo, error) {
	var result ChatFolderInviteLinkInfo

	request := &CheckChatFolderInviteLinkRequest{
		InviteLink: invitelink,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
