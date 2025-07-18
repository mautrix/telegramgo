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

// MessagesDialogs represents TL type `messages.dialogs#15ba6c40`.
// Full list of chats with messages and auxiliary data.
//
// See https://core.telegram.org/constructor/messages.dialogs for reference.
type MessagesDialogs struct {
	// List of chats
	Dialogs []DialogClass
	// List of last messages from each chat
	Messages []MessageClass
	// List of groups mentioned in the chats
	Chats []ChatClass
	// List of users mentioned in messages and groups
	Users []UserClass
}

// MessagesDialogsTypeID is TL type id of MessagesDialogs.
const MessagesDialogsTypeID = 0x15ba6c40

// construct implements constructor of MessagesDialogsClass.
func (d MessagesDialogs) construct() MessagesDialogsClass { return &d }

// Ensuring interfaces in compile-time for MessagesDialogs.
var (
	_ bin.Encoder     = &MessagesDialogs{}
	_ bin.Decoder     = &MessagesDialogs{}
	_ bin.BareEncoder = &MessagesDialogs{}
	_ bin.BareDecoder = &MessagesDialogs{}

	_ MessagesDialogsClass = &MessagesDialogs{}
)

func (d *MessagesDialogs) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Dialogs == nil) {
		return false
	}
	if !(d.Messages == nil) {
		return false
	}
	if !(d.Chats == nil) {
		return false
	}
	if !(d.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDialogs) String() string {
	if d == nil {
		return "MessagesDialogs(nil)"
	}
	type Alias MessagesDialogs
	return fmt.Sprintf("MessagesDialogs%+v", Alias(*d))
}

// FillFrom fills MessagesDialogs from given interface.
func (d *MessagesDialogs) FillFrom(from interface {
	GetDialogs() (value []DialogClass)
	GetMessages() (value []MessageClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	d.Dialogs = from.GetDialogs()
	d.Messages = from.GetMessages()
	d.Chats = from.GetChats()
	d.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDialogs) TypeID() uint32 {
	return MessagesDialogsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDialogs) TypeName() string {
	return "messages.dialogs"
}

// TypeInfo returns info about TL type.
func (d *MessagesDialogs) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.dialogs",
		ID:   MessagesDialogsTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Dialogs",
			SchemaName: "dialogs",
		},
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *MessagesDialogs) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dialogs#15ba6c40 as nil")
	}
	b.PutID(MessagesDialogsTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDialogs) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dialogs#15ba6c40 as nil")
	}
	b.PutVectorHeader(len(d.Dialogs))
	for idx, v := range d.Dialogs {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field dialogs element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field dialogs element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Messages))
	for idx, v := range d.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Chats))
	for idx, v := range d.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Users))
	for idx, v := range d.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogs#15ba6c40: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDialogs) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dialogs#15ba6c40 to nil")
	}
	if err := b.ConsumeID(MessagesDialogsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDialogs) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dialogs#15ba6c40 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field dialogs: %w", err)
		}

		if headerLen > 0 {
			d.Dialogs = make([]DialogClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDialog(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field dialogs: %w", err)
			}
			d.Dialogs = append(d.Dialogs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field messages: %w", err)
		}

		if headerLen > 0 {
			d.Messages = make([]MessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field messages: %w", err)
			}
			d.Messages = append(d.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field chats: %w", err)
		}

		if headerLen > 0 {
			d.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field chats: %w", err)
			}
			d.Chats = append(d.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field users: %w", err)
		}

		if headerLen > 0 {
			d.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogs#15ba6c40: field users: %w", err)
			}
			d.Users = append(d.Users, value)
		}
	}
	return nil
}

// GetDialogs returns value of Dialogs field.
func (d *MessagesDialogs) GetDialogs() (value []DialogClass) {
	if d == nil {
		return
	}
	return d.Dialogs
}

// GetMessages returns value of Messages field.
func (d *MessagesDialogs) GetMessages() (value []MessageClass) {
	if d == nil {
		return
	}
	return d.Messages
}

// GetChats returns value of Chats field.
func (d *MessagesDialogs) GetChats() (value []ChatClass) {
	if d == nil {
		return
	}
	return d.Chats
}

