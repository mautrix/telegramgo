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

// InputPeerEmpty represents TL type `inputPeerEmpty#7f3b18ea`.
// An empty constructor, no user or chat is defined.
//
// See https://core.telegram.org/constructor/inputPeerEmpty for reference.
type InputPeerEmpty struct {
}

// InputPeerEmptyTypeID is TL type id of InputPeerEmpty.
const InputPeerEmptyTypeID = 0x7f3b18ea

// construct implements constructor of InputPeerClass.
func (i InputPeerEmpty) construct() InputPeerClass { return &i }

// Ensuring interfaces in compile-time for InputPeerEmpty.
var (
	_ bin.Encoder     = &InputPeerEmpty{}
	_ bin.Decoder     = &InputPeerEmpty{}
	_ bin.BareEncoder = &InputPeerEmpty{}
	_ bin.BareDecoder = &InputPeerEmpty{}

	_ InputPeerClass = &InputPeerEmpty{}
)

func (i *InputPeerEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerEmpty) String() string {
	if i == nil {
		return "InputPeerEmpty(nil)"
	}
	type Alias InputPeerEmpty
	return fmt.Sprintf("InputPeerEmpty%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPeerEmpty) TypeID() uint32 {
	return InputPeerEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPeerEmpty) TypeName() string {
	return "inputPeerEmpty"
}

// TypeInfo returns info about TL type.
func (i *InputPeerEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPeerEmpty",
		ID:   InputPeerEmptyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPeerEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerEmpty#7f3b18ea as nil")
	}
	b.PutID(InputPeerEmptyTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPeerEmpty) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerEmpty#7f3b18ea as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPeerEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerEmpty#7f3b18ea to nil")
	}
	if err := b.ConsumeID(InputPeerEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerEmpty#7f3b18ea: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPeerEmpty) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerEmpty#7f3b18ea to nil")
	}
	return nil
}

// InputPeerSelf represents TL type `inputPeerSelf#7da07ec9`.
// Defines the current user.
//
// See https://core.telegram.org/constructor/inputPeerSelf for reference.
type InputPeerSelf struct {
}

// InputPeerSelfTypeID is TL type id of InputPeerSelf.
const InputPeerSelfTypeID = 0x7da07ec9

// construct implements constructor of InputPeerClass.
func (i InputPeerSelf) construct() InputPeerClass { return &i }

// Ensuring interfaces in compile-time for InputPeerSelf.
var (
	_ bin.Encoder     = &InputPeerSelf{}
	_ bin.Decoder     = &InputPeerSelf{}
	_ bin.BareEncoder = &InputPeerSelf{}
	_ bin.BareDecoder = &InputPeerSelf{}

	_ InputPeerClass = &InputPeerSelf{}
)

func (i *InputPeerSelf) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerSelf) String() string {
	if i == nil {
		return "InputPeerSelf(nil)"
	}
	type Alias InputPeerSelf
	return fmt.Sprintf("InputPeerSelf%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPeerSelf) TypeID() uint32 {
	return InputPeerSelfTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPeerSelf) TypeName() string {
	return "inputPeerSelf"
}

// TypeInfo returns info about TL type.
func (i *InputPeerSelf) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPeerSelf",
		ID:   InputPeerSelfTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPeerSelf) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerSelf#7da07ec9 as nil")
	}
	b.PutID(InputPeerSelfTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPeerSelf) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerSelf#7da07ec9 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPeerSelf) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerSelf#7da07ec9 to nil")
	}
	if err := b.ConsumeID(InputPeerSelfTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerSelf#7da07ec9: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPeerSelf) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerSelf#7da07ec9 to nil")
	}
	return nil
}

// InputPeerChat represents TL type `inputPeerChat#35a95cb9`.
// Defines a chat for further interaction.
//
// See https://core.telegram.org/constructor/inputPeerChat for reference.
type InputPeerChat struct {
	// Chat identifier
	ChatID int64
}

// InputPeerChatTypeID is TL type id of InputPeerChat.
const InputPeerChatTypeID = 0x35a95cb9

// construct implements constructor of InputPeerClass.
func (i InputPeerChat) construct() InputPeerClass { return &i }

