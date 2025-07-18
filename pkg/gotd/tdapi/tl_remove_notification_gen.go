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

// RemoveNotificationRequest represents TL type `removeNotification#336ab34e`.
type RemoveNotificationRequest struct {
	// Identifier of notification group to which the notification belongs
	NotificationGroupID int32
	// Identifier of removed notification
	NotificationID int32
}

// RemoveNotificationRequestTypeID is TL type id of RemoveNotificationRequest.
const RemoveNotificationRequestTypeID = 0x336ab34e

// Ensuring interfaces in compile-time for RemoveNotificationRequest.
var (
	_ bin.Encoder     = &RemoveNotificationRequest{}
	_ bin.Decoder     = &RemoveNotificationRequest{}
	_ bin.BareEncoder = &RemoveNotificationRequest{}
	_ bin.BareDecoder = &RemoveNotificationRequest{}
)

func (r *RemoveNotificationRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.NotificationGroupID == 0) {
		return false
	}
	if !(r.NotificationID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RemoveNotificationRequest) String() string {
	if r == nil {
		return "RemoveNotificationRequest(nil)"
	}
	type Alias RemoveNotificationRequest
	return fmt.Sprintf("RemoveNotificationRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RemoveNotificationRequest) TypeID() uint32 {
	return RemoveNotificationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RemoveNotificationRequest) TypeName() string {
	return "removeNotification"
}

// TypeInfo returns info about TL type.
func (r *RemoveNotificationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "removeNotification",
		ID:   RemoveNotificationRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "NotificationGroupID",
			SchemaName: "notification_group_id",
		},
		{
			Name:       "NotificationID",
			SchemaName: "notification_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RemoveNotificationRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeNotification#336ab34e as nil")
	}
	b.PutID(RemoveNotificationRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RemoveNotificationRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeNotification#336ab34e as nil")
	}
	b.PutInt32(r.NotificationGroupID)
	b.PutInt32(r.NotificationID)
	return nil
}

// Decode implements bin.Decoder.
func (r *RemoveNotificationRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeNotification#336ab34e to nil")
	}
	if err := b.ConsumeID(RemoveNotificationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode removeNotification#336ab34e: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RemoveNotificationRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeNotification#336ab34e to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode removeNotification#336ab34e: field notification_group_id: %w", err)
		}
		r.NotificationGroupID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode removeNotification#336ab34e: field notification_id: %w", err)
		}
		r.NotificationID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RemoveNotificationRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode removeNotification#336ab34e as nil")
	}
	b.ObjStart()
	b.PutID("removeNotification")
	b.Comma()
	b.FieldStart("notification_group_id")
	b.PutInt32(r.NotificationGroupID)
	b.Comma()
	b.FieldStart("notification_id")
	b.PutInt32(r.NotificationID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RemoveNotificationRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode removeNotification#336ab34e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("removeNotification"); err != nil {
				return fmt.Errorf("unable to decode removeNotification#336ab34e: %w", err)
			}
		case "notification_group_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode removeNotification#336ab34e: field notification_group_id: %w", err)
			}
			r.NotificationGroupID = value
		case "notification_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode removeNotification#336ab34e: field notification_id: %w", err)
			}
			r.NotificationID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetNotificationGroupID returns value of NotificationGroupID field.
func (r *RemoveNotificationRequest) GetNotificationGroupID() (value int32) {
	if r == nil {
		return
	}
	return r.NotificationGroupID
}

// GetNotificationID returns value of NotificationID field.
func (r *RemoveNotificationRequest) GetNotificationID() (value int32) {
	if r == nil {
		return
	}
	return r.NotificationID
}

// RemoveNotification invokes method removeNotification#336ab34e returning error if any.
func (c *Client) RemoveNotification(ctx context.Context, request *RemoveNotificationRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
