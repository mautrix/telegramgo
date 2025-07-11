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

// SetBirthdateRequest represents TL type `setBirthdate#4ea9dd98`.
type SetBirthdateRequest struct {
	// The new value of the current user's birthdate; pass null to remove the birthdate
	Birthdate Birthdate
}

// SetBirthdateRequestTypeID is TL type id of SetBirthdateRequest.
const SetBirthdateRequestTypeID = 0x4ea9dd98

// Ensuring interfaces in compile-time for SetBirthdateRequest.
var (
	_ bin.Encoder     = &SetBirthdateRequest{}
	_ bin.Decoder     = &SetBirthdateRequest{}
	_ bin.BareEncoder = &SetBirthdateRequest{}
	_ bin.BareDecoder = &SetBirthdateRequest{}
)

func (s *SetBirthdateRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Birthdate.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetBirthdateRequest) String() string {
	if s == nil {
		return "SetBirthdateRequest(nil)"
	}
	type Alias SetBirthdateRequest
	return fmt.Sprintf("SetBirthdateRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetBirthdateRequest) TypeID() uint32 {
	return SetBirthdateRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetBirthdateRequest) TypeName() string {
	return "setBirthdate"
}

// TypeInfo returns info about TL type.
func (s *SetBirthdateRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setBirthdate",
		ID:   SetBirthdateRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Birthdate",
			SchemaName: "birthdate",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetBirthdateRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBirthdate#4ea9dd98 as nil")
	}
	b.PutID(SetBirthdateRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetBirthdateRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBirthdate#4ea9dd98 as nil")
	}
	if err := s.Birthdate.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setBirthdate#4ea9dd98: field birthdate: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetBirthdateRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBirthdate#4ea9dd98 to nil")
	}
	if err := b.ConsumeID(SetBirthdateRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setBirthdate#4ea9dd98: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetBirthdateRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBirthdate#4ea9dd98 to nil")
	}
	{
		if err := s.Birthdate.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setBirthdate#4ea9dd98: field birthdate: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetBirthdateRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setBirthdate#4ea9dd98 as nil")
	}
	b.ObjStart()
	b.PutID("setBirthdate")
	b.Comma()
	b.FieldStart("birthdate")
	if err := s.Birthdate.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setBirthdate#4ea9dd98: field birthdate: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetBirthdateRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setBirthdate#4ea9dd98 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setBirthdate"); err != nil {
				return fmt.Errorf("unable to decode setBirthdate#4ea9dd98: %w", err)
			}
		case "birthdate":
			if err := s.Birthdate.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode setBirthdate#4ea9dd98: field birthdate: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBirthdate returns value of Birthdate field.
func (s *SetBirthdateRequest) GetBirthdate() (value Birthdate) {
	if s == nil {
		return
	}
	return s.Birthdate
}

// SetBirthdate invokes method setBirthdate#4ea9dd98 returning error if any.
func (c *Client) SetBirthdate(ctx context.Context, birthdate Birthdate) error {
	var ok Ok

	request := &SetBirthdateRequest{
		Birthdate: birthdate,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
