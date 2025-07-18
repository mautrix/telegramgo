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

// UserSupportInfo represents TL type `userSupportInfo#b50e1c29`.
type UserSupportInfo struct {
	// Information message
	Message FormattedText
	// Information author
	Author string
	// Information change date
	Date int32
}

// UserSupportInfoTypeID is TL type id of UserSupportInfo.
const UserSupportInfoTypeID = 0xb50e1c29

// Ensuring interfaces in compile-time for UserSupportInfo.
var (
	_ bin.Encoder     = &UserSupportInfo{}
	_ bin.Decoder     = &UserSupportInfo{}
	_ bin.BareEncoder = &UserSupportInfo{}
	_ bin.BareDecoder = &UserSupportInfo{}
)

func (u *UserSupportInfo) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Message.Zero()) {
		return false
	}
	if !(u.Author == "") {
		return false
	}
	if !(u.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserSupportInfo) String() string {
	if u == nil {
		return "UserSupportInfo(nil)"
	}
	type Alias UserSupportInfo
	return fmt.Sprintf("UserSupportInfo%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserSupportInfo) TypeID() uint32 {
	return UserSupportInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*UserSupportInfo) TypeName() string {
	return "userSupportInfo"
}

// TypeInfo returns info about TL type.
func (u *UserSupportInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userSupportInfo",
		ID:   UserSupportInfoTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Author",
			SchemaName: "author",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UserSupportInfo) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userSupportInfo#b50e1c29 as nil")
	}
	b.PutID(UserSupportInfoTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserSupportInfo) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userSupportInfo#b50e1c29 as nil")
	}
	if err := u.Message.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userSupportInfo#b50e1c29: field message: %w", err)
	}
	b.PutString(u.Author)
	b.PutInt32(u.Date)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserSupportInfo) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userSupportInfo#b50e1c29 to nil")
	}
	if err := b.ConsumeID(UserSupportInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode userSupportInfo#b50e1c29: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserSupportInfo) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userSupportInfo#b50e1c29 to nil")
	}
	{
		if err := u.Message.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userSupportInfo#b50e1c29: field message: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode userSupportInfo#b50e1c29: field author: %w", err)
		}
		u.Author = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userSupportInfo#b50e1c29: field date: %w", err)
		}
		u.Date = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *UserSupportInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode userSupportInfo#b50e1c29 as nil")
	}
	b.ObjStart()
	b.PutID("userSupportInfo")
	b.Comma()
	b.FieldStart("message")
	if err := u.Message.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode userSupportInfo#b50e1c29: field message: %w", err)
	}
	b.Comma()
	b.FieldStart("author")
	b.PutString(u.Author)
	b.Comma()
	b.FieldStart("date")
	b.PutInt32(u.Date)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *UserSupportInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode userSupportInfo#b50e1c29 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("userSupportInfo"); err != nil {
				return fmt.Errorf("unable to decode userSupportInfo#b50e1c29: %w", err)
			}
		case "message":
			if err := u.Message.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode userSupportInfo#b50e1c29: field message: %w", err)
			}
		case "author":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode userSupportInfo#b50e1c29: field author: %w", err)
			}
			u.Author = value
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode userSupportInfo#b50e1c29: field date: %w", err)
			}
			u.Date = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetMessage returns value of Message field.
func (u *UserSupportInfo) GetMessage() (value FormattedText) {
	if u == nil {
		return
	}
	return u.Message
}

// GetAuthor returns value of Author field.
func (u *UserSupportInfo) GetAuthor() (value string) {
	if u == nil {
		return
	}
	return u.Author
}

// GetDate returns value of Date field.
func (u *UserSupportInfo) GetDate() (value int32) {
	if u == nil {
		return
	}
	return u.Date
}
