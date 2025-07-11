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

// AvailableReaction represents TL type `availableReaction#c077ec01`.
// Animations associated with a message reaction
//
// See https://core.telegram.org/constructor/availableReaction for reference.
type AvailableReaction struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If not set, the reaction can be added to new messages and enabled in chats.
	Inactive bool
	// Whether this reaction can only be used by Telegram Premium users
	Premium bool
	// Reaction emoji
	Reaction string
	// Reaction description
	Title string
	// Static icon for the reaction
	StaticIcon DocumentClass
	// The animated sticker to show when the user opens the reaction dropdown
	AppearAnimation DocumentClass
	// The animated sticker to show when the user hovers over the reaction
	SelectAnimation DocumentClass
	// The animated sticker to show when the reaction is chosen and activated
	ActivateAnimation DocumentClass
	// The background effect (still an animated sticker) to play under the activate_animation
	// when the reaction is chosen and activated
	EffectAnimation DocumentClass
	// The animation that plays around the button when you press an existing reaction (played
	// together with center_icon).
	//
	// Use SetAroundAnimation and GetAroundAnimation helpers.
	AroundAnimation DocumentClass
	// The animation of the emoji inside the button when you press an existing reaction
	// (played together with around_animation).
	//
	// Use SetCenterIcon and GetCenterIcon helpers.
	CenterIcon DocumentClass
}

// AvailableReactionTypeID is TL type id of AvailableReaction.
const AvailableReactionTypeID = 0xc077ec01

// Ensuring interfaces in compile-time for AvailableReaction.
var (
	_ bin.Encoder     = &AvailableReaction{}
	_ bin.Decoder     = &AvailableReaction{}
	_ bin.BareEncoder = &AvailableReaction{}
	_ bin.BareDecoder = &AvailableReaction{}
)

func (a *AvailableReaction) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.Inactive == false) {
		return false
	}
	if !(a.Premium == false) {
		return false
	}
	if !(a.Reaction == "") {
		return false
	}
	if !(a.Title == "") {
		return false
	}
	if !(a.StaticIcon == nil) {
		return false
	}
	if !(a.AppearAnimation == nil) {
		return false
	}
	if !(a.SelectAnimation == nil) {
		return false
	}
	if !(a.ActivateAnimation == nil) {
		return false
	}
	if !(a.EffectAnimation == nil) {
		return false
	}
	if !(a.AroundAnimation == nil) {
		return false
	}
	if !(a.CenterIcon == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AvailableReaction) String() string {
	if a == nil {
		return "AvailableReaction(nil)"
	}
	type Alias AvailableReaction
	return fmt.Sprintf("AvailableReaction%+v", Alias(*a))
}

