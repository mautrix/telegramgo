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

// WebPageAttributeClassArray is adapter for slice of WebPageAttributeClass.
type WebPageAttributeClassArray []WebPageAttributeClass

// Sort sorts slice of WebPageAttributeClass.
func (s WebPageAttributeClassArray) Sort(less func(a, b WebPageAttributeClass) bool) WebPageAttributeClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of WebPageAttributeClass.
func (s WebPageAttributeClassArray) SortStable(less func(a, b WebPageAttributeClass) bool) WebPageAttributeClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of WebPageAttributeClass.
func (s WebPageAttributeClassArray) Retain(keep func(x WebPageAttributeClass) bool) WebPageAttributeClassArray {
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
func (s WebPageAttributeClassArray) First() (v WebPageAttributeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s WebPageAttributeClassArray) Last() (v WebPageAttributeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *WebPageAttributeClassArray) PopFirst() (v WebPageAttributeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero WebPageAttributeClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *WebPageAttributeClassArray) Pop() (v WebPageAttributeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsWebPageAttributeTheme returns copy with only WebPageAttributeTheme constructors.
func (s WebPageAttributeClassArray) AsWebPageAttributeTheme() (to WebPageAttributeThemeArray) {
	for _, elem := range s {
		value, ok := elem.(*WebPageAttributeTheme)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsWebPageAttributeStory returns copy with only WebPageAttributeStory constructors.
func (s WebPageAttributeClassArray) AsWebPageAttributeStory() (to WebPageAttributeStoryArray) {
	for _, elem := range s {
		value, ok := elem.(*WebPageAttributeStory)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsWebPageAttributeStickerSet returns copy with only WebPageAttributeStickerSet constructors.
func (s WebPageAttributeClassArray) AsWebPageAttributeStickerSet() (to WebPageAttributeStickerSetArray) {
	for _, elem := range s {
		value, ok := elem.(*WebPageAttributeStickerSet)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsWebPageAttributeUniqueStarGift returns copy with only WebPageAttributeUniqueStarGift constructors.
func (s WebPageAttributeClassArray) AsWebPageAttributeUniqueStarGift() (to WebPageAttributeUniqueStarGiftArray) {
	for _, elem := range s {
		value, ok := elem.(*WebPageAttributeUniqueStarGift)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// WebPageAttributeThemeArray is adapter for slice of WebPageAttributeTheme.
type WebPageAttributeThemeArray []WebPageAttributeTheme

// Sort sorts slice of WebPageAttributeTheme.
func (s WebPageAttributeThemeArray) Sort(less func(a, b WebPageAttributeTheme) bool) WebPageAttributeThemeArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of WebPageAttributeTheme.
func (s WebPageAttributeThemeArray) SortStable(less func(a, b WebPageAttributeTheme) bool) WebPageAttributeThemeArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of WebPageAttributeTheme.
func (s WebPageAttributeThemeArray) Retain(keep func(x WebPageAttributeTheme) bool) WebPageAttributeThemeArray {
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
func (s WebPageAttributeThemeArray) First() (v WebPageAttributeTheme, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s WebPageAttributeThemeArray) Last() (v WebPageAttributeTheme, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *WebPageAttributeThemeArray) PopFirst() (v WebPageAttributeTheme, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero WebPageAttributeTheme
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *WebPageAttributeThemeArray) Pop() (v WebPageAttributeTheme, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// WebPageAttributeStoryArray is adapter for slice of WebPageAttributeStory.
type WebPageAttributeStoryArray []WebPageAttributeStory

// Sort sorts slice of WebPageAttributeStory.
func (s WebPageAttributeStoryArray) Sort(less func(a, b WebPageAttributeStory) bool) WebPageAttributeStoryArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of WebPageAttributeStory.
func (s WebPageAttributeStoryArray) SortStable(less func(a, b WebPageAttributeStory) bool) WebPageAttributeStoryArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of WebPageAttributeStory.
func (s WebPageAttributeStoryArray) Retain(keep func(x WebPageAttributeStory) bool) WebPageAttributeStoryArray {
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
func (s WebPageAttributeStoryArray) First() (v WebPageAttributeStory, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s WebPageAttributeStoryArray) Last() (v WebPageAttributeStory, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *WebPageAttributeStoryArray) PopFirst() (v WebPageAttributeStory, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero WebPageAttributeStory
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *WebPageAttributeStoryArray) Pop() (v WebPageAttributeStory, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of WebPageAttributeStory by ID.
func (s WebPageAttributeStoryArray) SortByID() WebPageAttributeStoryArray {
	return s.Sort(func(a, b WebPageAttributeStory) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of WebPageAttributeStory by ID.
func (s WebPageAttributeStoryArray) SortStableByID() WebPageAttributeStoryArray {
	return s.SortStable(func(a, b WebPageAttributeStory) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s WebPageAttributeStoryArray) FillMap(to map[int]WebPageAttributeStory) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s WebPageAttributeStoryArray) ToMap() map[int]WebPageAttributeStory {
	r := make(map[int]WebPageAttributeStory, len(s))
	s.FillMap(r)
	return r
}

// WebPageAttributeStickerSetArray is adapter for slice of WebPageAttributeStickerSet.
type WebPageAttributeStickerSetArray []WebPageAttributeStickerSet

// Sort sorts slice of WebPageAttributeStickerSet.
func (s WebPageAttributeStickerSetArray) Sort(less func(a, b WebPageAttributeStickerSet) bool) WebPageAttributeStickerSetArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of WebPageAttributeStickerSet.
func (s WebPageAttributeStickerSetArray) SortStable(less func(a, b WebPageAttributeStickerSet) bool) WebPageAttributeStickerSetArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of WebPageAttributeStickerSet.
func (s WebPageAttributeStickerSetArray) Retain(keep func(x WebPageAttributeStickerSet) bool) WebPageAttributeStickerSetArray {
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
func (s WebPageAttributeStickerSetArray) First() (v WebPageAttributeStickerSet, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s WebPageAttributeStickerSetArray) Last() (v WebPageAttributeStickerSet, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *WebPageAttributeStickerSetArray) PopFirst() (v WebPageAttributeStickerSet, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero WebPageAttributeStickerSet
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *WebPageAttributeStickerSetArray) Pop() (v WebPageAttributeStickerSet, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// WebPageAttributeUniqueStarGiftArray is adapter for slice of WebPageAttributeUniqueStarGift.
type WebPageAttributeUniqueStarGiftArray []WebPageAttributeUniqueStarGift

// Sort sorts slice of WebPageAttributeUniqueStarGift.
func (s WebPageAttributeUniqueStarGiftArray) Sort(less func(a, b WebPageAttributeUniqueStarGift) bool) WebPageAttributeUniqueStarGiftArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of WebPageAttributeUniqueStarGift.
func (s WebPageAttributeUniqueStarGiftArray) SortStable(less func(a, b WebPageAttributeUniqueStarGift) bool) WebPageAttributeUniqueStarGiftArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of WebPageAttributeUniqueStarGift.
func (s WebPageAttributeUniqueStarGiftArray) Retain(keep func(x WebPageAttributeUniqueStarGift) bool) WebPageAttributeUniqueStarGiftArray {
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
func (s WebPageAttributeUniqueStarGiftArray) First() (v WebPageAttributeUniqueStarGift, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s WebPageAttributeUniqueStarGiftArray) Last() (v WebPageAttributeUniqueStarGift, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *WebPageAttributeUniqueStarGiftArray) PopFirst() (v WebPageAttributeUniqueStarGift, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero WebPageAttributeUniqueStarGift
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *WebPageAttributeUniqueStarGiftArray) Pop() (v WebPageAttributeUniqueStarGift, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
