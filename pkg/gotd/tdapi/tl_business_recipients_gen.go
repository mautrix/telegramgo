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

// BusinessRecipients represents TL type `businessRecipients#802011e2`.
type BusinessRecipients struct {
	// Identifiers of selected private chats
	ChatIDs []int64
	// Identifiers of private chats that are always excluded; for businessConnectedBot only
	ExcludedChatIDs []int64
	// True, if all existing private chats are selected
	SelectExistingChats bool
	// True, if all new private chats are selected
	SelectNewChats bool
	// True, if all private chats with contacts are selected
	SelectContacts bool
	// True, if all private chats with non-contacts are selected
	SelectNonContacts bool
	// If true, then all private chats except the selected are chosen. Otherwise, only the
	// selected chats are chosen
	ExcludeSelected bool
}

// BusinessRecipientsTypeID is TL type id of BusinessRecipients.
const BusinessRecipientsTypeID = 0x802011e2

// Ensuring interfaces in compile-time for BusinessRecipients.
var (
	_ bin.Encoder     = &BusinessRecipients{}
	_ bin.Decoder     = &BusinessRecipients{}
	_ bin.BareEncoder = &BusinessRecipients{}
	_ bin.BareDecoder = &BusinessRecipients{}
)

func (b *BusinessRecipients) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.ChatIDs == nil) {
		return false
	}
	if !(b.ExcludedChatIDs == nil) {
		return false
	}
	if !(b.SelectExistingChats == false) {
		return false
	}
	if !(b.SelectNewChats == false) {
		return false
	}
	if !(b.SelectContacts == false) {
		return false
	}
	if !(b.SelectNonContacts == false) {
		return false
	}
	if !(b.ExcludeSelected == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BusinessRecipients) String() string {
	if b == nil {
		return "BusinessRecipients(nil)"
	}
	type Alias BusinessRecipients
	return fmt.Sprintf("BusinessRecipients%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BusinessRecipients) TypeID() uint32 {
	return BusinessRecipientsTypeID
}

// TypeName returns name of type in TL schema.
func (*BusinessRecipients) TypeName() string {
	return "businessRecipients"
}

// TypeInfo returns info about TL type.
func (b *BusinessRecipients) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "businessRecipients",
		ID:   BusinessRecipientsTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatIDs",
			SchemaName: "chat_ids",
		},
		{
			Name:       "ExcludedChatIDs",
			SchemaName: "excluded_chat_ids",
		},
		{
			Name:       "SelectExistingChats",
			SchemaName: "select_existing_chats",
		},
		{
			Name:       "SelectNewChats",
			SchemaName: "select_new_chats",
		},
		{
			Name:       "SelectContacts",
			SchemaName: "select_contacts",
		},
		{
			Name:       "SelectNonContacts",
			SchemaName: "select_non_contacts",
		},
		{
			Name:       "ExcludeSelected",
			SchemaName: "exclude_selected",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BusinessRecipients) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessRecipients#802011e2 as nil")
	}
	buf.PutID(BusinessRecipientsTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BusinessRecipients) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessRecipients#802011e2 as nil")
	}
	buf.PutInt(len(b.ChatIDs))
	for _, v := range b.ChatIDs {
		buf.PutInt53(v)
	}
	buf.PutInt(len(b.ExcludedChatIDs))
	for _, v := range b.ExcludedChatIDs {
		buf.PutInt53(v)
	}
	buf.PutBool(b.SelectExistingChats)
	buf.PutBool(b.SelectNewChats)
	buf.PutBool(b.SelectContacts)
	buf.PutBool(b.SelectNonContacts)
	buf.PutBool(b.ExcludeSelected)
	return nil
}

