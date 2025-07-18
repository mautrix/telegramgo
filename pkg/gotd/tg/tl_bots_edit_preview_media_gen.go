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

// BotsEditPreviewMediaRequest represents TL type `bots.editPreviewMedia#8525606f`.
// Edit a main mini app preview, see here »¹ for more info.
// Only owners of bots with a configured Main Mini App can use this method, see see here
// »¹ for more info on how to check if you can invoke this method.
//
// Links:
//  1. https://core.telegram.org/api/bots/webapps#main-mini-app-previews
//  2. https://core.telegram.org/api/bots/webapps#main-mini-app-previews
//
// See https://core.telegram.org/method/bots.editPreviewMedia for reference.
type BotsEditPreviewMediaRequest struct {
	// The bot that owns the Main Mini App.
	Bot InputUserClass
	// ISO 639-1 language code, indicating the localization of the preview to edit.
	LangCode string
	// The photo/video preview to replace, previously fetched as specified here »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/webapps#main-mini-app-previews
	Media InputMediaClass
	// The new photo/video preview, uploaded using messages.uploadMedia¹.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.uploadMedia
	NewMedia InputMediaClass
}

// BotsEditPreviewMediaRequestTypeID is TL type id of BotsEditPreviewMediaRequest.
const BotsEditPreviewMediaRequestTypeID = 0x8525606f

// Ensuring interfaces in compile-time for BotsEditPreviewMediaRequest.
var (
	_ bin.Encoder     = &BotsEditPreviewMediaRequest{}
	_ bin.Decoder     = &BotsEditPreviewMediaRequest{}
	_ bin.BareEncoder = &BotsEditPreviewMediaRequest{}
	_ bin.BareDecoder = &BotsEditPreviewMediaRequest{}
)

func (e *BotsEditPreviewMediaRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Bot == nil) {
		return false
	}
	if !(e.LangCode == "") {
		return false
	}
	if !(e.Media == nil) {
		return false
	}
	if !(e.NewMedia == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *BotsEditPreviewMediaRequest) String() string {
	if e == nil {
		return "BotsEditPreviewMediaRequest(nil)"
	}
	type Alias BotsEditPreviewMediaRequest
	return fmt.Sprintf("BotsEditPreviewMediaRequest%+v", Alias(*e))
}

// FillFrom fills BotsEditPreviewMediaRequest from given interface.
func (e *BotsEditPreviewMediaRequest) FillFrom(from interface {
	GetBot() (value InputUserClass)
	GetLangCode() (value string)
	GetMedia() (value InputMediaClass)
	GetNewMedia() (value InputMediaClass)
}) {
	e.Bot = from.GetBot()
	e.LangCode = from.GetLangCode()
	e.Media = from.GetMedia()
	e.NewMedia = from.GetNewMedia()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsEditPreviewMediaRequest) TypeID() uint32 {
	return BotsEditPreviewMediaRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsEditPreviewMediaRequest) TypeName() string {
	return "bots.editPreviewMedia"
}

// TypeInfo returns info about TL type.
func (e *BotsEditPreviewMediaRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.editPreviewMedia",
		ID:   BotsEditPreviewMediaRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
		{
			Name:       "Media",
			SchemaName: "media",
		},
		{
			Name:       "NewMedia",
			SchemaName: "new_media",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *BotsEditPreviewMediaRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode bots.editPreviewMedia#8525606f as nil")
	}
	b.PutID(BotsEditPreviewMediaRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *BotsEditPreviewMediaRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode bots.editPreviewMedia#8525606f as nil")
	}
	if e.Bot == nil {
		return fmt.Errorf("unable to encode bots.editPreviewMedia#8525606f: field bot is nil")
	}
	if err := e.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.editPreviewMedia#8525606f: field bot: %w", err)
	}
	b.PutString(e.LangCode)
	if e.Media == nil {
		return fmt.Errorf("unable to encode bots.editPreviewMedia#8525606f: field media is nil")
	}
	if err := e.Media.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.editPreviewMedia#8525606f: field media: %w", err)
	}
	if e.NewMedia == nil {
		return fmt.Errorf("unable to encode bots.editPreviewMedia#8525606f: field new_media is nil")
	}
	if err := e.NewMedia.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.editPreviewMedia#8525606f: field new_media: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *BotsEditPreviewMediaRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode bots.editPreviewMedia#8525606f to nil")
	}
	if err := b.ConsumeID(BotsEditPreviewMediaRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.editPreviewMedia#8525606f: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *BotsEditPreviewMediaRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode bots.editPreviewMedia#8525606f to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.editPreviewMedia#8525606f: field bot: %w", err)
		}
		e.Bot = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.editPreviewMedia#8525606f: field lang_code: %w", err)
		}
		e.LangCode = value
	}
	{
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.editPreviewMedia#8525606f: field media: %w", err)
		}
		e.Media = value
	}
	{
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.editPreviewMedia#8525606f: field new_media: %w", err)
		}
		e.NewMedia = value
	}
	return nil
}

// GetBot returns value of Bot field.
func (e *BotsEditPreviewMediaRequest) GetBot() (value InputUserClass) {
	if e == nil {
		return
	}
	return e.Bot
}

// GetLangCode returns value of LangCode field.
func (e *BotsEditPreviewMediaRequest) GetLangCode() (value string) {
	if e == nil {
		return
	}
	return e.LangCode
}

// GetMedia returns value of Media field.
func (e *BotsEditPreviewMediaRequest) GetMedia() (value InputMediaClass) {
	if e == nil {
		return
	}
	return e.Media
}

// GetNewMedia returns value of NewMedia field.
func (e *BotsEditPreviewMediaRequest) GetNewMedia() (value InputMediaClass) {
	if e == nil {
		return
	}
	return e.NewMedia
}

// BotsEditPreviewMedia invokes method bots.editPreviewMedia#8525606f returning error if any.
// Edit a main mini app preview, see here »¹ for more info.
// Only owners of bots with a configured Main Mini App can use this method, see see here
// »¹ for more info on how to check if you can invoke this method.
//
// Links:
//  1. https://core.telegram.org/api/bots/webapps#main-mini-app-previews
//  2. https://core.telegram.org/api/bots/webapps#main-mini-app-previews
//
// Possible errors:
//
//	400 BOT_INVALID: This is not a valid bot.
//
// See https://core.telegram.org/method/bots.editPreviewMedia for reference.
func (c *Client) BotsEditPreviewMedia(ctx context.Context, request *BotsEditPreviewMediaRequest) (*BotPreviewMedia, error) {
	var result BotPreviewMedia

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
