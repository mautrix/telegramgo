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

// ChatRevenueAmount represents TL type `chatRevenueAmount#a648ce58`.
type ChatRevenueAmount struct {
	// Cryptocurrency in which revenue is calculated
	Cryptocurrency string
	// Total amount of the cryptocurrency earned, in the smallest units of the cryptocurrency
	TotalAmount int64
	// Amount of the cryptocurrency that isn't withdrawn yet, in the smallest units of the
	// cryptocurrency
	BalanceAmount int64
	// Amount of the cryptocurrency available for withdrawal, in the smallest units of the
	// cryptocurrency
	AvailableAmount int64
	// True, if Telegram Stars can be withdrawn now or later
	WithdrawalEnabled bool
}

// ChatRevenueAmountTypeID is TL type id of ChatRevenueAmount.
const ChatRevenueAmountTypeID = 0xa648ce58

// Ensuring interfaces in compile-time for ChatRevenueAmount.
var (
	_ bin.Encoder     = &ChatRevenueAmount{}
	_ bin.Decoder     = &ChatRevenueAmount{}
	_ bin.BareEncoder = &ChatRevenueAmount{}
	_ bin.BareDecoder = &ChatRevenueAmount{}
)

func (c *ChatRevenueAmount) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Cryptocurrency == "") {
		return false
	}
	if !(c.TotalAmount == 0) {
		return false
	}
	if !(c.BalanceAmount == 0) {
		return false
	}
	if !(c.AvailableAmount == 0) {
		return false
	}
	if !(c.WithdrawalEnabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatRevenueAmount) String() string {
	if c == nil {
		return "ChatRevenueAmount(nil)"
	}
	type Alias ChatRevenueAmount
	return fmt.Sprintf("ChatRevenueAmount%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatRevenueAmount) TypeID() uint32 {
	return ChatRevenueAmountTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatRevenueAmount) TypeName() string {
	return "chatRevenueAmount"
}

// TypeInfo returns info about TL type.
func (c *ChatRevenueAmount) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatRevenueAmount",
		ID:   ChatRevenueAmountTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Cryptocurrency",
			SchemaName: "cryptocurrency",
		},
		{
			Name:       "TotalAmount",
			SchemaName: "total_amount",
		},
		{
			Name:       "BalanceAmount",
			SchemaName: "balance_amount",
		},
		{
			Name:       "AvailableAmount",
			SchemaName: "available_amount",
		},
		{
			Name:       "WithdrawalEnabled",
			SchemaName: "withdrawal_enabled",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatRevenueAmount) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueAmount#a648ce58 as nil")
	}
	b.PutID(ChatRevenueAmountTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatRevenueAmount) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueAmount#a648ce58 as nil")
	}
	b.PutString(c.Cryptocurrency)
	b.PutLong(c.TotalAmount)
	b.PutLong(c.BalanceAmount)
	b.PutLong(c.AvailableAmount)
	b.PutBool(c.WithdrawalEnabled)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatRevenueAmount) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueAmount#a648ce58 to nil")
	}
	if err := b.ConsumeID(ChatRevenueAmountTypeID); err != nil {
		return fmt.Errorf("unable to decode chatRevenueAmount#a648ce58: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatRevenueAmount) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueAmount#a648ce58 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueAmount#a648ce58: field cryptocurrency: %w", err)
		}
		c.Cryptocurrency = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueAmount#a648ce58: field total_amount: %w", err)
		}
		c.TotalAmount = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueAmount#a648ce58: field balance_amount: %w", err)
		}
		c.BalanceAmount = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueAmount#a648ce58: field available_amount: %w", err)
		}
		c.AvailableAmount = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueAmount#a648ce58: field withdrawal_enabled: %w", err)
		}
		c.WithdrawalEnabled = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatRevenueAmount) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueAmount#a648ce58 as nil")
	}
	b.ObjStart()
	b.PutID("chatRevenueAmount")
	b.Comma()
	b.FieldStart("cryptocurrency")
	b.PutString(c.Cryptocurrency)
	b.Comma()
	b.FieldStart("total_amount")
	b.PutLong(c.TotalAmount)
	b.Comma()
	b.FieldStart("balance_amount")
	b.PutLong(c.BalanceAmount)
	b.Comma()
	b.FieldStart("available_amount")
	b.PutLong(c.AvailableAmount)
	b.Comma()
	b.FieldStart("withdrawal_enabled")
	b.PutBool(c.WithdrawalEnabled)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatRevenueAmount) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueAmount#a648ce58 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatRevenueAmount"); err != nil {
				return fmt.Errorf("unable to decode chatRevenueAmount#a648ce58: %w", err)
			}
		case "cryptocurrency":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueAmount#a648ce58: field cryptocurrency: %w", err)
			}
			c.Cryptocurrency = value
		case "total_amount":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueAmount#a648ce58: field total_amount: %w", err)
			}
			c.TotalAmount = value
		case "balance_amount":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueAmount#a648ce58: field balance_amount: %w", err)
			}
			c.BalanceAmount = value
		case "available_amount":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueAmount#a648ce58: field available_amount: %w", err)
			}
			c.AvailableAmount = value
		case "withdrawal_enabled":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueAmount#a648ce58: field withdrawal_enabled: %w", err)
			}
			c.WithdrawalEnabled = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCryptocurrency returns value of Cryptocurrency field.
func (c *ChatRevenueAmount) GetCryptocurrency() (value string) {
	if c == nil {
		return
	}
	return c.Cryptocurrency
}

// GetTotalAmount returns value of TotalAmount field.
func (c *ChatRevenueAmount) GetTotalAmount() (value int64) {
	if c == nil {
		return
	}
	return c.TotalAmount
}

// GetBalanceAmount returns value of BalanceAmount field.
func (c *ChatRevenueAmount) GetBalanceAmount() (value int64) {
	if c == nil {
		return
	}
	return c.BalanceAmount
}

// GetAvailableAmount returns value of AvailableAmount field.
func (c *ChatRevenueAmount) GetAvailableAmount() (value int64) {
	if c == nil {
		return
	}
	return c.AvailableAmount
}

// GetWithdrawalEnabled returns value of WithdrawalEnabled field.
func (c *ChatRevenueAmount) GetWithdrawalEnabled() (value bool) {
	if c == nil {
		return
	}
	return c.WithdrawalEnabled
}
