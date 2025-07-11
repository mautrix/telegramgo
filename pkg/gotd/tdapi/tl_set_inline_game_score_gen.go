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

// SetInlineGameScoreRequest represents TL type `setInlineGameScore#c7715d8b`.
type SetInlineGameScoreRequest struct {
	// Inline message identifier
	InlineMessageID string
	// Pass true to edit the game message to include the current scoreboard
	EditMessage bool
	// User identifier
	UserID int64
	// The new score
	Score int32
	// Pass true to update the score even if it decreases. If the score is 0, the user will
	// be deleted from the high score table
	Force bool
}

// SetInlineGameScoreRequestTypeID is TL type id of SetInlineGameScoreRequest.
const SetInlineGameScoreRequestTypeID = 0xc7715d8b

// Ensuring interfaces in compile-time for SetInlineGameScoreRequest.
var (
	_ bin.Encoder     = &SetInlineGameScoreRequest{}
	_ bin.Decoder     = &SetInlineGameScoreRequest{}
	_ bin.BareEncoder = &SetInlineGameScoreRequest{}
	_ bin.BareDecoder = &SetInlineGameScoreRequest{}
)

func (s *SetInlineGameScoreRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.InlineMessageID == "") {
		return false
	}
	if !(s.EditMessage == false) {
		return false
	}
	if !(s.UserID == 0) {
		return false
	}
	if !(s.Score == 0) {
		return false
	}
	if !(s.Force == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetInlineGameScoreRequest) String() string {
	if s == nil {
		return "SetInlineGameScoreRequest(nil)"
	}
	type Alias SetInlineGameScoreRequest
	return fmt.Sprintf("SetInlineGameScoreRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetInlineGameScoreRequest) TypeID() uint32 {
	return SetInlineGameScoreRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetInlineGameScoreRequest) TypeName() string {
	return "setInlineGameScore"
}

// TypeInfo returns info about TL type.
func (s *SetInlineGameScoreRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setInlineGameScore",
		ID:   SetInlineGameScoreRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InlineMessageID",
			SchemaName: "inline_message_id",
		},
		{
			Name:       "EditMessage",
			SchemaName: "edit_message",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Score",
			SchemaName: "score",
		},
		{
			Name:       "Force",
			SchemaName: "force",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetInlineGameScoreRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setInlineGameScore#c7715d8b as nil")
	}
	b.PutID(SetInlineGameScoreRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetInlineGameScoreRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setInlineGameScore#c7715d8b as nil")
	}
	b.PutString(s.InlineMessageID)
	b.PutBool(s.EditMessage)
	b.PutInt53(s.UserID)
	b.PutInt32(s.Score)
	b.PutBool(s.Force)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetInlineGameScoreRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setInlineGameScore#c7715d8b to nil")
	}
	if err := b.ConsumeID(SetInlineGameScoreRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setInlineGameScore#c7715d8b: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetInlineGameScoreRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setInlineGameScore#c7715d8b to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setInlineGameScore#c7715d8b: field inline_message_id: %w", err)
		}
		s.InlineMessageID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode setInlineGameScore#c7715d8b: field edit_message: %w", err)
		}
		s.EditMessage = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setInlineGameScore#c7715d8b: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode setInlineGameScore#c7715d8b: field score: %w", err)
		}
		s.Score = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode setInlineGameScore#c7715d8b: field force: %w", err)
		}
		s.Force = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetInlineGameScoreRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setInlineGameScore#c7715d8b as nil")
	}
	b.ObjStart()
	b.PutID("setInlineGameScore")
	b.Comma()
	b.FieldStart("inline_message_id")
	b.PutString(s.InlineMessageID)
	b.Comma()
	b.FieldStart("edit_message")
	b.PutBool(s.EditMessage)
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(s.UserID)
	b.Comma()
	b.FieldStart("score")
	b.PutInt32(s.Score)
	b.Comma()
	b.FieldStart("force")
	b.PutBool(s.Force)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetInlineGameScoreRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setInlineGameScore#c7715d8b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setInlineGameScore"); err != nil {
				return fmt.Errorf("unable to decode setInlineGameScore#c7715d8b: %w", err)
			}
		case "inline_message_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setInlineGameScore#c7715d8b: field inline_message_id: %w", err)
			}
			s.InlineMessageID = value
		case "edit_message":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode setInlineGameScore#c7715d8b: field edit_message: %w", err)
			}
			s.EditMessage = value
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setInlineGameScore#c7715d8b: field user_id: %w", err)
			}
			s.UserID = value
		case "score":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode setInlineGameScore#c7715d8b: field score: %w", err)
			}
			s.Score = value
		case "force":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode setInlineGameScore#c7715d8b: field force: %w", err)
			}
			s.Force = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInlineMessageID returns value of InlineMessageID field.
func (s *SetInlineGameScoreRequest) GetInlineMessageID() (value string) {
	if s == nil {
		return
	}
	return s.InlineMessageID
}

// GetEditMessage returns value of EditMessage field.
func (s *SetInlineGameScoreRequest) GetEditMessage() (value bool) {
	if s == nil {
		return
	}
	return s.EditMessage
}

// GetUserID returns value of UserID field.
func (s *SetInlineGameScoreRequest) GetUserID() (value int64) {
	if s == nil {
		return
	}
	return s.UserID
}

// GetScore returns value of Score field.
func (s *SetInlineGameScoreRequest) GetScore() (value int32) {
	if s == nil {
		return
	}
	return s.Score
}

// GetForce returns value of Force field.
func (s *SetInlineGameScoreRequest) GetForce() (value bool) {
	if s == nil {
		return
	}
	return s.Force
}

// SetInlineGameScore invokes method setInlineGameScore#c7715d8b returning error if any.
func (c *Client) SetInlineGameScore(ctx context.Context, request *SetInlineGameScoreRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
