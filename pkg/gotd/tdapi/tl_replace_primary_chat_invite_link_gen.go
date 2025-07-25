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

// ReplacePrimaryChatInviteLinkRequest represents TL type `replacePrimaryChatInviteLink#3f9e7b9d`.
type ReplacePrimaryChatInviteLinkRequest struct {
	// Chat identifier
	ChatID int64
}

// ReplacePrimaryChatInviteLinkRequestTypeID is TL type id of ReplacePrimaryChatInviteLinkRequest.
const ReplacePrimaryChatInviteLinkRequestTypeID = 0x3f9e7b9d

// Ensuring interfaces in compile-time for ReplacePrimaryChatInviteLinkRequest.
var (
	_ bin.Encoder     = &ReplacePrimaryChatInviteLinkRequest{}
	_ bin.Decoder     = &ReplacePrimaryChatInviteLinkRequest{}
	_ bin.BareEncoder = &ReplacePrimaryChatInviteLinkRequest{}
	_ bin.BareDecoder = &ReplacePrimaryChatInviteLinkRequest{}
)

func (r *ReplacePrimaryChatInviteLinkRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplacePrimaryChatInviteLinkRequest) String() string {
	if r == nil {
		return "ReplacePrimaryChatInviteLinkRequest(nil)"
	}
	type Alias ReplacePrimaryChatInviteLinkRequest
	return fmt.Sprintf("ReplacePrimaryChatInviteLinkRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplacePrimaryChatInviteLinkRequest) TypeID() uint32 {
	return ReplacePrimaryChatInviteLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplacePrimaryChatInviteLinkRequest) TypeName() string {
	return "replacePrimaryChatInviteLink"
}

// TypeInfo returns info about TL type.
func (r *ReplacePrimaryChatInviteLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replacePrimaryChatInviteLink",
		ID:   ReplacePrimaryChatInviteLinkRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReplacePrimaryChatInviteLinkRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replacePrimaryChatInviteLink#3f9e7b9d as nil")
	}
	b.PutID(ReplacePrimaryChatInviteLinkRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReplacePrimaryChatInviteLinkRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replacePrimaryChatInviteLink#3f9e7b9d as nil")
	}
	b.PutInt53(r.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReplacePrimaryChatInviteLinkRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replacePrimaryChatInviteLink#3f9e7b9d to nil")
	}
	if err := b.ConsumeID(ReplacePrimaryChatInviteLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode replacePrimaryChatInviteLink#3f9e7b9d: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReplacePrimaryChatInviteLinkRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replacePrimaryChatInviteLink#3f9e7b9d to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode replacePrimaryChatInviteLink#3f9e7b9d: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReplacePrimaryChatInviteLinkRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode replacePrimaryChatInviteLink#3f9e7b9d as nil")
	}
	b.ObjStart()
	b.PutID("replacePrimaryChatInviteLink")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(r.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReplacePrimaryChatInviteLinkRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode replacePrimaryChatInviteLink#3f9e7b9d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("replacePrimaryChatInviteLink"); err != nil {
				return fmt.Errorf("unable to decode replacePrimaryChatInviteLink#3f9e7b9d: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode replacePrimaryChatInviteLink#3f9e7b9d: field chat_id: %w", err)
			}
			r.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (r *ReplacePrimaryChatInviteLinkRequest) GetChatID() (value int64) {
	if r == nil {
		return
	}
	return r.ChatID
}

// ReplacePrimaryChatInviteLink invokes method replacePrimaryChatInviteLink#3f9e7b9d returning error if any.
func (c *Client) ReplacePrimaryChatInviteLink(ctx context.Context, chatid int64) (*ChatInviteLink, error) {
	var result ChatInviteLink

	request := &ReplacePrimaryChatInviteLinkRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
