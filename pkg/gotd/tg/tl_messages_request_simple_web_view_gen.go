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

// MessagesRequestSimpleWebViewRequest represents TL type `messages.requestSimpleWebView#413a3e73`.
// Open a bot mini app¹.
//
// Links:
//  1. https://core.telegram.org/api/bots/webapps
//
// See https://core.telegram.org/method/messages.requestSimpleWebView for reference.
type MessagesRequestSimpleWebViewRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the webapp was opened by clicking on the switch_webview button shown on top of
	// the inline results list returned by messages.getInlineBotResults¹.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.getInlineBotResults
	FromSwitchWebview bool
	// Set this flag if opening the Mini App from the installed side menu entry »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/attach
	FromSideMenu bool
	// Deprecated.
	Compact bool
	// Fullscreen field of MessagesRequestSimpleWebViewRequest.
	Fullscreen bool
	// Bot that owns the mini app
	Bot InputUserClass
	// Web app URL, if opening from a keyboard button or inline result
	//
	// Use SetURL and GetURL helpers.
	URL string
	// Deprecated.
	//
	// Use SetStartParam and GetStartParam helpers.
	StartParam string
	// Theme parameters »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/webapps#theme-parameters
	//
	// Use SetThemeParams and GetThemeParams helpers.
	ThemeParams DataJSON
	// Short name of the application; 0-64 English letters, digits, and underscores
	Platform string
}

// MessagesRequestSimpleWebViewRequestTypeID is TL type id of MessagesRequestSimpleWebViewRequest.
const MessagesRequestSimpleWebViewRequestTypeID = 0x413a3e73

// Ensuring interfaces in compile-time for MessagesRequestSimpleWebViewRequest.
var (
	_ bin.Encoder     = &MessagesRequestSimpleWebViewRequest{}
	_ bin.Decoder     = &MessagesRequestSimpleWebViewRequest{}
	_ bin.BareEncoder = &MessagesRequestSimpleWebViewRequest{}
	_ bin.BareDecoder = &MessagesRequestSimpleWebViewRequest{}
)

func (r *MessagesRequestSimpleWebViewRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.FromSwitchWebview == false) {
		return false
	}
	if !(r.FromSideMenu == false) {
		return false
	}
	if !(r.Compact == false) {
		return false
	}
	if !(r.Fullscreen == false) {
		return false
	}
	if !(r.Bot == nil) {
		return false
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.StartParam == "") {
		return false
	}
	if !(r.ThemeParams.Zero()) {
		return false
	}
	if !(r.Platform == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesRequestSimpleWebViewRequest) String() string {
	if r == nil {
		return "MessagesRequestSimpleWebViewRequest(nil)"
	}
	type Alias MessagesRequestSimpleWebViewRequest
	return fmt.Sprintf("MessagesRequestSimpleWebViewRequest%+v", Alias(*r))
}

// FillFrom fills MessagesRequestSimpleWebViewRequest from given interface.
func (r *MessagesRequestSimpleWebViewRequest) FillFrom(from interface {
	GetFromSwitchWebview() (value bool)
	GetFromSideMenu() (value bool)
	GetCompact() (value bool)
	GetFullscreen() (value bool)
	GetBot() (value InputUserClass)
	GetURL() (value string, ok bool)
	GetStartParam() (value string, ok bool)
	GetThemeParams() (value DataJSON, ok bool)
	GetPlatform() (value string)
}) {
	r.FromSwitchWebview = from.GetFromSwitchWebview()
	r.FromSideMenu = from.GetFromSideMenu()
	r.Compact = from.GetCompact()
	r.Fullscreen = from.GetFullscreen()
	r.Bot = from.GetBot()
	if val, ok := from.GetURL(); ok {
		r.URL = val
	}

	if val, ok := from.GetStartParam(); ok {
		r.StartParam = val
	}

	if val, ok := from.GetThemeParams(); ok {
		r.ThemeParams = val
	}

	r.Platform = from.GetPlatform()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesRequestSimpleWebViewRequest) TypeID() uint32 {
	return MessagesRequestSimpleWebViewRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesRequestSimpleWebViewRequest) TypeName() string {
	return "messages.requestSimpleWebView"
}

