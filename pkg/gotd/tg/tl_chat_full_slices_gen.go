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

// ChatFullClassArray is adapter for slice of ChatFullClass.
type ChatFullClassArray []ChatFullClass

// Sort sorts slice of ChatFullClass.
func (s ChatFullClassArray) Sort(less func(a, b ChatFullClass) bool) ChatFullClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatFullClass.
func (s ChatFullClassArray) SortStable(less func(a, b ChatFullClass) bool) ChatFullClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatFullClass.
func (s ChatFullClassArray) Retain(keep func(x ChatFullClass) bool) ChatFullClassArray {
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
func (s ChatFullClassArray) First() (v ChatFullClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatFullClassArray) Last() (v ChatFullClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatFullClassArray) PopFirst() (v ChatFullClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatFullClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatFullClassArray) Pop() (v ChatFullClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of ChatFullClass by ID.
func (s ChatFullClassArray) SortByID() ChatFullClassArray {
	return s.Sort(func(a, b ChatFullClass) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of ChatFullClass by ID.
func (s ChatFullClassArray) SortStableByID() ChatFullClassArray {
	return s.SortStable(func(a, b ChatFullClass) bool {
		return a.GetID() < b.GetID()
	})
}

// FillChatFullMap fills only ChatFull constructors to given map.
func (s ChatFullClassArray) FillChatFullMap(to map[int64]*ChatFull) {
	for _, elem := range s {
		value, ok := elem.(*ChatFull)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// ChatFullToMap collects only ChatFull constructors to map.
func (s ChatFullClassArray) ChatFullToMap() map[int64]*ChatFull {
	r := make(map[int64]*ChatFull, len(s))
	s.FillChatFullMap(r)
	return r
}

// AsChatFull returns copy with only ChatFull constructors.
func (s ChatFullClassArray) AsChatFull() (to ChatFullArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatFull)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillChannelFullMap fills only ChannelFull constructors to given map.
func (s ChatFullClassArray) FillChannelFullMap(to map[int64]*ChannelFull) {
	for _, elem := range s {
		value, ok := elem.(*ChannelFull)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// ChannelFullToMap collects only ChannelFull constructors to map.
func (s ChatFullClassArray) ChannelFullToMap() map[int64]*ChannelFull {
	r := make(map[int64]*ChannelFull, len(s))
	s.FillChannelFullMap(r)
	return r
}

// AsChannelFull returns copy with only ChannelFull constructors.
func (s ChatFullClassArray) AsChannelFull() (to ChannelFullArray) {
	for _, elem := range s {
		value, ok := elem.(*ChannelFull)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// ChatFullArray is adapter for slice of ChatFull.
type ChatFullArray []ChatFull

// Sort sorts slice of ChatFull.
func (s ChatFullArray) Sort(less func(a, b ChatFull) bool) ChatFullArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatFull.
func (s ChatFullArray) SortStable(less func(a, b ChatFull) bool) ChatFullArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatFull.
func (s ChatFullArray) Retain(keep func(x ChatFull) bool) ChatFullArray {
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
func (s ChatFullArray) First() (v ChatFull, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatFullArray) Last() (v ChatFull, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatFullArray) PopFirst() (v ChatFull, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatFull
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatFullArray) Pop() (v ChatFull, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of ChatFull by ID.
func (s ChatFullArray) SortByID() ChatFullArray {
	return s.Sort(func(a, b ChatFull) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of ChatFull by ID.
func (s ChatFullArray) SortStableByID() ChatFullArray {
	return s.SortStable(func(a, b ChatFull) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s ChatFullArray) FillMap(to map[int64]ChatFull) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s ChatFullArray) ToMap() map[int64]ChatFull {
	r := make(map[int64]ChatFull, len(s))
	s.FillMap(r)
	return r
}

// ChannelFullArray is adapter for slice of ChannelFull.
type ChannelFullArray []ChannelFull

// Sort sorts slice of ChannelFull.
func (s ChannelFullArray) Sort(less func(a, b ChannelFull) bool) ChannelFullArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChannelFull.
func (s ChannelFullArray) SortStable(less func(a, b ChannelFull) bool) ChannelFullArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChannelFull.
func (s ChannelFullArray) Retain(keep func(x ChannelFull) bool) ChannelFullArray {
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
func (s ChannelFullArray) First() (v ChannelFull, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChannelFullArray) Last() (v ChannelFull, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChannelFullArray) PopFirst() (v ChannelFull, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChannelFull
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChannelFullArray) Pop() (v ChannelFull, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of ChannelFull by ID.
func (s ChannelFullArray) SortByID() ChannelFullArray {
	return s.Sort(func(a, b ChannelFull) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of ChannelFull by ID.
func (s ChannelFullArray) SortStableByID() ChannelFullArray {
	return s.SortStable(func(a, b ChannelFull) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s ChannelFullArray) FillMap(to map[int64]ChannelFull) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s ChannelFullArray) ToMap() map[int64]ChannelFull {
	r := make(map[int64]ChannelFull, len(s))
	s.FillMap(r)
	return r
}