// FillFrom fills AvailableReaction from given interface.
func (a *AvailableReaction) FillFrom(from interface {
	GetInactive() (value bool)
	GetPremium() (value bool)
	GetReaction() (value string)
	GetTitle() (value string)
	GetStaticIcon() (value DocumentClass)
	GetAppearAnimation() (value DocumentClass)
	GetSelectAnimation() (value DocumentClass)
	GetActivateAnimation() (value DocumentClass)
	GetEffectAnimation() (value DocumentClass)
	GetAroundAnimation() (value DocumentClass, ok bool)
	GetCenterIcon() (value DocumentClass, ok bool)
}) {
	a.Inactive = from.GetInactive()
	a.Premium = from.GetPremium()
	a.Reaction = from.GetReaction()
	a.Title = from.GetTitle()
	a.StaticIcon = from.GetStaticIcon()
	a.AppearAnimation = from.GetAppearAnimation()
	a.SelectAnimation = from.GetSelectAnimation()
	a.ActivateAnimation = from.GetActivateAnimation()
	a.EffectAnimation = from.GetEffectAnimation()
	if val, ok := from.GetAroundAnimation(); ok {
		a.AroundAnimation = val
	}

	if val, ok := from.GetCenterIcon(); ok {
		a.CenterIcon = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AvailableReaction) TypeID() uint32 {
	return AvailableReactionTypeID
}

// TypeName returns name of type in TL schema.
func (*AvailableReaction) TypeName() string {
	return "availableReaction"
}

// TypeInfo returns info about TL type.
func (a *AvailableReaction) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "availableReaction",
		ID:   AvailableReactionTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Inactive",
			SchemaName: "inactive",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "Premium",
			SchemaName: "premium",
			Null:       !a.Flags.Has(2),
		},
		{
			Name:       "Reaction",
			SchemaName: "reaction",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "StaticIcon",
			SchemaName: "static_icon",
		},
		{
			Name:       "AppearAnimation",
			SchemaName: "appear_animation",
		},
		{
			Name:       "SelectAnimation",
			SchemaName: "select_animation",
		},
		{
			Name:       "ActivateAnimation",
			SchemaName: "activate_animation",
		},
		{
			Name:       "EffectAnimation",
			SchemaName: "effect_animation",
		},
		{
			Name:       "AroundAnimation",
			SchemaName: "around_animation",
			Null:       !a.Flags.Has(1),
		},
		{
			Name:       "CenterIcon",
			SchemaName: "center_icon",
			Null:       !a.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *AvailableReaction) SetFlags() {
	if !(a.Inactive == false) {
		a.Flags.Set(0)
	}
	if !(a.Premium == false) {
		a.Flags.Set(2)
	}
	if !(a.AroundAnimation == nil) {
		a.Flags.Set(1)
	}
	if !(a.CenterIcon == nil) {
		a.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (a *AvailableReaction) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode availableReaction#c077ec01 as nil")
	}
	b.PutID(AvailableReactionTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AvailableReaction) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode availableReaction#c077ec01 as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode availableReaction#c077ec01: field flags: %w", err)
	}
	b.PutString(a.Reaction)
	b.PutString(a.Title)
	if a.StaticIcon == nil {
		return fmt.Errorf("unable to encode availableReaction#c077ec01: field static_icon is nil")
	}
	if err := a.StaticIcon.Encode(b); err != nil {
		return fmt.Errorf("unable to encode availableReaction#c077ec01: field static_icon: %w", err)
	}
	if a.AppearAnimation == nil {
		return fmt.Errorf("unable to encode availableReaction#c077ec01: field appear_animation is nil")
	}
	if err := a.AppearAnimation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode availableReaction#c077ec01: field appear_animation: %w", err)
	}
	if a.SelectAnimation == nil {
		return fmt.Errorf("unable to encode availableReaction#c077ec01: field select_animation is nil")
	}
	if err := a.SelectAnimation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode availableReaction#c077ec01: field select_animation: %w", err)
	}
	if a.ActivateAnimation == nil {
		return fmt.Errorf("unable to encode availableReaction#c077ec01: field activate_animation is nil")
	}
	if err := a.ActivateAnimation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode availableReaction#c077ec01: field activate_animation: %w", err)
	}
	if a.EffectAnimation == nil {
		return fmt.Errorf("unable to encode availableReaction#c077ec01: field effect_animation is nil")
	}
	if err := a.EffectAnimation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode availableReaction#c077ec01: field effect_animation: %w", err)
	}
	if a.Flags.Has(1) {
		if a.AroundAnimation == nil {
			return fmt.Errorf("unable to encode availableReaction#c077ec01: field around_animation is nil")
		}
		if err := a.AroundAnimation.Encode(b); err != nil {
			return fmt.Errorf("unable to encode availableReaction#c077ec01: field around_animation: %w", err)
		}
	}
	if a.Flags.Has(1) {
		if a.CenterIcon == nil {
			return fmt.Errorf("unable to encode availableReaction#c077ec01: field center_icon is nil")
		}
		if err := a.CenterIcon.Encode(b); err != nil {
			return fmt.Errorf("unable to encode availableReaction#c077ec01: field center_icon: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AvailableReaction) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode availableReaction#c077ec01 to nil")
	}
	if err := b.ConsumeID(AvailableReactionTypeID); err != nil {
		return fmt.Errorf("unable to decode availableReaction#c077ec01: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AvailableReaction) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode availableReaction#c077ec01 to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode availableReaction#c077ec01: field flags: %w", err)
		}
	}
	a.Inactive = a.Flags.Has(0)
	a.Premium = a.Flags.Has(2)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode availableReaction#c077ec01: field reaction: %w", err)
		}
		a.Reaction = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode availableReaction#c077ec01: field title: %w", err)
		}
		a.Title = value
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode availableReaction#c077ec01: field static_icon: %w", err)
		}
		a.StaticIcon = value
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode availableReaction#c077ec01: field appear_animation: %w", err)
		}
		a.AppearAnimation = value
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode availableReaction#c077ec01: field select_animation: %w", err)
		}
		a.SelectAnimation = value
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode availableReaction#c077ec01: field activate_animation: %w", err)
		}
		a.ActivateAnimation = value
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode availableReaction#c077ec01: field effect_animation: %w", err)
		}
		a.EffectAnimation = value
	}
	if a.Flags.Has(1) {
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode availableReaction#c077ec01: field around_animation: %w", err)
		}
		a.AroundAnimation = value
	}
	if a.Flags.Has(1) {
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode availableReaction#c077ec01: field center_icon: %w", err)
		}
		a.CenterIcon = value
	}
	return nil
}

// SetInactive sets value of Inactive conditional field.
func (a *AvailableReaction) SetInactive(value bool) {
	if value {
		a.Flags.Set(0)
		a.Inactive = true
	} else {
		a.Flags.Unset(0)
		a.Inactive = false
	}
}

