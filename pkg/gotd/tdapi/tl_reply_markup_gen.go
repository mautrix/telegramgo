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

// ReplyMarkupRemoveKeyboard represents TL type `replyMarkupRemoveKeyboard#d6cc5171`.
type ReplyMarkupRemoveKeyboard struct {
	// True, if the keyboard is removed only for the mentioned users or the target user of a
	// reply
	IsPersonal bool
}

// ReplyMarkupRemoveKeyboardTypeID is TL type id of ReplyMarkupRemoveKeyboard.
const ReplyMarkupRemoveKeyboardTypeID = 0xd6cc5171

// construct implements constructor of ReplyMarkupClass.
func (r ReplyMarkupRemoveKeyboard) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyMarkupRemoveKeyboard.
var (
	_ bin.Encoder     = &ReplyMarkupRemoveKeyboard{}
	_ bin.Decoder     = &ReplyMarkupRemoveKeyboard{}
	_ bin.BareEncoder = &ReplyMarkupRemoveKeyboard{}
	_ bin.BareDecoder = &ReplyMarkupRemoveKeyboard{}

	_ ReplyMarkupClass = &ReplyMarkupRemoveKeyboard{}
)

func (r *ReplyMarkupRemoveKeyboard) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.IsPersonal == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyMarkupRemoveKeyboard) String() string {
	if r == nil {
		return "ReplyMarkupRemoveKeyboard(nil)"
	}
	type Alias ReplyMarkupRemoveKeyboard
	return fmt.Sprintf("ReplyMarkupRemoveKeyboard%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplyMarkupRemoveKeyboard) TypeID() uint32 {
	return ReplyMarkupRemoveKeyboardTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplyMarkupRemoveKeyboard) TypeName() string {
	return "replyMarkupRemoveKeyboard"
}

// TypeInfo returns info about TL type.
func (r *ReplyMarkupRemoveKeyboard) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replyMarkupRemoveKeyboard",
		ID:   ReplyMarkupRemoveKeyboardTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "IsPersonal",
			SchemaName: "is_personal",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReplyMarkupRemoveKeyboard) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyMarkupRemoveKeyboard#d6cc5171 as nil")
	}
	b.PutID(ReplyMarkupRemoveKeyboardTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReplyMarkupRemoveKeyboard) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyMarkupRemoveKeyboard#d6cc5171 as nil")
	}
	b.PutBool(r.IsPersonal)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReplyMarkupRemoveKeyboard) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyMarkupRemoveKeyboard#d6cc5171 to nil")
	}
	if err := b.ConsumeID(ReplyMarkupRemoveKeyboardTypeID); err != nil {
		return fmt.Errorf("unable to decode replyMarkupRemoveKeyboard#d6cc5171: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReplyMarkupRemoveKeyboard) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyMarkupRemoveKeyboard#d6cc5171 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode replyMarkupRemoveKeyboard#d6cc5171: field is_personal: %w", err)
		}
		r.IsPersonal = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReplyMarkupRemoveKeyboard) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode replyMarkupRemoveKeyboard#d6cc5171 as nil")
	}
	b.ObjStart()
	b.PutID("replyMarkupRemoveKeyboard")
	b.Comma()
	b.FieldStart("is_personal")
	b.PutBool(r.IsPersonal)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReplyMarkupRemoveKeyboard) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode replyMarkupRemoveKeyboard#d6cc5171 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("replyMarkupRemoveKeyboard"); err != nil {
				return fmt.Errorf("unable to decode replyMarkupRemoveKeyboard#d6cc5171: %w", err)
			}
		case "is_personal":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode replyMarkupRemoveKeyboard#d6cc5171: field is_personal: %w", err)
			}
			r.IsPersonal = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetIsPersonal returns value of IsPersonal field.
func (r *ReplyMarkupRemoveKeyboard) GetIsPersonal() (value bool) {
	if r == nil {
		return
	}
	return r.IsPersonal
}

// ReplyMarkupForceReply represents TL type `replyMarkupForceReply#41a6f99f`.
type ReplyMarkupForceReply struct {
	// True, if a forced reply must automatically be shown to the current user. For outgoing
	// messages, specify true to show the forced reply only for the mentioned users and for
	// the target user of a reply
	IsPersonal bool
	// If non-empty, the placeholder to be shown in the input field when the reply is active;
	// 0-64 characters
	InputFieldPlaceholder string
}

