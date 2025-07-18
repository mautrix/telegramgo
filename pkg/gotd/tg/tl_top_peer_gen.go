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

// TopPeer represents TL type `topPeer#edcdc05b`.
// Top peer
//
// See https://core.telegram.org/constructor/topPeer for reference.
type TopPeer struct {
	// Peer
	Peer PeerClass
	// Rating as computed in top peer rating »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/top-rating
	Rating float64
}

// TopPeerTypeID is TL type id of TopPeer.
const TopPeerTypeID = 0xedcdc05b

// Ensuring interfaces in compile-time for TopPeer.
var (
	_ bin.Encoder     = &TopPeer{}
	_ bin.Decoder     = &TopPeer{}
	_ bin.BareEncoder = &TopPeer{}
	_ bin.BareDecoder = &TopPeer{}
)

func (t *TopPeer) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Peer == nil) {
		return false
	}
	if !(t.Rating == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopPeer) String() string {
	if t == nil {
		return "TopPeer(nil)"
	}
	type Alias TopPeer
	return fmt.Sprintf("TopPeer%+v", Alias(*t))
}

// FillFrom fills TopPeer from given interface.
func (t *TopPeer) FillFrom(from interface {
	GetPeer() (value PeerClass)
	GetRating() (value float64)
}) {
	t.Peer = from.GetPeer()
	t.Rating = from.GetRating()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TopPeer) TypeID() uint32 {
	return TopPeerTypeID
}

// TypeName returns name of type in TL schema.
func (*TopPeer) TypeName() string {
	return "topPeer"
}

// TypeInfo returns info about TL type.
func (t *TopPeer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "topPeer",
		ID:   TopPeerTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Rating",
			SchemaName: "rating",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TopPeer) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeer#edcdc05b as nil")
	}
	b.PutID(TopPeerTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TopPeer) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeer#edcdc05b as nil")
	}
	if t.Peer == nil {
		return fmt.Errorf("unable to encode topPeer#edcdc05b: field peer is nil")
	}
	if err := t.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode topPeer#edcdc05b: field peer: %w", err)
	}
	b.PutDouble(t.Rating)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeer) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeer#edcdc05b to nil")
	}
	if err := b.ConsumeID(TopPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeer#edcdc05b: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TopPeer) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeer#edcdc05b to nil")
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode topPeer#edcdc05b: field peer: %w", err)
		}
		t.Peer = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode topPeer#edcdc05b: field rating: %w", err)
		}
		t.Rating = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (t *TopPeer) GetPeer() (value PeerClass) {
	if t == nil {
		return
	}
	return t.Peer
}

// GetRating returns value of Rating field.
func (t *TopPeer) GetRating() (value float64) {
	if t == nil {
		return
	}
	return t.Rating
}
