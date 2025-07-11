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

// SetDefaultBackgroundRequest represents TL type `setDefaultBackground#89d1a8a1`.
type SetDefaultBackgroundRequest struct {
	// The input background to use; pass null to create a new filled background
	Background InputBackgroundClass
	// Background type; pass null to use the default type of the remote background;
	// backgroundTypeChatTheme isn't supported
	Type BackgroundTypeClass
	// Pass true if the background is set for a dark theme
	ForDarkTheme bool
}

// SetDefaultBackgroundRequestTypeID is TL type id of SetDefaultBackgroundRequest.
const SetDefaultBackgroundRequestTypeID = 0x89d1a8a1

// Ensuring interfaces in compile-time for SetDefaultBackgroundRequest.
var (
	_ bin.Encoder     = &SetDefaultBackgroundRequest{}
	_ bin.Decoder     = &SetDefaultBackgroundRequest{}
	_ bin.BareEncoder = &SetDefaultBackgroundRequest{}
	_ bin.BareDecoder = &SetDefaultBackgroundRequest{}
)

func (s *SetDefaultBackgroundRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Background == nil) {
		return false
	}
	if !(s.Type == nil) {
		return false
	}
	if !(s.ForDarkTheme == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetDefaultBackgroundRequest) String() string {
	if s == nil {
		return "SetDefaultBackgroundRequest(nil)"
	}
	type Alias SetDefaultBackgroundRequest
	return fmt.Sprintf("SetDefaultBackgroundRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetDefaultBackgroundRequest) TypeID() uint32 {
	return SetDefaultBackgroundRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetDefaultBackgroundRequest) TypeName() string {
	return "setDefaultBackground"
}

// TypeInfo returns info about TL type.
func (s *SetDefaultBackgroundRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setDefaultBackground",
		ID:   SetDefaultBackgroundRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Background",
			SchemaName: "background",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "ForDarkTheme",
			SchemaName: "for_dark_theme",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetDefaultBackgroundRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setDefaultBackground#89d1a8a1 as nil")
	}
	b.PutID(SetDefaultBackgroundRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetDefaultBackgroundRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setDefaultBackground#89d1a8a1 as nil")
	}
	if s.Background == nil {
		return fmt.Errorf("unable to encode setDefaultBackground#89d1a8a1: field background is nil")
	}
	if err := s.Background.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setDefaultBackground#89d1a8a1: field background: %w", err)
	}
	if s.Type == nil {
		return fmt.Errorf("unable to encode setDefaultBackground#89d1a8a1: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setDefaultBackground#89d1a8a1: field type: %w", err)
	}
	b.PutBool(s.ForDarkTheme)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetDefaultBackgroundRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setDefaultBackground#89d1a8a1 to nil")
	}
	if err := b.ConsumeID(SetDefaultBackgroundRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setDefaultBackground#89d1a8a1: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetDefaultBackgroundRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setDefaultBackground#89d1a8a1 to nil")
	}
	{
		value, err := DecodeInputBackground(b)
		if err != nil {
			return fmt.Errorf("unable to decode setDefaultBackground#89d1a8a1: field background: %w", err)
		}
		s.Background = value
	}
	{
		value, err := DecodeBackgroundType(b)
		if err != nil {
			return fmt.Errorf("unable to decode setDefaultBackground#89d1a8a1: field type: %w", err)
		}
		s.Type = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode setDefaultBackground#89d1a8a1: field for_dark_theme: %w", err)
		}
		s.ForDarkTheme = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetDefaultBackgroundRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setDefaultBackground#89d1a8a1 as nil")
	}
	b.ObjStart()
	b.PutID("setDefaultBackground")
	b.Comma()
	b.FieldStart("background")
	if s.Background == nil {
		return fmt.Errorf("unable to encode setDefaultBackground#89d1a8a1: field background is nil")
	}
	if err := s.Background.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setDefaultBackground#89d1a8a1: field background: %w", err)
	}
	b.Comma()
	b.FieldStart("type")
	if s.Type == nil {
		return fmt.Errorf("unable to encode setDefaultBackground#89d1a8a1: field type is nil")
	}
	if err := s.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setDefaultBackground#89d1a8a1: field type: %w", err)
	}
	b.Comma()
	b.FieldStart("for_dark_theme")
	b.PutBool(s.ForDarkTheme)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetDefaultBackgroundRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setDefaultBackground#89d1a8a1 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setDefaultBackground"); err != nil {
				return fmt.Errorf("unable to decode setDefaultBackground#89d1a8a1: %w", err)
			}
		case "background":
			value, err := DecodeTDLibJSONInputBackground(b)
			if err != nil {
				return fmt.Errorf("unable to decode setDefaultBackground#89d1a8a1: field background: %w", err)
			}
			s.Background = value
		case "type":
			value, err := DecodeTDLibJSONBackgroundType(b)
			if err != nil {
				return fmt.Errorf("unable to decode setDefaultBackground#89d1a8a1: field type: %w", err)
			}
			s.Type = value
		case "for_dark_theme":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode setDefaultBackground#89d1a8a1: field for_dark_theme: %w", err)
			}
			s.ForDarkTheme = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBackground returns value of Background field.
func (s *SetDefaultBackgroundRequest) GetBackground() (value InputBackgroundClass) {
	if s == nil {
		return
	}
	return s.Background
}

// GetType returns value of Type field.
func (s *SetDefaultBackgroundRequest) GetType() (value BackgroundTypeClass) {
	if s == nil {
		return
	}
	return s.Type
}

// GetForDarkTheme returns value of ForDarkTheme field.
func (s *SetDefaultBackgroundRequest) GetForDarkTheme() (value bool) {
	if s == nil {
		return
	}
	return s.ForDarkTheme
}

// SetDefaultBackground invokes method setDefaultBackground#89d1a8a1 returning error if any.
func (c *Client) SetDefaultBackground(ctx context.Context, request *SetDefaultBackgroundRequest) (*Background, error) {
	var result Background

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
