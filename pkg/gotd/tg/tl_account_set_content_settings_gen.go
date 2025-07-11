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

// AccountSetContentSettingsRequest represents TL type `account.setContentSettings#b574b16b`.
// Set sensitive content settings (for viewing or hiding NSFW content)
//
// See https://core.telegram.org/method/account.setContentSettings for reference.
type AccountSetContentSettingsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Enable NSFW content
	SensitiveEnabled bool
}

// AccountSetContentSettingsRequestTypeID is TL type id of AccountSetContentSettingsRequest.
const AccountSetContentSettingsRequestTypeID = 0xb574b16b

// Ensuring interfaces in compile-time for AccountSetContentSettingsRequest.
var (
	_ bin.Encoder     = &AccountSetContentSettingsRequest{}
	_ bin.Decoder     = &AccountSetContentSettingsRequest{}
	_ bin.BareEncoder = &AccountSetContentSettingsRequest{}
	_ bin.BareDecoder = &AccountSetContentSettingsRequest{}
)

func (s *AccountSetContentSettingsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.SensitiveEnabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSetContentSettingsRequest) String() string {
	if s == nil {
		return "AccountSetContentSettingsRequest(nil)"
	}
	type Alias AccountSetContentSettingsRequest
	return fmt.Sprintf("AccountSetContentSettingsRequest%+v", Alias(*s))
}

// FillFrom fills AccountSetContentSettingsRequest from given interface.
func (s *AccountSetContentSettingsRequest) FillFrom(from interface {
	GetSensitiveEnabled() (value bool)
}) {
	s.SensitiveEnabled = from.GetSensitiveEnabled()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSetContentSettingsRequest) TypeID() uint32 {
	return AccountSetContentSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSetContentSettingsRequest) TypeName() string {
	return "account.setContentSettings"
}

// TypeInfo returns info about TL type.
func (s *AccountSetContentSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.setContentSettings",
		ID:   AccountSetContentSettingsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SensitiveEnabled",
			SchemaName: "sensitive_enabled",
			Null:       !s.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *AccountSetContentSettingsRequest) SetFlags() {
	if !(s.SensitiveEnabled == false) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *AccountSetContentSettingsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setContentSettings#b574b16b as nil")
	}
	b.PutID(AccountSetContentSettingsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSetContentSettingsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setContentSettings#b574b16b as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.setContentSettings#b574b16b: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSetContentSettingsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setContentSettings#b574b16b to nil")
	}
	if err := b.ConsumeID(AccountSetContentSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.setContentSettings#b574b16b: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSetContentSettingsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setContentSettings#b574b16b to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.setContentSettings#b574b16b: field flags: %w", err)
		}
	}
	s.SensitiveEnabled = s.Flags.Has(0)
	return nil
}

// SetSensitiveEnabled sets value of SensitiveEnabled conditional field.
func (s *AccountSetContentSettingsRequest) SetSensitiveEnabled(value bool) {
	if value {
		s.Flags.Set(0)
		s.SensitiveEnabled = true
	} else {
		s.Flags.Unset(0)
		s.SensitiveEnabled = false
	}
}

// GetSensitiveEnabled returns value of SensitiveEnabled conditional field.
func (s *AccountSetContentSettingsRequest) GetSensitiveEnabled() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(0)
}

// AccountSetContentSettings invokes method account.setContentSettings#b574b16b returning error if any.
// Set sensitive content settings (for viewing or hiding NSFW content)
//
// Possible errors:
//
//	403 SENSITIVE_CHANGE_FORBIDDEN: You can't change your sensitive content settings.
//
// See https://core.telegram.org/method/account.setContentSettings for reference.
func (c *Client) AccountSetContentSettings(ctx context.Context, request *AccountSetContentSettingsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
