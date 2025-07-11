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

// AccountFinishTakeoutSessionRequest represents TL type `account.finishTakeoutSession#1d2652ee`.
// Terminate a takeout session, see here » for more info¹.
//
// Links:
//  1. https://core.telegram.org/api/takeout
//
// See https://core.telegram.org/method/account.finishTakeoutSession for reference.
type AccountFinishTakeoutSessionRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Data exported successfully
	Success bool
}

// AccountFinishTakeoutSessionRequestTypeID is TL type id of AccountFinishTakeoutSessionRequest.
const AccountFinishTakeoutSessionRequestTypeID = 0x1d2652ee

// Ensuring interfaces in compile-time for AccountFinishTakeoutSessionRequest.
var (
	_ bin.Encoder     = &AccountFinishTakeoutSessionRequest{}
	_ bin.Decoder     = &AccountFinishTakeoutSessionRequest{}
	_ bin.BareEncoder = &AccountFinishTakeoutSessionRequest{}
	_ bin.BareDecoder = &AccountFinishTakeoutSessionRequest{}
)

func (f *AccountFinishTakeoutSessionRequest) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Flags.Zero()) {
		return false
	}
	if !(f.Success == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *AccountFinishTakeoutSessionRequest) String() string {
	if f == nil {
		return "AccountFinishTakeoutSessionRequest(nil)"
	}
	type Alias AccountFinishTakeoutSessionRequest
	return fmt.Sprintf("AccountFinishTakeoutSessionRequest%+v", Alias(*f))
}

// FillFrom fills AccountFinishTakeoutSessionRequest from given interface.
func (f *AccountFinishTakeoutSessionRequest) FillFrom(from interface {
	GetSuccess() (value bool)
}) {
	f.Success = from.GetSuccess()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountFinishTakeoutSessionRequest) TypeID() uint32 {
	return AccountFinishTakeoutSessionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountFinishTakeoutSessionRequest) TypeName() string {
	return "account.finishTakeoutSession"
}

// TypeInfo returns info about TL type.
func (f *AccountFinishTakeoutSessionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.finishTakeoutSession",
		ID:   AccountFinishTakeoutSessionRequestTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Success",
			SchemaName: "success",
			Null:       !f.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (f *AccountFinishTakeoutSessionRequest) SetFlags() {
	if !(f.Success == false) {
		f.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (f *AccountFinishTakeoutSessionRequest) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode account.finishTakeoutSession#1d2652ee as nil")
	}
	b.PutID(AccountFinishTakeoutSessionRequestTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *AccountFinishTakeoutSessionRequest) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode account.finishTakeoutSession#1d2652ee as nil")
	}
	f.SetFlags()
	if err := f.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.finishTakeoutSession#1d2652ee: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *AccountFinishTakeoutSessionRequest) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode account.finishTakeoutSession#1d2652ee to nil")
	}
	if err := b.ConsumeID(AccountFinishTakeoutSessionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.finishTakeoutSession#1d2652ee: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *AccountFinishTakeoutSessionRequest) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode account.finishTakeoutSession#1d2652ee to nil")
	}
	{
		if err := f.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.finishTakeoutSession#1d2652ee: field flags: %w", err)
		}
	}
	f.Success = f.Flags.Has(0)
	return nil
}

// SetSuccess sets value of Success conditional field.
func (f *AccountFinishTakeoutSessionRequest) SetSuccess(value bool) {
	if value {
		f.Flags.Set(0)
		f.Success = true
	} else {
		f.Flags.Unset(0)
		f.Success = false
	}
}

// GetSuccess returns value of Success conditional field.
func (f *AccountFinishTakeoutSessionRequest) GetSuccess() (value bool) {
	if f == nil {
		return
	}
	return f.Flags.Has(0)
}

// AccountFinishTakeoutSession invokes method account.finishTakeoutSession#1d2652ee returning error if any.
// Terminate a takeout session, see here » for more info¹.
//
// Links:
//  1. https://core.telegram.org/api/takeout
//
// Possible errors:
//
//	403 TAKEOUT_REQUIRED: A takeout session needs to be initialized first, see here » for more info.
//
// See https://core.telegram.org/method/account.finishTakeoutSession for reference.
func (c *Client) AccountFinishTakeoutSession(ctx context.Context, request *AccountFinishTakeoutSessionRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
