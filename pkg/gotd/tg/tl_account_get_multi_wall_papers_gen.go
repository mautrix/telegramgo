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

// AccountGetMultiWallPapersRequest represents TL type `account.getMultiWallPapers#65ad71dc`.
// Get info about multiple wallpapers¹
//
// Links:
//  1. https://core.telegram.org/api/wallpapers
//
// See https://core.telegram.org/method/account.getMultiWallPapers for reference.
type AccountGetMultiWallPapersRequest struct {
	// Wallpapers¹ to fetch info about
	//
	// Links:
	//  1) https://core.telegram.org/api/wallpapers
	Wallpapers []InputWallPaperClass
}

// AccountGetMultiWallPapersRequestTypeID is TL type id of AccountGetMultiWallPapersRequest.
const AccountGetMultiWallPapersRequestTypeID = 0x65ad71dc

// Ensuring interfaces in compile-time for AccountGetMultiWallPapersRequest.
var (
	_ bin.Encoder     = &AccountGetMultiWallPapersRequest{}
	_ bin.Decoder     = &AccountGetMultiWallPapersRequest{}
	_ bin.BareEncoder = &AccountGetMultiWallPapersRequest{}
	_ bin.BareDecoder = &AccountGetMultiWallPapersRequest{}
)

func (g *AccountGetMultiWallPapersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Wallpapers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetMultiWallPapersRequest) String() string {
	if g == nil {
		return "AccountGetMultiWallPapersRequest(nil)"
	}
	type Alias AccountGetMultiWallPapersRequest
	return fmt.Sprintf("AccountGetMultiWallPapersRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetMultiWallPapersRequest from given interface.
func (g *AccountGetMultiWallPapersRequest) FillFrom(from interface {
	GetWallpapers() (value []InputWallPaperClass)
}) {
	g.Wallpapers = from.GetWallpapers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetMultiWallPapersRequest) TypeID() uint32 {
	return AccountGetMultiWallPapersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetMultiWallPapersRequest) TypeName() string {
	return "account.getMultiWallPapers"
}

// TypeInfo returns info about TL type.
func (g *AccountGetMultiWallPapersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getMultiWallPapers",
		ID:   AccountGetMultiWallPapersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Wallpapers",
			SchemaName: "wallpapers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetMultiWallPapersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getMultiWallPapers#65ad71dc as nil")
	}
	b.PutID(AccountGetMultiWallPapersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetMultiWallPapersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getMultiWallPapers#65ad71dc as nil")
	}
	b.PutVectorHeader(len(g.Wallpapers))
	for idx, v := range g.Wallpapers {
		if v == nil {
			return fmt.Errorf("unable to encode account.getMultiWallPapers#65ad71dc: field wallpapers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.getMultiWallPapers#65ad71dc: field wallpapers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetMultiWallPapersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getMultiWallPapers#65ad71dc to nil")
	}
	if err := b.ConsumeID(AccountGetMultiWallPapersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getMultiWallPapers#65ad71dc: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetMultiWallPapersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getMultiWallPapers#65ad71dc to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.getMultiWallPapers#65ad71dc: field wallpapers: %w", err)
		}

		if headerLen > 0 {
			g.Wallpapers = make([]InputWallPaperClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputWallPaper(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.getMultiWallPapers#65ad71dc: field wallpapers: %w", err)
			}
			g.Wallpapers = append(g.Wallpapers, value)
		}
	}
	return nil
}

// GetWallpapers returns value of Wallpapers field.
func (g *AccountGetMultiWallPapersRequest) GetWallpapers() (value []InputWallPaperClass) {
	if g == nil {
		return
	}
	return g.Wallpapers
}

// MapWallpapers returns field Wallpapers wrapped in InputWallPaperClassArray helper.
func (g *AccountGetMultiWallPapersRequest) MapWallpapers() (value InputWallPaperClassArray) {
	return InputWallPaperClassArray(g.Wallpapers)
}

// AccountGetMultiWallPapers invokes method account.getMultiWallPapers#65ad71dc returning error if any.
// Get info about multiple wallpapers¹
//
// Links:
//  1. https://core.telegram.org/api/wallpapers
//
// Possible errors:
//
//	400 WALLPAPER_INVALID: The specified wallpaper is invalid.
//
// See https://core.telegram.org/method/account.getMultiWallPapers for reference.
func (c *Client) AccountGetMultiWallPapers(ctx context.Context, wallpapers []InputWallPaperClass) ([]WallPaperClass, error) {
	var result WallPaperClassVector

	request := &AccountGetMultiWallPapersRequest{
		Wallpapers: wallpapers,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []WallPaperClass(result.Elems), nil
}