// Ensuring interfaces in compile-time for InputPeerChat.
var (
	_ bin.Encoder     = &InputPeerChat{}
	_ bin.Decoder     = &InputPeerChat{}
	_ bin.BareEncoder = &InputPeerChat{}
	_ bin.BareDecoder = &InputPeerChat{}

	_ InputPeerClass = &InputPeerChat{}
)

func (i *InputPeerChat) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerChat) String() string {
	if i == nil {
		return "InputPeerChat(nil)"
	}
	type Alias InputPeerChat
	return fmt.Sprintf("InputPeerChat%+v", Alias(*i))
}

// FillFrom fills InputPeerChat from given interface.
func (i *InputPeerChat) FillFrom(from interface {
	GetChatID() (value int64)
}) {
	i.ChatID = from.GetChatID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPeerChat) TypeID() uint32 {
	return InputPeerChatTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPeerChat) TypeName() string {
	return "inputPeerChat"
}

// TypeInfo returns info about TL type.
func (i *InputPeerChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPeerChat",
		ID:   InputPeerChatTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPeerChat) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerChat#35a95cb9 as nil")
	}
	b.PutID(InputPeerChatTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPeerChat) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerChat#35a95cb9 as nil")
	}
	b.PutLong(i.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPeerChat) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerChat#35a95cb9 to nil")
	}
	if err := b.ConsumeID(InputPeerChatTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerChat#35a95cb9: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPeerChat) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerChat#35a95cb9 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerChat#35a95cb9: field chat_id: %w", err)
		}
		i.ChatID = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (i *InputPeerChat) GetChatID() (value int64) {
	if i == nil {
		return
	}
	return i.ChatID
}

// InputPeerUser represents TL type `inputPeerUser#dde8a54c`.
// Defines a user for further interaction.
//
// See https://core.telegram.org/constructor/inputPeerUser for reference.
type InputPeerUser struct {
	// User identifier
	UserID int64
	// access_hash value from the user¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/constructor/user
	AccessHash int64
}

// InputPeerUserTypeID is TL type id of InputPeerUser.
const InputPeerUserTypeID = 0xdde8a54c

// construct implements constructor of InputPeerClass.
func (i InputPeerUser) construct() InputPeerClass { return &i }

// Ensuring interfaces in compile-time for InputPeerUser.
var (
	_ bin.Encoder     = &InputPeerUser{}
	_ bin.Decoder     = &InputPeerUser{}
	_ bin.BareEncoder = &InputPeerUser{}
	_ bin.BareDecoder = &InputPeerUser{}

	_ InputPeerClass = &InputPeerUser{}
)

func (i *InputPeerUser) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.UserID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerUser) String() string {
	if i == nil {
		return "InputPeerUser(nil)"
	}
	type Alias InputPeerUser
	return fmt.Sprintf("InputPeerUser%+v", Alias(*i))
}

// FillFrom fills InputPeerUser from given interface.
func (i *InputPeerUser) FillFrom(from interface {
	GetUserID() (value int64)
	GetAccessHash() (value int64)
}) {
	i.UserID = from.GetUserID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPeerUser) TypeID() uint32 {
	return InputPeerUserTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPeerUser) TypeName() string {
	return "inputPeerUser"
}

// TypeInfo returns info about TL type.
func (i *InputPeerUser) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPeerUser",
		ID:   InputPeerUserTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPeerUser) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerUser#dde8a54c as nil")
	}
	b.PutID(InputPeerUserTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPeerUser) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerUser#dde8a54c as nil")
	}
	b.PutLong(i.UserID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPeerUser) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerUser#dde8a54c to nil")
	}
	if err := b.ConsumeID(InputPeerUserTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerUser#dde8a54c: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPeerUser) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerUser#dde8a54c to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerUser#dde8a54c: field user_id: %w", err)
		}
		i.UserID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerUser#dde8a54c: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (i *InputPeerUser) GetUserID() (value int64) {
	if i == nil {
		return
	}
	return i.UserID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputPeerUser) GetAccessHash() (value int64) {
	if i == nil {
		return
	}
	return i.AccessHash
}

// InputPeerChannel represents TL type `inputPeerChannel#27bcbbfc`.
// Defines a channel for further interaction.
//
// See https://core.telegram.org/constructor/inputPeerChannel for reference.
type InputPeerChannel struct {
	// Channel identifier
	ChannelID int64
	// access_hash value from the channel¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channel
	AccessHash int64
}

