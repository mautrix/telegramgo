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

// MessagesEditInlineBotMessageRequest represents TL type `messages.editInlineBotMessage#83557dba`.
// Edit an inline bot message
//
// See https://core.telegram.org/method/messages.editInlineBotMessage for reference.
type MessagesEditInlineBotMessageRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Disable webpage preview
	NoWebpage bool
	// If set, any eventual webpage preview will be shown on top of the message instead of at
	// the bottom.
	InvertMedia bool
	// Sent inline message ID
	ID InputBotInlineMessageIDClass
	// Message
	//
	// Use SetMessage and GetMessage helpers.
	Message string
	// Media
	//
	// Use SetMedia and GetMedia helpers.
	Media InputMediaClass
	// Reply markup for inline keyboards
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
}

// MessagesEditInlineBotMessageRequestTypeID is TL type id of MessagesEditInlineBotMessageRequest.
const MessagesEditInlineBotMessageRequestTypeID = 0x83557dba

// Ensuring interfaces in compile-time for MessagesEditInlineBotMessageRequest.
var (
	_ bin.Encoder     = &MessagesEditInlineBotMessageRequest{}
	_ bin.Decoder     = &MessagesEditInlineBotMessageRequest{}
	_ bin.BareEncoder = &MessagesEditInlineBotMessageRequest{}
	_ bin.BareDecoder = &MessagesEditInlineBotMessageRequest{}
)

func (e *MessagesEditInlineBotMessageRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.NoWebpage == false) {
		return false
	}
	if !(e.InvertMedia == false) {
		return false
	}
	if !(e.ID == nil) {
		return false
	}
	if !(e.Message == "") {
		return false
	}
	if !(e.Media == nil) {
		return false
	}
	if !(e.ReplyMarkup == nil) {
		return false
	}
	if !(e.Entities == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesEditInlineBotMessageRequest) String() string {
	if e == nil {
		return "MessagesEditInlineBotMessageRequest(nil)"
	}
	type Alias MessagesEditInlineBotMessageRequest
	return fmt.Sprintf("MessagesEditInlineBotMessageRequest%+v", Alias(*e))
}

// FillFrom fills MessagesEditInlineBotMessageRequest from given interface.
func (e *MessagesEditInlineBotMessageRequest) FillFrom(from interface {
	GetNoWebpage() (value bool)
	GetInvertMedia() (value bool)
	GetID() (value InputBotInlineMessageIDClass)
	GetMessage() (value string, ok bool)
	GetMedia() (value InputMediaClass, ok bool)
	GetReplyMarkup() (value ReplyMarkupClass, ok bool)
	GetEntities() (value []MessageEntityClass, ok bool)
}) {
	e.NoWebpage = from.GetNoWebpage()
	e.InvertMedia = from.GetInvertMedia()
	e.ID = from.GetID()
	if val, ok := from.GetMessage(); ok {
		e.Message = val
	}

	if val, ok := from.GetMedia(); ok {
		e.Media = val
	}

	if val, ok := from.GetReplyMarkup(); ok {
		e.ReplyMarkup = val
	}

	if val, ok := from.GetEntities(); ok {
		e.Entities = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesEditInlineBotMessageRequest) TypeID() uint32 {
	return MessagesEditInlineBotMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesEditInlineBotMessageRequest) TypeName() string {
	return "messages.editInlineBotMessage"
}

// TypeInfo returns info about TL type.
func (e *MessagesEditInlineBotMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.editInlineBotMessage",
		ID:   MessagesEditInlineBotMessageRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "NoWebpage",
			SchemaName: "no_webpage",
			Null:       !e.Flags.Has(1),
		},
		{
			Name:       "InvertMedia",
			SchemaName: "invert_media",
			Null:       !e.Flags.Has(16),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Message",
			SchemaName: "message",
			Null:       !e.Flags.Has(11),
		},
		{
			Name:       "Media",
			SchemaName: "media",
			Null:       !e.Flags.Has(14),
		},
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
			Null:       !e.Flags.Has(2),
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !e.Flags.Has(3),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (e *MessagesEditInlineBotMessageRequest) SetFlags() {
	if !(e.NoWebpage == false) {
		e.Flags.Set(1)
	}
	if !(e.InvertMedia == false) {
		e.Flags.Set(16)
	}
	if !(e.Message == "") {
		e.Flags.Set(11)
	}
	if !(e.Media == nil) {
		e.Flags.Set(14)
	}
	if !(e.ReplyMarkup == nil) {
		e.Flags.Set(2)
	}
	if !(e.Entities == nil) {
		e.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (e *MessagesEditInlineBotMessageRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editInlineBotMessage#83557dba as nil")
	}
	b.PutID(MessagesEditInlineBotMessageRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *MessagesEditInlineBotMessageRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editInlineBotMessage#83557dba as nil")
	}
	e.SetFlags()
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field flags: %w", err)
	}
	if e.ID == nil {
		return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field id is nil")
	}
	if err := e.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field id: %w", err)
	}
	if e.Flags.Has(11) {
		b.PutString(e.Message)
	}
	if e.Flags.Has(14) {
		if e.Media == nil {
			return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field media is nil")
		}
		if err := e.Media.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field media: %w", err)
		}
	}
	if e.Flags.Has(2) {
		if e.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field reply_markup is nil")
		}
		if err := e.ReplyMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field reply_markup: %w", err)
		}
	}
	if e.Flags.Has(3) {
		b.PutVectorHeader(len(e.Entities))
		for idx, v := range e.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field entities element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *MessagesEditInlineBotMessageRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editInlineBotMessage#83557dba to nil")
	}
	if err := b.ConsumeID(MessagesEditInlineBotMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *MessagesEditInlineBotMessageRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editInlineBotMessage#83557dba to nil")
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: field flags: %w", err)
		}
	}
	e.NoWebpage = e.Flags.Has(1)
	e.InvertMedia = e.Flags.Has(16)
	{
		value, err := DecodeInputBotInlineMessageID(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: field id: %w", err)
		}
		e.ID = value
	}
	if e.Flags.Has(11) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: field message: %w", err)
		}
		e.Message = value
	}
	if e.Flags.Has(14) {
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: field media: %w", err)
		}
		e.Media = value
	}
	if e.Flags.Has(2) {
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: field reply_markup: %w", err)
		}
		e.ReplyMarkup = value
	}
	if e.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: field entities: %w", err)
		}

		if headerLen > 0 {
			e.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: field entities: %w", err)
			}
			e.Entities = append(e.Entities, value)
		}
	}
	return nil
}

