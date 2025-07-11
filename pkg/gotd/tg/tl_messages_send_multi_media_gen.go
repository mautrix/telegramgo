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

// MessagesSendMultiMediaRequest represents TL type `messages.sendMultiMedia#1bf89d74`.
// Send an album or grouped media¹
//
// Links:
//  1. https://core.telegram.org/api/files#albums-grouped-media
//
// See https://core.telegram.org/method/messages.sendMultiMedia for reference.
type MessagesSendMultiMediaRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to send the album silently (no notification triggered)
	Silent bool
	// Send in background?
	Background bool
	// Whether to clear drafts¹
	//
	// Links:
	//  1) https://core.telegram.org/api/drafts
	ClearDraft bool
	// Only for bots, disallows forwarding and saving of the messages, even if the
	// destination chat doesn't have content protection¹ enabled
	//
	// Links:
	//  1) https://telegram.org/blog/protected-content-delete-by-date-and-more
	Noforwards bool
	// Whether to move used stickersets to top, see here for more info on this flag »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/stickers#recent-stickersets
	UpdateStickersetsOrder bool
	// If set, any eventual webpage preview will be shown on top of the message instead of at
	// the bottom.
	InvertMedia bool
	// Bots only: if set, allows sending up to 1000 messages per second, ignoring
	// broadcasting limits¹ for a fee of 0.1 Telegram Stars per message. The relevant Stars
	// will be withdrawn from the bot's balance.
	//
	// Links:
	//  1) https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once
	AllowPaidFloodskip bool
	// The destination chat
	Peer InputPeerClass
	// If set, indicates that the message should be sent in reply to the specified message or
	// story.
	//
	// Use SetReplyTo and GetReplyTo helpers.
	ReplyTo InputReplyToClass
	// The medias to send: note that they must be separately uploaded using messages
	// uploadMedia¹ first, using raw inputMediaUploaded* constructors is not supported.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.uploadMedia
	MultiMedia []InputSingleMedia
	// Scheduled message date for scheduled messages
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
	// Send this message as the specified peer
	//
	// Use SetSendAs and GetSendAs helpers.
	SendAs InputPeerClass
	// Add the message to the specified quick reply shortcut »¹, instead.
	//
	// Links:
	//  1) https://core.telegram.org/api/business#quick-reply-shortcuts
	//
	// Use SetQuickReplyShortcut and GetQuickReplyShortcut helpers.
	QuickReplyShortcut InputQuickReplyShortcutClass
	// Specifies a message effect »¹ to use for the message.
	//
	// Links:
	//  1) https://core.telegram.org/api/effects
	//
	// Use SetEffect and GetEffect helpers.
	Effect int64
	// AllowPaidStars field of MessagesSendMultiMediaRequest.
	//
	// Use SetAllowPaidStars and GetAllowPaidStars helpers.
	AllowPaidStars int64
}

// MessagesSendMultiMediaRequestTypeID is TL type id of MessagesSendMultiMediaRequest.
const MessagesSendMultiMediaRequestTypeID = 0x1bf89d74

// Ensuring interfaces in compile-time for MessagesSendMultiMediaRequest.
var (
	_ bin.Encoder     = &MessagesSendMultiMediaRequest{}
	_ bin.Decoder     = &MessagesSendMultiMediaRequest{}
	_ bin.BareEncoder = &MessagesSendMultiMediaRequest{}
	_ bin.BareDecoder = &MessagesSendMultiMediaRequest{}
)

func (s *MessagesSendMultiMediaRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Silent == false) {
		return false
	}
	if !(s.Background == false) {
		return false
	}
	if !(s.ClearDraft == false) {
		return false
	}
	if !(s.Noforwards == false) {
		return false
	}
	if !(s.UpdateStickersetsOrder == false) {
		return false
	}
	if !(s.InvertMedia == false) {
		return false
	}
	if !(s.AllowPaidFloodskip == false) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.ReplyTo == nil) {
		return false
	}
	if !(s.MultiMedia == nil) {
		return false
	}
	if !(s.ScheduleDate == 0) {
		return false
	}
	if !(s.SendAs == nil) {
		return false
	}
	if !(s.QuickReplyShortcut == nil) {
		return false
	}
	if !(s.Effect == 0) {
		return false
	}
	if !(s.AllowPaidStars == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendMultiMediaRequest) String() string {
	if s == nil {
		return "MessagesSendMultiMediaRequest(nil)"
	}
	type Alias MessagesSendMultiMediaRequest
	return fmt.Sprintf("MessagesSendMultiMediaRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSendMultiMediaRequest from given interface.
func (s *MessagesSendMultiMediaRequest) FillFrom(from interface {
	GetSilent() (value bool)
	GetBackground() (value bool)
	GetClearDraft() (value bool)
	GetNoforwards() (value bool)
	GetUpdateStickersetsOrder() (value bool)
	GetInvertMedia() (value bool)
	GetAllowPaidFloodskip() (value bool)
	GetPeer() (value InputPeerClass)
	GetReplyTo() (value InputReplyToClass, ok bool)
	GetMultiMedia() (value []InputSingleMedia)
	GetScheduleDate() (value int, ok bool)
	GetSendAs() (value InputPeerClass, ok bool)
	GetQuickReplyShortcut() (value InputQuickReplyShortcutClass, ok bool)
	GetEffect() (value int64, ok bool)
	GetAllowPaidStars() (value int64, ok bool)
}) {
	s.Silent = from.GetSilent()
	s.Background = from.GetBackground()
	s.ClearDraft = from.GetClearDraft()
	s.Noforwards = from.GetNoforwards()
	s.UpdateStickersetsOrder = from.GetUpdateStickersetsOrder()
	s.InvertMedia = from.GetInvertMedia()
	s.AllowPaidFloodskip = from.GetAllowPaidFloodskip()
	s.Peer = from.GetPeer()
	if val, ok := from.GetReplyTo(); ok {
		s.ReplyTo = val
	}

	s.MultiMedia = from.GetMultiMedia()
	if val, ok := from.GetScheduleDate(); ok {
		s.ScheduleDate = val
	}

	if val, ok := from.GetSendAs(); ok {
		s.SendAs = val
	}

	if val, ok := from.GetQuickReplyShortcut(); ok {
		s.QuickReplyShortcut = val
	}

	if val, ok := from.GetEffect(); ok {
		s.Effect = val
	}

	if val, ok := from.GetAllowPaidStars(); ok {
		s.AllowPaidStars = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSendMultiMediaRequest) TypeID() uint32 {
	return MessagesSendMultiMediaRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSendMultiMediaRequest) TypeName() string {
	return "messages.sendMultiMedia"
}

// TypeInfo returns info about TL type.
func (s *MessagesSendMultiMediaRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sendMultiMedia",
		ID:   MessagesSendMultiMediaRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !s.Flags.Has(5),
		},
		{
			Name:       "Background",
			SchemaName: "background",
			Null:       !s.Flags.Has(6),
		},
		{
			Name:       "ClearDraft",
			SchemaName: "clear_draft",
			Null:       !s.Flags.Has(7),
		},
		{
			Name:       "Noforwards",
			SchemaName: "noforwards",
			Null:       !s.Flags.Has(14),
		},
		{
			Name:       "UpdateStickersetsOrder",
			SchemaName: "update_stickersets_order",
			Null:       !s.Flags.Has(15),
		},
		{
			Name:       "InvertMedia",
			SchemaName: "invert_media",
			Null:       !s.Flags.Has(16),
		},
		{
			Name:       "AllowPaidFloodskip",
			SchemaName: "allow_paid_floodskip",
			Null:       !s.Flags.Has(19),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ReplyTo",
			SchemaName: "reply_to",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "MultiMedia",
			SchemaName: "multi_media",
		},
		{
			Name:       "ScheduleDate",
			SchemaName: "schedule_date",
			Null:       !s.Flags.Has(10),
		},
		{
			Name:       "SendAs",
			SchemaName: "send_as",
			Null:       !s.Flags.Has(13),
		},
		{
			Name:       "QuickReplyShortcut",
			SchemaName: "quick_reply_shortcut",
			Null:       !s.Flags.Has(17),
		},
		{
			Name:       "Effect",
			SchemaName: "effect",
			Null:       !s.Flags.Has(18),
		},
		{
			Name:       "AllowPaidStars",
			SchemaName: "allow_paid_stars",
			Null:       !s.Flags.Has(21),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSendMultiMediaRequest) SetFlags() {
	if !(s.Silent == false) {
		s.Flags.Set(5)
	}
	if !(s.Background == false) {
		s.Flags.Set(6)
	}
	if !(s.ClearDraft == false) {
		s.Flags.Set(7)
	}
	if !(s.Noforwards == false) {
		s.Flags.Set(14)
	}
	if !(s.UpdateStickersetsOrder == false) {
		s.Flags.Set(15)
	}
	if !(s.InvertMedia == false) {
		s.Flags.Set(16)
	}
	if !(s.AllowPaidFloodskip == false) {
		s.Flags.Set(19)
	}
	if !(s.ReplyTo == nil) {
		s.Flags.Set(0)
	}
	if !(s.ScheduleDate == 0) {
		s.Flags.Set(10)
	}
	if !(s.SendAs == nil) {
		s.Flags.Set(13)
	}
	if !(s.QuickReplyShortcut == nil) {
		s.Flags.Set(17)
	}
	if !(s.Effect == 0) {
		s.Flags.Set(18)
	}
	if !(s.AllowPaidStars == 0) {
		s.Flags.Set(21)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSendMultiMediaRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendMultiMedia#1bf89d74 as nil")
	}
	b.PutID(MessagesSendMultiMediaRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSendMultiMediaRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendMultiMedia#1bf89d74 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMultiMedia#1bf89d74: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendMultiMedia#1bf89d74: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMultiMedia#1bf89d74: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		if s.ReplyTo == nil {
			return fmt.Errorf("unable to encode messages.sendMultiMedia#1bf89d74: field reply_to is nil")
		}
		if err := s.ReplyTo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendMultiMedia#1bf89d74: field reply_to: %w", err)
		}
	}
	b.PutVectorHeader(len(s.MultiMedia))
	for idx, v := range s.MultiMedia {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendMultiMedia#1bf89d74: field multi_media element with index %d: %w", idx, err)
		}
	}
	if s.Flags.Has(10) {
		b.PutInt(s.ScheduleDate)
	}
	if s.Flags.Has(13) {
		if s.SendAs == nil {
			return fmt.Errorf("unable to encode messages.sendMultiMedia#1bf89d74: field send_as is nil")
		}
		if err := s.SendAs.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendMultiMedia#1bf89d74: field send_as: %w", err)
		}
	}
	if s.Flags.Has(17) {
		if s.QuickReplyShortcut == nil {
			return fmt.Errorf("unable to encode messages.sendMultiMedia#1bf89d74: field quick_reply_shortcut is nil")
		}
		if err := s.QuickReplyShortcut.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendMultiMedia#1bf89d74: field quick_reply_shortcut: %w", err)
		}
	}
	if s.Flags.Has(18) {
		b.PutLong(s.Effect)
	}
	if s.Flags.Has(21) {
		b.PutLong(s.AllowPaidStars)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendMultiMediaRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendMultiMedia#1bf89d74 to nil")
	}
	if err := b.ConsumeID(MessagesSendMultiMediaRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendMultiMedia#1bf89d74: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSendMultiMediaRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendMultiMedia#1bf89d74 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#1bf89d74: field flags: %w", err)
		}
	}
	s.Silent = s.Flags.Has(5)
	s.Background = s.Flags.Has(6)
	s.ClearDraft = s.Flags.Has(7)
	s.Noforwards = s.Flags.Has(14)
	s.UpdateStickersetsOrder = s.Flags.Has(15)
	s.InvertMedia = s.Flags.Has(16)
	s.AllowPaidFloodskip = s.Flags.Has(19)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#1bf89d74: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := DecodeInputReplyTo(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#1bf89d74: field reply_to: %w", err)
		}
		s.ReplyTo = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#1bf89d74: field multi_media: %w", err)
		}

		if headerLen > 0 {
			s.MultiMedia = make([]InputSingleMedia, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value InputSingleMedia
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.sendMultiMedia#1bf89d74: field multi_media: %w", err)
			}
			s.MultiMedia = append(s.MultiMedia, value)
		}
	}
	if s.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#1bf89d74: field schedule_date: %w", err)
		}
		s.ScheduleDate = value
	}
	if s.Flags.Has(13) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#1bf89d74: field send_as: %w", err)
		}
		s.SendAs = value
	}
	if s.Flags.Has(17) {
		value, err := DecodeInputQuickReplyShortcut(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#1bf89d74: field quick_reply_shortcut: %w", err)
		}
		s.QuickReplyShortcut = value
	}
	if s.Flags.Has(18) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#1bf89d74: field effect: %w", err)
		}
		s.Effect = value
	}
	if s.Flags.Has(21) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#1bf89d74: field allow_paid_stars: %w", err)
		}
		s.AllowPaidStars = value
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendMultiMediaRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(5)
		s.Silent = true
	} else {
		s.Flags.Unset(5)
		s.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (s *MessagesSendMultiMediaRequest) GetSilent() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(5)
}