// InputPeerChannelTypeID is TL type id of InputPeerChannel.
const InputPeerChannelTypeID = 0x27bcbbfc

// construct implements constructor of InputPeerClass.
func (i InputPeerChannel) construct() InputPeerClass { return &i }

// Ensuring interfaces in compile-time for InputPeerChannel.
var (
	_ bin.Encoder     = &InputPeerChannel{}
	_ bin.Decoder     = &InputPeerChannel{}
	_ bin.BareEncoder = &InputPeerChannel{}
	_ bin.BareDecoder = &InputPeerChannel{}

	_ InputPeerClass = &InputPeerChannel{}
)

func (i *InputPeerChannel) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ChannelID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerChannel) String() string {
	if i == nil {
		return "InputPeerChannel(nil)"
	}
	type Alias InputPeerChannel
	return fmt.Sprintf("InputPeerChannel%+v", Alias(*i))
}

// FillFrom fills InputPeerChannel from given interface.
func (i *InputPeerChannel) FillFrom(from interface {
	GetChannelID() (value int64)
	GetAccessHash() (value int64)
}) {
	i.ChannelID = from.GetChannelID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPeerChannel) TypeID() uint32 {
	return InputPeerChannelTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPeerChannel) TypeName() string {
	return "inputPeerChannel"
}

// TypeInfo returns info about TL type.
func (i *InputPeerChannel) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPeerChannel",
		ID:   InputPeerChannelTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChannelID",
			SchemaName: "channel_id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPeerChannel) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerChannel#27bcbbfc as nil")
	}
	b.PutID(InputPeerChannelTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPeerChannel) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerChannel#27bcbbfc as nil")
	}
	b.PutLong(i.ChannelID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPeerChannel) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerChannel#27bcbbfc to nil")
	}
	if err := b.ConsumeID(InputPeerChannelTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerChannel#27bcbbfc: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPeerChannel) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerChannel#27bcbbfc to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerChannel#27bcbbfc: field channel_id: %w", err)
		}
		i.ChannelID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerChannel#27bcbbfc: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// GetChannelID returns value of ChannelID field.
func (i *InputPeerChannel) GetChannelID() (value int64) {
	if i == nil {
		return
	}
	return i.ChannelID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputPeerChannel) GetAccessHash() (value int64) {
	if i == nil {
		return
	}
	return i.AccessHash
}

// InputPeerUserFromMessage represents TL type `inputPeerUserFromMessage#a87b0a1c`.
// Defines a min¹ user that was seen in a certain message of a certain chat.
//
// Links:
//  1. https://core.telegram.org/api/min
//
// See https://core.telegram.org/constructor/inputPeerUserFromMessage for reference.
type InputPeerUserFromMessage struct {
	// The chat where the user was seen
	Peer InputPeerClass
	// The message ID
	MsgID int
	// The identifier of the user that was seen
	UserID int64
}

// InputPeerUserFromMessageTypeID is TL type id of InputPeerUserFromMessage.
const InputPeerUserFromMessageTypeID = 0xa87b0a1c

// construct implements constructor of InputPeerClass.
func (i InputPeerUserFromMessage) construct() InputPeerClass { return &i }

// Ensuring interfaces in compile-time for InputPeerUserFromMessage.
var (
	_ bin.Encoder     = &InputPeerUserFromMessage{}
	_ bin.Decoder     = &InputPeerUserFromMessage{}
	_ bin.BareEncoder = &InputPeerUserFromMessage{}
	_ bin.BareDecoder = &InputPeerUserFromMessage{}

	_ InputPeerClass = &InputPeerUserFromMessage{}
)

func (i *InputPeerUserFromMessage) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Peer == nil) {
		return false
	}
	if !(i.MsgID == 0) {
		return false
	}
	if !(i.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerUserFromMessage) String() string {
	if i == nil {
		return "InputPeerUserFromMessage(nil)"
	}
	type Alias InputPeerUserFromMessage
	return fmt.Sprintf("InputPeerUserFromMessage%+v", Alias(*i))
}

// FillFrom fills InputPeerUserFromMessage from given interface.
func (i *InputPeerUserFromMessage) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetUserID() (value int64)
}) {
	i.Peer = from.GetPeer()
	i.MsgID = from.GetMsgID()
	i.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPeerUserFromMessage) TypeID() uint32 {
	return InputPeerUserFromMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPeerUserFromMessage) TypeName() string {
	return "inputPeerUserFromMessage"
}

