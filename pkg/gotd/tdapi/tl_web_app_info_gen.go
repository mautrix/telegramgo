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

// WebAppInfo represents TL type `webAppInfo#2efdb2e8`.
type WebAppInfo struct {
	// Unique identifier for the Web App launch
	LaunchID int64
	// A Web App URL to open in a web view
	URL string
}

// WebAppInfoTypeID is TL type id of WebAppInfo.
const WebAppInfoTypeID = 0x2efdb2e8

// Ensuring interfaces in compile-time for WebAppInfo.
var (
	_ bin.Encoder     = &WebAppInfo{}
	_ bin.Decoder     = &WebAppInfo{}
	_ bin.BareEncoder = &WebAppInfo{}
	_ bin.BareDecoder = &WebAppInfo{}
)

func (w *WebAppInfo) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.LaunchID == 0) {
		return false
	}
	if !(w.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebAppInfo) String() string {
	if w == nil {
		return "WebAppInfo(nil)"
	}
	type Alias WebAppInfo
	return fmt.Sprintf("WebAppInfo%+v", Alias(*w))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebAppInfo) TypeID() uint32 {
	return WebAppInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*WebAppInfo) TypeName() string {
	return "webAppInfo"
}

// TypeInfo returns info about TL type.
func (w *WebAppInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webAppInfo",
		ID:   WebAppInfoTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LaunchID",
			SchemaName: "launch_id",
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *WebAppInfo) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webAppInfo#2efdb2e8 as nil")
	}
	b.PutID(WebAppInfoTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WebAppInfo) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webAppInfo#2efdb2e8 as nil")
	}
	b.PutLong(w.LaunchID)
	b.PutString(w.URL)
	return nil
}

// Decode implements bin.Decoder.
func (w *WebAppInfo) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webAppInfo#2efdb2e8 to nil")
	}
	if err := b.ConsumeID(WebAppInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode webAppInfo#2efdb2e8: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WebAppInfo) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webAppInfo#2efdb2e8 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode webAppInfo#2efdb2e8: field launch_id: %w", err)
		}
		w.LaunchID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webAppInfo#2efdb2e8: field url: %w", err)
		}
		w.URL = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (w *WebAppInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if w == nil {
		return fmt.Errorf("can't encode webAppInfo#2efdb2e8 as nil")
	}
	b.ObjStart()
	b.PutID("webAppInfo")
	b.Comma()
	b.FieldStart("launch_id")
	b.PutLong(w.LaunchID)
	b.Comma()
	b.FieldStart("url")
	b.PutString(w.URL)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (w *WebAppInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if w == nil {
		return fmt.Errorf("can't decode webAppInfo#2efdb2e8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("webAppInfo"); err != nil {
				return fmt.Errorf("unable to decode webAppInfo#2efdb2e8: %w", err)
			}
		case "launch_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode webAppInfo#2efdb2e8: field launch_id: %w", err)
			}
			w.LaunchID = value
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode webAppInfo#2efdb2e8: field url: %w", err)
			}
			w.URL = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLaunchID returns value of LaunchID field.
func (w *WebAppInfo) GetLaunchID() (value int64) {
	if w == nil {
		return
	}
	return w.LaunchID
}

// GetURL returns value of URL field.
func (w *WebAppInfo) GetURL() (value string) {
	if w == nil {
		return
	}
	return w.URL
}
