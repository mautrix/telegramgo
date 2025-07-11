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

// GroupCall represents TL type `groupCall#9da75b8b`.
type GroupCall struct {
	// Group call identifier
	ID int32
	// Group call title; for video chats only
	Title string
	// Invite link for the group call; for group calls that aren't bound to a chat. For video
	// chats call getVideoChatInviteLink to get the link
	InviteLink string
	// Point in time (Unix timestamp) when the group call is expected to be started by an
	// administrator; 0 if it is already active or was ended; for video chats only
	ScheduledStartDate int32
	// True, if the group call is scheduled and the current user will receive a notification
	// when the group call starts; for video chats only
	EnabledStartNotification bool
	// True, if the call is active
	IsActive bool
	// True, if the call is bound to a chat
	IsVideoChat bool
	// True, if the call is an RTMP stream instead of an ordinary video chat; for video chats
	// only
	IsRtmpStream bool
	// True, if the call is joined
	IsJoined bool
	// True, if user was kicked from the call because of network loss and the call needs to
	// be rejoined
	NeedRejoin bool
	// True, if the user is the owner of the call and can end the call, change volume level
	// of other users, or ban users there; for group calls that aren't bound to a chat
	IsOwned bool
	// True, if the current user can manage the group call; for video chats only
	CanBeManaged bool
	// Number of participants in the group call
	ParticipantCount int32
	// True, if group call participants, which are muted, aren't returned in participant
	// list; for video chats only
	HasHiddenListeners bool
	// True, if all group call participants are loaded
	LoadedAllParticipants bool
	// At most 3 recently speaking users in the group call
	RecentSpeakers []GroupCallRecentSpeaker
	// True, if the current user's video is enabled
	IsMyVideoEnabled bool
	// True, if the current user's video is paused
	IsMyVideoPaused bool
	// True, if the current user can broadcast video or share screen
	CanEnableVideo bool
	// True, if only group call administrators can unmute new participants; for video chats
	// only
	MuteNewParticipants bool
	// True, if the current user can enable or disable mute_new_participants setting; for
	// video chats only
	CanToggleMuteNewParticipants bool
	// Duration of the ongoing group call recording, in seconds; 0 if none. An
	// updateGroupCall update is not triggered when value of this field changes, but the same
	// recording goes on
	RecordDuration int32
	// True, if a video file is being recorded for the call
	IsVideoRecorded bool
	// Call duration, in seconds; for ended calls only
	Duration int32
}

// GroupCallTypeID is TL type id of GroupCall.
const GroupCallTypeID = 0x9da75b8b

// Ensuring interfaces in compile-time for GroupCall.
var (
	_ bin.Encoder     = &GroupCall{}
	_ bin.Decoder     = &GroupCall{}
	_ bin.BareEncoder = &GroupCall{}
	_ bin.BareDecoder = &GroupCall{}
)

