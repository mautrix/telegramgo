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

// BusinessInfo represents TL type `businessInfo#5520498e`.
type BusinessInfo struct {
	// Location of the business; may be null if none
	Location BusinessLocation
	// Opening hours of the business; may be null if none. The hours are guaranteed to be
	// valid and has already been split by week days
	OpeningHours BusinessOpeningHours
	// Opening hours of the business in the local time; may be null if none. The hours are
	// guaranteed to be valid and has already been split by week days.
	LocalOpeningHours BusinessOpeningHours
	// Time left before the business will open the next time, in seconds; 0 if unknown. An
	// updateUserFullInfo update is not triggered when value of this field changes
	NextOpenIn int32
	// Time left before the business will close the next time, in seconds; 0 if unknown. An
	// updateUserFullInfo update is not triggered when value of this field changes
	NextCloseIn int32
	// The greeting message; may be null if none or the Business account is not of the
	// current user
	GreetingMessageSettings BusinessGreetingMessageSettings
	// The away message; may be null if none or the Business account is not of the current
	// user
	AwayMessageSettings BusinessAwayMessageSettings
	// Information about start page of the account; may be null if none
	StartPage BusinessStartPage
}

// BusinessInfoTypeID is TL type id of BusinessInfo.
const BusinessInfoTypeID = 0x5520498e

// Ensuring interfaces in compile-time for BusinessInfo.
var (
	_ bin.Encoder     = &BusinessInfo{}
	_ bin.Decoder     = &BusinessInfo{}
	_ bin.BareEncoder = &BusinessInfo{}
	_ bin.BareDecoder = &BusinessInfo{}
)

