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

// InputStickeredMediaClassArray is adapter for slice of InputStickeredMediaClass.
type InputStickeredMediaClassArray []InputStickeredMediaClass

// Sort sorts slice of InputStickeredMediaClass.
func (s InputStickeredMediaClassArray) Sort(less func(a, b InputStickeredMediaClass) bool) InputStickeredMediaClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputStickeredMediaClass.
func (s InputStickeredMediaClassArray) SortStable(less func(a, b InputStickeredMediaClass) bool) InputStickeredMediaClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputStickeredMediaClass.
func (s InputStickeredMediaClassArray) Retain(keep func(x InputStickeredMediaClass) bool) InputStickeredMediaClassArray {
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
func (s InputStickeredMediaClassArray) First() (v InputStickeredMediaClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputStickeredMediaClassArray) Last() (v InputStickeredMediaClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputStickeredMediaClassArray) PopFirst() (v InputStickeredMediaClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputStickeredMediaClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputStickeredMediaClassArray) Pop() (v InputStickeredMediaClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputStickeredMediaPhoto returns copy with only InputStickeredMediaPhoto constructors.
func (s InputStickeredMediaClassArray) AsInputStickeredMediaPhoto() (to InputStickeredMediaPhotoArray) {
	for _, elem := range s {
		value, ok := elem.(*InputStickeredMediaPhoto)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputStickeredMediaDocument returns copy with only InputStickeredMediaDocument constructors.
func (s InputStickeredMediaClassArray) AsInputStickeredMediaDocument() (to InputStickeredMediaDocumentArray) {
	for _, elem := range s {
		value, ok := elem.(*InputStickeredMediaDocument)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputStickeredMediaPhotoArray is adapter for slice of InputStickeredMediaPhoto.
type InputStickeredMediaPhotoArray []InputStickeredMediaPhoto

// Sort sorts slice of InputStickeredMediaPhoto.
func (s InputStickeredMediaPhotoArray) Sort(less func(a, b InputStickeredMediaPhoto) bool) InputStickeredMediaPhotoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputStickeredMediaPhoto.
func (s InputStickeredMediaPhotoArray) SortStable(less func(a, b InputStickeredMediaPhoto) bool) InputStickeredMediaPhotoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputStickeredMediaPhoto.
func (s InputStickeredMediaPhotoArray) Retain(keep func(x InputStickeredMediaPhoto) bool) InputStickeredMediaPhotoArray {
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
func (s InputStickeredMediaPhotoArray) First() (v InputStickeredMediaPhoto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputStickeredMediaPhotoArray) Last() (v InputStickeredMediaPhoto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputStickeredMediaPhotoArray) PopFirst() (v InputStickeredMediaPhoto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputStickeredMediaPhoto
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputStickeredMediaPhotoArray) Pop() (v InputStickeredMediaPhoto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputStickeredMediaDocumentArray is adapter for slice of InputStickeredMediaDocument.
type InputStickeredMediaDocumentArray []InputStickeredMediaDocument

// Sort sorts slice of InputStickeredMediaDocument.
func (s InputStickeredMediaDocumentArray) Sort(less func(a, b InputStickeredMediaDocument) bool) InputStickeredMediaDocumentArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputStickeredMediaDocument.
func (s InputStickeredMediaDocumentArray) SortStable(less func(a, b InputStickeredMediaDocument) bool) InputStickeredMediaDocumentArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputStickeredMediaDocument.
func (s InputStickeredMediaDocumentArray) Retain(keep func(x InputStickeredMediaDocument) bool) InputStickeredMediaDocumentArray {
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
func (s InputStickeredMediaDocumentArray) First() (v InputStickeredMediaDocument, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputStickeredMediaDocumentArray) Last() (v InputStickeredMediaDocument, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputStickeredMediaDocumentArray) PopFirst() (v InputStickeredMediaDocument, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputStickeredMediaDocument
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputStickeredMediaDocumentArray) Pop() (v InputStickeredMediaDocument, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
