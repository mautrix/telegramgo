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

// MessagesStartHistoryImportRequest represents TL type `messages.startHistoryImport#b43df344`.
// Complete the history import process¹, importing all messages into the chat.
// To be called only after initializing the import with messages.initHistoryImport² and
// uploading all files using messages.uploadImportedMedia³.
//
// Links:
//  1. https://core.telegram.org/api/import
//  2. https://core.telegram.org/method/messages.initHistoryImport
//  3. https://core.telegram.org/method/messages.uploadImportedMedia
//
// See https://core.telegram.org/method/messages.startHistoryImport for reference.
type MessagesStartHistoryImportRequest struct {
	// The Telegram chat where the messages should be imported, click here for more info »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/import
	Peer InputPeerClass
	// Identifier of a history import session, returned by messages.initHistoryImport¹.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.initHistoryImport
	ImportID int64
}

// MessagesStartHistoryImportRequestTypeID is TL type id of MessagesStartHistoryImportRequest.
const MessagesStartHistoryImportRequestTypeID = 0xb43df344

// Ensuring interfaces in compile-time for MessagesStartHistoryImportRequest.
var (
	_ bin.Encoder     = &MessagesStartHistoryImportRequest{}
	_ bin.Decoder     = &MessagesStartHistoryImportRequest{}
	_ bin.BareEncoder = &MessagesStartHistoryImportRequest{}
	_ bin.BareDecoder = &MessagesStartHistoryImportRequest{}
)

func (s *MessagesStartHistoryImportRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.ImportID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesStartHistoryImportRequest) String() string {
	if s == nil {
		return "MessagesStartHistoryImportRequest(nil)"
	}
	type Alias MessagesStartHistoryImportRequest
	return fmt.Sprintf("MessagesStartHistoryImportRequest%+v", Alias(*s))
}

// FillFrom fills MessagesStartHistoryImportRequest from given interface.
func (s *MessagesStartHistoryImportRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetImportID() (value int64)
}) {
	s.Peer = from.GetPeer()
	s.ImportID = from.GetImportID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesStartHistoryImportRequest) TypeID() uint32 {
	return MessagesStartHistoryImportRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesStartHistoryImportRequest) TypeName() string {
	return "messages.startHistoryImport"
}

// TypeInfo returns info about TL type.
func (s *MessagesStartHistoryImportRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.startHistoryImport",
		ID:   MessagesStartHistoryImportRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ImportID",
			SchemaName: "import_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesStartHistoryImportRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.startHistoryImport#b43df344 as nil")
	}
	b.PutID(MessagesStartHistoryImportRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesStartHistoryImportRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.startHistoryImport#b43df344 as nil")
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.startHistoryImport#b43df344: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.startHistoryImport#b43df344: field peer: %w", err)
	}
	b.PutLong(s.ImportID)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesStartHistoryImportRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.startHistoryImport#b43df344 to nil")
	}
	if err := b.ConsumeID(MessagesStartHistoryImportRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.startHistoryImport#b43df344: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesStartHistoryImportRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.startHistoryImport#b43df344 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.startHistoryImport#b43df344: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.startHistoryImport#b43df344: field import_id: %w", err)
		}
		s.ImportID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesStartHistoryImportRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetImportID returns value of ImportID field.
func (s *MessagesStartHistoryImportRequest) GetImportID() (value int64) {
	if s == nil {
		return
	}
	return s.ImportID
}

// MessagesStartHistoryImport invokes method messages.startHistoryImport#b43df344 returning error if any.
// Complete the history import process¹, importing all messages into the chat.
// To be called only after initializing the import with messages.initHistoryImport² and
// uploading all files using messages.uploadImportedMedia³.
//
// Links:
//  1. https://core.telegram.org/api/import
//  2. https://core.telegram.org/method/messages.initHistoryImport
//  3. https://core.telegram.org/method/messages.uploadImportedMedia
//
// Possible errors:
//
//	400 IMPORT_ID_INVALID: The specified import ID is invalid.
//
// See https://core.telegram.org/method/messages.startHistoryImport for reference.
func (c *Client) MessagesStartHistoryImport(ctx context.Context, request *MessagesStartHistoryImportRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
