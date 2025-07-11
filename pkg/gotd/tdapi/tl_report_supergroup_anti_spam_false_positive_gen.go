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

// ReportSupergroupAntiSpamFalsePositiveRequest represents TL type `reportSupergroupAntiSpamFalsePositive#e13db048`.
type ReportSupergroupAntiSpamFalsePositiveRequest struct {
	// Supergroup identifier
	SupergroupID int64
	// Identifier of the erroneously deleted message from chatEventMessageDeleted
	MessageID int64
}

// ReportSupergroupAntiSpamFalsePositiveRequestTypeID is TL type id of ReportSupergroupAntiSpamFalsePositiveRequest.
const ReportSupergroupAntiSpamFalsePositiveRequestTypeID = 0xe13db048

// Ensuring interfaces in compile-time for ReportSupergroupAntiSpamFalsePositiveRequest.
var (
	_ bin.Encoder     = &ReportSupergroupAntiSpamFalsePositiveRequest{}
	_ bin.Decoder     = &ReportSupergroupAntiSpamFalsePositiveRequest{}
	_ bin.BareEncoder = &ReportSupergroupAntiSpamFalsePositiveRequest{}
	_ bin.BareDecoder = &ReportSupergroupAntiSpamFalsePositiveRequest{}
)

func (r *ReportSupergroupAntiSpamFalsePositiveRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.SupergroupID == 0) {
		return false
	}
	if !(r.MessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportSupergroupAntiSpamFalsePositiveRequest) String() string {
	if r == nil {
		return "ReportSupergroupAntiSpamFalsePositiveRequest(nil)"
	}
	type Alias ReportSupergroupAntiSpamFalsePositiveRequest
	return fmt.Sprintf("ReportSupergroupAntiSpamFalsePositiveRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportSupergroupAntiSpamFalsePositiveRequest) TypeID() uint32 {
	return ReportSupergroupAntiSpamFalsePositiveRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportSupergroupAntiSpamFalsePositiveRequest) TypeName() string {
	return "reportSupergroupAntiSpamFalsePositive"
}

// TypeInfo returns info about TL type.
func (r *ReportSupergroupAntiSpamFalsePositiveRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportSupergroupAntiSpamFalsePositive",
		ID:   ReportSupergroupAntiSpamFalsePositiveRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportSupergroupAntiSpamFalsePositiveRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportSupergroupAntiSpamFalsePositive#e13db048 as nil")
	}
	b.PutID(ReportSupergroupAntiSpamFalsePositiveRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportSupergroupAntiSpamFalsePositiveRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportSupergroupAntiSpamFalsePositive#e13db048 as nil")
	}
	b.PutInt53(r.SupergroupID)
	b.PutInt53(r.MessageID)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportSupergroupAntiSpamFalsePositiveRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportSupergroupAntiSpamFalsePositive#e13db048 to nil")
	}
	if err := b.ConsumeID(ReportSupergroupAntiSpamFalsePositiveRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode reportSupergroupAntiSpamFalsePositive#e13db048: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportSupergroupAntiSpamFalsePositiveRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportSupergroupAntiSpamFalsePositive#e13db048 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode reportSupergroupAntiSpamFalsePositive#e13db048: field supergroup_id: %w", err)
		}
		r.SupergroupID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode reportSupergroupAntiSpamFalsePositive#e13db048: field message_id: %w", err)
		}
		r.MessageID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReportSupergroupAntiSpamFalsePositiveRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reportSupergroupAntiSpamFalsePositive#e13db048 as nil")
	}
	b.ObjStart()
	b.PutID("reportSupergroupAntiSpamFalsePositive")
	b.Comma()
	b.FieldStart("supergroup_id")
	b.PutInt53(r.SupergroupID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(r.MessageID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReportSupergroupAntiSpamFalsePositiveRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reportSupergroupAntiSpamFalsePositive#e13db048 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reportSupergroupAntiSpamFalsePositive"); err != nil {
				return fmt.Errorf("unable to decode reportSupergroupAntiSpamFalsePositive#e13db048: %w", err)
			}
		case "supergroup_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode reportSupergroupAntiSpamFalsePositive#e13db048: field supergroup_id: %w", err)
			}
			r.SupergroupID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode reportSupergroupAntiSpamFalsePositive#e13db048: field message_id: %w", err)
			}
			r.MessageID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (r *ReportSupergroupAntiSpamFalsePositiveRequest) GetSupergroupID() (value int64) {
	if r == nil {
		return
	}
	return r.SupergroupID
}

// GetMessageID returns value of MessageID field.
func (r *ReportSupergroupAntiSpamFalsePositiveRequest) GetMessageID() (value int64) {
	if r == nil {
		return
	}
	return r.MessageID
}

// ReportSupergroupAntiSpamFalsePositive invokes method reportSupergroupAntiSpamFalsePositive#e13db048 returning error if any.
func (c *Client) ReportSupergroupAntiSpamFalsePositive(ctx context.Context, request *ReportSupergroupAntiSpamFalsePositiveRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
