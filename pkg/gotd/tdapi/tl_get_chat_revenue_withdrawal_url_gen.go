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

// GetChatRevenueWithdrawalURLRequest represents TL type `getChatRevenueWithdrawalUrl#1e320720`.
type GetChatRevenueWithdrawalURLRequest struct {
	// Chat identifier
	ChatID int64
	// The 2-step verification password of the current user
	Password string
}

// GetChatRevenueWithdrawalURLRequestTypeID is TL type id of GetChatRevenueWithdrawalURLRequest.
const GetChatRevenueWithdrawalURLRequestTypeID = 0x1e320720

// Ensuring interfaces in compile-time for GetChatRevenueWithdrawalURLRequest.
var (
	_ bin.Encoder     = &GetChatRevenueWithdrawalURLRequest{}
	_ bin.Decoder     = &GetChatRevenueWithdrawalURLRequest{}
	_ bin.BareEncoder = &GetChatRevenueWithdrawalURLRequest{}
	_ bin.BareDecoder = &GetChatRevenueWithdrawalURLRequest{}
)

func (g *GetChatRevenueWithdrawalURLRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.Password == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatRevenueWithdrawalURLRequest) String() string {
	if g == nil {
		return "GetChatRevenueWithdrawalURLRequest(nil)"
	}
	type Alias GetChatRevenueWithdrawalURLRequest
	return fmt.Sprintf("GetChatRevenueWithdrawalURLRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatRevenueWithdrawalURLRequest) TypeID() uint32 {
	return GetChatRevenueWithdrawalURLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatRevenueWithdrawalURLRequest) TypeName() string {
	return "getChatRevenueWithdrawalUrl"
}

// TypeInfo returns info about TL type.
func (g *GetChatRevenueWithdrawalURLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatRevenueWithdrawalUrl",
		ID:   GetChatRevenueWithdrawalURLRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Password",
			SchemaName: "password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatRevenueWithdrawalURLRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatRevenueWithdrawalUrl#1e320720 as nil")
	}
	b.PutID(GetChatRevenueWithdrawalURLRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatRevenueWithdrawalURLRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatRevenueWithdrawalUrl#1e320720 as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutString(g.Password)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatRevenueWithdrawalURLRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatRevenueWithdrawalUrl#1e320720 to nil")
	}
	if err := b.ConsumeID(GetChatRevenueWithdrawalURLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatRevenueWithdrawalUrl#1e320720: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatRevenueWithdrawalURLRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatRevenueWithdrawalUrl#1e320720 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatRevenueWithdrawalUrl#1e320720: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getChatRevenueWithdrawalUrl#1e320720: field password: %w", err)
		}
		g.Password = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatRevenueWithdrawalURLRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatRevenueWithdrawalUrl#1e320720 as nil")
	}
	b.ObjStart()
	b.PutID("getChatRevenueWithdrawalUrl")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("password")
	b.PutString(g.Password)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatRevenueWithdrawalURLRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatRevenueWithdrawalUrl#1e320720 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatRevenueWithdrawalUrl"); err != nil {
				return fmt.Errorf("unable to decode getChatRevenueWithdrawalUrl#1e320720: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatRevenueWithdrawalUrl#1e320720: field chat_id: %w", err)
			}
			g.ChatID = value
		case "password":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getChatRevenueWithdrawalUrl#1e320720: field password: %w", err)
			}
			g.Password = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatRevenueWithdrawalURLRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetPassword returns value of Password field.
func (g *GetChatRevenueWithdrawalURLRequest) GetPassword() (value string) {
	if g == nil {
		return
	}
	return g.Password
}

// GetChatRevenueWithdrawalURL invokes method getChatRevenueWithdrawalUrl#1e320720 returning error if any.
func (c *Client) GetChatRevenueWithdrawalURL(ctx context.Context, request *GetChatRevenueWithdrawalURLRequest) (*HTTPURL, error) {
	var result HTTPURL

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