// ReplyMarkupForceReplyTypeID is TL type id of ReplyMarkupForceReply.
const ReplyMarkupForceReplyTypeID = 0x41a6f99f

// construct implements constructor of ReplyMarkupClass.
func (r ReplyMarkupForceReply) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyMarkupForceReply.
var (
	_ bin.Encoder     = &ReplyMarkupForceReply{}
	_ bin.Decoder     = &ReplyMarkupForceReply{}
	_ bin.BareEncoder = &ReplyMarkupForceReply{}
	_ bin.BareDecoder = &ReplyMarkupForceReply{}

	_ ReplyMarkupClass = &ReplyMarkupForceReply{}
)

func (r *ReplyMarkupForceReply) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.IsPersonal == false) {
		return false
	}
	if !(r.InputFieldPlaceholder == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyMarkupForceReply) String() string {
	if r == nil {
		return "ReplyMarkupForceReply(nil)"
	}
	type Alias ReplyMarkupForceReply
	return fmt.Sprintf("ReplyMarkupForceReply%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplyMarkupForceReply) TypeID() uint32 {
	return ReplyMarkupForceReplyTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplyMarkupForceReply) TypeName() string {
	return "replyMarkupForceReply"
}

// TypeInfo returns info about TL type.
func (r *ReplyMarkupForceReply) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replyMarkupForceReply",
		ID:   ReplyMarkupForceReplyTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "IsPersonal",
			SchemaName: "is_personal",
		},
		{
			Name:       "InputFieldPlaceholder",
			SchemaName: "input_field_placeholder",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReplyMarkupForceReply) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyMarkupForceReply#41a6f99f as nil")
	}
	b.PutID(ReplyMarkupForceReplyTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReplyMarkupForceReply) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyMarkupForceReply#41a6f99f as nil")
	}
	b.PutBool(r.IsPersonal)
	b.PutString(r.InputFieldPlaceholder)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReplyMarkupForceReply) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyMarkupForceReply#41a6f99f to nil")
	}
	if err := b.ConsumeID(ReplyMarkupForceReplyTypeID); err != nil {
		return fmt.Errorf("unable to decode replyMarkupForceReply#41a6f99f: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReplyMarkupForceReply) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyMarkupForceReply#41a6f99f to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode replyMarkupForceReply#41a6f99f: field is_personal: %w", err)
		}
		r.IsPersonal = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode replyMarkupForceReply#41a6f99f: field input_field_placeholder: %w", err)
		}
		r.InputFieldPlaceholder = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReplyMarkupForceReply) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode replyMarkupForceReply#41a6f99f as nil")
	}
	b.ObjStart()
	b.PutID("replyMarkupForceReply")
	b.Comma()
	b.FieldStart("is_personal")
	b.PutBool(r.IsPersonal)
	b.Comma()
	b.FieldStart("input_field_placeholder")
	b.PutString(r.InputFieldPlaceholder)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReplyMarkupForceReply) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode replyMarkupForceReply#41a6f99f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("replyMarkupForceReply"); err != nil {
				return fmt.Errorf("unable to decode replyMarkupForceReply#41a6f99f: %w", err)
			}
		case "is_personal":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode replyMarkupForceReply#41a6f99f: field is_personal: %w", err)
			}
			r.IsPersonal = value
		case "input_field_placeholder":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode replyMarkupForceReply#41a6f99f: field input_field_placeholder: %w", err)
			}
			r.InputFieldPlaceholder = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetIsPersonal returns value of IsPersonal field.
func (r *ReplyMarkupForceReply) GetIsPersonal() (value bool) {
	if r == nil {
		return
	}
	return r.IsPersonal
}

// GetInputFieldPlaceholder returns value of InputFieldPlaceholder field.
func (r *ReplyMarkupForceReply) GetInputFieldPlaceholder() (value string) {
	if r == nil {
		return
	}
	return r.InputFieldPlaceholder
}

// ReplyMarkupShowKeyboard represents TL type `replyMarkupShowKeyboard#802461d3`.
type ReplyMarkupShowKeyboard struct {
	// A list of rows of bot keyboard buttons
	Rows [][]KeyboardButton
	// True, if the keyboard is expected to always be shown when the ordinary keyboard is
	// hidden
	IsPersistent bool
	// True, if the application needs to resize the keyboard vertically
	ResizeKeyboard bool
	// True, if the application needs to hide the keyboard after use
	OneTime bool
	// True, if the keyboard must automatically be shown to the current user. For outgoing
	// messages, specify true to show the keyboard only for the mentioned users and for the
	// target user of a reply
	IsPersonal bool
	// If non-empty, the placeholder to be shown in the input field when the keyboard is
	// active; 0-64 characters
	InputFieldPlaceholder string
}