func (g *GroupCall) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ID == 0) {
		return false
	}
	if !(g.Title == "") {
		return false
	}
	if !(g.InviteLink == "") {
		return false
	}
	if !(g.ScheduledStartDate == 0) {
		return false
	}
	if !(g.EnabledStartNotification == false) {
		return false
	}
	if !(g.IsActive == false) {
		return false
	}
	if !(g.IsVideoChat == false) {
		return false
	}
	if !(g.IsRtmpStream == false) {
		return false
	}
	if !(g.IsJoined == false) {
		return false
	}
	if !(g.NeedRejoin == false) {
		return false
	}
	if !(g.IsOwned == false) {
		return false
	}
	if !(g.CanBeManaged == false) {
		return false
	}
	if !(g.ParticipantCount == 0) {
		return false
	}
	if !(g.HasHiddenListeners == false) {
		return false
	}
	if !(g.LoadedAllParticipants == false) {
		return false
	}
	if !(g.RecentSpeakers == nil) {
		return false
	}
	if !(g.IsMyVideoEnabled == false) {
		return false
	}
	if !(g.IsMyVideoPaused == false) {
		return false
	}
	if !(g.CanEnableVideo == false) {
		return false
	}
	if !(g.MuteNewParticipants == false) {
		return false
	}
	if !(g.CanToggleMuteNewParticipants == false) {
		return false
	}
	if !(g.RecordDuration == 0) {
		return false
	}
	if !(g.IsVideoRecorded == false) {
		return false
	}
	if !(g.Duration == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GroupCall) String() string {
	if g == nil {
		return "GroupCall(nil)"
	}
	type Alias GroupCall
	return fmt.Sprintf("GroupCall%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GroupCall) TypeID() uint32 {
	return GroupCallTypeID
}

// TypeName returns name of type in TL schema.
func (*GroupCall) TypeName() string {
	return "groupCall"
}

// TypeInfo returns info about TL type.
func (g *GroupCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "groupCall",
		ID:   GroupCallTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "InviteLink",
			SchemaName: "invite_link",
		},
		{
			Name:       "ScheduledStartDate",
			SchemaName: "scheduled_start_date",
		},
		{
			Name:       "EnabledStartNotification",
			SchemaName: "enabled_start_notification",
		},
		{
			Name:       "IsActive",
			SchemaName: "is_active",
		},
		{
			Name:       "IsVideoChat",
			SchemaName: "is_video_chat",
		},
		{
			Name:       "IsRtmpStream",
			SchemaName: "is_rtmp_stream",
		},
		{
			Name:       "IsJoined",
			SchemaName: "is_joined",
		},
		{
			Name:       "NeedRejoin",
			SchemaName: "need_rejoin",
		},
		{
			Name:       "IsOwned",
			SchemaName: "is_owned",
		},
		{
			Name:       "CanBeManaged",
			SchemaName: "can_be_managed",
		},
		{
			Name:       "ParticipantCount",
			SchemaName: "participant_count",
		},
		{
			Name:       "HasHiddenListeners",
			SchemaName: "has_hidden_listeners",
		},
		{
			Name:       "LoadedAllParticipants",
			SchemaName: "loaded_all_participants",
		},
		{
			Name:       "RecentSpeakers",
			SchemaName: "recent_speakers",
		},
		{
			Name:       "IsMyVideoEnabled",
			SchemaName: "is_my_video_enabled",
		},
		{
			Name:       "IsMyVideoPaused",
			SchemaName: "is_my_video_paused",
		},
		{
			Name:       "CanEnableVideo",
			SchemaName: "can_enable_video",
		},
		{
			Name:       "MuteNewParticipants",
			SchemaName: "mute_new_participants",
		},
		{
			Name:       "CanToggleMuteNewParticipants",
			SchemaName: "can_toggle_mute_new_participants",
		},
		{
			Name:       "RecordDuration",
			SchemaName: "record_duration",
		},
		{
			Name:       "IsVideoRecorded",
			SchemaName: "is_video_recorded",
		},
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GroupCall) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCall#9da75b8b as nil")
	}
	b.PutID(GroupCallTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GroupCall) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCall#9da75b8b as nil")
	}
	b.PutInt32(g.ID)
	b.PutString(g.Title)
	b.PutString(g.InviteLink)
	b.PutInt32(g.ScheduledStartDate)
	b.PutBool(g.EnabledStartNotification)
	b.PutBool(g.IsActive)
	b.PutBool(g.IsVideoChat)
	b.PutBool(g.IsRtmpStream)
	b.PutBool(g.IsJoined)
	b.PutBool(g.NeedRejoin)
	b.PutBool(g.IsOwned)
	b.PutBool(g.CanBeManaged)
	b.PutInt32(g.ParticipantCount)
	b.PutBool(g.HasHiddenListeners)
	b.PutBool(g.LoadedAllParticipants)
	b.PutInt(len(g.RecentSpeakers))
	for idx, v := range g.RecentSpeakers {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare groupCall#9da75b8b: field recent_speakers element with index %d: %w", idx, err)
		}
	}
	b.PutBool(g.IsMyVideoEnabled)
	b.PutBool(g.IsMyVideoPaused)
	b.PutBool(g.CanEnableVideo)
	b.PutBool(g.MuteNewParticipants)
	b.PutBool(g.CanToggleMuteNewParticipants)
	b.PutInt32(g.RecordDuration)
	b.PutBool(g.IsVideoRecorded)
	b.PutInt32(g.Duration)
	return nil
}

