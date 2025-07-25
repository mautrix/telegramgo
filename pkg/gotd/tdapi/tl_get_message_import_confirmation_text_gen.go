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

// GetMessageImportConfirmationTextRequest represents TL type `getMessageImportConfirmationText#174881a8`.
type GetMessageImportConfirmationTextRequest struct {
	// Identifier of a chat to which the messages will be imported. It must be an identifier
	// of a private chat with a mutual contact or an identifier of a supergroup chat with
	// can_change_info member right
	ChatID int64
}

// GetMessageImportConfirmationTextRequestTypeID is TL type id of GetMessageImportConfirmationTextRequest.
const GetMessageImportConfirmationTextRequestTypeID = 0x174881a8

// Ensuring interfaces in compile-time for GetMessageImportConfirmationTextRequest.
var (
	_ bin.Encoder     = &GetMessageImportConfirmationTextRequest{}
	_ bin.Decoder     = &GetMessageImportConfirmationTextRequest{}
	_ bin.BareEncoder = &GetMessageImportConfirmationTextRequest{}
	_ bin.BareDecoder = &GetMessageImportConfirmationTextRequest{}
)

func (g *GetMessageImportConfirmationTextRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetMessageImportConfirmationTextRequest) String() string {
	if g == nil {
		return "GetMessageImportConfirmationTextRequest(nil)"
	}
	type Alias GetMessageImportConfirmationTextRequest
	return fmt.Sprintf("GetMessageImportConfirmationTextRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMessageImportConfirmationTextRequest) TypeID() uint32 {
	return GetMessageImportConfirmationTextRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMessageImportConfirmationTextRequest) TypeName() string {
	return "getMessageImportConfirmationText"
}

// TypeInfo returns info about TL type.
func (g *GetMessageImportConfirmationTextRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMessageImportConfirmationText",
		ID:   GetMessageImportConfirmationTextRequestTypeID,
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetMessageImportConfirmationTextRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageImportConfirmationText#174881a8 as nil")
	}
	b.PutID(GetMessageImportConfirmationTextRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMessageImportConfirmationTextRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageImportConfirmationText#174881a8 as nil")
	}
	b.PutInt53(g.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMessageImportConfirmationTextRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageImportConfirmationText#174881a8 to nil")
	}
	if err := b.ConsumeID(GetMessageImportConfirmationTextRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMessageImportConfirmationText#174881a8: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMessageImportConfirmationTextRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageImportConfirmationText#174881a8 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageImportConfirmationText#174881a8: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetMessageImportConfirmationTextRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageImportConfirmationText#174881a8 as nil")
	}
	b.ObjStart()
	b.PutID("getMessageImportConfirmationText")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetMessageImportConfirmationTextRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageImportConfirmationText#174881a8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getMessageImportConfirmationText"); err != nil {
				return fmt.Errorf("unable to decode getMessageImportConfirmationText#174881a8: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageImportConfirmationText#174881a8: field chat_id: %w", err)
			}
			g.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetMessageImportConfirmationTextRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetMessageImportConfirmationText invokes method getMessageImportConfirmationText#174881a8 returning error if any.
func (c *Client) GetMessageImportConfirmationText(ctx context.Context, chatid int64) (*Text, error) {
	var result Text

	request := &GetMessageImportConfirmationTextRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
