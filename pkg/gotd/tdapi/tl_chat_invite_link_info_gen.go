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

// ChatInviteLinkInfo represents TL type `chatInviteLinkInfo#c3fe73a`.
type ChatInviteLinkInfo struct {
	// Chat identifier of the invite link; 0 if the user has no access to the chat before
	// joining
	ChatID int64
	// If non-zero, the amount of time for which read access to the chat will remain
	// available, in seconds
	AccessibleFor int32
	// Type of the chat
	Type InviteLinkChatTypeClass
	// Title of the chat
	Title string
	// Chat photo; may be null
	Photo ChatPhotoInfo
	// Identifier of the accent color for chat title and background of chat photo
	AccentColorID int32
	// Contains information about a chat invite link
	Description string
	// Number of members in the chat
	MemberCount int32
	// User identifiers of some chat members that may be known to the current user
	MemberUserIDs []int64
	// Information about subscription plan that must be paid by the user to use the link; may
	// be null if the link doesn't require subscription
	SubscriptionInfo ChatInviteLinkSubscriptionInfo
	// True, if the link only creates join request
	CreatesJoinRequest bool
	// True, if the chat is a public supergroup or channel, i.e. it has a username or it is a
	// location-based supergroup
	IsPublic bool
	// Information about verification status of the chat; may be null if none
	VerificationStatus VerificationStatus
}

// ChatInviteLinkInfoTypeID is TL type id of ChatInviteLinkInfo.
const ChatInviteLinkInfoTypeID = 0xc3fe73a

// Ensuring interfaces in compile-time for ChatInviteLinkInfo.
var (
	_ bin.Encoder     = &ChatInviteLinkInfo{}
	_ bin.Decoder     = &ChatInviteLinkInfo{}
	_ bin.BareEncoder = &ChatInviteLinkInfo{}
	_ bin.BareDecoder = &ChatInviteLinkInfo{}
)