// Decode implements bin.Decoder.
func (g *GroupCall) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCall#9da75b8b to nil")
	}
	if err := b.ConsumeID(GroupCallTypeID); err != nil {
		return fmt.Errorf("unable to decode groupCall#9da75b8b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GroupCall) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCall#9da75b8b to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field id: %w", err)
		}
		g.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field title: %w", err)
		}
		g.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field invite_link: %w", err)
		}
		g.InviteLink = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field scheduled_start_date: %w", err)
		}
		g.ScheduledStartDate = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field enabled_start_notification: %w", err)
		}
		g.EnabledStartNotification = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_active: %w", err)
		}
		g.IsActive = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_video_chat: %w", err)
		}
		g.IsVideoChat = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_rtmp_stream: %w", err)
		}
		g.IsRtmpStream = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_joined: %w", err)
		}
		g.IsJoined = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field need_rejoin: %w", err)
		}
		g.NeedRejoin = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_owned: %w", err)
		}
		g.IsOwned = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field can_be_managed: %w", err)
		}
		g.CanBeManaged = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field participant_count: %w", err)
		}
		g.ParticipantCount = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field has_hidden_listeners: %w", err)
		}
		g.HasHiddenListeners = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field loaded_all_participants: %w", err)
		}
		g.LoadedAllParticipants = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field recent_speakers: %w", err)
		}

		if headerLen > 0 {
			g.RecentSpeakers = make([]GroupCallRecentSpeaker, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value GroupCallRecentSpeaker
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare groupCall#9da75b8b: field recent_speakers: %w", err)
			}
			g.RecentSpeakers = append(g.RecentSpeakers, value)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_my_video_enabled: %w", err)
		}
		g.IsMyVideoEnabled = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_my_video_paused: %w", err)
		}
		g.IsMyVideoPaused = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field can_enable_video: %w", err)
		}
		g.CanEnableVideo = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field mute_new_participants: %w", err)
		}
		g.MuteNewParticipants = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field can_toggle_mute_new_participants: %w", err)
		}
		g.CanToggleMuteNewParticipants = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field record_duration: %w", err)
		}
		g.RecordDuration = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_video_recorded: %w", err)
		}
		g.IsVideoRecorded = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#9da75b8b: field duration: %w", err)
		}
		g.Duration = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GroupCall) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCall#9da75b8b as nil")
	}
	b.ObjStart()
	b.PutID("groupCall")
	b.Comma()
	b.FieldStart("id")
	b.PutInt32(g.ID)
	b.Comma()
	b.FieldStart("title")
	b.PutString(g.Title)
	b.Comma()
	b.FieldStart("invite_link")
	b.PutString(g.InviteLink)
	b.Comma()
	b.FieldStart("scheduled_start_date")
	b.PutInt32(g.ScheduledStartDate)
	b.Comma()
	b.FieldStart("enabled_start_notification")
	b.PutBool(g.EnabledStartNotification)
	b.Comma()
	b.FieldStart("is_active")
	b.PutBool(g.IsActive)
	b.Comma()
	b.FieldStart("is_video_chat")
	b.PutBool(g.IsVideoChat)
	b.Comma()
	b.FieldStart("is_rtmp_stream")
	b.PutBool(g.IsRtmpStream)
	b.Comma()
	b.FieldStart("is_joined")
	b.PutBool(g.IsJoined)
	b.Comma()
	b.FieldStart("need_rejoin")
	b.PutBool(g.NeedRejoin)
	b.Comma()
	b.FieldStart("is_owned")
	b.PutBool(g.IsOwned)
	b.Comma()
	b.FieldStart("can_be_managed")
	b.PutBool(g.CanBeManaged)
	b.Comma()
	b.FieldStart("participant_count")
	b.PutInt32(g.ParticipantCount)
	b.Comma()
	b.FieldStart("has_hidden_listeners")
	b.PutBool(g.HasHiddenListeners)
	b.Comma()
	b.FieldStart("loaded_all_participants")
	b.PutBool(g.LoadedAllParticipants)
	b.Comma()
	b.FieldStart("recent_speakers")
	b.ArrStart()
	for idx, v := range g.RecentSpeakers {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode groupCall#9da75b8b: field recent_speakers element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("is_my_video_enabled")
	b.PutBool(g.IsMyVideoEnabled)
	b.Comma()
	b.FieldStart("is_my_video_paused")
	b.PutBool(g.IsMyVideoPaused)
	b.Comma()
	b.FieldStart("can_enable_video")
	b.PutBool(g.CanEnableVideo)
	b.Comma()
	b.FieldStart("mute_new_participants")
	b.PutBool(g.MuteNewParticipants)
	b.Comma()
	b.FieldStart("can_toggle_mute_new_participants")
	b.PutBool(g.CanToggleMuteNewParticipants)
	b.Comma()
	b.FieldStart("record_duration")
	b.PutInt32(g.RecordDuration)
	b.Comma()
	b.FieldStart("is_video_recorded")
	b.PutBool(g.IsVideoRecorded)
	b.Comma()
	b.FieldStart("duration")
	b.PutInt32(g.Duration)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GroupCall) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCall#9da75b8b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("groupCall"); err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: %w", err)
			}
		case "id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field id: %w", err)
			}
			g.ID = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field title: %w", err)
			}
			g.Title = value
		case "invite_link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field invite_link: %w", err)
			}
			g.InviteLink = value
		case "scheduled_start_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field scheduled_start_date: %w", err)
			}
			g.ScheduledStartDate = value
		case "enabled_start_notification":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field enabled_start_notification: %w", err)
			}
			g.EnabledStartNotification = value
		case "is_active":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_active: %w", err)
			}
			g.IsActive = value
		case "is_video_chat":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_video_chat: %w", err)
			}
			g.IsVideoChat = value
		case "is_rtmp_stream":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_rtmp_stream: %w", err)
			}
			g.IsRtmpStream = value
		case "is_joined":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_joined: %w", err)
			}
			g.IsJoined = value
		case "need_rejoin":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field need_rejoin: %w", err)
			}
			g.NeedRejoin = value
		case "is_owned":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_owned: %w", err)
			}
			g.IsOwned = value
		case "can_be_managed":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field can_be_managed: %w", err)
			}
			g.CanBeManaged = value
		case "participant_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field participant_count: %w", err)
			}
			g.ParticipantCount = value
		case "has_hidden_listeners":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field has_hidden_listeners: %w", err)
			}
			g.HasHiddenListeners = value
		case "loaded_all_participants":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field loaded_all_participants: %w", err)
			}
			g.LoadedAllParticipants = value
		case "recent_speakers":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value GroupCallRecentSpeaker
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode groupCall#9da75b8b: field recent_speakers: %w", err)
				}
				g.RecentSpeakers = append(g.RecentSpeakers, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field recent_speakers: %w", err)
			}
		case "is_my_video_enabled":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_my_video_enabled: %w", err)
			}
			g.IsMyVideoEnabled = value
		case "is_my_video_paused":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_my_video_paused: %w", err)
			}
			g.IsMyVideoPaused = value
		case "can_enable_video":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field can_enable_video: %w", err)
			}
			g.CanEnableVideo = value
		case "mute_new_participants":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field mute_new_participants: %w", err)
			}
			g.MuteNewParticipants = value
		case "can_toggle_mute_new_participants":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field can_toggle_mute_new_participants: %w", err)
			}
			g.CanToggleMuteNewParticipants = value
		case "record_duration":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field record_duration: %w", err)
			}
			g.RecordDuration = value
		case "is_video_recorded":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field is_video_recorded: %w", err)
			}
			g.IsVideoRecorded = value
		case "duration":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode groupCall#9da75b8b: field duration: %w", err)
			}
			g.Duration = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (g *GroupCall) GetID() (value int32) {
	if g == nil {
		return
	}
	return g.ID
}

