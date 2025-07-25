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

// CheckCreatedPublicChatsLimitRequest represents TL type `checkCreatedPublicChatsLimit#e5717fa1`.
type CheckCreatedPublicChatsLimitRequest struct {
	// Type of the public chats, for which to check the limit
	Type PublicChatTypeClass
}

// CheckCreatedPublicChatsLimitRequestTypeID is TL type id of CheckCreatedPublicChatsLimitRequest.
const CheckCreatedPublicChatsLimitRequestTypeID = 0xe5717fa1

// Ensuring interfaces in compile-time for CheckCreatedPublicChatsLimitRequest.
var (
	_ bin.Encoder     = &CheckCreatedPublicChatsLimitRequest{}
	_ bin.Decoder     = &CheckCreatedPublicChatsLimitRequest{}
	_ bin.BareEncoder = &CheckCreatedPublicChatsLimitRequest{}
	_ bin.BareDecoder = &CheckCreatedPublicChatsLimitRequest{}
)

func (c *CheckCreatedPublicChatsLimitRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckCreatedPublicChatsLimitRequest) String() string {
	if c == nil {
		return "CheckCreatedPublicChatsLimitRequest(nil)"
	}
	type Alias CheckCreatedPublicChatsLimitRequest
	return fmt.Sprintf("CheckCreatedPublicChatsLimitRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckCreatedPublicChatsLimitRequest) TypeID() uint32 {
	return CheckCreatedPublicChatsLimitRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckCreatedPublicChatsLimitRequest) TypeName() string {
	return "checkCreatedPublicChatsLimit"
}

// TypeInfo returns info about TL type.
func (c *CheckCreatedPublicChatsLimitRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkCreatedPublicChatsLimit",
		ID:   CheckCreatedPublicChatsLimitRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckCreatedPublicChatsLimitRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkCreatedPublicChatsLimit#e5717fa1 as nil")
	}
	b.PutID(CheckCreatedPublicChatsLimitRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckCreatedPublicChatsLimitRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkCreatedPublicChatsLimit#e5717fa1 as nil")
	}
	if c.Type == nil {
		return fmt.Errorf("unable to encode checkCreatedPublicChatsLimit#e5717fa1: field type is nil")
	}
	if err := c.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode checkCreatedPublicChatsLimit#e5717fa1: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckCreatedPublicChatsLimitRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkCreatedPublicChatsLimit#e5717fa1 to nil")
	}
	if err := b.ConsumeID(CheckCreatedPublicChatsLimitRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode checkCreatedPublicChatsLimit#e5717fa1: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckCreatedPublicChatsLimitRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkCreatedPublicChatsLimit#e5717fa1 to nil")
	}
	{
		value, err := DecodePublicChatType(b)
		if err != nil {
			return fmt.Errorf("unable to decode checkCreatedPublicChatsLimit#e5717fa1: field type: %w", err)
		}
		c.Type = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckCreatedPublicChatsLimitRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkCreatedPublicChatsLimit#e5717fa1 as nil")
	}
	b.ObjStart()
	b.PutID("checkCreatedPublicChatsLimit")
	b.Comma()
	b.FieldStart("type")
	if c.Type == nil {
		return fmt.Errorf("unable to encode checkCreatedPublicChatsLimit#e5717fa1: field type is nil")
	}
	if err := c.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode checkCreatedPublicChatsLimit#e5717fa1: field type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckCreatedPublicChatsLimitRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkCreatedPublicChatsLimit#e5717fa1 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkCreatedPublicChatsLimit"); err != nil {
				return fmt.Errorf("unable to decode checkCreatedPublicChatsLimit#e5717fa1: %w", err)
			}
		case "type":
			value, err := DecodeTDLibJSONPublicChatType(b)
			if err != nil {
				return fmt.Errorf("unable to decode checkCreatedPublicChatsLimit#e5717fa1: field type: %w", err)
			}
			c.Type = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetType returns value of Type field.
func (c *CheckCreatedPublicChatsLimitRequest) GetType() (value PublicChatTypeClass) {
	if c == nil {
		return
	}
	return c.Type
}

// CheckCreatedPublicChatsLimit invokes method checkCreatedPublicChatsLimit#e5717fa1 returning error if any.
func (c *Client) CheckCreatedPublicChatsLimit(ctx context.Context, type_ PublicChatTypeClass) error {
	var ok Ok

	request := &CheckCreatedPublicChatsLimitRequest{
		Type: type_,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
