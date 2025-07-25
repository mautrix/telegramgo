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

// MessagesHistoryImportParsed represents TL type `messages.historyImportParsed#5e0fb7b9`.
// Contains information about a chat export file generated by a foreign chat app, click
// here for more info¹.
// If neither the pm or group flags are set, the specified chat export was generated from
// a chat of unknown type.
//
// Links:
//  1. https://core.telegram.org/api/import
//
// See https://core.telegram.org/constructor/messages.historyImportParsed for reference.
type MessagesHistoryImportParsed struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// The chat export file was generated from a private chat.
	Pm bool
	// The chat export file was generated from a group chat.
	Group bool
	// Title of the chat.
	//
	// Use SetTitle and GetTitle helpers.
	Title string
}

// MessagesHistoryImportParsedTypeID is TL type id of MessagesHistoryImportParsed.
const MessagesHistoryImportParsedTypeID = 0x5e0fb7b9

// Ensuring interfaces in compile-time for MessagesHistoryImportParsed.
var (
	_ bin.Encoder     = &MessagesHistoryImportParsed{}
	_ bin.Decoder     = &MessagesHistoryImportParsed{}
	_ bin.BareEncoder = &MessagesHistoryImportParsed{}
	_ bin.BareDecoder = &MessagesHistoryImportParsed{}
)

func (h *MessagesHistoryImportParsed) Zero() bool {
	if h == nil {
		return true
	}
	if !(h.Flags.Zero()) {
		return false
	}
	if !(h.Pm == false) {
		return false
	}
	if !(h.Group == false) {
		return false
	}
	if !(h.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (h *MessagesHistoryImportParsed) String() string {
	if h == nil {
		return "MessagesHistoryImportParsed(nil)"
	}
	type Alias MessagesHistoryImportParsed
	return fmt.Sprintf("MessagesHistoryImportParsed%+v", Alias(*h))
}

// FillFrom fills MessagesHistoryImportParsed from given interface.
func (h *MessagesHistoryImportParsed) FillFrom(from interface {
	GetPm() (value bool)
	GetGroup() (value bool)
	GetTitle() (value string, ok bool)
}) {
	h.Pm = from.GetPm()
	h.Group = from.GetGroup()
	if val, ok := from.GetTitle(); ok {
		h.Title = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesHistoryImportParsed) TypeID() uint32 {
	return MessagesHistoryImportParsedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesHistoryImportParsed) TypeName() string {
	return "messages.historyImportParsed"
}

// TypeInfo returns info about TL type.
func (h *MessagesHistoryImportParsed) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.historyImportParsed",
		ID:   MessagesHistoryImportParsedTypeID,
	}
	if h == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Pm",
			SchemaName: "pm",
			Null:       !h.Flags.Has(0),
		},
		{
			Name:       "Group",
			SchemaName: "group",
			Null:       !h.Flags.Has(1),
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !h.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (h *MessagesHistoryImportParsed) SetFlags() {
	if !(h.Pm == false) {
		h.Flags.Set(0)
	}
	if !(h.Group == false) {
		h.Flags.Set(1)
	}
	if !(h.Title == "") {
		h.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (h *MessagesHistoryImportParsed) Encode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode messages.historyImportParsed#5e0fb7b9 as nil")
	}
	b.PutID(MessagesHistoryImportParsedTypeID)
	return h.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (h *MessagesHistoryImportParsed) EncodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode messages.historyImportParsed#5e0fb7b9 as nil")
	}
	h.SetFlags()
	if err := h.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.historyImportParsed#5e0fb7b9: field flags: %w", err)
	}
	if h.Flags.Has(2) {
		b.PutString(h.Title)
	}
	return nil
}

// Decode implements bin.Decoder.
func (h *MessagesHistoryImportParsed) Decode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode messages.historyImportParsed#5e0fb7b9 to nil")
	}
	if err := b.ConsumeID(MessagesHistoryImportParsedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.historyImportParsed#5e0fb7b9: %w", err)
	}
	return h.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (h *MessagesHistoryImportParsed) DecodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode messages.historyImportParsed#5e0fb7b9 to nil")
	}
	{
		if err := h.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.historyImportParsed#5e0fb7b9: field flags: %w", err)
		}
	}
	h.Pm = h.Flags.Has(0)
	h.Group = h.Flags.Has(1)
	if h.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.historyImportParsed#5e0fb7b9: field title: %w", err)
		}
		h.Title = value
	}
	return nil
}

// SetPm sets value of Pm conditional field.
func (h *MessagesHistoryImportParsed) SetPm(value bool) {
	if value {
		h.Flags.Set(0)
		h.Pm = true
	} else {
		h.Flags.Unset(0)
		h.Pm = false
	}
}

// GetPm returns value of Pm conditional field.
func (h *MessagesHistoryImportParsed) GetPm() (value bool) {
	if h == nil {
		return
	}
	return h.Flags.Has(0)
}

// SetGroup sets value of Group conditional field.
func (h *MessagesHistoryImportParsed) SetGroup(value bool) {
	if value {
		h.Flags.Set(1)
		h.Group = true
	} else {
		h.Flags.Unset(1)
		h.Group = false
	}
}

// GetGroup returns value of Group conditional field.
func (h *MessagesHistoryImportParsed) GetGroup() (value bool) {
	if h == nil {
		return
	}
	return h.Flags.Has(1)
}

// SetTitle sets value of Title conditional field.
func (h *MessagesHistoryImportParsed) SetTitle(value string) {
	h.Flags.Set(2)
	h.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (h *MessagesHistoryImportParsed) GetTitle() (value string, ok bool) {
	if h == nil {
		return
	}
	if !h.Flags.Has(2) {
		return value, false
	}
	return h.Title, true
}
