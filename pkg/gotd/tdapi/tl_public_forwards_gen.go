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

// PublicForwards represents TL type `publicForwards#3ed29447`.
type PublicForwards struct {
	// Approximate total number of messages and stories found
	TotalCount int32
	// List of found public forwards and reposts
	Forwards []PublicForwardClass
	// The offset for the next request. If empty, then there are no more results
	NextOffset string
}

// PublicForwardsTypeID is TL type id of PublicForwards.
const PublicForwardsTypeID = 0x3ed29447

// Ensuring interfaces in compile-time for PublicForwards.
var (
	_ bin.Encoder     = &PublicForwards{}
	_ bin.Decoder     = &PublicForwards{}
	_ bin.BareEncoder = &PublicForwards{}
	_ bin.BareDecoder = &PublicForwards{}
)

func (p *PublicForwards) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.TotalCount == 0) {
		return false
	}
	if !(p.Forwards == nil) {
		return false
	}
	if !(p.NextOffset == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PublicForwards) String() string {
	if p == nil {
		return "PublicForwards(nil)"
	}
	type Alias PublicForwards
	return fmt.Sprintf("PublicForwards%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PublicForwards) TypeID() uint32 {
	return PublicForwardsTypeID
}

// TypeName returns name of type in TL schema.
func (*PublicForwards) TypeName() string {
	return "publicForwards"
}

// TypeInfo returns info about TL type.
func (p *PublicForwards) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "publicForwards",
		ID:   PublicForwardsTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "Forwards",
			SchemaName: "forwards",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PublicForwards) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode publicForwards#3ed29447 as nil")
	}
	b.PutID(PublicForwardsTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PublicForwards) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode publicForwards#3ed29447 as nil")
	}
	b.PutInt32(p.TotalCount)
	b.PutInt(len(p.Forwards))
	for idx, v := range p.Forwards {
		if v == nil {
			return fmt.Errorf("unable to encode publicForwards#3ed29447: field forwards element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare publicForwards#3ed29447: field forwards element with index %d: %w", idx, err)
		}
	}
	b.PutString(p.NextOffset)
	return nil
}

// Decode implements bin.Decoder.
func (p *PublicForwards) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode publicForwards#3ed29447 to nil")
	}
	if err := b.ConsumeID(PublicForwardsTypeID); err != nil {
		return fmt.Errorf("unable to decode publicForwards#3ed29447: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PublicForwards) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode publicForwards#3ed29447 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode publicForwards#3ed29447: field total_count: %w", err)
		}
		p.TotalCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode publicForwards#3ed29447: field forwards: %w", err)
		}

		if headerLen > 0 {
			p.Forwards = make([]PublicForwardClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePublicForward(b)
			if err != nil {
				return fmt.Errorf("unable to decode publicForwards#3ed29447: field forwards: %w", err)
			}
			p.Forwards = append(p.Forwards, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode publicForwards#3ed29447: field next_offset: %w", err)
		}
		p.NextOffset = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PublicForwards) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode publicForwards#3ed29447 as nil")
	}
	b.ObjStart()
	b.PutID("publicForwards")
	b.Comma()
	b.FieldStart("total_count")
	b.PutInt32(p.TotalCount)
	b.Comma()
	b.FieldStart("forwards")
	b.ArrStart()
	for idx, v := range p.Forwards {
		if v == nil {
			return fmt.Errorf("unable to encode publicForwards#3ed29447: field forwards element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode publicForwards#3ed29447: field forwards element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("next_offset")
	b.PutString(p.NextOffset)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PublicForwards) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode publicForwards#3ed29447 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("publicForwards"); err != nil {
				return fmt.Errorf("unable to decode publicForwards#3ed29447: %w", err)
			}
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode publicForwards#3ed29447: field total_count: %w", err)
			}
			p.TotalCount = value
		case "forwards":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := DecodeTDLibJSONPublicForward(b)
				if err != nil {
					return fmt.Errorf("unable to decode publicForwards#3ed29447: field forwards: %w", err)
				}
				p.Forwards = append(p.Forwards, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode publicForwards#3ed29447: field forwards: %w", err)
			}
		case "next_offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode publicForwards#3ed29447: field next_offset: %w", err)
			}
			p.NextOffset = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTotalCount returns value of TotalCount field.
func (p *PublicForwards) GetTotalCount() (value int32) {
	if p == nil {
		return
	}
	return p.TotalCount
}

// GetForwards returns value of Forwards field.
func (p *PublicForwards) GetForwards() (value []PublicForwardClass) {
	if p == nil {
		return
	}
	return p.Forwards
}

// GetNextOffset returns value of NextOffset field.
func (p *PublicForwards) GetNextOffset() (value string) {
	if p == nil {
		return
	}
	return p.NextOffset
}