// TypeInfo returns info about TL type.
func (r *MessagesRequestSimpleWebViewRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.requestSimpleWebView",
		ID:   MessagesRequestSimpleWebViewRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FromSwitchWebview",
			SchemaName: "from_switch_webview",
			Null:       !r.Flags.Has(1),
		},
		{
			Name:       "FromSideMenu",
			SchemaName: "from_side_menu",
			Null:       !r.Flags.Has(2),
		},
		{
			Name:       "Compact",
			SchemaName: "compact",
			Null:       !r.Flags.Has(7),
		},
		{
			Name:       "Fullscreen",
			SchemaName: "fullscreen",
			Null:       !r.Flags.Has(8),
		},
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "URL",
			SchemaName: "url",
			Null:       !r.Flags.Has(3),
		},
		{
			Name:       "StartParam",
			SchemaName: "start_param",
			Null:       !r.Flags.Has(4),
		},
		{
			Name:       "ThemeParams",
			SchemaName: "theme_params",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "Platform",
			SchemaName: "platform",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *MessagesRequestSimpleWebViewRequest) SetFlags() {
	if !(r.FromSwitchWebview == false) {
		r.Flags.Set(1)
	}
	if !(r.FromSideMenu == false) {
		r.Flags.Set(2)
	}
	if !(r.Compact == false) {
		r.Flags.Set(7)
	}
	if !(r.Fullscreen == false) {
		r.Flags.Set(8)
	}
	if !(r.URL == "") {
		r.Flags.Set(3)
	}
	if !(r.StartParam == "") {
		r.Flags.Set(4)
	}
	if !(r.ThemeParams.Zero()) {
		r.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (r *MessagesRequestSimpleWebViewRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestSimpleWebView#413a3e73 as nil")
	}
	b.PutID(MessagesRequestSimpleWebViewRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesRequestSimpleWebViewRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestSimpleWebView#413a3e73 as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestSimpleWebView#413a3e73: field flags: %w", err)
	}
	if r.Bot == nil {
		return fmt.Errorf("unable to encode messages.requestSimpleWebView#413a3e73: field bot is nil")
	}
	if err := r.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestSimpleWebView#413a3e73: field bot: %w", err)
	}
	if r.Flags.Has(3) {
		b.PutString(r.URL)
	}
	if r.Flags.Has(4) {
		b.PutString(r.StartParam)
	}
	if r.Flags.Has(0) {
		if err := r.ThemeParams.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.requestSimpleWebView#413a3e73: field theme_params: %w", err)
		}
	}
	b.PutString(r.Platform)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesRequestSimpleWebViewRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestSimpleWebView#413a3e73 to nil")
	}
	if err := b.ConsumeID(MessagesRequestSimpleWebViewRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.requestSimpleWebView#413a3e73: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesRequestSimpleWebViewRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestSimpleWebView#413a3e73 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.requestSimpleWebView#413a3e73: field flags: %w", err)
		}
	}
	r.FromSwitchWebview = r.Flags.Has(1)
	r.FromSideMenu = r.Flags.Has(2)
	r.Compact = r.Flags.Has(7)
	r.Fullscreen = r.Flags.Has(8)
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestSimpleWebView#413a3e73: field bot: %w", err)
		}
		r.Bot = value
	}
	if r.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestSimpleWebView#413a3e73: field url: %w", err)
		}
		r.URL = value
	}
	if r.Flags.Has(4) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestSimpleWebView#413a3e73: field start_param: %w", err)
		}
		r.StartParam = value
	}
	if r.Flags.Has(0) {
		if err := r.ThemeParams.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.requestSimpleWebView#413a3e73: field theme_params: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestSimpleWebView#413a3e73: field platform: %w", err)
		}
		r.Platform = value
	}
	return nil
}

// SetFromSwitchWebview sets value of FromSwitchWebview conditional field.
func (r *MessagesRequestSimpleWebViewRequest) SetFromSwitchWebview(value bool) {
	if value {
		r.Flags.Set(1)
		r.FromSwitchWebview = true
	} else {
		r.Flags.Unset(1)
		r.FromSwitchWebview = false
	}
}

