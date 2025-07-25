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

// MessageOriginUser represents TL type `messageOriginUser#9c009043`.
type MessageOriginUser struct {
	// Identifier of the user that originally sent the message
	SenderUserID int64
}

// MessageOriginUserTypeID is TL type id of MessageOriginUser.
const MessageOriginUserTypeID = 0x9c009043

// construct implements constructor of MessageOriginClass.
func (m MessageOriginUser) construct() MessageOriginClass { return &m }

// Ensuring interfaces in compile-time for MessageOriginUser.
var (
	_ bin.Encoder     = &MessageOriginUser{}
	_ bin.Decoder     = &MessageOriginUser{}
	_ bin.BareEncoder = &MessageOriginUser{}
	_ bin.BareDecoder = &MessageOriginUser{}

	_ MessageOriginClass = &MessageOriginUser{}
)

func (m *MessageOriginUser) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.SenderUserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageOriginUser) String() string {
	if m == nil {
		return "MessageOriginUser(nil)"
	}
	type Alias MessageOriginUser
	return fmt.Sprintf("MessageOriginUser%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageOriginUser) TypeID() uint32 {
	return MessageOriginUserTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageOriginUser) TypeName() string {
	return "messageOriginUser"
}

// TypeInfo returns info about TL type.
func (m *MessageOriginUser) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageOriginUser",
		ID:   MessageOriginUserTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SenderUserID",
			SchemaName: "sender_user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageOriginUser) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageOriginUser#9c009043 as nil")
	}
	b.PutID(MessageOriginUserTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageOriginUser) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageOriginUser#9c009043 as nil")
	}
	b.PutInt53(m.SenderUserID)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageOriginUser) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageOriginUser#9c009043 to nil")
	}
	if err := b.ConsumeID(MessageOriginUserTypeID); err != nil {
		return fmt.Errorf("unable to decode messageOriginUser#9c009043: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageOriginUser) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageOriginUser#9c009043 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode messageOriginUser#9c009043: field sender_user_id: %w", err)
		}
		m.SenderUserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageOriginUser) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageOriginUser#9c009043 as nil")
	}
	b.ObjStart()
	b.PutID("messageOriginUser")
	b.Comma()
	b.FieldStart("sender_user_id")
	b.PutInt53(m.SenderUserID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageOriginUser) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageOriginUser#9c009043 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageOriginUser"); err != nil {
				return fmt.Errorf("unable to decode messageOriginUser#9c009043: %w", err)
			}
		case "sender_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode messageOriginUser#9c009043: field sender_user_id: %w", err)
			}
			m.SenderUserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSenderUserID returns value of SenderUserID field.
func (m *MessageOriginUser) GetSenderUserID() (value int64) {
	if m == nil {
		return
	}
	return m.SenderUserID
}

// MessageOriginHiddenUser represents TL type `messageOriginHiddenUser#ed0c23da`.
type MessageOriginHiddenUser struct {
	// Name of the sender
	SenderName string
}

// MessageOriginHiddenUserTypeID is TL type id of MessageOriginHiddenUser.
const MessageOriginHiddenUserTypeID = 0xed0c23da

// construct implements constructor of MessageOriginClass.
func (m MessageOriginHiddenUser) construct() MessageOriginClass { return &m }

// Ensuring interfaces in compile-time for MessageOriginHiddenUser.
var (
	_ bin.Encoder     = &MessageOriginHiddenUser{}
	_ bin.Decoder     = &MessageOriginHiddenUser{}
	_ bin.BareEncoder = &MessageOriginHiddenUser{}
	_ bin.BareDecoder = &MessageOriginHiddenUser{}

	_ MessageOriginClass = &MessageOriginHiddenUser{}
)

