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

// SetUsernameRequest represents TL type `setUsername#1a385c1e`.
type SetUsernameRequest struct {
	// The new value of the username. Use an empty string to remove the username. The
	// username can't be completely removed if there is another active or disabled username
	Username string
}

// SetUsernameRequestTypeID is TL type id of SetUsernameRequest.
const SetUsernameRequestTypeID = 0x1a385c1e

// Ensuring interfaces in compile-time for SetUsernameRequest.
var (
	_ bin.Encoder     = &SetUsernameRequest{}
	_ bin.Decoder     = &SetUsernameRequest{}
	_ bin.BareEncoder = &SetUsernameRequest{}
	_ bin.BareDecoder = &SetUsernameRequest{}
)

func (s *SetUsernameRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Username == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetUsernameRequest) String() string {
	if s == nil {
		return "SetUsernameRequest(nil)"
	}
	type Alias SetUsernameRequest
	return fmt.Sprintf("SetUsernameRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetUsernameRequest) TypeID() uint32 {
	return SetUsernameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetUsernameRequest) TypeName() string {
	return "setUsername"
}

// TypeInfo returns info about TL type.
func (s *SetUsernameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setUsername",
		ID:   SetUsernameRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Username",
			SchemaName: "username",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetUsernameRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setUsername#1a385c1e as nil")
	}
	b.PutID(SetUsernameRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetUsernameRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setUsername#1a385c1e as nil")
	}
	b.PutString(s.Username)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetUsernameRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setUsername#1a385c1e to nil")
	}
	if err := b.ConsumeID(SetUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setUsername#1a385c1e: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetUsernameRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setUsername#1a385c1e to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setUsername#1a385c1e: field username: %w", err)
		}
		s.Username = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetUsernameRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setUsername#1a385c1e as nil")
	}
	b.ObjStart()
	b.PutID("setUsername")
	b.Comma()
	b.FieldStart("username")
	b.PutString(s.Username)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetUsernameRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setUsername#1a385c1e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setUsername"); err != nil {
				return fmt.Errorf("unable to decode setUsername#1a385c1e: %w", err)
			}
		case "username":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setUsername#1a385c1e: field username: %w", err)
			}
			s.Username = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUsername returns value of Username field.
func (s *SetUsernameRequest) GetUsername() (value string) {
	if s == nil {
		return
	}
	return s.Username
}

// SetUsername invokes method setUsername#1a385c1e returning error if any.
func (c *Client) SetUsername(ctx context.Context, username string) error {
	var ok Ok

	request := &SetUsernameRequest{
		Username: username,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
