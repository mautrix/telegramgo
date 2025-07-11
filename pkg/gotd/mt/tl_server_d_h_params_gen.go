// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// ServerDHParamsFail represents TL type `server_DH_params_fail#79cb045d`.
type ServerDHParamsFail struct {
	// Nonce field of ServerDHParamsFail.
	Nonce bin.Int128
	// ServerNonce field of ServerDHParamsFail.
	ServerNonce bin.Int128
	// NewNonceHash field of ServerDHParamsFail.
	NewNonceHash bin.Int128
}

// ServerDHParamsFailTypeID is TL type id of ServerDHParamsFail.
const ServerDHParamsFailTypeID = 0x79cb045d

// construct implements constructor of ServerDHParamsClass.
func (s ServerDHParamsFail) construct() ServerDHParamsClass { return &s }

// Ensuring interfaces in compile-time for ServerDHParamsFail.
var (
	_ bin.Encoder     = &ServerDHParamsFail{}
	_ bin.Decoder     = &ServerDHParamsFail{}
	_ bin.BareEncoder = &ServerDHParamsFail{}
	_ bin.BareDecoder = &ServerDHParamsFail{}

	_ ServerDHParamsClass = &ServerDHParamsFail{}
)

func (s *ServerDHParamsFail) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Nonce == bin.Int128{}) {
		return false
	}
	if !(s.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(s.NewNonceHash == bin.Int128{}) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ServerDHParamsFail) String() string {
	if s == nil {
		return "ServerDHParamsFail(nil)"
	}
	type Alias ServerDHParamsFail
	return fmt.Sprintf("ServerDHParamsFail%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ServerDHParamsFail) TypeID() uint32 {
	return ServerDHParamsFailTypeID
}

// TypeName returns name of type in TL schema.
func (*ServerDHParamsFail) TypeName() string {
	return "server_DH_params_fail"
}

// TypeInfo returns info about TL type.
func (s *ServerDHParamsFail) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "server_DH_params_fail",
		ID:   ServerDHParamsFailTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Nonce",
			SchemaName: "nonce",
		},
		{
			Name:       "ServerNonce",
			SchemaName: "server_nonce",
		},
		{
			Name:       "NewNonceHash",
			SchemaName: "new_nonce_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *ServerDHParamsFail) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode server_DH_params_fail#79cb045d as nil")
	}
	b.PutID(ServerDHParamsFailTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *ServerDHParamsFail) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode server_DH_params_fail#79cb045d as nil")
	}
	b.PutInt128(s.Nonce)
	b.PutInt128(s.ServerNonce)
	b.PutInt128(s.NewNonceHash)
	return nil
}

// Decode implements bin.Decoder.
func (s *ServerDHParamsFail) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode server_DH_params_fail#79cb045d to nil")
	}
	if err := b.ConsumeID(ServerDHParamsFailTypeID); err != nil {
		return fmt.Errorf("unable to decode server_DH_params_fail#79cb045d: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *ServerDHParamsFail) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode server_DH_params_fail#79cb045d to nil")
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_fail#79cb045d: field nonce: %w", err)
		}
		s.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_fail#79cb045d: field server_nonce: %w", err)
		}
		s.ServerNonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_fail#79cb045d: field new_nonce_hash: %w", err)
		}
		s.NewNonceHash = value
	}
	return nil
}

