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

// MessagesRecentStickersClassArray is adapter for slice of MessagesRecentStickersClass.
type MessagesRecentStickersClassArray []MessagesRecentStickersClass

// Sort sorts slice of MessagesRecentStickersClass.
func (s MessagesRecentStickersClassArray) Sort(less func(a, b MessagesRecentStickersClass) bool) MessagesRecentStickersClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesRecentStickersClass.
func (s MessagesRecentStickersClassArray) SortStable(less func(a, b MessagesRecentStickersClass) bool) MessagesRecentStickersClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesRecentStickersClass.
func (s MessagesRecentStickersClassArray) Retain(keep func(x MessagesRecentStickersClass) bool) MessagesRecentStickersClassArray {
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
func (s MessagesRecentStickersClassArray) First() (v MessagesRecentStickersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesRecentStickersClassArray) Last() (v MessagesRecentStickersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesRecentStickersClassArray) PopFirst() (v MessagesRecentStickersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesRecentStickersClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesRecentStickersClassArray) Pop() (v MessagesRecentStickersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessagesRecentStickers returns copy with only MessagesRecentStickers constructors.
func (s MessagesRecentStickersClassArray) AsMessagesRecentStickers() (to MessagesRecentStickersArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesRecentStickers)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s MessagesRecentStickersClassArray) AppendOnlyModified(to []*MessagesRecentStickers) []*MessagesRecentStickers {
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
func (s MessagesRecentStickersClassArray) AsModified() (to []*MessagesRecentStickers) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s MessagesRecentStickersClassArray) FirstAsModified() (v *MessagesRecentStickers, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s MessagesRecentStickersClassArray) LastAsModified() (v *MessagesRecentStickers, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *MessagesRecentStickersClassArray) PopFirstAsModified() (v *MessagesRecentStickers, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *MessagesRecentStickersClassArray) PopAsModified() (v *MessagesRecentStickers, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// MessagesRecentStickersArray is adapter for slice of MessagesRecentStickers.
type MessagesRecentStickersArray []MessagesRecentStickers

// Sort sorts slice of MessagesRecentStickers.
func (s MessagesRecentStickersArray) Sort(less func(a, b MessagesRecentStickers) bool) MessagesRecentStickersArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesRecentStickers.
func (s MessagesRecentStickersArray) SortStable(less func(a, b MessagesRecentStickers) bool) MessagesRecentStickersArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesRecentStickers.
func (s MessagesRecentStickersArray) Retain(keep func(x MessagesRecentStickers) bool) MessagesRecentStickersArray {
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
func (s MessagesRecentStickersArray) First() (v MessagesRecentStickers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesRecentStickersArray) Last() (v MessagesRecentStickers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesRecentStickersArray) PopFirst() (v MessagesRecentStickers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesRecentStickers
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesRecentStickersArray) Pop() (v MessagesRecentStickers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
