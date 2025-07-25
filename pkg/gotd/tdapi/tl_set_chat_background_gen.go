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

// SetChatBackgroundRequest represents TL type `setChatBackground#eb4c3fe`.
type SetChatBackgroundRequest struct {
	// Chat identifier
	ChatID int64
	// The input background to use; pass null to create a new filled or chat theme background
	Background InputBackgroundClass
	// Background type; pass null to use default background type for the chosen background;
	// backgroundTypeChatTheme isn't supported for private and secret chats.
	Type BackgroundTypeClass
	// Dimming of the background in dark themes, as a percentage; 0-100. Applied only to
	// Wallpaper and Fill types of background
	DarkThemeDimming int32
	// Pass true to set background only for self; pass false to set background for all chat
	// users. Always false for backgrounds set in boosted chats. Background can be set for
	// both users only by Telegram Premium users and if set background isn't of the type
	// inputBackgroundPrevious
	OnlyForSelf bool
}

// SetChatBackgroundRequestTypeID is TL type id of SetChatBackgroundRequest.
const SetChatBackgroundRequestTypeID = 0xeb4c3fe

// Ensuring interfaces in compile-time for SetChatBackgroundRequest.
var (
	_ bin.Encoder     = &SetChatBackgroundRequest{}
	_ bin.Decoder     = &SetChatBackgroundRequest{}
	_ bin.BareEncoder = &SetChatBackgroundRequest{}
	_ bin.BareDecoder = &SetChatBackgroundRequest{}
)

func (s *SetChatBackgroundRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.Background == nil) {
		return false
	}
	if !(s.Type == nil) {
		return false
	}
	if !(s.DarkThemeDimming == 0) {
		return false
	}
	if !(s.OnlyForSelf == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetChatBackgroundRequest) String() string {
	if s == nil {
		return "SetChatBackgroundRequest(nil)"
	}
	type Alias SetChatBackgroundRequest
	return fmt.Sprintf("SetChatBackgroundRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetChatBackgroundRequest) TypeID() uint32 {
	return SetChatBackgroundRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetChatBackgroundRequest) TypeName() string {
	return "setChatBackground"
}

// TypeInfo returns info about TL type.
func (s *SetChatBackgroundRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setChatBackground",
		ID:   SetChatBackgroundRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Background",
			SchemaName: "background",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "DarkThemeDimming",
			SchemaName: "dark_theme_dimming",
		},
		{
			Name:       "OnlyForSelf",
			SchemaName: "only_for_self",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetChatBackgroundRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatBackground#eb4c3fe as nil")
	}
	b.PutID(SetChatBackgroundRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetChatBackgroundRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatBackground#eb4c3fe as nil")
	}
	b.PutInt53(s.ChatID)
	if s.Background == nil {
		return fmt.Errorf("unable to encode setChatBackground#eb4c3fe: field background is nil")
	}
	if err := s.Background.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setChatBackground#eb4c3fe: field background: %w", err)
	}
	if s.Type == nil {
		return fmt.Errorf("unable to encode setChatBackground#eb4c3fe: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setChatBackground#eb4c3fe: field type: %w", err)
	}
	b.PutInt32(s.DarkThemeDimming)
	b.PutBool(s.OnlyForSelf)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetChatBackgroundRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatBackground#eb4c3fe to nil")
	}
	if err := b.ConsumeID(SetChatBackgroundRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setChatBackground#eb4c3fe: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetChatBackgroundRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatBackground#eb4c3fe to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setChatBackground#eb4c3fe: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := DecodeInputBackground(b)
		if err != nil {
			return fmt.Errorf("unable to decode setChatBackground#eb4c3fe: field background: %w", err)
		}
		s.Background = value
	}
	{
		value, err := DecodeBackgroundType(b)
		if err != nil {
			return fmt.Errorf("unable to decode setChatBackground#eb4c3fe: field type: %w", err)
		}
		s.Type = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode setChatBackground#eb4c3fe: field dark_theme_dimming: %w", err)
		}
		s.DarkThemeDimming = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode setChatBackground#eb4c3fe: field only_for_self: %w", err)
		}
		s.OnlyForSelf = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetChatBackgroundRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatBackground#eb4c3fe as nil")
	}
	b.ObjStart()
	b.PutID("setChatBackground")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("background")
	if s.Background == nil {
		return fmt.Errorf("unable to encode setChatBackground#eb4c3fe: field background is nil")
	}
	if err := s.Background.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setChatBackground#eb4c3fe: field background: %w", err)
	}
	b.Comma()
	b.FieldStart("type")
	if s.Type == nil {
		return fmt.Errorf("unable to encode setChatBackground#eb4c3fe: field type is nil")
	}
	if err := s.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setChatBackground#eb4c3fe: field type: %w", err)
	}
	b.Comma()
	b.FieldStart("dark_theme_dimming")
	b.PutInt32(s.DarkThemeDimming)
	b.Comma()
	b.FieldStart("only_for_self")
	b.PutBool(s.OnlyForSelf)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetChatBackgroundRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatBackground#eb4c3fe to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setChatBackground"); err != nil {
				return fmt.Errorf("unable to decode setChatBackground#eb4c3fe: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setChatBackground#eb4c3fe: field chat_id: %w", err)
			}
			s.ChatID = value
		case "background":
			value, err := DecodeTDLibJSONInputBackground(b)
			if err != nil {
				return fmt.Errorf("unable to decode setChatBackground#eb4c3fe: field background: %w", err)
			}
			s.Background = value
		case "type":
			value, err := DecodeTDLibJSONBackgroundType(b)
			if err != nil {
				return fmt.Errorf("unable to decode setChatBackground#eb4c3fe: field type: %w", err)
			}
			s.Type = value
		case "dark_theme_dimming":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode setChatBackground#eb4c3fe: field dark_theme_dimming: %w", err)
			}
			s.DarkThemeDimming = value
		case "only_for_self":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode setChatBackground#eb4c3fe: field only_for_self: %w", err)
			}
			s.OnlyForSelf = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetChatBackgroundRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetBackground returns value of Background field.
func (s *SetChatBackgroundRequest) GetBackground() (value InputBackgroundClass) {
	if s == nil {
		return
	}
	return s.Background
}

// GetType returns value of Type field.
func (s *SetChatBackgroundRequest) GetType() (value BackgroundTypeClass) {
	if s == nil {
		return
	}
	return s.Type
}

// GetDarkThemeDimming returns value of DarkThemeDimming field.
func (s *SetChatBackgroundRequest) GetDarkThemeDimming() (value int32) {
	if s == nil {
		return
	}
	return s.DarkThemeDimming
}

// GetOnlyForSelf returns value of OnlyForSelf field.
func (s *SetChatBackgroundRequest) GetOnlyForSelf() (value bool) {
	if s == nil {
		return
	}
	return s.OnlyForSelf
}

// SetChatBackground invokes method setChatBackground#eb4c3fe returning error if any.
func (c *Client) SetChatBackground(ctx context.Context, request *SetChatBackgroundRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
