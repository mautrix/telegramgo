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

// SetBotInfoDescriptionRequest represents TL type `setBotInfoDescription#29571d48`.
type SetBotInfoDescriptionRequest struct {
	// Identifier of the target bot
	BotUserID int64
	// A two-letter ISO 639-1 language code. If empty, the description will be shown to all
	// users for whose languages there is no dedicated description
	LanguageCode string
	// Sets the text shown in the chat with a bot if the chat is empty. Can be called only if
	// userTypeBot.can_be_edited == true
	Description string
}

// SetBotInfoDescriptionRequestTypeID is TL type id of SetBotInfoDescriptionRequest.
const SetBotInfoDescriptionRequestTypeID = 0x29571d48

// Ensuring interfaces in compile-time for SetBotInfoDescriptionRequest.
var (
	_ bin.Encoder     = &SetBotInfoDescriptionRequest{}
	_ bin.Decoder     = &SetBotInfoDescriptionRequest{}
	_ bin.BareEncoder = &SetBotInfoDescriptionRequest{}
	_ bin.BareDecoder = &SetBotInfoDescriptionRequest{}
)

func (s *SetBotInfoDescriptionRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.BotUserID == 0) {
		return false
	}
	if !(s.LanguageCode == "") {
		return false
	}
	if !(s.Description == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetBotInfoDescriptionRequest) String() string {
	if s == nil {
		return "SetBotInfoDescriptionRequest(nil)"
	}
	type Alias SetBotInfoDescriptionRequest
	return fmt.Sprintf("SetBotInfoDescriptionRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetBotInfoDescriptionRequest) TypeID() uint32 {
	return SetBotInfoDescriptionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetBotInfoDescriptionRequest) TypeName() string {
	return "setBotInfoDescription"
}

// TypeInfo returns info about TL type.
func (s *SetBotInfoDescriptionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setBotInfoDescription",
		ID:   SetBotInfoDescriptionRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "LanguageCode",
			SchemaName: "language_code",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetBotInfoDescriptionRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBotInfoDescription#29571d48 as nil")
	}
	b.PutID(SetBotInfoDescriptionRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetBotInfoDescriptionRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBotInfoDescription#29571d48 as nil")
	}
	b.PutInt53(s.BotUserID)
	b.PutString(s.LanguageCode)
	b.PutString(s.Description)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetBotInfoDescriptionRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBotInfoDescription#29571d48 to nil")
	}
	if err := b.ConsumeID(SetBotInfoDescriptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setBotInfoDescription#29571d48: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetBotInfoDescriptionRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBotInfoDescription#29571d48 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setBotInfoDescription#29571d48: field bot_user_id: %w", err)
		}
		s.BotUserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setBotInfoDescription#29571d48: field language_code: %w", err)
		}
		s.LanguageCode = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setBotInfoDescription#29571d48: field description: %w", err)
		}
		s.Description = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetBotInfoDescriptionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setBotInfoDescription#29571d48 as nil")
	}
	b.ObjStart()
	b.PutID("setBotInfoDescription")
	b.Comma()
	b.FieldStart("bot_user_id")
	b.PutInt53(s.BotUserID)
	b.Comma()
	b.FieldStart("language_code")
	b.PutString(s.LanguageCode)
	b.Comma()
	b.FieldStart("description")
	b.PutString(s.Description)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetBotInfoDescriptionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setBotInfoDescription#29571d48 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setBotInfoDescription"); err != nil {
				return fmt.Errorf("unable to decode setBotInfoDescription#29571d48: %w", err)
			}
		case "bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setBotInfoDescription#29571d48: field bot_user_id: %w", err)
			}
			s.BotUserID = value
		case "language_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setBotInfoDescription#29571d48: field language_code: %w", err)
			}
			s.LanguageCode = value
		case "description":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setBotInfoDescription#29571d48: field description: %w", err)
			}
			s.Description = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBotUserID returns value of BotUserID field.
func (s *SetBotInfoDescriptionRequest) GetBotUserID() (value int64) {
	if s == nil {
		return
	}
	return s.BotUserID
}

// GetLanguageCode returns value of LanguageCode field.
func (s *SetBotInfoDescriptionRequest) GetLanguageCode() (value string) {
	if s == nil {
		return
	}
	return s.LanguageCode
}

// GetDescription returns value of Description field.
func (s *SetBotInfoDescriptionRequest) GetDescription() (value string) {
	if s == nil {
		return
	}
	return s.Description
}

// SetBotInfoDescription invokes method setBotInfoDescription#29571d48 returning error if any.
func (c *Client) SetBotInfoDescription(ctx context.Context, request *SetBotInfoDescriptionRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
