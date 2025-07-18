// Code generated by gotdgen, DO NOT EDIT.

package td

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

// TextEntities represents TL type `textEntities#cf89c258`.
//
// See https://localhost:80/doc/constructor/textEntities for reference.
type TextEntities struct {
	// List of text entities
	Entities []TextEntity
}

// TextEntitiesTypeID is TL type id of TextEntities.
const TextEntitiesTypeID = 0xcf89c258

// Ensuring interfaces in compile-time for TextEntities.
var (
	_ bin.Encoder     = &TextEntities{}
	_ bin.Decoder     = &TextEntities{}
	_ bin.BareEncoder = &TextEntities{}
	_ bin.BareDecoder = &TextEntities{}
)

func (t *TextEntities) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Entities == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TextEntities) String() string {
	if t == nil {
		return "TextEntities(nil)"
	}
	type Alias TextEntities
	return fmt.Sprintf("TextEntities%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TextEntities) TypeID() uint32 {
	return TextEntitiesTypeID
}

// TypeName returns name of type in TL schema.
func (*TextEntities) TypeName() string {
	return "textEntities"
}

// TypeInfo returns info about TL type.
func (t *TextEntities) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "textEntities",
		ID:   TextEntitiesTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Entities",
			SchemaName: "entities",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TextEntities) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textEntities#cf89c258 as nil")
	}
	b.PutID(TextEntitiesTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TextEntities) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textEntities#cf89c258 as nil")
	}
	b.PutInt(len(t.Entities))
	for idx, v := range t.Entities {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare textEntities#cf89c258: field entities element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TextEntities) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textEntities#cf89c258 to nil")
	}
	if err := b.ConsumeID(TextEntitiesTypeID); err != nil {
		return fmt.Errorf("unable to decode textEntities#cf89c258: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TextEntities) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textEntities#cf89c258 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode textEntities#cf89c258: field entities: %w", err)
		}

		if headerLen > 0 {
			t.Entities = make([]TextEntity, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TextEntity
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare textEntities#cf89c258: field entities: %w", err)
			}
			t.Entities = append(t.Entities, value)
		}
	}
	return nil
}

// GetEntities returns value of Entities field.
func (t *TextEntities) GetEntities() (value []TextEntity) {
	if t == nil {
		return
	}
	return t.Entities
}