func (c *ChatInviteLinkInfo) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ChatID == 0) {
		return false
	}
	if !(c.AccessibleFor == 0) {
		return false
	}
	if !(c.Type == nil) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.Photo.Zero()) {
		return false
	}
	if !(c.AccentColorID == 0) {
		return false
	}
	if !(c.Description == "") {
		return false
	}
	if !(c.MemberCount == 0) {
		return false
	}
	if !(c.MemberUserIDs == nil) {
		return false
	}
	if !(c.SubscriptionInfo.Zero()) {
		return false
	}
	if !(c.CreatesJoinRequest == false) {
		return false
	}
	if !(c.IsPublic == false) {
		return false
	}
	if !(c.VerificationStatus.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInviteLinkInfo) String() string {
	if c == nil {
		return "ChatInviteLinkInfo(nil)"
	}
	type Alias ChatInviteLinkInfo
	return fmt.Sprintf("ChatInviteLinkInfo%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInviteLinkInfo) TypeID() uint32 {
	return ChatInviteLinkInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInviteLinkInfo) TypeName() string {
	return "chatInviteLinkInfo"
}

// TypeInfo returns info about TL type.
func (c *ChatInviteLinkInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInviteLinkInfo",
		ID:   ChatInviteLinkInfoTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "AccessibleFor",
			SchemaName: "accessible_for",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "AccentColorID",
			SchemaName: "accent_color_id",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
		{
			Name:       "MemberCount",
			SchemaName: "member_count",
		},
		{
			Name:       "MemberUserIDs",
			SchemaName: "member_user_ids",
		},
		{
			Name:       "SubscriptionInfo",
			SchemaName: "subscription_info",
		},
		{
			Name:       "CreatesJoinRequest",
			SchemaName: "creates_join_request",
		},
		{
			Name:       "IsPublic",
			SchemaName: "is_public",
		},
		{
			Name:       "VerificationStatus",
			SchemaName: "verification_status",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatInviteLinkInfo) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinkInfo#c3fe73a as nil")
	}
	b.PutID(ChatInviteLinkInfoTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatInviteLinkInfo) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinkInfo#c3fe73a as nil")
	}
	b.PutInt53(c.ChatID)
	b.PutInt32(c.AccessibleFor)
	if c.Type == nil {
		return fmt.Errorf("unable to encode chatInviteLinkInfo#c3fe73a: field type is nil")
	}
	if err := c.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInviteLinkInfo#c3fe73a: field type: %w", err)
	}
	b.PutString(c.Title)
	if err := c.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInviteLinkInfo#c3fe73a: field photo: %w", err)
	}
	b.PutInt32(c.AccentColorID)
	b.PutString(c.Description)
	b.PutInt32(c.MemberCount)
	b.PutInt(len(c.MemberUserIDs))
	for _, v := range c.MemberUserIDs {
		b.PutInt53(v)
	}
	if err := c.SubscriptionInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInviteLinkInfo#c3fe73a: field subscription_info: %w", err)
	}
	b.PutBool(c.CreatesJoinRequest)
	b.PutBool(c.IsPublic)
	if err := c.VerificationStatus.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInviteLinkInfo#c3fe73a: field verification_status: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInviteLinkInfo) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinkInfo#c3fe73a to nil")
	}
	if err := b.ConsumeID(ChatInviteLinkInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatInviteLinkInfo) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinkInfo#c3fe73a to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field chat_id: %w", err)
		}
		c.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field accessible_for: %w", err)
		}
		c.AccessibleFor = value
	}
	{
		value, err := DecodeInviteLinkChatType(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field type: %w", err)
		}
		c.Type = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field title: %w", err)
		}
		c.Title = value
	}
	{
		if err := c.Photo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field photo: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field accent_color_id: %w", err)
		}
		c.AccentColorID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field description: %w", err)
		}
		c.Description = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field member_count: %w", err)
		}
		c.MemberCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field member_user_ids: %w", err)
		}

		if headerLen > 0 {
			c.MemberUserIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field member_user_ids: %w", err)
			}
			c.MemberUserIDs = append(c.MemberUserIDs, value)
		}
	}
	{
		if err := c.SubscriptionInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field subscription_info: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field creates_join_request: %w", err)
		}
		c.CreatesJoinRequest = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field is_public: %w", err)
		}
		c.IsPublic = value
	}
	{
		if err := c.VerificationStatus.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field verification_status: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatInviteLinkInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinkInfo#c3fe73a as nil")
	}
	b.ObjStart()
	b.PutID("chatInviteLinkInfo")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(c.ChatID)
	b.Comma()
	b.FieldStart("accessible_for")
	b.PutInt32(c.AccessibleFor)
	b.Comma()
	b.FieldStart("type")
	if c.Type == nil {
		return fmt.Errorf("unable to encode chatInviteLinkInfo#c3fe73a: field type is nil")
	}
	if err := c.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatInviteLinkInfo#c3fe73a: field type: %w", err)
	}
	b.Comma()
	b.FieldStart("title")
	b.PutString(c.Title)
	b.Comma()
	b.FieldStart("photo")
	if err := c.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatInviteLinkInfo#c3fe73a: field photo: %w", err)
	}
	b.Comma()
	b.FieldStart("accent_color_id")
	b.PutInt32(c.AccentColorID)
	b.Comma()
	b.FieldStart("description")
	b.PutString(c.Description)
	b.Comma()
	b.FieldStart("member_count")
	b.PutInt32(c.MemberCount)
	b.Comma()
	b.FieldStart("member_user_ids")
	b.ArrStart()
	for _, v := range c.MemberUserIDs {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("subscription_info")
	if err := c.SubscriptionInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatInviteLinkInfo#c3fe73a: field subscription_info: %w", err)
	}
	b.Comma()
	b.FieldStart("creates_join_request")
	b.PutBool(c.CreatesJoinRequest)
	b.Comma()
	b.FieldStart("is_public")
	b.PutBool(c.IsPublic)
	b.Comma()
	b.FieldStart("verification_status")
	if err := c.VerificationStatus.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatInviteLinkInfo#c3fe73a: field verification_status: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatInviteLinkInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinkInfo#c3fe73a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatInviteLinkInfo"); err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field chat_id: %w", err)
			}
			c.ChatID = value
		case "accessible_for":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field accessible_for: %w", err)
			}
			c.AccessibleFor = value
		case "type":
			value, err := DecodeTDLibJSONInviteLinkChatType(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field type: %w", err)
			}
			c.Type = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field title: %w", err)
			}
			c.Title = value
		case "photo":
			if err := c.Photo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field photo: %w", err)
			}
		case "accent_color_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field accent_color_id: %w", err)
			}
			c.AccentColorID = value
		case "description":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field description: %w", err)
			}
			c.Description = value
		case "member_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field member_count: %w", err)
			}
			c.MemberCount = value
		case "member_user_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field member_user_ids: %w", err)
				}
				c.MemberUserIDs = append(c.MemberUserIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field member_user_ids: %w", err)
			}
		case "subscription_info":
			if err := c.SubscriptionInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field subscription_info: %w", err)
			}
		case "creates_join_request":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field creates_join_request: %w", err)
			}
			c.CreatesJoinRequest = value
		case "is_public":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field is_public: %w", err)
			}
			c.IsPublic = value
		case "verification_status":
			if err := c.VerificationStatus.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkInfo#c3fe73a: field verification_status: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (c *ChatInviteLinkInfo) GetChatID() (value int64) {
	if c == nil {
		return
	}
	return c.ChatID
}

