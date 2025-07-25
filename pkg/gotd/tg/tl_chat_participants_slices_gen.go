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

// ChatParticipantsClassArray is adapter for slice of ChatParticipantsClass.
type ChatParticipantsClassArray []ChatParticipantsClass

// Sort sorts slice of ChatParticipantsClass.
func (s ChatParticipantsClassArray) Sort(less func(a, b ChatParticipantsClass) bool) ChatParticipantsClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatParticipantsClass.
func (s ChatParticipantsClassArray) SortStable(less func(a, b ChatParticipantsClass) bool) ChatParticipantsClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatParticipantsClass.
func (s ChatParticipantsClassArray) Retain(keep func(x ChatParticipantsClass) bool) ChatParticipantsClassArray {
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
func (s ChatParticipantsClassArray) First() (v ChatParticipantsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatParticipantsClassArray) Last() (v ChatParticipantsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatParticipantsClassArray) PopFirst() (v ChatParticipantsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatParticipantsClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatParticipantsClassArray) Pop() (v ChatParticipantsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsChatParticipantsForbidden returns copy with only ChatParticipantsForbidden constructors.
func (s ChatParticipantsClassArray) AsChatParticipantsForbidden() (to ChatParticipantsForbiddenArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatParticipantsForbidden)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsChatParticipants returns copy with only ChatParticipants constructors.
func (s ChatParticipantsClassArray) AsChatParticipants() (to ChatParticipantsArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatParticipants)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotForbidden appends only NotForbidden constructors to
// given slice.
func (s ChatParticipantsClassArray) AppendOnlyNotForbidden(to []*ChatParticipants) []*ChatParticipants {
	for _, elem := range s {
		value, ok := elem.AsNotForbidden()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotForbidden returns copy with only NotForbidden constructors.
func (s ChatParticipantsClassArray) AsNotForbidden() (to []*ChatParticipants) {
	return s.AppendOnlyNotForbidden(to)
}

// FirstAsNotForbidden returns first element of slice (if exists).
func (s ChatParticipantsClassArray) FirstAsNotForbidden() (v *ChatParticipants, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotForbidden()
}

// LastAsNotForbidden returns last element of slice (if exists).
func (s ChatParticipantsClassArray) LastAsNotForbidden() (v *ChatParticipants, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotForbidden()
}

// PopFirstAsNotForbidden returns element of slice (if exists).
func (s *ChatParticipantsClassArray) PopFirstAsNotForbidden() (v *ChatParticipants, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotForbidden()
}

// PopAsNotForbidden returns element of slice (if exists).
func (s *ChatParticipantsClassArray) PopAsNotForbidden() (v *ChatParticipants, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotForbidden()
}

// ChatParticipantsForbiddenArray is adapter for slice of ChatParticipantsForbidden.
type ChatParticipantsForbiddenArray []ChatParticipantsForbidden

// Sort sorts slice of ChatParticipantsForbidden.
func (s ChatParticipantsForbiddenArray) Sort(less func(a, b ChatParticipantsForbidden) bool) ChatParticipantsForbiddenArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatParticipantsForbidden.
func (s ChatParticipantsForbiddenArray) SortStable(less func(a, b ChatParticipantsForbidden) bool) ChatParticipantsForbiddenArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatParticipantsForbidden.
func (s ChatParticipantsForbiddenArray) Retain(keep func(x ChatParticipantsForbidden) bool) ChatParticipantsForbiddenArray {
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
func (s ChatParticipantsForbiddenArray) First() (v ChatParticipantsForbidden, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatParticipantsForbiddenArray) Last() (v ChatParticipantsForbidden, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatParticipantsForbiddenArray) PopFirst() (v ChatParticipantsForbidden, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatParticipantsForbidden
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatParticipantsForbiddenArray) Pop() (v ChatParticipantsForbidden, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ChatParticipantsArray is adapter for slice of ChatParticipants.
type ChatParticipantsArray []ChatParticipants

// Sort sorts slice of ChatParticipants.
func (s ChatParticipantsArray) Sort(less func(a, b ChatParticipants) bool) ChatParticipantsArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatParticipants.
func (s ChatParticipantsArray) SortStable(less func(a, b ChatParticipants) bool) ChatParticipantsArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatParticipants.
func (s ChatParticipantsArray) Retain(keep func(x ChatParticipants) bool) ChatParticipantsArray {
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
func (s ChatParticipantsArray) First() (v ChatParticipants, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatParticipantsArray) Last() (v ChatParticipants, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatParticipantsArray) PopFirst() (v ChatParticipants, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatParticipants
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatParticipantsArray) Pop() (v ChatParticipants, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
