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

// WallPaperClassArray is adapter for slice of WallPaperClass.
type WallPaperClassArray []WallPaperClass

// Sort sorts slice of WallPaperClass.
func (s WallPaperClassArray) Sort(less func(a, b WallPaperClass) bool) WallPaperClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of WallPaperClass.
func (s WallPaperClassArray) SortStable(less func(a, b WallPaperClass) bool) WallPaperClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of WallPaperClass.
func (s WallPaperClassArray) Retain(keep func(x WallPaperClass) bool) WallPaperClassArray {
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
func (s WallPaperClassArray) First() (v WallPaperClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s WallPaperClassArray) Last() (v WallPaperClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *WallPaperClassArray) PopFirst() (v WallPaperClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero WallPaperClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *WallPaperClassArray) Pop() (v WallPaperClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of WallPaperClass by ID.
func (s WallPaperClassArray) SortByID() WallPaperClassArray {
	return s.Sort(func(a, b WallPaperClass) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of WallPaperClass by ID.
func (s WallPaperClassArray) SortStableByID() WallPaperClassArray {
	return s.SortStable(func(a, b WallPaperClass) bool {
		return a.GetID() < b.GetID()
	})
}

// FillWallPaperMap fills only WallPaper constructors to given map.
func (s WallPaperClassArray) FillWallPaperMap(to map[int64]*WallPaper) {
	for _, elem := range s {
		value, ok := elem.(*WallPaper)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// WallPaperToMap collects only WallPaper constructors to map.
func (s WallPaperClassArray) WallPaperToMap() map[int64]*WallPaper {
	r := make(map[int64]*WallPaper, len(s))
	s.FillWallPaperMap(r)
	return r
}

// AsWallPaper returns copy with only WallPaper constructors.
func (s WallPaperClassArray) AsWallPaper() (to WallPaperArray) {
	for _, elem := range s {
		value, ok := elem.(*WallPaper)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillWallPaperNoFileMap fills only WallPaperNoFile constructors to given map.
func (s WallPaperClassArray) FillWallPaperNoFileMap(to map[int64]*WallPaperNoFile) {
	for _, elem := range s {
		value, ok := elem.(*WallPaperNoFile)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// WallPaperNoFileToMap collects only WallPaperNoFile constructors to map.
func (s WallPaperClassArray) WallPaperNoFileToMap() map[int64]*WallPaperNoFile {
	r := make(map[int64]*WallPaperNoFile, len(s))
	s.FillWallPaperNoFileMap(r)
	return r
}

// AsWallPaperNoFile returns copy with only WallPaperNoFile constructors.
func (s WallPaperClassArray) AsWallPaperNoFile() (to WallPaperNoFileArray) {
	for _, elem := range s {
		value, ok := elem.(*WallPaperNoFile)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// WallPaperArray is adapter for slice of WallPaper.
type WallPaperArray []WallPaper

// Sort sorts slice of WallPaper.
func (s WallPaperArray) Sort(less func(a, b WallPaper) bool) WallPaperArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of WallPaper.
func (s WallPaperArray) SortStable(less func(a, b WallPaper) bool) WallPaperArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of WallPaper.
func (s WallPaperArray) Retain(keep func(x WallPaper) bool) WallPaperArray {
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
func (s WallPaperArray) First() (v WallPaper, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s WallPaperArray) Last() (v WallPaper, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *WallPaperArray) PopFirst() (v WallPaper, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero WallPaper
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *WallPaperArray) Pop() (v WallPaper, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of WallPaper by ID.
func (s WallPaperArray) SortByID() WallPaperArray {
	return s.Sort(func(a, b WallPaper) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of WallPaper by ID.
func (s WallPaperArray) SortStableByID() WallPaperArray {
	return s.SortStable(func(a, b WallPaper) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s WallPaperArray) FillMap(to map[int64]WallPaper) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s WallPaperArray) ToMap() map[int64]WallPaper {
	r := make(map[int64]WallPaper, len(s))
	s.FillMap(r)
	return r
}

// WallPaperNoFileArray is adapter for slice of WallPaperNoFile.
type WallPaperNoFileArray []WallPaperNoFile

// Sort sorts slice of WallPaperNoFile.
func (s WallPaperNoFileArray) Sort(less func(a, b WallPaperNoFile) bool) WallPaperNoFileArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of WallPaperNoFile.
func (s WallPaperNoFileArray) SortStable(less func(a, b WallPaperNoFile) bool) WallPaperNoFileArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of WallPaperNoFile.
func (s WallPaperNoFileArray) Retain(keep func(x WallPaperNoFile) bool) WallPaperNoFileArray {
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
func (s WallPaperNoFileArray) First() (v WallPaperNoFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s WallPaperNoFileArray) Last() (v WallPaperNoFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *WallPaperNoFileArray) PopFirst() (v WallPaperNoFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero WallPaperNoFile
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *WallPaperNoFileArray) Pop() (v WallPaperNoFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of WallPaperNoFile by ID.
func (s WallPaperNoFileArray) SortByID() WallPaperNoFileArray {
	return s.Sort(func(a, b WallPaperNoFile) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of WallPaperNoFile by ID.
func (s WallPaperNoFileArray) SortStableByID() WallPaperNoFileArray {
	return s.SortStable(func(a, b WallPaperNoFile) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s WallPaperNoFileArray) FillMap(to map[int64]WallPaperNoFile) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s WallPaperNoFileArray) ToMap() map[int64]WallPaperNoFile {
	r := make(map[int64]WallPaperNoFile, len(s))
	s.FillMap(r)
	return r
}
