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

// SearchUserByTokenRequest represents TL type `searchUserByToken#d841f436`.
type SearchUserByTokenRequest struct {
	// Token to search for
	Token string
}

// SearchUserByTokenRequestTypeID is TL type id of SearchUserByTokenRequest.
const SearchUserByTokenRequestTypeID = 0xd841f436

// Ensuring interfaces in compile-time for SearchUserByTokenRequest.
var (
	_ bin.Encoder     = &SearchUserByTokenRequest{}
	_ bin.Decoder     = &SearchUserByTokenRequest{}
	_ bin.BareEncoder = &SearchUserByTokenRequest{}
	_ bin.BareDecoder = &SearchUserByTokenRequest{}
)

func (s *SearchUserByTokenRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Token == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchUserByTokenRequest) String() string {
	if s == nil {
		return "SearchUserByTokenRequest(nil)"
	}
	type Alias SearchUserByTokenRequest
	return fmt.Sprintf("SearchUserByTokenRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchUserByTokenRequest) TypeID() uint32 {
	return SearchUserByTokenRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchUserByTokenRequest) TypeName() string {
	return "searchUserByToken"
}

// TypeInfo returns info about TL type.
func (s *SearchUserByTokenRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchUserByToken",
		ID:   SearchUserByTokenRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Token",
			SchemaName: "token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchUserByTokenRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchUserByToken#d841f436 as nil")
	}
	b.PutID(SearchUserByTokenRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchUserByTokenRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchUserByToken#d841f436 as nil")
	}
	b.PutString(s.Token)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchUserByTokenRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchUserByToken#d841f436 to nil")
	}
	if err := b.ConsumeID(SearchUserByTokenRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchUserByToken#d841f436: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchUserByTokenRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchUserByToken#d841f436 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchUserByToken#d841f436: field token: %w", err)
		}
		s.Token = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SearchUserByTokenRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchUserByToken#d841f436 as nil")
	}
	b.ObjStart()
	b.PutID("searchUserByToken")
	b.Comma()
	b.FieldStart("token")
	b.PutString(s.Token)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SearchUserByTokenRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchUserByToken#d841f436 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("searchUserByToken"); err != nil {
				return fmt.Errorf("unable to decode searchUserByToken#d841f436: %w", err)
			}
		case "token":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchUserByToken#d841f436: field token: %w", err)
			}
			s.Token = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetToken returns value of Token field.
func (s *SearchUserByTokenRequest) GetToken() (value string) {
	if s == nil {
		return
	}
	return s.Token
}

// SearchUserByToken invokes method searchUserByToken#d841f436 returning error if any.
func (c *Client) SearchUserByToken(ctx context.Context, token string) (*User, error) {
	var result User

	request := &SearchUserByTokenRequest{
		Token: token,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
