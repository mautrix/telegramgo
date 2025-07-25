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

// GetEmojiCategoriesRequest represents TL type `getEmojiCategories#7f86c16e`.
type GetEmojiCategoriesRequest struct {
	// Type of emoji categories to return; pass null to get default emoji categories
	Type EmojiCategoryTypeClass
}

// GetEmojiCategoriesRequestTypeID is TL type id of GetEmojiCategoriesRequest.
const GetEmojiCategoriesRequestTypeID = 0x7f86c16e

// Ensuring interfaces in compile-time for GetEmojiCategoriesRequest.
var (
	_ bin.Encoder     = &GetEmojiCategoriesRequest{}
	_ bin.Decoder     = &GetEmojiCategoriesRequest{}
	_ bin.BareEncoder = &GetEmojiCategoriesRequest{}
	_ bin.BareDecoder = &GetEmojiCategoriesRequest{}
)

func (g *GetEmojiCategoriesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetEmojiCategoriesRequest) String() string {
	if g == nil {
		return "GetEmojiCategoriesRequest(nil)"
	}
	type Alias GetEmojiCategoriesRequest
	return fmt.Sprintf("GetEmojiCategoriesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetEmojiCategoriesRequest) TypeID() uint32 {
	return GetEmojiCategoriesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetEmojiCategoriesRequest) TypeName() string {
	return "getEmojiCategories"
}

// TypeInfo returns info about TL type.
func (g *GetEmojiCategoriesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getEmojiCategories",
		ID:   GetEmojiCategoriesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetEmojiCategoriesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getEmojiCategories#7f86c16e as nil")
	}
	b.PutID(GetEmojiCategoriesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetEmojiCategoriesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getEmojiCategories#7f86c16e as nil")
	}
	if g.Type == nil {
		return fmt.Errorf("unable to encode getEmojiCategories#7f86c16e: field type is nil")
	}
	if err := g.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getEmojiCategories#7f86c16e: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetEmojiCategoriesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getEmojiCategories#7f86c16e to nil")
	}
	if err := b.ConsumeID(GetEmojiCategoriesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getEmojiCategories#7f86c16e: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetEmojiCategoriesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getEmojiCategories#7f86c16e to nil")
	}
	{
		value, err := DecodeEmojiCategoryType(b)
		if err != nil {
			return fmt.Errorf("unable to decode getEmojiCategories#7f86c16e: field type: %w", err)
		}
		g.Type = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetEmojiCategoriesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getEmojiCategories#7f86c16e as nil")
	}
	b.ObjStart()
	b.PutID("getEmojiCategories")
	b.Comma()
	b.FieldStart("type")
	if g.Type == nil {
		return fmt.Errorf("unable to encode getEmojiCategories#7f86c16e: field type is nil")
	}
	if err := g.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getEmojiCategories#7f86c16e: field type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetEmojiCategoriesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getEmojiCategories#7f86c16e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getEmojiCategories"); err != nil {
				return fmt.Errorf("unable to decode getEmojiCategories#7f86c16e: %w", err)
			}
		case "type":
			value, err := DecodeTDLibJSONEmojiCategoryType(b)
			if err != nil {
				return fmt.Errorf("unable to decode getEmojiCategories#7f86c16e: field type: %w", err)
			}
			g.Type = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetType returns value of Type field.
func (g *GetEmojiCategoriesRequest) GetType() (value EmojiCategoryTypeClass) {
	if g == nil {
		return
	}
	return g.Type
}

// GetEmojiCategories invokes method getEmojiCategories#7f86c16e returning error if any.
func (c *Client) GetEmojiCategories(ctx context.Context, type_ EmojiCategoryTypeClass) (*EmojiCategories, error) {
	var result EmojiCategories

	request := &GetEmojiCategoriesRequest{
		Type: type_,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
