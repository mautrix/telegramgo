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

// ChannelsCreateChannelRequest represents TL type `channels.createChannel#91006707`.
// Create a supergroup/channel¹.
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.createChannel for reference.
type ChannelsCreateChannelRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to create a channel¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Broadcast bool
	// Whether to create a supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Megagroup bool
	// Whether the supergroup is being created to import messages from a foreign chat service
	// using messages.initHistoryImport¹
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.initHistoryImport
	ForImport bool
	// Whether to create a forum¹
	//
	// Links:
	//  1) https://core.telegram.org/api/forum
	Forum bool
	// Channel title
	Title string
	// Channel description
	About string
	// Geogroup location, see here »¹ for more info on geogroups.
	//
	// Links:
	//  1) https://core.telegram.org/api/nearby
	//
	// Use SetGeoPoint and GetGeoPoint helpers.
	GeoPoint InputGeoPointClass
	// Geogroup address, see here »¹ for more info on geogroups.
	//
	// Links:
	//  1) https://core.telegram.org/api/nearby
	//
	// Use SetAddress and GetAddress helpers.
	Address string
	// Time-to-live of all messages that will be sent in the supergroup: once message
	// date+message.ttl_period === time(), the message will be deleted on the server, and
	// must be deleted locally as well. You can use messages.setDefaultHistoryTTL¹ to edit
	// this value later.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.setDefaultHistoryTTL
	//
	// Use SetTTLPeriod and GetTTLPeriod helpers.
	TTLPeriod int
}

// ChannelsCreateChannelRequestTypeID is TL type id of ChannelsCreateChannelRequest.
const ChannelsCreateChannelRequestTypeID = 0x91006707

// Ensuring interfaces in compile-time for ChannelsCreateChannelRequest.
var (
	_ bin.Encoder     = &ChannelsCreateChannelRequest{}
	_ bin.Decoder     = &ChannelsCreateChannelRequest{}
	_ bin.BareEncoder = &ChannelsCreateChannelRequest{}
	_ bin.BareDecoder = &ChannelsCreateChannelRequest{}
)

func (c *ChannelsCreateChannelRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Broadcast == false) {
		return false
	}
	if !(c.Megagroup == false) {
		return false
	}
	if !(c.ForImport == false) {
		return false
	}
	if !(c.Forum == false) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.About == "") {
		return false
	}
	if !(c.GeoPoint == nil) {
		return false
	}
	if !(c.Address == "") {
		return false
	}
	if !(c.TTLPeriod == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelsCreateChannelRequest) String() string {
	if c == nil {
		return "ChannelsCreateChannelRequest(nil)"
	}
	type Alias ChannelsCreateChannelRequest
	return fmt.Sprintf("ChannelsCreateChannelRequest%+v", Alias(*c))
}

