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

// ReplaceVideoChatRtmpURLRequest represents TL type `replaceVideoChatRtmpUrl#214f8fe0`.
type ReplaceVideoChatRtmpURLRequest struct {
	// Chat identifier
	ChatID int64
}

// ReplaceVideoChatRtmpURLRequestTypeID is TL type id of ReplaceVideoChatRtmpURLRequest.
const ReplaceVideoChatRtmpURLRequestTypeID = 0x214f8fe0

// Ensuring interfaces in compile-time for ReplaceVideoChatRtmpURLRequest.
var (
	_ bin.Encoder     = &ReplaceVideoChatRtmpURLRequest{}
	_ bin.Decoder     = &ReplaceVideoChatRtmpURLRequest{}
	_ bin.BareEncoder = &ReplaceVideoChatRtmpURLRequest{}
	_ bin.BareDecoder = &ReplaceVideoChatRtmpURLRequest{}
)

func (r *ReplaceVideoChatRtmpURLRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplaceVideoChatRtmpURLRequest) String() string {
	if r == nil {
		return "ReplaceVideoChatRtmpURLRequest(nil)"
	}
	type Alias ReplaceVideoChatRtmpURLRequest
	return fmt.Sprintf("ReplaceVideoChatRtmpURLRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplaceVideoChatRtmpURLRequest) TypeID() uint32 {
	return ReplaceVideoChatRtmpURLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplaceVideoChatRtmpURLRequest) TypeName() string {
	return "replaceVideoChatRtmpUrl"
}

// TypeInfo returns info about TL type.
func (r *ReplaceVideoChatRtmpURLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replaceVideoChatRtmpUrl",
		ID:   ReplaceVideoChatRtmpURLRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReplaceVideoChatRtmpURLRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replaceVideoChatRtmpUrl#214f8fe0 as nil")
	}
	b.PutID(ReplaceVideoChatRtmpURLRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReplaceVideoChatRtmpURLRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replaceVideoChatRtmpUrl#214f8fe0 as nil")
	}
	b.PutInt53(r.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReplaceVideoChatRtmpURLRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replaceVideoChatRtmpUrl#214f8fe0 to nil")
	}
	if err := b.ConsumeID(ReplaceVideoChatRtmpURLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode replaceVideoChatRtmpUrl#214f8fe0: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReplaceVideoChatRtmpURLRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replaceVideoChatRtmpUrl#214f8fe0 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode replaceVideoChatRtmpUrl#214f8fe0: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReplaceVideoChatRtmpURLRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode replaceVideoChatRtmpUrl#214f8fe0 as nil")
	}
	b.ObjStart()
	b.PutID("replaceVideoChatRtmpUrl")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(r.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReplaceVideoChatRtmpURLRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode replaceVideoChatRtmpUrl#214f8fe0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("replaceVideoChatRtmpUrl"); err != nil {
				return fmt.Errorf("unable to decode replaceVideoChatRtmpUrl#214f8fe0: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode replaceVideoChatRtmpUrl#214f8fe0: field chat_id: %w", err)
			}
			r.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (r *ReplaceVideoChatRtmpURLRequest) GetChatID() (value int64) {
	if r == nil {
		return
	}
	return r.ChatID
}

// ReplaceVideoChatRtmpURL invokes method replaceVideoChatRtmpUrl#214f8fe0 returning error if any.
func (c *Client) ReplaceVideoChatRtmpURL(ctx context.Context, chatid int64) (*RtmpURL, error) {
	var result RtmpURL

	request := &ReplaceVideoChatRtmpURLRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
