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

// TestVectorStringObject represents TL type `testVectorStringObject#e5ecc0d`.
//
// See https://localhost:80/doc/constructor/testVectorStringObject for reference.
type TestVectorStringObject struct {
	// Vector of objects
	Value []TestString
}

// TestVectorStringObjectTypeID is TL type id of TestVectorStringObject.
const TestVectorStringObjectTypeID = 0xe5ecc0d

// Ensuring interfaces in compile-time for TestVectorStringObject.
var (
	_ bin.Encoder     = &TestVectorStringObject{}
	_ bin.Decoder     = &TestVectorStringObject{}
	_ bin.BareEncoder = &TestVectorStringObject{}
	_ bin.BareDecoder = &TestVectorStringObject{}
)

func (t *TestVectorStringObject) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Value == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestVectorStringObject) String() string {
	if t == nil {
		return "TestVectorStringObject(nil)"
	}
	type Alias TestVectorStringObject
	return fmt.Sprintf("TestVectorStringObject%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestVectorStringObject) TypeID() uint32 {
	return TestVectorStringObjectTypeID
}

// TypeName returns name of type in TL schema.
func (*TestVectorStringObject) TypeName() string {
	return "testVectorStringObject"
}

// TypeInfo returns info about TL type.
func (t *TestVectorStringObject) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testVectorStringObject",
		ID:   TestVectorStringObjectTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TestVectorStringObject) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testVectorStringObject#e5ecc0d as nil")
	}
	b.PutID(TestVectorStringObjectTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TestVectorStringObject) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testVectorStringObject#e5ecc0d as nil")
	}
	b.PutInt(len(t.Value))
	for idx, v := range t.Value {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare testVectorStringObject#e5ecc0d: field value element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TestVectorStringObject) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testVectorStringObject#e5ecc0d to nil")
	}
	if err := b.ConsumeID(TestVectorStringObjectTypeID); err != nil {
		return fmt.Errorf("unable to decode testVectorStringObject#e5ecc0d: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TestVectorStringObject) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testVectorStringObject#e5ecc0d to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode testVectorStringObject#e5ecc0d: field value: %w", err)
		}

		if headerLen > 0 {
			t.Value = make([]TestString, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TestString
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare testVectorStringObject#e5ecc0d: field value: %w", err)
			}
			t.Value = append(t.Value, value)
		}
	}
	return nil
}

// GetValue returns value of Value field.
func (t *TestVectorStringObject) GetValue() (value []TestString) {
	if t == nil {
		return
	}
	return t.Value
}