// SetNoWebpage sets value of NoWebpage conditional field.
func (e *MessagesEditInlineBotMessageRequest) SetNoWebpage(value bool) {
	if value {
		e.Flags.Set(1)
		e.NoWebpage = true
	} else {
		e.Flags.Unset(1)
		e.NoWebpage = false
	}
}

// GetNoWebpage returns value of NoWebpage conditional field.
func (e *MessagesEditInlineBotMessageRequest) GetNoWebpage() (value bool) {
	if e == nil {
		return
	}
	return e.Flags.Has(1)
}

// SetInvertMedia sets value of InvertMedia conditional field.
func (e *MessagesEditInlineBotMessageRequest) SetInvertMedia(value bool) {
	if value {
		e.Flags.Set(16)
		e.InvertMedia = true
	} else {
		e.Flags.Unset(16)
		e.InvertMedia = false
	}
}

// GetInvertMedia returns value of InvertMedia conditional field.
func (e *MessagesEditInlineBotMessageRequest) GetInvertMedia() (value bool) {
	if e == nil {
		return
	}
	return e.Flags.Has(16)
}

// GetID returns value of ID field.
func (e *MessagesEditInlineBotMessageRequest) GetID() (value InputBotInlineMessageIDClass) {
	if e == nil {
		return
	}
	return e.ID
}

// SetMessage sets value of Message conditional field.
func (e *MessagesEditInlineBotMessageRequest) SetMessage(value string) {
	e.Flags.Set(11)
	e.Message = value
}

// GetMessage returns value of Message conditional field and
// boolean which is true if field was set.
func (e *MessagesEditInlineBotMessageRequest) GetMessage() (value string, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(11) {
		return value, false
	}
	return e.Message, true
}

// SetMedia sets value of Media conditional field.
func (e *MessagesEditInlineBotMessageRequest) SetMedia(value InputMediaClass) {
	e.Flags.Set(14)
	e.Media = value
}

// GetMedia returns value of Media conditional field and
// boolean which is true if field was set.
func (e *MessagesEditInlineBotMessageRequest) GetMedia() (value InputMediaClass, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(14) {
		return value, false
	}
	return e.Media, true
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (e *MessagesEditInlineBotMessageRequest) SetReplyMarkup(value ReplyMarkupClass) {
	e.Flags.Set(2)
	e.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (e *MessagesEditInlineBotMessageRequest) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(2) {
		return value, false
	}
	return e.ReplyMarkup, true
}

// SetEntities sets value of Entities conditional field.
func (e *MessagesEditInlineBotMessageRequest) SetEntities(value []MessageEntityClass) {
	e.Flags.Set(3)
	e.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (e *MessagesEditInlineBotMessageRequest) GetEntities() (value []MessageEntityClass, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(3) {
		return value, false
	}
	return e.Entities, true
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (e *MessagesEditInlineBotMessageRequest) MapEntities() (value MessageEntityClassArray, ok bool) {
	if !e.Flags.Has(3) {
		return value, false
	}
	return MessageEntityClassArray(e.Entities), true
}

// MessagesEditInlineBotMessage invokes method messages.editInlineBotMessage#83557dba returning error if any.
// Edit an inline bot message
//
// Possible errors:
//
//	400 BUTTON_DATA_INVALID: The data of one or more of the buttons you provided is invalid.
//	400 ENTITY_BOUNDS_INVALID: A specified entity offset or length is invalid, see here » for info on how to properly compute the entity offset/length.
//	400 MESSAGE_ID_INVALID: The provided message id is invalid.
//	400 MESSAGE_NOT_MODIFIED: The provided message data is identical to the previous message data, the message wasn't modified.
//
// See https://core.telegram.org/method/messages.editInlineBotMessage for reference.
// Can be used by bots.
func (c *Client) MessagesEditInlineBotMessage(ctx context.Context, request *MessagesEditInlineBotMessageRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
