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

// WebViewMessageSent represents TL type `webViewMessageSent#c94511c`.
// Info about a sent inline webview message
//
// See https://core.telegram.org/constructor/webViewMessageSent for reference.
type WebViewMessageSent struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Message ID
	//
	// Use SetMsgID and GetMsgID helpers.
	MsgID InputBotInlineMessageIDClass
}

// WebViewMessageSentTypeID is TL type id of WebViewMessageSent.
const WebViewMessageSentTypeID = 0xc94511c

// Ensuring interfaces in compile-time for WebViewMessageSent.
var (
	_ bin.Encoder     = &WebViewMessageSent{}
	_ bin.Decoder     = &WebViewMessageSent{}
	_ bin.BareEncoder = &WebViewMessageSent{}
	_ bin.BareDecoder = &WebViewMessageSent{}
)

func (w *WebViewMessageSent) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Flags.Zero()) {
		return false
	}
	if !(w.MsgID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebViewMessageSent) String() string {
	if w == nil {
		return "WebViewMessageSent(nil)"
	}
	type Alias WebViewMessageSent
	return fmt.Sprintf("WebViewMessageSent%+v", Alias(*w))
}

// FillFrom fills WebViewMessageSent from given interface.
func (w *WebViewMessageSent) FillFrom(from interface {
	GetMsgID() (value InputBotInlineMessageIDClass, ok bool)
}) {
	if val, ok := from.GetMsgID(); ok {
		w.MsgID = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebViewMessageSent) TypeID() uint32 {
	return WebViewMessageSentTypeID
}

// TypeName returns name of type in TL schema.
func (*WebViewMessageSent) TypeName() string {
	return "webViewMessageSent"
}

// TypeInfo returns info about TL type.
func (w *WebViewMessageSent) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webViewMessageSent",
		ID:   WebViewMessageSentTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
			Null:       !w.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (w *WebViewMessageSent) SetFlags() {
	if !(w.MsgID == nil) {
		w.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (w *WebViewMessageSent) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webViewMessageSent#c94511c as nil")
	}
	b.PutID(WebViewMessageSentTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WebViewMessageSent) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webViewMessageSent#c94511c as nil")
	}
	w.SetFlags()
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webViewMessageSent#c94511c: field flags: %w", err)
	}
	if w.Flags.Has(0) {
		if w.MsgID == nil {
			return fmt.Errorf("unable to encode webViewMessageSent#c94511c: field msg_id is nil")
		}
		if err := w.MsgID.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webViewMessageSent#c94511c: field msg_id: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *WebViewMessageSent) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webViewMessageSent#c94511c to nil")
	}
	if err := b.ConsumeID(WebViewMessageSentTypeID); err != nil {
		return fmt.Errorf("unable to decode webViewMessageSent#c94511c: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WebViewMessageSent) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webViewMessageSent#c94511c to nil")
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webViewMessageSent#c94511c: field flags: %w", err)
		}
	}
	if w.Flags.Has(0) {
		value, err := DecodeInputBotInlineMessageID(b)
		if err != nil {
			return fmt.Errorf("unable to decode webViewMessageSent#c94511c: field msg_id: %w", err)
		}
		w.MsgID = value
	}
	return nil
}

// SetMsgID sets value of MsgID conditional field.
func (w *WebViewMessageSent) SetMsgID(value InputBotInlineMessageIDClass) {
	w.Flags.Set(0)
	w.MsgID = value
}

// GetMsgID returns value of MsgID conditional field and
// boolean which is true if field was set.
func (w *WebViewMessageSent) GetMsgID() (value InputBotInlineMessageIDClass, ok bool) {
	if w == nil {
		return
	}
	if !w.Flags.Has(0) {
		return value, false
	}
	return w.MsgID, true
}
