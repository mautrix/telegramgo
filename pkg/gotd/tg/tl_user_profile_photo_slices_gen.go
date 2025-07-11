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

// UserProfilePhotoClassArray is adapter for slice of UserProfilePhotoClass.
type UserProfilePhotoClassArray []UserProfilePhotoClass

// Sort sorts slice of UserProfilePhotoClass.
func (s UserProfilePhotoClassArray) Sort(less func(a, b UserProfilePhotoClass) bool) UserProfilePhotoClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UserProfilePhotoClass.
func (s UserProfilePhotoClassArray) SortStable(less func(a, b UserProfilePhotoClass) bool) UserProfilePhotoClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UserProfilePhotoClass.
func (s UserProfilePhotoClassArray) Retain(keep func(x UserProfilePhotoClass) bool) UserProfilePhotoClassArray {
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
func (s UserProfilePhotoClassArray) First() (v UserProfilePhotoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UserProfilePhotoClassArray) Last() (v UserProfilePhotoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UserProfilePhotoClassArray) PopFirst() (v UserProfilePhotoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UserProfilePhotoClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UserProfilePhotoClassArray) Pop() (v UserProfilePhotoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsUserProfilePhoto returns copy with only UserProfilePhoto constructors.
func (s UserProfilePhotoClassArray) AsUserProfilePhoto() (to UserProfilePhotoArray) {
	for _, elem := range s {
		value, ok := elem.(*UserProfilePhoto)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s UserProfilePhotoClassArray) AppendOnlyNotEmpty(to []*UserProfilePhoto) []*UserProfilePhoto {
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
func (s UserProfilePhotoClassArray) AsNotEmpty() (to []*UserProfilePhoto) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s UserProfilePhotoClassArray) FirstAsNotEmpty() (v *UserProfilePhoto, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s UserProfilePhotoClassArray) LastAsNotEmpty() (v *UserProfilePhoto, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *UserProfilePhotoClassArray) PopFirstAsNotEmpty() (v *UserProfilePhoto, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *UserProfilePhotoClassArray) PopAsNotEmpty() (v *UserProfilePhoto, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// UserProfilePhotoArray is adapter for slice of UserProfilePhoto.
type UserProfilePhotoArray []UserProfilePhoto

// Sort sorts slice of UserProfilePhoto.
func (s UserProfilePhotoArray) Sort(less func(a, b UserProfilePhoto) bool) UserProfilePhotoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UserProfilePhoto.
func (s UserProfilePhotoArray) SortStable(less func(a, b UserProfilePhoto) bool) UserProfilePhotoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UserProfilePhoto.
func (s UserProfilePhotoArray) Retain(keep func(x UserProfilePhoto) bool) UserProfilePhotoArray {
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
func (s UserProfilePhotoArray) First() (v UserProfilePhoto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UserProfilePhotoArray) Last() (v UserProfilePhoto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UserProfilePhotoArray) PopFirst() (v UserProfilePhoto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UserProfilePhoto
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UserProfilePhotoArray) Pop() (v UserProfilePhoto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
