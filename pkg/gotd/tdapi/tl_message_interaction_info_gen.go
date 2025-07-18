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

// MessageInteractionInfo represents TL type `messageInteractionInfo#2bbcde05`.
type MessageInteractionInfo struct {
	// Number of times the message was viewed
	ViewCount int32
	// Number of times the message was forwarded
	ForwardCount int32
	// Information about direct or indirect replies to the message; may be null. Currently,
	// available only in channels with a discussion supergroup and discussion supergroups for
	// messages, which are not replies itself
	ReplyInfo MessageReplyInfo
	// The list of reactions or tags added to the message; may be null
	Reactions MessageReactions
}

// MessageInteractionInfoTypeID is TL type id of MessageInteractionInfo.
const MessageInteractionInfoTypeID = 0x2bbcde05

// Ensuring interfaces in compile-time for MessageInteractionInfo.
var (
	_ bin.Encoder     = &MessageInteractionInfo{}
	_ bin.Decoder     = &MessageInteractionInfo{}
	_ bin.BareEncoder = &MessageInteractionInfo{}
	_ bin.BareDecoder = &MessageInteractionInfo{}
)

func (m *MessageInteractionInfo) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.ViewCount == 0) {
		return false
	}
	if !(m.ForwardCount == 0) {
		return false
	}
	if !(m.ReplyInfo.Zero()) {
		return false
	}
	if !(m.Reactions.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageInteractionInfo) String() string {
	if m == nil {
		return "MessageInteractionInfo(nil)"
	}
	type Alias MessageInteractionInfo
	return fmt.Sprintf("MessageInteractionInfo%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageInteractionInfo) TypeID() uint32 {
	return MessageInteractionInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageInteractionInfo) TypeName() string {
	return "messageInteractionInfo"
}

// TypeInfo returns info about TL type.
func (m *MessageInteractionInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageInteractionInfo",
		ID:   MessageInteractionInfoTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ViewCount",
			SchemaName: "view_count",
		},
		{
			Name:       "ForwardCount",
			SchemaName: "forward_count",
		},
		{
			Name:       "ReplyInfo",
			SchemaName: "reply_info",
		},
		{
			Name:       "Reactions",
			SchemaName: "reactions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageInteractionInfo) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageInteractionInfo#2bbcde05 as nil")
	}
	b.PutID(MessageInteractionInfoTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageInteractionInfo) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageInteractionInfo#2bbcde05 as nil")
	}
	b.PutInt32(m.ViewCount)
	b.PutInt32(m.ForwardCount)
	if err := m.ReplyInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageInteractionInfo#2bbcde05: field reply_info: %w", err)
	}
	if err := m.Reactions.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageInteractionInfo#2bbcde05: field reactions: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageInteractionInfo) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageInteractionInfo#2bbcde05 to nil")
	}
	if err := b.ConsumeID(MessageInteractionInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode messageInteractionInfo#2bbcde05: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageInteractionInfo) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageInteractionInfo#2bbcde05 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messageInteractionInfo#2bbcde05: field view_count: %w", err)
		}
		m.ViewCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messageInteractionInfo#2bbcde05: field forward_count: %w", err)
		}
		m.ForwardCount = value
	}
	{
		if err := m.ReplyInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageInteractionInfo#2bbcde05: field reply_info: %w", err)
		}
	}
	{
		if err := m.Reactions.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageInteractionInfo#2bbcde05: field reactions: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageInteractionInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageInteractionInfo#2bbcde05 as nil")
	}
	b.ObjStart()
	b.PutID("messageInteractionInfo")
	b.Comma()
	b.FieldStart("view_count")
	b.PutInt32(m.ViewCount)
	b.Comma()
	b.FieldStart("forward_count")
	b.PutInt32(m.ForwardCount)
	b.Comma()
	b.FieldStart("reply_info")
	if err := m.ReplyInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode messageInteractionInfo#2bbcde05: field reply_info: %w", err)
	}
	b.Comma()
	b.FieldStart("reactions")
	if err := m.Reactions.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode messageInteractionInfo#2bbcde05: field reactions: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageInteractionInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageInteractionInfo#2bbcde05 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageInteractionInfo"); err != nil {
				return fmt.Errorf("unable to decode messageInteractionInfo#2bbcde05: %w", err)
			}
		case "view_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messageInteractionInfo#2bbcde05: field view_count: %w", err)
			}
			m.ViewCount = value
		case "forward_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messageInteractionInfo#2bbcde05: field forward_count: %w", err)
			}
			m.ForwardCount = value
		case "reply_info":
			if err := m.ReplyInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode messageInteractionInfo#2bbcde05: field reply_info: %w", err)
			}
		case "reactions":
			if err := m.Reactions.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode messageInteractionInfo#2bbcde05: field reactions: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetViewCount returns value of ViewCount field.
func (m *MessageInteractionInfo) GetViewCount() (value int32) {
	if m == nil {
		return
	}
	return m.ViewCount
}

// GetForwardCount returns value of ForwardCount field.
func (m *MessageInteractionInfo) GetForwardCount() (value int32) {
	if m == nil {
		return
	}
	return m.ForwardCount
}

// GetReplyInfo returns value of ReplyInfo field.
func (m *MessageInteractionInfo) GetReplyInfo() (value MessageReplyInfo) {
	if m == nil {
		return
	}
	return m.ReplyInfo
}

// GetReactions returns value of Reactions field.
func (m *MessageInteractionInfo) GetReactions() (value MessageReactions) {
	if m == nil {
		return
	}
	return m.Reactions
}
