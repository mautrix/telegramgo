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

// GetProxyLinkRequest represents TL type `getProxyLink#c125ae78`.
type GetProxyLinkRequest struct {
	// Proxy identifier
	ProxyID int32
}

// GetProxyLinkRequestTypeID is TL type id of GetProxyLinkRequest.
const GetProxyLinkRequestTypeID = 0xc125ae78

// Ensuring interfaces in compile-time for GetProxyLinkRequest.
var (
	_ bin.Encoder     = &GetProxyLinkRequest{}
	_ bin.Decoder     = &GetProxyLinkRequest{}
	_ bin.BareEncoder = &GetProxyLinkRequest{}
	_ bin.BareDecoder = &GetProxyLinkRequest{}
)

func (g *GetProxyLinkRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ProxyID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetProxyLinkRequest) String() string {
	if g == nil {
		return "GetProxyLinkRequest(nil)"
	}
	type Alias GetProxyLinkRequest
	return fmt.Sprintf("GetProxyLinkRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetProxyLinkRequest) TypeID() uint32 {
	return GetProxyLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetProxyLinkRequest) TypeName() string {
	return "getProxyLink"
}

// TypeInfo returns info about TL type.
func (g *GetProxyLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getProxyLink",
		ID:   GetProxyLinkRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ProxyID",
			SchemaName: "proxy_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetProxyLinkRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getProxyLink#c125ae78 as nil")
	}
	b.PutID(GetProxyLinkRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetProxyLinkRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getProxyLink#c125ae78 as nil")
	}
	b.PutInt32(g.ProxyID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetProxyLinkRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getProxyLink#c125ae78 to nil")
	}
	if err := b.ConsumeID(GetProxyLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getProxyLink#c125ae78: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetProxyLinkRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getProxyLink#c125ae78 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getProxyLink#c125ae78: field proxy_id: %w", err)
		}
		g.ProxyID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetProxyLinkRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getProxyLink#c125ae78 as nil")
	}
	b.ObjStart()
	b.PutID("getProxyLink")
	b.Comma()
	b.FieldStart("proxy_id")
	b.PutInt32(g.ProxyID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetProxyLinkRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getProxyLink#c125ae78 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getProxyLink"); err != nil {
				return fmt.Errorf("unable to decode getProxyLink#c125ae78: %w", err)
			}
		case "proxy_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getProxyLink#c125ae78: field proxy_id: %w", err)
			}
			g.ProxyID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetProxyID returns value of ProxyID field.
func (g *GetProxyLinkRequest) GetProxyID() (value int32) {
	if g == nil {
		return
	}
	return g.ProxyID
}

// GetProxyLink invokes method getProxyLink#c125ae78 returning error if any.
func (c *Client) GetProxyLink(ctx context.Context, proxyid int32) (*HTTPURL, error) {
	var result HTTPURL

	request := &GetProxyLinkRequest{
		ProxyID: proxyid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
