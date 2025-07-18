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

// Users represents TL type `users#9d955a12`.
type Users struct {
	// Approximate total number of users found
	TotalCount int32
	// A list of user identifiers
	UserIDs []int64
}

// UsersTypeID is TL type id of Users.
const UsersTypeID = 0x9d955a12

// Ensuring interfaces in compile-time for Users.
var (
	_ bin.Encoder     = &Users{}
	_ bin.Decoder     = &Users{}
	_ bin.BareEncoder = &Users{}
	_ bin.BareDecoder = &Users{}
)

func (u *Users) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.TotalCount == 0) {
		return false
	}
	if !(u.UserIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *Users) String() string {
	if u == nil {
		return "Users(nil)"
	}
	type Alias Users
	return fmt.Sprintf("Users%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Users) TypeID() uint32 {
	return UsersTypeID
}

// TypeName returns name of type in TL schema.
func (*Users) TypeName() string {
	return "users"
}

// TypeInfo returns info about TL type.
func (u *Users) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "users",
		ID:   UsersTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "UserIDs",
			SchemaName: "user_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *Users) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode users#9d955a12 as nil")
	}
	b.PutID(UsersTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *Users) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode users#9d955a12 as nil")
	}
	b.PutInt32(u.TotalCount)
	b.PutInt(len(u.UserIDs))
	for _, v := range u.UserIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *Users) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode users#9d955a12 to nil")
	}
	if err := b.ConsumeID(UsersTypeID); err != nil {
		return fmt.Errorf("unable to decode users#9d955a12: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *Users) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode users#9d955a12 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode users#9d955a12: field total_count: %w", err)
		}
		u.TotalCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode users#9d955a12: field user_ids: %w", err)
		}

		if headerLen > 0 {
			u.UserIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode users#9d955a12: field user_ids: %w", err)
			}
			u.UserIDs = append(u.UserIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *Users) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode users#9d955a12 as nil")
	}
	b.ObjStart()
	b.PutID("users")
	b.Comma()
	b.FieldStart("total_count")
	b.PutInt32(u.TotalCount)
	b.Comma()
	b.FieldStart("user_ids")
	b.ArrStart()
	for _, v := range u.UserIDs {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *Users) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode users#9d955a12 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("users"); err != nil {
				return fmt.Errorf("unable to decode users#9d955a12: %w", err)
			}
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode users#9d955a12: field total_count: %w", err)
			}
			u.TotalCount = value
		case "user_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode users#9d955a12: field user_ids: %w", err)
				}
				u.UserIDs = append(u.UserIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode users#9d955a12: field user_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTotalCount returns value of TotalCount field.
func (u *Users) GetTotalCount() (value int32) {
	if u == nil {
		return
	}
	return u.TotalCount
}

// GetUserIDs returns value of UserIDs field.
func (u *Users) GetUserIDs() (value []int64) {
	if u == nil {
		return
	}
	return u.UserIDs
}
