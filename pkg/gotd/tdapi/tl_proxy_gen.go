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

// Proxy represents TL type `proxy#baf7b73`.
type Proxy struct {
	// Unique identifier of the proxy
	ID int32
	// Proxy server domain or IP address
	Server string
	// Proxy server port
	Port int32
	// Point in time (Unix timestamp) when the proxy was last used; 0 if never
	LastUsedDate int32
	// True, if the proxy is enabled now
	IsEnabled bool
	// Type of the proxy
	Type ProxyTypeClass
}

// ProxyTypeID is TL type id of Proxy.
const ProxyTypeID = 0xbaf7b73

// Ensuring interfaces in compile-time for Proxy.
var (
	_ bin.Encoder     = &Proxy{}
	_ bin.Decoder     = &Proxy{}
	_ bin.BareEncoder = &Proxy{}
	_ bin.BareDecoder = &Proxy{}
)

func (p *Proxy) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ID == 0) {
		return false
	}
	if !(p.Server == "") {
		return false
	}
	if !(p.Port == 0) {
		return false
	}
	if !(p.LastUsedDate == 0) {
		return false
	}
	if !(p.IsEnabled == false) {
		return false
	}
	if !(p.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *Proxy) String() string {
	if p == nil {
		return "Proxy(nil)"
	}
	type Alias Proxy
	return fmt.Sprintf("Proxy%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Proxy) TypeID() uint32 {
	return ProxyTypeID
}

// TypeName returns name of type in TL schema.
func (*Proxy) TypeName() string {
	return "proxy"
}

// TypeInfo returns info about TL type.
func (p *Proxy) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "proxy",
		ID:   ProxyTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Server",
			SchemaName: "server",
		},
		{
			Name:       "Port",
			SchemaName: "port",
		},
		{
			Name:       "LastUsedDate",
			SchemaName: "last_used_date",
		},
		{
			Name:       "IsEnabled",
			SchemaName: "is_enabled",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *Proxy) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode proxy#baf7b73 as nil")
	}
	b.PutID(ProxyTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *Proxy) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode proxy#baf7b73 as nil")
	}
	b.PutInt32(p.ID)
	b.PutString(p.Server)
	b.PutInt32(p.Port)
	b.PutInt32(p.LastUsedDate)
	b.PutBool(p.IsEnabled)
	if p.Type == nil {
		return fmt.Errorf("unable to encode proxy#baf7b73: field type is nil")
	}
	if err := p.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode proxy#baf7b73: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *Proxy) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode proxy#baf7b73 to nil")
	}
	if err := b.ConsumeID(ProxyTypeID); err != nil {
		return fmt.Errorf("unable to decode proxy#baf7b73: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *Proxy) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode proxy#baf7b73 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode proxy#baf7b73: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode proxy#baf7b73: field server: %w", err)
		}
		p.Server = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode proxy#baf7b73: field port: %w", err)
		}
		p.Port = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode proxy#baf7b73: field last_used_date: %w", err)
		}
		p.LastUsedDate = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode proxy#baf7b73: field is_enabled: %w", err)
		}
		p.IsEnabled = value
	}
	{
		value, err := DecodeProxyType(b)
		if err != nil {
			return fmt.Errorf("unable to decode proxy#baf7b73: field type: %w", err)
		}
		p.Type = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *Proxy) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode proxy#baf7b73 as nil")
	}
	b.ObjStart()
	b.PutID("proxy")
	b.Comma()
	b.FieldStart("id")
	b.PutInt32(p.ID)
	b.Comma()
	b.FieldStart("server")
	b.PutString(p.Server)
	b.Comma()
	b.FieldStart("port")
	b.PutInt32(p.Port)
	b.Comma()
	b.FieldStart("last_used_date")
	b.PutInt32(p.LastUsedDate)
	b.Comma()
	b.FieldStart("is_enabled")
	b.PutBool(p.IsEnabled)
	b.Comma()
	b.FieldStart("type")
	if p.Type == nil {
		return fmt.Errorf("unable to encode proxy#baf7b73: field type is nil")
	}
	if err := p.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode proxy#baf7b73: field type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *Proxy) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode proxy#baf7b73 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("proxy"); err != nil {
				return fmt.Errorf("unable to decode proxy#baf7b73: %w", err)
			}
		case "id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode proxy#baf7b73: field id: %w", err)
			}
			p.ID = value
		case "server":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode proxy#baf7b73: field server: %w", err)
			}
			p.Server = value
		case "port":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode proxy#baf7b73: field port: %w", err)
			}
			p.Port = value
		case "last_used_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode proxy#baf7b73: field last_used_date: %w", err)
			}
			p.LastUsedDate = value
		case "is_enabled":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode proxy#baf7b73: field is_enabled: %w", err)
			}
			p.IsEnabled = value
		case "type":
			value, err := DecodeTDLibJSONProxyType(b)
			if err != nil {
				return fmt.Errorf("unable to decode proxy#baf7b73: field type: %w", err)
			}
			p.Type = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (p *Proxy) GetID() (value int32) {
	if p == nil {
		return
	}
	return p.ID
}

// GetServer returns value of Server field.
func (p *Proxy) GetServer() (value string) {
	if p == nil {
		return
	}
	return p.Server
}

// GetPort returns value of Port field.
func (p *Proxy) GetPort() (value int32) {
	if p == nil {
		return
	}
	return p.Port
}

// GetLastUsedDate returns value of LastUsedDate field.
func (p *Proxy) GetLastUsedDate() (value int32) {
	if p == nil {
		return
	}
	return p.LastUsedDate
}

// GetIsEnabled returns value of IsEnabled field.
func (p *Proxy) GetIsEnabled() (value bool) {
	if p == nil {
		return
	}
	return p.IsEnabled
}

// GetType returns value of Type field.
func (p *Proxy) GetType() (value ProxyTypeClass) {
	if p == nil {
		return
	}
	return p.Type
}
