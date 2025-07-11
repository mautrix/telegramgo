// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagePeerVote represents TL type `messagePeerVote#b6cc2d5c`.
// How a peer voted in a poll
//
// See https://core.telegram.org/constructor/messagePeerVote for reference.
type MessagePeerVote struct {
	// Peer ID
	Peer PeerClass
	// The option chosen by the peer
	Option []byte
	// When did the peer cast the vote
	Date int
}

// MessagePeerVoteTypeID is TL type id of MessagePeerVote.
const MessagePeerVoteTypeID = 0xb6cc2d5c

// construct implements constructor of MessagePeerVoteClass.
func (m MessagePeerVote) construct() MessagePeerVoteClass { return &m }

// Ensuring interfaces in compile-time for MessagePeerVote.
var (
	_ bin.Encoder     = &MessagePeerVote{}
	_ bin.Decoder     = &MessagePeerVote{}
	_ bin.BareEncoder = &MessagePeerVote{}
	_ bin.BareDecoder = &MessagePeerVote{}

	_ MessagePeerVoteClass = &MessagePeerVote{}
)

func (m *MessagePeerVote) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Peer == nil) {
		return false
	}
	if !(m.Option == nil) {
		return false
	}
	if !(m.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagePeerVote) String() string {
	if m == nil {
		return "MessagePeerVote(nil)"
	}
	type Alias MessagePeerVote
	return fmt.Sprintf("MessagePeerVote%+v", Alias(*m))
}

// FillFrom fills MessagePeerVote from given interface.
func (m *MessagePeerVote) FillFrom(from interface {
	GetPeer() (value PeerClass)
	GetOption() (value []byte)
	GetDate() (value int)
}) {
	m.Peer = from.GetPeer()
	m.Option = from.GetOption()
	m.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagePeerVote) TypeID() uint32 {
	return MessagePeerVoteTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagePeerVote) TypeName() string {
	return "messagePeerVote"
}

// TypeInfo returns info about TL type.
func (m *MessagePeerVote) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messagePeerVote",
		ID:   MessagePeerVoteTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Option",
			SchemaName: "option",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessagePeerVote) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messagePeerVote#b6cc2d5c as nil")
	}
	b.PutID(MessagePeerVoteTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessagePeerVote) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messagePeerVote#b6cc2d5c as nil")
	}
	if m.Peer == nil {
		return fmt.Errorf("unable to encode messagePeerVote#b6cc2d5c: field peer is nil")
	}
	if err := m.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messagePeerVote#b6cc2d5c: field peer: %w", err)
	}
	b.PutBytes(m.Option)
	b.PutInt(m.Date)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessagePeerVote) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messagePeerVote#b6cc2d5c to nil")
	}
	if err := b.ConsumeID(MessagePeerVoteTypeID); err != nil {
		return fmt.Errorf("unable to decode messagePeerVote#b6cc2d5c: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessagePeerVote) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messagePeerVote#b6cc2d5c to nil")
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messagePeerVote#b6cc2d5c: field peer: %w", err)
		}
		m.Peer = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messagePeerVote#b6cc2d5c: field option: %w", err)
		}
		m.Option = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messagePeerVote#b6cc2d5c: field date: %w", err)
		}
		m.Date = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (m *MessagePeerVote) GetPeer() (value PeerClass) {
	if m == nil {
		return
	}
	return m.Peer
}

// GetOption returns value of Option field.
func (m *MessagePeerVote) GetOption() (value []byte) {
	if m == nil {
		return
	}
	return m.Option
}

// GetDate returns value of Date field.
func (m *MessagePeerVote) GetDate() (value int) {
	if m == nil {
		return
	}
	return m.Date
}

// MessagePeerVoteInputOption represents TL type `messagePeerVoteInputOption#74cda504`.
// How a peer voted in a poll (reduced constructor, returned if an option was provided to
// messages.getPollVotes¹)
//
// Links:
//  1. https://core.telegram.org/method/messages.getPollVotes
//
// See https://core.telegram.org/constructor/messagePeerVoteInputOption for reference.
type MessagePeerVoteInputOption struct {
	// The peer that voted for the queried option
	Peer PeerClass
	// When did the peer cast the vote
	Date int
}

// MessagePeerVoteInputOptionTypeID is TL type id of MessagePeerVoteInputOption.
const MessagePeerVoteInputOptionTypeID = 0x74cda504

// construct implements constructor of MessagePeerVoteClass.
func (m MessagePeerVoteInputOption) construct() MessagePeerVoteClass { return &m }

