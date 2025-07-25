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

// StorageStatisticsByChat represents TL type `storageStatisticsByChat#a5498fe4`.
type StorageStatisticsByChat struct {
	// Chat identifier; 0 if none
	ChatID int64
	// Total size of the files in the chat, in bytes
	Size int64
	// Total number of files in the chat
	Count int32
	// Statistics split by file types
	ByFileType []StorageStatisticsByFileType
}

// StorageStatisticsByChatTypeID is TL type id of StorageStatisticsByChat.
const StorageStatisticsByChatTypeID = 0xa5498fe4

// Ensuring interfaces in compile-time for StorageStatisticsByChat.
var (
	_ bin.Encoder     = &StorageStatisticsByChat{}
	_ bin.Decoder     = &StorageStatisticsByChat{}
	_ bin.BareEncoder = &StorageStatisticsByChat{}
	_ bin.BareDecoder = &StorageStatisticsByChat{}
)

func (s *StorageStatisticsByChat) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.Size == 0) {
		return false
	}
	if !(s.Count == 0) {
		return false
	}
	if !(s.ByFileType == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StorageStatisticsByChat) String() string {
	if s == nil {
		return "StorageStatisticsByChat(nil)"
	}
	type Alias StorageStatisticsByChat
	return fmt.Sprintf("StorageStatisticsByChat%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StorageStatisticsByChat) TypeID() uint32 {
	return StorageStatisticsByChatTypeID
}

// TypeName returns name of type in TL schema.
func (*StorageStatisticsByChat) TypeName() string {
	return "storageStatisticsByChat"
}

// TypeInfo returns info about TL type.
func (s *StorageStatisticsByChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storageStatisticsByChat",
		ID:   StorageStatisticsByChatTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "ByFileType",
			SchemaName: "by_file_type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StorageStatisticsByChat) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storageStatisticsByChat#a5498fe4 as nil")
	}
	b.PutID(StorageStatisticsByChatTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StorageStatisticsByChat) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storageStatisticsByChat#a5498fe4 as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutInt53(s.Size)
	b.PutInt32(s.Count)
	b.PutInt(len(s.ByFileType))
	for idx, v := range s.ByFileType {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare storageStatisticsByChat#a5498fe4: field by_file_type element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StorageStatisticsByChat) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storageStatisticsByChat#a5498fe4 to nil")
	}
	if err := b.ConsumeID(StorageStatisticsByChatTypeID); err != nil {
		return fmt.Errorf("unable to decode storageStatisticsByChat#a5498fe4: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StorageStatisticsByChat) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storageStatisticsByChat#a5498fe4 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode storageStatisticsByChat#a5498fe4: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode storageStatisticsByChat#a5498fe4: field size: %w", err)
		}
		s.Size = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode storageStatisticsByChat#a5498fe4: field count: %w", err)
		}
		s.Count = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storageStatisticsByChat#a5498fe4: field by_file_type: %w", err)
		}

		if headerLen > 0 {
			s.ByFileType = make([]StorageStatisticsByFileType, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StorageStatisticsByFileType
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare storageStatisticsByChat#a5498fe4: field by_file_type: %w", err)
			}
			s.ByFileType = append(s.ByFileType, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StorageStatisticsByChat) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storageStatisticsByChat#a5498fe4 as nil")
	}
	b.ObjStart()
	b.PutID("storageStatisticsByChat")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("size")
	b.PutInt53(s.Size)
	b.Comma()
	b.FieldStart("count")
	b.PutInt32(s.Count)
	b.Comma()
	b.FieldStart("by_file_type")
	b.ArrStart()
	for idx, v := range s.ByFileType {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode storageStatisticsByChat#a5498fe4: field by_file_type element with index %d: %w", idx, err)
		}
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
func (s *StorageStatisticsByChat) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storageStatisticsByChat#a5498fe4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storageStatisticsByChat"); err != nil {
				return fmt.Errorf("unable to decode storageStatisticsByChat#a5498fe4: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode storageStatisticsByChat#a5498fe4: field chat_id: %w", err)
			}
			s.ChatID = value
		case "size":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode storageStatisticsByChat#a5498fe4: field size: %w", err)
			}
			s.Size = value
		case "count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode storageStatisticsByChat#a5498fe4: field count: %w", err)
			}
			s.Count = value
		case "by_file_type":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value StorageStatisticsByFileType
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode storageStatisticsByChat#a5498fe4: field by_file_type: %w", err)
				}
				s.ByFileType = append(s.ByFileType, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode storageStatisticsByChat#a5498fe4: field by_file_type: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *StorageStatisticsByChat) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetSize returns value of Size field.
func (s *StorageStatisticsByChat) GetSize() (value int64) {
	if s == nil {
		return
	}
	return s.Size
}

// GetCount returns value of Count field.
func (s *StorageStatisticsByChat) GetCount() (value int32) {
	if s == nil {
		return
	}
	return s.Count
}

// GetByFileType returns value of ByFileType field.
func (s *StorageStatisticsByChat) GetByFileType() (value []StorageStatisticsByFileType) {
	if s == nil {
		return
	}
	return s.ByFileType
}
