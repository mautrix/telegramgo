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

// SearchPublicStoriesByVenueRequest represents TL type `searchPublicStoriesByVenue#d71a622a`.
type SearchPublicStoriesByVenueRequest struct {
	// Provider of the venue
	VenueProvider string
	// Identifier of the venue in the provider database
	VenueID string
	// Offset of the first entry to return as received from the previous request; use empty
	// string to get the first chunk of results
	Offset string
	// The maximum number of stories to be returned; up to 100. For optimal performance, the
	// number of returned stories is chosen by TDLib and can be smaller than the specified
	// limit
	Limit int32
}

// SearchPublicStoriesByVenueRequestTypeID is TL type id of SearchPublicStoriesByVenueRequest.
const SearchPublicStoriesByVenueRequestTypeID = 0xd71a622a

// Ensuring interfaces in compile-time for SearchPublicStoriesByVenueRequest.
var (
	_ bin.Encoder     = &SearchPublicStoriesByVenueRequest{}
	_ bin.Decoder     = &SearchPublicStoriesByVenueRequest{}
	_ bin.BareEncoder = &SearchPublicStoriesByVenueRequest{}
	_ bin.BareDecoder = &SearchPublicStoriesByVenueRequest{}
)

func (s *SearchPublicStoriesByVenueRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.VenueProvider == "") {
		return false
	}
	if !(s.VenueID == "") {
		return false
	}
	if !(s.Offset == "") {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchPublicStoriesByVenueRequest) String() string {
	if s == nil {
		return "SearchPublicStoriesByVenueRequest(nil)"
	}
	type Alias SearchPublicStoriesByVenueRequest
	return fmt.Sprintf("SearchPublicStoriesByVenueRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchPublicStoriesByVenueRequest) TypeID() uint32 {
	return SearchPublicStoriesByVenueRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchPublicStoriesByVenueRequest) TypeName() string {
	return "searchPublicStoriesByVenue"
}

// TypeInfo returns info about TL type.
func (s *SearchPublicStoriesByVenueRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchPublicStoriesByVenue",
		ID:   SearchPublicStoriesByVenueRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "VenueProvider",
			SchemaName: "venue_provider",
		},
		{
			Name:       "VenueID",
			SchemaName: "venue_id",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchPublicStoriesByVenueRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchPublicStoriesByVenue#d71a622a as nil")
	}
	b.PutID(SearchPublicStoriesByVenueRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchPublicStoriesByVenueRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchPublicStoriesByVenue#d71a622a as nil")
	}
	b.PutString(s.VenueProvider)
	b.PutString(s.VenueID)
	b.PutString(s.Offset)
	b.PutInt32(s.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchPublicStoriesByVenueRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchPublicStoriesByVenue#d71a622a to nil")
	}
	if err := b.ConsumeID(SearchPublicStoriesByVenueRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchPublicStoriesByVenue#d71a622a: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchPublicStoriesByVenueRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchPublicStoriesByVenue#d71a622a to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchPublicStoriesByVenue#d71a622a: field venue_provider: %w", err)
		}
		s.VenueProvider = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchPublicStoriesByVenue#d71a622a: field venue_id: %w", err)
		}
		s.VenueID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchPublicStoriesByVenue#d71a622a: field offset: %w", err)
		}
		s.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchPublicStoriesByVenue#d71a622a: field limit: %w", err)
		}
		s.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SearchPublicStoriesByVenueRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchPublicStoriesByVenue#d71a622a as nil")
	}
	b.ObjStart()
	b.PutID("searchPublicStoriesByVenue")
	b.Comma()
	b.FieldStart("venue_provider")
	b.PutString(s.VenueProvider)
	b.Comma()
	b.FieldStart("venue_id")
	b.PutString(s.VenueID)
	b.Comma()
	b.FieldStart("offset")
	b.PutString(s.Offset)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(s.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SearchPublicStoriesByVenueRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchPublicStoriesByVenue#d71a622a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("searchPublicStoriesByVenue"); err != nil {
				return fmt.Errorf("unable to decode searchPublicStoriesByVenue#d71a622a: %w", err)
			}
		case "venue_provider":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchPublicStoriesByVenue#d71a622a: field venue_provider: %w", err)
			}
			s.VenueProvider = value
		case "venue_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchPublicStoriesByVenue#d71a622a: field venue_id: %w", err)
			}
			s.VenueID = value
		case "offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchPublicStoriesByVenue#d71a622a: field offset: %w", err)
			}
			s.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode searchPublicStoriesByVenue#d71a622a: field limit: %w", err)
			}
			s.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetVenueProvider returns value of VenueProvider field.
func (s *SearchPublicStoriesByVenueRequest) GetVenueProvider() (value string) {
	if s == nil {
		return
	}
	return s.VenueProvider
}

// GetVenueID returns value of VenueID field.
func (s *SearchPublicStoriesByVenueRequest) GetVenueID() (value string) {
	if s == nil {
		return
	}
	return s.VenueID
}

// GetOffset returns value of Offset field.
func (s *SearchPublicStoriesByVenueRequest) GetOffset() (value string) {
	if s == nil {
		return
	}
	return s.Offset
}

// GetLimit returns value of Limit field.
func (s *SearchPublicStoriesByVenueRequest) GetLimit() (value int32) {
	if s == nil {
		return
	}
	return s.Limit
}

// SearchPublicStoriesByVenue invokes method searchPublicStoriesByVenue#d71a622a returning error if any.
func (c *Client) SearchPublicStoriesByVenue(ctx context.Context, request *SearchPublicStoriesByVenueRequest) (*FoundStories, error) {
	var result FoundStories

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
