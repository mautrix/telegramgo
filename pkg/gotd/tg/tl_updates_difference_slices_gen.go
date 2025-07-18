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

// UpdatesDifferenceClassArray is adapter for slice of UpdatesDifferenceClass.
type UpdatesDifferenceClassArray []UpdatesDifferenceClass

// Sort sorts slice of UpdatesDifferenceClass.
func (s UpdatesDifferenceClassArray) Sort(less func(a, b UpdatesDifferenceClass) bool) UpdatesDifferenceClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UpdatesDifferenceClass.
func (s UpdatesDifferenceClassArray) SortStable(less func(a, b UpdatesDifferenceClass) bool) UpdatesDifferenceClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UpdatesDifferenceClass.
func (s UpdatesDifferenceClassArray) Retain(keep func(x UpdatesDifferenceClass) bool) UpdatesDifferenceClassArray {
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
func (s UpdatesDifferenceClassArray) First() (v UpdatesDifferenceClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesDifferenceClassArray) Last() (v UpdatesDifferenceClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceClassArray) PopFirst() (v UpdatesDifferenceClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UpdatesDifferenceClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceClassArray) Pop() (v UpdatesDifferenceClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsUpdatesDifferenceEmpty returns copy with only UpdatesDifferenceEmpty constructors.
func (s UpdatesDifferenceClassArray) AsUpdatesDifferenceEmpty() (to UpdatesDifferenceEmptyArray) {
	for _, elem := range s {
		value, ok := elem.(*UpdatesDifferenceEmpty)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUpdatesDifference returns copy with only UpdatesDifference constructors.
func (s UpdatesDifferenceClassArray) AsUpdatesDifference() (to UpdatesDifferenceArray) {
	for _, elem := range s {
		value, ok := elem.(*UpdatesDifference)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUpdatesDifferenceSlice returns copy with only UpdatesDifferenceSlice constructors.
func (s UpdatesDifferenceClassArray) AsUpdatesDifferenceSlice() (to UpdatesDifferenceSliceArray) {
	for _, elem := range s {
		value, ok := elem.(*UpdatesDifferenceSlice)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUpdatesDifferenceTooLong returns copy with only UpdatesDifferenceTooLong constructors.
func (s UpdatesDifferenceClassArray) AsUpdatesDifferenceTooLong() (to UpdatesDifferenceTooLongArray) {
	for _, elem := range s {
		value, ok := elem.(*UpdatesDifferenceTooLong)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// UpdatesDifferenceEmptyArray is adapter for slice of UpdatesDifferenceEmpty.
type UpdatesDifferenceEmptyArray []UpdatesDifferenceEmpty

// Sort sorts slice of UpdatesDifferenceEmpty.
func (s UpdatesDifferenceEmptyArray) Sort(less func(a, b UpdatesDifferenceEmpty) bool) UpdatesDifferenceEmptyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UpdatesDifferenceEmpty.
func (s UpdatesDifferenceEmptyArray) SortStable(less func(a, b UpdatesDifferenceEmpty) bool) UpdatesDifferenceEmptyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UpdatesDifferenceEmpty.
func (s UpdatesDifferenceEmptyArray) Retain(keep func(x UpdatesDifferenceEmpty) bool) UpdatesDifferenceEmptyArray {
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
func (s UpdatesDifferenceEmptyArray) First() (v UpdatesDifferenceEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesDifferenceEmptyArray) Last() (v UpdatesDifferenceEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceEmptyArray) PopFirst() (v UpdatesDifferenceEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UpdatesDifferenceEmpty
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceEmptyArray) Pop() (v UpdatesDifferenceEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of UpdatesDifferenceEmpty by Date.
func (s UpdatesDifferenceEmptyArray) SortByDate() UpdatesDifferenceEmptyArray {
	return s.Sort(func(a, b UpdatesDifferenceEmpty) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of UpdatesDifferenceEmpty by Date.
func (s UpdatesDifferenceEmptyArray) SortStableByDate() UpdatesDifferenceEmptyArray {
	return s.SortStable(func(a, b UpdatesDifferenceEmpty) bool {
		return a.GetDate() < b.GetDate()
	})
}

// UpdatesDifferenceArray is adapter for slice of UpdatesDifference.
type UpdatesDifferenceArray []UpdatesDifference

// Sort sorts slice of UpdatesDifference.
func (s UpdatesDifferenceArray) Sort(less func(a, b UpdatesDifference) bool) UpdatesDifferenceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UpdatesDifference.
func (s UpdatesDifferenceArray) SortStable(less func(a, b UpdatesDifference) bool) UpdatesDifferenceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UpdatesDifference.
func (s UpdatesDifferenceArray) Retain(keep func(x UpdatesDifference) bool) UpdatesDifferenceArray {
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
func (s UpdatesDifferenceArray) First() (v UpdatesDifference, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesDifferenceArray) Last() (v UpdatesDifference, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceArray) PopFirst() (v UpdatesDifference, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UpdatesDifference
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceArray) Pop() (v UpdatesDifference, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// UpdatesDifferenceSliceArray is adapter for slice of UpdatesDifferenceSlice.
type UpdatesDifferenceSliceArray []UpdatesDifferenceSlice

// Sort sorts slice of UpdatesDifferenceSlice.
func (s UpdatesDifferenceSliceArray) Sort(less func(a, b UpdatesDifferenceSlice) bool) UpdatesDifferenceSliceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UpdatesDifferenceSlice.
func (s UpdatesDifferenceSliceArray) SortStable(less func(a, b UpdatesDifferenceSlice) bool) UpdatesDifferenceSliceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UpdatesDifferenceSlice.
func (s UpdatesDifferenceSliceArray) Retain(keep func(x UpdatesDifferenceSlice) bool) UpdatesDifferenceSliceArray {
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
func (s UpdatesDifferenceSliceArray) First() (v UpdatesDifferenceSlice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesDifferenceSliceArray) Last() (v UpdatesDifferenceSlice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceSliceArray) PopFirst() (v UpdatesDifferenceSlice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UpdatesDifferenceSlice
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceSliceArray) Pop() (v UpdatesDifferenceSlice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// UpdatesDifferenceTooLongArray is adapter for slice of UpdatesDifferenceTooLong.
type UpdatesDifferenceTooLongArray []UpdatesDifferenceTooLong

// Sort sorts slice of UpdatesDifferenceTooLong.
func (s UpdatesDifferenceTooLongArray) Sort(less func(a, b UpdatesDifferenceTooLong) bool) UpdatesDifferenceTooLongArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UpdatesDifferenceTooLong.
func (s UpdatesDifferenceTooLongArray) SortStable(less func(a, b UpdatesDifferenceTooLong) bool) UpdatesDifferenceTooLongArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UpdatesDifferenceTooLong.
func (s UpdatesDifferenceTooLongArray) Retain(keep func(x UpdatesDifferenceTooLong) bool) UpdatesDifferenceTooLongArray {
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
func (s UpdatesDifferenceTooLongArray) First() (v UpdatesDifferenceTooLong, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesDifferenceTooLongArray) Last() (v UpdatesDifferenceTooLong, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceTooLongArray) PopFirst() (v UpdatesDifferenceTooLong, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UpdatesDifferenceTooLong
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceTooLongArray) Pop() (v UpdatesDifferenceTooLong, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
