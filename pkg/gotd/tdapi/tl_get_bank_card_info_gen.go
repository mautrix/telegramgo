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

// GetBankCardInfoRequest represents TL type `getBankCardInfo#b1e31db0`.
type GetBankCardInfoRequest struct {
	// The bank card number
	BankCardNumber string
}

// GetBankCardInfoRequestTypeID is TL type id of GetBankCardInfoRequest.
const GetBankCardInfoRequestTypeID = 0xb1e31db0

// Ensuring interfaces in compile-time for GetBankCardInfoRequest.
var (
	_ bin.Encoder     = &GetBankCardInfoRequest{}
	_ bin.Decoder     = &GetBankCardInfoRequest{}
	_ bin.BareEncoder = &GetBankCardInfoRequest{}
	_ bin.BareDecoder = &GetBankCardInfoRequest{}
)

func (g *GetBankCardInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.BankCardNumber == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetBankCardInfoRequest) String() string {
	if g == nil {
		return "GetBankCardInfoRequest(nil)"
	}
	type Alias GetBankCardInfoRequest
	return fmt.Sprintf("GetBankCardInfoRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetBankCardInfoRequest) TypeID() uint32 {
	return GetBankCardInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetBankCardInfoRequest) TypeName() string {
	return "getBankCardInfo"
}

// TypeInfo returns info about TL type.
func (g *GetBankCardInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getBankCardInfo",
		ID:   GetBankCardInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BankCardNumber",
			SchemaName: "bank_card_number",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetBankCardInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getBankCardInfo#b1e31db0 as nil")
	}
	b.PutID(GetBankCardInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetBankCardInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getBankCardInfo#b1e31db0 as nil")
	}
	b.PutString(g.BankCardNumber)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetBankCardInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getBankCardInfo#b1e31db0 to nil")
	}
	if err := b.ConsumeID(GetBankCardInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getBankCardInfo#b1e31db0: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetBankCardInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getBankCardInfo#b1e31db0 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getBankCardInfo#b1e31db0: field bank_card_number: %w", err)
		}
		g.BankCardNumber = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetBankCardInfoRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getBankCardInfo#b1e31db0 as nil")
	}
	b.ObjStart()
	b.PutID("getBankCardInfo")
	b.Comma()
	b.FieldStart("bank_card_number")
	b.PutString(g.BankCardNumber)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetBankCardInfoRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getBankCardInfo#b1e31db0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getBankCardInfo"); err != nil {
				return fmt.Errorf("unable to decode getBankCardInfo#b1e31db0: %w", err)
			}
		case "bank_card_number":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getBankCardInfo#b1e31db0: field bank_card_number: %w", err)
			}
			g.BankCardNumber = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBankCardNumber returns value of BankCardNumber field.
func (g *GetBankCardInfoRequest) GetBankCardNumber() (value string) {
	if g == nil {
		return
	}
	return g.BankCardNumber
}

// GetBankCardInfo invokes method getBankCardInfo#b1e31db0 returning error if any.
func (c *Client) GetBankCardInfo(ctx context.Context, bankcardnumber string) (*BankCardInfo, error) {
	var result BankCardInfo

	request := &GetBankCardInfoRequest{
		BankCardNumber: bankcardnumber,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