// GetAccessibleFor returns value of AccessibleFor field.
func (c *ChatInviteLinkInfo) GetAccessibleFor() (value int32) {
	if c == nil {
		return
	}
	return c.AccessibleFor
}

// GetType returns value of Type field.
func (c *ChatInviteLinkInfo) GetType() (value InviteLinkChatTypeClass) {
	if c == nil {
		return
	}
	return c.Type
}

// GetTitle returns value of Title field.
func (c *ChatInviteLinkInfo) GetTitle() (value string) {
	if c == nil {
		return
	}
	return c.Title
}

// GetPhoto returns value of Photo field.
func (c *ChatInviteLinkInfo) GetPhoto() (value ChatPhotoInfo) {
	if c == nil {
		return
	}
	return c.Photo
}

// GetAccentColorID returns value of AccentColorID field.
func (c *ChatInviteLinkInfo) GetAccentColorID() (value int32) {
	if c == nil {
		return
	}
	return c.AccentColorID
}

// GetDescription returns value of Description field.
func (c *ChatInviteLinkInfo) GetDescription() (value string) {
	if c == nil {
		return
	}
	return c.Description
}

// GetMemberCount returns value of MemberCount field.
func (c *ChatInviteLinkInfo) GetMemberCount() (value int32) {
	if c == nil {
		return
	}
	return c.MemberCount
}

// GetMemberUserIDs returns value of MemberUserIDs field.
func (c *ChatInviteLinkInfo) GetMemberUserIDs() (value []int64) {
	if c == nil {
		return
	}
	return c.MemberUserIDs
}

// GetSubscriptionInfo returns value of SubscriptionInfo field.
func (c *ChatInviteLinkInfo) GetSubscriptionInfo() (value ChatInviteLinkSubscriptionInfo) {
	if c == nil {
		return
	}
	return c.SubscriptionInfo
}

// GetCreatesJoinRequest returns value of CreatesJoinRequest field.
func (c *ChatInviteLinkInfo) GetCreatesJoinRequest() (value bool) {
	if c == nil {
		return
	}
	return c.CreatesJoinRequest
}

// GetIsPublic returns value of IsPublic field.
func (c *ChatInviteLinkInfo) GetIsPublic() (value bool) {
	if c == nil {
		return
	}
	return c.IsPublic
}

// GetVerificationStatus returns value of VerificationStatus field.
func (c *ChatInviteLinkInfo) GetVerificationStatus() (value VerificationStatus) {
	if c == nil {
		return
	}
	return c.VerificationStatus
}
