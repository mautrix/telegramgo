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

// StatsPercentValue represents TL type `statsPercentValue#cbce2fe0`.
// Channel statistics percentage¹.
// Compute the percentage simply by doing part * total / 100
//
// Links:
//  1. https://core.telegram.org/api/stats
//
// See https://core.telegram.org/constructor/statsPercentValue for reference.
type StatsPercentValue struct {
	// Partial value
	Part float64
	// Total value
	Total float64
}

// StatsPercentValueTypeID is TL type id of StatsPercentValue.
const StatsPercentValueTypeID = 0xcbce2fe0

// Ensuring interfaces in compile-time for StatsPercentValue.
var (
	_ bin.Encoder     = &StatsPercentValue{}
	_ bin.Decoder     = &StatsPercentValue{}
	_ bin.BareEncoder = &StatsPercentValue{}
	_ bin.BareDecoder = &StatsPercentValue{}
)

func (s *StatsPercentValue) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Part == 0) {
		return false
	}
	if !(s.Total == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsPercentValue) String() string {
	if s == nil {
		return "StatsPercentValue(nil)"
	}
	type Alias StatsPercentValue
	return fmt.Sprintf("StatsPercentValue%+v", Alias(*s))
}

// FillFrom fills StatsPercentValue from given interface.
func (s *StatsPercentValue) FillFrom(from interface {
	GetPart() (value float64)
	GetTotal() (value float64)
}) {
	s.Part = from.GetPart()
	s.Total = from.GetTotal()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsPercentValue) TypeID() uint32 {
	return StatsPercentValueTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsPercentValue) TypeName() string {
	return "statsPercentValue"
}

// TypeInfo returns info about TL type.
func (s *StatsPercentValue) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "statsPercentValue",
		ID:   StatsPercentValueTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Part",
			SchemaName: "part",
		},
		{
			Name:       "Total",
			SchemaName: "total",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StatsPercentValue) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsPercentValue#cbce2fe0 as nil")
	}
	b.PutID(StatsPercentValueTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StatsPercentValue) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsPercentValue#cbce2fe0 as nil")
	}
	b.PutDouble(s.Part)
	b.PutDouble(s.Total)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsPercentValue) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsPercentValue#cbce2fe0 to nil")
	}
	if err := b.ConsumeID(StatsPercentValueTypeID); err != nil {
		return fmt.Errorf("unable to decode statsPercentValue#cbce2fe0: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StatsPercentValue) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsPercentValue#cbce2fe0 to nil")
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode statsPercentValue#cbce2fe0: field part: %w", err)
		}
		s.Part = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode statsPercentValue#cbce2fe0: field total: %w", err)
		}
		s.Total = value
	}
	return nil
}

// GetPart returns value of Part field.
func (s *StatsPercentValue) GetPart() (value float64) {
	if s == nil {
		return
	}
	return s.Part
}

// GetTotal returns value of Total field.
func (s *StatsPercentValue) GetTotal() (value float64) {
	if s == nil {
		return
	}
	return s.Total
}
