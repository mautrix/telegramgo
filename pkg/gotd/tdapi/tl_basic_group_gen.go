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

// BasicGroup represents TL type `basicGroup#f464168f`.
type BasicGroup struct {
	// Group identifier
	ID int64
	// Number of members in the group
	MemberCount int32
	// Status of the current user in the group
	Status ChatMemberStatusClass
	// True, if the group is active
	IsActive bool
	// Identifier of the supergroup to which this group was upgraded; 0 if none
	UpgradedToSupergroupID int64
}

// BasicGroupTypeID is TL type id of BasicGroup.
const BasicGroupTypeID = 0xf464168f

// Ensuring interfaces in compile-time for BasicGroup.
var (
	_ bin.Encoder     = &BasicGroup{}
	_ bin.Decoder     = &BasicGroup{}
	_ bin.BareEncoder = &BasicGroup{}
	_ bin.BareDecoder = &BasicGroup{}
)

func (b *BasicGroup) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.ID == 0) {
		return false
	}
	if !(b.MemberCount == 0) {
		return false
	}
	if !(b.Status == nil) {
		return false
	}
	if !(b.IsActive == false) {
		return false
	}
	if !(b.UpgradedToSupergroupID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BasicGroup) String() string {
	if b == nil {
		return "BasicGroup(nil)"
	}
	type Alias BasicGroup
	return fmt.Sprintf("BasicGroup%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BasicGroup) TypeID() uint32 {
	return BasicGroupTypeID
}

// TypeName returns name of type in TL schema.
func (*BasicGroup) TypeName() string {
	return "basicGroup"
}

// TypeInfo returns info about TL type.
func (b *BasicGroup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "basicGroup",
		ID:   BasicGroupTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "MemberCount",
			SchemaName: "member_count",
		},
		{
			Name:       "Status",
			SchemaName: "status",
		},
		{
			Name:       "IsActive",
			SchemaName: "is_active",
		},
		{
			Name:       "UpgradedToSupergroupID",
			SchemaName: "upgraded_to_supergroup_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BasicGroup) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode basicGroup#f464168f as nil")
	}
	buf.PutID(BasicGroupTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BasicGroup) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode basicGroup#f464168f as nil")
	}
	buf.PutInt53(b.ID)
	buf.PutInt32(b.MemberCount)
	if b.Status == nil {
		return fmt.Errorf("unable to encode basicGroup#f464168f: field status is nil")
	}
	if err := b.Status.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode basicGroup#f464168f: field status: %w", err)
	}
	buf.PutBool(b.IsActive)
	buf.PutInt53(b.UpgradedToSupergroupID)
	return nil
}

// Decode implements bin.Decoder.
func (b *BasicGroup) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode basicGroup#f464168f to nil")
	}
	if err := buf.ConsumeID(BasicGroupTypeID); err != nil {
		return fmt.Errorf("unable to decode basicGroup#f464168f: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BasicGroup) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode basicGroup#f464168f to nil")
	}
	{
		value, err := buf.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode basicGroup#f464168f: field id: %w", err)
		}
		b.ID = value
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode basicGroup#f464168f: field member_count: %w", err)
		}
		b.MemberCount = value
	}
	{
		value, err := DecodeChatMemberStatus(buf)
		if err != nil {
			return fmt.Errorf("unable to decode basicGroup#f464168f: field status: %w", err)
		}
		b.Status = value
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode basicGroup#f464168f: field is_active: %w", err)
		}
		b.IsActive = value
	}
	{
		value, err := buf.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode basicGroup#f464168f: field upgraded_to_supergroup_id: %w", err)
		}
		b.UpgradedToSupergroupID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BasicGroup) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode basicGroup#f464168f as nil")
	}
	buf.ObjStart()
	buf.PutID("basicGroup")
	buf.Comma()
	buf.FieldStart("id")
	buf.PutInt53(b.ID)
	buf.Comma()
	buf.FieldStart("member_count")
	buf.PutInt32(b.MemberCount)
	buf.Comma()
	buf.FieldStart("status")
	if b.Status == nil {
		return fmt.Errorf("unable to encode basicGroup#f464168f: field status is nil")
	}
	if err := b.Status.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode basicGroup#f464168f: field status: %w", err)
	}
	buf.Comma()
	buf.FieldStart("is_active")
	buf.PutBool(b.IsActive)
	buf.Comma()
	buf.FieldStart("upgraded_to_supergroup_id")
	buf.PutInt53(b.UpgradedToSupergroupID)
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BasicGroup) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode basicGroup#f464168f to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("basicGroup"); err != nil {
				return fmt.Errorf("unable to decode basicGroup#f464168f: %w", err)
			}
		case "id":
			value, err := buf.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode basicGroup#f464168f: field id: %w", err)
			}
			b.ID = value
		case "member_count":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode basicGroup#f464168f: field member_count: %w", err)
			}
			b.MemberCount = value
		case "status":
			value, err := DecodeTDLibJSONChatMemberStatus(buf)
			if err != nil {
				return fmt.Errorf("unable to decode basicGroup#f464168f: field status: %w", err)
			}
			b.Status = value
		case "is_active":
			value, err := buf.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode basicGroup#f464168f: field is_active: %w", err)
			}
			b.IsActive = value
		case "upgraded_to_supergroup_id":
			value, err := buf.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode basicGroup#f464168f: field upgraded_to_supergroup_id: %w", err)
			}
			b.UpgradedToSupergroupID = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (b *BasicGroup) GetID() (value int64) {
	if b == nil {
		return
	}
	return b.ID
}

// GetMemberCount returns value of MemberCount field.
func (b *BasicGroup) GetMemberCount() (value int32) {
	if b == nil {
		return
	}
	return b.MemberCount
}

// GetStatus returns value of Status field.
func (b *BasicGroup) GetStatus() (value ChatMemberStatusClass) {
	if b == nil {
		return
	}
	return b.Status
}

// GetIsActive returns value of IsActive field.
func (b *BasicGroup) GetIsActive() (value bool) {
	if b == nil {
		return
	}
	return b.IsActive
}

// GetUpgradedToSupergroupID returns value of UpgradedToSupergroupID field.
func (b *BasicGroup) GetUpgradedToSupergroupID() (value int64) {
	if b == nil {
		return
	}
	return b.UpgradedToSupergroupID
}
