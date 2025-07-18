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

// MessageCalendarDay represents TL type `messageCalendarDay#e98f8f62`.
type MessageCalendarDay struct {
	// Total number of found messages sent on the day
	TotalCount int32
	// First message sent on the day
	Message Message
}

// MessageCalendarDayTypeID is TL type id of MessageCalendarDay.
const MessageCalendarDayTypeID = 0xe98f8f62

// Ensuring interfaces in compile-time for MessageCalendarDay.
var (
	_ bin.Encoder     = &MessageCalendarDay{}
	_ bin.Decoder     = &MessageCalendarDay{}
	_ bin.BareEncoder = &MessageCalendarDay{}
	_ bin.BareDecoder = &MessageCalendarDay{}
)

func (m *MessageCalendarDay) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.TotalCount == 0) {
		return false
	}
	if !(m.Message.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageCalendarDay) String() string {
	if m == nil {
		return "MessageCalendarDay(nil)"
	}
	type Alias MessageCalendarDay
	return fmt.Sprintf("MessageCalendarDay%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageCalendarDay) TypeID() uint32 {
	return MessageCalendarDayTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageCalendarDay) TypeName() string {
	return "messageCalendarDay"
}

// TypeInfo returns info about TL type.
func (m *MessageCalendarDay) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageCalendarDay",
		ID:   MessageCalendarDayTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageCalendarDay) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageCalendarDay#e98f8f62 as nil")
	}
	b.PutID(MessageCalendarDayTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageCalendarDay) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageCalendarDay#e98f8f62 as nil")
	}
	b.PutInt32(m.TotalCount)
	if err := m.Message.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageCalendarDay#e98f8f62: field message: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageCalendarDay) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageCalendarDay#e98f8f62 to nil")
	}
	if err := b.ConsumeID(MessageCalendarDayTypeID); err != nil {
		return fmt.Errorf("unable to decode messageCalendarDay#e98f8f62: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageCalendarDay) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageCalendarDay#e98f8f62 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messageCalendarDay#e98f8f62: field total_count: %w", err)
		}
		m.TotalCount = value
	}
	{
		if err := m.Message.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageCalendarDay#e98f8f62: field message: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageCalendarDay) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageCalendarDay#e98f8f62 as nil")
	}
	b.ObjStart()
	b.PutID("messageCalendarDay")
	b.Comma()
	b.FieldStart("total_count")
	b.PutInt32(m.TotalCount)
	b.Comma()
	b.FieldStart("message")
	if err := m.Message.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode messageCalendarDay#e98f8f62: field message: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageCalendarDay) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageCalendarDay#e98f8f62 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageCalendarDay"); err != nil {
				return fmt.Errorf("unable to decode messageCalendarDay#e98f8f62: %w", err)
			}
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messageCalendarDay#e98f8f62: field total_count: %w", err)
			}
			m.TotalCount = value
		case "message":
			if err := m.Message.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode messageCalendarDay#e98f8f62: field message: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTotalCount returns value of TotalCount field.
func (m *MessageCalendarDay) GetTotalCount() (value int32) {
	if m == nil {
		return
	}
	return m.TotalCount
}

// GetMessage returns value of Message field.
func (m *MessageCalendarDay) GetMessage() (value Message) {
	if m == nil {
		return
	}
	return m.Message
}