// GetInactive returns value of Inactive conditional field.
func (a *AvailableReaction) GetInactive() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(0)
}

// SetPremium sets value of Premium conditional field.
func (a *AvailableReaction) SetPremium(value bool) {
	if value {
		a.Flags.Set(2)
		a.Premium = true
	} else {
		a.Flags.Unset(2)
		a.Premium = false
	}
}

// GetPremium returns value of Premium conditional field.
func (a *AvailableReaction) GetPremium() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(2)
}

// GetReaction returns value of Reaction field.
func (a *AvailableReaction) GetReaction() (value string) {
	if a == nil {
		return
	}
	return a.Reaction
}

// GetTitle returns value of Title field.
func (a *AvailableReaction) GetTitle() (value string) {
	if a == nil {
		return
	}
	return a.Title
}

// GetStaticIcon returns value of StaticIcon field.
func (a *AvailableReaction) GetStaticIcon() (value DocumentClass) {
	if a == nil {
		return
	}
	return a.StaticIcon
}

// GetAppearAnimation returns value of AppearAnimation field.
func (a *AvailableReaction) GetAppearAnimation() (value DocumentClass) {
	if a == nil {
		return
	}
	return a.AppearAnimation
}

// GetSelectAnimation returns value of SelectAnimation field.
func (a *AvailableReaction) GetSelectAnimation() (value DocumentClass) {
	if a == nil {
		return
	}
	return a.SelectAnimation
}

// GetActivateAnimation returns value of ActivateAnimation field.
func (a *AvailableReaction) GetActivateAnimation() (value DocumentClass) {
	if a == nil {
		return
	}
	return a.ActivateAnimation
}

// GetEffectAnimation returns value of EffectAnimation field.
func (a *AvailableReaction) GetEffectAnimation() (value DocumentClass) {
	if a == nil {
		return
	}
	return a.EffectAnimation
}

// SetAroundAnimation sets value of AroundAnimation conditional field.
func (a *AvailableReaction) SetAroundAnimation(value DocumentClass) {
	a.Flags.Set(1)
	a.AroundAnimation = value
}

// GetAroundAnimation returns value of AroundAnimation conditional field and
// boolean which is true if field was set.
func (a *AvailableReaction) GetAroundAnimation() (value DocumentClass, ok bool) {
	if a == nil {
		return
	}
	if !a.Flags.Has(1) {
		return value, false
	}
	return a.AroundAnimation, true
}

// SetCenterIcon sets value of CenterIcon conditional field.
func (a *AvailableReaction) SetCenterIcon(value DocumentClass) {
	a.Flags.Set(1)
	a.CenterIcon = value
}

// GetCenterIcon returns value of CenterIcon conditional field and
// boolean which is true if field was set.
func (a *AvailableReaction) GetCenterIcon() (value DocumentClass, ok bool) {
	if a == nil {
		return
	}
	if !a.Flags.Has(1) {
		return value, false
	}
	return a.CenterIcon, true
}

// GetStaticIconAsNotEmpty returns mapped value of StaticIcon field.
func (a *AvailableReaction) GetStaticIconAsNotEmpty() (*Document, bool) {
	return a.StaticIcon.AsNotEmpty()
}

// GetAppearAnimationAsNotEmpty returns mapped value of AppearAnimation field.
func (a *AvailableReaction) GetAppearAnimationAsNotEmpty() (*Document, bool) {
	return a.AppearAnimation.AsNotEmpty()
}

// GetSelectAnimationAsNotEmpty returns mapped value of SelectAnimation field.
func (a *AvailableReaction) GetSelectAnimationAsNotEmpty() (*Document, bool) {
	return a.SelectAnimation.AsNotEmpty()
}

// GetActivateAnimationAsNotEmpty returns mapped value of ActivateAnimation field.
func (a *AvailableReaction) GetActivateAnimationAsNotEmpty() (*Document, bool) {
	return a.ActivateAnimation.AsNotEmpty()
}

// GetEffectAnimationAsNotEmpty returns mapped value of EffectAnimation field.
func (a *AvailableReaction) GetEffectAnimationAsNotEmpty() (*Document, bool) {
	return a.EffectAnimation.AsNotEmpty()
}

// GetAroundAnimationAsNotEmpty returns mapped value of AroundAnimation conditional field and
// boolean which is true if field was set.
func (a *AvailableReaction) GetAroundAnimationAsNotEmpty() (*Document, bool) {
	if value, ok := a.GetAroundAnimation(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// GetCenterIconAsNotEmpty returns mapped value of CenterIcon conditional field and
// boolean which is true if field was set.
func (a *AvailableReaction) GetCenterIconAsNotEmpty() (*Document, bool) {
	if value, ok := a.GetCenterIcon(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}
