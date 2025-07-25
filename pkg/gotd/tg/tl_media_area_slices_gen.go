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

// MediaAreaClassArray is adapter for slice of MediaAreaClass.
type MediaAreaClassArray []MediaAreaClass

// Sort sorts slice of MediaAreaClass.
func (s MediaAreaClassArray) Sort(less func(a, b MediaAreaClass) bool) MediaAreaClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MediaAreaClass.
func (s MediaAreaClassArray) SortStable(less func(a, b MediaAreaClass) bool) MediaAreaClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MediaAreaClass.
func (s MediaAreaClassArray) Retain(keep func(x MediaAreaClass) bool) MediaAreaClassArray {
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
func (s MediaAreaClassArray) First() (v MediaAreaClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MediaAreaClassArray) Last() (v MediaAreaClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MediaAreaClassArray) PopFirst() (v MediaAreaClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MediaAreaClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MediaAreaClassArray) Pop() (v MediaAreaClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMediaAreaVenue returns copy with only MediaAreaVenue constructors.
func (s MediaAreaClassArray) AsMediaAreaVenue() (to MediaAreaVenueArray) {
	for _, elem := range s {
		value, ok := elem.(*MediaAreaVenue)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaAreaVenue returns copy with only InputMediaAreaVenue constructors.
func (s MediaAreaClassArray) AsInputMediaAreaVenue() (to InputMediaAreaVenueArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaAreaVenue)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMediaAreaGeoPoint returns copy with only MediaAreaGeoPoint constructors.
func (s MediaAreaClassArray) AsMediaAreaGeoPoint() (to MediaAreaGeoPointArray) {
	for _, elem := range s {
		value, ok := elem.(*MediaAreaGeoPoint)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMediaAreaSuggestedReaction returns copy with only MediaAreaSuggestedReaction constructors.
func (s MediaAreaClassArray) AsMediaAreaSuggestedReaction() (to MediaAreaSuggestedReactionArray) {
	for _, elem := range s {
		value, ok := elem.(*MediaAreaSuggestedReaction)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMediaAreaChannelPost returns copy with only MediaAreaChannelPost constructors.
func (s MediaAreaClassArray) AsMediaAreaChannelPost() (to MediaAreaChannelPostArray) {
	for _, elem := range s {
		value, ok := elem.(*MediaAreaChannelPost)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaAreaChannelPost returns copy with only InputMediaAreaChannelPost constructors.
func (s MediaAreaClassArray) AsInputMediaAreaChannelPost() (to InputMediaAreaChannelPostArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaAreaChannelPost)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMediaAreaURL returns copy with only MediaAreaURL constructors.
func (s MediaAreaClassArray) AsMediaAreaURL() (to MediaAreaURLArray) {
	for _, elem := range s {
		value, ok := elem.(*MediaAreaURL)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMediaAreaWeather returns copy with only MediaAreaWeather constructors.
func (s MediaAreaClassArray) AsMediaAreaWeather() (to MediaAreaWeatherArray) {
	for _, elem := range s {
		value, ok := elem.(*MediaAreaWeather)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMediaAreaStarGift returns copy with only MediaAreaStarGift constructors.
func (s MediaAreaClassArray) AsMediaAreaStarGift() (to MediaAreaStarGiftArray) {
	for _, elem := range s {
		value, ok := elem.(*MediaAreaStarGift)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// MediaAreaVenueArray is adapter for slice of MediaAreaVenue.
type MediaAreaVenueArray []MediaAreaVenue

// Sort sorts slice of MediaAreaVenue.
func (s MediaAreaVenueArray) Sort(less func(a, b MediaAreaVenue) bool) MediaAreaVenueArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MediaAreaVenue.
func (s MediaAreaVenueArray) SortStable(less func(a, b MediaAreaVenue) bool) MediaAreaVenueArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MediaAreaVenue.
func (s MediaAreaVenueArray) Retain(keep func(x MediaAreaVenue) bool) MediaAreaVenueArray {
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
func (s MediaAreaVenueArray) First() (v MediaAreaVenue, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MediaAreaVenueArray) Last() (v MediaAreaVenue, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MediaAreaVenueArray) PopFirst() (v MediaAreaVenue, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MediaAreaVenue
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MediaAreaVenueArray) Pop() (v MediaAreaVenue, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaAreaVenueArray is adapter for slice of InputMediaAreaVenue.
type InputMediaAreaVenueArray []InputMediaAreaVenue

// Sort sorts slice of InputMediaAreaVenue.
func (s InputMediaAreaVenueArray) Sort(less func(a, b InputMediaAreaVenue) bool) InputMediaAreaVenueArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaAreaVenue.
func (s InputMediaAreaVenueArray) SortStable(less func(a, b InputMediaAreaVenue) bool) InputMediaAreaVenueArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaAreaVenue.
func (s InputMediaAreaVenueArray) Retain(keep func(x InputMediaAreaVenue) bool) InputMediaAreaVenueArray {
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
func (s InputMediaAreaVenueArray) First() (v InputMediaAreaVenue, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaAreaVenueArray) Last() (v InputMediaAreaVenue, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaAreaVenueArray) PopFirst() (v InputMediaAreaVenue, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaAreaVenue
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaAreaVenueArray) Pop() (v InputMediaAreaVenue, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MediaAreaGeoPointArray is adapter for slice of MediaAreaGeoPoint.
type MediaAreaGeoPointArray []MediaAreaGeoPoint

// Sort sorts slice of MediaAreaGeoPoint.
func (s MediaAreaGeoPointArray) Sort(less func(a, b MediaAreaGeoPoint) bool) MediaAreaGeoPointArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MediaAreaGeoPoint.
func (s MediaAreaGeoPointArray) SortStable(less func(a, b MediaAreaGeoPoint) bool) MediaAreaGeoPointArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MediaAreaGeoPoint.
func (s MediaAreaGeoPointArray) Retain(keep func(x MediaAreaGeoPoint) bool) MediaAreaGeoPointArray {
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
func (s MediaAreaGeoPointArray) First() (v MediaAreaGeoPoint, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MediaAreaGeoPointArray) Last() (v MediaAreaGeoPoint, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MediaAreaGeoPointArray) PopFirst() (v MediaAreaGeoPoint, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MediaAreaGeoPoint
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MediaAreaGeoPointArray) Pop() (v MediaAreaGeoPoint, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MediaAreaSuggestedReactionArray is adapter for slice of MediaAreaSuggestedReaction.
type MediaAreaSuggestedReactionArray []MediaAreaSuggestedReaction

// Sort sorts slice of MediaAreaSuggestedReaction.
func (s MediaAreaSuggestedReactionArray) Sort(less func(a, b MediaAreaSuggestedReaction) bool) MediaAreaSuggestedReactionArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MediaAreaSuggestedReaction.
func (s MediaAreaSuggestedReactionArray) SortStable(less func(a, b MediaAreaSuggestedReaction) bool) MediaAreaSuggestedReactionArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MediaAreaSuggestedReaction.
func (s MediaAreaSuggestedReactionArray) Retain(keep func(x MediaAreaSuggestedReaction) bool) MediaAreaSuggestedReactionArray {
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
func (s MediaAreaSuggestedReactionArray) First() (v MediaAreaSuggestedReaction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MediaAreaSuggestedReactionArray) Last() (v MediaAreaSuggestedReaction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MediaAreaSuggestedReactionArray) PopFirst() (v MediaAreaSuggestedReaction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MediaAreaSuggestedReaction
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MediaAreaSuggestedReactionArray) Pop() (v MediaAreaSuggestedReaction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MediaAreaChannelPostArray is adapter for slice of MediaAreaChannelPost.
type MediaAreaChannelPostArray []MediaAreaChannelPost

// Sort sorts slice of MediaAreaChannelPost.
func (s MediaAreaChannelPostArray) Sort(less func(a, b MediaAreaChannelPost) bool) MediaAreaChannelPostArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MediaAreaChannelPost.
func (s MediaAreaChannelPostArray) SortStable(less func(a, b MediaAreaChannelPost) bool) MediaAreaChannelPostArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MediaAreaChannelPost.
func (s MediaAreaChannelPostArray) Retain(keep func(x MediaAreaChannelPost) bool) MediaAreaChannelPostArray {
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
func (s MediaAreaChannelPostArray) First() (v MediaAreaChannelPost, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MediaAreaChannelPostArray) Last() (v MediaAreaChannelPost, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MediaAreaChannelPostArray) PopFirst() (v MediaAreaChannelPost, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MediaAreaChannelPost
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MediaAreaChannelPostArray) Pop() (v MediaAreaChannelPost, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaAreaChannelPostArray is adapter for slice of InputMediaAreaChannelPost.
type InputMediaAreaChannelPostArray []InputMediaAreaChannelPost

// Sort sorts slice of InputMediaAreaChannelPost.
func (s InputMediaAreaChannelPostArray) Sort(less func(a, b InputMediaAreaChannelPost) bool) InputMediaAreaChannelPostArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaAreaChannelPost.
func (s InputMediaAreaChannelPostArray) SortStable(less func(a, b InputMediaAreaChannelPost) bool) InputMediaAreaChannelPostArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaAreaChannelPost.
func (s InputMediaAreaChannelPostArray) Retain(keep func(x InputMediaAreaChannelPost) bool) InputMediaAreaChannelPostArray {
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
func (s InputMediaAreaChannelPostArray) First() (v InputMediaAreaChannelPost, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaAreaChannelPostArray) Last() (v InputMediaAreaChannelPost, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaAreaChannelPostArray) PopFirst() (v InputMediaAreaChannelPost, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaAreaChannelPost
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaAreaChannelPostArray) Pop() (v InputMediaAreaChannelPost, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MediaAreaURLArray is adapter for slice of MediaAreaURL.
type MediaAreaURLArray []MediaAreaURL

// Sort sorts slice of MediaAreaURL.
func (s MediaAreaURLArray) Sort(less func(a, b MediaAreaURL) bool) MediaAreaURLArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MediaAreaURL.
func (s MediaAreaURLArray) SortStable(less func(a, b MediaAreaURL) bool) MediaAreaURLArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MediaAreaURL.
func (s MediaAreaURLArray) Retain(keep func(x MediaAreaURL) bool) MediaAreaURLArray {
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
func (s MediaAreaURLArray) First() (v MediaAreaURL, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MediaAreaURLArray) Last() (v MediaAreaURL, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MediaAreaURLArray) PopFirst() (v MediaAreaURL, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MediaAreaURL
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MediaAreaURLArray) Pop() (v MediaAreaURL, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MediaAreaWeatherArray is adapter for slice of MediaAreaWeather.
type MediaAreaWeatherArray []MediaAreaWeather

// Sort sorts slice of MediaAreaWeather.
func (s MediaAreaWeatherArray) Sort(less func(a, b MediaAreaWeather) bool) MediaAreaWeatherArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MediaAreaWeather.
func (s MediaAreaWeatherArray) SortStable(less func(a, b MediaAreaWeather) bool) MediaAreaWeatherArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MediaAreaWeather.
func (s MediaAreaWeatherArray) Retain(keep func(x MediaAreaWeather) bool) MediaAreaWeatherArray {
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
func (s MediaAreaWeatherArray) First() (v MediaAreaWeather, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MediaAreaWeatherArray) Last() (v MediaAreaWeather, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MediaAreaWeatherArray) PopFirst() (v MediaAreaWeather, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MediaAreaWeather
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MediaAreaWeatherArray) Pop() (v MediaAreaWeather, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MediaAreaStarGiftArray is adapter for slice of MediaAreaStarGift.
type MediaAreaStarGiftArray []MediaAreaStarGift

// Sort sorts slice of MediaAreaStarGift.
func (s MediaAreaStarGiftArray) Sort(less func(a, b MediaAreaStarGift) bool) MediaAreaStarGiftArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MediaAreaStarGift.
func (s MediaAreaStarGiftArray) SortStable(less func(a, b MediaAreaStarGift) bool) MediaAreaStarGiftArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MediaAreaStarGift.
func (s MediaAreaStarGiftArray) Retain(keep func(x MediaAreaStarGift) bool) MediaAreaStarGiftArray {
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
func (s MediaAreaStarGiftArray) First() (v MediaAreaStarGift, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MediaAreaStarGiftArray) Last() (v MediaAreaStarGift, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MediaAreaStarGiftArray) PopFirst() (v MediaAreaStarGift, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MediaAreaStarGift
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MediaAreaStarGiftArray) Pop() (v MediaAreaStarGift, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
