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

// URLAuthResultClassArray is adapter for slice of URLAuthResultClass.
type URLAuthResultClassArray []URLAuthResultClass

// Sort sorts slice of URLAuthResultClass.
func (s URLAuthResultClassArray) Sort(less func(a, b URLAuthResultClass) bool) URLAuthResultClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of URLAuthResultClass.
func (s URLAuthResultClassArray) SortStable(less func(a, b URLAuthResultClass) bool) URLAuthResultClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of URLAuthResultClass.
func (s URLAuthResultClassArray) Retain(keep func(x URLAuthResultClass) bool) URLAuthResultClassArray {
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
func (s URLAuthResultClassArray) First() (v URLAuthResultClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s URLAuthResultClassArray) Last() (v URLAuthResultClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *URLAuthResultClassArray) PopFirst() (v URLAuthResultClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero URLAuthResultClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *URLAuthResultClassArray) Pop() (v URLAuthResultClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsURLAuthResultRequest returns copy with only URLAuthResultRequest constructors.
func (s URLAuthResultClassArray) AsURLAuthResultRequest() (to URLAuthResultRequestArray) {
	for _, elem := range s {
		value, ok := elem.(*URLAuthResultRequest)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsURLAuthResultAccepted returns copy with only URLAuthResultAccepted constructors.
func (s URLAuthResultClassArray) AsURLAuthResultAccepted() (to URLAuthResultAcceptedArray) {
	for _, elem := range s {
		value, ok := elem.(*URLAuthResultAccepted)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// URLAuthResultRequestArray is adapter for slice of URLAuthResultRequest.
type URLAuthResultRequestArray []URLAuthResultRequest

// Sort sorts slice of URLAuthResultRequest.
func (s URLAuthResultRequestArray) Sort(less func(a, b URLAuthResultRequest) bool) URLAuthResultRequestArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of URLAuthResultRequest.
func (s URLAuthResultRequestArray) SortStable(less func(a, b URLAuthResultRequest) bool) URLAuthResultRequestArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of URLAuthResultRequest.
func (s URLAuthResultRequestArray) Retain(keep func(x URLAuthResultRequest) bool) URLAuthResultRequestArray {
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
func (s URLAuthResultRequestArray) First() (v URLAuthResultRequest, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s URLAuthResultRequestArray) Last() (v URLAuthResultRequest, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *URLAuthResultRequestArray) PopFirst() (v URLAuthResultRequest, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero URLAuthResultRequest
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *URLAuthResultRequestArray) Pop() (v URLAuthResultRequest, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// URLAuthResultAcceptedArray is adapter for slice of URLAuthResultAccepted.
type URLAuthResultAcceptedArray []URLAuthResultAccepted

// Sort sorts slice of URLAuthResultAccepted.
func (s URLAuthResultAcceptedArray) Sort(less func(a, b URLAuthResultAccepted) bool) URLAuthResultAcceptedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of URLAuthResultAccepted.
func (s URLAuthResultAcceptedArray) SortStable(less func(a, b URLAuthResultAccepted) bool) URLAuthResultAcceptedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of URLAuthResultAccepted.
func (s URLAuthResultAcceptedArray) Retain(keep func(x URLAuthResultAccepted) bool) URLAuthResultAcceptedArray {
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
func (s URLAuthResultAcceptedArray) First() (v URLAuthResultAccepted, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s URLAuthResultAcceptedArray) Last() (v URLAuthResultAccepted, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *URLAuthResultAcceptedArray) PopFirst() (v URLAuthResultAccepted, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero URLAuthResultAccepted
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *URLAuthResultAcceptedArray) Pop() (v URLAuthResultAccepted, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
