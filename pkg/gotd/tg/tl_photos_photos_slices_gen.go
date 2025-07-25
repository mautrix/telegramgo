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

// PhotosPhotosClassArray is adapter for slice of PhotosPhotosClass.
type PhotosPhotosClassArray []PhotosPhotosClass

// Sort sorts slice of PhotosPhotosClass.
func (s PhotosPhotosClassArray) Sort(less func(a, b PhotosPhotosClass) bool) PhotosPhotosClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotosPhotosClass.
func (s PhotosPhotosClassArray) SortStable(less func(a, b PhotosPhotosClass) bool) PhotosPhotosClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotosPhotosClass.
func (s PhotosPhotosClassArray) Retain(keep func(x PhotosPhotosClass) bool) PhotosPhotosClassArray {
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
func (s PhotosPhotosClassArray) First() (v PhotosPhotosClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotosPhotosClassArray) Last() (v PhotosPhotosClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotosPhotosClassArray) PopFirst() (v PhotosPhotosClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotosPhotosClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotosPhotosClassArray) Pop() (v PhotosPhotosClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsPhotosPhotos returns copy with only PhotosPhotos constructors.
func (s PhotosPhotosClassArray) AsPhotosPhotos() (to PhotosPhotosArray) {
	for _, elem := range s {
		value, ok := elem.(*PhotosPhotos)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsPhotosPhotosSlice returns copy with only PhotosPhotosSlice constructors.
func (s PhotosPhotosClassArray) AsPhotosPhotosSlice() (to PhotosPhotosSliceArray) {
	for _, elem := range s {
		value, ok := elem.(*PhotosPhotosSlice)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// PhotosPhotosArray is adapter for slice of PhotosPhotos.
type PhotosPhotosArray []PhotosPhotos

// Sort sorts slice of PhotosPhotos.
func (s PhotosPhotosArray) Sort(less func(a, b PhotosPhotos) bool) PhotosPhotosArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotosPhotos.
func (s PhotosPhotosArray) SortStable(less func(a, b PhotosPhotos) bool) PhotosPhotosArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotosPhotos.
func (s PhotosPhotosArray) Retain(keep func(x PhotosPhotos) bool) PhotosPhotosArray {
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
func (s PhotosPhotosArray) First() (v PhotosPhotos, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotosPhotosArray) Last() (v PhotosPhotos, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotosPhotosArray) PopFirst() (v PhotosPhotos, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotosPhotos
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotosPhotosArray) Pop() (v PhotosPhotos, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// PhotosPhotosSliceArray is adapter for slice of PhotosPhotosSlice.
type PhotosPhotosSliceArray []PhotosPhotosSlice

// Sort sorts slice of PhotosPhotosSlice.
func (s PhotosPhotosSliceArray) Sort(less func(a, b PhotosPhotosSlice) bool) PhotosPhotosSliceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotosPhotosSlice.
func (s PhotosPhotosSliceArray) SortStable(less func(a, b PhotosPhotosSlice) bool) PhotosPhotosSliceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotosPhotosSlice.
func (s PhotosPhotosSliceArray) Retain(keep func(x PhotosPhotosSlice) bool) PhotosPhotosSliceArray {
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
func (s PhotosPhotosSliceArray) First() (v PhotosPhotosSlice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotosPhotosSliceArray) Last() (v PhotosPhotosSlice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotosPhotosSliceArray) PopFirst() (v PhotosPhotosSlice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotosPhotosSlice
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotosPhotosSliceArray) Pop() (v PhotosPhotosSlice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