// ReplyMarkupShowKeyboardTypeID is TL type id of ReplyMarkupShowKeyboard.
const ReplyMarkupShowKeyboardTypeID = 0x802461d3

// construct implements constructor of ReplyMarkupClass.
func (r ReplyMarkupShowKeyboard) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyMarkupShowKeyboard.
var (
	_ bin.Encoder     = &ReplyMarkupShowKeyboard{}
	_ bin.Decoder     = &ReplyMarkupShowKeyboard{}
	_ bin.BareEncoder = &ReplyMarkupShowKeyboard{}
	_ bin.BareDecoder = &ReplyMarkupShowKeyboard{}

	_ ReplyMarkupClass = &ReplyMarkupShowKeyboard{}
)

func (r *ReplyMarkupShowKeyboard) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Rows == nil) {
		return false
	}
	if !(r.IsPersistent == false) {
		return false
	}
	if !(r.ResizeKeyboard == false) {
		return false
	}
	if !(r.OneTime == false) {
		return false
	}
	if !(r.IsPersonal == false) {
		return false
	}
	if !(r.InputFieldPlaceholder == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyMarkupShowKeyboard) String() string {
	if r == nil {
		return "ReplyMarkupShowKeyboard(nil)"
	}
	type Alias ReplyMarkupShowKeyboard
	return fmt.Sprintf("ReplyMarkupShowKeyboard%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplyMarkupShowKeyboard) TypeID() uint32 {
	return ReplyMarkupShowKeyboardTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplyMarkupShowKeyboard) TypeName() string {
	return "replyMarkupShowKeyboard"
}

// TypeInfo returns info about TL type.
func (r *ReplyMarkupShowKeyboard) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replyMarkupShowKeyboard",
		ID:   ReplyMarkupShowKeyboardTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Rows",
			SchemaName: "rows",
		},
		{
			Name:       "IsPersistent",
			SchemaName: "is_persistent",
		},
		{
			Name:       "ResizeKeyboard",
			SchemaName: "resize_keyboard",
		},
		{
			Name:       "OneTime",
			SchemaName: "one_time",
		},
		{
			Name:       "IsPersonal",
			SchemaName: "is_personal",
		},
		{
			Name:       "InputFieldPlaceholder",
			SchemaName: "input_field_placeholder",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReplyMarkupShowKeyboard) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyMarkupShowKeyboard#802461d3 as nil")
	}
	b.PutID(ReplyMarkupShowKeyboardTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReplyMarkupShowKeyboard) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyMarkupShowKeyboard#802461d3 as nil")
	}
	b.PutInt(len(r.Rows))
	for idx, row := range r.Rows {
		b.PutVectorHeader(len(row))
		for _, v := range row {
			if err := v.EncodeBare(b); err != nil {
				return fmt.Errorf("unable to encode bare replyMarkupShowKeyboard#802461d3: field rows element with index %d: %w", idx, err)
			}
		}
	}
	b.PutBool(r.IsPersistent)
	b.PutBool(r.ResizeKeyboard)
	b.PutBool(r.OneTime)
	b.PutBool(r.IsPersonal)
	b.PutString(r.InputFieldPlaceholder)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReplyMarkupShowKeyboard) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyMarkupShowKeyboard#802461d3 to nil")
	}
	if err := b.ConsumeID(ReplyMarkupShowKeyboardTypeID); err != nil {
		return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReplyMarkupShowKeyboard) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyMarkupShowKeyboard#802461d3 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field rows: %w", err)
		}

		if headerLen > 0 {
			r.Rows = make([][]KeyboardButton, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			innerLen, err := b.VectorHeader()
			if err != nil {
				return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field rows: %w", err)
			}

			var row []KeyboardButton
			if innerLen > 0 {
				row = make([]KeyboardButton, 0, innerLen%bin.PreallocateLimit)
			}
			for innerIndex := 0; innerIndex < innerLen; innerLen++ {
				var value KeyboardButton
				if err := value.DecodeBare(b); err != nil {
					return fmt.Errorf("unable to decode bare replyMarkupShowKeyboard#802461d3: field rows: %w", err)
				}
				row = append(row, value)
			}
			r.Rows = append(r.Rows, row)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field is_persistent: %w", err)
		}
		r.IsPersistent = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field resize_keyboard: %w", err)
		}
		r.ResizeKeyboard = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field one_time: %w", err)
		}
		r.OneTime = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field is_personal: %w", err)
		}
		r.IsPersonal = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field input_field_placeholder: %w", err)
		}
		r.InputFieldPlaceholder = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReplyMarkupShowKeyboard) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode replyMarkupShowKeyboard#802461d3 as nil")
	}
	b.ObjStart()
	b.PutID("replyMarkupShowKeyboard")
	b.Comma()
	b.FieldStart("rows")
	b.ArrStart()
	for idx, row := range r.Rows {
		b.ArrStart()
		for _, v := range row {
			if err := v.EncodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to encode replyMarkupShowKeyboard#802461d3: field rows element with index %d: %w", idx, err)
			}
			b.Comma()
		}
		b.StripComma()
		b.ArrEnd()
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("is_persistent")
	b.PutBool(r.IsPersistent)
	b.Comma()
	b.FieldStart("resize_keyboard")
	b.PutBool(r.ResizeKeyboard)
	b.Comma()
	b.FieldStart("one_time")
	b.PutBool(r.OneTime)
	b.Comma()
	b.FieldStart("is_personal")
	b.PutBool(r.IsPersonal)
	b.Comma()
	b.FieldStart("input_field_placeholder")
	b.PutString(r.InputFieldPlaceholder)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReplyMarkupShowKeyboard) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode replyMarkupShowKeyboard#802461d3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("replyMarkupShowKeyboard"); err != nil {
				return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: %w", err)
			}
		case "rows":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var row []KeyboardButton
				if err := b.Arr(func(b tdjson.Decoder) error {
					var value KeyboardButton
					if err := value.DecodeTDLibJSON(b); err != nil {
						return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field rows: %w", err)
					}
					row = append(row, value)
					return nil
				}); err != nil {
					return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field rows: %w", err)
				}
				r.Rows = append(r.Rows, row)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field rows: %w", err)
			}
		case "is_persistent":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field is_persistent: %w", err)
			}
			r.IsPersistent = value
		case "resize_keyboard":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field resize_keyboard: %w", err)
			}
			r.ResizeKeyboard = value
		case "one_time":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field one_time: %w", err)
			}
			r.OneTime = value
		case "is_personal":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field is_personal: %w", err)
			}
			r.IsPersonal = value
		case "input_field_placeholder":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode replyMarkupShowKeyboard#802461d3: field input_field_placeholder: %w", err)
			}
			r.InputFieldPlaceholder = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetRows returns value of Rows field.
