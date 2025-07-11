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

// EncryptedCredentials represents TL type `encryptedCredentials#4f5713ce`.
type EncryptedCredentials struct {
	// The encrypted credentials
	Data []byte
	// The decrypted data hash
	Hash []byte
	// Secret for data decryption, encrypted with the service's public key
	Secret []byte
}

// EncryptedCredentialsTypeID is TL type id of EncryptedCredentials.
const EncryptedCredentialsTypeID = 0x4f5713ce

// Ensuring interfaces in compile-time for EncryptedCredentials.
var (
	_ bin.Encoder     = &EncryptedCredentials{}
	_ bin.Decoder     = &EncryptedCredentials{}
	_ bin.BareEncoder = &EncryptedCredentials{}
	_ bin.BareDecoder = &EncryptedCredentials{}
)

func (e *EncryptedCredentials) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Data == nil) {
		return false
	}
	if !(e.Hash == nil) {
		return false
	}
	if !(e.Secret == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedCredentials) String() string {
	if e == nil {
		return "EncryptedCredentials(nil)"
	}
	type Alias EncryptedCredentials
	return fmt.Sprintf("EncryptedCredentials%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EncryptedCredentials) TypeID() uint32 {
	return EncryptedCredentialsTypeID
}

// TypeName returns name of type in TL schema.
func (*EncryptedCredentials) TypeName() string {
	return "encryptedCredentials"
}

// TypeInfo returns info about TL type.
func (e *EncryptedCredentials) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "encryptedCredentials",
		ID:   EncryptedCredentialsTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Data",
			SchemaName: "data",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Secret",
			SchemaName: "secret",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EncryptedCredentials) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedCredentials#4f5713ce as nil")
	}
	b.PutID(EncryptedCredentialsTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EncryptedCredentials) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedCredentials#4f5713ce as nil")
	}
	b.PutBytes(e.Data)
	b.PutBytes(e.Hash)
	b.PutBytes(e.Secret)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedCredentials) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedCredentials#4f5713ce to nil")
	}
	if err := b.ConsumeID(EncryptedCredentialsTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedCredentials#4f5713ce: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EncryptedCredentials) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedCredentials#4f5713ce to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedCredentials#4f5713ce: field data: %w", err)
		}
		e.Data = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedCredentials#4f5713ce: field hash: %w", err)
		}
		e.Hash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedCredentials#4f5713ce: field secret: %w", err)
		}
		e.Secret = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EncryptedCredentials) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedCredentials#4f5713ce as nil")
	}
	b.ObjStart()
	b.PutID("encryptedCredentials")
	b.Comma()
	b.FieldStart("data")
	b.PutBytes(e.Data)
	b.Comma()
	b.FieldStart("hash")
	b.PutBytes(e.Hash)
	b.Comma()
	b.FieldStart("secret")
	b.PutBytes(e.Secret)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EncryptedCredentials) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedCredentials#4f5713ce to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("encryptedCredentials"); err != nil {
				return fmt.Errorf("unable to decode encryptedCredentials#4f5713ce: %w", err)
			}
		case "data":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode encryptedCredentials#4f5713ce: field data: %w", err)
			}
			e.Data = value
		case "hash":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode encryptedCredentials#4f5713ce: field hash: %w", err)
			}
			e.Hash = value
		case "secret":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode encryptedCredentials#4f5713ce: field secret: %w", err)
			}
			e.Secret = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetData returns value of Data field.
func (e *EncryptedCredentials) GetData() (value []byte) {
	if e == nil {
		return
	}
	return e.Data
}

// GetHash returns value of Hash field.
func (e *EncryptedCredentials) GetHash() (value []byte) {
	if e == nil {
		return
	}
	return e.Hash
}

// GetSecret returns value of Secret field.
func (e *EncryptedCredentials) GetSecret() (value []byte) {
	if e == nil {
		return
	}
	return e.Secret
}
