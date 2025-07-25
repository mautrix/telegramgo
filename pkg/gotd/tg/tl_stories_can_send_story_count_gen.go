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

// StoriesCanSendStoryCount represents TL type `stories.canSendStoryCount#c387c04e`.
//
// See https://core.telegram.org/constructor/stories.canSendStoryCount for reference.
type StoriesCanSendStoryCount struct {
	// CountRemains field of StoriesCanSendStoryCount.
	CountRemains int
}

// StoriesCanSendStoryCountTypeID is TL type id of StoriesCanSendStoryCount.
const StoriesCanSendStoryCountTypeID = 0xc387c04e

// Ensuring interfaces in compile-time for StoriesCanSendStoryCount.
var (
	_ bin.Encoder     = &StoriesCanSendStoryCount{}
	_ bin.Decoder     = &StoriesCanSendStoryCount{}
	_ bin.BareEncoder = &StoriesCanSendStoryCount{}
	_ bin.BareDecoder = &StoriesCanSendStoryCount{}
)

func (c *StoriesCanSendStoryCount) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.CountRemains == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *StoriesCanSendStoryCount) String() string {
	if c == nil {
		return "StoriesCanSendStoryCount(nil)"
	}
	type Alias StoriesCanSendStoryCount
	return fmt.Sprintf("StoriesCanSendStoryCount%+v", Alias(*c))
}

// FillFrom fills StoriesCanSendStoryCount from given interface.
func (c *StoriesCanSendStoryCount) FillFrom(from interface {
	GetCountRemains() (value int)
}) {
	c.CountRemains = from.GetCountRemains()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesCanSendStoryCount) TypeID() uint32 {
	return StoriesCanSendStoryCountTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesCanSendStoryCount) TypeName() string {
	return "stories.canSendStoryCount"
}

// TypeInfo returns info about TL type.
func (c *StoriesCanSendStoryCount) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.canSendStoryCount",
		ID:   StoriesCanSendStoryCountTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CountRemains",
			SchemaName: "count_remains",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *StoriesCanSendStoryCount) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode stories.canSendStoryCount#c387c04e as nil")
	}
	b.PutID(StoriesCanSendStoryCountTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *StoriesCanSendStoryCount) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode stories.canSendStoryCount#c387c04e as nil")
	}
	b.PutInt(c.CountRemains)
	return nil
}

// Decode implements bin.Decoder.
func (c *StoriesCanSendStoryCount) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode stories.canSendStoryCount#c387c04e to nil")
	}
	if err := b.ConsumeID(StoriesCanSendStoryCountTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.canSendStoryCount#c387c04e: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *StoriesCanSendStoryCount) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode stories.canSendStoryCount#c387c04e to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stories.canSendStoryCount#c387c04e: field count_remains: %w", err)
		}
		c.CountRemains = value
	}
	return nil
}

// GetCountRemains returns value of CountRemains field.
func (c *StoriesCanSendStoryCount) GetCountRemains() (value int) {
	if c == nil {
		return
	}
	return c.CountRemains
}
