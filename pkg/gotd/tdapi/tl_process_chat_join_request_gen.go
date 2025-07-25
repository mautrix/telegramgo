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

// ProcessChatJoinRequestRequest represents TL type `processChatJoinRequest#3be534a3`.
type ProcessChatJoinRequestRequest struct {
	// Chat identifier
	ChatID int64
	// Identifier of the user that sent the request
	UserID int64
	// Pass true to approve the request; pass false to decline it
	Approve bool
}

// ProcessChatJoinRequestRequestTypeID is TL type id of ProcessChatJoinRequestRequest.
const ProcessChatJoinRequestRequestTypeID = 0x3be534a3

// Ensuring interfaces in compile-time for ProcessChatJoinRequestRequest.
var (
	_ bin.Encoder     = &ProcessChatJoinRequestRequest{}
	_ bin.Decoder     = &ProcessChatJoinRequestRequest{}
	_ bin.BareEncoder = &ProcessChatJoinRequestRequest{}
	_ bin.BareDecoder = &ProcessChatJoinRequestRequest{}
)

func (p *ProcessChatJoinRequestRequest) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ChatID == 0) {
		return false
	}
	if !(p.UserID == 0) {
		return false
	}
	if !(p.Approve == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *ProcessChatJoinRequestRequest) String() string {
	if p == nil {
		return "ProcessChatJoinRequestRequest(nil)"
	}
	type Alias ProcessChatJoinRequestRequest
	return fmt.Sprintf("ProcessChatJoinRequestRequest%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ProcessChatJoinRequestRequest) TypeID() uint32 {
	return ProcessChatJoinRequestRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ProcessChatJoinRequestRequest) TypeName() string {
	return "processChatJoinRequest"
}

// TypeInfo returns info about TL type.
func (p *ProcessChatJoinRequestRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "processChatJoinRequest",
		ID:   ProcessChatJoinRequestRequestTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Approve",
			SchemaName: "approve",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *ProcessChatJoinRequestRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode processChatJoinRequest#3be534a3 as nil")
	}
	b.PutID(ProcessChatJoinRequestRequestTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *ProcessChatJoinRequestRequest) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode processChatJoinRequest#3be534a3 as nil")
	}
	b.PutInt53(p.ChatID)
	b.PutInt53(p.UserID)
	b.PutBool(p.Approve)
	return nil
}

// Decode implements bin.Decoder.
func (p *ProcessChatJoinRequestRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode processChatJoinRequest#3be534a3 to nil")
	}
	if err := b.ConsumeID(ProcessChatJoinRequestRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode processChatJoinRequest#3be534a3: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *ProcessChatJoinRequestRequest) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode processChatJoinRequest#3be534a3 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode processChatJoinRequest#3be534a3: field chat_id: %w", err)
		}
		p.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode processChatJoinRequest#3be534a3: field user_id: %w", err)
		}
		p.UserID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode processChatJoinRequest#3be534a3: field approve: %w", err)
		}
		p.Approve = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *ProcessChatJoinRequestRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode processChatJoinRequest#3be534a3 as nil")
	}
	b.ObjStart()
	b.PutID("processChatJoinRequest")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(p.ChatID)
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(p.UserID)
	b.Comma()
	b.FieldStart("approve")
	b.PutBool(p.Approve)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *ProcessChatJoinRequestRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode processChatJoinRequest#3be534a3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("processChatJoinRequest"); err != nil {
				return fmt.Errorf("unable to decode processChatJoinRequest#3be534a3: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode processChatJoinRequest#3be534a3: field chat_id: %w", err)
			}
			p.ChatID = value
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode processChatJoinRequest#3be534a3: field user_id: %w", err)
			}
			p.UserID = value
		case "approve":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode processChatJoinRequest#3be534a3: field approve: %w", err)
			}
			p.Approve = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (p *ProcessChatJoinRequestRequest) GetChatID() (value int64) {
	if p == nil {
		return
	}
	return p.ChatID
}

// GetUserID returns value of UserID field.
func (p *ProcessChatJoinRequestRequest) GetUserID() (value int64) {
	if p == nil {
		return
	}
	return p.UserID
}

// GetApprove returns value of Approve field.
func (p *ProcessChatJoinRequestRequest) GetApprove() (value bool) {
	if p == nil {
		return
	}
	return p.Approve
}

// ProcessChatJoinRequest invokes method processChatJoinRequest#3be534a3 returning error if any.
func (c *Client) ProcessChatJoinRequest(ctx context.Context, request *ProcessChatJoinRequestRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
