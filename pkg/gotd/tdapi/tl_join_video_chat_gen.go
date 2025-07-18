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

// JoinVideoChatRequest represents TL type `joinVideoChat#4edb39af`.
type JoinVideoChatRequest struct {
	// Group call identifier
	GroupCallID int32
	// Identifier of a group call participant, which will be used to join the call; pass null
	// to join as self; video chats only
	ParticipantID MessageSenderClass
	// Parameters to join the call
	JoinParameters GroupCallJoinParameters
	// Invite hash as received from internalLinkTypeVideoChat
	InviteHash string
}

// JoinVideoChatRequestTypeID is TL type id of JoinVideoChatRequest.
const JoinVideoChatRequestTypeID = 0x4edb39af

// Ensuring interfaces in compile-time for JoinVideoChatRequest.
var (
	_ bin.Encoder     = &JoinVideoChatRequest{}
	_ bin.Decoder     = &JoinVideoChatRequest{}
	_ bin.BareEncoder = &JoinVideoChatRequest{}
	_ bin.BareDecoder = &JoinVideoChatRequest{}
)

func (j *JoinVideoChatRequest) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.GroupCallID == 0) {
		return false
	}
	if !(j.ParticipantID == nil) {
		return false
	}
	if !(j.JoinParameters.Zero()) {
		return false
	}
	if !(j.InviteHash == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *JoinVideoChatRequest) String() string {
	if j == nil {
		return "JoinVideoChatRequest(nil)"
	}
	type Alias JoinVideoChatRequest
	return fmt.Sprintf("JoinVideoChatRequest%+v", Alias(*j))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*JoinVideoChatRequest) TypeID() uint32 {
	return JoinVideoChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*JoinVideoChatRequest) TypeName() string {
	return "joinVideoChat"
}

// TypeInfo returns info about TL type.
func (j *JoinVideoChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "joinVideoChat",
		ID:   JoinVideoChatRequestTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
		{
			Name:       "ParticipantID",
			SchemaName: "participant_id",
		},
		{
			Name:       "JoinParameters",
			SchemaName: "join_parameters",
		},
		{
			Name:       "InviteHash",
			SchemaName: "invite_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (j *JoinVideoChatRequest) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode joinVideoChat#4edb39af as nil")
	}
	b.PutID(JoinVideoChatRequestTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *JoinVideoChatRequest) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode joinVideoChat#4edb39af as nil")
	}
	b.PutInt32(j.GroupCallID)
	if j.ParticipantID == nil {
		return fmt.Errorf("unable to encode joinVideoChat#4edb39af: field participant_id is nil")
	}
	if err := j.ParticipantID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode joinVideoChat#4edb39af: field participant_id: %w", err)
	}
	if err := j.JoinParameters.Encode(b); err != nil {
		return fmt.Errorf("unable to encode joinVideoChat#4edb39af: field join_parameters: %w", err)
	}
	b.PutString(j.InviteHash)
	return nil
}

// Decode implements bin.Decoder.
func (j *JoinVideoChatRequest) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode joinVideoChat#4edb39af to nil")
	}
	if err := b.ConsumeID(JoinVideoChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode joinVideoChat#4edb39af: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *JoinVideoChatRequest) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode joinVideoChat#4edb39af to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode joinVideoChat#4edb39af: field group_call_id: %w", err)
		}
		j.GroupCallID = value
	}
	{
		value, err := DecodeMessageSender(b)
		if err != nil {
			return fmt.Errorf("unable to decode joinVideoChat#4edb39af: field participant_id: %w", err)
		}
		j.ParticipantID = value
	}
	{
		if err := j.JoinParameters.Decode(b); err != nil {
			return fmt.Errorf("unable to decode joinVideoChat#4edb39af: field join_parameters: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode joinVideoChat#4edb39af: field invite_hash: %w", err)
		}
		j.InviteHash = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (j *JoinVideoChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if j == nil {
		return fmt.Errorf("can't encode joinVideoChat#4edb39af as nil")
	}
	b.ObjStart()
	b.PutID("joinVideoChat")
	b.Comma()
	b.FieldStart("group_call_id")
	b.PutInt32(j.GroupCallID)
	b.Comma()
	b.FieldStart("participant_id")
	if j.ParticipantID == nil {
		return fmt.Errorf("unable to encode joinVideoChat#4edb39af: field participant_id is nil")
	}
	if err := j.ParticipantID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode joinVideoChat#4edb39af: field participant_id: %w", err)
	}
	b.Comma()
	b.FieldStart("join_parameters")
	if err := j.JoinParameters.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode joinVideoChat#4edb39af: field join_parameters: %w", err)
	}
	b.Comma()
	b.FieldStart("invite_hash")
	b.PutString(j.InviteHash)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (j *JoinVideoChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if j == nil {
		return fmt.Errorf("can't decode joinVideoChat#4edb39af to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("joinVideoChat"); err != nil {
				return fmt.Errorf("unable to decode joinVideoChat#4edb39af: %w", err)
			}
		case "group_call_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode joinVideoChat#4edb39af: field group_call_id: %w", err)
			}
			j.GroupCallID = value
		case "participant_id":
			value, err := DecodeTDLibJSONMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode joinVideoChat#4edb39af: field participant_id: %w", err)
			}
			j.ParticipantID = value
		case "join_parameters":
			if err := j.JoinParameters.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode joinVideoChat#4edb39af: field join_parameters: %w", err)
			}
		case "invite_hash":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode joinVideoChat#4edb39af: field invite_hash: %w", err)
			}
			j.InviteHash = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGroupCallID returns value of GroupCallID field.
func (j *JoinVideoChatRequest) GetGroupCallID() (value int32) {
	if j == nil {
		return
	}
	return j.GroupCallID
}

// GetParticipantID returns value of ParticipantID field.
func (j *JoinVideoChatRequest) GetParticipantID() (value MessageSenderClass) {
	if j == nil {
		return
	}
	return j.ParticipantID
}

// GetJoinParameters returns value of JoinParameters field.
func (j *JoinVideoChatRequest) GetJoinParameters() (value GroupCallJoinParameters) {
	if j == nil {
		return
	}
	return j.JoinParameters
}

// GetInviteHash returns value of InviteHash field.
func (j *JoinVideoChatRequest) GetInviteHash() (value string) {
	if j == nil {
		return
	}
	return j.InviteHash
}

// JoinVideoChat invokes method joinVideoChat#4edb39af returning error if any.
func (c *Client) JoinVideoChat(ctx context.Context, request *JoinVideoChatRequest) (*Text, error) {
	var result Text

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
