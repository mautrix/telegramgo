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

// PhoneSetCallRatingRequest represents TL type `phone.setCallRating#59ead627`.
// Rate a call, returns info about the rating message sent to the official VoIP bot.
//
// See https://core.telegram.org/method/phone.setCallRating for reference.
type PhoneSetCallRatingRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the user decided on their own initiative to rate the call
	UserInitiative bool
	// The call to rate
	Peer InputPhoneCall
	// Rating in 1-5 stars
	Rating int
	// An additional comment
	Comment string
}

// PhoneSetCallRatingRequestTypeID is TL type id of PhoneSetCallRatingRequest.
const PhoneSetCallRatingRequestTypeID = 0x59ead627

// Ensuring interfaces in compile-time for PhoneSetCallRatingRequest.
var (
	_ bin.Encoder     = &PhoneSetCallRatingRequest{}
	_ bin.Decoder     = &PhoneSetCallRatingRequest{}
	_ bin.BareEncoder = &PhoneSetCallRatingRequest{}
	_ bin.BareDecoder = &PhoneSetCallRatingRequest{}
)

func (s *PhoneSetCallRatingRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.UserInitiative == false) {
		return false
	}
	if !(s.Peer.Zero()) {
		return false
	}
	if !(s.Rating == 0) {
		return false
	}
	if !(s.Comment == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PhoneSetCallRatingRequest) String() string {
	if s == nil {
		return "PhoneSetCallRatingRequest(nil)"
	}
	type Alias PhoneSetCallRatingRequest
	return fmt.Sprintf("PhoneSetCallRatingRequest%+v", Alias(*s))
}

// FillFrom fills PhoneSetCallRatingRequest from given interface.
func (s *PhoneSetCallRatingRequest) FillFrom(from interface {
	GetUserInitiative() (value bool)
	GetPeer() (value InputPhoneCall)
	GetRating() (value int)
	GetComment() (value string)
}) {
	s.UserInitiative = from.GetUserInitiative()
	s.Peer = from.GetPeer()
	s.Rating = from.GetRating()
	s.Comment = from.GetComment()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneSetCallRatingRequest) TypeID() uint32 {
	return PhoneSetCallRatingRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneSetCallRatingRequest) TypeName() string {
	return "phone.setCallRating"
}

// TypeInfo returns info about TL type.
func (s *PhoneSetCallRatingRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.setCallRating",
		ID:   PhoneSetCallRatingRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserInitiative",
			SchemaName: "user_initiative",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Rating",
			SchemaName: "rating",
		},
		{
			Name:       "Comment",
			SchemaName: "comment",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *PhoneSetCallRatingRequest) SetFlags() {
	if !(s.UserInitiative == false) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *PhoneSetCallRatingRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode phone.setCallRating#59ead627 as nil")
	}
	b.PutID(PhoneSetCallRatingRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *PhoneSetCallRatingRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode phone.setCallRating#59ead627 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.setCallRating#59ead627: field flags: %w", err)
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.setCallRating#59ead627: field peer: %w", err)
	}
	b.PutInt(s.Rating)
	b.PutString(s.Comment)
	return nil
}

// Decode implements bin.Decoder.
func (s *PhoneSetCallRatingRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode phone.setCallRating#59ead627 to nil")
	}
	if err := b.ConsumeID(PhoneSetCallRatingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.setCallRating#59ead627: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *PhoneSetCallRatingRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode phone.setCallRating#59ead627 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.setCallRating#59ead627: field flags: %w", err)
		}
	}
	s.UserInitiative = s.Flags.Has(0)
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.setCallRating#59ead627: field peer: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.setCallRating#59ead627: field rating: %w", err)
		}
		s.Rating = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phone.setCallRating#59ead627: field comment: %w", err)
		}
		s.Comment = value
	}
	return nil
}

// SetUserInitiative sets value of UserInitiative conditional field.
func (s *PhoneSetCallRatingRequest) SetUserInitiative(value bool) {
	if value {
		s.Flags.Set(0)
		s.UserInitiative = true
	} else {
		s.Flags.Unset(0)
		s.UserInitiative = false
	}
}

// GetUserInitiative returns value of UserInitiative conditional field.
func (s *PhoneSetCallRatingRequest) GetUserInitiative() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (s *PhoneSetCallRatingRequest) GetPeer() (value InputPhoneCall) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetRating returns value of Rating field.
func (s *PhoneSetCallRatingRequest) GetRating() (value int) {
	if s == nil {
		return
	}
	return s.Rating
}

// GetComment returns value of Comment field.
func (s *PhoneSetCallRatingRequest) GetComment() (value string) {
	if s == nil {
		return
	}
	return s.Comment
}

// PhoneSetCallRating invokes method phone.setCallRating#59ead627 returning error if any.
// Rate a call, returns info about the rating message sent to the official VoIP bot.
//
// Possible errors:
//
//	400 CALL_PEER_INVALID: The provided call peer object is invalid.
//
// See https://core.telegram.org/method/phone.setCallRating for reference.
func (c *Client) PhoneSetCallRating(ctx context.Context, request *PhoneSetCallRatingRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
