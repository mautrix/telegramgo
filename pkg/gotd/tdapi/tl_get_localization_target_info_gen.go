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

// GetLocalizationTargetInfoRequest represents TL type `getLocalizationTargetInfo#6e3d1f86`.
type GetLocalizationTargetInfoRequest struct {
	// Pass true to get only locally available information without sending network requests
	OnlyLocal bool
}

// GetLocalizationTargetInfoRequestTypeID is TL type id of GetLocalizationTargetInfoRequest.
const GetLocalizationTargetInfoRequestTypeID = 0x6e3d1f86

// Ensuring interfaces in compile-time for GetLocalizationTargetInfoRequest.
var (
	_ bin.Encoder     = &GetLocalizationTargetInfoRequest{}
	_ bin.Decoder     = &GetLocalizationTargetInfoRequest{}
	_ bin.BareEncoder = &GetLocalizationTargetInfoRequest{}
	_ bin.BareDecoder = &GetLocalizationTargetInfoRequest{}
)

func (g *GetLocalizationTargetInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.OnlyLocal == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetLocalizationTargetInfoRequest) String() string {
	if g == nil {
		return "GetLocalizationTargetInfoRequest(nil)"
	}
	type Alias GetLocalizationTargetInfoRequest
	return fmt.Sprintf("GetLocalizationTargetInfoRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetLocalizationTargetInfoRequest) TypeID() uint32 {
	return GetLocalizationTargetInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetLocalizationTargetInfoRequest) TypeName() string {
	return "getLocalizationTargetInfo"
}

// TypeInfo returns info about TL type.
func (g *GetLocalizationTargetInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getLocalizationTargetInfo",
		ID:   GetLocalizationTargetInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "OnlyLocal",
			SchemaName: "only_local",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetLocalizationTargetInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLocalizationTargetInfo#6e3d1f86 as nil")
	}
	b.PutID(GetLocalizationTargetInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetLocalizationTargetInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLocalizationTargetInfo#6e3d1f86 as nil")
	}
	b.PutBool(g.OnlyLocal)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetLocalizationTargetInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLocalizationTargetInfo#6e3d1f86 to nil")
	}
	if err := b.ConsumeID(GetLocalizationTargetInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getLocalizationTargetInfo#6e3d1f86: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetLocalizationTargetInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLocalizationTargetInfo#6e3d1f86 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getLocalizationTargetInfo#6e3d1f86: field only_local: %w", err)
		}
		g.OnlyLocal = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetLocalizationTargetInfoRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getLocalizationTargetInfo#6e3d1f86 as nil")
	}
	b.ObjStart()
	b.PutID("getLocalizationTargetInfo")
	b.Comma()
	b.FieldStart("only_local")
	b.PutBool(g.OnlyLocal)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetLocalizationTargetInfoRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getLocalizationTargetInfo#6e3d1f86 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getLocalizationTargetInfo"); err != nil {
				return fmt.Errorf("unable to decode getLocalizationTargetInfo#6e3d1f86: %w", err)
			}
		case "only_local":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getLocalizationTargetInfo#6e3d1f86: field only_local: %w", err)
			}
			g.OnlyLocal = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetOnlyLocal returns value of OnlyLocal field.
func (g *GetLocalizationTargetInfoRequest) GetOnlyLocal() (value bool) {
	if g == nil {
		return
	}
	return g.OnlyLocal
}

// GetLocalizationTargetInfo invokes method getLocalizationTargetInfo#6e3d1f86 returning error if any.
func (c *Client) GetLocalizationTargetInfo(ctx context.Context, onlylocal bool) (*LocalizationTargetInfo, error) {
	var result LocalizationTargetInfo

	request := &GetLocalizationTargetInfoRequest{
		OnlyLocal: onlylocal,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
