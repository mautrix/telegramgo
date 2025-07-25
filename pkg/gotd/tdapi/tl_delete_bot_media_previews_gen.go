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

// DeleteBotMediaPreviewsRequest represents TL type `deleteBotMediaPreviews#f7d259db`.
type DeleteBotMediaPreviewsRequest struct {
	// Identifier of the target bot. The bot must be owned and must have the main Web App
	BotUserID int64
	// Language code of the media previews to delete
	LanguageCode string
	// File identifiers of the media to delete
	FileIDs []int32
}

// DeleteBotMediaPreviewsRequestTypeID is TL type id of DeleteBotMediaPreviewsRequest.
const DeleteBotMediaPreviewsRequestTypeID = 0xf7d259db

// Ensuring interfaces in compile-time for DeleteBotMediaPreviewsRequest.
var (
	_ bin.Encoder     = &DeleteBotMediaPreviewsRequest{}
	_ bin.Decoder     = &DeleteBotMediaPreviewsRequest{}
	_ bin.BareEncoder = &DeleteBotMediaPreviewsRequest{}
	_ bin.BareDecoder = &DeleteBotMediaPreviewsRequest{}
)

func (d *DeleteBotMediaPreviewsRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.BotUserID == 0) {
		return false
	}
	if !(d.LanguageCode == "") {
		return false
	}
	if !(d.FileIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteBotMediaPreviewsRequest) String() string {
	if d == nil {
		return "DeleteBotMediaPreviewsRequest(nil)"
	}
	type Alias DeleteBotMediaPreviewsRequest
	return fmt.Sprintf("DeleteBotMediaPreviewsRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteBotMediaPreviewsRequest) TypeID() uint32 {
	return DeleteBotMediaPreviewsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteBotMediaPreviewsRequest) TypeName() string {
	return "deleteBotMediaPreviews"
}

// TypeInfo returns info about TL type.
func (d *DeleteBotMediaPreviewsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteBotMediaPreviews",
		ID:   DeleteBotMediaPreviewsRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "LanguageCode",
			SchemaName: "language_code",
		},
		{
			Name:       "FileIDs",
			SchemaName: "file_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteBotMediaPreviewsRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteBotMediaPreviews#f7d259db as nil")
	}
	b.PutID(DeleteBotMediaPreviewsRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteBotMediaPreviewsRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteBotMediaPreviews#f7d259db as nil")
	}
	b.PutInt53(d.BotUserID)
	b.PutString(d.LanguageCode)
	b.PutInt(len(d.FileIDs))
	for _, v := range d.FileIDs {
		b.PutInt32(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteBotMediaPreviewsRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteBotMediaPreviews#f7d259db to nil")
	}
	if err := b.ConsumeID(DeleteBotMediaPreviewsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteBotMediaPreviews#f7d259db: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteBotMediaPreviewsRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteBotMediaPreviews#f7d259db to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode deleteBotMediaPreviews#f7d259db: field bot_user_id: %w", err)
		}
		d.BotUserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode deleteBotMediaPreviews#f7d259db: field language_code: %w", err)
		}
		d.LanguageCode = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode deleteBotMediaPreviews#f7d259db: field file_ids: %w", err)
		}

		if headerLen > 0 {
			d.FileIDs = make([]int32, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode deleteBotMediaPreviews#f7d259db: field file_ids: %w", err)
			}
			d.FileIDs = append(d.FileIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteBotMediaPreviewsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteBotMediaPreviews#f7d259db as nil")
	}
	b.ObjStart()
	b.PutID("deleteBotMediaPreviews")
	b.Comma()
	b.FieldStart("bot_user_id")
	b.PutInt53(d.BotUserID)
	b.Comma()
	b.FieldStart("language_code")
	b.PutString(d.LanguageCode)
	b.Comma()
	b.FieldStart("file_ids")
	b.ArrStart()
	for _, v := range d.FileIDs {
		b.PutInt32(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteBotMediaPreviewsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteBotMediaPreviews#f7d259db to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteBotMediaPreviews"); err != nil {
				return fmt.Errorf("unable to decode deleteBotMediaPreviews#f7d259db: %w", err)
			}
		case "bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode deleteBotMediaPreviews#f7d259db: field bot_user_id: %w", err)
			}
			d.BotUserID = value
		case "language_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode deleteBotMediaPreviews#f7d259db: field language_code: %w", err)
			}
			d.LanguageCode = value
		case "file_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int32()
				if err != nil {
					return fmt.Errorf("unable to decode deleteBotMediaPreviews#f7d259db: field file_ids: %w", err)
				}
				d.FileIDs = append(d.FileIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode deleteBotMediaPreviews#f7d259db: field file_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBotUserID returns value of BotUserID field.
func (d *DeleteBotMediaPreviewsRequest) GetBotUserID() (value int64) {
	if d == nil {
		return
	}
	return d.BotUserID
}

// GetLanguageCode returns value of LanguageCode field.
func (d *DeleteBotMediaPreviewsRequest) GetLanguageCode() (value string) {
	if d == nil {
		return
	}
	return d.LanguageCode
}

// GetFileIDs returns value of FileIDs field.
func (d *DeleteBotMediaPreviewsRequest) GetFileIDs() (value []int32) {
	if d == nil {
		return
	}
	return d.FileIDs
}

// DeleteBotMediaPreviews invokes method deleteBotMediaPreviews#f7d259db returning error if any.
func (c *Client) DeleteBotMediaPreviews(ctx context.Context, request *DeleteBotMediaPreviewsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
