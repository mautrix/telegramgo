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

// AddProxyRequest represents TL type `addProxy#13c2bcd8`.
type AddProxyRequest struct {
	// Proxy server domain or IP address
	Server string
	// Proxy server port
	Port int32
	// Pass true to immediately enable the proxy
	Enable bool
	// Proxy type
	Type ProxyTypeClass
}

// AddProxyRequestTypeID is TL type id of AddProxyRequest.
const AddProxyRequestTypeID = 0x13c2bcd8

// Ensuring interfaces in compile-time for AddProxyRequest.
var (
	_ bin.Encoder     = &AddProxyRequest{}
	_ bin.Decoder     = &AddProxyRequest{}
	_ bin.BareEncoder = &AddProxyRequest{}
	_ bin.BareDecoder = &AddProxyRequest{}
)

func (a *AddProxyRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Server == "") {
		return false
	}
	if !(a.Port == 0) {
		return false
	}
	if !(a.Enable == false) {
		return false
	}
	if !(a.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AddProxyRequest) String() string {
	if a == nil {
		return "AddProxyRequest(nil)"
	}
	type Alias AddProxyRequest
	return fmt.Sprintf("AddProxyRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AddProxyRequest) TypeID() uint32 {
	return AddProxyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AddProxyRequest) TypeName() string {
	return "addProxy"
}

// TypeInfo returns info about TL type.
func (a *AddProxyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "addProxy",
		ID:   AddProxyRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Server",
			SchemaName: "server",
		},
		{
			Name:       "Port",
			SchemaName: "port",
		},
		{
			Name:       "Enable",
			SchemaName: "enable",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AddProxyRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addProxy#13c2bcd8 as nil")
	}
	b.PutID(AddProxyRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AddProxyRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addProxy#13c2bcd8 as nil")
	}
	b.PutString(a.Server)
	b.PutInt32(a.Port)
	b.PutBool(a.Enable)
	if a.Type == nil {
		return fmt.Errorf("unable to encode addProxy#13c2bcd8: field type is nil")
	}
	if err := a.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode addProxy#13c2bcd8: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AddProxyRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addProxy#13c2bcd8 to nil")
	}
	if err := b.ConsumeID(AddProxyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode addProxy#13c2bcd8: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AddProxyRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addProxy#13c2bcd8 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode addProxy#13c2bcd8: field server: %w", err)
		}
		a.Server = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode addProxy#13c2bcd8: field port: %w", err)
		}
		a.Port = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode addProxy#13c2bcd8: field enable: %w", err)
		}
		a.Enable = value
	}
	{
		value, err := DecodeProxyType(b)
		if err != nil {
			return fmt.Errorf("unable to decode addProxy#13c2bcd8: field type: %w", err)
		}
		a.Type = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AddProxyRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode addProxy#13c2bcd8 as nil")
	}
	b.ObjStart()
	b.PutID("addProxy")
	b.Comma()
	b.FieldStart("server")
	b.PutString(a.Server)
	b.Comma()
	b.FieldStart("port")
	b.PutInt32(a.Port)
	b.Comma()
	b.FieldStart("enable")
	b.PutBool(a.Enable)
	b.Comma()
	b.FieldStart("type")
	if a.Type == nil {
		return fmt.Errorf("unable to encode addProxy#13c2bcd8: field type is nil")
	}
	if err := a.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode addProxy#13c2bcd8: field type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AddProxyRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode addProxy#13c2bcd8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("addProxy"); err != nil {
				return fmt.Errorf("unable to decode addProxy#13c2bcd8: %w", err)
			}
		case "server":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode addProxy#13c2bcd8: field server: %w", err)
			}
			a.Server = value
		case "port":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode addProxy#13c2bcd8: field port: %w", err)
			}
			a.Port = value
		case "enable":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode addProxy#13c2bcd8: field enable: %w", err)
			}
			a.Enable = value
		case "type":
			value, err := DecodeTDLibJSONProxyType(b)
			if err != nil {
				return fmt.Errorf("unable to decode addProxy#13c2bcd8: field type: %w", err)
			}
			a.Type = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetServer returns value of Server field.
func (a *AddProxyRequest) GetServer() (value string) {
	if a == nil {
		return
	}
	return a.Server
}

// GetPort returns value of Port field.
func (a *AddProxyRequest) GetPort() (value int32) {
	if a == nil {
		return
	}
	return a.Port
}

// GetEnable returns value of Enable field.
func (a *AddProxyRequest) GetEnable() (value bool) {
	if a == nil {
		return
	}
	return a.Enable
}

// GetType returns value of Type field.
func (a *AddProxyRequest) GetType() (value ProxyTypeClass) {
	if a == nil {
		return
	}
	return a.Type
}

// AddProxy invokes method addProxy#13c2bcd8 returning error if any.
func (c *Client) AddProxy(ctx context.Context, request *AddProxyRequest) (*Proxy, error) {
	var result Proxy

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