func (m *MessageOriginHiddenUser) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.SenderName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageOriginHiddenUser) String() string {
	if m == nil {
		return "MessageOriginHiddenUser(nil)"
	}
	type Alias MessageOriginHiddenUser
	return fmt.Sprintf("MessageOriginHiddenUser%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageOriginHiddenUser) TypeID() uint32 {
	return MessageOriginHiddenUserTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageOriginHiddenUser) TypeName() string {
	return "messageOriginHiddenUser"
}

// TypeInfo returns info about TL type.
func (m *MessageOriginHiddenUser) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageOriginHiddenUser",
		ID:   MessageOriginHiddenUserTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SenderName",
			SchemaName: "sender_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageOriginHiddenUser) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageOriginHiddenUser#ed0c23da as nil")
	}
	b.PutID(MessageOriginHiddenUserTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageOriginHiddenUser) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageOriginHiddenUser#ed0c23da as nil")
	}
	b.PutString(m.SenderName)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageOriginHiddenUser) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageOriginHiddenUser#ed0c23da to nil")
	}
	if err := b.ConsumeID(MessageOriginHiddenUserTypeID); err != nil {
		return fmt.Errorf("unable to decode messageOriginHiddenUser#ed0c23da: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageOriginHiddenUser) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageOriginHiddenUser#ed0c23da to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageOriginHiddenUser#ed0c23da: field sender_name: %w", err)
		}
		m.SenderName = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageOriginHiddenUser) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageOriginHiddenUser#ed0c23da as nil")
	}
	b.ObjStart()
	b.PutID("messageOriginHiddenUser")
	b.Comma()
	b.FieldStart("sender_name")
	b.PutString(m.SenderName)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageOriginHiddenUser) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageOriginHiddenUser#ed0c23da to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageOriginHiddenUser"); err != nil {
				return fmt.Errorf("unable to decode messageOriginHiddenUser#ed0c23da: %w", err)
			}
		case "sender_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode messageOriginHiddenUser#ed0c23da: field sender_name: %w", err)
			}
			m.SenderName = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSenderName returns value of SenderName field.
func (m *MessageOriginHiddenUser) GetSenderName() (value string) {
	if m == nil {
		return
	}
	return m.SenderName
}

// MessageOriginChat represents TL type `messageOriginChat#f3bb5eb4`.
type MessageOriginChat struct {
	// Identifier of the chat that originally sent the message
	SenderChatID int64
	// For messages originally sent by an anonymous chat administrator, original message
	// author signature
	AuthorSignature string
}

// MessageOriginChatTypeID is TL type id of MessageOriginChat.
const MessageOriginChatTypeID = 0xf3bb5eb4

// construct implements constructor of MessageOriginClass.
func (m MessageOriginChat) construct() MessageOriginClass { return &m }

// Ensuring interfaces in compile-time for MessageOriginChat.
var (
	_ bin.Encoder     = &MessageOriginChat{}
	_ bin.Decoder     = &MessageOriginChat{}
	_ bin.BareEncoder = &MessageOriginChat{}
	_ bin.BareDecoder = &MessageOriginChat{}

	_ MessageOriginClass = &MessageOriginChat{}
)

func (m *MessageOriginChat) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.SenderChatID == 0) {
		return false
	}
	if !(m.AuthorSignature == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageOriginChat) String() string {
	if m == nil {
		return "MessageOriginChat(nil)"
	}
	type Alias MessageOriginChat
	return fmt.Sprintf("MessageOriginChat%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageOriginChat) TypeID() uint32 {
	return MessageOriginChatTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageOriginChat) TypeName() string {
	return "messageOriginChat"
}

// TypeInfo returns info about TL type.
func (m *MessageOriginChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageOriginChat",
		ID:   MessageOriginChatTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SenderChatID",
			SchemaName: "sender_chat_id",
		},
		{
			Name:       "AuthorSignature",
			SchemaName: "author_signature",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageOriginChat) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageOriginChat#f3bb5eb4 as nil")
	}
	b.PutID(MessageOriginChatTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageOriginChat) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageOriginChat#f3bb5eb4 as nil")
	}
	b.PutInt53(m.SenderChatID)
	b.PutString(m.AuthorSignature)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageOriginChat) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageOriginChat#f3bb5eb4 to nil")
	}
	if err := b.ConsumeID(MessageOriginChatTypeID); err != nil {
		return fmt.Errorf("unable to decode messageOriginChat#f3bb5eb4: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageOriginChat) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageOriginChat#f3bb5eb4 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode messageOriginChat#f3bb5eb4: field sender_chat_id: %w", err)
		}
		m.SenderChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageOriginChat#f3bb5eb4: field author_signature: %w", err)
		}
		m.AuthorSignature = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageOriginChat) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageOriginChat#f3bb5eb4 as nil")
	}
	b.ObjStart()
	b.PutID("messageOriginChat")
	b.Comma()
	b.FieldStart("sender_chat_id")
	b.PutInt53(m.SenderChatID)
	b.Comma()
	b.FieldStart("author_signature")
	b.PutString(m.AuthorSignature)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageOriginChat) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageOriginChat#f3bb5eb4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageOriginChat"); err != nil {
				return fmt.Errorf("unable to decode messageOriginChat#f3bb5eb4: %w", err)
			}
		case "sender_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode messageOriginChat#f3bb5eb4: field sender_chat_id: %w", err)
			}
			m.SenderChatID = value
		case "author_signature":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode messageOriginChat#f3bb5eb4: field author_signature: %w", err)
			}
			m.AuthorSignature = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSenderChatID returns value of SenderChatID field.
