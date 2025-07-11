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

// AvailableGift represents TL type `availableGift#32b210c6`.
type AvailableGift struct {
	// The gift
	Gift Gift
	// Number of gifts that are available for resale
	ResaleCount int32
	// The minimum price for the gifts available for resale; 0 if there are no such gifts
	MinResaleStarCount int64
	// The title of the upgraded gift; empty if the gift isn't available for resale
	Title string
}

// AvailableGiftTypeID is TL type id of AvailableGift.
const AvailableGiftTypeID = 0x32b210c6

// Ensuring interfaces in compile-time for AvailableGift.
var (
	_ bin.Encoder     = &AvailableGift{}
	_ bin.Decoder     = &AvailableGift{}
	_ bin.BareEncoder = &AvailableGift{}
	_ bin.BareDecoder = &AvailableGift{}
)

func (a *AvailableGift) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Gift.Zero()) {
		return false
	}
	if !(a.ResaleCount == 0) {
		return false
	}
	if !(a.MinResaleStarCount == 0) {
		return false
	}
	if !(a.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AvailableGift) String() string {
	if a == nil {
		return "AvailableGift(nil)"
	}
	type Alias AvailableGift
	return fmt.Sprintf("AvailableGift%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AvailableGift) TypeID() uint32 {
	return AvailableGiftTypeID
}

// TypeName returns name of type in TL schema.
func (*AvailableGift) TypeName() string {
	return "availableGift"
}

// TypeInfo returns info about TL type.
func (a *AvailableGift) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "availableGift",
		ID:   AvailableGiftTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Gift",
			SchemaName: "gift",
		},
		{
			Name:       "ResaleCount",
			SchemaName: "resale_count",
		},
		{
			Name:       "MinResaleStarCount",
			SchemaName: "min_resale_star_count",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AvailableGift) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode availableGift#32b210c6 as nil")
	}
	b.PutID(AvailableGiftTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AvailableGift) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode availableGift#32b210c6 as nil")
	}
	if err := a.Gift.Encode(b); err != nil {
		return fmt.Errorf("unable to encode availableGift#32b210c6: field gift: %w", err)
	}
	b.PutInt32(a.ResaleCount)
	b.PutInt53(a.MinResaleStarCount)
	b.PutString(a.Title)
	return nil
}

// Decode implements bin.Decoder.
func (a *AvailableGift) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode availableGift#32b210c6 to nil")
	}
	if err := b.ConsumeID(AvailableGiftTypeID); err != nil {
		return fmt.Errorf("unable to decode availableGift#32b210c6: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AvailableGift) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode availableGift#32b210c6 to nil")
	}
	{
		if err := a.Gift.Decode(b); err != nil {
			return fmt.Errorf("unable to decode availableGift#32b210c6: field gift: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode availableGift#32b210c6: field resale_count: %w", err)
		}
		a.ResaleCount = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode availableGift#32b210c6: field min_resale_star_count: %w", err)
		}
		a.MinResaleStarCount = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode availableGift#32b210c6: field title: %w", err)
		}
		a.Title = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AvailableGift) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode availableGift#32b210c6 as nil")
	}
	b.ObjStart()
	b.PutID("availableGift")
	b.Comma()
	b.FieldStart("gift")
	if err := a.Gift.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode availableGift#32b210c6: field gift: %w", err)
	}
	b.Comma()
	b.FieldStart("resale_count")
	b.PutInt32(a.ResaleCount)
	b.Comma()
	b.FieldStart("min_resale_star_count")
	b.PutInt53(a.MinResaleStarCount)
	b.Comma()
	b.FieldStart("title")
	b.PutString(a.Title)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AvailableGift) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode availableGift#32b210c6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("availableGift"); err != nil {
				return fmt.Errorf("unable to decode availableGift#32b210c6: %w", err)
			}
		case "gift":
			if err := a.Gift.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode availableGift#32b210c6: field gift: %w", err)
			}
		case "resale_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode availableGift#32b210c6: field resale_count: %w", err)
			}
			a.ResaleCount = value
		case "min_resale_star_count":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode availableGift#32b210c6: field min_resale_star_count: %w", err)
			}
			a.MinResaleStarCount = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode availableGift#32b210c6: field title: %w", err)
			}
			a.Title = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGift returns value of Gift field.
func (a *AvailableGift) GetGift() (value Gift) {
	if a == nil {
		return
	}
	return a.Gift
}

// GetResaleCount returns value of ResaleCount field.
func (a *AvailableGift) GetResaleCount() (value int32) {
	if a == nil {
		return
	}
	return a.ResaleCount
}

// GetMinResaleStarCount returns value of MinResaleStarCount field.
func (a *AvailableGift) GetMinResaleStarCount() (value int64) {
	if a == nil {
		return
	}
	return a.MinResaleStarCount
}

// GetTitle returns value of Title field.
func (a *AvailableGift) GetTitle() (value string) {
	if a == nil {
		return
	}
	return a.Title
}
