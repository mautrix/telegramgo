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

// FolderPeer represents TL type `folderPeer#e9baa668`.
// Peer in a folder
//
// See https://core.telegram.org/constructor/folderPeer for reference.
type FolderPeer struct {
	// Folder peer info
	Peer PeerClass
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	FolderID int
}

// FolderPeerTypeID is TL type id of FolderPeer.
const FolderPeerTypeID = 0xe9baa668

// Ensuring interfaces in compile-time for FolderPeer.
var (
	_ bin.Encoder     = &FolderPeer{}
	_ bin.Decoder     = &FolderPeer{}
	_ bin.BareEncoder = &FolderPeer{}
	_ bin.BareDecoder = &FolderPeer{}
)

func (f *FolderPeer) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Peer == nil) {
		return false
	}
	if !(f.FolderID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FolderPeer) String() string {
	if f == nil {
		return "FolderPeer(nil)"
	}
	type Alias FolderPeer
	return fmt.Sprintf("FolderPeer%+v", Alias(*f))
}

// FillFrom fills FolderPeer from given interface.
func (f *FolderPeer) FillFrom(from interface {
	GetPeer() (value PeerClass)
	GetFolderID() (value int)
}) {
	f.Peer = from.GetPeer()
	f.FolderID = from.GetFolderID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FolderPeer) TypeID() uint32 {
	return FolderPeerTypeID
}

// TypeName returns name of type in TL schema.
func (*FolderPeer) TypeName() string {
	return "folderPeer"
}

// TypeInfo returns info about TL type.
func (f *FolderPeer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "folderPeer",
		ID:   FolderPeerTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "FolderID",
			SchemaName: "folder_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *FolderPeer) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode folderPeer#e9baa668 as nil")
	}
	b.PutID(FolderPeerTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FolderPeer) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode folderPeer#e9baa668 as nil")
	}
	if f.Peer == nil {
		return fmt.Errorf("unable to encode folderPeer#e9baa668: field peer is nil")
	}
	if err := f.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode folderPeer#e9baa668: field peer: %w", err)
	}
	b.PutInt(f.FolderID)
	return nil
}

// Decode implements bin.Decoder.
func (f *FolderPeer) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode folderPeer#e9baa668 to nil")
	}
	if err := b.ConsumeID(FolderPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode folderPeer#e9baa668: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FolderPeer) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode folderPeer#e9baa668 to nil")
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode folderPeer#e9baa668: field peer: %w", err)
		}
		f.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode folderPeer#e9baa668: field folder_id: %w", err)
		}
		f.FolderID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (f *FolderPeer) GetPeer() (value PeerClass) {
	if f == nil {
		return
	}
	return f.Peer
}

// GetFolderID returns value of FolderID field.
func (f *FolderPeer) GetFolderID() (value int) {
	if f == nil {
		return
	}
	return f.FolderID
}
