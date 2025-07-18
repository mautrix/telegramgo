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

// InputMessageClassArray is adapter for slice of InputMessageClass.
type InputMessageClassArray []InputMessageClass

// Sort sorts slice of InputMessageClass.
func (s InputMessageClassArray) Sort(less func(a, b InputMessageClass) bool) InputMessageClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMessageClass.
func (s InputMessageClassArray) SortStable(less func(a, b InputMessageClass) bool) InputMessageClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMessageClass.
func (s InputMessageClassArray) Retain(keep func(x InputMessageClass) bool) InputMessageClassArray {
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
func (s InputMessageClassArray) First() (v InputMessageClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMessageClassArray) Last() (v InputMessageClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMessageClassArray) PopFirst() (v InputMessageClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMessageClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMessageClassArray) Pop() (v InputMessageClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputMessageID returns copy with only InputMessageID constructors.
func (s InputMessageClassArray) AsInputMessageID() (to InputMessageIDArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMessageID)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMessageReplyTo returns copy with only InputMessageReplyTo constructors.
func (s InputMessageClassArray) AsInputMessageReplyTo() (to InputMessageReplyToArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMessageReplyTo)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMessageCallbackQuery returns copy with only InputMessageCallbackQuery constructors.
func (s InputMessageClassArray) AsInputMessageCallbackQuery() (to InputMessageCallbackQueryArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMessageCallbackQuery)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputMessageIDArray is adapter for slice of InputMessageID.
type InputMessageIDArray []InputMessageID

// Sort sorts slice of InputMessageID.
func (s InputMessageIDArray) Sort(less func(a, b InputMessageID) bool) InputMessageIDArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMessageID.
func (s InputMessageIDArray) SortStable(less func(a, b InputMessageID) bool) InputMessageIDArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMessageID.
func (s InputMessageIDArray) Retain(keep func(x InputMessageID) bool) InputMessageIDArray {
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
func (s InputMessageIDArray) First() (v InputMessageID, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMessageIDArray) Last() (v InputMessageID, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMessageIDArray) PopFirst() (v InputMessageID, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMessageID
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMessageIDArray) Pop() (v InputMessageID, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputMessageID by ID.
func (s InputMessageIDArray) SortByID() InputMessageIDArray {
	return s.Sort(func(a, b InputMessageID) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputMessageID by ID.
func (s InputMessageIDArray) SortStableByID() InputMessageIDArray {
	return s.SortStable(func(a, b InputMessageID) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputMessageIDArray) FillMap(to map[int]InputMessageID) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputMessageIDArray) ToMap() map[int]InputMessageID {
	r := make(map[int]InputMessageID, len(s))
	s.FillMap(r)
	return r
}

// InputMessageReplyToArray is adapter for slice of InputMessageReplyTo.
type InputMessageReplyToArray []InputMessageReplyTo

// Sort sorts slice of InputMessageReplyTo.
func (s InputMessageReplyToArray) Sort(less func(a, b InputMessageReplyTo) bool) InputMessageReplyToArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMessageReplyTo.
func (s InputMessageReplyToArray) SortStable(less func(a, b InputMessageReplyTo) bool) InputMessageReplyToArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMessageReplyTo.
func (s InputMessageReplyToArray) Retain(keep func(x InputMessageReplyTo) bool) InputMessageReplyToArray {
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
func (s InputMessageReplyToArray) First() (v InputMessageReplyTo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMessageReplyToArray) Last() (v InputMessageReplyTo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMessageReplyToArray) PopFirst() (v InputMessageReplyTo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMessageReplyTo
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMessageReplyToArray) Pop() (v InputMessageReplyTo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputMessageReplyTo by ID.
func (s InputMessageReplyToArray) SortByID() InputMessageReplyToArray {
	return s.Sort(func(a, b InputMessageReplyTo) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputMessageReplyTo by ID.
func (s InputMessageReplyToArray) SortStableByID() InputMessageReplyToArray {
	return s.SortStable(func(a, b InputMessageReplyTo) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputMessageReplyToArray) FillMap(to map[int]InputMessageReplyTo) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputMessageReplyToArray) ToMap() map[int]InputMessageReplyTo {
	r := make(map[int]InputMessageReplyTo, len(s))
	s.FillMap(r)
	return r
}

// InputMessageCallbackQueryArray is adapter for slice of InputMessageCallbackQuery.
type InputMessageCallbackQueryArray []InputMessageCallbackQuery

// Sort sorts slice of InputMessageCallbackQuery.
func (s InputMessageCallbackQueryArray) Sort(less func(a, b InputMessageCallbackQuery) bool) InputMessageCallbackQueryArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMessageCallbackQuery.
func (s InputMessageCallbackQueryArray) SortStable(less func(a, b InputMessageCallbackQuery) bool) InputMessageCallbackQueryArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMessageCallbackQuery.
func (s InputMessageCallbackQueryArray) Retain(keep func(x InputMessageCallbackQuery) bool) InputMessageCallbackQueryArray {
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
func (s InputMessageCallbackQueryArray) First() (v InputMessageCallbackQuery, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMessageCallbackQueryArray) Last() (v InputMessageCallbackQuery, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMessageCallbackQueryArray) PopFirst() (v InputMessageCallbackQuery, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMessageCallbackQuery
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMessageCallbackQueryArray) Pop() (v InputMessageCallbackQuery, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputMessageCallbackQuery by ID.
func (s InputMessageCallbackQueryArray) SortByID() InputMessageCallbackQueryArray {
	return s.Sort(func(a, b InputMessageCallbackQuery) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputMessageCallbackQuery by ID.
func (s InputMessageCallbackQueryArray) SortStableByID() InputMessageCallbackQueryArray {
	return s.SortStable(func(a, b InputMessageCallbackQuery) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputMessageCallbackQueryArray) FillMap(to map[int]InputMessageCallbackQuery) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputMessageCallbackQueryArray) ToMap() map[int]InputMessageCallbackQuery {
	r := make(map[int]InputMessageCallbackQuery, len(s))
	s.FillMap(r)
	return r
}
