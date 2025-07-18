// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// MsgContainer represents TL type `msg_container#73f1f8dc`.
type MsgContainer struct {
	// Messages field of MsgContainer.
	Messages []Message
}

// MsgContainerTypeID is TL type id of MsgContainer.
const MsgContainerTypeID = 0x73f1f8dc

// Ensuring interfaces in compile-time for MsgContainer.
var (
	_ bin.Encoder     = &MsgContainer{}
	_ bin.Decoder     = &MsgContainer{}
	_ bin.BareEncoder = &MsgContainer{}
	_ bin.BareDecoder = &MsgContainer{}
)

func (m *MsgContainer) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Messages == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MsgContainer) String() string {
	if m == nil {
		return "MsgContainer(nil)"
	}
	type Alias MsgContainer
	return fmt.Sprintf("MsgContainer%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MsgContainer) TypeID() uint32 {
	return MsgContainerTypeID
}

// TypeName returns name of type in TL schema.
func (*MsgContainer) TypeName() string {
	return "msg_container"
}

// TypeInfo returns info about TL type.
func (m *MsgContainer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "msg_container",
		ID:   MsgContainerTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MsgContainer) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msg_container#73f1f8dc as nil")
	}
	b.PutID(MsgContainerTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MsgContainer) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msg_container#73f1f8dc as nil")
	}
	b.PutInt(len(m.Messages))
	for idx, v := range m.Messages {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare msg_container#73f1f8dc: field messages element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MsgContainer) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msg_container#73f1f8dc to nil")
	}
	if err := b.ConsumeID(MsgContainerTypeID); err != nil {
		return fmt.Errorf("unable to decode msg_container#73f1f8dc: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MsgContainer) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msg_container#73f1f8dc to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode msg_container#73f1f8dc: field messages: %w", err)
		}

		if headerLen > 0 {
			m.Messages = make([]Message, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Message
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare msg_container#73f1f8dc: field messages: %w", err)
			}
			m.Messages = append(m.Messages, value)
		}
	}
	return nil
}

// GetMessages returns value of Messages field.
func (m *MsgContainer) GetMessages() (value []Message) {
	if m == nil {
		return
	}
	return m.Messages
}
