// Code generated by gotdgen, DO NOT EDIT.

package tgtrace

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

// TypesMap returns mapping from type ids to TL type names.
func TypesMap() map[uint32]string {
	return map[uint32]string{
		InvokeWithTraceRequestTypeID: "invokeWithTrace#104be9ab",
	}
}

// NamesMap returns mapping from type names to TL type ids.
func NamesMap() map[string]uint32 {
	return map[string]uint32{
		"invokeWithTrace": InvokeWithTraceRequestTypeID,
	}
}

// TypesConstructorMap maps type ids to constructors.
func TypesConstructorMap() map[uint32]func() bin.Object {
	return map[uint32]func() bin.Object{
		InvokeWithTraceRequestTypeID: func() bin.Object { return &InvokeWithTraceRequest{} },
	}
}

// ClassConstructorsMap maps class schema name to constructors type ids.
func ClassConstructorsMap() map[string][]uint32 {
	return map[string][]uint32{}
}
