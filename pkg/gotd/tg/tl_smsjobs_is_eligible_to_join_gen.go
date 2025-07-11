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

// SMSJobsIsEligibleToJoinRequest represents TL type `smsjobs.isEligibleToJoin#edc39d0`.
// Check if we can process SMS jobs (official clients only).
//
// See https://core.telegram.org/method/smsjobs.isEligibleToJoin for reference.
type SMSJobsIsEligibleToJoinRequest struct {
}

// SMSJobsIsEligibleToJoinRequestTypeID is TL type id of SMSJobsIsEligibleToJoinRequest.
const SMSJobsIsEligibleToJoinRequestTypeID = 0xedc39d0

// Ensuring interfaces in compile-time for SMSJobsIsEligibleToJoinRequest.
var (
	_ bin.Encoder     = &SMSJobsIsEligibleToJoinRequest{}
	_ bin.Decoder     = &SMSJobsIsEligibleToJoinRequest{}
	_ bin.BareEncoder = &SMSJobsIsEligibleToJoinRequest{}
	_ bin.BareDecoder = &SMSJobsIsEligibleToJoinRequest{}
)

func (i *SMSJobsIsEligibleToJoinRequest) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *SMSJobsIsEligibleToJoinRequest) String() string {
	if i == nil {
		return "SMSJobsIsEligibleToJoinRequest(nil)"
	}
	type Alias SMSJobsIsEligibleToJoinRequest
	return fmt.Sprintf("SMSJobsIsEligibleToJoinRequest%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SMSJobsIsEligibleToJoinRequest) TypeID() uint32 {
	return SMSJobsIsEligibleToJoinRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SMSJobsIsEligibleToJoinRequest) TypeName() string {
	return "smsjobs.isEligibleToJoin"
}

// TypeInfo returns info about TL type.
func (i *SMSJobsIsEligibleToJoinRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "smsjobs.isEligibleToJoin",
		ID:   SMSJobsIsEligibleToJoinRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *SMSJobsIsEligibleToJoinRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode smsjobs.isEligibleToJoin#edc39d0 as nil")
	}
	b.PutID(SMSJobsIsEligibleToJoinRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *SMSJobsIsEligibleToJoinRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode smsjobs.isEligibleToJoin#edc39d0 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *SMSJobsIsEligibleToJoinRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode smsjobs.isEligibleToJoin#edc39d0 to nil")
	}
	if err := b.ConsumeID(SMSJobsIsEligibleToJoinRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode smsjobs.isEligibleToJoin#edc39d0: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *SMSJobsIsEligibleToJoinRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode smsjobs.isEligibleToJoin#edc39d0 to nil")
	}
	return nil
}

// SMSJobsIsEligibleToJoin invokes method smsjobs.isEligibleToJoin#edc39d0 returning error if any.
// Check if we can process SMS jobs (official clients only).
//
// Possible errors:
//
//	403 NOT_ELIGIBLE: The current user is not eligible to join the Peer-to-Peer Login Program.
//
// See https://core.telegram.org/method/smsjobs.isEligibleToJoin for reference.
func (c *Client) SMSJobsIsEligibleToJoin(ctx context.Context) (*SMSJobsEligibleToJoin, error) {
	var result SMSJobsEligibleToJoin

	request := &SMSJobsIsEligibleToJoinRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
