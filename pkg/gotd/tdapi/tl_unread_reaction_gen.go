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

// UnreadReaction represents TL type `unreadReaction#8c5b3b82`.
type UnreadReaction struct {
	// Type of the reaction
	Type ReactionTypeClass
	// Identifier of the sender, added the reaction
	SenderID MessageSenderClass
	// True, if the reaction was added with a big animation
	IsBig bool
}

// UnreadReactionTypeID is TL type id of UnreadReaction.
const UnreadReactionTypeID = 0x8c5b3b82

// Ensuring interfaces in compile-time for UnreadReaction.
var (
	_ bin.Encoder     = &UnreadReaction{}
	_ bin.Decoder     = &UnreadReaction{}
	_ bin.BareEncoder = &UnreadReaction{}
	_ bin.BareDecoder = &UnreadReaction{}
)

func (u *UnreadReaction) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Type == nil) {
		return false
	}
	if !(u.SenderID == nil) {
		return false
	}
	if !(u.IsBig == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UnreadReaction) String() string {
	if u == nil {
		return "UnreadReaction(nil)"
	}
	type Alias UnreadReaction
	return fmt.Sprintf("UnreadReaction%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UnreadReaction) TypeID() uint32 {
	return UnreadReactionTypeID
}

// TypeName returns name of type in TL schema.
func (*UnreadReaction) TypeName() string {
	return "unreadReaction"
}

// TypeInfo returns info about TL type.
func (u *UnreadReaction) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "unreadReaction",
		ID:   UnreadReactionTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "SenderID",
			SchemaName: "sender_id",
		},
		{
			Name:       "IsBig",
			SchemaName: "is_big",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UnreadReaction) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode unreadReaction#8c5b3b82 as nil")
	}
	b.PutID(UnreadReactionTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UnreadReaction) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode unreadReaction#8c5b3b82 as nil")
	}
	if u.Type == nil {
		return fmt.Errorf("unable to encode unreadReaction#8c5b3b82: field type is nil")
	}
	if err := u.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode unreadReaction#8c5b3b82: field type: %w", err)
	}
	if u.SenderID == nil {
		return fmt.Errorf("unable to encode unreadReaction#8c5b3b82: field sender_id is nil")
	}
	if err := u.SenderID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode unreadReaction#8c5b3b82: field sender_id: %w", err)
	}
	b.PutBool(u.IsBig)
	return nil
}

// Decode implements bin.Decoder.
func (u *UnreadReaction) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode unreadReaction#8c5b3b82 to nil")
	}
	if err := b.ConsumeID(UnreadReactionTypeID); err != nil {
		return fmt.Errorf("unable to decode unreadReaction#8c5b3b82: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UnreadReaction) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode unreadReaction#8c5b3b82 to nil")
	}
	{
		value, err := DecodeReactionType(b)
		if err != nil {
			return fmt.Errorf("unable to decode unreadReaction#8c5b3b82: field type: %w", err)
		}
		u.Type = value
	}
	{
		value, err := DecodeMessageSender(b)
		if err != nil {
			return fmt.Errorf("unable to decode unreadReaction#8c5b3b82: field sender_id: %w", err)
		}
		u.SenderID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode unreadReaction#8c5b3b82: field is_big: %w", err)
		}
		u.IsBig = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *UnreadReaction) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode unreadReaction#8c5b3b82 as nil")
	}
	b.ObjStart()
	b.PutID("unreadReaction")
	b.Comma()
	b.FieldStart("type")
	if u.Type == nil {
		return fmt.Errorf("unable to encode unreadReaction#8c5b3b82: field type is nil")
	}
	if err := u.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode unreadReaction#8c5b3b82: field type: %w", err)
	}
	b.Comma()
	b.FieldStart("sender_id")
	if u.SenderID == nil {
		return fmt.Errorf("unable to encode unreadReaction#8c5b3b82: field sender_id is nil")
	}
	if err := u.SenderID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode unreadReaction#8c5b3b82: field sender_id: %w", err)
	}
	b.Comma()
	b.FieldStart("is_big")
	b.PutBool(u.IsBig)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *UnreadReaction) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode unreadReaction#8c5b3b82 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("unreadReaction"); err != nil {
				return fmt.Errorf("unable to decode unreadReaction#8c5b3b82: %w", err)
			}
		case "type":
			value, err := DecodeTDLibJSONReactionType(b)
			if err != nil {
				return fmt.Errorf("unable to decode unreadReaction#8c5b3b82: field type: %w", err)
			}
			u.Type = value
		case "sender_id":
			value, err := DecodeTDLibJSONMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode unreadReaction#8c5b3b82: field sender_id: %w", err)
			}
			u.SenderID = value
		case "is_big":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode unreadReaction#8c5b3b82: field is_big: %w", err)
			}
			u.IsBig = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetType returns value of Type field.
func (u *UnreadReaction) GetType() (value ReactionTypeClass) {
	if u == nil {
		return
	}
	return u.Type
}

// GetSenderID returns value of SenderID field.
func (u *UnreadReaction) GetSenderID() (value MessageSenderClass) {
	if u == nil {
		return
	}
	return u.SenderID
}

// GetIsBig returns value of IsBig field.
func (u *UnreadReaction) GetIsBig() (value bool) {
	if u == nil {
		return
	}
	return u.IsBig
}
