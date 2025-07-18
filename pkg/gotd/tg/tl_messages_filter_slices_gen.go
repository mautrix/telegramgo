//go:build !no_gotd_slices
// +build !no_gotd_slices

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

// MessagesFilterClassArray is adapter for slice of MessagesFilterClass.
type MessagesFilterClassArray []MessagesFilterClass

// Sort sorts slice of MessagesFilterClass.
func (s MessagesFilterClassArray) Sort(less func(a, b MessagesFilterClass) bool) MessagesFilterClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesFilterClass.
func (s MessagesFilterClassArray) SortStable(less func(a, b MessagesFilterClass) bool) MessagesFilterClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesFilterClass.
func (s MessagesFilterClassArray) Retain(keep func(x MessagesFilterClass) bool) MessagesFilterClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessagesFilterClassArray) First() (v MessagesFilterClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesFilterClassArray) Last() (v MessagesFilterClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesFilterClassArray) PopFirst() (v MessagesFilterClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesFilterClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesFilterClassArray) Pop() (v MessagesFilterClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputMessagesFilterPhoneCalls returns copy with only InputMessagesFilterPhoneCalls constructors.
func (s MessagesFilterClassArray) AsInputMessagesFilterPhoneCalls() (to InputMessagesFilterPhoneCallsArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMessagesFilterPhoneCalls)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputMessagesFilterPhoneCallsArray is adapter for slice of InputMessagesFilterPhoneCalls.
type InputMessagesFilterPhoneCallsArray []InputMessagesFilterPhoneCalls

// Sort sorts slice of InputMessagesFilterPhoneCalls.
func (s InputMessagesFilterPhoneCallsArray) Sort(less func(a, b InputMessagesFilterPhoneCalls) bool) InputMessagesFilterPhoneCallsArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMessagesFilterPhoneCalls.
func (s InputMessagesFilterPhoneCallsArray) SortStable(less func(a, b InputMessagesFilterPhoneCalls) bool) InputMessagesFilterPhoneCallsArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMessagesFilterPhoneCalls.
func (s InputMessagesFilterPhoneCallsArray) Retain(keep func(x InputMessagesFilterPhoneCalls) bool) InputMessagesFilterPhoneCallsArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputMessagesFilterPhoneCallsArray) First() (v InputMessagesFilterPhoneCalls, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMessagesFilterPhoneCallsArray) Last() (v InputMessagesFilterPhoneCalls, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMessagesFilterPhoneCallsArray) PopFirst() (v InputMessagesFilterPhoneCalls, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMessagesFilterPhoneCalls
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMessagesFilterPhoneCallsArray) Pop() (v InputMessagesFilterPhoneCalls, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