// Ensuring interfaces in compile-time for MessagePeerVoteInputOption.
var (
	_ bin.Encoder     = &MessagePeerVoteInputOption{}
	_ bin.Decoder     = &MessagePeerVoteInputOption{}
	_ bin.BareEncoder = &MessagePeerVoteInputOption{}
	_ bin.BareDecoder = &MessagePeerVoteInputOption{}

	_ MessagePeerVoteClass = &MessagePeerVoteInputOption{}
)

func (m *MessagePeerVoteInputOption) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Peer == nil) {
		return false
	}
	if !(m.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagePeerVoteInputOption) String() string {
	if m == nil {
		return "MessagePeerVoteInputOption(nil)"
	}
	type Alias MessagePeerVoteInputOption
	return fmt.Sprintf("MessagePeerVoteInputOption%+v", Alias(*m))
}

// FillFrom fills MessagePeerVoteInputOption from given interface.
func (m *MessagePeerVoteInputOption) FillFrom(from interface {
	GetPeer() (value PeerClass)
	GetDate() (value int)
}) {
	m.Peer = from.GetPeer()
	m.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagePeerVoteInputOption) TypeID() uint32 {
	return MessagePeerVoteInputOptionTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagePeerVoteInputOption) TypeName() string {
	return "messagePeerVoteInputOption"
}

// TypeInfo returns info about TL type.
func (m *MessagePeerVoteInputOption) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messagePeerVoteInputOption",
		ID:   MessagePeerVoteInputOptionTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessagePeerVoteInputOption) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messagePeerVoteInputOption#74cda504 as nil")
	}
	b.PutID(MessagePeerVoteInputOptionTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessagePeerVoteInputOption) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messagePeerVoteInputOption#74cda504 as nil")
	}
	if m.Peer == nil {
		return fmt.Errorf("unable to encode messagePeerVoteInputOption#74cda504: field peer is nil")
	}
	if err := m.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messagePeerVoteInputOption#74cda504: field peer: %w", err)
	}
	b.PutInt(m.Date)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessagePeerVoteInputOption) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messagePeerVoteInputOption#74cda504 to nil")
	}
	if err := b.ConsumeID(MessagePeerVoteInputOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode messagePeerVoteInputOption#74cda504: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessagePeerVoteInputOption) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messagePeerVoteInputOption#74cda504 to nil")
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messagePeerVoteInputOption#74cda504: field peer: %w", err)
		}
		m.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messagePeerVoteInputOption#74cda504: field date: %w", err)
		}
		m.Date = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (m *MessagePeerVoteInputOption) GetPeer() (value PeerClass) {
	if m == nil {
		return
	}
	return m.Peer
}

// GetDate returns value of Date field.
func (m *MessagePeerVoteInputOption) GetDate() (value int) {
	if m == nil {
		return
	}
	return m.Date
}

// MessagePeerVoteMultiple represents TL type `messagePeerVoteMultiple#4628f6e6`.
// How a peer voted in a multiple-choice poll
//
// See https://core.telegram.org/constructor/messagePeerVoteMultiple for reference.
type MessagePeerVoteMultiple struct {
	// Peer ID
	Peer PeerClass
	// Options chosen by the peer
	Options [][]byte
	// When did the peer cast their votes
	Date int
}

// MessagePeerVoteMultipleTypeID is TL type id of MessagePeerVoteMultiple.
const MessagePeerVoteMultipleTypeID = 0x4628f6e6

// construct implements constructor of MessagePeerVoteClass.
func (m MessagePeerVoteMultiple) construct() MessagePeerVoteClass { return &m }

// Ensuring interfaces in compile-time for MessagePeerVoteMultiple.
var (
	_ bin.Encoder     = &MessagePeerVoteMultiple{}
	_ bin.Decoder     = &MessagePeerVoteMultiple{}
	_ bin.BareEncoder = &MessagePeerVoteMultiple{}
	_ bin.BareDecoder = &MessagePeerVoteMultiple{}

	_ MessagePeerVoteClass = &MessagePeerVoteMultiple{}
)

func (m *MessagePeerVoteMultiple) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Peer == nil) {
		return false
	}
	if !(m.Options == nil) {
		return false
	}
	if !(m.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagePeerVoteMultiple) String() string {
	if m == nil {
		return "MessagePeerVoteMultiple(nil)"
	}
	type Alias MessagePeerVoteMultiple
	return fmt.Sprintf("MessagePeerVoteMultiple%+v", Alias(*m))
}

// FillFrom fills MessagePeerVoteMultiple from given interface.
func (m *MessagePeerVoteMultiple) FillFrom(from interface {
	GetPeer() (value PeerClass)
	GetOptions() (value [][]byte)
	GetDate() (value int)
}) {
	m.Peer = from.GetPeer()
	m.Options = from.GetOptions()
	m.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagePeerVoteMultiple) TypeID() uint32 {
	return MessagePeerVoteMultipleTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagePeerVoteMultiple) TypeName() string {
	return "messagePeerVoteMultiple"
}

