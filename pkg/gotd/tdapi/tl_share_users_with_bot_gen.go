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

// ShareUsersWithBotRequest represents TL type `shareUsersWithBot#58448550`.
type ShareUsersWithBotRequest struct {
	// Identifier of the chat with the bot
	ChatID int64
	// Identifier of the message with the button
	MessageID int64
	// Identifier of the button
	ButtonID int32
	// Identifiers of the shared users
	SharedUserIDs []int64
	// Pass true to check that the users can be shared by the button instead of actually
	// sharing them
	OnlyCheck bool
}

// ShareUsersWithBotRequestTypeID is TL type id of ShareUsersWithBotRequest.
const ShareUsersWithBotRequestTypeID = 0x58448550

// Ensuring interfaces in compile-time for ShareUsersWithBotRequest.
var (
	_ bin.Encoder     = &ShareUsersWithBotRequest{}
	_ bin.Decoder     = &ShareUsersWithBotRequest{}
	_ bin.BareEncoder = &ShareUsersWithBotRequest{}
	_ bin.BareDecoder = &ShareUsersWithBotRequest{}
)

func (s *ShareUsersWithBotRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.MessageID == 0) {
		return false
	}
	if !(s.ButtonID == 0) {
		return false
	}
	if !(s.SharedUserIDs == nil) {
		return false
	}
	if !(s.OnlyCheck == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ShareUsersWithBotRequest) String() string {
	if s == nil {
		return "ShareUsersWithBotRequest(nil)"
	}
	type Alias ShareUsersWithBotRequest
	return fmt.Sprintf("ShareUsersWithBotRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ShareUsersWithBotRequest) TypeID() uint32 {
	return ShareUsersWithBotRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ShareUsersWithBotRequest) TypeName() string {
	return "shareUsersWithBot"
}

// TypeInfo returns info about TL type.
func (s *ShareUsersWithBotRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "shareUsersWithBot",
		ID:   ShareUsersWithBotRequestTypeID,
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
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "ButtonID",
			SchemaName: "button_id",
		},
		{
			Name:       "SharedUserIDs",
			SchemaName: "shared_user_ids",
		},
		{
			Name:       "OnlyCheck",
			SchemaName: "only_check",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *ShareUsersWithBotRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode shareUsersWithBot#58448550 as nil")
	}
	b.PutID(ShareUsersWithBotRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *ShareUsersWithBotRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode shareUsersWithBot#58448550 as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutInt53(s.MessageID)
	b.PutInt32(s.ButtonID)
	b.PutInt(len(s.SharedUserIDs))
	for _, v := range s.SharedUserIDs {
		b.PutInt53(v)
	}
	b.PutBool(s.OnlyCheck)
	return nil
}

// Decode implements bin.Decoder.
func (s *ShareUsersWithBotRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode shareUsersWithBot#58448550 to nil")
	}
	if err := b.ConsumeID(ShareUsersWithBotRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode shareUsersWithBot#58448550: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *ShareUsersWithBotRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode shareUsersWithBot#58448550 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode shareUsersWithBot#58448550: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode shareUsersWithBot#58448550: field message_id: %w", err)
		}
		s.MessageID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode shareUsersWithBot#58448550: field button_id: %w", err)
		}
		s.ButtonID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode shareUsersWithBot#58448550: field shared_user_ids: %w", err)
		}

		if headerLen > 0 {
			s.SharedUserIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode shareUsersWithBot#58448550: field shared_user_ids: %w", err)
			}
			s.SharedUserIDs = append(s.SharedUserIDs, value)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode shareUsersWithBot#58448550: field only_check: %w", err)
		}
		s.OnlyCheck = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *ShareUsersWithBotRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode shareUsersWithBot#58448550 as nil")
	}
	b.ObjStart()
	b.PutID("shareUsersWithBot")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(s.MessageID)
	b.Comma()
	b.FieldStart("button_id")
	b.PutInt32(s.ButtonID)
	b.Comma()
	b.FieldStart("shared_user_ids")
	b.ArrStart()
	for _, v := range s.SharedUserIDs {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("only_check")
	b.PutBool(s.OnlyCheck)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *ShareUsersWithBotRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode shareUsersWithBot#58448550 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("shareUsersWithBot"); err != nil {
				return fmt.Errorf("unable to decode shareUsersWithBot#58448550: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode shareUsersWithBot#58448550: field chat_id: %w", err)
			}
			s.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode shareUsersWithBot#58448550: field message_id: %w", err)
			}
			s.MessageID = value
		case "button_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode shareUsersWithBot#58448550: field button_id: %w", err)
			}
			s.ButtonID = value
		case "shared_user_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode shareUsersWithBot#58448550: field shared_user_ids: %w", err)
				}
				s.SharedUserIDs = append(s.SharedUserIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode shareUsersWithBot#58448550: field shared_user_ids: %w", err)
			}
		case "only_check":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode shareUsersWithBot#58448550: field only_check: %w", err)
			}
			s.OnlyCheck = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *ShareUsersWithBotRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetMessageID returns value of MessageID field.
func (s *ShareUsersWithBotRequest) GetMessageID() (value int64) {
	if s == nil {
		return
	}
	return s.MessageID
}

// GetButtonID returns value of ButtonID field.
func (s *ShareUsersWithBotRequest) GetButtonID() (value int32) {
	if s == nil {
		return
	}
	return s.ButtonID
}

// GetSharedUserIDs returns value of SharedUserIDs field.
func (s *ShareUsersWithBotRequest) GetSharedUserIDs() (value []int64) {
	if s == nil {
		return
	}
	return s.SharedUserIDs
}

// GetOnlyCheck returns value of OnlyCheck field.
func (s *ShareUsersWithBotRequest) GetOnlyCheck() (value bool) {
	if s == nil {
		return
	}
	return s.OnlyCheck
}

// ShareUsersWithBot invokes method shareUsersWithBot#58448550 returning error if any.
func (c *Client) ShareUsersWithBot(ctx context.Context, request *ShareUsersWithBotRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
