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

// PaymentsPaymentResultClassArray is adapter for slice of PaymentsPaymentResultClass.
type PaymentsPaymentResultClassArray []PaymentsPaymentResultClass

// Sort sorts slice of PaymentsPaymentResultClass.
func (s PaymentsPaymentResultClassArray) Sort(less func(a, b PaymentsPaymentResultClass) bool) PaymentsPaymentResultClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PaymentsPaymentResultClass.
func (s PaymentsPaymentResultClassArray) SortStable(less func(a, b PaymentsPaymentResultClass) bool) PaymentsPaymentResultClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PaymentsPaymentResultClass.
func (s PaymentsPaymentResultClassArray) Retain(keep func(x PaymentsPaymentResultClass) bool) PaymentsPaymentResultClassArray {
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
func (s PaymentsPaymentResultClassArray) First() (v PaymentsPaymentResultClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PaymentsPaymentResultClassArray) Last() (v PaymentsPaymentResultClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PaymentsPaymentResultClassArray) PopFirst() (v PaymentsPaymentResultClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PaymentsPaymentResultClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PaymentsPaymentResultClassArray) Pop() (v PaymentsPaymentResultClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsPaymentsPaymentResult returns copy with only PaymentsPaymentResult constructors.
func (s PaymentsPaymentResultClassArray) AsPaymentsPaymentResult() (to PaymentsPaymentResultArray) {
	for _, elem := range s {
		value, ok := elem.(*PaymentsPaymentResult)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsPaymentsPaymentVerificationNeeded returns copy with only PaymentsPaymentVerificationNeeded constructors.
func (s PaymentsPaymentResultClassArray) AsPaymentsPaymentVerificationNeeded() (to PaymentsPaymentVerificationNeededArray) {
	for _, elem := range s {
		value, ok := elem.(*PaymentsPaymentVerificationNeeded)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// PaymentsPaymentResultArray is adapter for slice of PaymentsPaymentResult.
type PaymentsPaymentResultArray []PaymentsPaymentResult

// Sort sorts slice of PaymentsPaymentResult.
func (s PaymentsPaymentResultArray) Sort(less func(a, b PaymentsPaymentResult) bool) PaymentsPaymentResultArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PaymentsPaymentResult.
func (s PaymentsPaymentResultArray) SortStable(less func(a, b PaymentsPaymentResult) bool) PaymentsPaymentResultArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PaymentsPaymentResult.
func (s PaymentsPaymentResultArray) Retain(keep func(x PaymentsPaymentResult) bool) PaymentsPaymentResultArray {
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
func (s PaymentsPaymentResultArray) First() (v PaymentsPaymentResult, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PaymentsPaymentResultArray) Last() (v PaymentsPaymentResult, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PaymentsPaymentResultArray) PopFirst() (v PaymentsPaymentResult, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PaymentsPaymentResult
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PaymentsPaymentResultArray) Pop() (v PaymentsPaymentResult, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// PaymentsPaymentVerificationNeededArray is adapter for slice of PaymentsPaymentVerificationNeeded.
type PaymentsPaymentVerificationNeededArray []PaymentsPaymentVerificationNeeded

// Sort sorts slice of PaymentsPaymentVerificationNeeded.
func (s PaymentsPaymentVerificationNeededArray) Sort(less func(a, b PaymentsPaymentVerificationNeeded) bool) PaymentsPaymentVerificationNeededArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PaymentsPaymentVerificationNeeded.
func (s PaymentsPaymentVerificationNeededArray) SortStable(less func(a, b PaymentsPaymentVerificationNeeded) bool) PaymentsPaymentVerificationNeededArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PaymentsPaymentVerificationNeeded.
func (s PaymentsPaymentVerificationNeededArray) Retain(keep func(x PaymentsPaymentVerificationNeeded) bool) PaymentsPaymentVerificationNeededArray {
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
func (s PaymentsPaymentVerificationNeededArray) First() (v PaymentsPaymentVerificationNeeded, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PaymentsPaymentVerificationNeededArray) Last() (v PaymentsPaymentVerificationNeeded, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PaymentsPaymentVerificationNeededArray) PopFirst() (v PaymentsPaymentVerificationNeeded, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PaymentsPaymentVerificationNeeded
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PaymentsPaymentVerificationNeededArray) Pop() (v PaymentsPaymentVerificationNeeded, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
