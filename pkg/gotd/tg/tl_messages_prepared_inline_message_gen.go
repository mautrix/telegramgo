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

// MessagesPreparedInlineMessage represents TL type `messages.preparedInlineMessage#ff57708d`.
//
// See https://core.telegram.org/constructor/messages.preparedInlineMessage for reference.
type MessagesPreparedInlineMessage struct {
	// QueryID field of MessagesPreparedInlineMessage.
	QueryID int64
	// Result field of MessagesPreparedInlineMessage.
	Result BotInlineResultClass
	// PeerTypes field of MessagesPreparedInlineMessage.
	PeerTypes []InlineQueryPeerTypeClass
	// CacheTime field of MessagesPreparedInlineMessage.
	CacheTime int
	// Users field of MessagesPreparedInlineMessage.
	Users []UserClass
}

// MessagesPreparedInlineMessageTypeID is TL type id of MessagesPreparedInlineMessage.
const MessagesPreparedInlineMessageTypeID = 0xff57708d

// Ensuring interfaces in compile-time for MessagesPreparedInlineMessage.
var (
	_ bin.Encoder     = &MessagesPreparedInlineMessage{}
	_ bin.Decoder     = &MessagesPreparedInlineMessage{}
	_ bin.BareEncoder = &MessagesPreparedInlineMessage{}
	_ bin.BareDecoder = &MessagesPreparedInlineMessage{}
)

func (p *MessagesPreparedInlineMessage) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.QueryID == 0) {
		return false
	}
	if !(p.Result == nil) {
		return false
	}
	if !(p.PeerTypes == nil) {
		return false
	}
	if !(p.CacheTime == 0) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *MessagesPreparedInlineMessage) String() string {
	if p == nil {
		return "MessagesPreparedInlineMessage(nil)"
	}
	type Alias MessagesPreparedInlineMessage
	return fmt.Sprintf("MessagesPreparedInlineMessage%+v", Alias(*p))
}

// FillFrom fills MessagesPreparedInlineMessage from given interface.
func (p *MessagesPreparedInlineMessage) FillFrom(from interface {
	GetQueryID() (value int64)
	GetResult() (value BotInlineResultClass)
	GetPeerTypes() (value []InlineQueryPeerTypeClass)
	GetCacheTime() (value int)
	GetUsers() (value []UserClass)
}) {
	p.QueryID = from.GetQueryID()
	p.Result = from.GetResult()
	p.PeerTypes = from.GetPeerTypes()
	p.CacheTime = from.GetCacheTime()
	p.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesPreparedInlineMessage) TypeID() uint32 {
	return MessagesPreparedInlineMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesPreparedInlineMessage) TypeName() string {
	return "messages.preparedInlineMessage"
}

// TypeInfo returns info about TL type.
func (p *MessagesPreparedInlineMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.preparedInlineMessage",
		ID:   MessagesPreparedInlineMessageTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "QueryID",
			SchemaName: "query_id",
		},
		{
			Name:       "Result",
			SchemaName: "result",
		},
		{
			Name:       "PeerTypes",
			SchemaName: "peer_types",
		},
		{
			Name:       "CacheTime",
			SchemaName: "cache_time",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *MessagesPreparedInlineMessage) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode messages.preparedInlineMessage#ff57708d as nil")
	}
	b.PutID(MessagesPreparedInlineMessageTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *MessagesPreparedInlineMessage) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode messages.preparedInlineMessage#ff57708d as nil")
	}
	b.PutLong(p.QueryID)
	if p.Result == nil {
		return fmt.Errorf("unable to encode messages.preparedInlineMessage#ff57708d: field result is nil")
	}
	if err := p.Result.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.preparedInlineMessage#ff57708d: field result: %w", err)
	}
	b.PutVectorHeader(len(p.PeerTypes))
	for idx, v := range p.PeerTypes {
		if v == nil {
			return fmt.Errorf("unable to encode messages.preparedInlineMessage#ff57708d: field peer_types element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.preparedInlineMessage#ff57708d: field peer_types element with index %d: %w", idx, err)
		}
	}
	b.PutInt(p.CacheTime)
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.preparedInlineMessage#ff57708d: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.preparedInlineMessage#ff57708d: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *MessagesPreparedInlineMessage) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode messages.preparedInlineMessage#ff57708d to nil")
	}
	if err := b.ConsumeID(MessagesPreparedInlineMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.preparedInlineMessage#ff57708d: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *MessagesPreparedInlineMessage) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode messages.preparedInlineMessage#ff57708d to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.preparedInlineMessage#ff57708d: field query_id: %w", err)
		}
		p.QueryID = value
	}
	{
		value, err := DecodeBotInlineResult(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.preparedInlineMessage#ff57708d: field result: %w", err)
		}
		p.Result = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.preparedInlineMessage#ff57708d: field peer_types: %w", err)
		}

		if headerLen > 0 {
			p.PeerTypes = make([]InlineQueryPeerTypeClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInlineQueryPeerType(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.preparedInlineMessage#ff57708d: field peer_types: %w", err)
			}
			p.PeerTypes = append(p.PeerTypes, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.preparedInlineMessage#ff57708d: field cache_time: %w", err)
		}
		p.CacheTime = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.preparedInlineMessage#ff57708d: field users: %w", err)
		}

		if headerLen > 0 {
			p.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.preparedInlineMessage#ff57708d: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// GetQueryID returns value of QueryID field.
func (p *MessagesPreparedInlineMessage) GetQueryID() (value int64) {
	if p == nil {
		return
	}
	return p.QueryID
}

// GetResult returns value of Result field.
func (p *MessagesPreparedInlineMessage) GetResult() (value BotInlineResultClass) {
	if p == nil {
		return
	}
	return p.Result
}

// GetPeerTypes returns value of PeerTypes field.
func (p *MessagesPreparedInlineMessage) GetPeerTypes() (value []InlineQueryPeerTypeClass) {
	if p == nil {
		return
	}
	return p.PeerTypes
}

// GetCacheTime returns value of CacheTime field.
func (p *MessagesPreparedInlineMessage) GetCacheTime() (value int) {
	if p == nil {
		return
	}
	return p.CacheTime
}

// GetUsers returns value of Users field.
func (p *MessagesPreparedInlineMessage) GetUsers() (value []UserClass) {
	if p == nil {
		return
	}
	return p.Users
}

// MapPeerTypes returns field PeerTypes wrapped in InlineQueryPeerTypeClassArray helper.
func (p *MessagesPreparedInlineMessage) MapPeerTypes() (value InlineQueryPeerTypeClassArray) {
	return InlineQueryPeerTypeClassArray(p.PeerTypes)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (p *MessagesPreparedInlineMessage) MapUsers() (value UserClassArray) {
	return UserClassArray(p.Users)
}
