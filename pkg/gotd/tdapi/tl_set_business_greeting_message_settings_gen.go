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

// SetBusinessGreetingMessageSettingsRequest represents TL type `setBusinessGreetingMessageSettings#cbf53c3d`.
type SetBusinessGreetingMessageSettingsRequest struct {
	// The new settings for the greeting message of the business; pass null to disable the
	// greeting message
	GreetingMessageSettings BusinessGreetingMessageSettings
}

// SetBusinessGreetingMessageSettingsRequestTypeID is TL type id of SetBusinessGreetingMessageSettingsRequest.
const SetBusinessGreetingMessageSettingsRequestTypeID = 0xcbf53c3d

// Ensuring interfaces in compile-time for SetBusinessGreetingMessageSettingsRequest.
var (
	_ bin.Encoder     = &SetBusinessGreetingMessageSettingsRequest{}
	_ bin.Decoder     = &SetBusinessGreetingMessageSettingsRequest{}
	_ bin.BareEncoder = &SetBusinessGreetingMessageSettingsRequest{}
	_ bin.BareDecoder = &SetBusinessGreetingMessageSettingsRequest{}
)

func (s *SetBusinessGreetingMessageSettingsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.GreetingMessageSettings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetBusinessGreetingMessageSettingsRequest) String() string {
	if s == nil {
		return "SetBusinessGreetingMessageSettingsRequest(nil)"
	}
	type Alias SetBusinessGreetingMessageSettingsRequest
	return fmt.Sprintf("SetBusinessGreetingMessageSettingsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetBusinessGreetingMessageSettingsRequest) TypeID() uint32 {
	return SetBusinessGreetingMessageSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetBusinessGreetingMessageSettingsRequest) TypeName() string {
	return "setBusinessGreetingMessageSettings"
}

// TypeInfo returns info about TL type.
func (s *SetBusinessGreetingMessageSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setBusinessGreetingMessageSettings",
		ID:   SetBusinessGreetingMessageSettingsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GreetingMessageSettings",
			SchemaName: "greeting_message_settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetBusinessGreetingMessageSettingsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBusinessGreetingMessageSettings#cbf53c3d as nil")
	}
	b.PutID(SetBusinessGreetingMessageSettingsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetBusinessGreetingMessageSettingsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBusinessGreetingMessageSettings#cbf53c3d as nil")
	}
	if err := s.GreetingMessageSettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setBusinessGreetingMessageSettings#cbf53c3d: field greeting_message_settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetBusinessGreetingMessageSettingsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBusinessGreetingMessageSettings#cbf53c3d to nil")
	}
	if err := b.ConsumeID(SetBusinessGreetingMessageSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setBusinessGreetingMessageSettings#cbf53c3d: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetBusinessGreetingMessageSettingsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBusinessGreetingMessageSettings#cbf53c3d to nil")
	}
	{
		if err := s.GreetingMessageSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setBusinessGreetingMessageSettings#cbf53c3d: field greeting_message_settings: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetBusinessGreetingMessageSettingsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setBusinessGreetingMessageSettings#cbf53c3d as nil")
	}
	b.ObjStart()
	b.PutID("setBusinessGreetingMessageSettings")
	b.Comma()
	b.FieldStart("greeting_message_settings")
	if err := s.GreetingMessageSettings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setBusinessGreetingMessageSettings#cbf53c3d: field greeting_message_settings: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetBusinessGreetingMessageSettingsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setBusinessGreetingMessageSettings#cbf53c3d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setBusinessGreetingMessageSettings"); err != nil {
				return fmt.Errorf("unable to decode setBusinessGreetingMessageSettings#cbf53c3d: %w", err)
			}
		case "greeting_message_settings":
			if err := s.GreetingMessageSettings.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode setBusinessGreetingMessageSettings#cbf53c3d: field greeting_message_settings: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGreetingMessageSettings returns value of GreetingMessageSettings field.
func (s *SetBusinessGreetingMessageSettingsRequest) GetGreetingMessageSettings() (value BusinessGreetingMessageSettings) {
	if s == nil {
		return
	}
	return s.GreetingMessageSettings
}

// SetBusinessGreetingMessageSettings invokes method setBusinessGreetingMessageSettings#cbf53c3d returning error if any.
func (c *Client) SetBusinessGreetingMessageSettings(ctx context.Context, greetingmessagesettings BusinessGreetingMessageSettings) error {
	var ok Ok

	request := &SetBusinessGreetingMessageSettingsRequest{
		GreetingMessageSettings: greetingmessagesettings,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
