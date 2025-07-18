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

// DialogPeerClassArray is adapter for slice of DialogPeerClass.
type DialogPeerClassArray []DialogPeerClass

// Sort sorts slice of DialogPeerClass.
func (s DialogPeerClassArray) Sort(less func(a, b DialogPeerClass) bool) DialogPeerClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DialogPeerClass.
func (s DialogPeerClassArray) SortStable(less func(a, b DialogPeerClass) bool) DialogPeerClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DialogPeerClass.
func (s DialogPeerClassArray) Retain(keep func(x DialogPeerClass) bool) DialogPeerClassArray {
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
func (s DialogPeerClassArray) First() (v DialogPeerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DialogPeerClassArray) Last() (v DialogPeerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DialogPeerClassArray) PopFirst() (v DialogPeerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DialogPeerClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DialogPeerClassArray) Pop() (v DialogPeerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsDialogPeer returns copy with only DialogPeer constructors.
func (s DialogPeerClassArray) AsDialogPeer() (to DialogPeerArray) {
	for _, elem := range s {
		value, ok := elem.(*DialogPeer)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsDialogPeerFolder returns copy with only DialogPeerFolder constructors.
func (s DialogPeerClassArray) AsDialogPeerFolder() (to DialogPeerFolderArray) {
	for _, elem := range s {
		value, ok := elem.(*DialogPeerFolder)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// DialogPeerArray is adapter for slice of DialogPeer.
type DialogPeerArray []DialogPeer

// Sort sorts slice of DialogPeer.
func (s DialogPeerArray) Sort(less func(a, b DialogPeer) bool) DialogPeerArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DialogPeer.
func (s DialogPeerArray) SortStable(less func(a, b DialogPeer) bool) DialogPeerArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DialogPeer.
func (s DialogPeerArray) Retain(keep func(x DialogPeer) bool) DialogPeerArray {
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
func (s DialogPeerArray) First() (v DialogPeer, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DialogPeerArray) Last() (v DialogPeer, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DialogPeerArray) PopFirst() (v DialogPeer, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DialogPeer
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DialogPeerArray) Pop() (v DialogPeer, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// DialogPeerFolderArray is adapter for slice of DialogPeerFolder.
type DialogPeerFolderArray []DialogPeerFolder

// Sort sorts slice of DialogPeerFolder.
func (s DialogPeerFolderArray) Sort(less func(a, b DialogPeerFolder) bool) DialogPeerFolderArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DialogPeerFolder.
func (s DialogPeerFolderArray) SortStable(less func(a, b DialogPeerFolder) bool) DialogPeerFolderArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DialogPeerFolder.
func (s DialogPeerFolderArray) Retain(keep func(x DialogPeerFolder) bool) DialogPeerFolderArray {
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
func (s DialogPeerFolderArray) First() (v DialogPeerFolder, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DialogPeerFolderArray) Last() (v DialogPeerFolder, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DialogPeerFolderArray) PopFirst() (v DialogPeerFolder, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DialogPeerFolder
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DialogPeerFolderArray) Pop() (v DialogPeerFolder, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
