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

// SetChatThemeRequest represents TL type `setChatTheme#8f090293`.
type SetChatThemeRequest struct {
	// Chat identifier
	ChatID int64
	// Name of the new chat theme; pass an empty string to return the default theme
	ThemeName string
}

// SetChatThemeRequestTypeID is TL type id of SetChatThemeRequest.
const SetChatThemeRequestTypeID = 0x8f090293

// Ensuring interfaces in compile-time for SetChatThemeRequest.
var (
	_ bin.Encoder     = &SetChatThemeRequest{}
	_ bin.Decoder     = &SetChatThemeRequest{}
	_ bin.BareEncoder = &SetChatThemeRequest{}
	_ bin.BareDecoder = &SetChatThemeRequest{}
)

func (s *SetChatThemeRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.ThemeName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetChatThemeRequest) String() string {
	if s == nil {
		return "SetChatThemeRequest(nil)"
	}
	type Alias SetChatThemeRequest
	return fmt.Sprintf("SetChatThemeRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetChatThemeRequest) TypeID() uint32 {
	return SetChatThemeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetChatThemeRequest) TypeName() string {
	return "setChatTheme"
}

// TypeInfo returns info about TL type.
func (s *SetChatThemeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setChatTheme",
		ID:   SetChatThemeRequestTypeID,
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
			Name:       "ThemeName",
			SchemaName: "theme_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetChatThemeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatTheme#8f090293 as nil")
	}
	b.PutID(SetChatThemeRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetChatThemeRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatTheme#8f090293 as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutString(s.ThemeName)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetChatThemeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatTheme#8f090293 to nil")
	}
	if err := b.ConsumeID(SetChatThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setChatTheme#8f090293: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetChatThemeRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatTheme#8f090293 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setChatTheme#8f090293: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setChatTheme#8f090293: field theme_name: %w", err)
		}
		s.ThemeName = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetChatThemeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatTheme#8f090293 as nil")
	}
	b.ObjStart()
	b.PutID("setChatTheme")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("theme_name")
	b.PutString(s.ThemeName)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetChatThemeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatTheme#8f090293 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setChatTheme"); err != nil {
				return fmt.Errorf("unable to decode setChatTheme#8f090293: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setChatTheme#8f090293: field chat_id: %w", err)
			}
			s.ChatID = value
		case "theme_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setChatTheme#8f090293: field theme_name: %w", err)
			}
			s.ThemeName = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetChatThemeRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetThemeName returns value of ThemeName field.
func (s *SetChatThemeRequest) GetThemeName() (value string) {
	if s == nil {
		return
	}
	return s.ThemeName
}

// SetChatTheme invokes method setChatTheme#8f090293 returning error if any.
func (c *Client) SetChatTheme(ctx context.Context, request *SetChatThemeRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