func (m *MessageOriginChat) GetSenderChatID() (value int64) {
	if m == nil {
		return
	}
	return m.SenderChatID
}

// GetAuthorSignature returns value of AuthorSignature field.
func (m *MessageOriginChat) GetAuthorSignature() (value string) {
	if m == nil {
		return
	}
	return m.AuthorSignature
}

// MessageOriginChannel represents TL type `messageOriginChannel#a97b51be`.
type MessageOriginChannel struct {
	// Identifier of the channel chat to which the message was originally sent
	ChatID int64
	// Message identifier of the original message
	MessageID int64
	// Original post author signature
	AuthorSignature string
}

// MessageOriginChannelTypeID is TL type id of MessageOriginChannel.
const MessageOriginChannelTypeID = 0xa97b51be

// construct implements constructor of MessageOriginClass.
func (m MessageOriginChannel) construct() MessageOriginClass { return &m }

// Ensuring interfaces in compile-time for MessageOriginChannel.
var (
	_ bin.Encoder     = &MessageOriginChannel{}
	_ bin.Decoder     = &MessageOriginChannel{}
	_ bin.BareEncoder = &MessageOriginChannel{}
	_ bin.BareDecoder = &MessageOriginChannel{}

	_ MessageOriginClass = &MessageOriginChannel{}
)

func (m *MessageOriginChannel) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.ChatID == 0) {
		return false
	}
	if !(m.MessageID == 0) {
		return false
	}
	if !(m.AuthorSignature == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageOriginChannel) String() string {
	if m == nil {
		return "MessageOriginChannel(nil)"
	}
	type Alias MessageOriginChannel
	return fmt.Sprintf("MessageOriginChannel%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageOriginChannel) TypeID() uint32 {
	return MessageOriginChannelTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageOriginChannel) TypeName() string {
	return "messageOriginChannel"
}

// TypeInfo returns info about TL type.
func (m *MessageOriginChannel) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageOriginChannel",
		ID:   MessageOriginChannelTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "AuthorSignature",
			SchemaName: "author_signature",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageOriginChannel) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageOriginChannel#a97b51be as nil")
	}
	b.PutID(MessageOriginChannelTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageOriginChannel) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageOriginChannel#a97b51be as nil")
	}
	b.PutInt53(m.ChatID)
	b.PutInt53(m.MessageID)
	b.PutString(m.AuthorSignature)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageOriginChannel) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageOriginChannel#a97b51be to nil")
	}
	if err := b.ConsumeID(MessageOriginChannelTypeID); err != nil {
		return fmt.Errorf("unable to decode messageOriginChannel#a97b51be: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageOriginChannel) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageOriginChannel#a97b51be to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode messageOriginChannel#a97b51be: field chat_id: %w", err)
		}
		m.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode messageOriginChannel#a97b51be: field message_id: %w", err)
		}
		m.MessageID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageOriginChannel#a97b51be: field author_signature: %w", err)
		}
		m.AuthorSignature = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageOriginChannel) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageOriginChannel#a97b51be as nil")
	}
	b.ObjStart()
	b.PutID("messageOriginChannel")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(m.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(m.MessageID)
	b.Comma()
	b.FieldStart("author_signature")
	b.PutString(m.AuthorSignature)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageOriginChannel) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageOriginChannel#a97b51be to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageOriginChannel"); err != nil {
				return fmt.Errorf("unable to decode messageOriginChannel#a97b51be: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode messageOriginChannel#a97b51be: field chat_id: %w", err)
			}
			m.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode messageOriginChannel#a97b51be: field message_id: %w", err)
			}
			m.MessageID = value
		case "author_signature":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode messageOriginChannel#a97b51be: field author_signature: %w", err)
			}
			m.AuthorSignature = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (m *MessageOriginChannel) GetChatID() (value int64) {
	if m == nil {
		return
	}
	return m.ChatID
}