// TypeInfo returns info about TL type.
func (i *InputPeerUserFromMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPeerUserFromMessage",
		ID:   InputPeerUserFromMessageTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPeerUserFromMessage) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerUserFromMessage#a87b0a1c as nil")
	}
	b.PutID(InputPeerUserFromMessageTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPeerUserFromMessage) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerUserFromMessage#a87b0a1c as nil")
	}
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputPeerUserFromMessage#a87b0a1c: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPeerUserFromMessage#a87b0a1c: field peer: %w", err)
	}
	b.PutInt(i.MsgID)
	b.PutLong(i.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPeerUserFromMessage) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerUserFromMessage#a87b0a1c to nil")
	}
	if err := b.ConsumeID(InputPeerUserFromMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerUserFromMessage#a87b0a1c: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPeerUserFromMessage) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerUserFromMessage#a87b0a1c to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerUserFromMessage#a87b0a1c: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerUserFromMessage#a87b0a1c: field msg_id: %w", err)
		}
		i.MsgID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerUserFromMessage#a87b0a1c: field user_id: %w", err)
		}
		i.UserID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (i *InputPeerUserFromMessage) GetPeer() (value InputPeerClass) {
	if i == nil {
		return
	}
	return i.Peer
}

// GetMsgID returns value of MsgID field.
func (i *InputPeerUserFromMessage) GetMsgID() (value int) {
	if i == nil {
		return
	}
	return i.MsgID
}

// GetUserID returns value of UserID field.
func (i *InputPeerUserFromMessage) GetUserID() (value int64) {
	if i == nil {
		return
	}
	return i.UserID
}

// InputPeerChannelFromMessage represents TL type `inputPeerChannelFromMessage#bd2a0840`.
// Defines a min¹ channel that was seen in a certain message of a certain chat.
//
// Links:
//  1. https://core.telegram.org/api/min
//
// See https://core.telegram.org/constructor/inputPeerChannelFromMessage for reference.
type InputPeerChannelFromMessage struct {
	// The chat where the channel's message was seen
	Peer InputPeerClass
	// The message ID
	MsgID int
	// The identifier of the channel that was seen
	ChannelID int64
}

// InputPeerChannelFromMessageTypeID is TL type id of InputPeerChannelFromMessage.
const InputPeerChannelFromMessageTypeID = 0xbd2a0840

// construct implements constructor of InputPeerClass.
func (i InputPeerChannelFromMessage) construct() InputPeerClass { return &i }

// Ensuring interfaces in compile-time for InputPeerChannelFromMessage.
var (
	_ bin.Encoder     = &InputPeerChannelFromMessage{}
	_ bin.Decoder     = &InputPeerChannelFromMessage{}
	_ bin.BareEncoder = &InputPeerChannelFromMessage{}
	_ bin.BareDecoder = &InputPeerChannelFromMessage{}

	_ InputPeerClass = &InputPeerChannelFromMessage{}
)

func (i *InputPeerChannelFromMessage) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Peer == nil) {
		return false
	}
	if !(i.MsgID == 0) {
		return false
	}
	if !(i.ChannelID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerChannelFromMessage) String() string {
	if i == nil {
		return "InputPeerChannelFromMessage(nil)"
	}
	type Alias InputPeerChannelFromMessage
	return fmt.Sprintf("InputPeerChannelFromMessage%+v", Alias(*i))
}

// FillFrom fills InputPeerChannelFromMessage from given interface.
func (i *InputPeerChannelFromMessage) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetChannelID() (value int64)
}) {
	i.Peer = from.GetPeer()
	i.MsgID = from.GetMsgID()
	i.ChannelID = from.GetChannelID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPeerChannelFromMessage) TypeID() uint32 {
	return InputPeerChannelFromMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPeerChannelFromMessage) TypeName() string {
	return "inputPeerChannelFromMessage"
}