// Decode implements bin.Decoder.
func (b *BusinessRecipients) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessRecipients#802011e2 to nil")
	}
	if err := buf.ConsumeID(BusinessRecipientsTypeID); err != nil {
		return fmt.Errorf("unable to decode businessRecipients#802011e2: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BusinessRecipients) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessRecipients#802011e2 to nil")
	}
	{
		headerLen, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode businessRecipients#802011e2: field chat_ids: %w", err)
		}

		if headerLen > 0 {
			b.ChatIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := buf.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode businessRecipients#802011e2: field chat_ids: %w", err)
			}
			b.ChatIDs = append(b.ChatIDs, value)
		}
	}
	{
		headerLen, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode businessRecipients#802011e2: field excluded_chat_ids: %w", err)
		}

		if headerLen > 0 {
			b.ExcludedChatIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := buf.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode businessRecipients#802011e2: field excluded_chat_ids: %w", err)
			}
			b.ExcludedChatIDs = append(b.ExcludedChatIDs, value)
		}
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode businessRecipients#802011e2: field select_existing_chats: %w", err)
		}
		b.SelectExistingChats = value
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode businessRecipients#802011e2: field select_new_chats: %w", err)
		}
		b.SelectNewChats = value
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode businessRecipients#802011e2: field select_contacts: %w", err)
		}
		b.SelectContacts = value
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode businessRecipients#802011e2: field select_non_contacts: %w", err)
		}
		b.SelectNonContacts = value
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode businessRecipients#802011e2: field exclude_selected: %w", err)
		}
		b.ExcludeSelected = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BusinessRecipients) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode businessRecipients#802011e2 as nil")
	}
	buf.ObjStart()
	buf.PutID("businessRecipients")
	buf.Comma()
	buf.FieldStart("chat_ids")
	buf.ArrStart()
	for _, v := range b.ChatIDs {
		buf.PutInt53(v)
		buf.Comma()
	}
	buf.StripComma()
	buf.ArrEnd()
	buf.Comma()
	buf.FieldStart("excluded_chat_ids")
	buf.ArrStart()
	for _, v := range b.ExcludedChatIDs {
		buf.PutInt53(v)
		buf.Comma()
	}
	buf.StripComma()
	buf.ArrEnd()
	buf.Comma()
	buf.FieldStart("select_existing_chats")
	buf.PutBool(b.SelectExistingChats)
	buf.Comma()
	buf.FieldStart("select_new_chats")
	buf.PutBool(b.SelectNewChats)
	buf.Comma()
	buf.FieldStart("select_contacts")
	buf.PutBool(b.SelectContacts)
	buf.Comma()
	buf.FieldStart("select_non_contacts")
	buf.PutBool(b.SelectNonContacts)
	buf.Comma()
	buf.FieldStart("exclude_selected")
	buf.PutBool(b.ExcludeSelected)
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BusinessRecipients) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode businessRecipients#802011e2 to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("businessRecipients"); err != nil {
				return fmt.Errorf("unable to decode businessRecipients#802011e2: %w", err)
			}
		case "chat_ids":
			if err := buf.Arr(func(buf tdjson.Decoder) error {
				value, err := buf.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode businessRecipients#802011e2: field chat_ids: %w", err)
				}
				b.ChatIDs = append(b.ChatIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode businessRecipients#802011e2: field chat_ids: %w", err)
			}
		case "excluded_chat_ids":
			if err := buf.Arr(func(buf tdjson.Decoder) error {
				value, err := buf.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode businessRecipients#802011e2: field excluded_chat_ids: %w", err)
				}
				b.ExcludedChatIDs = append(b.ExcludedChatIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode businessRecipients#802011e2: field excluded_chat_ids: %w", err)
			}
		case "select_existing_chats":
			value, err := buf.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode businessRecipients#802011e2: field select_existing_chats: %w", err)
			}
			b.SelectExistingChats = value
		case "select_new_chats":
			value, err := buf.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode businessRecipients#802011e2: field select_new_chats: %w", err)
			}
			b.SelectNewChats = value
		case "select_contacts":
			value, err := buf.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode businessRecipients#802011e2: field select_contacts: %w", err)
			}
			b.SelectContacts = value
		case "select_non_contacts":
			value, err := buf.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode businessRecipients#802011e2: field select_non_contacts: %w", err)
			}
			b.SelectNonContacts = value
		case "exclude_selected":
			value, err := buf.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode businessRecipients#802011e2: field exclude_selected: %w", err)
			}
			b.ExcludeSelected = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetChatIDs returns value of ChatIDs field.
func (b *BusinessRecipients) GetChatIDs() (value []int64) {
	if b == nil {
		return
	}
	return b.ChatIDs
}

// GetExcludedChatIDs returns value of ExcludedChatIDs field.
func (b *BusinessRecipients) GetExcludedChatIDs() (value []int64) {
	if b == nil {
		return
	}
	return b.ExcludedChatIDs
}

// GetSelectExistingChats returns value of SelectExistingChats field.
func (b *BusinessRecipients) GetSelectExistingChats() (value bool) {
	if b == nil {
		return
	}
	return b.SelectExistingChats
}

// GetSelectNewChats returns value of SelectNewChats field.
func (b *BusinessRecipients) GetSelectNewChats() (value bool) {
	if b == nil {
		return
	}
	return b.SelectNewChats
}

// GetSelectContacts returns value of SelectContacts field.
func (b *BusinessRecipients) GetSelectContacts() (value bool) {
	if b == nil {
		return
	}
	return b.SelectContacts
}

// GetSelectNonContacts returns value of SelectNonContacts field.
func (b *BusinessRecipients) GetSelectNonContacts() (value bool) {
	if b == nil {
		return
	}
	return b.SelectNonContacts
}

// GetExcludeSelected returns value of ExcludeSelected field.
func (b *BusinessRecipients) GetExcludeSelected() (value bool) {
	if b == nil {
		return
	}
	return b.ExcludeSelected
}