func (r *ReplyMarkupShowKeyboard) GetRows() (value [][]KeyboardButton) {
	if r == nil {
		return
	}
	return r.Rows
}

// GetIsPersistent returns value of IsPersistent field.
func (r *ReplyMarkupShowKeyboard) GetIsPersistent() (value bool) {
	if r == nil {
		return
	}
	return r.IsPersistent
}

// GetResizeKeyboard returns value of ResizeKeyboard field.
func (r *ReplyMarkupShowKeyboard) GetResizeKeyboard() (value bool) {
	if r == nil {
		return
	}
	return r.ResizeKeyboard
}

// GetOneTime returns value of OneTime field.
func (r *ReplyMarkupShowKeyboard) GetOneTime() (value bool) {
	if r == nil {
		return
	}
	return r.OneTime
}

// GetIsPersonal returns value of IsPersonal field.
func (r *ReplyMarkupShowKeyboard) GetIsPersonal() (value bool) {
	if r == nil {
		return
	}
	return r.IsPersonal
}

// GetInputFieldPlaceholder returns value of InputFieldPlaceholder field.
func (r *ReplyMarkupShowKeyboard) GetInputFieldPlaceholder() (value string) {
	if r == nil {
		return
	}
	return r.InputFieldPlaceholder
}

// ReplyMarkupInlineKeyboard represents TL type `replyMarkupInlineKeyboard#92ac0efb`.
type ReplyMarkupInlineKeyboard struct {
	// A list of rows of inline keyboard buttons
	Rows [][]InlineKeyboardButton
}

