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

// ReorderChatFoldersRequest represents TL type `reorderChatFolders#59544c32`.
type ReorderChatFoldersRequest struct {
	// Identifiers of chat folders in the new correct order
	ChatFolderIDs []int32
	// Position of the main chat list among chat folders, 0-based. Can be non-zero only for
	// Premium users
	MainChatListPosition int32
}

// ReorderChatFoldersRequestTypeID is TL type id of ReorderChatFoldersRequest.
const ReorderChatFoldersRequestTypeID = 0x59544c32

// Ensuring interfaces in compile-time for ReorderChatFoldersRequest.
var (
	_ bin.Encoder     = &ReorderChatFoldersRequest{}
	_ bin.Decoder     = &ReorderChatFoldersRequest{}
	_ bin.BareEncoder = &ReorderChatFoldersRequest{}
	_ bin.BareDecoder = &ReorderChatFoldersRequest{}
)

func (r *ReorderChatFoldersRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatFolderIDs == nil) {
		return false
	}
	if !(r.MainChatListPosition == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReorderChatFoldersRequest) String() string {
	if r == nil {
		return "ReorderChatFoldersRequest(nil)"
	}
	type Alias ReorderChatFoldersRequest
	return fmt.Sprintf("ReorderChatFoldersRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReorderChatFoldersRequest) TypeID() uint32 {
	return ReorderChatFoldersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReorderChatFoldersRequest) TypeName() string {
	return "reorderChatFolders"
}

// TypeInfo returns info about TL type.
func (r *ReorderChatFoldersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reorderChatFolders",
		ID:   ReorderChatFoldersRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatFolderIDs",
			SchemaName: "chat_folder_ids",
		},
		{
			Name:       "MainChatListPosition",
			SchemaName: "main_chat_list_position",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReorderChatFoldersRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reorderChatFolders#59544c32 as nil")
	}
	b.PutID(ReorderChatFoldersRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReorderChatFoldersRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reorderChatFolders#59544c32 as nil")
	}
	b.PutInt(len(r.ChatFolderIDs))
	for _, v := range r.ChatFolderIDs {
		b.PutInt32(v)
	}
	b.PutInt32(r.MainChatListPosition)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReorderChatFoldersRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reorderChatFolders#59544c32 to nil")
	}
	if err := b.ConsumeID(ReorderChatFoldersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode reorderChatFolders#59544c32: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReorderChatFoldersRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reorderChatFolders#59544c32 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode reorderChatFolders#59544c32: field chat_folder_ids: %w", err)
		}

		if headerLen > 0 {
			r.ChatFolderIDs = make([]int32, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode reorderChatFolders#59544c32: field chat_folder_ids: %w", err)
			}
			r.ChatFolderIDs = append(r.ChatFolderIDs, value)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode reorderChatFolders#59544c32: field main_chat_list_position: %w", err)
		}
		r.MainChatListPosition = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReorderChatFoldersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reorderChatFolders#59544c32 as nil")
	}
	b.ObjStart()
	b.PutID("reorderChatFolders")
	b.Comma()
	b.FieldStart("chat_folder_ids")
	b.ArrStart()
	for _, v := range r.ChatFolderIDs {
		b.PutInt32(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("main_chat_list_position")
	b.PutInt32(r.MainChatListPosition)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReorderChatFoldersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reorderChatFolders#59544c32 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reorderChatFolders"); err != nil {
				return fmt.Errorf("unable to decode reorderChatFolders#59544c32: %w", err)
			}
		case "chat_folder_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int32()
				if err != nil {
					return fmt.Errorf("unable to decode reorderChatFolders#59544c32: field chat_folder_ids: %w", err)
				}
				r.ChatFolderIDs = append(r.ChatFolderIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode reorderChatFolders#59544c32: field chat_folder_ids: %w", err)
			}
		case "main_chat_list_position":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode reorderChatFolders#59544c32: field main_chat_list_position: %w", err)
			}
			r.MainChatListPosition = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatFolderIDs returns value of ChatFolderIDs field.
func (r *ReorderChatFoldersRequest) GetChatFolderIDs() (value []int32) {
	if r == nil {
		return
	}
	return r.ChatFolderIDs
}

// GetMainChatListPosition returns value of MainChatListPosition field.
func (r *ReorderChatFoldersRequest) GetMainChatListPosition() (value int32) {
	if r == nil {
		return
	}
	return r.MainChatListPosition
}

// ReorderChatFolders invokes method reorderChatFolders#59544c32 returning error if any.
func (c *Client) ReorderChatFolders(ctx context.Context, request *ReorderChatFoldersRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
