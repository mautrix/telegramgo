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

// ChatParticipantClassArray is adapter for slice of ChatParticipantClass.
type ChatParticipantClassArray []ChatParticipantClass

// Sort sorts slice of ChatParticipantClass.
func (s ChatParticipantClassArray) Sort(less func(a, b ChatParticipantClass) bool) ChatParticipantClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatParticipantClass.
func (s ChatParticipantClassArray) SortStable(less func(a, b ChatParticipantClass) bool) ChatParticipantClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatParticipantClass.
func (s ChatParticipantClassArray) Retain(keep func(x ChatParticipantClass) bool) ChatParticipantClassArray {
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
func (s ChatParticipantClassArray) First() (v ChatParticipantClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatParticipantClassArray) Last() (v ChatParticipantClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatParticipantClassArray) PopFirst() (v ChatParticipantClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatParticipantClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatParticipantClassArray) Pop() (v ChatParticipantClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsChatParticipant returns copy with only ChatParticipant constructors.
func (s ChatParticipantClassArray) AsChatParticipant() (to ChatParticipantArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatParticipant)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsChatParticipantCreator returns copy with only ChatParticipantCreator constructors.
func (s ChatParticipantClassArray) AsChatParticipantCreator() (to ChatParticipantCreatorArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatParticipantCreator)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsChatParticipantAdmin returns copy with only ChatParticipantAdmin constructors.
func (s ChatParticipantClassArray) AsChatParticipantAdmin() (to ChatParticipantAdminArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatParticipantAdmin)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// ChatParticipantArray is adapter for slice of ChatParticipant.
type ChatParticipantArray []ChatParticipant

// Sort sorts slice of ChatParticipant.
func (s ChatParticipantArray) Sort(less func(a, b ChatParticipant) bool) ChatParticipantArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatParticipant.
func (s ChatParticipantArray) SortStable(less func(a, b ChatParticipant) bool) ChatParticipantArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatParticipant.
func (s ChatParticipantArray) Retain(keep func(x ChatParticipant) bool) ChatParticipantArray {
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
func (s ChatParticipantArray) First() (v ChatParticipant, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatParticipantArray) Last() (v ChatParticipant, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatParticipantArray) PopFirst() (v ChatParticipant, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatParticipant
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatParticipantArray) Pop() (v ChatParticipant, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of ChatParticipant by Date.
func (s ChatParticipantArray) SortByDate() ChatParticipantArray {
	return s.Sort(func(a, b ChatParticipant) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of ChatParticipant by Date.
func (s ChatParticipantArray) SortStableByDate() ChatParticipantArray {
	return s.SortStable(func(a, b ChatParticipant) bool {
		return a.GetDate() < b.GetDate()
	})
}

// ChatParticipantCreatorArray is adapter for slice of ChatParticipantCreator.
type ChatParticipantCreatorArray []ChatParticipantCreator

// Sort sorts slice of ChatParticipantCreator.
func (s ChatParticipantCreatorArray) Sort(less func(a, b ChatParticipantCreator) bool) ChatParticipantCreatorArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatParticipantCreator.
func (s ChatParticipantCreatorArray) SortStable(less func(a, b ChatParticipantCreator) bool) ChatParticipantCreatorArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatParticipantCreator.
func (s ChatParticipantCreatorArray) Retain(keep func(x ChatParticipantCreator) bool) ChatParticipantCreatorArray {
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
func (s ChatParticipantCreatorArray) First() (v ChatParticipantCreator, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatParticipantCreatorArray) Last() (v ChatParticipantCreator, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatParticipantCreatorArray) PopFirst() (v ChatParticipantCreator, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatParticipantCreator
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatParticipantCreatorArray) Pop() (v ChatParticipantCreator, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ChatParticipantAdminArray is adapter for slice of ChatParticipantAdmin.
type ChatParticipantAdminArray []ChatParticipantAdmin

// Sort sorts slice of ChatParticipantAdmin.
func (s ChatParticipantAdminArray) Sort(less func(a, b ChatParticipantAdmin) bool) ChatParticipantAdminArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatParticipantAdmin.
func (s ChatParticipantAdminArray) SortStable(less func(a, b ChatParticipantAdmin) bool) ChatParticipantAdminArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatParticipantAdmin.
func (s ChatParticipantAdminArray) Retain(keep func(x ChatParticipantAdmin) bool) ChatParticipantAdminArray {
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
func (s ChatParticipantAdminArray) First() (v ChatParticipantAdmin, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatParticipantAdminArray) Last() (v ChatParticipantAdmin, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatParticipantAdminArray) PopFirst() (v ChatParticipantAdmin, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatParticipantAdmin
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatParticipantAdminArray) Pop() (v ChatParticipantAdmin, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of ChatParticipantAdmin by Date.
func (s ChatParticipantAdminArray) SortByDate() ChatParticipantAdminArray {
	return s.Sort(func(a, b ChatParticipantAdmin) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of ChatParticipantAdmin by Date.
func (s ChatParticipantAdminArray) SortStableByDate() ChatParticipantAdminArray {
	return s.SortStable(func(a, b ChatParticipantAdmin) bool {
		return a.GetDate() < b.GetDate()
	})
}
