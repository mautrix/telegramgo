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

// DeleteAllRevokedChatInviteLinksRequest represents TL type `deleteAllRevokedChatInviteLinks#424816da`.
type DeleteAllRevokedChatInviteLinksRequest struct {
	// Chat identifier
	ChatID int64
	// User identifier of a chat administrator, which links will be deleted. Must be an
	// identifier of the current user for non-owner
	CreatorUserID int64
}

// DeleteAllRevokedChatInviteLinksRequestTypeID is TL type id of DeleteAllRevokedChatInviteLinksRequest.
const DeleteAllRevokedChatInviteLinksRequestTypeID = 0x424816da

// Ensuring interfaces in compile-time for DeleteAllRevokedChatInviteLinksRequest.
var (
	_ bin.Encoder     = &DeleteAllRevokedChatInviteLinksRequest{}
	_ bin.Decoder     = &DeleteAllRevokedChatInviteLinksRequest{}
	_ bin.BareEncoder = &DeleteAllRevokedChatInviteLinksRequest{}
	_ bin.BareDecoder = &DeleteAllRevokedChatInviteLinksRequest{}
)

func (d *DeleteAllRevokedChatInviteLinksRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ChatID == 0) {
		return false
	}
	if !(d.CreatorUserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteAllRevokedChatInviteLinksRequest) String() string {
	if d == nil {
		return "DeleteAllRevokedChatInviteLinksRequest(nil)"
	}
	type Alias DeleteAllRevokedChatInviteLinksRequest
	return fmt.Sprintf("DeleteAllRevokedChatInviteLinksRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteAllRevokedChatInviteLinksRequest) TypeID() uint32 {
	return DeleteAllRevokedChatInviteLinksRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteAllRevokedChatInviteLinksRequest) TypeName() string {
	return "deleteAllRevokedChatInviteLinks"
}

// TypeInfo returns info about TL type.
func (d *DeleteAllRevokedChatInviteLinksRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteAllRevokedChatInviteLinks",
		ID:   DeleteAllRevokedChatInviteLinksRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "CreatorUserID",
			SchemaName: "creator_user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteAllRevokedChatInviteLinksRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteAllRevokedChatInviteLinks#424816da as nil")
	}
	b.PutID(DeleteAllRevokedChatInviteLinksRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteAllRevokedChatInviteLinksRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteAllRevokedChatInviteLinks#424816da as nil")
	}
	b.PutInt53(d.ChatID)
	b.PutInt53(d.CreatorUserID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteAllRevokedChatInviteLinksRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteAllRevokedChatInviteLinks#424816da to nil")
	}
	if err := b.ConsumeID(DeleteAllRevokedChatInviteLinksRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteAllRevokedChatInviteLinks#424816da: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteAllRevokedChatInviteLinksRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteAllRevokedChatInviteLinks#424816da to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode deleteAllRevokedChatInviteLinks#424816da: field chat_id: %w", err)
		}
		d.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode deleteAllRevokedChatInviteLinks#424816da: field creator_user_id: %w", err)
		}
		d.CreatorUserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteAllRevokedChatInviteLinksRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteAllRevokedChatInviteLinks#424816da as nil")
	}
	b.ObjStart()
	b.PutID("deleteAllRevokedChatInviteLinks")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(d.ChatID)
	b.Comma()
	b.FieldStart("creator_user_id")
	b.PutInt53(d.CreatorUserID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteAllRevokedChatInviteLinksRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteAllRevokedChatInviteLinks#424816da to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteAllRevokedChatInviteLinks"); err != nil {
				return fmt.Errorf("unable to decode deleteAllRevokedChatInviteLinks#424816da: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode deleteAllRevokedChatInviteLinks#424816da: field chat_id: %w", err)
			}
			d.ChatID = value
		case "creator_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode deleteAllRevokedChatInviteLinks#424816da: field creator_user_id: %w", err)
			}
			d.CreatorUserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (d *DeleteAllRevokedChatInviteLinksRequest) GetChatID() (value int64) {
	if d == nil {
		return
	}
	return d.ChatID
}

// GetCreatorUserID returns value of CreatorUserID field.
func (d *DeleteAllRevokedChatInviteLinksRequest) GetCreatorUserID() (value int64) {
	if d == nil {
		return
	}
	return d.CreatorUserID
}

// DeleteAllRevokedChatInviteLinks invokes method deleteAllRevokedChatInviteLinks#424816da returning error if any.
func (c *Client) DeleteAllRevokedChatInviteLinks(ctx context.Context, request *DeleteAllRevokedChatInviteLinksRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