// GetUsers returns value of Users field.
func (d *MessagesDialogs) GetUsers() (value []UserClass) {
	if d == nil {
		return
	}
	return d.Users
}

// MapDialogs returns field Dialogs wrapped in DialogClassArray helper.
func (d *MessagesDialogs) MapDialogs() (value DialogClassArray) {
	return DialogClassArray(d.Dialogs)
}

// MapMessages returns field Messages wrapped in MessageClassArray helper.
func (d *MessagesDialogs) MapMessages() (value MessageClassArray) {
	return MessageClassArray(d.Messages)
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (d *MessagesDialogs) MapChats() (value ChatClassArray) {
	return ChatClassArray(d.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (d *MessagesDialogs) MapUsers() (value UserClassArray) {
	return UserClassArray(d.Users)
}

// MessagesDialogsSlice represents TL type `messages.dialogsSlice#71e094f3`.
// Incomplete list of dialogs with messages and auxiliary data.
//
// See https://core.telegram.org/constructor/messages.dialogsSlice for reference.
type MessagesDialogsSlice struct {
	// Total number of dialogs
	Count int
	// List of dialogs
	Dialogs []DialogClass
	// List of last messages from dialogs
	Messages []MessageClass
	// List of chats mentioned in dialogs
	Chats []ChatClass
	// List of users mentioned in messages and chats
	Users []UserClass
}

// MessagesDialogsSliceTypeID is TL type id of MessagesDialogsSlice.
const MessagesDialogsSliceTypeID = 0x71e094f3

// construct implements constructor of MessagesDialogsClass.
func (d MessagesDialogsSlice) construct() MessagesDialogsClass { return &d }

// Ensuring interfaces in compile-time for MessagesDialogsSlice.
var (
	_ bin.Encoder     = &MessagesDialogsSlice{}
	_ bin.Decoder     = &MessagesDialogsSlice{}
	_ bin.BareEncoder = &MessagesDialogsSlice{}
	_ bin.BareDecoder = &MessagesDialogsSlice{}

	_ MessagesDialogsClass = &MessagesDialogsSlice{}
)

func (d *MessagesDialogsSlice) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Count == 0) {
		return false
	}
	if !(d.Dialogs == nil) {
		return false
	}
	if !(d.Messages == nil) {
		return false
	}
	if !(d.Chats == nil) {
		return false
	}
	if !(d.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDialogsSlice) String() string {
	if d == nil {
		return "MessagesDialogsSlice(nil)"
	}
	type Alias MessagesDialogsSlice
	return fmt.Sprintf("MessagesDialogsSlice%+v", Alias(*d))
}

// FillFrom fills MessagesDialogsSlice from given interface.
func (d *MessagesDialogsSlice) FillFrom(from interface {
	GetCount() (value int)
	GetDialogs() (value []DialogClass)
	GetMessages() (value []MessageClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	d.Count = from.GetCount()
	d.Dialogs = from.GetDialogs()
	d.Messages = from.GetMessages()
	d.Chats = from.GetChats()
	d.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDialogsSlice) TypeID() uint32 {
	return MessagesDialogsSliceTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDialogsSlice) TypeName() string {
	return "messages.dialogsSlice"
}

// TypeInfo returns info about TL type.
func (d *MessagesDialogsSlice) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.dialogsSlice",
		ID:   MessagesDialogsSliceTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Dialogs",
			SchemaName: "dialogs",
		},
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *MessagesDialogsSlice) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dialogsSlice#71e094f3 as nil")
	}
	b.PutID(MessagesDialogsSliceTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDialogsSlice) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dialogsSlice#71e094f3 as nil")
	}
	b.PutInt(d.Count)
	b.PutVectorHeader(len(d.Dialogs))
	for idx, v := range d.Dialogs {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field dialogs element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field dialogs element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Messages))
	for idx, v := range d.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Chats))
	for idx, v := range d.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Users))
	for idx, v := range d.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.dialogsSlice#71e094f3: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDialogsSlice) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dialogsSlice#71e094f3 to nil")
	}
	if err := b.ConsumeID(MessagesDialogsSliceTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDialogsSlice) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dialogsSlice#71e094f3 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field count: %w", err)
		}
		d.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field dialogs: %w", err)
		}

		if headerLen > 0 {
			d.Dialogs = make([]DialogClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDialog(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field dialogs: %w", err)
			}
			d.Dialogs = append(d.Dialogs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field messages: %w", err)
		}

		if headerLen > 0 {
			d.Messages = make([]MessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field messages: %w", err)
			}
			d.Messages = append(d.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field chats: %w", err)
		}

		if headerLen > 0 {
			d.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field chats: %w", err)
			}
			d.Chats = append(d.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field users: %w", err)
		}

		if headerLen > 0 {
			d.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.dialogsSlice#71e094f3: field users: %w", err)
			}
			d.Users = append(d.Users, value)
		}
	}
	return nil
}

