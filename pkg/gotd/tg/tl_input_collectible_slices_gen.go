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

// InputCollectibleClassArray is adapter for slice of InputCollectibleClass.
type InputCollectibleClassArray []InputCollectibleClass

// Sort sorts slice of InputCollectibleClass.
func (s InputCollectibleClassArray) Sort(less func(a, b InputCollectibleClass) bool) InputCollectibleClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputCollectibleClass.
func (s InputCollectibleClassArray) SortStable(less func(a, b InputCollectibleClass) bool) InputCollectibleClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputCollectibleClass.
func (s InputCollectibleClassArray) Retain(keep func(x InputCollectibleClass) bool) InputCollectibleClassArray {
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
func (s InputCollectibleClassArray) First() (v InputCollectibleClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputCollectibleClassArray) Last() (v InputCollectibleClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputCollectibleClassArray) PopFirst() (v InputCollectibleClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputCollectibleClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputCollectibleClassArray) Pop() (v InputCollectibleClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputCollectibleUsername returns copy with only InputCollectibleUsername constructors.
func (s InputCollectibleClassArray) AsInputCollectibleUsername() (to InputCollectibleUsernameArray) {
	for _, elem := range s {
		value, ok := elem.(*InputCollectibleUsername)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputCollectiblePhone returns copy with only InputCollectiblePhone constructors.
func (s InputCollectibleClassArray) AsInputCollectiblePhone() (to InputCollectiblePhoneArray) {
	for _, elem := range s {
		value, ok := elem.(*InputCollectiblePhone)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputCollectibleUsernameArray is adapter for slice of InputCollectibleUsername.
type InputCollectibleUsernameArray []InputCollectibleUsername

// Sort sorts slice of InputCollectibleUsername.
func (s InputCollectibleUsernameArray) Sort(less func(a, b InputCollectibleUsername) bool) InputCollectibleUsernameArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputCollectibleUsername.
func (s InputCollectibleUsernameArray) SortStable(less func(a, b InputCollectibleUsername) bool) InputCollectibleUsernameArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputCollectibleUsername.
func (s InputCollectibleUsernameArray) Retain(keep func(x InputCollectibleUsername) bool) InputCollectibleUsernameArray {
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
func (s InputCollectibleUsernameArray) First() (v InputCollectibleUsername, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputCollectibleUsernameArray) Last() (v InputCollectibleUsername, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputCollectibleUsernameArray) PopFirst() (v InputCollectibleUsername, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputCollectibleUsername
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputCollectibleUsernameArray) Pop() (v InputCollectibleUsername, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputCollectiblePhoneArray is adapter for slice of InputCollectiblePhone.
type InputCollectiblePhoneArray []InputCollectiblePhone

// Sort sorts slice of InputCollectiblePhone.
func (s InputCollectiblePhoneArray) Sort(less func(a, b InputCollectiblePhone) bool) InputCollectiblePhoneArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputCollectiblePhone.
func (s InputCollectiblePhoneArray) SortStable(less func(a, b InputCollectiblePhone) bool) InputCollectiblePhoneArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputCollectiblePhone.
func (s InputCollectiblePhoneArray) Retain(keep func(x InputCollectiblePhone) bool) InputCollectiblePhoneArray {
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
func (s InputCollectiblePhoneArray) First() (v InputCollectiblePhone, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputCollectiblePhoneArray) Last() (v InputCollectiblePhone, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputCollectiblePhoneArray) PopFirst() (v InputCollectiblePhone, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputCollectiblePhone
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputCollectiblePhoneArray) Pop() (v InputCollectiblePhone, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
