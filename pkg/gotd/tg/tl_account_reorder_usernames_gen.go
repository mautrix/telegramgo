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

// AccountReorderUsernamesRequest represents TL type `account.reorderUsernames#ef500eab`.
// Reorder usernames associated with the currently logged-in user.
//
// See https://core.telegram.org/method/account.reorderUsernames for reference.
type AccountReorderUsernamesRequest struct {
	// The new order for active usernames. All active usernames must be specified.
	Order []string
}

// AccountReorderUsernamesRequestTypeID is TL type id of AccountReorderUsernamesRequest.
const AccountReorderUsernamesRequestTypeID = 0xef500eab

// Ensuring interfaces in compile-time for AccountReorderUsernamesRequest.
var (
	_ bin.Encoder     = &AccountReorderUsernamesRequest{}
	_ bin.Decoder     = &AccountReorderUsernamesRequest{}
	_ bin.BareEncoder = &AccountReorderUsernamesRequest{}
	_ bin.BareDecoder = &AccountReorderUsernamesRequest{}
)

func (r *AccountReorderUsernamesRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Order == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountReorderUsernamesRequest) String() string {
	if r == nil {
		return "AccountReorderUsernamesRequest(nil)"
	}
	type Alias AccountReorderUsernamesRequest
	return fmt.Sprintf("AccountReorderUsernamesRequest%+v", Alias(*r))
}

// FillFrom fills AccountReorderUsernamesRequest from given interface.
func (r *AccountReorderUsernamesRequest) FillFrom(from interface {
	GetOrder() (value []string)
}) {
	r.Order = from.GetOrder()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountReorderUsernamesRequest) TypeID() uint32 {
	return AccountReorderUsernamesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountReorderUsernamesRequest) TypeName() string {
	return "account.reorderUsernames"
}

// TypeInfo returns info about TL type.
func (r *AccountReorderUsernamesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.reorderUsernames",
		ID:   AccountReorderUsernamesRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Order",
			SchemaName: "order",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *AccountReorderUsernamesRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.reorderUsernames#ef500eab as nil")
	}
	b.PutID(AccountReorderUsernamesRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AccountReorderUsernamesRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.reorderUsernames#ef500eab as nil")
	}
	b.PutVectorHeader(len(r.Order))
	for _, v := range r.Order {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountReorderUsernamesRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.reorderUsernames#ef500eab to nil")
	}
	if err := b.ConsumeID(AccountReorderUsernamesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.reorderUsernames#ef500eab: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AccountReorderUsernamesRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.reorderUsernames#ef500eab to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.reorderUsernames#ef500eab: field order: %w", err)
		}

		if headerLen > 0 {
			r.Order = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode account.reorderUsernames#ef500eab: field order: %w", err)
			}
			r.Order = append(r.Order, value)
		}
	}
	return nil
}

// GetOrder returns value of Order field.
func (r *AccountReorderUsernamesRequest) GetOrder() (value []string) {
	if r == nil {
		return
	}
	return r.Order
}

// AccountReorderUsernames invokes method account.reorderUsernames#ef500eab returning error if any.
// Reorder usernames associated with the currently logged-in user.
//
// Possible errors:
//
//	400 ORDER_INVALID: The specified username order is invalid.
//	400 USERNAME_NOT_MODIFIED: The username was not modified.
//
// See https://core.telegram.org/method/account.reorderUsernames for reference.
func (c *Client) AccountReorderUsernames(ctx context.Context, order []string) (bool, error) {
	var result BoolBox

	request := &AccountReorderUsernamesRequest{
		Order: order,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
