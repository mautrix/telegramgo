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

// TestReturnErrorRequest represents TL type `testReturnError#1b217cf2`.
type TestReturnErrorRequest struct {
	// The error to be returned
	Error Error
}

// TestReturnErrorRequestTypeID is TL type id of TestReturnErrorRequest.
const TestReturnErrorRequestTypeID = 0x1b217cf2

// Ensuring interfaces in compile-time for TestReturnErrorRequest.
var (
	_ bin.Encoder     = &TestReturnErrorRequest{}
	_ bin.Decoder     = &TestReturnErrorRequest{}
	_ bin.BareEncoder = &TestReturnErrorRequest{}
	_ bin.BareDecoder = &TestReturnErrorRequest{}
)

func (t *TestReturnErrorRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Error.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestReturnErrorRequest) String() string {
	if t == nil {
		return "TestReturnErrorRequest(nil)"
	}
	type Alias TestReturnErrorRequest
	return fmt.Sprintf("TestReturnErrorRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestReturnErrorRequest) TypeID() uint32 {
	return TestReturnErrorRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*TestReturnErrorRequest) TypeName() string {
	return "testReturnError"
}

// TypeInfo returns info about TL type.
func (t *TestReturnErrorRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testReturnError",
		ID:   TestReturnErrorRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Error",
			SchemaName: "error",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TestReturnErrorRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testReturnError#1b217cf2 as nil")
	}
	b.PutID(TestReturnErrorRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TestReturnErrorRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testReturnError#1b217cf2 as nil")
	}
	if err := t.Error.Encode(b); err != nil {
		return fmt.Errorf("unable to encode testReturnError#1b217cf2: field error: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TestReturnErrorRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testReturnError#1b217cf2 to nil")
	}
	if err := b.ConsumeID(TestReturnErrorRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode testReturnError#1b217cf2: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TestReturnErrorRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testReturnError#1b217cf2 to nil")
	}
	{
		if err := t.Error.Decode(b); err != nil {
			return fmt.Errorf("unable to decode testReturnError#1b217cf2: field error: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TestReturnErrorRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode testReturnError#1b217cf2 as nil")
	}
	b.ObjStart()
	b.PutID("testReturnError")
	b.Comma()
	b.FieldStart("error")
	if err := t.Error.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode testReturnError#1b217cf2: field error: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TestReturnErrorRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode testReturnError#1b217cf2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("testReturnError"); err != nil {
				return fmt.Errorf("unable to decode testReturnError#1b217cf2: %w", err)
			}
		case "error":
			if err := t.Error.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode testReturnError#1b217cf2: field error: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetError returns value of Error field.
func (t *TestReturnErrorRequest) GetError() (value Error) {
	if t == nil {
		return
	}
	return t.Error
}

// TestReturnError invokes method testReturnError#1b217cf2 returning error if any.
func (c *Client) TestReturnError(ctx context.Context, error Error) (*Error, error) {
	var result Error

	request := &TestReturnErrorRequest{
		Error: error,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