func (b *BusinessInfo) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Location.Zero()) {
		return false
	}
	if !(b.OpeningHours.Zero()) {
		return false
	}
	if !(b.LocalOpeningHours.Zero()) {
		return false
	}
	if !(b.NextOpenIn == 0) {
		return false
	}
	if !(b.NextCloseIn == 0) {
		return false
	}
	if !(b.GreetingMessageSettings.Zero()) {
		return false
	}
	if !(b.AwayMessageSettings.Zero()) {
		return false
	}
	if !(b.StartPage.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BusinessInfo) String() string {
	if b == nil {
		return "BusinessInfo(nil)"
	}
	type Alias BusinessInfo
	return fmt.Sprintf("BusinessInfo%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BusinessInfo) TypeID() uint32 {
	return BusinessInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*BusinessInfo) TypeName() string {
	return "businessInfo"
}

// TypeInfo returns info about TL type.
func (b *BusinessInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "businessInfo",
		ID:   BusinessInfoTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Location",
			SchemaName: "location",
		},
		{
			Name:       "OpeningHours",
			SchemaName: "opening_hours",
		},
		{
			Name:       "LocalOpeningHours",
			SchemaName: "local_opening_hours",
		},
		{
			Name:       "NextOpenIn",
			SchemaName: "next_open_in",
		},
		{
			Name:       "NextCloseIn",
			SchemaName: "next_close_in",
		},
		{
			Name:       "GreetingMessageSettings",
			SchemaName: "greeting_message_settings",
		},
		{
			Name:       "AwayMessageSettings",
			SchemaName: "away_message_settings",
		},
		{
			Name:       "StartPage",
			SchemaName: "start_page",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BusinessInfo) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessInfo#5520498e as nil")
	}
	buf.PutID(BusinessInfoTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BusinessInfo) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessInfo#5520498e as nil")
	}
	if err := b.Location.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode businessInfo#5520498e: field location: %w", err)
	}
	if err := b.OpeningHours.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode businessInfo#5520498e: field opening_hours: %w", err)
	}
	if err := b.LocalOpeningHours.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode businessInfo#5520498e: field local_opening_hours: %w", err)
	}
	buf.PutInt32(b.NextOpenIn)
	buf.PutInt32(b.NextCloseIn)
	if err := b.GreetingMessageSettings.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode businessInfo#5520498e: field greeting_message_settings: %w", err)
	}
	if err := b.AwayMessageSettings.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode businessInfo#5520498e: field away_message_settings: %w", err)
	}
	if err := b.StartPage.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode businessInfo#5520498e: field start_page: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BusinessInfo) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessInfo#5520498e to nil")
	}
	if err := buf.ConsumeID(BusinessInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode businessInfo#5520498e: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BusinessInfo) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessInfo#5520498e to nil")
	}
	{
		if err := b.Location.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode businessInfo#5520498e: field location: %w", err)
		}
	}
	{
		if err := b.OpeningHours.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode businessInfo#5520498e: field opening_hours: %w", err)
		}
	}
	{
		if err := b.LocalOpeningHours.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode businessInfo#5520498e: field local_opening_hours: %w", err)
		}
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode businessInfo#5520498e: field next_open_in: %w", err)
		}
		b.NextOpenIn = value
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode businessInfo#5520498e: field next_close_in: %w", err)
		}
		b.NextCloseIn = value
	}
	{
		if err := b.GreetingMessageSettings.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode businessInfo#5520498e: field greeting_message_settings: %w", err)
		}
	}
	{
		if err := b.AwayMessageSettings.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode businessInfo#5520498e: field away_message_settings: %w", err)
		}
	}
	{
		if err := b.StartPage.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode businessInfo#5520498e: field start_page: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BusinessInfo) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode businessInfo#5520498e as nil")
	}
	buf.ObjStart()
	buf.PutID("businessInfo")
	buf.Comma()
	buf.FieldStart("location")
	if err := b.Location.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode businessInfo#5520498e: field location: %w", err)
	}
	buf.Comma()
	buf.FieldStart("opening_hours")
	if err := b.OpeningHours.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode businessInfo#5520498e: field opening_hours: %w", err)
	}
	buf.Comma()
	buf.FieldStart("local_opening_hours")
	if err := b.LocalOpeningHours.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode businessInfo#5520498e: field local_opening_hours: %w", err)
	}
	buf.Comma()
	buf.FieldStart("next_open_in")
	buf.PutInt32(b.NextOpenIn)
	buf.Comma()
	buf.FieldStart("next_close_in")
	buf.PutInt32(b.NextCloseIn)
	buf.Comma()
	buf.FieldStart("greeting_message_settings")
	if err := b.GreetingMessageSettings.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode businessInfo#5520498e: field greeting_message_settings: %w", err)
	}
	buf.Comma()
	buf.FieldStart("away_message_settings")
	if err := b.AwayMessageSettings.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode businessInfo#5520498e: field away_message_settings: %w", err)
	}
	buf.Comma()
	buf.FieldStart("start_page")
	if err := b.StartPage.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode businessInfo#5520498e: field start_page: %w", err)
	}
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BusinessInfo) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode businessInfo#5520498e to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("businessInfo"); err != nil {
				return fmt.Errorf("unable to decode businessInfo#5520498e: %w", err)
			}
		case "location":
			if err := b.Location.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode businessInfo#5520498e: field location: %w", err)
			}
		case "opening_hours":
			if err := b.OpeningHours.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode businessInfo#5520498e: field opening_hours: %w", err)
			}
		case "local_opening_hours":
			if err := b.LocalOpeningHours.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode businessInfo#5520498e: field local_opening_hours: %w", err)
			}
		case "next_open_in":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode businessInfo#5520498e: field next_open_in: %w", err)
			}
			b.NextOpenIn = value
		case "next_close_in":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode businessInfo#5520498e: field next_close_in: %w", err)
			}
			b.NextCloseIn = value
		case "greeting_message_settings":
			if err := b.GreetingMessageSettings.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode businessInfo#5520498e: field greeting_message_settings: %w", err)
			}
		case "away_message_settings":
			if err := b.AwayMessageSettings.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode businessInfo#5520498e: field away_message_settings: %w", err)
			}
		case "start_page":
			if err := b.StartPage.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode businessInfo#5520498e: field start_page: %w", err)
			}
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetLocation returns value of Location field.
func (b *BusinessInfo) GetLocation() (value BusinessLocation) {
	if b == nil {
		return
	}
	return b.Location
}

// GetOpeningHours returns value of OpeningHours field.
func (b *BusinessInfo) GetOpeningHours() (value BusinessOpeningHours) {
	if b == nil {
		return
	}
	return b.OpeningHours
}

// GetLocalOpeningHours returns value of LocalOpeningHours field.
func (b *BusinessInfo) GetLocalOpeningHours() (value BusinessOpeningHours) {
	if b == nil {
		return
	}
	return b.LocalOpeningHours
}

// GetNextOpenIn returns value of NextOpenIn field.
func (b *BusinessInfo) GetNextOpenIn() (value int32) {
	if b == nil {
		return
	}
	return b.NextOpenIn
}

// GetNextCloseIn returns value of NextCloseIn field.
func (b *BusinessInfo) GetNextCloseIn() (value int32) {
	if b == nil {
		return
	}
	return b.NextCloseIn
}

// GetGreetingMessageSettings returns value of GreetingMessageSettings field.
func (b *BusinessInfo) GetGreetingMessageSettings() (value BusinessGreetingMessageSettings) {
	if b == nil {
		return
	}
	return b.GreetingMessageSettings
}

// GetAwayMessageSettings returns value of AwayMessageSettings field.
func (b *BusinessInfo) GetAwayMessageSettings() (value BusinessAwayMessageSettings) {
	if b == nil {
		return
	}
	return b.AwayMessageSettings
}

// GetStartPage returns value of StartPage field.
func (b *BusinessInfo) GetStartPage() (value BusinessStartPage) {
	if b == nil {
		return
	}
	return b.StartPage
}
