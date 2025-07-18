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

// FoundWebApp represents TL type `foundWebApp#eea8d01e`.
type FoundWebApp struct {
	// The Web App
	WebApp WebApp
	// True, if the user must be asked for the permission to the bot to send them messages
	RequestWriteAccess bool
	// True, if there is no need to show an ordinary open URL confirmation before opening the
	// Web App. The field must be ignored and confirmation must be shown anyway if the Web
	// App link was hidden
	SkipConfirmation bool
}

// FoundWebAppTypeID is TL type id of FoundWebApp.
const FoundWebAppTypeID = 0xeea8d01e

// Ensuring interfaces in compile-time for FoundWebApp.
var (
	_ bin.Encoder     = &FoundWebApp{}
	_ bin.Decoder     = &FoundWebApp{}
	_ bin.BareEncoder = &FoundWebApp{}
	_ bin.BareDecoder = &FoundWebApp{}
)

func (f *FoundWebApp) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.WebApp.Zero()) {
		return false
	}
	if !(f.RequestWriteAccess == false) {
		return false
	}
	if !(f.SkipConfirmation == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FoundWebApp) String() string {
	if f == nil {
		return "FoundWebApp(nil)"
	}
	type Alias FoundWebApp
	return fmt.Sprintf("FoundWebApp%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FoundWebApp) TypeID() uint32 {
	return FoundWebAppTypeID
}

// TypeName returns name of type in TL schema.
func (*FoundWebApp) TypeName() string {
	return "foundWebApp"
}

// TypeInfo returns info about TL type.
func (f *FoundWebApp) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "foundWebApp",
		ID:   FoundWebAppTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "WebApp",
			SchemaName: "web_app",
		},
		{
			Name:       "RequestWriteAccess",
			SchemaName: "request_write_access",
		},
		{
			Name:       "SkipConfirmation",
			SchemaName: "skip_confirmation",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *FoundWebApp) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode foundWebApp#eea8d01e as nil")
	}
	b.PutID(FoundWebAppTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FoundWebApp) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode foundWebApp#eea8d01e as nil")
	}
	if err := f.WebApp.Encode(b); err != nil {
		return fmt.Errorf("unable to encode foundWebApp#eea8d01e: field web_app: %w", err)
	}
	b.PutBool(f.RequestWriteAccess)
	b.PutBool(f.SkipConfirmation)
	return nil
}

// Decode implements bin.Decoder.
func (f *FoundWebApp) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode foundWebApp#eea8d01e to nil")
	}
	if err := b.ConsumeID(FoundWebAppTypeID); err != nil {
		return fmt.Errorf("unable to decode foundWebApp#eea8d01e: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FoundWebApp) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode foundWebApp#eea8d01e to nil")
	}
	{
		if err := f.WebApp.Decode(b); err != nil {
			return fmt.Errorf("unable to decode foundWebApp#eea8d01e: field web_app: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode foundWebApp#eea8d01e: field request_write_access: %w", err)
		}
		f.RequestWriteAccess = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode foundWebApp#eea8d01e: field skip_confirmation: %w", err)
		}
		f.SkipConfirmation = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *FoundWebApp) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode foundWebApp#eea8d01e as nil")
	}
	b.ObjStart()
	b.PutID("foundWebApp")
	b.Comma()
	b.FieldStart("web_app")
	if err := f.WebApp.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode foundWebApp#eea8d01e: field web_app: %w", err)
	}
	b.Comma()
	b.FieldStart("request_write_access")
	b.PutBool(f.RequestWriteAccess)
	b.Comma()
	b.FieldStart("skip_confirmation")
	b.PutBool(f.SkipConfirmation)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (f *FoundWebApp) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode foundWebApp#eea8d01e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("foundWebApp"); err != nil {
				return fmt.Errorf("unable to decode foundWebApp#eea8d01e: %w", err)
			}
		case "web_app":
			if err := f.WebApp.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode foundWebApp#eea8d01e: field web_app: %w", err)
			}
		case "request_write_access":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode foundWebApp#eea8d01e: field request_write_access: %w", err)
			}
			f.RequestWriteAccess = value
		case "skip_confirmation":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode foundWebApp#eea8d01e: field skip_confirmation: %w", err)
			}
			f.SkipConfirmation = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetWebApp returns value of WebApp field.
func (f *FoundWebApp) GetWebApp() (value WebApp) {
	if f == nil {
		return
	}
	return f.WebApp
}

// GetRequestWriteAccess returns value of RequestWriteAccess field.
func (f *FoundWebApp) GetRequestWriteAccess() (value bool) {
	if f == nil {
		return
	}
	return f.RequestWriteAccess
}

// GetSkipConfirmation returns value of SkipConfirmation field.
func (f *FoundWebApp) GetSkipConfirmation() (value bool) {
	if f == nil {
		return
	}
	return f.SkipConfirmation
}