// SetBackground sets value of Background conditional field.
func (s *MessagesSendMultiMediaRequest) SetBackground(value bool) {
	if value {
		s.Flags.Set(6)
		s.Background = true
	} else {
		s.Flags.Unset(6)
		s.Background = false
	}
}

// GetBackground returns value of Background conditional field.
func (s *MessagesSendMultiMediaRequest) GetBackground() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(6)
}

// SetClearDraft sets value of ClearDraft conditional field.
func (s *MessagesSendMultiMediaRequest) SetClearDraft(value bool) {
	if value {
		s.Flags.Set(7)
		s.ClearDraft = true
	} else {
		s.Flags.Unset(7)
		s.ClearDraft = false
	}
}

// GetClearDraft returns value of ClearDraft conditional field.
func (s *MessagesSendMultiMediaRequest) GetClearDraft() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(7)
}

// SetNoforwards sets value of Noforwards conditional field.
func (s *MessagesSendMultiMediaRequest) SetNoforwards(value bool) {
	if value {
		s.Flags.Set(14)
		s.Noforwards = true
	} else {
		s.Flags.Unset(14)
		s.Noforwards = false
	}
}

// GetNoforwards returns value of Noforwards conditional field.
func (s *MessagesSendMultiMediaRequest) GetNoforwards() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(14)
}