// GetTitle returns value of Title field.
func (g *GroupCall) GetTitle() (value string) {
	if g == nil {
		return
	}
	return g.Title
}

// GetInviteLink returns value of InviteLink field.
func (g *GroupCall) GetInviteLink() (value string) {
	if g == nil {
		return
	}
	return g.InviteLink
}

// GetScheduledStartDate returns value of ScheduledStartDate field.
func (g *GroupCall) GetScheduledStartDate() (value int32) {
	if g == nil {
		return
	}
	return g.ScheduledStartDate
}

// GetEnabledStartNotification returns value of EnabledStartNotification field.
func (g *GroupCall) GetEnabledStartNotification() (value bool) {
	if g == nil {
		return
	}
	return g.EnabledStartNotification
}

// GetIsActive returns value of IsActive field.
func (g *GroupCall) GetIsActive() (value bool) {
	if g == nil {
		return
	}
	return g.IsActive
}

// GetIsVideoChat returns value of IsVideoChat field.
func (g *GroupCall) GetIsVideoChat() (value bool) {
	if g == nil {
		return
	}
	return g.IsVideoChat
}

// GetIsRtmpStream returns value of IsRtmpStream field.
func (g *GroupCall) GetIsRtmpStream() (value bool) {
	if g == nil {
		return
	}
	return g.IsRtmpStream
}

