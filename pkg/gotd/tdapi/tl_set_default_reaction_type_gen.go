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

// SetDefaultReactionTypeRequest represents TL type `setDefaultReactionType#65038a3d`.
type SetDefaultReactionTypeRequest struct {
	// New type of the default reaction. The paid reaction can't be set as default
	ReactionType ReactionTypeClass
}

// SetDefaultReactionTypeRequestTypeID is TL type id of SetDefaultReactionTypeRequest.
const SetDefaultReactionTypeRequestTypeID = 0x65038a3d

// Ensuring interfaces in compile-time for SetDefaultReactionTypeRequest.
var (
	_ bin.Encoder     = &SetDefaultReactionTypeRequest{}
	_ bin.Decoder     = &SetDefaultReactionTypeRequest{}
	_ bin.BareEncoder = &SetDefaultReactionTypeRequest{}
	_ bin.BareDecoder = &SetDefaultReactionTypeRequest{}
)

func (s *SetDefaultReactionTypeRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ReactionType == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetDefaultReactionTypeRequest) String() string {
	if s == nil {
		return "SetDefaultReactionTypeRequest(nil)"
	}
	type Alias SetDefaultReactionTypeRequest
	return fmt.Sprintf("SetDefaultReactionTypeRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetDefaultReactionTypeRequest) TypeID() uint32 {
	return SetDefaultReactionTypeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetDefaultReactionTypeRequest) TypeName() string {
	return "setDefaultReactionType"
}

// TypeInfo returns info about TL type.
func (s *SetDefaultReactionTypeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setDefaultReactionType",
		ID:   SetDefaultReactionTypeRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ReactionType",
			SchemaName: "reaction_type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetDefaultReactionTypeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setDefaultReactionType#65038a3d as nil")
	}
	b.PutID(SetDefaultReactionTypeRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetDefaultReactionTypeRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setDefaultReactionType#65038a3d as nil")
	}
	if s.ReactionType == nil {
		return fmt.Errorf("unable to encode setDefaultReactionType#65038a3d: field reaction_type is nil")
	}
	if err := s.ReactionType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setDefaultReactionType#65038a3d: field reaction_type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetDefaultReactionTypeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setDefaultReactionType#65038a3d to nil")
	}
	if err := b.ConsumeID(SetDefaultReactionTypeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setDefaultReactionType#65038a3d: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetDefaultReactionTypeRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setDefaultReactionType#65038a3d to nil")
	}
	{
		value, err := DecodeReactionType(b)
		if err != nil {
			return fmt.Errorf("unable to decode setDefaultReactionType#65038a3d: field reaction_type: %w", err)
		}
		s.ReactionType = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetDefaultReactionTypeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setDefaultReactionType#65038a3d as nil")
	}
	b.ObjStart()
	b.PutID("setDefaultReactionType")
	b.Comma()
	b.FieldStart("reaction_type")
	if s.ReactionType == nil {
		return fmt.Errorf("unable to encode setDefaultReactionType#65038a3d: field reaction_type is nil")
	}
	if err := s.ReactionType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setDefaultReactionType#65038a3d: field reaction_type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetDefaultReactionTypeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setDefaultReactionType#65038a3d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setDefaultReactionType"); err != nil {
				return fmt.Errorf("unable to decode setDefaultReactionType#65038a3d: %w", err)
			}
		case "reaction_type":
			value, err := DecodeTDLibJSONReactionType(b)
			if err != nil {
				return fmt.Errorf("unable to decode setDefaultReactionType#65038a3d: field reaction_type: %w", err)
			}
			s.ReactionType = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetReactionType returns value of ReactionType field.
func (s *SetDefaultReactionTypeRequest) GetReactionType() (value ReactionTypeClass) {
	if s == nil {
		return
	}
	return s.ReactionType
}

// SetDefaultReactionType invokes method setDefaultReactionType#65038a3d returning error if any.
func (c *Client) SetDefaultReactionType(ctx context.Context, reactiontype ReactionTypeClass) error {
	var ok Ok

	request := &SetDefaultReactionTypeRequest{
		ReactionType: reactiontype,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