// TypeInfo returns info about TL type.
func (i *InputPeerChannelFromMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPeerChannelFromMessage",
		ID:   InputPeerChannelFromMessageTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "ChannelID",
			SchemaName: "channel_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPeerChannelFromMessage) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerChannelFromMessage#bd2a0840 as nil")
	}
	b.PutID(InputPeerChannelFromMessageTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPeerChannelFromMessage) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerChannelFromMessage#bd2a0840 as nil")
	}
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputPeerChannelFromMessage#bd2a0840: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPeerChannelFromMessage#bd2a0840: field peer: %w", err)
	}
	b.PutInt(i.MsgID)
	b.PutLong(i.ChannelID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPeerChannelFromMessage) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerChannelFromMessage#bd2a0840 to nil")
	}
	if err := b.ConsumeID(InputPeerChannelFromMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerChannelFromMessage#bd2a0840: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPeerChannelFromMessage) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerChannelFromMessage#bd2a0840 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerChannelFromMessage#bd2a0840: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerChannelFromMessage#bd2a0840: field msg_id: %w", err)
		}
		i.MsgID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerChannelFromMessage#bd2a0840: field channel_id: %w", err)
		}
		i.ChannelID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (i *InputPeerChannelFromMessage) GetPeer() (value InputPeerClass) {
	if i == nil {
		return
	}
	return i.Peer
}

// GetMsgID returns value of MsgID field.
func (i *InputPeerChannelFromMessage) GetMsgID() (value int) {
	if i == nil {
		return
	}
	return i.MsgID
}

// GetChannelID returns value of ChannelID field.
func (i *InputPeerChannelFromMessage) GetChannelID() (value int64) {
	if i == nil {
		return
	}
	return i.ChannelID
}

// InputPeerClassName is schema name of InputPeerClass.
const InputPeerClassName = "InputPeer"

// InputPeerClass represents InputPeer generic type.
//
// See https://core.telegram.org/type/InputPeer for reference.
//
// Example:
//
//	g, err := tg.DecodeInputPeer(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.InputPeerEmpty: // inputPeerEmpty#7f3b18ea
//	case *tg.InputPeerSelf: // inputPeerSelf#7da07ec9
//	case *tg.InputPeerChat: // inputPeerChat#35a95cb9
//	case *tg.InputPeerUser: // inputPeerUser#dde8a54c
//	case *tg.InputPeerChannel: // inputPeerChannel#27bcbbfc
//	case *tg.InputPeerUserFromMessage: // inputPeerUserFromMessage#a87b0a1c
//	case *tg.InputPeerChannelFromMessage: // inputPeerChannelFromMessage#bd2a0840
//	default: panic(v)
//	}
type InputPeerClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputPeerClass

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
}

// DecodeInputPeer implements binary de-serialization for InputPeerClass.
func DecodeInputPeer(buf *bin.Buffer) (InputPeerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPeerEmptyTypeID:
		// Decoding inputPeerEmpty#7f3b18ea.
		v := InputPeerEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPeerClass: %w", err)
		}
		return &v, nil
	case InputPeerSelfTypeID:
		// Decoding inputPeerSelf#7da07ec9.
		v := InputPeerSelf{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPeerClass: %w", err)
		}
		return &v, nil
	case InputPeerChatTypeID:
		// Decoding inputPeerChat#35a95cb9.
		v := InputPeerChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPeerClass: %w", err)
		}
		return &v, nil
	case InputPeerUserTypeID:
		// Decoding inputPeerUser#dde8a54c.
		v := InputPeerUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPeerClass: %w", err)
		}
		return &v, nil
	case InputPeerChannelTypeID:
		// Decoding inputPeerChannel#27bcbbfc.
		v := InputPeerChannel{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPeerClass: %w", err)
		}
		return &v, nil
	case InputPeerUserFromMessageTypeID:
		// Decoding inputPeerUserFromMessage#a87b0a1c.
		v := InputPeerUserFromMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPeerClass: %w", err)
		}
		return &v, nil
	case InputPeerChannelFromMessageTypeID:
		// Decoding inputPeerChannelFromMessage#bd2a0840.
		v := InputPeerChannelFromMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPeerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPeerClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputPeer boxes the InputPeerClass providing a helper.
type InputPeerBox struct {
	InputPeer InputPeerClass
}

// Decode implements bin.Decoder for InputPeerBox.
func (b *InputPeerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPeerBox to nil")
	}
	v, err := DecodeInputPeer(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPeer = v
	return nil
}

// Encode implements bin.Encode for InputPeerBox.
func (b *InputPeerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputPeer == nil {
		return fmt.Errorf("unable to encode InputPeerClass as nil")
	}
	return b.InputPeer.Encode(buf)
}
