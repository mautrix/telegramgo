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

// AddQuickReplyShortcutInlineQueryResultMessageRequest represents TL type `addQuickReplyShortcutInlineQueryResultMessage#87c02a04`.
type AddQuickReplyShortcutInlineQueryResultMessageRequest struct {
	// Name of the target shortcut
	ShortcutName string
	// Identifier of a quick reply message in the same shortcut to be replied; pass 0 if none
	ReplyToMessageID int64
	// Identifier of the inline query
	QueryID int64
	// Identifier of the inline query result
	ResultID string
	// Pass true to hide the bot, via which the message is sent. Can be used only for bots
	// getOption("animation_search_bot_username"), getOption("photo_search_bot_username"),
	// and getOption("venue_search_bot_username")
	HideViaBot bool
}

// AddQuickReplyShortcutInlineQueryResultMessageRequestTypeID is TL type id of AddQuickReplyShortcutInlineQueryResultMessageRequest.
const AddQuickReplyShortcutInlineQueryResultMessageRequestTypeID = 0x87c02a04

// Ensuring interfaces in compile-time for AddQuickReplyShortcutInlineQueryResultMessageRequest.
var (
	_ bin.Encoder     = &AddQuickReplyShortcutInlineQueryResultMessageRequest{}
	_ bin.Decoder     = &AddQuickReplyShortcutInlineQueryResultMessageRequest{}
	_ bin.BareEncoder = &AddQuickReplyShortcutInlineQueryResultMessageRequest{}
	_ bin.BareDecoder = &AddQuickReplyShortcutInlineQueryResultMessageRequest{}
)

func (a *AddQuickReplyShortcutInlineQueryResultMessageRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ShortcutName == "") {
		return false
	}
	if !(a.ReplyToMessageID == 0) {
		return false
	}
	if !(a.QueryID == 0) {
		return false
	}
	if !(a.ResultID == "") {
		return false
	}
	if !(a.HideViaBot == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AddQuickReplyShortcutInlineQueryResultMessageRequest) String() string {
	if a == nil {
		return "AddQuickReplyShortcutInlineQueryResultMessageRequest(nil)"
	}
	type Alias AddQuickReplyShortcutInlineQueryResultMessageRequest
	return fmt.Sprintf("AddQuickReplyShortcutInlineQueryResultMessageRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AddQuickReplyShortcutInlineQueryResultMessageRequest) TypeID() uint32 {
	return AddQuickReplyShortcutInlineQueryResultMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AddQuickReplyShortcutInlineQueryResultMessageRequest) TypeName() string {
	return "addQuickReplyShortcutInlineQueryResultMessage"
}

// TypeInfo returns info about TL type.
func (a *AddQuickReplyShortcutInlineQueryResultMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "addQuickReplyShortcutInlineQueryResultMessage",
		ID:   AddQuickReplyShortcutInlineQueryResultMessageRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ShortcutName",
			SchemaName: "shortcut_name",
		},
		{
			Name:       "ReplyToMessageID",
			SchemaName: "reply_to_message_id",
		},
		{
			Name:       "QueryID",
			SchemaName: "query_id",
		},
		{
			Name:       "ResultID",
			SchemaName: "result_id",
		},
		{
			Name:       "HideViaBot",
			SchemaName: "hide_via_bot",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AddQuickReplyShortcutInlineQueryResultMessageRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addQuickReplyShortcutInlineQueryResultMessage#87c02a04 as nil")
	}
	b.PutID(AddQuickReplyShortcutInlineQueryResultMessageRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AddQuickReplyShortcutInlineQueryResultMessageRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addQuickReplyShortcutInlineQueryResultMessage#87c02a04 as nil")
	}
	b.PutString(a.ShortcutName)
	b.PutInt53(a.ReplyToMessageID)
	b.PutLong(a.QueryID)
	b.PutString(a.ResultID)
	b.PutBool(a.HideViaBot)
	return nil
}

// Decode implements bin.Decoder.
func (a *AddQuickReplyShortcutInlineQueryResultMessageRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04 to nil")
	}
	if err := b.ConsumeID(AddQuickReplyShortcutInlineQueryResultMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AddQuickReplyShortcutInlineQueryResultMessageRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04: field shortcut_name: %w", err)
		}
		a.ShortcutName = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04: field reply_to_message_id: %w", err)
		}
		a.ReplyToMessageID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04: field query_id: %w", err)
		}
		a.QueryID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04: field result_id: %w", err)
		}
		a.ResultID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04: field hide_via_bot: %w", err)
		}
		a.HideViaBot = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AddQuickReplyShortcutInlineQueryResultMessageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode addQuickReplyShortcutInlineQueryResultMessage#87c02a04 as nil")
	}
	b.ObjStart()
	b.PutID("addQuickReplyShortcutInlineQueryResultMessage")
	b.Comma()
	b.FieldStart("shortcut_name")
	b.PutString(a.ShortcutName)
	b.Comma()
	b.FieldStart("reply_to_message_id")
	b.PutInt53(a.ReplyToMessageID)
	b.Comma()
	b.FieldStart("query_id")
	b.PutLong(a.QueryID)
	b.Comma()
	b.FieldStart("result_id")
	b.PutString(a.ResultID)
	b.Comma()
	b.FieldStart("hide_via_bot")
	b.PutBool(a.HideViaBot)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AddQuickReplyShortcutInlineQueryResultMessageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("addQuickReplyShortcutInlineQueryResultMessage"); err != nil {
				return fmt.Errorf("unable to decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04: %w", err)
			}
		case "shortcut_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04: field shortcut_name: %w", err)
			}
			a.ShortcutName = value
		case "reply_to_message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04: field reply_to_message_id: %w", err)
			}
			a.ReplyToMessageID = value
		case "query_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04: field query_id: %w", err)
			}
			a.QueryID = value
		case "result_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04: field result_id: %w", err)
			}
			a.ResultID = value
		case "hide_via_bot":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode addQuickReplyShortcutInlineQueryResultMessage#87c02a04: field hide_via_bot: %w", err)
			}
			a.HideViaBot = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetShortcutName returns value of ShortcutName field.
func (a *AddQuickReplyShortcutInlineQueryResultMessageRequest) GetShortcutName() (value string) {
	if a == nil {
		return
	}
	return a.ShortcutName
}

// GetReplyToMessageID returns value of ReplyToMessageID field.
func (a *AddQuickReplyShortcutInlineQueryResultMessageRequest) GetReplyToMessageID() (value int64) {
	if a == nil {
		return
	}
	return a.ReplyToMessageID
}

// GetQueryID returns value of QueryID field.
func (a *AddQuickReplyShortcutInlineQueryResultMessageRequest) GetQueryID() (value int64) {
	if a == nil {
		return
	}
	return a.QueryID
}

// GetResultID returns value of ResultID field.
func (a *AddQuickReplyShortcutInlineQueryResultMessageRequest) GetResultID() (value string) {
	if a == nil {
		return
	}
	return a.ResultID
}

// GetHideViaBot returns value of HideViaBot field.
func (a *AddQuickReplyShortcutInlineQueryResultMessageRequest) GetHideViaBot() (value bool) {
	if a == nil {
		return
	}
	return a.HideViaBot
}

// AddQuickReplyShortcutInlineQueryResultMessage invokes method addQuickReplyShortcutInlineQueryResultMessage#87c02a04 returning error if any.
func (c *Client) AddQuickReplyShortcutInlineQueryResultMessage(ctx context.Context, request *AddQuickReplyShortcutInlineQueryResultMessageRequest) (*QuickReplyMessage, error) {
	var result QuickReplyMessage

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
