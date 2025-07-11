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

// ToggleChatGiftNotificationsRequest represents TL type `toggleChatGiftNotifications#84a7045e`.
type ToggleChatGiftNotificationsRequest struct {
	// Identifier of the channel chat
	ChatID int64
	// Pass true to enable notifications about new gifts owned by the channel chat; pass
	// false to disable the notifications
	AreEnabled bool
}

// ToggleChatGiftNotificationsRequestTypeID is TL type id of ToggleChatGiftNotificationsRequest.
const ToggleChatGiftNotificationsRequestTypeID = 0x84a7045e

// Ensuring interfaces in compile-time for ToggleChatGiftNotificationsRequest.
var (
	_ bin.Encoder     = &ToggleChatGiftNotificationsRequest{}
	_ bin.Decoder     = &ToggleChatGiftNotificationsRequest{}
	_ bin.BareEncoder = &ToggleChatGiftNotificationsRequest{}
	_ bin.BareDecoder = &ToggleChatGiftNotificationsRequest{}
)

func (t *ToggleChatGiftNotificationsRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.ChatID == 0) {
		return false
	}
	if !(t.AreEnabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleChatGiftNotificationsRequest) String() string {
	if t == nil {
		return "ToggleChatGiftNotificationsRequest(nil)"
	}
	type Alias ToggleChatGiftNotificationsRequest
	return fmt.Sprintf("ToggleChatGiftNotificationsRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleChatGiftNotificationsRequest) TypeID() uint32 {
	return ToggleChatGiftNotificationsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleChatGiftNotificationsRequest) TypeName() string {
	return "toggleChatGiftNotifications"
}

// TypeInfo returns info about TL type.
func (t *ToggleChatGiftNotificationsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleChatGiftNotifications",
		ID:   ToggleChatGiftNotificationsRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "AreEnabled",
			SchemaName: "are_enabled",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleChatGiftNotificationsRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatGiftNotifications#84a7045e as nil")
	}
	b.PutID(ToggleChatGiftNotificationsRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleChatGiftNotificationsRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatGiftNotifications#84a7045e as nil")
	}
	b.PutInt53(t.ChatID)
	b.PutBool(t.AreEnabled)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleChatGiftNotificationsRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatGiftNotifications#84a7045e to nil")
	}
	if err := b.ConsumeID(ToggleChatGiftNotificationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleChatGiftNotifications#84a7045e: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleChatGiftNotificationsRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatGiftNotifications#84a7045e to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode toggleChatGiftNotifications#84a7045e: field chat_id: %w", err)
		}
		t.ChatID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleChatGiftNotifications#84a7045e: field are_enabled: %w", err)
		}
		t.AreEnabled = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleChatGiftNotificationsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatGiftNotifications#84a7045e as nil")
	}
	b.ObjStart()
	b.PutID("toggleChatGiftNotifications")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(t.ChatID)
	b.Comma()
	b.FieldStart("are_enabled")
	b.PutBool(t.AreEnabled)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleChatGiftNotificationsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatGiftNotifications#84a7045e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleChatGiftNotifications"); err != nil {
				return fmt.Errorf("unable to decode toggleChatGiftNotifications#84a7045e: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode toggleChatGiftNotifications#84a7045e: field chat_id: %w", err)
			}
			t.ChatID = value
		case "are_enabled":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleChatGiftNotifications#84a7045e: field are_enabled: %w", err)
			}
			t.AreEnabled = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (t *ToggleChatGiftNotificationsRequest) GetChatID() (value int64) {
	if t == nil {
		return
	}
	return t.ChatID
}

// GetAreEnabled returns value of AreEnabled field.
func (t *ToggleChatGiftNotificationsRequest) GetAreEnabled() (value bool) {
	if t == nil {
		return
	}
	return t.AreEnabled
}

// ToggleChatGiftNotifications invokes method toggleChatGiftNotifications#84a7045e returning error if any.
func (c *Client) ToggleChatGiftNotifications(ctx context.Context, request *ToggleChatGiftNotificationsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
