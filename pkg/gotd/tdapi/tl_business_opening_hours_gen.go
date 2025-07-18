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

// BusinessOpeningHours represents TL type `businessOpeningHours#a623e64a`.
type BusinessOpeningHours struct {
	// Unique time zone identifier
	TimeZoneID string
	// Intervals of the time when the business is open
	OpeningHours []BusinessOpeningHoursInterval
}

// BusinessOpeningHoursTypeID is TL type id of BusinessOpeningHours.
const BusinessOpeningHoursTypeID = 0xa623e64a

// Ensuring interfaces in compile-time for BusinessOpeningHours.
var (
	_ bin.Encoder     = &BusinessOpeningHours{}
	_ bin.Decoder     = &BusinessOpeningHours{}
	_ bin.BareEncoder = &BusinessOpeningHours{}
	_ bin.BareDecoder = &BusinessOpeningHours{}
)

func (b *BusinessOpeningHours) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.TimeZoneID == "") {
		return false
	}
	if !(b.OpeningHours == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BusinessOpeningHours) String() string {
	if b == nil {
		return "BusinessOpeningHours(nil)"
	}
	type Alias BusinessOpeningHours
	return fmt.Sprintf("BusinessOpeningHours%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BusinessOpeningHours) TypeID() uint32 {
	return BusinessOpeningHoursTypeID
}

// TypeName returns name of type in TL schema.
func (*BusinessOpeningHours) TypeName() string {
	return "businessOpeningHours"
}

// TypeInfo returns info about TL type.
func (b *BusinessOpeningHours) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "businessOpeningHours",
		ID:   BusinessOpeningHoursTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TimeZoneID",
			SchemaName: "time_zone_id",
		},
		{
			Name:       "OpeningHours",
			SchemaName: "opening_hours",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BusinessOpeningHours) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessOpeningHours#a623e64a as nil")
	}
	buf.PutID(BusinessOpeningHoursTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BusinessOpeningHours) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessOpeningHours#a623e64a as nil")
	}
	buf.PutString(b.TimeZoneID)
	buf.PutInt(len(b.OpeningHours))
	for idx, v := range b.OpeningHours {
		if err := v.EncodeBare(buf); err != nil {
			return fmt.Errorf("unable to encode bare businessOpeningHours#a623e64a: field opening_hours element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BusinessOpeningHours) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessOpeningHours#a623e64a to nil")
	}
	if err := buf.ConsumeID(BusinessOpeningHoursTypeID); err != nil {
		return fmt.Errorf("unable to decode businessOpeningHours#a623e64a: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BusinessOpeningHours) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessOpeningHours#a623e64a to nil")
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode businessOpeningHours#a623e64a: field time_zone_id: %w", err)
		}
		b.TimeZoneID = value
	}
	{
		headerLen, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode businessOpeningHours#a623e64a: field opening_hours: %w", err)
		}

		if headerLen > 0 {
			b.OpeningHours = make([]BusinessOpeningHoursInterval, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BusinessOpeningHoursInterval
			if err := value.DecodeBare(buf); err != nil {
				return fmt.Errorf("unable to decode bare businessOpeningHours#a623e64a: field opening_hours: %w", err)
			}
			b.OpeningHours = append(b.OpeningHours, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BusinessOpeningHours) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode businessOpeningHours#a623e64a as nil")
	}
	buf.ObjStart()
	buf.PutID("businessOpeningHours")
	buf.Comma()
	buf.FieldStart("time_zone_id")
	buf.PutString(b.TimeZoneID)
	buf.Comma()
	buf.FieldStart("opening_hours")
	buf.ArrStart()
	for idx, v := range b.OpeningHours {
		if err := v.EncodeTDLibJSON(buf); err != nil {
			return fmt.Errorf("unable to encode businessOpeningHours#a623e64a: field opening_hours element with index %d: %w", idx, err)
		}
		buf.Comma()
	}
	buf.StripComma()
	buf.ArrEnd()
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BusinessOpeningHours) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode businessOpeningHours#a623e64a to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("businessOpeningHours"); err != nil {
				return fmt.Errorf("unable to decode businessOpeningHours#a623e64a: %w", err)
			}
		case "time_zone_id":
			value, err := buf.String()
			if err != nil {
				return fmt.Errorf("unable to decode businessOpeningHours#a623e64a: field time_zone_id: %w", err)
			}
			b.TimeZoneID = value
		case "opening_hours":
			if err := buf.Arr(func(buf tdjson.Decoder) error {
				var value BusinessOpeningHoursInterval
				if err := value.DecodeTDLibJSON(buf); err != nil {
					return fmt.Errorf("unable to decode businessOpeningHours#a623e64a: field opening_hours: %w", err)
				}
				b.OpeningHours = append(b.OpeningHours, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode businessOpeningHours#a623e64a: field opening_hours: %w", err)
			}
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetTimeZoneID returns value of TimeZoneID field.
func (b *BusinessOpeningHours) GetTimeZoneID() (value string) {
	if b == nil {
		return
	}
	return b.TimeZoneID
}

// GetOpeningHours returns value of OpeningHours field.
func (b *BusinessOpeningHours) GetOpeningHours() (value []BusinessOpeningHoursInterval) {
	if b == nil {
		return
	}
	return b.OpeningHours
}