// GetFromSwitchWebview returns value of FromSwitchWebview conditional field.
func (r *MessagesRequestSimpleWebViewRequest) GetFromSwitchWebview() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(1)
}

// SetFromSideMenu sets value of FromSideMenu conditional field.
func (r *MessagesRequestSimpleWebViewRequest) SetFromSideMenu(value bool) {
	if value {
		r.Flags.Set(2)
		r.FromSideMenu = true
	} else {
		r.Flags.Unset(2)
		r.FromSideMenu = false
	}
}

// GetFromSideMenu returns value of FromSideMenu conditional field.
func (r *MessagesRequestSimpleWebViewRequest) GetFromSideMenu() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(2)
}

// SetCompact sets value of Compact conditional field.
func (r *MessagesRequestSimpleWebViewRequest) SetCompact(value bool) {
	if value {
		r.Flags.Set(7)
		r.Compact = true
	} else {
		r.Flags.Unset(7)
		r.Compact = false
	}
}

// GetCompact returns value of Compact conditional field.
func (r *MessagesRequestSimpleWebViewRequest) GetCompact() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(7)
}

// SetFullscreen sets value of Fullscreen conditional field.
func (r *MessagesRequestSimpleWebViewRequest) SetFullscreen(value bool) {
	if value {
		r.Flags.Set(8)
		r.Fullscreen = true
	} else {
		r.Flags.Unset(8)
		r.Fullscreen = false
	}
}

// GetFullscreen returns value of Fullscreen conditional field.
func (r *MessagesRequestSimpleWebViewRequest) GetFullscreen() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(8)
}

// GetBot returns value of Bot field.
func (r *MessagesRequestSimpleWebViewRequest) GetBot() (value InputUserClass) {
	if r == nil {
		return
	}
	return r.Bot
}

// SetURL sets value of URL conditional field.
func (r *MessagesRequestSimpleWebViewRequest) SetURL(value string) {
	r.Flags.Set(3)
	r.URL = value
}

// GetURL returns value of URL conditional field and
// boolean which is true if field was set.
func (r *MessagesRequestSimpleWebViewRequest) GetURL() (value string, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(3) {
		return value, false
	}
	return r.URL, true
}

// SetStartParam sets value of StartParam conditional field.
func (r *MessagesRequestSimpleWebViewRequest) SetStartParam(value string) {
	r.Flags.Set(4)
	r.StartParam = value
}

// GetStartParam returns value of StartParam conditional field and
// boolean which is true if field was set.
func (r *MessagesRequestSimpleWebViewRequest) GetStartParam() (value string, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(4) {
		return value, false
	}
	return r.StartParam, true
}

// SetThemeParams sets value of ThemeParams conditional field.
func (r *MessagesRequestSimpleWebViewRequest) SetThemeParams(value DataJSON) {
	r.Flags.Set(0)
	r.ThemeParams = value
}

// GetThemeParams returns value of ThemeParams conditional field and
// boolean which is true if field was set.
func (r *MessagesRequestSimpleWebViewRequest) GetThemeParams() (value DataJSON, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(0) {
		return value, false
	}
	return r.ThemeParams, true
}

// GetPlatform returns value of Platform field.
func (r *MessagesRequestSimpleWebViewRequest) GetPlatform() (value string) {
	if r == nil {
		return
	}
	return r.Platform
}

// MessagesRequestSimpleWebView invokes method messages.requestSimpleWebView#413a3e73 returning error if any.
// Open a bot mini app¹.
//
// Links:
//  1. https://core.telegram.org/api/bots/webapps
//
// Possible errors:
//
//	400 BOT_INVALID: This is not a valid bot.
//	400 URL_INVALID: Invalid URL provided.
//
// See https://core.telegram.org/method/messages.requestSimpleWebView for reference.
func (c *Client) MessagesRequestSimpleWebView(ctx context.Context, request *MessagesRequestSimpleWebViewRequest) (*WebViewResultURL, error) {
	var result WebViewResultURL

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
