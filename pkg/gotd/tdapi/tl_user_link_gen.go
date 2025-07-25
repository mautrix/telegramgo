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

// UserLink represents TL type `userLink#1db0fef8`.
type UserLink struct {
	// The URL
	URL string
	// Left time for which the link is valid, in seconds; 0 if the link is a public username
	// link
	ExpiresIn int32
}

// UserLinkTypeID is TL type id of UserLink.
const UserLinkTypeID = 0x1db0fef8

// Ensuring interfaces in compile-time for UserLink.
var (
	_ bin.Encoder     = &UserLink{}
	_ bin.Decoder     = &UserLink{}
	_ bin.BareEncoder = &UserLink{}
	_ bin.BareDecoder = &UserLink{}
)

func (u *UserLink) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.URL == "") {
		return false
	}
	if !(u.ExpiresIn == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserLink) String() string {
	if u == nil {
		return "UserLink(nil)"
	}
	type Alias UserLink
	return fmt.Sprintf("UserLink%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserLink) TypeID() uint32 {
	return UserLinkTypeID
}

// TypeName returns name of type in TL schema.
func (*UserLink) TypeName() string {
	return "userLink"
}

// TypeInfo returns info about TL type.
func (u *UserLink) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userLink",
		ID:   UserLinkTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "ExpiresIn",
			SchemaName: "expires_in",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UserLink) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userLink#1db0fef8 as nil")
	}
	b.PutID(UserLinkTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserLink) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userLink#1db0fef8 as nil")
	}
	b.PutString(u.URL)
	b.PutInt32(u.ExpiresIn)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserLink) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userLink#1db0fef8 to nil")
	}
	if err := b.ConsumeID(UserLinkTypeID); err != nil {
		return fmt.Errorf("unable to decode userLink#1db0fef8: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserLink) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userLink#1db0fef8 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode userLink#1db0fef8: field url: %w", err)
		}
		u.URL = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userLink#1db0fef8: field expires_in: %w", err)
		}
		u.ExpiresIn = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *UserLink) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode userLink#1db0fef8 as nil")
	}
	b.ObjStart()
	b.PutID("userLink")
	b.Comma()
	b.FieldStart("url")
	b.PutString(u.URL)
	b.Comma()
	b.FieldStart("expires_in")
	b.PutInt32(u.ExpiresIn)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *UserLink) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode userLink#1db0fef8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("userLink"); err != nil {
				return fmt.Errorf("unable to decode userLink#1db0fef8: %w", err)
			}
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode userLink#1db0fef8: field url: %w", err)
			}
			u.URL = value
		case "expires_in":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode userLink#1db0fef8: field expires_in: %w", err)
			}
			u.ExpiresIn = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetURL returns value of URL field.
func (u *UserLink) GetURL() (value string) {
	if u == nil {
		return
	}
	return u.URL
}

// GetExpiresIn returns value of ExpiresIn field.
func (u *UserLink) GetExpiresIn() (value int32) {
	if u == nil {
		return
	}
	return u.ExpiresIn
}