// SetUpdateStickersetsOrder sets value of UpdateStickersetsOrder conditional field.
func (s *MessagesSendMultiMediaRequest) SetUpdateStickersetsOrder(value bool) {
	if value {
		s.Flags.Set(15)
		s.UpdateStickersetsOrder = true
	} else {
		s.Flags.Unset(15)
		s.UpdateStickersetsOrder = false
	}
}

// GetUpdateStickersetsOrder returns value of UpdateStickersetsOrder conditional field.
func (s *MessagesSendMultiMediaRequest) GetUpdateStickersetsOrder() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(15)
}

// SetInvertMedia sets value of InvertMedia conditional field.
func (s *MessagesSendMultiMediaRequest) SetInvertMedia(value bool) {
	if value {
		s.Flags.Set(16)
		s.InvertMedia = true
	} else {
		s.Flags.Unset(16)
		s.InvertMedia = false
	}
}

// GetInvertMedia returns value of InvertMedia conditional field.
func (s *MessagesSendMultiMediaRequest) GetInvertMedia() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(16)
}

// SetAllowPaidFloodskip sets value of AllowPaidFloodskip conditional field.
func (s *MessagesSendMultiMediaRequest) SetAllowPaidFloodskip(value bool) {
	if value {
		s.Flags.Set(19)
		s.AllowPaidFloodskip = true
	} else {
		s.Flags.Unset(19)
		s.AllowPaidFloodskip = false
	}
}

// GetAllowPaidFloodskip returns value of AllowPaidFloodskip conditional field.
func (s *MessagesSendMultiMediaRequest) GetAllowPaidFloodskip() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(19)
}

// GetPeer returns value of Peer field.
func (s *MessagesSendMultiMediaRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// SetReplyTo sets value of ReplyTo conditional field.
func (s *MessagesSendMultiMediaRequest) SetReplyTo(value InputReplyToClass) {
	s.Flags.Set(0)
	s.ReplyTo = value
}

// GetReplyTo returns value of ReplyTo conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMultiMediaRequest) GetReplyTo() (value InputReplyToClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ReplyTo, true
}

// GetMultiMedia returns value of MultiMedia field.
func (s *MessagesSendMultiMediaRequest) GetMultiMedia() (value []InputSingleMedia) {
	if s == nil {
		return
	}
	return s.MultiMedia
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (s *MessagesSendMultiMediaRequest) SetScheduleDate(value int) {
	s.Flags.Set(10)
	s.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMultiMediaRequest) GetScheduleDate() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(10) {
		return value, false
	}
	return s.ScheduleDate, true
}

