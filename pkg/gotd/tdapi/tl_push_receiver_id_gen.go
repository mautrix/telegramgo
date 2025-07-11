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

// PushReceiverID represents TL type `pushReceiverId#161ddf2c`.
type PushReceiverID struct {
	// The globally unique identifier of push notification subscription
	ID int64
}

// PushReceiverIDTypeID is TL type id of PushReceiverID.
const PushReceiverIDTypeID = 0x161ddf2c

// Ensuring interfaces in compile-time for PushReceiverID.
var (
	_ bin.Encoder     = &PushReceiverID{}
	_ bin.Decoder     = &PushReceiverID{}
	_ bin.BareEncoder = &PushReceiverID{}
	_ bin.BareDecoder = &PushReceiverID{}
)

func (p *PushReceiverID) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PushReceiverID) String() string {
	if p == nil {
		return "PushReceiverID(nil)"
	}
	type Alias PushReceiverID
	return fmt.Sprintf("PushReceiverID%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PushReceiverID) TypeID() uint32 {
	return PushReceiverIDTypeID
}

// TypeName returns name of type in TL schema.
func (*PushReceiverID) TypeName() string {
	return "pushReceiverId"
}

// TypeInfo returns info about TL type.
func (p *PushReceiverID) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pushReceiverId",
		ID:   PushReceiverIDTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PushReceiverID) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pushReceiverId#161ddf2c as nil")
	}
	b.PutID(PushReceiverIDTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PushReceiverID) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pushReceiverId#161ddf2c as nil")
	}
	b.PutLong(p.ID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PushReceiverID) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pushReceiverId#161ddf2c to nil")
	}
	if err := b.ConsumeID(PushReceiverIDTypeID); err != nil {
		return fmt.Errorf("unable to decode pushReceiverId#161ddf2c: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PushReceiverID) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pushReceiverId#161ddf2c to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode pushReceiverId#161ddf2c: field id: %w", err)
		}
		p.ID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PushReceiverID) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode pushReceiverId#161ddf2c as nil")
	}
	b.ObjStart()
	b.PutID("pushReceiverId")
	b.Comma()
	b.FieldStart("id")
	b.PutLong(p.ID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PushReceiverID) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode pushReceiverId#161ddf2c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("pushReceiverId"); err != nil {
				return fmt.Errorf("unable to decode pushReceiverId#161ddf2c: %w", err)
			}
		case "id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode pushReceiverId#161ddf2c: field id: %w", err)
			}
			p.ID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (p *PushReceiverID) GetID() (value int64) {
	if p == nil {
		return
	}
	return p.ID
}
