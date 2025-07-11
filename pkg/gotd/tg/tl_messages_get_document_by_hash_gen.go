// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesGetDocumentByHashRequest represents TL type `messages.getDocumentByHash#b1f2061f`.
// Get a document by its SHA256 hash, mainly used for gifs
//
// See https://core.telegram.org/method/messages.getDocumentByHash for reference.
type MessagesGetDocumentByHashRequest struct {
	// SHA256 of file
	SHA256 []byte
	// Size of the file in bytes
	Size int64
	// Mime type
	MimeType string
}

// MessagesGetDocumentByHashRequestTypeID is TL type id of MessagesGetDocumentByHashRequest.
const MessagesGetDocumentByHashRequestTypeID = 0xb1f2061f

// Ensuring interfaces in compile-time for MessagesGetDocumentByHashRequest.
var (
	_ bin.Encoder     = &MessagesGetDocumentByHashRequest{}
	_ bin.Decoder     = &MessagesGetDocumentByHashRequest{}
	_ bin.BareEncoder = &MessagesGetDocumentByHashRequest{}
	_ bin.BareDecoder = &MessagesGetDocumentByHashRequest{}
)

func (g *MessagesGetDocumentByHashRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.SHA256 == nil) {
		return false
	}
	if !(g.Size == 0) {
		return false
	}
	if !(g.MimeType == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetDocumentByHashRequest) String() string {
	if g == nil {
		return "MessagesGetDocumentByHashRequest(nil)"
	}
	type Alias MessagesGetDocumentByHashRequest
	return fmt.Sprintf("MessagesGetDocumentByHashRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetDocumentByHashRequest from given interface.
func (g *MessagesGetDocumentByHashRequest) FillFrom(from interface {
	GetSHA256() (value []byte)
	GetSize() (value int64)
	GetMimeType() (value string)
}) {
	g.SHA256 = from.GetSHA256()
	g.Size = from.GetSize()
	g.MimeType = from.GetMimeType()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetDocumentByHashRequest) TypeID() uint32 {
	return MessagesGetDocumentByHashRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetDocumentByHashRequest) TypeName() string {
	return "messages.getDocumentByHash"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetDocumentByHashRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getDocumentByHash",
		ID:   MessagesGetDocumentByHashRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SHA256",
			SchemaName: "sha256",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "MimeType",
			SchemaName: "mime_type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetDocumentByHashRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDocumentByHash#b1f2061f as nil")
	}
	b.PutID(MessagesGetDocumentByHashRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetDocumentByHashRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDocumentByHash#b1f2061f as nil")
	}
	b.PutBytes(g.SHA256)
	b.PutLong(g.Size)
	b.PutString(g.MimeType)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetDocumentByHashRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDocumentByHash#b1f2061f to nil")
	}
	if err := b.ConsumeID(MessagesGetDocumentByHashRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDocumentByHash#b1f2061f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetDocumentByHashRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDocumentByHash#b1f2061f to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDocumentByHash#b1f2061f: field sha256: %w", err)
		}
		g.SHA256 = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDocumentByHash#b1f2061f: field size: %w", err)
		}
		g.Size = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDocumentByHash#b1f2061f: field mime_type: %w", err)
		}
		g.MimeType = value
	}
	return nil
}

// GetSHA256 returns value of SHA256 field.
func (g *MessagesGetDocumentByHashRequest) GetSHA256() (value []byte) {
	if g == nil {
		return
	}
	return g.SHA256
}

// GetSize returns value of Size field.
func (g *MessagesGetDocumentByHashRequest) GetSize() (value int64) {
	if g == nil {
		return
	}
	return g.Size
}

// GetMimeType returns value of MimeType field.
func (g *MessagesGetDocumentByHashRequest) GetMimeType() (value string) {
	if g == nil {
		return
	}
	return g.MimeType
}

// MessagesGetDocumentByHash invokes method messages.getDocumentByHash#b1f2061f returning error if any.
// Get a document by its SHA256 hash, mainly used for gifs
//
// Possible errors:
//
//	400 SHA256_HASH_INVALID: The provided SHA256 hash is invalid.
//
// See https://core.telegram.org/method/messages.getDocumentByHash for reference.
// Can be used by bots.
func (c *Client) MessagesGetDocumentByHash(ctx context.Context, request *MessagesGetDocumentByHashRequest) (DocumentClass, error) {
	var result DocumentBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Document, nil
}
