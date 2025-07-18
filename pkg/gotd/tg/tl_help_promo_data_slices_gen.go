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

// HelpPromoDataClassArray is adapter for slice of HelpPromoDataClass.
type HelpPromoDataClassArray []HelpPromoDataClass

// Sort sorts slice of HelpPromoDataClass.
func (s HelpPromoDataClassArray) Sort(less func(a, b HelpPromoDataClass) bool) HelpPromoDataClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpPromoDataClass.
func (s HelpPromoDataClassArray) SortStable(less func(a, b HelpPromoDataClass) bool) HelpPromoDataClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpPromoDataClass.
func (s HelpPromoDataClassArray) Retain(keep func(x HelpPromoDataClass) bool) HelpPromoDataClassArray {
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
func (s HelpPromoDataClassArray) First() (v HelpPromoDataClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpPromoDataClassArray) Last() (v HelpPromoDataClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpPromoDataClassArray) PopFirst() (v HelpPromoDataClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpPromoDataClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpPromoDataClassArray) Pop() (v HelpPromoDataClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsHelpPromoDataEmpty returns copy with only HelpPromoDataEmpty constructors.
func (s HelpPromoDataClassArray) AsHelpPromoDataEmpty() (to HelpPromoDataEmptyArray) {
	for _, elem := range s {
		value, ok := elem.(*HelpPromoDataEmpty)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsHelpPromoData returns copy with only HelpPromoData constructors.
func (s HelpPromoDataClassArray) AsHelpPromoData() (to HelpPromoDataArray) {
	for _, elem := range s {
		value, ok := elem.(*HelpPromoData)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s HelpPromoDataClassArray) AppendOnlyNotEmpty(to []*HelpPromoData) []*HelpPromoData {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s HelpPromoDataClassArray) AsNotEmpty() (to []*HelpPromoData) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s HelpPromoDataClassArray) FirstAsNotEmpty() (v *HelpPromoData, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s HelpPromoDataClassArray) LastAsNotEmpty() (v *HelpPromoData, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *HelpPromoDataClassArray) PopFirstAsNotEmpty() (v *HelpPromoData, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *HelpPromoDataClassArray) PopAsNotEmpty() (v *HelpPromoData, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// HelpPromoDataEmptyArray is adapter for slice of HelpPromoDataEmpty.
type HelpPromoDataEmptyArray []HelpPromoDataEmpty

// Sort sorts slice of HelpPromoDataEmpty.
func (s HelpPromoDataEmptyArray) Sort(less func(a, b HelpPromoDataEmpty) bool) HelpPromoDataEmptyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpPromoDataEmpty.
func (s HelpPromoDataEmptyArray) SortStable(less func(a, b HelpPromoDataEmpty) bool) HelpPromoDataEmptyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpPromoDataEmpty.
func (s HelpPromoDataEmptyArray) Retain(keep func(x HelpPromoDataEmpty) bool) HelpPromoDataEmptyArray {
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
func (s HelpPromoDataEmptyArray) First() (v HelpPromoDataEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpPromoDataEmptyArray) Last() (v HelpPromoDataEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpPromoDataEmptyArray) PopFirst() (v HelpPromoDataEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpPromoDataEmpty
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpPromoDataEmptyArray) Pop() (v HelpPromoDataEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// HelpPromoDataArray is adapter for slice of HelpPromoData.
type HelpPromoDataArray []HelpPromoData

// Sort sorts slice of HelpPromoData.
func (s HelpPromoDataArray) Sort(less func(a, b HelpPromoData) bool) HelpPromoDataArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpPromoData.
func (s HelpPromoDataArray) SortStable(less func(a, b HelpPromoData) bool) HelpPromoDataArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpPromoData.
func (s HelpPromoDataArray) Retain(keep func(x HelpPromoData) bool) HelpPromoDataArray {
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
func (s HelpPromoDataArray) First() (v HelpPromoData, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpPromoDataArray) Last() (v HelpPromoData, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpPromoDataArray) PopFirst() (v HelpPromoData, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpPromoData
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpPromoDataArray) Pop() (v HelpPromoData, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
