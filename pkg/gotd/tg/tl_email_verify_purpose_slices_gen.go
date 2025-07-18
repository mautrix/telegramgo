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

// EmailVerifyPurposeClassArray is adapter for slice of EmailVerifyPurposeClass.
type EmailVerifyPurposeClassArray []EmailVerifyPurposeClass

// Sort sorts slice of EmailVerifyPurposeClass.
func (s EmailVerifyPurposeClassArray) Sort(less func(a, b EmailVerifyPurposeClass) bool) EmailVerifyPurposeClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of EmailVerifyPurposeClass.
func (s EmailVerifyPurposeClassArray) SortStable(less func(a, b EmailVerifyPurposeClass) bool) EmailVerifyPurposeClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of EmailVerifyPurposeClass.
func (s EmailVerifyPurposeClassArray) Retain(keep func(x EmailVerifyPurposeClass) bool) EmailVerifyPurposeClassArray {
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
func (s EmailVerifyPurposeClassArray) First() (v EmailVerifyPurposeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s EmailVerifyPurposeClassArray) Last() (v EmailVerifyPurposeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *EmailVerifyPurposeClassArray) PopFirst() (v EmailVerifyPurposeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero EmailVerifyPurposeClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *EmailVerifyPurposeClassArray) Pop() (v EmailVerifyPurposeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsEmailVerifyPurposeLoginSetup returns copy with only EmailVerifyPurposeLoginSetup constructors.
func (s EmailVerifyPurposeClassArray) AsEmailVerifyPurposeLoginSetup() (to EmailVerifyPurposeLoginSetupArray) {
	for _, elem := range s {
		value, ok := elem.(*EmailVerifyPurposeLoginSetup)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// EmailVerifyPurposeLoginSetupArray is adapter for slice of EmailVerifyPurposeLoginSetup.
type EmailVerifyPurposeLoginSetupArray []EmailVerifyPurposeLoginSetup

// Sort sorts slice of EmailVerifyPurposeLoginSetup.
func (s EmailVerifyPurposeLoginSetupArray) Sort(less func(a, b EmailVerifyPurposeLoginSetup) bool) EmailVerifyPurposeLoginSetupArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of EmailVerifyPurposeLoginSetup.
func (s EmailVerifyPurposeLoginSetupArray) SortStable(less func(a, b EmailVerifyPurposeLoginSetup) bool) EmailVerifyPurposeLoginSetupArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of EmailVerifyPurposeLoginSetup.
func (s EmailVerifyPurposeLoginSetupArray) Retain(keep func(x EmailVerifyPurposeLoginSetup) bool) EmailVerifyPurposeLoginSetupArray {
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
func (s EmailVerifyPurposeLoginSetupArray) First() (v EmailVerifyPurposeLoginSetup, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s EmailVerifyPurposeLoginSetupArray) Last() (v EmailVerifyPurposeLoginSetup, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *EmailVerifyPurposeLoginSetupArray) PopFirst() (v EmailVerifyPurposeLoginSetup, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero EmailVerifyPurposeLoginSetup
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *EmailVerifyPurposeLoginSetupArray) Pop() (v EmailVerifyPurposeLoginSetup, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
