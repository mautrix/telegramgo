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

// InvokeWithReCaptchaRequest represents TL type `invokeWithReCaptcha#adbb0f94`.
//
// See https://core.telegram.org/constructor/invokeWithReCaptcha for reference.
type InvokeWithReCaptchaRequest struct {
	// Token field of InvokeWithReCaptchaRequest.
	Token string
	// Query field of InvokeWithReCaptchaRequest.
	Query bin.Object
}

// InvokeWithReCaptchaRequestTypeID is TL type id of InvokeWithReCaptchaRequest.
const InvokeWithReCaptchaRequestTypeID = 0xadbb0f94

// Ensuring interfaces in compile-time for InvokeWithReCaptchaRequest.
var (
	_ bin.Encoder     = &InvokeWithReCaptchaRequest{}
	_ bin.Decoder     = &InvokeWithReCaptchaRequest{}
	_ bin.BareEncoder = &InvokeWithReCaptchaRequest{}
	_ bin.BareDecoder = &InvokeWithReCaptchaRequest{}
)

func (i *InvokeWithReCaptchaRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Token == "") {
		return false
	}
	if !(i.Query == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InvokeWithReCaptchaRequest) String() string {
	if i == nil {
		return "InvokeWithReCaptchaRequest(nil)"
	}
	type Alias InvokeWithReCaptchaRequest
	return fmt.Sprintf("InvokeWithReCaptchaRequest%+v", Alias(*i))
}

// FillFrom fills InvokeWithReCaptchaRequest from given interface.
func (i *InvokeWithReCaptchaRequest) FillFrom(from interface {
	GetToken() (value string)
	GetQuery() (value bin.Object)
}) {
	i.Token = from.GetToken()
	i.Query = from.GetQuery()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InvokeWithReCaptchaRequest) TypeID() uint32 {
	return InvokeWithReCaptchaRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*InvokeWithReCaptchaRequest) TypeName() string {
	return "invokeWithReCaptcha"
}

// TypeInfo returns info about TL type.
func (i *InvokeWithReCaptchaRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "invokeWithReCaptcha",
		ID:   InvokeWithReCaptchaRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Token",
			SchemaName: "token",
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InvokeWithReCaptchaRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeWithReCaptcha#adbb0f94 as nil")
	}
	b.PutID(InvokeWithReCaptchaRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InvokeWithReCaptchaRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeWithReCaptcha#adbb0f94 as nil")
	}
	b.PutString(i.Token)
	if err := i.Query.Encode(b); err != nil {
		return fmt.Errorf("unable to encode invokeWithReCaptcha#adbb0f94: field query: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InvokeWithReCaptchaRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeWithReCaptcha#adbb0f94 to nil")
	}
	if err := b.ConsumeID(InvokeWithReCaptchaRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode invokeWithReCaptcha#adbb0f94: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InvokeWithReCaptchaRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeWithReCaptcha#adbb0f94 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode invokeWithReCaptcha#adbb0f94: field token: %w", err)
		}
		i.Token = value
	}
	{
		if err := i.Query.Decode(b); err != nil {
			return fmt.Errorf("unable to decode invokeWithReCaptcha#adbb0f94: field query: %w", err)
		}
	}
	return nil
}

// GetToken returns value of Token field.
func (i *InvokeWithReCaptchaRequest) GetToken() (value string) {
	if i == nil {
		return
	}
	return i.Token
}

// GetQuery returns value of Query field.
func (i *InvokeWithReCaptchaRequest) GetQuery() (value bin.Object) {
	if i == nil {
		return
	}
	return i.Query
}
