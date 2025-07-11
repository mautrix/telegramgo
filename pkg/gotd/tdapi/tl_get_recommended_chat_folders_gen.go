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

// GetRecommendedChatFoldersRequest represents TL type `getRecommendedChatFolders#f7533b87`.
type GetRecommendedChatFoldersRequest struct {
}

// GetRecommendedChatFoldersRequestTypeID is TL type id of GetRecommendedChatFoldersRequest.
const GetRecommendedChatFoldersRequestTypeID = 0xf7533b87

// Ensuring interfaces in compile-time for GetRecommendedChatFoldersRequest.
var (
	_ bin.Encoder     = &GetRecommendedChatFoldersRequest{}
	_ bin.Decoder     = &GetRecommendedChatFoldersRequest{}
	_ bin.BareEncoder = &GetRecommendedChatFoldersRequest{}
	_ bin.BareDecoder = &GetRecommendedChatFoldersRequest{}
)

func (g *GetRecommendedChatFoldersRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetRecommendedChatFoldersRequest) String() string {
	if g == nil {
		return "GetRecommendedChatFoldersRequest(nil)"
	}
	type Alias GetRecommendedChatFoldersRequest
	return fmt.Sprintf("GetRecommendedChatFoldersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetRecommendedChatFoldersRequest) TypeID() uint32 {
	return GetRecommendedChatFoldersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetRecommendedChatFoldersRequest) TypeName() string {
	return "getRecommendedChatFolders"
}

// TypeInfo returns info about TL type.
func (g *GetRecommendedChatFoldersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getRecommendedChatFolders",
		ID:   GetRecommendedChatFoldersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetRecommendedChatFoldersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getRecommendedChatFolders#f7533b87 as nil")
	}
	b.PutID(GetRecommendedChatFoldersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetRecommendedChatFoldersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getRecommendedChatFolders#f7533b87 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetRecommendedChatFoldersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getRecommendedChatFolders#f7533b87 to nil")
	}
	if err := b.ConsumeID(GetRecommendedChatFoldersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getRecommendedChatFolders#f7533b87: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetRecommendedChatFoldersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getRecommendedChatFolders#f7533b87 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetRecommendedChatFoldersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getRecommendedChatFolders#f7533b87 as nil")
	}
	b.ObjStart()
	b.PutID("getRecommendedChatFolders")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetRecommendedChatFoldersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getRecommendedChatFolders#f7533b87 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getRecommendedChatFolders"); err != nil {
				return fmt.Errorf("unable to decode getRecommendedChatFolders#f7533b87: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetRecommendedChatFolders invokes method getRecommendedChatFolders#f7533b87 returning error if any.
func (c *Client) GetRecommendedChatFolders(ctx context.Context) (*RecommendedChatFolders, error) {
	var result RecommendedChatFolders

	request := &GetRecommendedChatFoldersRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