// FillFrom fills ChannelsCreateChannelRequest from given interface.
func (c *ChannelsCreateChannelRequest) FillFrom(from interface {
	GetBroadcast() (value bool)
	GetMegagroup() (value bool)
	GetForImport() (value bool)
	GetForum() (value bool)
	GetTitle() (value string)
	GetAbout() (value string)
	GetGeoPoint() (value InputGeoPointClass, ok bool)
	GetAddress() (value string, ok bool)
	GetTTLPeriod() (value int, ok bool)
}) {
	c.Broadcast = from.GetBroadcast()
	c.Megagroup = from.GetMegagroup()
	c.ForImport = from.GetForImport()
	c.Forum = from.GetForum()
	c.Title = from.GetTitle()
	c.About = from.GetAbout()
	if val, ok := from.GetGeoPoint(); ok {
		c.GeoPoint = val
	}

	if val, ok := from.GetAddress(); ok {
		c.Address = val
	}

	if val, ok := from.GetTTLPeriod(); ok {
		c.TTLPeriod = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsCreateChannelRequest) TypeID() uint32 {
	return ChannelsCreateChannelRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsCreateChannelRequest) TypeName() string {
	return "channels.createChannel"
}

// TypeInfo returns info about TL type.
func (c *ChannelsCreateChannelRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.createChannel",
		ID:   ChannelsCreateChannelRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Broadcast",
			SchemaName: "broadcast",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "Megagroup",
			SchemaName: "megagroup",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "ForImport",
			SchemaName: "for_import",
			Null:       !c.Flags.Has(3),
		},
		{
			Name:       "Forum",
			SchemaName: "forum",
			Null:       !c.Flags.Has(5),
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "About",
			SchemaName: "about",
		},
		{
			Name:       "GeoPoint",
			SchemaName: "geo_point",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "Address",
			SchemaName: "address",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "TTLPeriod",
			SchemaName: "ttl_period",
			Null:       !c.Flags.Has(4),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *ChannelsCreateChannelRequest) SetFlags() {
	if !(c.Broadcast == false) {
		c.Flags.Set(0)
	}
	if !(c.Megagroup == false) {
		c.Flags.Set(1)
	}
	if !(c.ForImport == false) {
		c.Flags.Set(3)
	}
	if !(c.Forum == false) {
		c.Flags.Set(5)
	}
	if !(c.GeoPoint == nil) {
		c.Flags.Set(2)
	}
	if !(c.Address == "") {
		c.Flags.Set(2)
	}
	if !(c.TTLPeriod == 0) {
		c.Flags.Set(4)
	}
}

// Encode implements bin.Encoder.
func (c *ChannelsCreateChannelRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.createChannel#91006707 as nil")
	}
	b.PutID(ChannelsCreateChannelRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelsCreateChannelRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.createChannel#91006707 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.createChannel#91006707: field flags: %w", err)
	}
	b.PutString(c.Title)
	b.PutString(c.About)
	if c.Flags.Has(2) {
		if c.GeoPoint == nil {
			return fmt.Errorf("unable to encode channels.createChannel#91006707: field geo_point is nil")
		}
		if err := c.GeoPoint.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.createChannel#91006707: field geo_point: %w", err)
		}
	}
	if c.Flags.Has(2) {
		b.PutString(c.Address)
	}
	if c.Flags.Has(4) {
		b.PutInt(c.TTLPeriod)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelsCreateChannelRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.createChannel#91006707 to nil")
	}
	if err := b.ConsumeID(ChannelsCreateChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.createChannel#91006707: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelsCreateChannelRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.createChannel#91006707 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#91006707: field flags: %w", err)
		}
	}
	c.Broadcast = c.Flags.Has(0)
	c.Megagroup = c.Flags.Has(1)
	c.ForImport = c.Flags.Has(3)
	c.Forum = c.Flags.Has(5)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#91006707: field title: %w", err)
		}
		c.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#91006707: field about: %w", err)
		}
		c.About = value
	}
	if c.Flags.Has(2) {
		value, err := DecodeInputGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#91006707: field geo_point: %w", err)
		}
		c.GeoPoint = value
	}
	if c.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#91006707: field address: %w", err)
		}
		c.Address = value
	}
	if c.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#91006707: field ttl_period: %w", err)
		}
		c.TTLPeriod = value
	}
	return nil
}

// SetBroadcast sets value of Broadcast conditional field.
func (c *ChannelsCreateChannelRequest) SetBroadcast(value bool) {
	if value {
		c.Flags.Set(0)
		c.Broadcast = true
	} else {
		c.Flags.Unset(0)
		c.Broadcast = false
	}
}

// GetBroadcast returns value of Broadcast conditional field.
func (c *ChannelsCreateChannelRequest) GetBroadcast() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(0)
}

// SetMegagroup sets value of Megagroup conditional field.
func (c *ChannelsCreateChannelRequest) SetMegagroup(value bool) {
	if value {
		c.Flags.Set(1)
		c.Megagroup = true
	} else {
		c.Flags.Unset(1)
		c.Megagroup = false
	}
}