// TypeInfo returns info about TL type.
func (m *MessagePeerVoteMultiple) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messagePeerVoteMultiple",
		ID:   MessagePeerVoteMultipleTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Options",
			SchemaName: "options",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessagePeerVoteMultiple) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messagePeerVoteMultiple#4628f6e6 as nil")
	}
	b.PutID(MessagePeerVoteMultipleTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessagePeerVoteMultiple) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messagePeerVoteMultiple#4628f6e6 as nil")
	}
	if m.Peer == nil {
		return fmt.Errorf("unable to encode messagePeerVoteMultiple#4628f6e6: field peer is nil")
	}
	if err := m.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messagePeerVoteMultiple#4628f6e6: field peer: %w", err)
	}
	b.PutVectorHeader(len(m.Options))
	for _, v := range m.Options {
		b.PutBytes(v)
	}
	b.PutInt(m.Date)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessagePeerVoteMultiple) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messagePeerVoteMultiple#4628f6e6 to nil")
	}
	if err := b.ConsumeID(MessagePeerVoteMultipleTypeID); err != nil {
		return fmt.Errorf("unable to decode messagePeerVoteMultiple#4628f6e6: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessagePeerVoteMultiple) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messagePeerVoteMultiple#4628f6e6 to nil")
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messagePeerVoteMultiple#4628f6e6: field peer: %w", err)
		}
		m.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messagePeerVoteMultiple#4628f6e6: field options: %w", err)
		}

		if headerLen > 0 {
			m.Options = make([][]byte, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode messagePeerVoteMultiple#4628f6e6: field options: %w", err)
			}
			m.Options = append(m.Options, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messagePeerVoteMultiple#4628f6e6: field date: %w", err)
		}
		m.Date = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (m *MessagePeerVoteMultiple) GetPeer() (value PeerClass) {
	if m == nil {
		return
	}
	return m.Peer
}

// GetOptions returns value of Options field.
func (m *MessagePeerVoteMultiple) GetOptions() (value [][]byte) {
	if m == nil {
		return
	}
	return m.Options
}

// GetDate returns value of Date field.
func (m *MessagePeerVoteMultiple) GetDate() (value int) {
	if m == nil {
		return
	}
	return m.Date
}

// MessagePeerVoteClassName is schema name of MessagePeerVoteClass.
const MessagePeerVoteClassName = "MessagePeerVote"

// MessagePeerVoteClass represents MessagePeerVote generic type.
//
// See https://core.telegram.org/type/MessagePeerVote for reference.
//
// Example:
//
//	g, err := tg.DecodeMessagePeerVote(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.MessagePeerVote: // messagePeerVote#b6cc2d5c
//	case *tg.MessagePeerVoteInputOption: // messagePeerVoteInputOption#74cda504
//	case *tg.MessagePeerVoteMultiple: // messagePeerVoteMultiple#4628f6e6
//	default: panic(v)
//	}
type MessagePeerVoteClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagePeerVoteClass

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

	// Peer ID
	GetPeer() (value PeerClass)

	// When did the peer cast the vote
	GetDate() (value int)
}

// DecodeMessagePeerVote implements binary de-serialization for MessagePeerVoteClass.
func DecodeMessagePeerVote(buf *bin.Buffer) (MessagePeerVoteClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagePeerVoteTypeID:
		// Decoding messagePeerVote#b6cc2d5c.
		v := MessagePeerVote{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagePeerVoteClass: %w", err)
		}
		return &v, nil
	case MessagePeerVoteInputOptionTypeID:
		// Decoding messagePeerVoteInputOption#74cda504.
		v := MessagePeerVoteInputOption{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagePeerVoteClass: %w", err)
		}
		return &v, nil
	case MessagePeerVoteMultipleTypeID:
		// Decoding messagePeerVoteMultiple#4628f6e6.
		v := MessagePeerVoteMultiple{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagePeerVoteClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagePeerVoteClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagePeerVote boxes the MessagePeerVoteClass providing a helper.
type MessagePeerVoteBox struct {
	MessagePeerVote MessagePeerVoteClass
}

// Decode implements bin.Decoder for MessagePeerVoteBox.
func (b *MessagePeerVoteBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagePeerVoteBox to nil")
	}
	v, err := DecodeMessagePeerVote(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessagePeerVote = v
	return nil
}

// Encode implements bin.Encode for MessagePeerVoteBox.
func (b *MessagePeerVoteBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MessagePeerVote == nil {
		return fmt.Errorf("unable to encode MessagePeerVoteClass as nil")
	}
	return b.MessagePeerVote.Encode(buf)
}
