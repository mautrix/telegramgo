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

// ThemeSettings represents TL type `themeSettings#fc4c1c12`.
type ThemeSettings struct {
	// Theme accent color in ARGB format
	AccentColor int32
	// The background to be used in chats; may be null
	Background Background
	// The fill to be used as a background for outgoing messages
	OutgoingMessageFill BackgroundFillClass
	// If true, the freeform gradient fill needs to be animated on every sent message
	AnimateOutgoingMessageFill bool
	// Accent color of outgoing messages in ARGB format
	OutgoingMessageAccentColor int32
}

// ThemeSettingsTypeID is TL type id of ThemeSettings.
const ThemeSettingsTypeID = 0xfc4c1c12

// Ensuring interfaces in compile-time for ThemeSettings.
var (
	_ bin.Encoder     = &ThemeSettings{}
	_ bin.Decoder     = &ThemeSettings{}
	_ bin.BareEncoder = &ThemeSettings{}
	_ bin.BareDecoder = &ThemeSettings{}
)

func (t *ThemeSettings) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.AccentColor == 0) {
		return false
	}
	if !(t.Background.Zero()) {
		return false
	}
	if !(t.OutgoingMessageFill == nil) {
		return false
	}
	if !(t.AnimateOutgoingMessageFill == false) {
		return false
	}
	if !(t.OutgoingMessageAccentColor == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ThemeSettings) String() string {
	if t == nil {
		return "ThemeSettings(nil)"
	}
	type Alias ThemeSettings
	return fmt.Sprintf("ThemeSettings%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ThemeSettings) TypeID() uint32 {
	return ThemeSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*ThemeSettings) TypeName() string {
	return "themeSettings"
}

// TypeInfo returns info about TL type.
func (t *ThemeSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "themeSettings",
		ID:   ThemeSettingsTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "AccentColor",
			SchemaName: "accent_color",
		},
		{
			Name:       "Background",
			SchemaName: "background",
		},
		{
			Name:       "OutgoingMessageFill",
			SchemaName: "outgoing_message_fill",
		},
		{
			Name:       "AnimateOutgoingMessageFill",
			SchemaName: "animate_outgoing_message_fill",
		},
		{
			Name:       "OutgoingMessageAccentColor",
			SchemaName: "outgoing_message_accent_color",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ThemeSettings) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode themeSettings#fc4c1c12 as nil")
	}
	b.PutID(ThemeSettingsTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ThemeSettings) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode themeSettings#fc4c1c12 as nil")
	}
	b.PutInt32(t.AccentColor)
	if err := t.Background.Encode(b); err != nil {
		return fmt.Errorf("unable to encode themeSettings#fc4c1c12: field background: %w", err)
	}
	if t.OutgoingMessageFill == nil {
		return fmt.Errorf("unable to encode themeSettings#fc4c1c12: field outgoing_message_fill is nil")
	}
	if err := t.OutgoingMessageFill.Encode(b); err != nil {
		return fmt.Errorf("unable to encode themeSettings#fc4c1c12: field outgoing_message_fill: %w", err)
	}
	b.PutBool(t.AnimateOutgoingMessageFill)
	b.PutInt32(t.OutgoingMessageAccentColor)
	return nil
}

// Decode implements bin.Decoder.
func (t *ThemeSettings) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode themeSettings#fc4c1c12 to nil")
	}
	if err := b.ConsumeID(ThemeSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode themeSettings#fc4c1c12: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ThemeSettings) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode themeSettings#fc4c1c12 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#fc4c1c12: field accent_color: %w", err)
		}
		t.AccentColor = value
	}
	{
		if err := t.Background.Decode(b); err != nil {
			return fmt.Errorf("unable to decode themeSettings#fc4c1c12: field background: %w", err)
		}
	}
	{
		value, err := DecodeBackgroundFill(b)
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#fc4c1c12: field outgoing_message_fill: %w", err)
		}
		t.OutgoingMessageFill = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#fc4c1c12: field animate_outgoing_message_fill: %w", err)
		}
		t.AnimateOutgoingMessageFill = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#fc4c1c12: field outgoing_message_accent_color: %w", err)
		}
		t.OutgoingMessageAccentColor = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ThemeSettings) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode themeSettings#fc4c1c12 as nil")
	}
	b.ObjStart()
	b.PutID("themeSettings")
	b.Comma()
	b.FieldStart("accent_color")
	b.PutInt32(t.AccentColor)
	b.Comma()
	b.FieldStart("background")
	if err := t.Background.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode themeSettings#fc4c1c12: field background: %w", err)
	}
	b.Comma()
	b.FieldStart("outgoing_message_fill")
	if t.OutgoingMessageFill == nil {
		return fmt.Errorf("unable to encode themeSettings#fc4c1c12: field outgoing_message_fill is nil")
	}
	if err := t.OutgoingMessageFill.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode themeSettings#fc4c1c12: field outgoing_message_fill: %w", err)
	}
	b.Comma()
	b.FieldStart("animate_outgoing_message_fill")
	b.PutBool(t.AnimateOutgoingMessageFill)
	b.Comma()
	b.FieldStart("outgoing_message_accent_color")
	b.PutInt32(t.OutgoingMessageAccentColor)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ThemeSettings) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode themeSettings#fc4c1c12 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("themeSettings"); err != nil {
				return fmt.Errorf("unable to decode themeSettings#fc4c1c12: %w", err)
			}
		case "accent_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeSettings#fc4c1c12: field accent_color: %w", err)
			}
			t.AccentColor = value
		case "background":
			if err := t.Background.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode themeSettings#fc4c1c12: field background: %w", err)
			}
		case "outgoing_message_fill":
			value, err := DecodeTDLibJSONBackgroundFill(b)
			if err != nil {
				return fmt.Errorf("unable to decode themeSettings#fc4c1c12: field outgoing_message_fill: %w", err)
			}
			t.OutgoingMessageFill = value
		case "animate_outgoing_message_fill":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode themeSettings#fc4c1c12: field animate_outgoing_message_fill: %w", err)
			}
			t.AnimateOutgoingMessageFill = value
		case "outgoing_message_accent_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeSettings#fc4c1c12: field outgoing_message_accent_color: %w", err)
			}
			t.OutgoingMessageAccentColor = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAccentColor returns value of AccentColor field.
func (t *ThemeSettings) GetAccentColor() (value int32) {
	if t == nil {
		return
	}
	return t.AccentColor
}

// GetBackground returns value of Background field.
func (t *ThemeSettings) GetBackground() (value Background) {
	if t == nil {
		return
	}
	return t.Background
}

// GetOutgoingMessageFill returns value of OutgoingMessageFill field.
func (t *ThemeSettings) GetOutgoingMessageFill() (value BackgroundFillClass) {
	if t == nil {
		return
	}
	return t.OutgoingMessageFill
}

// GetAnimateOutgoingMessageFill returns value of AnimateOutgoingMessageFill field.
func (t *ThemeSettings) GetAnimateOutgoingMessageFill() (value bool) {
	if t == nil {
		return
	}
	return t.AnimateOutgoingMessageFill
}

// GetOutgoingMessageAccentColor returns value of OutgoingMessageAccentColor field.
func (t *ThemeSettings) GetOutgoingMessageAccentColor() (value int32) {
	if t == nil {
		return
	}
	return t.OutgoingMessageAccentColor
}