// GetMegagroup returns value of Megagroup conditional field.
func (c *ChannelsCreateChannelRequest) GetMegagroup() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(1)
}

// SetForImport sets value of ForImport conditional field.
func (c *ChannelsCreateChannelRequest) SetForImport(value bool) {
	if value {
		c.Flags.Set(3)
		c.ForImport = true
	} else {
		c.Flags.Unset(3)
		c.ForImport = false
	}
}

// GetForImport returns value of ForImport conditional field.
func (c *ChannelsCreateChannelRequest) GetForImport() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(3)
}

// SetForum sets value of Forum conditional field.
func (c *ChannelsCreateChannelRequest) SetForum(value bool) {
	if value {
		c.Flags.Set(5)
		c.Forum = true
	} else {
		c.Flags.Unset(5)
		c.Forum = false
	}
}

// GetForum returns value of Forum conditional field.
func (c *ChannelsCreateChannelRequest) GetForum() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(5)
}

// GetTitle returns value of Title field.
func (c *ChannelsCreateChannelRequest) GetTitle() (value string) {
	if c == nil {
		return
	}
	return c.Title
}

// GetAbout returns value of About field.
func (c *ChannelsCreateChannelRequest) GetAbout() (value string) {
	if c == nil {
		return
	}
	return c.About
}

// SetGeoPoint sets value of GeoPoint conditional field.
func (c *ChannelsCreateChannelRequest) SetGeoPoint(value InputGeoPointClass) {
	c.Flags.Set(2)
	c.GeoPoint = value
}

// GetGeoPoint returns value of GeoPoint conditional field and
// boolean which is true if field was set.
func (c *ChannelsCreateChannelRequest) GetGeoPoint() (value InputGeoPointClass, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.GeoPoint, true
}

// SetAddress sets value of Address conditional field.
func (c *ChannelsCreateChannelRequest) SetAddress(value string) {
	c.Flags.Set(2)
	c.Address = value
}

// GetAddress returns value of Address conditional field and
// boolean which is true if field was set.
func (c *ChannelsCreateChannelRequest) GetAddress() (value string, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.Address, true
}

// SetTTLPeriod sets value of TTLPeriod conditional field.
func (c *ChannelsCreateChannelRequest) SetTTLPeriod(value int) {
	c.Flags.Set(4)
	c.TTLPeriod = value
}

// GetTTLPeriod returns value of TTLPeriod conditional field and
// boolean which is true if field was set.
func (c *ChannelsCreateChannelRequest) GetTTLPeriod() (value int, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(4) {
		return value, false
	}
	return c.TTLPeriod, true
}

// GetGeoPointAsNotEmpty returns mapped value of GeoPoint conditional field and
// boolean which is true if field was set.
func (c *ChannelsCreateChannelRequest) GetGeoPointAsNotEmpty() (*InputGeoPoint, bool) {
	if value, ok := c.GetGeoPoint(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// ChannelsCreateChannel invokes method channels.createChannel#91006707 returning error if any.
// Create a supergroup/channel¹.
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// Possible errors:
//
//	400 ADDRESS_INVALID: The specified geopoint address is invalid.
//	400 CHANNELS_ADMIN_LOCATED_TOO_MUCH: The user has reached the limit of public geogroups.
//	400 CHANNELS_TOO_MUCH: You have joined too many channels/supergroups.
//	400 CHAT_ABOUT_TOO_LONG: Chat about too long.
//	500 CHAT_INVALID: Invalid chat.
//	400 CHAT_TITLE_EMPTY: No chat title provided.
//	400 TTL_PERIOD_INVALID: The specified TTL period is invalid.
//	406 USER_RESTRICTED: You're spamreported, you can't create channels or chats.
//
// See https://core.telegram.org/method/channels.createChannel for reference.
func (c *Client) ChannelsCreateChannel(ctx context.Context, request *ChannelsCreateChannelRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
