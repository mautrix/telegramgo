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

// QuickReplyShortcut represents TL type `quickReplyShortcut#bdfd9a95`.
type QuickReplyShortcut struct {
	// Unique shortcut identifier
	ID int32
	// The name of the shortcut that can be used to use the shortcut
	Name string
	// The first shortcut message
	FirstMessage QuickReplyMessage
	// The total number of messages in the shortcut
	MessageCount int32
}

// QuickReplyShortcutTypeID is TL type id of QuickReplyShortcut.
const QuickReplyShortcutTypeID = 0xbdfd9a95

// Ensuring interfaces in compile-time for QuickReplyShortcut.
var (
	_ bin.Encoder     = &QuickReplyShortcut{}
	_ bin.Decoder     = &QuickReplyShortcut{}
	_ bin.BareEncoder = &QuickReplyShortcut{}
	_ bin.BareDecoder = &QuickReplyShortcut{}
)

func (q *QuickReplyShortcut) Zero() bool {
	if q == nil {
		return true
	}
	if !(q.ID == 0) {
		return false
	}
	if !(q.Name == "") {
		return false
	}
	if !(q.FirstMessage.Zero()) {
		return false
	}
	if !(q.MessageCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (q *QuickReplyShortcut) String() string {
	if q == nil {
		return "QuickReplyShortcut(nil)"
	}
	type Alias QuickReplyShortcut
	return fmt.Sprintf("QuickReplyShortcut%+v", Alias(*q))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*QuickReplyShortcut) TypeID() uint32 {
	return QuickReplyShortcutTypeID
}

// TypeName returns name of type in TL schema.
func (*QuickReplyShortcut) TypeName() string {
	return "quickReplyShortcut"
}

// TypeInfo returns info about TL type.
func (q *QuickReplyShortcut) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "quickReplyShortcut",
		ID:   QuickReplyShortcutTypeID,
	}
	if q == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "FirstMessage",
			SchemaName: "first_message",
		},
		{
			Name:       "MessageCount",
			SchemaName: "message_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (q *QuickReplyShortcut) Encode(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't encode quickReplyShortcut#bdfd9a95 as nil")
	}
	b.PutID(QuickReplyShortcutTypeID)
	return q.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (q *QuickReplyShortcut) EncodeBare(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't encode quickReplyShortcut#bdfd9a95 as nil")
	}
	b.PutInt32(q.ID)
	b.PutString(q.Name)
	if err := q.FirstMessage.Encode(b); err != nil {
		return fmt.Errorf("unable to encode quickReplyShortcut#bdfd9a95: field first_message: %w", err)
	}
	b.PutInt32(q.MessageCount)
	return nil
}

// Decode implements bin.Decoder.
func (q *QuickReplyShortcut) Decode(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't decode quickReplyShortcut#bdfd9a95 to nil")
	}
	if err := b.ConsumeID(QuickReplyShortcutTypeID); err != nil {
		return fmt.Errorf("unable to decode quickReplyShortcut#bdfd9a95: %w", err)
	}
	return q.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (q *QuickReplyShortcut) DecodeBare(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't decode quickReplyShortcut#bdfd9a95 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode quickReplyShortcut#bdfd9a95: field id: %w", err)
		}
		q.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode quickReplyShortcut#bdfd9a95: field name: %w", err)
		}
		q.Name = value
	}
	{
		if err := q.FirstMessage.Decode(b); err != nil {
			return fmt.Errorf("unable to decode quickReplyShortcut#bdfd9a95: field first_message: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode quickReplyShortcut#bdfd9a95: field message_count: %w", err)
		}
		q.MessageCount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (q *QuickReplyShortcut) EncodeTDLibJSON(b tdjson.Encoder) error {
	if q == nil {
		return fmt.Errorf("can't encode quickReplyShortcut#bdfd9a95 as nil")
	}
	b.ObjStart()
	b.PutID("quickReplyShortcut")
	b.Comma()
	b.FieldStart("id")
	b.PutInt32(q.ID)
	b.Comma()
	b.FieldStart("name")
	b.PutString(q.Name)
	b.Comma()
	b.FieldStart("first_message")
	if err := q.FirstMessage.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode quickReplyShortcut#bdfd9a95: field first_message: %w", err)
	}
	b.Comma()
	b.FieldStart("message_count")
	b.PutInt32(q.MessageCount)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (q *QuickReplyShortcut) DecodeTDLibJSON(b tdjson.Decoder) error {
	if q == nil {
		return fmt.Errorf("can't decode quickReplyShortcut#bdfd9a95 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("quickReplyShortcut"); err != nil {
				return fmt.Errorf("unable to decode quickReplyShortcut#bdfd9a95: %w", err)
			}
		case "id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode quickReplyShortcut#bdfd9a95: field id: %w", err)
			}
			q.ID = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode quickReplyShortcut#bdfd9a95: field name: %w", err)
			}
			q.Name = value
		case "first_message":
			if err := q.FirstMessage.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode quickReplyShortcut#bdfd9a95: field first_message: %w", err)
			}
		case "message_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode quickReplyShortcut#bdfd9a95: field message_count: %w", err)
			}
			q.MessageCount = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (q *QuickReplyShortcut) GetID() (value int32) {
	if q == nil {
		return
	}
	return q.ID
}

// GetName returns value of Name field.
func (q *QuickReplyShortcut) GetName() (value string) {
	if q == nil {
		return
	}
	return q.Name
}

// GetFirstMessage returns value of FirstMessage field.
func (q *QuickReplyShortcut) GetFirstMessage() (value QuickReplyMessage) {
	if q == nil {
		return
	}
	return q.FirstMessage
}

// GetMessageCount returns value of MessageCount field.
func (q *QuickReplyShortcut) GetMessageCount() (value int32) {
	if q == nil {
		return
	}
	return q.MessageCount
}
