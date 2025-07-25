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

// MessageAutoDeleteTime represents TL type `messageAutoDeleteTime#758b0715`.
type MessageAutoDeleteTime struct {
	// Message auto-delete time, in seconds. If 0, then messages aren't deleted automatically
	Time int32
}

// MessageAutoDeleteTimeTypeID is TL type id of MessageAutoDeleteTime.
const MessageAutoDeleteTimeTypeID = 0x758b0715

// Ensuring interfaces in compile-time for MessageAutoDeleteTime.
var (
	_ bin.Encoder     = &MessageAutoDeleteTime{}
	_ bin.Decoder     = &MessageAutoDeleteTime{}
	_ bin.BareEncoder = &MessageAutoDeleteTime{}
	_ bin.BareDecoder = &MessageAutoDeleteTime{}
)

func (m *MessageAutoDeleteTime) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Time == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageAutoDeleteTime) String() string {
	if m == nil {
		return "MessageAutoDeleteTime(nil)"
	}
	type Alias MessageAutoDeleteTime
	return fmt.Sprintf("MessageAutoDeleteTime%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageAutoDeleteTime) TypeID() uint32 {
	return MessageAutoDeleteTimeTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageAutoDeleteTime) TypeName() string {
	return "messageAutoDeleteTime"
}

// TypeInfo returns info about TL type.
func (m *MessageAutoDeleteTime) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageAutoDeleteTime",
		ID:   MessageAutoDeleteTimeTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Time",
			SchemaName: "time",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageAutoDeleteTime) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageAutoDeleteTime#758b0715 as nil")
	}
	b.PutID(MessageAutoDeleteTimeTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageAutoDeleteTime) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageAutoDeleteTime#758b0715 as nil")
	}
	b.PutInt32(m.Time)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageAutoDeleteTime) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageAutoDeleteTime#758b0715 to nil")
	}
	if err := b.ConsumeID(MessageAutoDeleteTimeTypeID); err != nil {
		return fmt.Errorf("unable to decode messageAutoDeleteTime#758b0715: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageAutoDeleteTime) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageAutoDeleteTime#758b0715 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messageAutoDeleteTime#758b0715: field time: %w", err)
		}
		m.Time = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageAutoDeleteTime) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageAutoDeleteTime#758b0715 as nil")
	}
	b.ObjStart()
	b.PutID("messageAutoDeleteTime")
	b.Comma()
	b.FieldStart("time")
	b.PutInt32(m.Time)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageAutoDeleteTime) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageAutoDeleteTime#758b0715 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageAutoDeleteTime"); err != nil {
				return fmt.Errorf("unable to decode messageAutoDeleteTime#758b0715: %w", err)
			}
		case "time":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messageAutoDeleteTime#758b0715: field time: %w", err)
			}
			m.Time = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTime returns value of Time field.
func (m *MessageAutoDeleteTime) GetTime() (value int32) {
	if m == nil {
		return
	}
	return m.Time
}