// GetNonce returns value of Nonce field.
func (s *ServerDHParamsFail) GetNonce() (value bin.Int128) {
	if s == nil {
		return
	}
	return s.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (s *ServerDHParamsFail) GetServerNonce() (value bin.Int128) {
	if s == nil {
		return
	}
	return s.ServerNonce
}

// GetNewNonceHash returns value of NewNonceHash field.
func (s *ServerDHParamsFail) GetNewNonceHash() (value bin.Int128) {
	if s == nil {
		return
	}
	return s.NewNonceHash
}

// ServerDHParamsOk represents TL type `server_DH_params_ok#d0e8075c`.
type ServerDHParamsOk struct {
	// Nonce field of ServerDHParamsOk.
	Nonce bin.Int128
	// ServerNonce field of ServerDHParamsOk.
	ServerNonce bin.Int128
	// EncryptedAnswer field of ServerDHParamsOk.
	EncryptedAnswer []byte
}

// ServerDHParamsOkTypeID is TL type id of ServerDHParamsOk.
const ServerDHParamsOkTypeID = 0xd0e8075c

// construct implements constructor of ServerDHParamsClass.
func (s ServerDHParamsOk) construct() ServerDHParamsClass { return &s }

// Ensuring interfaces in compile-time for ServerDHParamsOk.
var (
	_ bin.Encoder     = &ServerDHParamsOk{}
	_ bin.Decoder     = &ServerDHParamsOk{}
	_ bin.BareEncoder = &ServerDHParamsOk{}
	_ bin.BareDecoder = &ServerDHParamsOk{}

	_ ServerDHParamsClass = &ServerDHParamsOk{}
)

func (s *ServerDHParamsOk) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Nonce == bin.Int128{}) {
		return false
	}
	if !(s.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(s.EncryptedAnswer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ServerDHParamsOk) String() string {
	if s == nil {
		return "ServerDHParamsOk(nil)"
	}
	type Alias ServerDHParamsOk
	return fmt.Sprintf("ServerDHParamsOk%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ServerDHParamsOk) TypeID() uint32 {
	return ServerDHParamsOkTypeID
}

// TypeName returns name of type in TL schema.
func (*ServerDHParamsOk) TypeName() string {
	return "server_DH_params_ok"
}

// TypeInfo returns info about TL type.
func (s *ServerDHParamsOk) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "server_DH_params_ok",
		ID:   ServerDHParamsOkTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Nonce",
			SchemaName: "nonce",
		},
		{
			Name:       "ServerNonce",
			SchemaName: "server_nonce",
		},
		{
			Name:       "EncryptedAnswer",
			SchemaName: "encrypted_answer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *ServerDHParamsOk) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode server_DH_params_ok#d0e8075c as nil")
	}
	b.PutID(ServerDHParamsOkTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *ServerDHParamsOk) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode server_DH_params_ok#d0e8075c as nil")
	}
	b.PutInt128(s.Nonce)
	b.PutInt128(s.ServerNonce)
	b.PutBytes(s.EncryptedAnswer)
	return nil
}

// Decode implements bin.Decoder.
func (s *ServerDHParamsOk) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode server_DH_params_ok#d0e8075c to nil")
	}
	if err := b.ConsumeID(ServerDHParamsOkTypeID); err != nil {
		return fmt.Errorf("unable to decode server_DH_params_ok#d0e8075c: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *ServerDHParamsOk) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode server_DH_params_ok#d0e8075c to nil")
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_ok#d0e8075c: field nonce: %w", err)
		}
		s.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_ok#d0e8075c: field server_nonce: %w", err)
		}
		s.ServerNonce = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_ok#d0e8075c: field encrypted_answer: %w", err)
		}
		s.EncryptedAnswer = value
	}
	return nil
}

// GetNonce returns value of Nonce field.
func (s *ServerDHParamsOk) GetNonce() (value bin.Int128) {
	if s == nil {
		return
	}
	return s.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (s *ServerDHParamsOk) GetServerNonce() (value bin.Int128) {
	if s == nil {
		return
	}
	return s.ServerNonce
}

// GetEncryptedAnswer returns value of EncryptedAnswer field.
func (s *ServerDHParamsOk) GetEncryptedAnswer() (value []byte) {
	if s == nil {
		return
	}
	return s.EncryptedAnswer
}

// ServerDHParamsClassName is schema name of ServerDHParamsClass.
const ServerDHParamsClassName = "Server_DH_Params"

// ServerDHParamsClass represents Server_DH_Params generic type.
//
// Example:
//
//	g, err := mt.DecodeServerDHParams(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *mt.ServerDHParamsFail: // server_DH_params_fail#79cb045d
//	case *mt.ServerDHParamsOk: // server_DH_params_ok#d0e8075c
//	default: panic(v)
//	}
type ServerDHParamsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ServerDHParamsClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Nonce field of ServerDHParamsFail.
	GetNonce() (value bin.Int128)
	// ServerNonce field of ServerDHParamsFail.
	GetServerNonce() (value bin.Int128)
}

// DecodeServerDHParams implements binary de-serialization for ServerDHParamsClass.
func DecodeServerDHParams(buf *bin.Buffer) (ServerDHParamsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ServerDHParamsFailTypeID:
		// Decoding server_DH_params_fail#79cb045d.
		v := ServerDHParamsFail{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ServerDHParamsClass: %w", err)
		}
		return &v, nil
	case ServerDHParamsOkTypeID:
		// Decoding server_DH_params_ok#d0e8075c.
		v := ServerDHParamsOk{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ServerDHParamsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ServerDHParamsClass: %w", bin.NewUnexpectedID(id))
	}
}

// ServerDHParams boxes the ServerDHParamsClass providing a helper.
type ServerDHParamsBox struct {
	Server_DH_Params ServerDHParamsClass
}

// Decode implements bin.Decoder for ServerDHParamsBox.
func (b *ServerDHParamsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ServerDHParamsBox to nil")
	}
	v, err := DecodeServerDHParams(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Server_DH_Params = v
	return nil
}

// Encode implements bin.Encode for ServerDHParamsBox.
func (b *ServerDHParamsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Server_DH_Params == nil {
		return fmt.Errorf("unable to encode ServerDHParamsClass as nil")
	}
	return b.Server_DH_Params.Encode(buf)
}