// GetCount returns value of Count field.
func (d *MessagesDialogsSlice) GetCount() (value int) {
	if d == nil {
		return
	}
	return d.Count
}

// GetDialogs returns value of Dialogs field.
func (d *MessagesDialogsSlice) GetDialogs() (value []DialogClass) {
	if d == nil {
		return
	}
	return d.Dialogs
}

// GetMessages returns value of Messages field.
func (d *MessagesDialogsSlice) GetMessages() (value []MessageClass) {
	if d == nil {
		return
	}
	return d.Messages
}

// GetChats returns value of Chats field.
func (d *MessagesDialogsSlice) GetChats() (value []ChatClass) {
	if d == nil {
		return
	}
	return d.Chats
}

// GetUsers returns value of Users field.
func (d *MessagesDialogsSlice) GetUsers() (value []UserClass) {
	if d == nil {
		return
	}
	return d.Users
}

// MapDialogs returns field Dialogs wrapped in DialogClassArray helper.
func (d *MessagesDialogsSlice) MapDialogs() (value DialogClassArray) {
	return DialogClassArray(d.Dialogs)
}

// MapMessages returns field Messages wrapped in MessageClassArray helper.
func (d *MessagesDialogsSlice) MapMessages() (value MessageClassArray) {
	return MessageClassArray(d.Messages)
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (d *MessagesDialogsSlice) MapChats() (value ChatClassArray) {
	return ChatClassArray(d.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (d *MessagesDialogsSlice) MapUsers() (value UserClassArray) {
	return UserClassArray(d.Users)
}

// MessagesDialogsNotModified represents TL type `messages.dialogsNotModified#f0e3e596`.
// Dialogs haven't changed
//
// See https://core.telegram.org/constructor/messages.dialogsNotModified for reference.
type MessagesDialogsNotModified struct {
	// Number of dialogs found server-side by the query
	Count int
}

// MessagesDialogsNotModifiedTypeID is TL type id of MessagesDialogsNotModified.
const MessagesDialogsNotModifiedTypeID = 0xf0e3e596

// construct implements constructor of MessagesDialogsClass.
func (d MessagesDialogsNotModified) construct() MessagesDialogsClass { return &d }

// Ensuring interfaces in compile-time for MessagesDialogsNotModified.
var (
	_ bin.Encoder     = &MessagesDialogsNotModified{}
	_ bin.Decoder     = &MessagesDialogsNotModified{}
	_ bin.BareEncoder = &MessagesDialogsNotModified{}
	_ bin.BareDecoder = &MessagesDialogsNotModified{}

	_ MessagesDialogsClass = &MessagesDialogsNotModified{}
)

func (d *MessagesDialogsNotModified) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Count == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDialogsNotModified) String() string {
	if d == nil {
		return "MessagesDialogsNotModified(nil)"
	}
	type Alias MessagesDialogsNotModified
	return fmt.Sprintf("MessagesDialogsNotModified%+v", Alias(*d))
}

// FillFrom fills MessagesDialogsNotModified from given interface.
func (d *MessagesDialogsNotModified) FillFrom(from interface {
	GetCount() (value int)
}) {
	d.Count = from.GetCount()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDialogsNotModified) TypeID() uint32 {
	return MessagesDialogsNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDialogsNotModified) TypeName() string {
	return "messages.dialogsNotModified"
}

// TypeInfo returns info about TL type.
func (d *MessagesDialogsNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.dialogsNotModified",
		ID:   MessagesDialogsNotModifiedTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *MessagesDialogsNotModified) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dialogsNotModified#f0e3e596 as nil")
	}
	b.PutID(MessagesDialogsNotModifiedTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDialogsNotModified) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dialogsNotModified#f0e3e596 as nil")
	}
	b.PutInt(d.Count)
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDialogsNotModified) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dialogsNotModified#f0e3e596 to nil")
	}
	if err := b.ConsumeID(MessagesDialogsNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.dialogsNotModified#f0e3e596: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDialogsNotModified) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dialogsNotModified#f0e3e596 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dialogsNotModified#f0e3e596: field count: %w", err)
		}
		d.Count = value
	}
	return nil
}

