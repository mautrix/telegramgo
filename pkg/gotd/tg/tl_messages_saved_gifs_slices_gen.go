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

// MessagesSavedGifsClassArray is adapter for slice of MessagesSavedGifsClass.
type MessagesSavedGifsClassArray []MessagesSavedGifsClass

// Sort sorts slice of MessagesSavedGifsClass.
func (s MessagesSavedGifsClassArray) Sort(less func(a, b MessagesSavedGifsClass) bool) MessagesSavedGifsClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesSavedGifsClass.
func (s MessagesSavedGifsClassArray) SortStable(less func(a, b MessagesSavedGifsClass) bool) MessagesSavedGifsClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesSavedGifsClass.
func (s MessagesSavedGifsClassArray) Retain(keep func(x MessagesSavedGifsClass) bool) MessagesSavedGifsClassArray {
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
func (s MessagesSavedGifsClassArray) First() (v MessagesSavedGifsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesSavedGifsClassArray) Last() (v MessagesSavedGifsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesSavedGifsClassArray) PopFirst() (v MessagesSavedGifsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesSavedGifsClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesSavedGifsClassArray) Pop() (v MessagesSavedGifsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessagesSavedGifs returns copy with only MessagesSavedGifs constructors.
func (s MessagesSavedGifsClassArray) AsMessagesSavedGifs() (to MessagesSavedGifsArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesSavedGifs)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s MessagesSavedGifsClassArray) AppendOnlyModified(to []*MessagesSavedGifs) []*MessagesSavedGifs {
	for _, elem := range s {
		value, ok := elem.AsModified()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsModified returns copy with only Modified constructors.
func (s MessagesSavedGifsClassArray) AsModified() (to []*MessagesSavedGifs) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s MessagesSavedGifsClassArray) FirstAsModified() (v *MessagesSavedGifs, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s MessagesSavedGifsClassArray) LastAsModified() (v *MessagesSavedGifs, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *MessagesSavedGifsClassArray) PopFirstAsModified() (v *MessagesSavedGifs, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *MessagesSavedGifsClassArray) PopAsModified() (v *MessagesSavedGifs, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// MessagesSavedGifsArray is adapter for slice of MessagesSavedGifs.
type MessagesSavedGifsArray []MessagesSavedGifs

// Sort sorts slice of MessagesSavedGifs.
func (s MessagesSavedGifsArray) Sort(less func(a, b MessagesSavedGifs) bool) MessagesSavedGifsArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesSavedGifs.
func (s MessagesSavedGifsArray) SortStable(less func(a, b MessagesSavedGifs) bool) MessagesSavedGifsArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesSavedGifs.
func (s MessagesSavedGifsArray) Retain(keep func(x MessagesSavedGifs) bool) MessagesSavedGifsArray {
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
func (s MessagesSavedGifsArray) First() (v MessagesSavedGifs, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesSavedGifsArray) Last() (v MessagesSavedGifs, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesSavedGifsArray) PopFirst() (v MessagesSavedGifs, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesSavedGifs
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesSavedGifsArray) Pop() (v MessagesSavedGifs, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