// SetSendAs sets value of SendAs conditional field.
func (s *MessagesSendMultiMediaRequest) SetSendAs(value InputPeerClass) {
	s.Flags.Set(13)
	s.SendAs = value
}

// GetSendAs returns value of SendAs conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMultiMediaRequest) GetSendAs() (value InputPeerClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(13) {
		return value, false
	}
	return s.SendAs, true
}

// SetQuickReplyShortcut sets value of QuickReplyShortcut conditional field.
func (s *MessagesSendMultiMediaRequest) SetQuickReplyShortcut(value InputQuickReplyShortcutClass) {
	s.Flags.Set(17)
	s.QuickReplyShortcut = value
}

// GetQuickReplyShortcut returns value of QuickReplyShortcut conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMultiMediaRequest) GetQuickReplyShortcut() (value InputQuickReplyShortcutClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(17) {
		return value, false
	}
	return s.QuickReplyShortcut, true
}

// SetEffect sets value of Effect conditional field.
func (s *MessagesSendMultiMediaRequest) SetEffect(value int64) {
	s.Flags.Set(18)
	s.Effect = value
}

// GetEffect returns value of Effect conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMultiMediaRequest) GetEffect() (value int64, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(18) {
		return value, false
	}
	return s.Effect, true
}

// SetAllowPaidStars sets value of AllowPaidStars conditional field.
func (s *MessagesSendMultiMediaRequest) SetAllowPaidStars(value int64) {
	s.Flags.Set(21)
	s.AllowPaidStars = value
}

// GetAllowPaidStars returns value of AllowPaidStars conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMultiMediaRequest) GetAllowPaidStars() (value int64, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(21) {
		return value, false
	}
	return s.AllowPaidStars, true
}

// MessagesSendMultiMedia invokes method messages.sendMultiMedia#1bf89d74 returning error if any.
// Send an album or grouped media¹
//
// Links:
//  1. https://core.telegram.org/api/files#albums-grouped-media
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 CHAT_FORWARDS_RESTRICTED: You can't forward messages from a protected chat.
//	403 CHAT_SEND_MEDIA_FORBIDDEN: You can't send media in this chat.
//	403 CHAT_SEND_PHOTOS_FORBIDDEN: You can't send photos in this chat.
//	403 CHAT_SEND_VIDEOS_FORBIDDEN: You can't send videos in this chat.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	400 ENTITY_BOUNDS_INVALID: A specified entity offset or length is invalid, see here » for info on how to properly compute the entity offset/length.
//	400 FILE_REFERENCE_%d_EXPIRED: The file reference of the media file at index %d in the passed media array expired, it must be refreshed.
//	400 FILE_REFERENCE_%d_INVALID: The file reference of the media file at index %d in the passed media array is invalid.
//	400 MEDIA_CAPTION_TOO_LONG: The caption is too long.
//	400 MEDIA_EMPTY: The provided media object is invalid.
//	400 MEDIA_INVALID: Media invalid.
//	400 MULTI_MEDIA_TOO_LONG: Too many media files for album.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 QUICK_REPLIES_TOO_MUCH: A maximum of appConfig.quick_replies_limit shortcuts may be created, the limit was reached.
//	500 RANDOM_ID_DUPLICATE: You provided a random ID that was already used.
//	400 RANDOM_ID_EMPTY: Random ID empty.
//	400 REPLY_MESSAGES_TOO_MUCH: Each shortcut can contain a maximum of appConfig.quick_reply_messages_limit messages, the limit was reached.
//	400 SCHEDULE_DATE_TOO_LATE: You can't schedule a message this far in the future.
//	400 SCHEDULE_TOO_MUCH: There are too many scheduled messages.
//	400 SEND_AS_PEER_INVALID: You can't send messages as the specified peer.
//	420 SLOWMODE_WAIT_%d: Slowmode is enabled in this chat: wait %d seconds before sending another message to this chat.
//	400 TOPIC_CLOSED: This topic was closed, you can't send messages to it anymore.
//	400 TOPIC_DELETED: The specified topic was deleted.
//	400 USER_BANNED_IN_CHANNEL: You're banned from sending messages in supergroups/channels.
//
// See https://core.telegram.org/method/messages.sendMultiMedia for reference.
// Can be used by bots.
func (c *Client) MessagesSendMultiMedia(ctx context.Context, request *MessagesSendMultiMediaRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
