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

// MessagesGetBotCallbackAnswerRequest represents TL type `messages.getBotCallbackAnswer#9342ca07`.
// Press an inline callback button and get a callback answer from the bot
//
// See https://core.telegram.org/method/messages.getBotCallbackAnswer for reference.
type MessagesGetBotCallbackAnswerRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is a "play game" button
	Game bool
	// Where was the inline keyboard sent
	Peer InputPeerClass
	// ID of the Message with the inline keyboard
	MsgID int
	// Callback data
	//
	// Use SetData and GetData helpers.
	Data []byte
	// For buttons requiring you to verify your identity with your 2FA password¹, the SRP
	// payload generated using SRP².
	//
	// Links:
	//  1) https://core.telegram.org/constructor/keyboardButtonCallback
	//  2) https://core.telegram.org/api/srp
	//
	// Use SetPassword and GetPassword helpers.
	Password InputCheckPasswordSRPClass
}

// MessagesGetBotCallbackAnswerRequestTypeID is TL type id of MessagesGetBotCallbackAnswerRequest.
const MessagesGetBotCallbackAnswerRequestTypeID = 0x9342ca07

// Ensuring interfaces in compile-time for MessagesGetBotCallbackAnswerRequest.
var (
	_ bin.Encoder     = &MessagesGetBotCallbackAnswerRequest{}
	_ bin.Decoder     = &MessagesGetBotCallbackAnswerRequest{}
	_ bin.BareEncoder = &MessagesGetBotCallbackAnswerRequest{}
	_ bin.BareDecoder = &MessagesGetBotCallbackAnswerRequest{}
)

func (g *MessagesGetBotCallbackAnswerRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Game == false) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}
	if !(g.Data == nil) {
		return false
	}
	if !(g.Password == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetBotCallbackAnswerRequest) String() string {
	if g == nil {
		return "MessagesGetBotCallbackAnswerRequest(nil)"
	}
	type Alias MessagesGetBotCallbackAnswerRequest
	return fmt.Sprintf("MessagesGetBotCallbackAnswerRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetBotCallbackAnswerRequest from given interface.
func (g *MessagesGetBotCallbackAnswerRequest) FillFrom(from interface {
	GetGame() (value bool)
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetData() (value []byte, ok bool)
	GetPassword() (value InputCheckPasswordSRPClass, ok bool)
}) {
	g.Game = from.GetGame()
	g.Peer = from.GetPeer()
	g.MsgID = from.GetMsgID()
	if val, ok := from.GetData(); ok {
		g.Data = val
	}

	if val, ok := from.GetPassword(); ok {
		g.Password = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetBotCallbackAnswerRequest) TypeID() uint32 {
	return MessagesGetBotCallbackAnswerRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetBotCallbackAnswerRequest) TypeName() string {
	return "messages.getBotCallbackAnswer"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetBotCallbackAnswerRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getBotCallbackAnswer",
		ID:   MessagesGetBotCallbackAnswerRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Game",
			SchemaName: "game",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "Data",
			SchemaName: "data",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Password",
			SchemaName: "password",
			Null:       !g.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *MessagesGetBotCallbackAnswerRequest) SetFlags() {
	if !(g.Game == false) {
		g.Flags.Set(1)
	}
	if !(g.Data == nil) {
		g.Flags.Set(0)
	}
	if !(g.Password == nil) {
		g.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (g *MessagesGetBotCallbackAnswerRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getBotCallbackAnswer#9342ca07 as nil")
	}
	b.PutID(MessagesGetBotCallbackAnswerRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetBotCallbackAnswerRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getBotCallbackAnswer#9342ca07 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	if g.Flags.Has(0) {
		b.PutBytes(g.Data)
	}
	if g.Flags.Has(2) {
		if g.Password == nil {
			return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field password is nil")
		}
		if err := g.Password.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field password: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetBotCallbackAnswerRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getBotCallbackAnswer#9342ca07 to nil")
	}
	if err := b.ConsumeID(MessagesGetBotCallbackAnswerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetBotCallbackAnswerRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getBotCallbackAnswer#9342ca07 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field flags: %w", err)
		}
	}
	g.Game = g.Flags.Has(1)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	if g.Flags.Has(0) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field data: %w", err)
		}
		g.Data = value
	}
	if g.Flags.Has(2) {
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field password: %w", err)
		}
		g.Password = value
	}
	return nil
}

// SetGame sets value of Game conditional field.
func (g *MessagesGetBotCallbackAnswerRequest) SetGame(value bool) {
	if value {
		g.Flags.Set(1)
		g.Game = true
	} else {
		g.Flags.Unset(1)
		g.Game = false
	}
}

// GetGame returns value of Game conditional field.
func (g *MessagesGetBotCallbackAnswerRequest) GetGame() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(1)
}

// GetPeer returns value of Peer field.
func (g *MessagesGetBotCallbackAnswerRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetMsgID returns value of MsgID field.
func (g *MessagesGetBotCallbackAnswerRequest) GetMsgID() (value int) {
	if g == nil {
		return
	}
	return g.MsgID
}

// SetData sets value of Data conditional field.
func (g *MessagesGetBotCallbackAnswerRequest) SetData(value []byte) {
	g.Flags.Set(0)
	g.Data = value
}

// GetData returns value of Data conditional field and
// boolean which is true if field was set.
func (g *MessagesGetBotCallbackAnswerRequest) GetData() (value []byte, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.Data, true
}

// SetPassword sets value of Password conditional field.
func (g *MessagesGetBotCallbackAnswerRequest) SetPassword(value InputCheckPasswordSRPClass) {
	g.Flags.Set(2)
	g.Password = value
}

// GetPassword returns value of Password conditional field and
// boolean which is true if field was set.
func (g *MessagesGetBotCallbackAnswerRequest) GetPassword() (value InputCheckPasswordSRPClass, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(2) {
		return value, false
	}
	return g.Password, true
}

// GetPasswordAsNotEmpty returns mapped value of Password conditional field and
// boolean which is true if field was set.
func (g *MessagesGetBotCallbackAnswerRequest) GetPasswordAsNotEmpty() (*InputCheckPasswordSRP, bool) {
	if value, ok := g.GetPassword(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// MessagesGetBotCallbackAnswer invokes method messages.getBotCallbackAnswer#9342ca07 returning error if any.
// Press an inline callback button and get a callback answer from the bot
//
// Possible errors:
//
//	400 BOT_RESPONSE_TIMEOUT: A timeout occurred while fetching data from the bot.
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 DATA_INVALID: Encrypted data invalid.
//	400 MESSAGE_ID_INVALID: The provided message id is invalid.
//	400 PASSWORD_MISSING: You must enable 2FA before executing this operation.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	-503 Timeout: Timeout while fetching data.
//
// See https://core.telegram.org/method/messages.getBotCallbackAnswer for reference.
func (c *Client) MessagesGetBotCallbackAnswer(ctx context.Context, request *MessagesGetBotCallbackAnswerRequest) (*MessagesBotCallbackAnswer, error) {
	var result MessagesBotCallbackAnswer

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
