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

// GetLogStreamRequest represents TL type `getLogStream#45984b5b`.
type GetLogStreamRequest struct {
}

// GetLogStreamRequestTypeID is TL type id of GetLogStreamRequest.
const GetLogStreamRequestTypeID = 0x45984b5b

// Ensuring interfaces in compile-time for GetLogStreamRequest.
var (
	_ bin.Encoder     = &GetLogStreamRequest{}
	_ bin.Decoder     = &GetLogStreamRequest{}
	_ bin.BareEncoder = &GetLogStreamRequest{}
	_ bin.BareDecoder = &GetLogStreamRequest{}
)

func (g *GetLogStreamRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetLogStreamRequest) String() string {
	if g == nil {
		return "GetLogStreamRequest(nil)"
	}
	type Alias GetLogStreamRequest
	return fmt.Sprintf("GetLogStreamRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetLogStreamRequest) TypeID() uint32 {
	return GetLogStreamRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetLogStreamRequest) TypeName() string {
	return "getLogStream"
}

// TypeInfo returns info about TL type.
func (g *GetLogStreamRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getLogStream",
		ID:   GetLogStreamRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetLogStreamRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLogStream#45984b5b as nil")
	}
	b.PutID(GetLogStreamRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetLogStreamRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLogStream#45984b5b as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetLogStreamRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLogStream#45984b5b to nil")
	}
	if err := b.ConsumeID(GetLogStreamRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getLogStream#45984b5b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetLogStreamRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLogStream#45984b5b to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetLogStreamRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getLogStream#45984b5b as nil")
	}
	b.ObjStart()
	b.PutID("getLogStream")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetLogStreamRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getLogStream#45984b5b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getLogStream"); err != nil {
				return fmt.Errorf("unable to decode getLogStream#45984b5b: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLogStream invokes method getLogStream#45984b5b returning error if any.
func (c *Client) GetLogStream(ctx context.Context) (LogStreamClass, error) {
	var result LogStreamBox

	request := &GetLogStreamRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.LogStream, nil
}
