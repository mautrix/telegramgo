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

// ChannelsEditCreatorRequest represents TL type `channels.editCreator#8f38cd1f`.
// Transfer channel ownership
//
// See https://core.telegram.org/method/channels.editCreator for reference.
type ChannelsEditCreatorRequest struct {
	// Channel
	Channel InputChannelClass
	// New channel owner
	UserID InputUserClass
	// 2FA password¹ of account
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	Password InputCheckPasswordSRPClass
}

// ChannelsEditCreatorRequestTypeID is TL type id of ChannelsEditCreatorRequest.
const ChannelsEditCreatorRequestTypeID = 0x8f38cd1f

// Ensuring interfaces in compile-time for ChannelsEditCreatorRequest.
var (
	_ bin.Encoder     = &ChannelsEditCreatorRequest{}
	_ bin.Decoder     = &ChannelsEditCreatorRequest{}
	_ bin.BareEncoder = &ChannelsEditCreatorRequest{}
	_ bin.BareDecoder = &ChannelsEditCreatorRequest{}
)

func (e *ChannelsEditCreatorRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Channel == nil) {
		return false
	}
	if !(e.UserID == nil) {
		return false
	}
	if !(e.Password == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *ChannelsEditCreatorRequest) String() string {
	if e == nil {
		return "ChannelsEditCreatorRequest(nil)"
	}
	type Alias ChannelsEditCreatorRequest
	return fmt.Sprintf("ChannelsEditCreatorRequest%+v", Alias(*e))
}

// FillFrom fills ChannelsEditCreatorRequest from given interface.
func (e *ChannelsEditCreatorRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetUserID() (value InputUserClass)
	GetPassword() (value InputCheckPasswordSRPClass)
}) {
	e.Channel = from.GetChannel()
	e.UserID = from.GetUserID()
	e.Password = from.GetPassword()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsEditCreatorRequest) TypeID() uint32 {
	return ChannelsEditCreatorRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsEditCreatorRequest) TypeName() string {
	return "channels.editCreator"
}

// TypeInfo returns info about TL type.
func (e *ChannelsEditCreatorRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.editCreator",
		ID:   ChannelsEditCreatorRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Password",
			SchemaName: "password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *ChannelsEditCreatorRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editCreator#8f38cd1f as nil")
	}
	b.PutID(ChannelsEditCreatorRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *ChannelsEditCreatorRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editCreator#8f38cd1f as nil")
	}
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field channel: %w", err)
	}
	if e.UserID == nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field user_id is nil")
	}
	if err := e.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field user_id: %w", err)
	}
	if e.Password == nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field password is nil")
	}
	if err := e.Password.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editCreator#8f38cd1f: field password: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *ChannelsEditCreatorRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editCreator#8f38cd1f to nil")
	}
	if err := b.ConsumeID(ChannelsEditCreatorRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.editCreator#8f38cd1f: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *ChannelsEditCreatorRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editCreator#8f38cd1f to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editCreator#8f38cd1f: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editCreator#8f38cd1f: field user_id: %w", err)
		}
		e.UserID = value
	}
	{
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editCreator#8f38cd1f: field password: %w", err)
		}
		e.Password = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (e *ChannelsEditCreatorRequest) GetChannel() (value InputChannelClass) {
	if e == nil {
		return
	}
	return e.Channel
}

// GetUserID returns value of UserID field.
func (e *ChannelsEditCreatorRequest) GetUserID() (value InputUserClass) {
	if e == nil {
		return
	}
	return e.UserID
}

// GetPassword returns value of Password field.
func (e *ChannelsEditCreatorRequest) GetPassword() (value InputCheckPasswordSRPClass) {
	if e == nil {
		return
	}
	return e.Password
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (e *ChannelsEditCreatorRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return e.Channel.AsNotEmpty()
}

// GetPasswordAsNotEmpty returns mapped value of Password field.
func (e *ChannelsEditCreatorRequest) GetPasswordAsNotEmpty() (*InputCheckPasswordSRP, bool) {
	return e.Password.AsNotEmpty()
}

// ChannelsEditCreator invokes method channels.editCreator#8f38cd1f returning error if any.
// Transfer channel ownership
//
// Possible errors:
//
//	400 CHANNELS_ADMIN_PUBLIC_TOO_MUCH: You're admin of too many public channels, make some channels private to change the username of this channel.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 CHAT_NOT_MODIFIED: No changes were made to chat information because the new information you passed is identical to the current information.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	400 PASSWORD_HASH_INVALID: The provided password hash is invalid.
//	400 PASSWORD_MISSING: You must enable 2FA before executing this operation.
//	400 PASSWORD_TOO_FRESH_%d: The password was modified less than 24 hours ago, try again in %d seconds.
//	400 SESSION_TOO_FRESH_%d: This session was created less than 24 hours ago, try again in %d seconds.
//	400 SRP_ID_INVALID: Invalid SRP ID provided.
//	400 USER_ID_INVALID: The provided user ID is invalid.
//
// See https://core.telegram.org/method/channels.editCreator for reference.
func (c *Client) ChannelsEditCreator(ctx context.Context, request *ChannelsEditCreatorRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
