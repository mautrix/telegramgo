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

// AccountUpdateBirthdayRequest represents TL type `account.updateBirthday#cc6e0c11`.
// Update our birthday, see here »¹ for more info.
//
// Links:
//  1. https://core.telegram.org/api/profile#birthday
//
// See https://core.telegram.org/method/account.updateBirthday for reference.
type AccountUpdateBirthdayRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Birthday.
	//
	// Use SetBirthday and GetBirthday helpers.
	Birthday Birthday
}

// AccountUpdateBirthdayRequestTypeID is TL type id of AccountUpdateBirthdayRequest.
const AccountUpdateBirthdayRequestTypeID = 0xcc6e0c11

// Ensuring interfaces in compile-time for AccountUpdateBirthdayRequest.
var (
	_ bin.Encoder     = &AccountUpdateBirthdayRequest{}
	_ bin.Decoder     = &AccountUpdateBirthdayRequest{}
	_ bin.BareEncoder = &AccountUpdateBirthdayRequest{}
	_ bin.BareDecoder = &AccountUpdateBirthdayRequest{}
)

func (u *AccountUpdateBirthdayRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Birthday.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateBirthdayRequest) String() string {
	if u == nil {
		return "AccountUpdateBirthdayRequest(nil)"
	}
	type Alias AccountUpdateBirthdayRequest
	return fmt.Sprintf("AccountUpdateBirthdayRequest%+v", Alias(*u))
}

// FillFrom fills AccountUpdateBirthdayRequest from given interface.
func (u *AccountUpdateBirthdayRequest) FillFrom(from interface {
	GetBirthday() (value Birthday, ok bool)
}) {
	if val, ok := from.GetBirthday(); ok {
		u.Birthday = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdateBirthdayRequest) TypeID() uint32 {
	return AccountUpdateBirthdayRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdateBirthdayRequest) TypeName() string {
	return "account.updateBirthday"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdateBirthdayRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updateBirthday",
		ID:   AccountUpdateBirthdayRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Birthday",
			SchemaName: "birthday",
			Null:       !u.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *AccountUpdateBirthdayRequest) SetFlags() {
	if !(u.Birthday.Zero()) {
		u.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (u *AccountUpdateBirthdayRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateBirthday#cc6e0c11 as nil")
	}
	b.PutID(AccountUpdateBirthdayRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdateBirthdayRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateBirthday#cc6e0c11 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateBirthday#cc6e0c11: field flags: %w", err)
	}
	if u.Flags.Has(0) {
		if err := u.Birthday.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.updateBirthday#cc6e0c11: field birthday: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateBirthdayRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateBirthday#cc6e0c11 to nil")
	}
	if err := b.ConsumeID(AccountUpdateBirthdayRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateBirthday#cc6e0c11: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdateBirthdayRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateBirthday#cc6e0c11 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateBirthday#cc6e0c11: field flags: %w", err)
		}
	}
	if u.Flags.Has(0) {
		if err := u.Birthday.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateBirthday#cc6e0c11: field birthday: %w", err)
		}
	}
	return nil
}

// SetBirthday sets value of Birthday conditional field.
func (u *AccountUpdateBirthdayRequest) SetBirthday(value Birthday) {
	u.Flags.Set(0)
	u.Birthday = value
}

// GetBirthday returns value of Birthday conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateBirthdayRequest) GetBirthday() (value Birthday, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.Birthday, true
}

// AccountUpdateBirthday invokes method account.updateBirthday#cc6e0c11 returning error if any.
// Update our birthday, see here »¹ for more info.
//
// Links:
//  1. https://core.telegram.org/api/profile#birthday
//
// Possible errors:
//
//	400 BIRTHDAY_INVALID: An invalid age was specified, must be between 0 and 150 years.
//
// See https://core.telegram.org/method/account.updateBirthday for reference.
func (c *Client) AccountUpdateBirthday(ctx context.Context, request *AccountUpdateBirthdayRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
