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

// PremiumFeaturePromotionAnimation represents TL type `premiumFeaturePromotionAnimation#899dab1c`.
type PremiumFeaturePromotionAnimation struct {
	// Premium feature
	Feature PremiumFeatureClass
	// Promotion animation for the feature
	Animation Animation
}

// PremiumFeaturePromotionAnimationTypeID is TL type id of PremiumFeaturePromotionAnimation.
const PremiumFeaturePromotionAnimationTypeID = 0x899dab1c

// Ensuring interfaces in compile-time for PremiumFeaturePromotionAnimation.
var (
	_ bin.Encoder     = &PremiumFeaturePromotionAnimation{}
	_ bin.Decoder     = &PremiumFeaturePromotionAnimation{}
	_ bin.BareEncoder = &PremiumFeaturePromotionAnimation{}
	_ bin.BareDecoder = &PremiumFeaturePromotionAnimation{}
)

func (p *PremiumFeaturePromotionAnimation) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Feature == nil) {
		return false
	}
	if !(p.Animation.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumFeaturePromotionAnimation) String() string {
	if p == nil {
		return "PremiumFeaturePromotionAnimation(nil)"
	}
	type Alias PremiumFeaturePromotionAnimation
	return fmt.Sprintf("PremiumFeaturePromotionAnimation%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumFeaturePromotionAnimation) TypeID() uint32 {
	return PremiumFeaturePromotionAnimationTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumFeaturePromotionAnimation) TypeName() string {
	return "premiumFeaturePromotionAnimation"
}

// TypeInfo returns info about TL type.
func (p *PremiumFeaturePromotionAnimation) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumFeaturePromotionAnimation",
		ID:   PremiumFeaturePromotionAnimationTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Feature",
			SchemaName: "feature",
		},
		{
			Name:       "Animation",
			SchemaName: "animation",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PremiumFeaturePromotionAnimation) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumFeaturePromotionAnimation#899dab1c as nil")
	}
	b.PutID(PremiumFeaturePromotionAnimationTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumFeaturePromotionAnimation) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumFeaturePromotionAnimation#899dab1c as nil")
	}
	if p.Feature == nil {
		return fmt.Errorf("unable to encode premiumFeaturePromotionAnimation#899dab1c: field feature is nil")
	}
	if err := p.Feature.Encode(b); err != nil {
		return fmt.Errorf("unable to encode premiumFeaturePromotionAnimation#899dab1c: field feature: %w", err)
	}
	if err := p.Animation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode premiumFeaturePromotionAnimation#899dab1c: field animation: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumFeaturePromotionAnimation) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumFeaturePromotionAnimation#899dab1c to nil")
	}
	if err := b.ConsumeID(PremiumFeaturePromotionAnimationTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumFeaturePromotionAnimation#899dab1c: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumFeaturePromotionAnimation) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumFeaturePromotionAnimation#899dab1c to nil")
	}
	{
		value, err := DecodePremiumFeature(b)
		if err != nil {
			return fmt.Errorf("unable to decode premiumFeaturePromotionAnimation#899dab1c: field feature: %w", err)
		}
		p.Feature = value
	}
	{
		if err := p.Animation.Decode(b); err != nil {
			return fmt.Errorf("unable to decode premiumFeaturePromotionAnimation#899dab1c: field animation: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PremiumFeaturePromotionAnimation) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumFeaturePromotionAnimation#899dab1c as nil")
	}
	b.ObjStart()
	b.PutID("premiumFeaturePromotionAnimation")
	b.Comma()
	b.FieldStart("feature")
	if p.Feature == nil {
		return fmt.Errorf("unable to encode premiumFeaturePromotionAnimation#899dab1c: field feature is nil")
	}
	if err := p.Feature.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode premiumFeaturePromotionAnimation#899dab1c: field feature: %w", err)
	}
	b.Comma()
	b.FieldStart("animation")
	if err := p.Animation.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode premiumFeaturePromotionAnimation#899dab1c: field animation: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PremiumFeaturePromotionAnimation) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumFeaturePromotionAnimation#899dab1c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("premiumFeaturePromotionAnimation"); err != nil {
				return fmt.Errorf("unable to decode premiumFeaturePromotionAnimation#899dab1c: %w", err)
			}
		case "feature":
			value, err := DecodeTDLibJSONPremiumFeature(b)
			if err != nil {
				return fmt.Errorf("unable to decode premiumFeaturePromotionAnimation#899dab1c: field feature: %w", err)
			}
			p.Feature = value
		case "animation":
			if err := p.Animation.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode premiumFeaturePromotionAnimation#899dab1c: field animation: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFeature returns value of Feature field.
func (p *PremiumFeaturePromotionAnimation) GetFeature() (value PremiumFeatureClass) {
	if p == nil {
		return
	}
	return p.Feature
}

// GetAnimation returns value of Animation field.
func (p *PremiumFeaturePromotionAnimation) GetAnimation() (value Animation) {
	if p == nil {
		return
	}
	return p.Animation
}
