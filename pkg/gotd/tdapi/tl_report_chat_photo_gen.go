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

// ReportChatPhotoRequest represents TL type `reportChatPhoto#d9701288`.
type ReportChatPhotoRequest struct {
	// Chat identifier
	ChatID int64
	// Identifier of the photo to report. Only full photos from chatPhoto can be reported
	FileID int32
	// The reason for reporting the chat photo
	Reason ReportReasonClass
	// Additional report details; 0-1024 characters
	Text string
}

// ReportChatPhotoRequestTypeID is TL type id of ReportChatPhotoRequest.
const ReportChatPhotoRequestTypeID = 0xd9701288

// Ensuring interfaces in compile-time for ReportChatPhotoRequest.
var (
	_ bin.Encoder     = &ReportChatPhotoRequest{}
	_ bin.Decoder     = &ReportChatPhotoRequest{}
	_ bin.BareEncoder = &ReportChatPhotoRequest{}
	_ bin.BareDecoder = &ReportChatPhotoRequest{}
)

func (r *ReportChatPhotoRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatID == 0) {
		return false
	}
	if !(r.FileID == 0) {
		return false
	}
	if !(r.Reason == nil) {
		return false
	}
	if !(r.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportChatPhotoRequest) String() string {
	if r == nil {
		return "ReportChatPhotoRequest(nil)"
	}
	type Alias ReportChatPhotoRequest
	return fmt.Sprintf("ReportChatPhotoRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportChatPhotoRequest) TypeID() uint32 {
	return ReportChatPhotoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportChatPhotoRequest) TypeName() string {
	return "reportChatPhoto"
}

// TypeInfo returns info about TL type.
func (r *ReportChatPhotoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportChatPhoto",
		ID:   ReportChatPhotoRequestTypeID,
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
		{
			Name:       "FileID",
			SchemaName: "file_id",
		},
		{
			Name:       "Reason",
			SchemaName: "reason",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportChatPhotoRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatPhoto#d9701288 as nil")
	}
	b.PutID(ReportChatPhotoRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportChatPhotoRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatPhoto#d9701288 as nil")
	}
	b.PutInt53(r.ChatID)
	b.PutInt32(r.FileID)
	if r.Reason == nil {
		return fmt.Errorf("unable to encode reportChatPhoto#d9701288: field reason is nil")
	}
	if err := r.Reason.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reportChatPhoto#d9701288: field reason: %w", err)
	}
	b.PutString(r.Text)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportChatPhotoRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatPhoto#d9701288 to nil")
	}
	if err := b.ConsumeID(ReportChatPhotoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode reportChatPhoto#d9701288: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportChatPhotoRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatPhoto#d9701288 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode reportChatPhoto#d9701288: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode reportChatPhoto#d9701288: field file_id: %w", err)
		}
		r.FileID = value
	}
	{
		value, err := DecodeReportReason(b)
		if err != nil {
			return fmt.Errorf("unable to decode reportChatPhoto#d9701288: field reason: %w", err)
		}
		r.Reason = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode reportChatPhoto#d9701288: field text: %w", err)
		}
		r.Text = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReportChatPhotoRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatPhoto#d9701288 as nil")
	}
	b.ObjStart()
	b.PutID("reportChatPhoto")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(r.ChatID)
	b.Comma()
	b.FieldStart("file_id")
	b.PutInt32(r.FileID)
	b.Comma()
	b.FieldStart("reason")
	if r.Reason == nil {
		return fmt.Errorf("unable to encode reportChatPhoto#d9701288: field reason is nil")
	}
	if err := r.Reason.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode reportChatPhoto#d9701288: field reason: %w", err)
	}
	b.Comma()
	b.FieldStart("text")
	b.PutString(r.Text)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReportChatPhotoRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatPhoto#d9701288 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reportChatPhoto"); err != nil {
				return fmt.Errorf("unable to decode reportChatPhoto#d9701288: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode reportChatPhoto#d9701288: field chat_id: %w", err)
			}
			r.ChatID = value
		case "file_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode reportChatPhoto#d9701288: field file_id: %w", err)
			}
			r.FileID = value
		case "reason":
			value, err := DecodeTDLibJSONReportReason(b)
			if err != nil {
				return fmt.Errorf("unable to decode reportChatPhoto#d9701288: field reason: %w", err)
			}
			r.Reason = value
		case "text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode reportChatPhoto#d9701288: field text: %w", err)
			}
			r.Text = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (r *ReportChatPhotoRequest) GetChatID() (value int64) {
	if r == nil {
		return
	}
	return r.ChatID
}

// GetFileID returns value of FileID field.
func (r *ReportChatPhotoRequest) GetFileID() (value int32) {
	if r == nil {
		return
	}
	return r.FileID
}

// GetReason returns value of Reason field.
func (r *ReportChatPhotoRequest) GetReason() (value ReportReasonClass) {
	if r == nil {
		return
	}
	return r.Reason
}

// GetText returns value of Text field.
func (r *ReportChatPhotoRequest) GetText() (value string) {
	if r == nil {
		return
	}
	return r.Text
}

// ReportChatPhoto invokes method reportChatPhoto#d9701288 returning error if any.
func (c *Client) ReportChatPhoto(ctx context.Context, request *ReportChatPhotoRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
