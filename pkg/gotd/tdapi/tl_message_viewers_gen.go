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

// MessageViewers represents TL type `messageViewers#87c73f6c`.
type MessageViewers struct {
	// List of message viewers
	Viewers []MessageViewer
}

// MessageViewersTypeID is TL type id of MessageViewers.
const MessageViewersTypeID = 0x87c73f6c

// Ensuring interfaces in compile-time for MessageViewers.
var (
	_ bin.Encoder     = &MessageViewers{}
	_ bin.Decoder     = &MessageViewers{}
	_ bin.BareEncoder = &MessageViewers{}
	_ bin.BareDecoder = &MessageViewers{}
)

func (m *MessageViewers) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Viewers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageViewers) String() string {
	if m == nil {
		return "MessageViewers(nil)"
	}
	type Alias MessageViewers
	return fmt.Sprintf("MessageViewers%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageViewers) TypeID() uint32 {
	return MessageViewersTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageViewers) TypeName() string {
	return "messageViewers"
}

// TypeInfo returns info about TL type.
func (m *MessageViewers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageViewers",
		ID:   MessageViewersTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Viewers",
			SchemaName: "viewers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageViewers) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageViewers#87c73f6c as nil")
	}
	b.PutID(MessageViewersTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageViewers) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageViewers#87c73f6c as nil")
	}
	b.PutInt(len(m.Viewers))
	for idx, v := range m.Viewers {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare messageViewers#87c73f6c: field viewers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageViewers) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageViewers#87c73f6c to nil")
	}
	if err := b.ConsumeID(MessageViewersTypeID); err != nil {
		return fmt.Errorf("unable to decode messageViewers#87c73f6c: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageViewers) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageViewers#87c73f6c to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageViewers#87c73f6c: field viewers: %w", err)
		}

		if headerLen > 0 {
			m.Viewers = make([]MessageViewer, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MessageViewer
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare messageViewers#87c73f6c: field viewers: %w", err)
			}
			m.Viewers = append(m.Viewers, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageViewers) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageViewers#87c73f6c as nil")
	}
	b.ObjStart()
	b.PutID("messageViewers")
	b.Comma()
	b.FieldStart("viewers")
	b.ArrStart()
	for idx, v := range m.Viewers {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode messageViewers#87c73f6c: field viewers element with index %d: %w", idx, err)
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
func (m *MessageViewers) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageViewers#87c73f6c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageViewers"); err != nil {
				return fmt.Errorf("unable to decode messageViewers#87c73f6c: %w", err)
			}
		case "viewers":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value MessageViewer
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode messageViewers#87c73f6c: field viewers: %w", err)
				}
				m.Viewers = append(m.Viewers, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode messageViewers#87c73f6c: field viewers: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetViewers returns value of Viewers field.
func (m *MessageViewers) GetViewers() (value []MessageViewer) {
	if m == nil {
		return
	}
	return m.Viewers
}
