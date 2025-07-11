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

// UploadGetCDNFileHashesRequest represents TL type `upload.getCdnFileHashes#91dc3f31`.
// Get SHA256 hashes for verifying downloaded CDN¹ files
//
// Links:
//  1. https://core.telegram.org/cdn
//
// See https://core.telegram.org/method/upload.getCdnFileHashes for reference.
type UploadGetCDNFileHashesRequest struct {
	// File
	FileToken []byte
	// Offset from which to start getting hashes
	Offset int64
}

// UploadGetCDNFileHashesRequestTypeID is TL type id of UploadGetCDNFileHashesRequest.
const UploadGetCDNFileHashesRequestTypeID = 0x91dc3f31

// Ensuring interfaces in compile-time for UploadGetCDNFileHashesRequest.
var (
	_ bin.Encoder     = &UploadGetCDNFileHashesRequest{}
	_ bin.Decoder     = &UploadGetCDNFileHashesRequest{}
	_ bin.BareEncoder = &UploadGetCDNFileHashesRequest{}
	_ bin.BareDecoder = &UploadGetCDNFileHashesRequest{}
)

func (g *UploadGetCDNFileHashesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.FileToken == nil) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *UploadGetCDNFileHashesRequest) String() string {
	if g == nil {
		return "UploadGetCDNFileHashesRequest(nil)"
	}
	type Alias UploadGetCDNFileHashesRequest
	return fmt.Sprintf("UploadGetCDNFileHashesRequest%+v", Alias(*g))
}

// FillFrom fills UploadGetCDNFileHashesRequest from given interface.
func (g *UploadGetCDNFileHashesRequest) FillFrom(from interface {
	GetFileToken() (value []byte)
	GetOffset() (value int64)
}) {
	g.FileToken = from.GetFileToken()
	g.Offset = from.GetOffset()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadGetCDNFileHashesRequest) TypeID() uint32 {
	return UploadGetCDNFileHashesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadGetCDNFileHashesRequest) TypeName() string {
	return "upload.getCdnFileHashes"
}

// TypeInfo returns info about TL type.
func (g *UploadGetCDNFileHashesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upload.getCdnFileHashes",
		ID:   UploadGetCDNFileHashesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileToken",
			SchemaName: "file_token",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *UploadGetCDNFileHashesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode upload.getCdnFileHashes#91dc3f31 as nil")
	}
	b.PutID(UploadGetCDNFileHashesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *UploadGetCDNFileHashesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode upload.getCdnFileHashes#91dc3f31 as nil")
	}
	b.PutBytes(g.FileToken)
	b.PutLong(g.Offset)
	return nil
}

// Decode implements bin.Decoder.
func (g *UploadGetCDNFileHashesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode upload.getCdnFileHashes#91dc3f31 to nil")
	}
	if err := b.ConsumeID(UploadGetCDNFileHashesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.getCdnFileHashes#91dc3f31: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *UploadGetCDNFileHashesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode upload.getCdnFileHashes#91dc3f31 to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getCdnFileHashes#91dc3f31: field file_token: %w", err)
		}
		g.FileToken = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getCdnFileHashes#91dc3f31: field offset: %w", err)
		}
		g.Offset = value
	}
	return nil
}

// GetFileToken returns value of FileToken field.
func (g *UploadGetCDNFileHashesRequest) GetFileToken() (value []byte) {
	if g == nil {
		return
	}
	return g.FileToken
}

// GetOffset returns value of Offset field.
func (g *UploadGetCDNFileHashesRequest) GetOffset() (value int64) {
	if g == nil {
		return
	}
	return g.Offset
}

// UploadGetCDNFileHashes invokes method upload.getCdnFileHashes#91dc3f31 returning error if any.
// Get SHA256 hashes for verifying downloaded CDN¹ files
//
// Links:
//  1. https://core.telegram.org/cdn
//
// Possible errors:
//
//	400 CDN_METHOD_INVALID: You can't call this method in a CDN DC.
//	400 FILE_TOKEN_INVALID: The master DC did not accept the file_token (e.g., the token has expired). Continue downloading the file from the master DC using upload.getFile.
//	400 RSA_DECRYPT_FAILED: Internal RSA decryption failed.
//
// See https://core.telegram.org/method/upload.getCdnFileHashes for reference.
// Can be used by bots.
func (c *Client) UploadGetCDNFileHashes(ctx context.Context, request *UploadGetCDNFileHashesRequest) ([]FileHash, error) {
	var result FileHashVector

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []FileHash(result.Elems), nil
}
