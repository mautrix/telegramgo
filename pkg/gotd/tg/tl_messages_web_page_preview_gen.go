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

// MessagesWebPagePreview represents TL type `messages.webPagePreview#b53e8b21`.
//
// See https://core.telegram.org/constructor/messages.webPagePreview for reference.
type MessagesWebPagePreview struct {
	// Media field of MessagesWebPagePreview.
	Media MessageMediaClass
	// Users field of MessagesWebPagePreview.
	Users []UserClass
}

// MessagesWebPagePreviewTypeID is TL type id of MessagesWebPagePreview.
const MessagesWebPagePreviewTypeID = 0xb53e8b21

// Ensuring interfaces in compile-time for MessagesWebPagePreview.
var (
	_ bin.Encoder     = &MessagesWebPagePreview{}
	_ bin.Decoder     = &MessagesWebPagePreview{}
	_ bin.BareEncoder = &MessagesWebPagePreview{}
	_ bin.BareDecoder = &MessagesWebPagePreview{}
)

func (w *MessagesWebPagePreview) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Media == nil) {
		return false
	}
	if !(w.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *MessagesWebPagePreview) String() string {
	if w == nil {
		return "MessagesWebPagePreview(nil)"
	}
	type Alias MessagesWebPagePreview
	return fmt.Sprintf("MessagesWebPagePreview%+v", Alias(*w))
}

// FillFrom fills MessagesWebPagePreview from given interface.
func (w *MessagesWebPagePreview) FillFrom(from interface {
	GetMedia() (value MessageMediaClass)
	GetUsers() (value []UserClass)
}) {
	w.Media = from.GetMedia()
	w.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesWebPagePreview) TypeID() uint32 {
	return MessagesWebPagePreviewTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesWebPagePreview) TypeName() string {
	return "messages.webPagePreview"
}

// TypeInfo returns info about TL type.
func (w *MessagesWebPagePreview) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.webPagePreview",
		ID:   MessagesWebPagePreviewTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Media",
			SchemaName: "media",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *MessagesWebPagePreview) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode messages.webPagePreview#b53e8b21 as nil")
	}
	b.PutID(MessagesWebPagePreviewTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *MessagesWebPagePreview) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode messages.webPagePreview#b53e8b21 as nil")
	}
	if w.Media == nil {
		return fmt.Errorf("unable to encode messages.webPagePreview#b53e8b21: field media is nil")
	}
	if err := w.Media.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.webPagePreview#b53e8b21: field media: %w", err)
	}
	b.PutVectorHeader(len(w.Users))
	for idx, v := range w.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.webPagePreview#b53e8b21: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.webPagePreview#b53e8b21: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *MessagesWebPagePreview) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode messages.webPagePreview#b53e8b21 to nil")
	}
	if err := b.ConsumeID(MessagesWebPagePreviewTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.webPagePreview#b53e8b21: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *MessagesWebPagePreview) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode messages.webPagePreview#b53e8b21 to nil")
	}
	{
		value, err := DecodeMessageMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.webPagePreview#b53e8b21: field media: %w", err)
		}
		w.Media = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.webPagePreview#b53e8b21: field users: %w", err)
		}

		if headerLen > 0 {
			w.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.webPagePreview#b53e8b21: field users: %w", err)
			}
			w.Users = append(w.Users, value)
		}
	}
	return nil
}

// GetMedia returns value of Media field.
func (w *MessagesWebPagePreview) GetMedia() (value MessageMediaClass) {
	if w == nil {
		return
	}
	return w.Media
}

// GetUsers returns value of Users field.
func (w *MessagesWebPagePreview) GetUsers() (value []UserClass) {
	if w == nil {
		return
	}
	return w.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (w *MessagesWebPagePreview) MapUsers() (value UserClassArray) {
	return UserClassArray(w.Users)
}