// GetMessageID returns value of MessageID field.
func (m *MessageOriginChannel) GetMessageID() (value int64) {
	if m == nil {
		return
	}
	return m.MessageID
}

// GetAuthorSignature returns value of AuthorSignature field.
func (m *MessageOriginChannel) GetAuthorSignature() (value string) {
	if m == nil {
		return
	}
	return m.AuthorSignature
}

// MessageOriginClassName is schema name of MessageOriginClass.
const MessageOriginClassName = "MessageOrigin"

// MessageOriginClass represents MessageOrigin generic type.
//
// Example:
//
//	g, err := tdapi.DecodeMessageOrigin(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.MessageOriginUser: // messageOriginUser#9c009043
//	case *tdapi.MessageOriginHiddenUser: // messageOriginHiddenUser#ed0c23da
//	case *tdapi.MessageOriginChat: // messageOriginChat#f3bb5eb4
//	case *tdapi.MessageOriginChannel: // messageOriginChannel#a97b51be
//	default: panic(v)
//	}
type MessageOriginClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessageOriginClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeMessageOrigin implements binary de-serialization for MessageOriginClass.
func DecodeMessageOrigin(buf *bin.Buffer) (MessageOriginClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessageOriginUserTypeID:
		// Decoding messageOriginUser#9c009043.
		v := MessageOriginUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageOriginClass: %w", err)
		}
		return &v, nil
	case MessageOriginHiddenUserTypeID:
		// Decoding messageOriginHiddenUser#ed0c23da.
		v := MessageOriginHiddenUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageOriginClass: %w", err)
		}
		return &v, nil
	case MessageOriginChatTypeID:
		// Decoding messageOriginChat#f3bb5eb4.
		v := MessageOriginChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageOriginClass: %w", err)
		}
		return &v, nil
	case MessageOriginChannelTypeID:
		// Decoding messageOriginChannel#a97b51be.
		v := MessageOriginChannel{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageOriginClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageOriginClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONMessageOrigin implements binary de-serialization for MessageOriginClass.
func DecodeTDLibJSONMessageOrigin(buf tdjson.Decoder) (MessageOriginClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "messageOriginUser":
		// Decoding messageOriginUser#9c009043.
		v := MessageOriginUser{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageOriginClass: %w", err)
		}
		return &v, nil
	case "messageOriginHiddenUser":
		// Decoding messageOriginHiddenUser#ed0c23da.
		v := MessageOriginHiddenUser{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageOriginClass: %w", err)
		}
		return &v, nil
	case "messageOriginChat":
		// Decoding messageOriginChat#f3bb5eb4.
		v := MessageOriginChat{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageOriginClass: %w", err)
		}
		return &v, nil
	case "messageOriginChannel":
		// Decoding messageOriginChannel#a97b51be.
		v := MessageOriginChannel{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageOriginClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageOriginClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// MessageOrigin boxes the MessageOriginClass providing a helper.
type MessageOriginBox struct {
	MessageOrigin MessageOriginClass
}

// Decode implements bin.Decoder for MessageOriginBox.
func (b *MessageOriginBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageOriginBox to nil")
	}
	v, err := DecodeMessageOrigin(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageOrigin = v
	return nil
}

// Encode implements bin.Encode for MessageOriginBox.
func (b *MessageOriginBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MessageOrigin == nil {
		return fmt.Errorf("unable to encode MessageOriginClass as nil")
	}
	return b.MessageOrigin.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for MessageOriginBox.
func (b *MessageOriginBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageOriginBox to nil")
	}
	v, err := DecodeTDLibJSONMessageOrigin(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageOrigin = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for MessageOriginBox.
func (b *MessageOriginBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.MessageOrigin == nil {
		return fmt.Errorf("unable to encode MessageOriginClass as nil")
	}
	return b.MessageOrigin.EncodeTDLibJSON(buf)
}
