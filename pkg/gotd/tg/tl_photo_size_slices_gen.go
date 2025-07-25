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

// PhotoSizeClassArray is adapter for slice of PhotoSizeClass.
type PhotoSizeClassArray []PhotoSizeClass

// Sort sorts slice of PhotoSizeClass.
func (s PhotoSizeClassArray) Sort(less func(a, b PhotoSizeClass) bool) PhotoSizeClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotoSizeClass.
func (s PhotoSizeClassArray) SortStable(less func(a, b PhotoSizeClass) bool) PhotoSizeClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotoSizeClass.
func (s PhotoSizeClassArray) Retain(keep func(x PhotoSizeClass) bool) PhotoSizeClassArray {
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
func (s PhotoSizeClassArray) First() (v PhotoSizeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotoSizeClassArray) Last() (v PhotoSizeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotoSizeClassArray) PopFirst() (v PhotoSizeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotoSizeClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotoSizeClassArray) Pop() (v PhotoSizeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsPhotoSizeEmpty returns copy with only PhotoSizeEmpty constructors.
func (s PhotoSizeClassArray) AsPhotoSizeEmpty() (to PhotoSizeEmptyArray) {
	for _, elem := range s {
		value, ok := elem.(*PhotoSizeEmpty)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsPhotoSize returns copy with only PhotoSize constructors.
func (s PhotoSizeClassArray) AsPhotoSize() (to PhotoSizeArray) {
	for _, elem := range s {
		value, ok := elem.(*PhotoSize)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsPhotoCachedSize returns copy with only PhotoCachedSize constructors.
func (s PhotoSizeClassArray) AsPhotoCachedSize() (to PhotoCachedSizeArray) {
	for _, elem := range s {
		value, ok := elem.(*PhotoCachedSize)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsPhotoStrippedSize returns copy with only PhotoStrippedSize constructors.
func (s PhotoSizeClassArray) AsPhotoStrippedSize() (to PhotoStrippedSizeArray) {
	for _, elem := range s {
		value, ok := elem.(*PhotoStrippedSize)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsPhotoSizeProgressive returns copy with only PhotoSizeProgressive constructors.
func (s PhotoSizeClassArray) AsPhotoSizeProgressive() (to PhotoSizeProgressiveArray) {
	for _, elem := range s {
		value, ok := elem.(*PhotoSizeProgressive)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsPhotoPathSize returns copy with only PhotoPathSize constructors.
func (s PhotoSizeClassArray) AsPhotoPathSize() (to PhotoPathSizeArray) {
	for _, elem := range s {
		value, ok := elem.(*PhotoPathSize)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s PhotoSizeClassArray) AppendOnlyNotEmpty(to []NotEmptyPhotoSize) []NotEmptyPhotoSize {
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
func (s PhotoSizeClassArray) AsNotEmpty() (to []NotEmptyPhotoSize) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s PhotoSizeClassArray) FirstAsNotEmpty() (v NotEmptyPhotoSize, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s PhotoSizeClassArray) LastAsNotEmpty() (v NotEmptyPhotoSize, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *PhotoSizeClassArray) PopFirstAsNotEmpty() (v NotEmptyPhotoSize, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *PhotoSizeClassArray) PopAsNotEmpty() (v NotEmptyPhotoSize, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PhotoSizeEmptyArray is adapter for slice of PhotoSizeEmpty.
type PhotoSizeEmptyArray []PhotoSizeEmpty

// Sort sorts slice of PhotoSizeEmpty.
func (s PhotoSizeEmptyArray) Sort(less func(a, b PhotoSizeEmpty) bool) PhotoSizeEmptyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotoSizeEmpty.
func (s PhotoSizeEmptyArray) SortStable(less func(a, b PhotoSizeEmpty) bool) PhotoSizeEmptyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotoSizeEmpty.
func (s PhotoSizeEmptyArray) Retain(keep func(x PhotoSizeEmpty) bool) PhotoSizeEmptyArray {
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
func (s PhotoSizeEmptyArray) First() (v PhotoSizeEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotoSizeEmptyArray) Last() (v PhotoSizeEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotoSizeEmptyArray) PopFirst() (v PhotoSizeEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotoSizeEmpty
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotoSizeEmptyArray) Pop() (v PhotoSizeEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// PhotoSizeArray is adapter for slice of PhotoSize.
type PhotoSizeArray []PhotoSize

// Sort sorts slice of PhotoSize.
func (s PhotoSizeArray) Sort(less func(a, b PhotoSize) bool) PhotoSizeArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotoSize.
func (s PhotoSizeArray) SortStable(less func(a, b PhotoSize) bool) PhotoSizeArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotoSize.
func (s PhotoSizeArray) Retain(keep func(x PhotoSize) bool) PhotoSizeArray {
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
func (s PhotoSizeArray) First() (v PhotoSize, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotoSizeArray) Last() (v PhotoSize, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotoSizeArray) PopFirst() (v PhotoSize, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotoSize
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotoSizeArray) Pop() (v PhotoSize, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// PhotoCachedSizeArray is adapter for slice of PhotoCachedSize.
type PhotoCachedSizeArray []PhotoCachedSize

// Sort sorts slice of PhotoCachedSize.
func (s PhotoCachedSizeArray) Sort(less func(a, b PhotoCachedSize) bool) PhotoCachedSizeArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotoCachedSize.
func (s PhotoCachedSizeArray) SortStable(less func(a, b PhotoCachedSize) bool) PhotoCachedSizeArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotoCachedSize.
func (s PhotoCachedSizeArray) Retain(keep func(x PhotoCachedSize) bool) PhotoCachedSizeArray {
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
func (s PhotoCachedSizeArray) First() (v PhotoCachedSize, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotoCachedSizeArray) Last() (v PhotoCachedSize, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotoCachedSizeArray) PopFirst() (v PhotoCachedSize, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotoCachedSize
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotoCachedSizeArray) Pop() (v PhotoCachedSize, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// PhotoStrippedSizeArray is adapter for slice of PhotoStrippedSize.
type PhotoStrippedSizeArray []PhotoStrippedSize

// Sort sorts slice of PhotoStrippedSize.
func (s PhotoStrippedSizeArray) Sort(less func(a, b PhotoStrippedSize) bool) PhotoStrippedSizeArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotoStrippedSize.
func (s PhotoStrippedSizeArray) SortStable(less func(a, b PhotoStrippedSize) bool) PhotoStrippedSizeArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotoStrippedSize.
func (s PhotoStrippedSizeArray) Retain(keep func(x PhotoStrippedSize) bool) PhotoStrippedSizeArray {
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
func (s PhotoStrippedSizeArray) First() (v PhotoStrippedSize, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotoStrippedSizeArray) Last() (v PhotoStrippedSize, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotoStrippedSizeArray) PopFirst() (v PhotoStrippedSize, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotoStrippedSize
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotoStrippedSizeArray) Pop() (v PhotoStrippedSize, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// PhotoSizeProgressiveArray is adapter for slice of PhotoSizeProgressive.
type PhotoSizeProgressiveArray []PhotoSizeProgressive

// Sort sorts slice of PhotoSizeProgressive.
func (s PhotoSizeProgressiveArray) Sort(less func(a, b PhotoSizeProgressive) bool) PhotoSizeProgressiveArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotoSizeProgressive.
func (s PhotoSizeProgressiveArray) SortStable(less func(a, b PhotoSizeProgressive) bool) PhotoSizeProgressiveArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotoSizeProgressive.
func (s PhotoSizeProgressiveArray) Retain(keep func(x PhotoSizeProgressive) bool) PhotoSizeProgressiveArray {
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
func (s PhotoSizeProgressiveArray) First() (v PhotoSizeProgressive, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotoSizeProgressiveArray) Last() (v PhotoSizeProgressive, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotoSizeProgressiveArray) PopFirst() (v PhotoSizeProgressive, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotoSizeProgressive
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotoSizeProgressiveArray) Pop() (v PhotoSizeProgressive, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// PhotoPathSizeArray is adapter for slice of PhotoPathSize.
type PhotoPathSizeArray []PhotoPathSize

// Sort sorts slice of PhotoPathSize.
func (s PhotoPathSizeArray) Sort(less func(a, b PhotoPathSize) bool) PhotoPathSizeArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotoPathSize.
func (s PhotoPathSizeArray) SortStable(less func(a, b PhotoPathSize) bool) PhotoPathSizeArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotoPathSize.
func (s PhotoPathSizeArray) Retain(keep func(x PhotoPathSize) bool) PhotoPathSizeArray {
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
func (s PhotoPathSizeArray) First() (v PhotoPathSize, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotoPathSizeArray) Last() (v PhotoPathSize, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotoPathSizeArray) PopFirst() (v PhotoPathSize, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotoPathSize
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotoPathSizeArray) Pop() (v PhotoPathSize, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
