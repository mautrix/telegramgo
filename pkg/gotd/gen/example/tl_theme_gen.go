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

// Theme represents TL type `theme#28f1114`.
//
// See https://localhost:80/doc/constructor/theme for reference.
type Theme struct {
	// Name field of Theme.
	Name string
}

// ThemeTypeID is TL type id of Theme.
const ThemeTypeID = 0x28f1114

// Ensuring interfaces in compile-time for Theme.
var (
	_ bin.Encoder     = &Theme{}
	_ bin.Decoder     = &Theme{}
	_ bin.BareEncoder = &Theme{}
	_ bin.BareDecoder = &Theme{}
)

func (t *Theme) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *Theme) String() string {
	if t == nil {
		return "Theme(nil)"
	}
	type Alias Theme
	return fmt.Sprintf("Theme%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Theme) TypeID() uint32 {
	return ThemeTypeID
}

// TypeName returns name of type in TL schema.
func (*Theme) TypeName() string {
	return "theme"
}

// TypeInfo returns info about TL type.
func (t *Theme) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "theme",
		ID:   ThemeTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *Theme) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode theme#28f1114 as nil")
	}
	b.PutID(ThemeTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *Theme) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode theme#28f1114 as nil")
	}
	b.PutString(t.Name)
	return nil
}

// Decode implements bin.Decoder.
func (t *Theme) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode theme#28f1114 to nil")
	}
	if err := b.ConsumeID(ThemeTypeID); err != nil {
		return fmt.Errorf("unable to decode theme#28f1114: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *Theme) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode theme#28f1114 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field name: %w", err)
		}
		t.Name = value
	}
	return nil
}

// GetName returns value of Name field.
func (t *Theme) GetName() (value string) {
	if t == nil {
		return
	}
	return t.Name
}