// GetIsJoined returns value of IsJoined field.
func (g *GroupCall) GetIsJoined() (value bool) {
	if g == nil {
		return
	}
	return g.IsJoined
}

// GetNeedRejoin returns value of NeedRejoin field.
func (g *GroupCall) GetNeedRejoin() (value bool) {
	if g == nil {
		return
	}
	return g.NeedRejoin
}

// GetIsOwned returns value of IsOwned field.
func (g *GroupCall) GetIsOwned() (value bool) {
	if g == nil {
		return
	}
	return g.IsOwned
}

// GetCanBeManaged returns value of CanBeManaged field.
func (g *GroupCall) GetCanBeManaged() (value bool) {
	if g == nil {
		return
	}
	return g.CanBeManaged
}

// GetParticipantCount returns value of ParticipantCount field.
func (g *GroupCall) GetParticipantCount() (value int32) {
	if g == nil {
		return
	}
	return g.ParticipantCount
}

// GetHasHiddenListeners returns value of HasHiddenListeners field.
func (g *GroupCall) GetHasHiddenListeners() (value bool) {
	if g == nil {
		return
	}
	return g.HasHiddenListeners
}

// GetLoadedAllParticipants returns value of LoadedAllParticipants field.
func (g *GroupCall) GetLoadedAllParticipants() (value bool) {
	if g == nil {
		return
	}
	return g.LoadedAllParticipants
}

// GetRecentSpeakers returns value of RecentSpeakers field.
func (g *GroupCall) GetRecentSpeakers() (value []GroupCallRecentSpeaker) {
	if g == nil {
		return
	}
	return g.RecentSpeakers
}

// GetIsMyVideoEnabled returns value of IsMyVideoEnabled field.
func (g *GroupCall) GetIsMyVideoEnabled() (value bool) {
	if g == nil {
		return
	}
	return g.IsMyVideoEnabled
}

// GetIsMyVideoPaused returns value of IsMyVideoPaused field.
func (g *GroupCall) GetIsMyVideoPaused() (value bool) {
	if g == nil {
		return
	}
	return g.IsMyVideoPaused
}

// GetCanEnableVideo returns value of CanEnableVideo field.
func (g *GroupCall) GetCanEnableVideo() (value bool) {
	if g == nil {
		return
	}
	return g.CanEnableVideo
}

// GetMuteNewParticipants returns value of MuteNewParticipants field.
func (g *GroupCall) GetMuteNewParticipants() (value bool) {
	if g == nil {
		return
	}
	return g.MuteNewParticipants
}

// GetCanToggleMuteNewParticipants returns value of CanToggleMuteNewParticipants field.
func (g *GroupCall) GetCanToggleMuteNewParticipants() (value bool) {
	if g == nil {
		return
	}
	return g.CanToggleMuteNewParticipants
}

// GetRecordDuration returns value of RecordDuration field.
func (g *GroupCall) GetRecordDuration() (value int32) {
	if g == nil {
		return
	}
	return g.RecordDuration
}

// GetIsVideoRecorded returns value of IsVideoRecorded field.
func (g *GroupCall) GetIsVideoRecorded() (value bool) {
	if g == nil {
		return
	}
	return g.IsVideoRecorded
}

// GetDuration returns value of Duration field.
func (g *GroupCall) GetDuration() (value int32) {
	if g == nil {
		return
	}
	return g.Duration
}
