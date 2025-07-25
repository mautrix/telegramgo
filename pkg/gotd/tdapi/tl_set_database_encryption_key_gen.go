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

// SetDatabaseEncryptionKeyRequest represents TL type `setDatabaseEncryptionKey#b83345b5`.
type SetDatabaseEncryptionKeyRequest struct {
	// New encryption key
	NewEncryptionKey []byte
}

// SetDatabaseEncryptionKeyRequestTypeID is TL type id of SetDatabaseEncryptionKeyRequest.
const SetDatabaseEncryptionKeyRequestTypeID = 0xb83345b5

// Ensuring interfaces in compile-time for SetDatabaseEncryptionKeyRequest.
var (
	_ bin.Encoder     = &SetDatabaseEncryptionKeyRequest{}
	_ bin.Decoder     = &SetDatabaseEncryptionKeyRequest{}
	_ bin.BareEncoder = &SetDatabaseEncryptionKeyRequest{}
	_ bin.BareDecoder = &SetDatabaseEncryptionKeyRequest{}
)

func (s *SetDatabaseEncryptionKeyRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.NewEncryptionKey == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetDatabaseEncryptionKeyRequest) String() string {
	if s == nil {
		return "SetDatabaseEncryptionKeyRequest(nil)"
	}
	type Alias SetDatabaseEncryptionKeyRequest
	return fmt.Sprintf("SetDatabaseEncryptionKeyRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetDatabaseEncryptionKeyRequest) TypeID() uint32 {
	return SetDatabaseEncryptionKeyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetDatabaseEncryptionKeyRequest) TypeName() string {
	return "setDatabaseEncryptionKey"
}

// TypeInfo returns info about TL type.
func (s *SetDatabaseEncryptionKeyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setDatabaseEncryptionKey",
		ID:   SetDatabaseEncryptionKeyRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "NewEncryptionKey",
			SchemaName: "new_encryption_key",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetDatabaseEncryptionKeyRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setDatabaseEncryptionKey#b83345b5 as nil")
	}
	b.PutID(SetDatabaseEncryptionKeyRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetDatabaseEncryptionKeyRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setDatabaseEncryptionKey#b83345b5 as nil")
	}
	b.PutBytes(s.NewEncryptionKey)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetDatabaseEncryptionKeyRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setDatabaseEncryptionKey#b83345b5 to nil")
	}
	if err := b.ConsumeID(SetDatabaseEncryptionKeyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setDatabaseEncryptionKey#b83345b5: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetDatabaseEncryptionKeyRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setDatabaseEncryptionKey#b83345b5 to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode setDatabaseEncryptionKey#b83345b5: field new_encryption_key: %w", err)
		}
		s.NewEncryptionKey = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetDatabaseEncryptionKeyRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setDatabaseEncryptionKey#b83345b5 as nil")
	}
	b.ObjStart()
	b.PutID("setDatabaseEncryptionKey")
	b.Comma()
	b.FieldStart("new_encryption_key")
	b.PutBytes(s.NewEncryptionKey)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetDatabaseEncryptionKeyRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setDatabaseEncryptionKey#b83345b5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setDatabaseEncryptionKey"); err != nil {
				return fmt.Errorf("unable to decode setDatabaseEncryptionKey#b83345b5: %w", err)
			}
		case "new_encryption_key":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode setDatabaseEncryptionKey#b83345b5: field new_encryption_key: %w", err)
			}
			s.NewEncryptionKey = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetNewEncryptionKey returns value of NewEncryptionKey field.
func (s *SetDatabaseEncryptionKeyRequest) GetNewEncryptionKey() (value []byte) {
	if s == nil {
		return
	}
	return s.NewEncryptionKey
}

// SetDatabaseEncryptionKey invokes method setDatabaseEncryptionKey#b83345b5 returning error if any.
func (c *Client) SetDatabaseEncryptionKey(ctx context.Context, newencryptionkey []byte) error {
	var ok Ok

	request := &SetDatabaseEncryptionKeyRequest{
		NewEncryptionKey: newencryptionkey,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