// GetCount returns value of Count field.
func (d *MessagesDialogsNotModified) GetCount() (value int) {
	if d == nil {
		return
	}
	return d.Count
}

// MessagesDialogsClassName is schema name of MessagesDialogsClass.
const MessagesDialogsClassName = "messages.Dialogs"

// MessagesDialogsClass represents messages.Dialogs generic type.
//
// See https://core.telegram.org/type/messages.Dialogs for reference.
//
// Example:
//
//	g, err := tg.DecodeMessagesDialogs(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.MessagesDialogs: // messages.dialogs#15ba6c40
//	case *tg.MessagesDialogsSlice: // messages.dialogsSlice#71e094f3
//	case *tg.MessagesDialogsNotModified: // messages.dialogsNotModified#f0e3e596
//	default: panic(v)
//	}
type MessagesDialogsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesDialogsClass

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

	// AsModified tries to map MessagesDialogsClass to ModifiedMessagesDialogs.
	AsModified() (ModifiedMessagesDialogs, bool)
}

// ModifiedMessagesDialogs represents Modified subset of MessagesDialogsClass.
type ModifiedMessagesDialogs interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesDialogsClass

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

	// List of chats
	GetDialogs() (value []DialogClass)

	// List of last messages from each chat
	GetMessages() (value []MessageClass)

	// List of groups mentioned in the chats
	GetChats() (value []ChatClass)

	// List of users mentioned in messages and groups
	GetUsers() (value []UserClass)
}

// AsModified tries to map MessagesDialogs to ModifiedMessagesDialogs.
func (d *MessagesDialogs) AsModified() (ModifiedMessagesDialogs, bool) {
	value, ok := (MessagesDialogsClass(d)).(ModifiedMessagesDialogs)
	return value, ok
}

// AsModified tries to map MessagesDialogsSlice to ModifiedMessagesDialogs.
func (d *MessagesDialogsSlice) AsModified() (ModifiedMessagesDialogs, bool) {
	value, ok := (MessagesDialogsClass(d)).(ModifiedMessagesDialogs)
	return value, ok
}

// AsModified tries to map MessagesDialogsNotModified to ModifiedMessagesDialogs.
func (d *MessagesDialogsNotModified) AsModified() (ModifiedMessagesDialogs, bool) {
	value, ok := (MessagesDialogsClass(d)).(ModifiedMessagesDialogs)
	return value, ok
}

// DecodeMessagesDialogs implements binary de-serialization for MessagesDialogsClass.
func DecodeMessagesDialogs(buf *bin.Buffer) (MessagesDialogsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesDialogsTypeID:
		// Decoding messages.dialogs#15ba6c40.
		v := MessagesDialogs{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesDialogsClass: %w", err)
		}
		return &v, nil
	case MessagesDialogsSliceTypeID:
		// Decoding messages.dialogsSlice#71e094f3.
		v := MessagesDialogsSlice{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesDialogsClass: %w", err)
		}
		return &v, nil
	case MessagesDialogsNotModifiedTypeID:
		// Decoding messages.dialogsNotModified#f0e3e596.
		v := MessagesDialogsNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesDialogsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesDialogsClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesDialogs boxes the MessagesDialogsClass providing a helper.
type MessagesDialogsBox struct {
	Dialogs MessagesDialogsClass
}

// Decode implements bin.Decoder for MessagesDialogsBox.
func (b *MessagesDialogsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesDialogsBox to nil")
	}
	v, err := DecodeMessagesDialogs(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Dialogs = v
	return nil
}

// Encode implements bin.Encode for MessagesDialogsBox.
func (b *MessagesDialogsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Dialogs == nil {
		return fmt.Errorf("unable to encode MessagesDialogsClass as nil")
	}
	return b.Dialogs.Encode(buf)
}
