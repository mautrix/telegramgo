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

// UpgradedGiftBackdropCount represents TL type `upgradedGiftBackdropCount#de7c60b8`.
type UpgradedGiftBackdropCount struct {
	// The backdrop
	Backdrop UpgradedGiftBackdrop
	// Total number of gifts with the symbol
	TotalCount int32
}

// UpgradedGiftBackdropCountTypeID is TL type id of UpgradedGiftBackdropCount.
const UpgradedGiftBackdropCountTypeID = 0xde7c60b8

// Ensuring interfaces in compile-time for UpgradedGiftBackdropCount.
var (
	_ bin.Encoder     = &UpgradedGiftBackdropCount{}
	_ bin.Decoder     = &UpgradedGiftBackdropCount{}
	_ bin.BareEncoder = &UpgradedGiftBackdropCount{}
	_ bin.BareDecoder = &UpgradedGiftBackdropCount{}
)

func (u *UpgradedGiftBackdropCount) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Backdrop.Zero()) {
		return false
	}
	if !(u.TotalCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UpgradedGiftBackdropCount) String() string {
	if u == nil {
		return "UpgradedGiftBackdropCount(nil)"
	}
	type Alias UpgradedGiftBackdropCount
	return fmt.Sprintf("UpgradedGiftBackdropCount%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UpgradedGiftBackdropCount) TypeID() uint32 {
	return UpgradedGiftBackdropCountTypeID
}

// TypeName returns name of type in TL schema.
func (*UpgradedGiftBackdropCount) TypeName() string {
	return "upgradedGiftBackdropCount"
}

// TypeInfo returns info about TL type.
func (u *UpgradedGiftBackdropCount) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upgradedGiftBackdropCount",
		ID:   UpgradedGiftBackdropCountTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Backdrop",
			SchemaName: "backdrop",
		},
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UpgradedGiftBackdropCount) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode upgradedGiftBackdropCount#de7c60b8 as nil")
	}
	b.PutID(UpgradedGiftBackdropCountTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UpgradedGiftBackdropCount) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode upgradedGiftBackdropCount#de7c60b8 as nil")
	}
	if err := u.Backdrop.Encode(b); err != nil {
		return fmt.Errorf("unable to encode upgradedGiftBackdropCount#de7c60b8: field backdrop: %w", err)
	}
	b.PutInt32(u.TotalCount)
	return nil
}

// Decode implements bin.Decoder.
func (u *UpgradedGiftBackdropCount) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode upgradedGiftBackdropCount#de7c60b8 to nil")
	}
	if err := b.ConsumeID(UpgradedGiftBackdropCountTypeID); err != nil {
		return fmt.Errorf("unable to decode upgradedGiftBackdropCount#de7c60b8: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UpgradedGiftBackdropCount) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode upgradedGiftBackdropCount#de7c60b8 to nil")
	}
	{
		if err := u.Backdrop.Decode(b); err != nil {
			return fmt.Errorf("unable to decode upgradedGiftBackdropCount#de7c60b8: field backdrop: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode upgradedGiftBackdropCount#de7c60b8: field total_count: %w", err)
		}
		u.TotalCount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *UpgradedGiftBackdropCount) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode upgradedGiftBackdropCount#de7c60b8 as nil")
	}
	b.ObjStart()
	b.PutID("upgradedGiftBackdropCount")
	b.Comma()
	b.FieldStart("backdrop")
	if err := u.Backdrop.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode upgradedGiftBackdropCount#de7c60b8: field backdrop: %w", err)
	}
	b.Comma()
	b.FieldStart("total_count")
	b.PutInt32(u.TotalCount)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *UpgradedGiftBackdropCount) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode upgradedGiftBackdropCount#de7c60b8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("upgradedGiftBackdropCount"); err != nil {
				return fmt.Errorf("unable to decode upgradedGiftBackdropCount#de7c60b8: %w", err)
			}
		case "backdrop":
			if err := u.Backdrop.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode upgradedGiftBackdropCount#de7c60b8: field backdrop: %w", err)
			}
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode upgradedGiftBackdropCount#de7c60b8: field total_count: %w", err)
			}
			u.TotalCount = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBackdrop returns value of Backdrop field.
func (u *UpgradedGiftBackdropCount) GetBackdrop() (value UpgradedGiftBackdrop) {
	if u == nil {
		return
	}
	return u.Backdrop
}

// GetTotalCount returns value of TotalCount field.
func (u *UpgradedGiftBackdropCount) GetTotalCount() (value int32) {
	if u == nil {
		return
	}
	return u.TotalCount
}
