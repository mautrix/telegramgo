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

// EditBusinessChatLinkRequest represents TL type `editBusinessChatLink#5f10f626`.
type EditBusinessChatLinkRequest struct {
	// The link to edit
	Link string
	// New description of the link
	LinkInfo InputBusinessChatLink
}

// EditBusinessChatLinkRequestTypeID is TL type id of EditBusinessChatLinkRequest.
const EditBusinessChatLinkRequestTypeID = 0x5f10f626

// Ensuring interfaces in compile-time for EditBusinessChatLinkRequest.
var (
	_ bin.Encoder     = &EditBusinessChatLinkRequest{}
	_ bin.Decoder     = &EditBusinessChatLinkRequest{}
	_ bin.BareEncoder = &EditBusinessChatLinkRequest{}
	_ bin.BareDecoder = &EditBusinessChatLinkRequest{}
)

func (e *EditBusinessChatLinkRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Link == "") {
		return false
	}
	if !(e.LinkInfo.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditBusinessChatLinkRequest) String() string {
	if e == nil {
		return "EditBusinessChatLinkRequest(nil)"
	}
	type Alias EditBusinessChatLinkRequest
	return fmt.Sprintf("EditBusinessChatLinkRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditBusinessChatLinkRequest) TypeID() uint32 {
	return EditBusinessChatLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditBusinessChatLinkRequest) TypeName() string {
	return "editBusinessChatLink"
}

// TypeInfo returns info about TL type.
func (e *EditBusinessChatLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editBusinessChatLink",
		ID:   EditBusinessChatLinkRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Link",
			SchemaName: "link",
		},
		{
			Name:       "LinkInfo",
			SchemaName: "link_info",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditBusinessChatLinkRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editBusinessChatLink#5f10f626 as nil")
	}
	b.PutID(EditBusinessChatLinkRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditBusinessChatLinkRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editBusinessChatLink#5f10f626 as nil")
	}
	b.PutString(e.Link)
	if err := e.LinkInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editBusinessChatLink#5f10f626: field link_info: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EditBusinessChatLinkRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editBusinessChatLink#5f10f626 to nil")
	}
	if err := b.ConsumeID(EditBusinessChatLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editBusinessChatLink#5f10f626: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditBusinessChatLinkRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editBusinessChatLink#5f10f626 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode editBusinessChatLink#5f10f626: field link: %w", err)
		}
		e.Link = value
	}
	{
		if err := e.LinkInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode editBusinessChatLink#5f10f626: field link_info: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditBusinessChatLinkRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editBusinessChatLink#5f10f626 as nil")
	}
	b.ObjStart()
	b.PutID("editBusinessChatLink")
	b.Comma()
	b.FieldStart("link")
	b.PutString(e.Link)
	b.Comma()
	b.FieldStart("link_info")
	if err := e.LinkInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editBusinessChatLink#5f10f626: field link_info: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditBusinessChatLinkRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editBusinessChatLink#5f10f626 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editBusinessChatLink"); err != nil {
				return fmt.Errorf("unable to decode editBusinessChatLink#5f10f626: %w", err)
			}
		case "link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode editBusinessChatLink#5f10f626: field link: %w", err)
			}
			e.Link = value
		case "link_info":
			if err := e.LinkInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode editBusinessChatLink#5f10f626: field link_info: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLink returns value of Link field.
func (e *EditBusinessChatLinkRequest) GetLink() (value string) {
	if e == nil {
		return
	}
	return e.Link
}

// GetLinkInfo returns value of LinkInfo field.
func (e *EditBusinessChatLinkRequest) GetLinkInfo() (value InputBusinessChatLink) {
	if e == nil {
		return
	}
	return e.LinkInfo
}

// EditBusinessChatLink invokes method editBusinessChatLink#5f10f626 returning error if any.
func (c *Client) EditBusinessChatLink(ctx context.Context, request *EditBusinessChatLinkRequest) (*BusinessChatLink, error) {
	var result BusinessChatLink

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
