// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// AccountSaveThemeRequest represents TL type `account.saveTheme#f257106c`.
// Save a theme
//
// See https://core.telegram.org/method/account.saveTheme for reference.
type AccountSaveThemeRequest struct {
	// Theme to save
	Theme InputThemeClass
	// Unsave
	Unsave bool
}

// AccountSaveThemeRequestTypeID is TL type id of AccountSaveThemeRequest.
const AccountSaveThemeRequestTypeID = 0xf257106c

// Ensuring interfaces in compile-time for AccountSaveThemeRequest.
var (
	_ bin.Encoder     = &AccountSaveThemeRequest{}
	_ bin.Decoder     = &AccountSaveThemeRequest{}
	_ bin.BareEncoder = &AccountSaveThemeRequest{}
	_ bin.BareDecoder = &AccountSaveThemeRequest{}
)

func (s *AccountSaveThemeRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Theme == nil) {
		return false
	}
	if !(s.Unsave == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSaveThemeRequest) String() string {
	if s == nil {
		return "AccountSaveThemeRequest(nil)"
	}
	type Alias AccountSaveThemeRequest
	return fmt.Sprintf("AccountSaveThemeRequest%+v", Alias(*s))
}

// FillFrom fills AccountSaveThemeRequest from given interface.
func (s *AccountSaveThemeRequest) FillFrom(from interface {
	GetTheme() (value InputThemeClass)
	GetUnsave() (value bool)
}) {
	s.Theme = from.GetTheme()
	s.Unsave = from.GetUnsave()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSaveThemeRequest) TypeID() uint32 {
	return AccountSaveThemeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSaveThemeRequest) TypeName() string {
	return "account.saveTheme"
}

// TypeInfo returns info about TL type.
func (s *AccountSaveThemeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.saveTheme",
		ID:   AccountSaveThemeRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Theme",
			SchemaName: "theme",
		},
		{
			Name:       "Unsave",
			SchemaName: "unsave",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSaveThemeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.saveTheme#f257106c as nil")
	}
	b.PutID(AccountSaveThemeRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSaveThemeRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.saveTheme#f257106c as nil")
	}
	if s.Theme == nil {
		return fmt.Errorf("unable to encode account.saveTheme#f257106c: field theme is nil")
	}
	if err := s.Theme.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveTheme#f257106c: field theme: %w", err)
	}
	b.PutBool(s.Unsave)
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSaveThemeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.saveTheme#f257106c to nil")
	}
	if err := b.ConsumeID(AccountSaveThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.saveTheme#f257106c: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSaveThemeRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.saveTheme#f257106c to nil")
	}
	{
		value, err := DecodeInputTheme(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.saveTheme#f257106c: field theme: %w", err)
		}
		s.Theme = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.saveTheme#f257106c: field unsave: %w", err)
		}
		s.Unsave = value
	}
	return nil
}

// GetTheme returns value of Theme field.
func (s *AccountSaveThemeRequest) GetTheme() (value InputThemeClass) {
	if s == nil {
		return
	}
	return s.Theme
}

// GetUnsave returns value of Unsave field.
func (s *AccountSaveThemeRequest) GetUnsave() (value bool) {
	if s == nil {
		return
	}
	return s.Unsave
}

// AccountSaveTheme invokes method account.saveTheme#f257106c returning error if any.
// Save a theme
//
// Possible errors:
//
//	400 THEME_INVALID: Invalid theme provided.
//
// See https://core.telegram.org/method/account.saveTheme for reference.
func (c *Client) AccountSaveTheme(ctx context.Context, request *AccountSaveThemeRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
