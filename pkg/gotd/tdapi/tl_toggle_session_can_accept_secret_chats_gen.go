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

// ToggleSessionCanAcceptSecretChatsRequest represents TL type `toggleSessionCanAcceptSecretChats#3ba7a87e`.
type ToggleSessionCanAcceptSecretChatsRequest struct {
	// Session identifier
	SessionID int64
	// Pass true to allow accepting secret chats by the session; pass false otherwise
	CanAcceptSecretChats bool
}

// ToggleSessionCanAcceptSecretChatsRequestTypeID is TL type id of ToggleSessionCanAcceptSecretChatsRequest.
const ToggleSessionCanAcceptSecretChatsRequestTypeID = 0x3ba7a87e

// Ensuring interfaces in compile-time for ToggleSessionCanAcceptSecretChatsRequest.
var (
	_ bin.Encoder     = &ToggleSessionCanAcceptSecretChatsRequest{}
	_ bin.Decoder     = &ToggleSessionCanAcceptSecretChatsRequest{}
	_ bin.BareEncoder = &ToggleSessionCanAcceptSecretChatsRequest{}
	_ bin.BareDecoder = &ToggleSessionCanAcceptSecretChatsRequest{}
)

func (t *ToggleSessionCanAcceptSecretChatsRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.SessionID == 0) {
		return false
	}
	if !(t.CanAcceptSecretChats == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleSessionCanAcceptSecretChatsRequest) String() string {
	if t == nil {
		return "ToggleSessionCanAcceptSecretChatsRequest(nil)"
	}
	type Alias ToggleSessionCanAcceptSecretChatsRequest
	return fmt.Sprintf("ToggleSessionCanAcceptSecretChatsRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleSessionCanAcceptSecretChatsRequest) TypeID() uint32 {
	return ToggleSessionCanAcceptSecretChatsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleSessionCanAcceptSecretChatsRequest) TypeName() string {
	return "toggleSessionCanAcceptSecretChats"
}

// TypeInfo returns info about TL type.
func (t *ToggleSessionCanAcceptSecretChatsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleSessionCanAcceptSecretChats",
		ID:   ToggleSessionCanAcceptSecretChatsRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SessionID",
			SchemaName: "session_id",
		},
		{
			Name:       "CanAcceptSecretChats",
			SchemaName: "can_accept_secret_chats",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleSessionCanAcceptSecretChatsRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSessionCanAcceptSecretChats#3ba7a87e as nil")
	}
	b.PutID(ToggleSessionCanAcceptSecretChatsRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleSessionCanAcceptSecretChatsRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSessionCanAcceptSecretChats#3ba7a87e as nil")
	}
	b.PutLong(t.SessionID)
	b.PutBool(t.CanAcceptSecretChats)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleSessionCanAcceptSecretChatsRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSessionCanAcceptSecretChats#3ba7a87e to nil")
	}
	if err := b.ConsumeID(ToggleSessionCanAcceptSecretChatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleSessionCanAcceptSecretChats#3ba7a87e: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleSessionCanAcceptSecretChatsRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSessionCanAcceptSecretChats#3ba7a87e to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode toggleSessionCanAcceptSecretChats#3ba7a87e: field session_id: %w", err)
		}
		t.SessionID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleSessionCanAcceptSecretChats#3ba7a87e: field can_accept_secret_chats: %w", err)
		}
		t.CanAcceptSecretChats = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleSessionCanAcceptSecretChatsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSessionCanAcceptSecretChats#3ba7a87e as nil")
	}
	b.ObjStart()
	b.PutID("toggleSessionCanAcceptSecretChats")
	b.Comma()
	b.FieldStart("session_id")
	b.PutLong(t.SessionID)
	b.Comma()
	b.FieldStart("can_accept_secret_chats")
	b.PutBool(t.CanAcceptSecretChats)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleSessionCanAcceptSecretChatsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSessionCanAcceptSecretChats#3ba7a87e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleSessionCanAcceptSecretChats"); err != nil {
				return fmt.Errorf("unable to decode toggleSessionCanAcceptSecretChats#3ba7a87e: %w", err)
			}
		case "session_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode toggleSessionCanAcceptSecretChats#3ba7a87e: field session_id: %w", err)
			}
			t.SessionID = value
		case "can_accept_secret_chats":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleSessionCanAcceptSecretChats#3ba7a87e: field can_accept_secret_chats: %w", err)
			}
			t.CanAcceptSecretChats = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSessionID returns value of SessionID field.
func (t *ToggleSessionCanAcceptSecretChatsRequest) GetSessionID() (value int64) {
	if t == nil {
		return
	}
	return t.SessionID
}

// GetCanAcceptSecretChats returns value of CanAcceptSecretChats field.
func (t *ToggleSessionCanAcceptSecretChatsRequest) GetCanAcceptSecretChats() (value bool) {
	if t == nil {
		return
	}
	return t.CanAcceptSecretChats
}

// ToggleSessionCanAcceptSecretChats invokes method toggleSessionCanAcceptSecretChats#3ba7a87e returning error if any.
func (c *Client) ToggleSessionCanAcceptSecretChats(ctx context.Context, request *ToggleSessionCanAcceptSecretChatsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
