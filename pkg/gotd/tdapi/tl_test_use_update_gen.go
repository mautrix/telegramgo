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

// TestUseUpdateRequest represents TL type `testUseUpdate#2abdff1e`.
type TestUseUpdateRequest struct {
}

// TestUseUpdateRequestTypeID is TL type id of TestUseUpdateRequest.
const TestUseUpdateRequestTypeID = 0x2abdff1e

// Ensuring interfaces in compile-time for TestUseUpdateRequest.
var (
	_ bin.Encoder     = &TestUseUpdateRequest{}
	_ bin.Decoder     = &TestUseUpdateRequest{}
	_ bin.BareEncoder = &TestUseUpdateRequest{}
	_ bin.BareDecoder = &TestUseUpdateRequest{}
)

func (t *TestUseUpdateRequest) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestUseUpdateRequest) String() string {
	if t == nil {
		return "TestUseUpdateRequest(nil)"
	}
	type Alias TestUseUpdateRequest
	return fmt.Sprintf("TestUseUpdateRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestUseUpdateRequest) TypeID() uint32 {
	return TestUseUpdateRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*TestUseUpdateRequest) TypeName() string {
	return "testUseUpdate"
}

// TypeInfo returns info about TL type.
func (t *TestUseUpdateRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testUseUpdate",
		ID:   TestUseUpdateRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *TestUseUpdateRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testUseUpdate#2abdff1e as nil")
	}
	b.PutID(TestUseUpdateRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TestUseUpdateRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testUseUpdate#2abdff1e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TestUseUpdateRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testUseUpdate#2abdff1e to nil")
	}
	if err := b.ConsumeID(TestUseUpdateRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode testUseUpdate#2abdff1e: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TestUseUpdateRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testUseUpdate#2abdff1e to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TestUseUpdateRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode testUseUpdate#2abdff1e as nil")
	}
	b.ObjStart()
	b.PutID("testUseUpdate")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TestUseUpdateRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode testUseUpdate#2abdff1e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("testUseUpdate"); err != nil {
				return fmt.Errorf("unable to decode testUseUpdate#2abdff1e: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// TestUseUpdate invokes method testUseUpdate#2abdff1e returning error if any.
func (c *Client) TestUseUpdate(ctx context.Context) (UpdateClass, error) {
	var result UpdateBox

	request := &TestUseUpdateRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Update, nil
}