// ReplyMarkupInlineKeyboardTypeID is TL type id of ReplyMarkupInlineKeyboard.
const ReplyMarkupInlineKeyboardTypeID = 0x92ac0efb

// construct implements constructor of ReplyMarkupClass.
func (r ReplyMarkupInlineKeyboard) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyMarkupInlineKeyboard.
var (
	_ bin.Encoder     = &ReplyMarkupInlineKeyboard{}
	_ bin.Decoder     = &ReplyMarkupInlineKeyboard{}
	_ bin.BareEncoder = &ReplyMarkupInlineKeyboard{}
	_ bin.BareDecoder = &ReplyMarkupInlineKeyboard{}

	_ ReplyMarkupClass = &ReplyMarkupInlineKeyboard{}
)

func (r *ReplyMarkupInlineKeyboard) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Rows == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyMarkupInlineKeyboard) String() string {
	if r == nil {
		return "ReplyMarkupInlineKeyboard(nil)"
	}
	type Alias ReplyMarkupInlineKeyboard
	return fmt.Sprintf("ReplyMarkupInlineKeyboard%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplyMarkupInlineKeyboard) TypeID() uint32 {
	return ReplyMarkupInlineKeyboardTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplyMarkupInlineKeyboard) TypeName() string {
	return "replyMarkupInlineKeyboard"
}

// TypeInfo returns info about TL type.
func (r *ReplyMarkupInlineKeyboard) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replyMarkupInlineKeyboard",
		ID:   ReplyMarkupInlineKeyboardTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Rows",
			SchemaName: "rows",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReplyMarkupInlineKeyboard) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyMarkupInlineKeyboard#92ac0efb as nil")
	}
	b.PutID(ReplyMarkupInlineKeyboardTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReplyMarkupInlineKeyboard) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyMarkupInlineKeyboard#92ac0efb as nil")
	}
	b.PutInt(len(r.Rows))
	for idx, row := range r.Rows {
		b.PutVectorHeader(len(row))
		for _, v := range row {
			if err := v.EncodeBare(b); err != nil {
				return fmt.Errorf("unable to encode bare replyMarkupInlineKeyboard#92ac0efb: field rows element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReplyMarkupInlineKeyboard) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyMarkupInlineKeyboard#92ac0efb to nil")
	}
	if err := b.ConsumeID(ReplyMarkupInlineKeyboardTypeID); err != nil {
		return fmt.Errorf("unable to decode replyMarkupInlineKeyboard#92ac0efb: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReplyMarkupInlineKeyboard) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyMarkupInlineKeyboard#92ac0efb to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode replyMarkupInlineKeyboard#92ac0efb: field rows: %w", err)
		}

		if headerLen > 0 {
			r.Rows = make([][]InlineKeyboardButton, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			innerLen, err := b.VectorHeader()
			if err != nil {
				return fmt.Errorf("unable to decode replyMarkupInlineKeyboard#92ac0efb: field rows: %w", err)
			}

			var row []InlineKeyboardButton
			if innerLen > 0 {
				row = make([]InlineKeyboardButton, 0, innerLen%bin.PreallocateLimit)
			}
			for innerIndex := 0; innerIndex < innerLen; innerLen++ {
				var value InlineKeyboardButton
				if err := value.DecodeBare(b); err != nil {
					return fmt.Errorf("unable to decode bare replyMarkupInlineKeyboard#92ac0efb: field rows: %w", err)
				}
				row = append(row, value)
			}
			r.Rows = append(r.Rows, row)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReplyMarkupInlineKeyboard) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode replyMarkupInlineKeyboard#92ac0efb as nil")
	}
	b.ObjStart()
	b.PutID("replyMarkupInlineKeyboard")
	b.Comma()
	b.FieldStart("rows")
	b.ArrStart()
	for idx, row := range r.Rows {
		b.ArrStart()
		for _, v := range row {
			if err := v.EncodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to encode replyMarkupInlineKeyboard#92ac0efb: field rows element with index %d: %w", idx, err)
			}
			b.Comma()
		}
		b.StripComma()
		b.ArrEnd()
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
func (r *ReplyMarkupInlineKeyboard) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode replyMarkupInlineKeyboard#92ac0efb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("replyMarkupInlineKeyboard"); err != nil {
				return fmt.Errorf("unable to decode replyMarkupInlineKeyboard#92ac0efb: %w", err)
			}
		case "rows":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var row []InlineKeyboardButton
				if err := b.Arr(func(b tdjson.Decoder) error {
					var value InlineKeyboardButton
					if err := value.DecodeTDLibJSON(b); err != nil {
						return fmt.Errorf("unable to decode replyMarkupInlineKeyboard#92ac0efb: field rows: %w", err)
					}
					row = append(row, value)
					return nil
				}); err != nil {
					return fmt.Errorf("unable to decode replyMarkupInlineKeyboard#92ac0efb: field rows: %w", err)
				}
				r.Rows = append(r.Rows, row)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode replyMarkupInlineKeyboard#92ac0efb: field rows: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetRows returns value of Rows field.
func (r *ReplyMarkupInlineKeyboard) GetRows() (value [][]InlineKeyboardButton) {
	if r == nil {
		return
	}
	return r.Rows
}

// ReplyMarkupClassName is schema name of ReplyMarkupClass.
const ReplyMarkupClassName = "ReplyMarkup"

// ReplyMarkupClass represents ReplyMarkup generic type.
//
// Example:
//
//	g, err := tdapi.DecodeReplyMarkup(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.ReplyMarkupRemoveKeyboard: // replyMarkupRemoveKeyboard#d6cc5171
//	case *tdapi.ReplyMarkupForceReply: // replyMarkupForceReply#41a6f99f
//	case *tdapi.ReplyMarkupShowKeyboard: // replyMarkupShowKeyboard#802461d3
//	case *tdapi.ReplyMarkupInlineKeyboard: // replyMarkupInlineKeyboard#92ac0efb
//	default: panic(v)
//	}
type ReplyMarkupClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ReplyMarkupClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeReplyMarkup implements binary de-serialization for ReplyMarkupClass.
func DecodeReplyMarkup(buf *bin.Buffer) (ReplyMarkupClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ReplyMarkupRemoveKeyboardTypeID:
		// Decoding replyMarkupRemoveKeyboard#d6cc5171.
		v := ReplyMarkupRemoveKeyboard{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case ReplyMarkupForceReplyTypeID:
		// Decoding replyMarkupForceReply#41a6f99f.
		v := ReplyMarkupForceReply{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case ReplyMarkupShowKeyboardTypeID:
		// Decoding replyMarkupShowKeyboard#802461d3.
		v := ReplyMarkupShowKeyboard{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case ReplyMarkupInlineKeyboardTypeID:
		// Decoding replyMarkupInlineKeyboard#92ac0efb.
		v := ReplyMarkupInlineKeyboard{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONReplyMarkup implements binary de-serialization for ReplyMarkupClass.
func DecodeTDLibJSONReplyMarkup(buf tdjson.Decoder) (ReplyMarkupClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "replyMarkupRemoveKeyboard":
		// Decoding replyMarkupRemoveKeyboard#d6cc5171.
		v := ReplyMarkupRemoveKeyboard{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case "replyMarkupForceReply":
		// Decoding replyMarkupForceReply#41a6f99f.
		v := ReplyMarkupForceReply{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case "replyMarkupShowKeyboard":
		// Decoding replyMarkupShowKeyboard#802461d3.
		v := ReplyMarkupShowKeyboard{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case "replyMarkupInlineKeyboard":
		// Decoding replyMarkupInlineKeyboard#92ac0efb.
		v := ReplyMarkupInlineKeyboard{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// ReplyMarkup boxes the ReplyMarkupClass providing a helper.
type ReplyMarkupBox struct {
	ReplyMarkup ReplyMarkupClass
}

// Decode implements bin.Decoder for ReplyMarkupBox.
func (b *ReplyMarkupBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ReplyMarkupBox to nil")
	}
	v, err := DecodeReplyMarkup(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ReplyMarkup = v
	return nil
}

// Encode implements bin.Encode for ReplyMarkupBox.
func (b *ReplyMarkupBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode ReplyMarkupClass as nil")
	}
	return b.ReplyMarkup.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for ReplyMarkupBox.
func (b *ReplyMarkupBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode ReplyMarkupBox to nil")
	}
	v, err := DecodeTDLibJSONReplyMarkup(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ReplyMarkup = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for ReplyMarkupBox.
func (b *ReplyMarkupBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode ReplyMarkupClass as nil")
	}
	return b.ReplyMarkup.EncodeTDLibJSON(buf)
}
