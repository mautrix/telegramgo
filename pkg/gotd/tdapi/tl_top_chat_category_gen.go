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

// TopChatCategoryUsers represents TL type `topChatCategoryUsers#3d324d80`.
type TopChatCategoryUsers struct {
}

// TopChatCategoryUsersTypeID is TL type id of TopChatCategoryUsers.
const TopChatCategoryUsersTypeID = 0x3d324d80

// construct implements constructor of TopChatCategoryClass.
func (t TopChatCategoryUsers) construct() TopChatCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopChatCategoryUsers.
var (
	_ bin.Encoder     = &TopChatCategoryUsers{}
	_ bin.Decoder     = &TopChatCategoryUsers{}
	_ bin.BareEncoder = &TopChatCategoryUsers{}
	_ bin.BareDecoder = &TopChatCategoryUsers{}

	_ TopChatCategoryClass = &TopChatCategoryUsers{}
)

func (t *TopChatCategoryUsers) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopChatCategoryUsers) String() string {
	if t == nil {
		return "TopChatCategoryUsers(nil)"
	}
	type Alias TopChatCategoryUsers
	return fmt.Sprintf("TopChatCategoryUsers%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TopChatCategoryUsers) TypeID() uint32 {
	return TopChatCategoryUsersTypeID
}

// TypeName returns name of type in TL schema.
func (*TopChatCategoryUsers) TypeName() string {
	return "topChatCategoryUsers"
}

// TypeInfo returns info about TL type.
func (t *TopChatCategoryUsers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "topChatCategoryUsers",
		ID:   TopChatCategoryUsersTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *TopChatCategoryUsers) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryUsers#3d324d80 as nil")
	}
	b.PutID(TopChatCategoryUsersTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TopChatCategoryUsers) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryUsers#3d324d80 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TopChatCategoryUsers) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryUsers#3d324d80 to nil")
	}
	if err := b.ConsumeID(TopChatCategoryUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode topChatCategoryUsers#3d324d80: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TopChatCategoryUsers) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryUsers#3d324d80 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TopChatCategoryUsers) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryUsers#3d324d80 as nil")
	}
	b.ObjStart()
	b.PutID("topChatCategoryUsers")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TopChatCategoryUsers) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryUsers#3d324d80 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("topChatCategoryUsers"); err != nil {
				return fmt.Errorf("unable to decode topChatCategoryUsers#3d324d80: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// TopChatCategoryBots represents TL type `topChatCategoryBots#a1feeb15`.
type TopChatCategoryBots struct {
}

// TopChatCategoryBotsTypeID is TL type id of TopChatCategoryBots.
const TopChatCategoryBotsTypeID = 0xa1feeb15

// construct implements constructor of TopChatCategoryClass.
func (t TopChatCategoryBots) construct() TopChatCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopChatCategoryBots.
var (
	_ bin.Encoder     = &TopChatCategoryBots{}
	_ bin.Decoder     = &TopChatCategoryBots{}
	_ bin.BareEncoder = &TopChatCategoryBots{}
	_ bin.BareDecoder = &TopChatCategoryBots{}

	_ TopChatCategoryClass = &TopChatCategoryBots{}
)

func (t *TopChatCategoryBots) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopChatCategoryBots) String() string {
	if t == nil {
		return "TopChatCategoryBots(nil)"
	}
	type Alias TopChatCategoryBots
	return fmt.Sprintf("TopChatCategoryBots%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TopChatCategoryBots) TypeID() uint32 {
	return TopChatCategoryBotsTypeID
}

// TypeName returns name of type in TL schema.
func (*TopChatCategoryBots) TypeName() string {
	return "topChatCategoryBots"
}

// TypeInfo returns info about TL type.
func (t *TopChatCategoryBots) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "topChatCategoryBots",
		ID:   TopChatCategoryBotsTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *TopChatCategoryBots) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryBots#a1feeb15 as nil")
	}
	b.PutID(TopChatCategoryBotsTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TopChatCategoryBots) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryBots#a1feeb15 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TopChatCategoryBots) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryBots#a1feeb15 to nil")
	}
	if err := b.ConsumeID(TopChatCategoryBotsTypeID); err != nil {
		return fmt.Errorf("unable to decode topChatCategoryBots#a1feeb15: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TopChatCategoryBots) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryBots#a1feeb15 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TopChatCategoryBots) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryBots#a1feeb15 as nil")
	}
	b.ObjStart()
	b.PutID("topChatCategoryBots")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TopChatCategoryBots) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryBots#a1feeb15 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("topChatCategoryBots"); err != nil {
				return fmt.Errorf("unable to decode topChatCategoryBots#a1feeb15: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// TopChatCategoryGroups represents TL type `topChatCategoryGroups#5b32d08e`.
type TopChatCategoryGroups struct {
}

// TopChatCategoryGroupsTypeID is TL type id of TopChatCategoryGroups.
const TopChatCategoryGroupsTypeID = 0x5b32d08e

// construct implements constructor of TopChatCategoryClass.
func (t TopChatCategoryGroups) construct() TopChatCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopChatCategoryGroups.
var (
	_ bin.Encoder     = &TopChatCategoryGroups{}
	_ bin.Decoder     = &TopChatCategoryGroups{}
	_ bin.BareEncoder = &TopChatCategoryGroups{}
	_ bin.BareDecoder = &TopChatCategoryGroups{}

	_ TopChatCategoryClass = &TopChatCategoryGroups{}
)

func (t *TopChatCategoryGroups) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopChatCategoryGroups) String() string {
	if t == nil {
		return "TopChatCategoryGroups(nil)"
	}
	type Alias TopChatCategoryGroups
	return fmt.Sprintf("TopChatCategoryGroups%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TopChatCategoryGroups) TypeID() uint32 {
	return TopChatCategoryGroupsTypeID
}

// TypeName returns name of type in TL schema.
func (*TopChatCategoryGroups) TypeName() string {
	return "topChatCategoryGroups"
}

// TypeInfo returns info about TL type.
func (t *TopChatCategoryGroups) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "topChatCategoryGroups",
		ID:   TopChatCategoryGroupsTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *TopChatCategoryGroups) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryGroups#5b32d08e as nil")
	}
	b.PutID(TopChatCategoryGroupsTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TopChatCategoryGroups) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryGroups#5b32d08e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TopChatCategoryGroups) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryGroups#5b32d08e to nil")
	}
	if err := b.ConsumeID(TopChatCategoryGroupsTypeID); err != nil {
		return fmt.Errorf("unable to decode topChatCategoryGroups#5b32d08e: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TopChatCategoryGroups) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryGroups#5b32d08e to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TopChatCategoryGroups) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryGroups#5b32d08e as nil")
	}
	b.ObjStart()
	b.PutID("topChatCategoryGroups")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TopChatCategoryGroups) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryGroups#5b32d08e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("topChatCategoryGroups"); err != nil {
				return fmt.Errorf("unable to decode topChatCategoryGroups#5b32d08e: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// TopChatCategoryChannels represents TL type `topChatCategoryChannels#e22600e3`.
type TopChatCategoryChannels struct {
}

// TopChatCategoryChannelsTypeID is TL type id of TopChatCategoryChannels.
const TopChatCategoryChannelsTypeID = 0xe22600e3

// construct implements constructor of TopChatCategoryClass.
func (t TopChatCategoryChannels) construct() TopChatCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopChatCategoryChannels.
var (
	_ bin.Encoder     = &TopChatCategoryChannels{}
	_ bin.Decoder     = &TopChatCategoryChannels{}
	_ bin.BareEncoder = &TopChatCategoryChannels{}
	_ bin.BareDecoder = &TopChatCategoryChannels{}

	_ TopChatCategoryClass = &TopChatCategoryChannels{}
)

func (t *TopChatCategoryChannels) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopChatCategoryChannels) String() string {
	if t == nil {
		return "TopChatCategoryChannels(nil)"
	}
	type Alias TopChatCategoryChannels
	return fmt.Sprintf("TopChatCategoryChannels%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TopChatCategoryChannels) TypeID() uint32 {
	return TopChatCategoryChannelsTypeID
}

// TypeName returns name of type in TL schema.
func (*TopChatCategoryChannels) TypeName() string {
	return "topChatCategoryChannels"
}

// TypeInfo returns info about TL type.
func (t *TopChatCategoryChannels) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "topChatCategoryChannels",
		ID:   TopChatCategoryChannelsTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *TopChatCategoryChannels) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryChannels#e22600e3 as nil")
	}
	b.PutID(TopChatCategoryChannelsTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TopChatCategoryChannels) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryChannels#e22600e3 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TopChatCategoryChannels) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryChannels#e22600e3 to nil")
	}
	if err := b.ConsumeID(TopChatCategoryChannelsTypeID); err != nil {
		return fmt.Errorf("unable to decode topChatCategoryChannels#e22600e3: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TopChatCategoryChannels) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryChannels#e22600e3 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TopChatCategoryChannels) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryChannels#e22600e3 as nil")
	}
	b.ObjStart()
	b.PutID("topChatCategoryChannels")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TopChatCategoryChannels) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryChannels#e22600e3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("topChatCategoryChannels"); err != nil {
				return fmt.Errorf("unable to decode topChatCategoryChannels#e22600e3: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// TopChatCategoryInlineBots represents TL type `topChatCategoryInlineBots#1678eb7c`.
type TopChatCategoryInlineBots struct {
}

// TopChatCategoryInlineBotsTypeID is TL type id of TopChatCategoryInlineBots.
const TopChatCategoryInlineBotsTypeID = 0x1678eb7c

// construct implements constructor of TopChatCategoryClass.
func (t TopChatCategoryInlineBots) construct() TopChatCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopChatCategoryInlineBots.
var (
	_ bin.Encoder     = &TopChatCategoryInlineBots{}
	_ bin.Decoder     = &TopChatCategoryInlineBots{}
	_ bin.BareEncoder = &TopChatCategoryInlineBots{}
	_ bin.BareDecoder = &TopChatCategoryInlineBots{}

	_ TopChatCategoryClass = &TopChatCategoryInlineBots{}
)

func (t *TopChatCategoryInlineBots) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopChatCategoryInlineBots) String() string {
	if t == nil {
		return "TopChatCategoryInlineBots(nil)"
	}
	type Alias TopChatCategoryInlineBots
	return fmt.Sprintf("TopChatCategoryInlineBots%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TopChatCategoryInlineBots) TypeID() uint32 {
	return TopChatCategoryInlineBotsTypeID
}

// TypeName returns name of type in TL schema.
func (*TopChatCategoryInlineBots) TypeName() string {
	return "topChatCategoryInlineBots"
}

// TypeInfo returns info about TL type.
func (t *TopChatCategoryInlineBots) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "topChatCategoryInlineBots",
		ID:   TopChatCategoryInlineBotsTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *TopChatCategoryInlineBots) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryInlineBots#1678eb7c as nil")
	}
	b.PutID(TopChatCategoryInlineBotsTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TopChatCategoryInlineBots) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryInlineBots#1678eb7c as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TopChatCategoryInlineBots) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryInlineBots#1678eb7c to nil")
	}
	if err := b.ConsumeID(TopChatCategoryInlineBotsTypeID); err != nil {
		return fmt.Errorf("unable to decode topChatCategoryInlineBots#1678eb7c: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TopChatCategoryInlineBots) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryInlineBots#1678eb7c to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TopChatCategoryInlineBots) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryInlineBots#1678eb7c as nil")
	}
	b.ObjStart()
	b.PutID("topChatCategoryInlineBots")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TopChatCategoryInlineBots) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryInlineBots#1678eb7c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("topChatCategoryInlineBots"); err != nil {
				return fmt.Errorf("unable to decode topChatCategoryInlineBots#1678eb7c: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// TopChatCategoryWebAppBots represents TL type `topChatCategoryWebAppBots#5f6d6fd`.
type TopChatCategoryWebAppBots struct {
}

// TopChatCategoryWebAppBotsTypeID is TL type id of TopChatCategoryWebAppBots.
const TopChatCategoryWebAppBotsTypeID = 0x5f6d6fd

// construct implements constructor of TopChatCategoryClass.
func (t TopChatCategoryWebAppBots) construct() TopChatCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopChatCategoryWebAppBots.
var (
	_ bin.Encoder     = &TopChatCategoryWebAppBots{}
	_ bin.Decoder     = &TopChatCategoryWebAppBots{}
	_ bin.BareEncoder = &TopChatCategoryWebAppBots{}
	_ bin.BareDecoder = &TopChatCategoryWebAppBots{}

	_ TopChatCategoryClass = &TopChatCategoryWebAppBots{}
)

func (t *TopChatCategoryWebAppBots) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopChatCategoryWebAppBots) String() string {
	if t == nil {
		return "TopChatCategoryWebAppBots(nil)"
	}
	type Alias TopChatCategoryWebAppBots
	return fmt.Sprintf("TopChatCategoryWebAppBots%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TopChatCategoryWebAppBots) TypeID() uint32 {
	return TopChatCategoryWebAppBotsTypeID
}

// TypeName returns name of type in TL schema.
func (*TopChatCategoryWebAppBots) TypeName() string {
	return "topChatCategoryWebAppBots"
}

// TypeInfo returns info about TL type.
func (t *TopChatCategoryWebAppBots) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "topChatCategoryWebAppBots",
		ID:   TopChatCategoryWebAppBotsTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *TopChatCategoryWebAppBots) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryWebAppBots#5f6d6fd as nil")
	}
	b.PutID(TopChatCategoryWebAppBotsTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TopChatCategoryWebAppBots) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryWebAppBots#5f6d6fd as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TopChatCategoryWebAppBots) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryWebAppBots#5f6d6fd to nil")
	}
	if err := b.ConsumeID(TopChatCategoryWebAppBotsTypeID); err != nil {
		return fmt.Errorf("unable to decode topChatCategoryWebAppBots#5f6d6fd: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TopChatCategoryWebAppBots) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryWebAppBots#5f6d6fd to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TopChatCategoryWebAppBots) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryWebAppBots#5f6d6fd as nil")
	}
	b.ObjStart()
	b.PutID("topChatCategoryWebAppBots")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TopChatCategoryWebAppBots) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryWebAppBots#5f6d6fd to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("topChatCategoryWebAppBots"); err != nil {
				return fmt.Errorf("unable to decode topChatCategoryWebAppBots#5f6d6fd: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// TopChatCategoryCalls represents TL type `topChatCategoryCalls#153b50dd`.
type TopChatCategoryCalls struct {
}

// TopChatCategoryCallsTypeID is TL type id of TopChatCategoryCalls.
const TopChatCategoryCallsTypeID = 0x153b50dd

// construct implements constructor of TopChatCategoryClass.
func (t TopChatCategoryCalls) construct() TopChatCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopChatCategoryCalls.
var (
	_ bin.Encoder     = &TopChatCategoryCalls{}
	_ bin.Decoder     = &TopChatCategoryCalls{}
	_ bin.BareEncoder = &TopChatCategoryCalls{}
	_ bin.BareDecoder = &TopChatCategoryCalls{}

	_ TopChatCategoryClass = &TopChatCategoryCalls{}
)

func (t *TopChatCategoryCalls) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopChatCategoryCalls) String() string {
	if t == nil {
		return "TopChatCategoryCalls(nil)"
	}
	type Alias TopChatCategoryCalls
	return fmt.Sprintf("TopChatCategoryCalls%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TopChatCategoryCalls) TypeID() uint32 {
	return TopChatCategoryCallsTypeID
}

// TypeName returns name of type in TL schema.
func (*TopChatCategoryCalls) TypeName() string {
	return "topChatCategoryCalls"
}

// TypeInfo returns info about TL type.
func (t *TopChatCategoryCalls) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "topChatCategoryCalls",
		ID:   TopChatCategoryCallsTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *TopChatCategoryCalls) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryCalls#153b50dd as nil")
	}
	b.PutID(TopChatCategoryCallsTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TopChatCategoryCalls) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryCalls#153b50dd as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TopChatCategoryCalls) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryCalls#153b50dd to nil")
	}
	if err := b.ConsumeID(TopChatCategoryCallsTypeID); err != nil {
		return fmt.Errorf("unable to decode topChatCategoryCalls#153b50dd: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TopChatCategoryCalls) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryCalls#153b50dd to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TopChatCategoryCalls) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryCalls#153b50dd as nil")
	}
	b.ObjStart()
	b.PutID("topChatCategoryCalls")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TopChatCategoryCalls) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryCalls#153b50dd to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("topChatCategoryCalls"); err != nil {
				return fmt.Errorf("unable to decode topChatCategoryCalls#153b50dd: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// TopChatCategoryForwardChats represents TL type `topChatCategoryForwardChats#6515b7d5`.
type TopChatCategoryForwardChats struct {
}

// TopChatCategoryForwardChatsTypeID is TL type id of TopChatCategoryForwardChats.
const TopChatCategoryForwardChatsTypeID = 0x6515b7d5

// construct implements constructor of TopChatCategoryClass.
func (t TopChatCategoryForwardChats) construct() TopChatCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopChatCategoryForwardChats.
var (
	_ bin.Encoder     = &TopChatCategoryForwardChats{}
	_ bin.Decoder     = &TopChatCategoryForwardChats{}
	_ bin.BareEncoder = &TopChatCategoryForwardChats{}
	_ bin.BareDecoder = &TopChatCategoryForwardChats{}

	_ TopChatCategoryClass = &TopChatCategoryForwardChats{}
)

func (t *TopChatCategoryForwardChats) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopChatCategoryForwardChats) String() string {
	if t == nil {
		return "TopChatCategoryForwardChats(nil)"
	}
	type Alias TopChatCategoryForwardChats
	return fmt.Sprintf("TopChatCategoryForwardChats%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TopChatCategoryForwardChats) TypeID() uint32 {
	return TopChatCategoryForwardChatsTypeID
}

// TypeName returns name of type in TL schema.
func (*TopChatCategoryForwardChats) TypeName() string {
	return "topChatCategoryForwardChats"
}

// TypeInfo returns info about TL type.
func (t *TopChatCategoryForwardChats) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "topChatCategoryForwardChats",
		ID:   TopChatCategoryForwardChatsTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *TopChatCategoryForwardChats) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryForwardChats#6515b7d5 as nil")
	}
	b.PutID(TopChatCategoryForwardChatsTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TopChatCategoryForwardChats) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryForwardChats#6515b7d5 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TopChatCategoryForwardChats) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryForwardChats#6515b7d5 to nil")
	}
	if err := b.ConsumeID(TopChatCategoryForwardChatsTypeID); err != nil {
		return fmt.Errorf("unable to decode topChatCategoryForwardChats#6515b7d5: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TopChatCategoryForwardChats) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryForwardChats#6515b7d5 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TopChatCategoryForwardChats) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode topChatCategoryForwardChats#6515b7d5 as nil")
	}
	b.ObjStart()
	b.PutID("topChatCategoryForwardChats")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TopChatCategoryForwardChats) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode topChatCategoryForwardChats#6515b7d5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("topChatCategoryForwardChats"); err != nil {
				return fmt.Errorf("unable to decode topChatCategoryForwardChats#6515b7d5: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// TopChatCategoryClassName is schema name of TopChatCategoryClass.
const TopChatCategoryClassName = "TopChatCategory"

// TopChatCategoryClass represents TopChatCategory generic type.
//
// Example:
//
//	g, err := tdapi.DecodeTopChatCategory(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.TopChatCategoryUsers: // topChatCategoryUsers#3d324d80
//	case *tdapi.TopChatCategoryBots: // topChatCategoryBots#a1feeb15
//	case *tdapi.TopChatCategoryGroups: // topChatCategoryGroups#5b32d08e
//	case *tdapi.TopChatCategoryChannels: // topChatCategoryChannels#e22600e3
//	case *tdapi.TopChatCategoryInlineBots: // topChatCategoryInlineBots#1678eb7c
//	case *tdapi.TopChatCategoryWebAppBots: // topChatCategoryWebAppBots#5f6d6fd
//	case *tdapi.TopChatCategoryCalls: // topChatCategoryCalls#153b50dd
//	case *tdapi.TopChatCategoryForwardChats: // topChatCategoryForwardChats#6515b7d5
//	default: panic(v)
//	}
type TopChatCategoryClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() TopChatCategoryClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeTopChatCategory implements binary de-serialization for TopChatCategoryClass.
func DecodeTopChatCategory(buf *bin.Buffer) (TopChatCategoryClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case TopChatCategoryUsersTypeID:
		// Decoding topChatCategoryUsers#3d324d80.
		v := TopChatCategoryUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	case TopChatCategoryBotsTypeID:
		// Decoding topChatCategoryBots#a1feeb15.
		v := TopChatCategoryBots{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	case TopChatCategoryGroupsTypeID:
		// Decoding topChatCategoryGroups#5b32d08e.
		v := TopChatCategoryGroups{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	case TopChatCategoryChannelsTypeID:
		// Decoding topChatCategoryChannels#e22600e3.
		v := TopChatCategoryChannels{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	case TopChatCategoryInlineBotsTypeID:
		// Decoding topChatCategoryInlineBots#1678eb7c.
		v := TopChatCategoryInlineBots{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	case TopChatCategoryWebAppBotsTypeID:
		// Decoding topChatCategoryWebAppBots#5f6d6fd.
		v := TopChatCategoryWebAppBots{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	case TopChatCategoryCallsTypeID:
		// Decoding topChatCategoryCalls#153b50dd.
		v := TopChatCategoryCalls{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	case TopChatCategoryForwardChatsTypeID:
		// Decoding topChatCategoryForwardChats#6515b7d5.
		v := TopChatCategoryForwardChats{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONTopChatCategory implements binary de-serialization for TopChatCategoryClass.
func DecodeTDLibJSONTopChatCategory(buf tdjson.Decoder) (TopChatCategoryClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "topChatCategoryUsers":
		// Decoding topChatCategoryUsers#3d324d80.
		v := TopChatCategoryUsers{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	case "topChatCategoryBots":
		// Decoding topChatCategoryBots#a1feeb15.
		v := TopChatCategoryBots{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	case "topChatCategoryGroups":
		// Decoding topChatCategoryGroups#5b32d08e.
		v := TopChatCategoryGroups{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	case "topChatCategoryChannels":
		// Decoding topChatCategoryChannels#e22600e3.
		v := TopChatCategoryChannels{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	case "topChatCategoryInlineBots":
		// Decoding topChatCategoryInlineBots#1678eb7c.
		v := TopChatCategoryInlineBots{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	case "topChatCategoryWebAppBots":
		// Decoding topChatCategoryWebAppBots#5f6d6fd.
		v := TopChatCategoryWebAppBots{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	case "topChatCategoryCalls":
		// Decoding topChatCategoryCalls#153b50dd.
		v := TopChatCategoryCalls{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	case "topChatCategoryForwardChats":
		// Decoding topChatCategoryForwardChats#6515b7d5.
		v := TopChatCategoryForwardChats{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode TopChatCategoryClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// TopChatCategory boxes the TopChatCategoryClass providing a helper.
type TopChatCategoryBox struct {
	TopChatCategory TopChatCategoryClass
}

// Decode implements bin.Decoder for TopChatCategoryBox.
func (b *TopChatCategoryBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode TopChatCategoryBox to nil")
	}
	v, err := DecodeTopChatCategory(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.TopChatCategory = v
	return nil
}

// Encode implements bin.Encode for TopChatCategoryBox.
func (b *TopChatCategoryBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.TopChatCategory == nil {
		return fmt.Errorf("unable to encode TopChatCategoryClass as nil")
	}
	return b.TopChatCategory.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for TopChatCategoryBox.
func (b *TopChatCategoryBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode TopChatCategoryBox to nil")
	}
	v, err := DecodeTDLibJSONTopChatCategory(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.TopChatCategory = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for TopChatCategoryBox.
func (b *TopChatCategoryBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.TopChatCategory == nil {
		return fmt.Errorf("unable to encode TopChatCategoryClass as nil")
	}
	return b.TopChatCategory.EncodeTDLibJSON(buf)
}
