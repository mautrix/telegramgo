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

// RtmpURL represents TL type `rtmpUrl#3c28bc55`.
type RtmpURL struct {
	// The URL
	URL string
	// Stream key
	StreamKey string
}

// RtmpURLTypeID is TL type id of RtmpURL.
const RtmpURLTypeID = 0x3c28bc55

// Ensuring interfaces in compile-time for RtmpURL.
var (
	_ bin.Encoder     = &RtmpURL{}
	_ bin.Decoder     = &RtmpURL{}
	_ bin.BareEncoder = &RtmpURL{}
	_ bin.BareDecoder = &RtmpURL{}
)

func (r *RtmpURL) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.StreamKey == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RtmpURL) String() string {
	if r == nil {
		return "RtmpURL(nil)"
	}
	type Alias RtmpURL
	return fmt.Sprintf("RtmpURL%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RtmpURL) TypeID() uint32 {
	return RtmpURLTypeID
}

// TypeName returns name of type in TL schema.
func (*RtmpURL) TypeName() string {
	return "rtmpUrl"
}

// TypeInfo returns info about TL type.
func (r *RtmpURL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "rtmpUrl",
		ID:   RtmpURLTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "StreamKey",
			SchemaName: "stream_key",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RtmpURL) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rtmpUrl#3c28bc55 as nil")
	}
	b.PutID(RtmpURLTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RtmpURL) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rtmpUrl#3c28bc55 as nil")
	}
	b.PutString(r.URL)
	b.PutString(r.StreamKey)
	return nil
}

// Decode implements bin.Decoder.
func (r *RtmpURL) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rtmpUrl#3c28bc55 to nil")
	}
	if err := b.ConsumeID(RtmpURLTypeID); err != nil {
		return fmt.Errorf("unable to decode rtmpUrl#3c28bc55: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RtmpURL) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rtmpUrl#3c28bc55 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode rtmpUrl#3c28bc55: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode rtmpUrl#3c28bc55: field stream_key: %w", err)
		}
		r.StreamKey = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RtmpURL) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode rtmpUrl#3c28bc55 as nil")
	}
	b.ObjStart()
	b.PutID("rtmpUrl")
	b.Comma()
	b.FieldStart("url")
	b.PutString(r.URL)
	b.Comma()
	b.FieldStart("stream_key")
	b.PutString(r.StreamKey)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RtmpURL) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode rtmpUrl#3c28bc55 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("rtmpUrl"); err != nil {
				return fmt.Errorf("unable to decode rtmpUrl#3c28bc55: %w", err)
			}
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode rtmpUrl#3c28bc55: field url: %w", err)
			}
			r.URL = value
		case "stream_key":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode rtmpUrl#3c28bc55: field stream_key: %w", err)
			}
			r.StreamKey = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetURL returns value of URL field.
func (r *RtmpURL) GetURL() (value string) {
	if r == nil {
		return
	}
	return r.URL
}

// GetStreamKey returns value of StreamKey field.
func (r *RtmpURL) GetStreamKey() (value string) {
	if r == nil {
		return
	}
	return r.StreamKey
}
