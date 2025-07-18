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

// WebApp represents TL type `webApp#605ba8f3`.
type WebApp struct {
	// Web App short name
	ShortName string
	// Web App title
	Title string
	// Describes a Web App. Use getInternalLink with internalLinkTypeWebApp to share the Web
	// App
	Description string
	// Web App photo
	Photo Photo
	// Web App animation; may be null
	Animation Animation
}

// WebAppTypeID is TL type id of WebApp.
const WebAppTypeID = 0x605ba8f3

// Ensuring interfaces in compile-time for WebApp.
var (
	_ bin.Encoder     = &WebApp{}
	_ bin.Decoder     = &WebApp{}
	_ bin.BareEncoder = &WebApp{}
	_ bin.BareDecoder = &WebApp{}
)

func (w *WebApp) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.ShortName == "") {
		return false
	}
	if !(w.Title == "") {
		return false
	}
	if !(w.Description == "") {
		return false
	}
	if !(w.Photo.Zero()) {
		return false
	}
	if !(w.Animation.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebApp) String() string {
	if w == nil {
		return "WebApp(nil)"
	}
	type Alias WebApp
	return fmt.Sprintf("WebApp%+v", Alias(*w))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebApp) TypeID() uint32 {
	return WebAppTypeID
}

// TypeName returns name of type in TL schema.
func (*WebApp) TypeName() string {
	return "webApp"
}

// TypeInfo returns info about TL type.
func (w *WebApp) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webApp",
		ID:   WebAppTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ShortName",
			SchemaName: "short_name",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "Animation",
			SchemaName: "animation",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *WebApp) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webApp#605ba8f3 as nil")
	}
	b.PutID(WebAppTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WebApp) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webApp#605ba8f3 as nil")
	}
	b.PutString(w.ShortName)
	b.PutString(w.Title)
	b.PutString(w.Description)
	if err := w.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webApp#605ba8f3: field photo: %w", err)
	}
	if err := w.Animation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webApp#605ba8f3: field animation: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *WebApp) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webApp#605ba8f3 to nil")
	}
	if err := b.ConsumeID(WebAppTypeID); err != nil {
		return fmt.Errorf("unable to decode webApp#605ba8f3: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WebApp) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webApp#605ba8f3 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webApp#605ba8f3: field short_name: %w", err)
		}
		w.ShortName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webApp#605ba8f3: field title: %w", err)
		}
		w.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webApp#605ba8f3: field description: %w", err)
		}
		w.Description = value
	}
	{
		if err := w.Photo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webApp#605ba8f3: field photo: %w", err)
		}
	}
	{
		if err := w.Animation.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webApp#605ba8f3: field animation: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (w *WebApp) EncodeTDLibJSON(b tdjson.Encoder) error {
	if w == nil {
		return fmt.Errorf("can't encode webApp#605ba8f3 as nil")
	}
	b.ObjStart()
	b.PutID("webApp")
	b.Comma()
	b.FieldStart("short_name")
	b.PutString(w.ShortName)
	b.Comma()
	b.FieldStart("title")
	b.PutString(w.Title)
	b.Comma()
	b.FieldStart("description")
	b.PutString(w.Description)
	b.Comma()
	b.FieldStart("photo")
	if err := w.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode webApp#605ba8f3: field photo: %w", err)
	}
	b.Comma()
	b.FieldStart("animation")
	if err := w.Animation.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode webApp#605ba8f3: field animation: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (w *WebApp) DecodeTDLibJSON(b tdjson.Decoder) error {
	if w == nil {
		return fmt.Errorf("can't decode webApp#605ba8f3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("webApp"); err != nil {
				return fmt.Errorf("unable to decode webApp#605ba8f3: %w", err)
			}
		case "short_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode webApp#605ba8f3: field short_name: %w", err)
			}
			w.ShortName = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode webApp#605ba8f3: field title: %w", err)
			}
			w.Title = value
		case "description":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode webApp#605ba8f3: field description: %w", err)
			}
			w.Description = value
		case "photo":
			if err := w.Photo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode webApp#605ba8f3: field photo: %w", err)
			}
		case "animation":
			if err := w.Animation.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode webApp#605ba8f3: field animation: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetShortName returns value of ShortName field.
func (w *WebApp) GetShortName() (value string) {
	if w == nil {
		return
	}
	return w.ShortName
}

// GetTitle returns value of Title field.
func (w *WebApp) GetTitle() (value string) {
	if w == nil {
		return
	}
	return w.Title
}

// GetDescription returns value of Description field.
func (w *WebApp) GetDescription() (value string) {
	if w == nil {
		return
	}
	return w.Description
}

// GetPhoto returns value of Photo field.
func (w *WebApp) GetPhoto() (value Photo) {
	if w == nil {
		return
	}
	return w.Photo
}

// GetAnimation returns value of Animation field.
func (w *WebApp) GetAnimation() (value Animation) {
	if w == nil {
		return
	}
	return w.Animation
}
