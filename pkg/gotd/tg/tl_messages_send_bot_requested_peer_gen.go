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

// MessagesSendBotRequestedPeerRequest represents TL type `messages.sendBotRequestedPeer#91b2d060`.
// Send one or more chosen peers, as requested by a keyboardButtonRequestPeer¹ button.
//
// Links:
//  1. https://core.telegram.org/constructor/keyboardButtonRequestPeer
//
// See https://core.telegram.org/method/messages.sendBotRequestedPeer for reference.
type MessagesSendBotRequestedPeerRequest struct {
	// The bot that sent the keyboardButtonRequestPeer¹ button.
	//
	// Links:
	//  1) https://core.telegram.org/constructor/keyboardButtonRequestPeer
	Peer InputPeerClass
	// ID of the message that contained the reply keyboard with the
	// keyboardButtonRequestPeer¹ button.
	//
	// Links:
	//  1) https://core.telegram.org/constructor/keyboardButtonRequestPeer
	MsgID int
	// The button_id field from the keyboardButtonRequestPeer¹ constructor.
	//
	// Links:
	//  1) https://core.telegram.org/constructor/keyboardButtonRequestPeer
	ButtonID int
	// The chosen peers.
	RequestedPeers []InputPeerClass
}

// MessagesSendBotRequestedPeerRequestTypeID is TL type id of MessagesSendBotRequestedPeerRequest.
const MessagesSendBotRequestedPeerRequestTypeID = 0x91b2d060

// Ensuring interfaces in compile-time for MessagesSendBotRequestedPeerRequest.
var (
	_ bin.Encoder     = &MessagesSendBotRequestedPeerRequest{}
	_ bin.Decoder     = &MessagesSendBotRequestedPeerRequest{}
	_ bin.BareEncoder = &MessagesSendBotRequestedPeerRequest{}
	_ bin.BareDecoder = &MessagesSendBotRequestedPeerRequest{}
)

func (s *MessagesSendBotRequestedPeerRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.MsgID == 0) {
		return false
	}
	if !(s.ButtonID == 0) {
		return false
	}
	if !(s.RequestedPeers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendBotRequestedPeerRequest) String() string {
	if s == nil {
		return "MessagesSendBotRequestedPeerRequest(nil)"
	}
	type Alias MessagesSendBotRequestedPeerRequest
	return fmt.Sprintf("MessagesSendBotRequestedPeerRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSendBotRequestedPeerRequest from given interface.
func (s *MessagesSendBotRequestedPeerRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetButtonID() (value int)
	GetRequestedPeers() (value []InputPeerClass)
}) {
	s.Peer = from.GetPeer()
	s.MsgID = from.GetMsgID()
	s.ButtonID = from.GetButtonID()
	s.RequestedPeers = from.GetRequestedPeers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSendBotRequestedPeerRequest) TypeID() uint32 {
	return MessagesSendBotRequestedPeerRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSendBotRequestedPeerRequest) TypeName() string {
	return "messages.sendBotRequestedPeer"
}

// TypeInfo returns info about TL type.
func (s *MessagesSendBotRequestedPeerRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sendBotRequestedPeer",
		ID:   MessagesSendBotRequestedPeerRequestTypeID,
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
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "ButtonID",
			SchemaName: "button_id",
		},
		{
			Name:       "RequestedPeers",
			SchemaName: "requested_peers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSendBotRequestedPeerRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendBotRequestedPeer#91b2d060 as nil")
	}
	b.PutID(MessagesSendBotRequestedPeerRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSendBotRequestedPeerRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendBotRequestedPeer#91b2d060 as nil")
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendBotRequestedPeer#91b2d060: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendBotRequestedPeer#91b2d060: field peer: %w", err)
	}
	b.PutInt(s.MsgID)
	b.PutInt(s.ButtonID)
	b.PutVectorHeader(len(s.RequestedPeers))
	for idx, v := range s.RequestedPeers {
		if v == nil {
			return fmt.Errorf("unable to encode messages.sendBotRequestedPeer#91b2d060: field requested_peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendBotRequestedPeer#91b2d060: field requested_peers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendBotRequestedPeerRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendBotRequestedPeer#91b2d060 to nil")
	}
	if err := b.ConsumeID(MessagesSendBotRequestedPeerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendBotRequestedPeer#91b2d060: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSendBotRequestedPeerRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendBotRequestedPeer#91b2d060 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendBotRequestedPeer#91b2d060: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendBotRequestedPeer#91b2d060: field msg_id: %w", err)
		}
		s.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendBotRequestedPeer#91b2d060: field button_id: %w", err)
		}
		s.ButtonID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendBotRequestedPeer#91b2d060: field requested_peers: %w", err)
		}

		if headerLen > 0 {
			s.RequestedPeers = make([]InputPeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.sendBotRequestedPeer#91b2d060: field requested_peers: %w", err)
			}
			s.RequestedPeers = append(s.RequestedPeers, value)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSendBotRequestedPeerRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetMsgID returns value of MsgID field.
func (s *MessagesSendBotRequestedPeerRequest) GetMsgID() (value int) {
	if s == nil {
		return
	}
	return s.MsgID
}

// GetButtonID returns value of ButtonID field.
func (s *MessagesSendBotRequestedPeerRequest) GetButtonID() (value int) {
	if s == nil {
		return
	}
	return s.ButtonID
}

// GetRequestedPeers returns value of RequestedPeers field.
func (s *MessagesSendBotRequestedPeerRequest) GetRequestedPeers() (value []InputPeerClass) {
	if s == nil {
		return
	}
	return s.RequestedPeers
}

// MapRequestedPeers returns field RequestedPeers wrapped in InputPeerClassArray helper.
func (s *MessagesSendBotRequestedPeerRequest) MapRequestedPeers() (value InputPeerClassArray) {
	return InputPeerClassArray(s.RequestedPeers)
}

// MessagesSendBotRequestedPeer invokes method messages.sendBotRequestedPeer#91b2d060 returning error if any.
// Send one or more chosen peers, as requested by a keyboardButtonRequestPeer¹ button.
//
// Links:
//  1. https://core.telegram.org/constructor/keyboardButtonRequestPeer
//
// Possible errors:
//
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.sendBotRequestedPeer for reference.
func (c *Client) MessagesSendBotRequestedPeer(ctx context.Context, request *MessagesSendBotRequestedPeerRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
