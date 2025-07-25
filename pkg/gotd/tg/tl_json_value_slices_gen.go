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

// JSONValueClassArray is adapter for slice of JSONValueClass.
type JSONValueClassArray []JSONValueClass

// Sort sorts slice of JSONValueClass.
func (s JSONValueClassArray) Sort(less func(a, b JSONValueClass) bool) JSONValueClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of JSONValueClass.
func (s JSONValueClassArray) SortStable(less func(a, b JSONValueClass) bool) JSONValueClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of JSONValueClass.
func (s JSONValueClassArray) Retain(keep func(x JSONValueClass) bool) JSONValueClassArray {
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
func (s JSONValueClassArray) First() (v JSONValueClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s JSONValueClassArray) Last() (v JSONValueClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *JSONValueClassArray) PopFirst() (v JSONValueClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero JSONValueClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *JSONValueClassArray) Pop() (v JSONValueClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsJSONBool returns copy with only JSONBool constructors.
func (s JSONValueClassArray) AsJSONBool() (to JSONBoolArray) {
	for _, elem := range s {
		value, ok := elem.(*JSONBool)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsJSONNumber returns copy with only JSONNumber constructors.
func (s JSONValueClassArray) AsJSONNumber() (to JSONNumberArray) {
	for _, elem := range s {
		value, ok := elem.(*JSONNumber)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsJSONString returns copy with only JSONString constructors.
func (s JSONValueClassArray) AsJSONString() (to JSONStringArray) {
	for _, elem := range s {
		value, ok := elem.(*JSONString)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsJSONArray returns copy with only JSONArray constructors.
func (s JSONValueClassArray) AsJSONArray() (to JSONArrayArray) {
	for _, elem := range s {
		value, ok := elem.(*JSONArray)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsJSONObject returns copy with only JSONObject constructors.
func (s JSONValueClassArray) AsJSONObject() (to JSONObjectArray) {
	for _, elem := range s {
		value, ok := elem.(*JSONObject)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// JSONBoolArray is adapter for slice of JSONBool.
type JSONBoolArray []JSONBool

// Sort sorts slice of JSONBool.
func (s JSONBoolArray) Sort(less func(a, b JSONBool) bool) JSONBoolArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of JSONBool.
func (s JSONBoolArray) SortStable(less func(a, b JSONBool) bool) JSONBoolArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of JSONBool.
func (s JSONBoolArray) Retain(keep func(x JSONBool) bool) JSONBoolArray {
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
func (s JSONBoolArray) First() (v JSONBool, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s JSONBoolArray) Last() (v JSONBool, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *JSONBoolArray) PopFirst() (v JSONBool, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero JSONBool
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *JSONBoolArray) Pop() (v JSONBool, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// JSONNumberArray is adapter for slice of JSONNumber.
type JSONNumberArray []JSONNumber

// Sort sorts slice of JSONNumber.
func (s JSONNumberArray) Sort(less func(a, b JSONNumber) bool) JSONNumberArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of JSONNumber.
func (s JSONNumberArray) SortStable(less func(a, b JSONNumber) bool) JSONNumberArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of JSONNumber.
func (s JSONNumberArray) Retain(keep func(x JSONNumber) bool) JSONNumberArray {
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
func (s JSONNumberArray) First() (v JSONNumber, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s JSONNumberArray) Last() (v JSONNumber, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *JSONNumberArray) PopFirst() (v JSONNumber, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero JSONNumber
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *JSONNumberArray) Pop() (v JSONNumber, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// JSONStringArray is adapter for slice of JSONString.
type JSONStringArray []JSONString

// Sort sorts slice of JSONString.
func (s JSONStringArray) Sort(less func(a, b JSONString) bool) JSONStringArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of JSONString.
func (s JSONStringArray) SortStable(less func(a, b JSONString) bool) JSONStringArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of JSONString.
func (s JSONStringArray) Retain(keep func(x JSONString) bool) JSONStringArray {
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
func (s JSONStringArray) First() (v JSONString, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s JSONStringArray) Last() (v JSONString, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *JSONStringArray) PopFirst() (v JSONString, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero JSONString
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *JSONStringArray) Pop() (v JSONString, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// JSONArrayArray is adapter for slice of JSONArray.
type JSONArrayArray []JSONArray

// Sort sorts slice of JSONArray.
func (s JSONArrayArray) Sort(less func(a, b JSONArray) bool) JSONArrayArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of JSONArray.
func (s JSONArrayArray) SortStable(less func(a, b JSONArray) bool) JSONArrayArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of JSONArray.
func (s JSONArrayArray) Retain(keep func(x JSONArray) bool) JSONArrayArray {
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
func (s JSONArrayArray) First() (v JSONArray, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s JSONArrayArray) Last() (v JSONArray, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *JSONArrayArray) PopFirst() (v JSONArray, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero JSONArray
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *JSONArrayArray) Pop() (v JSONArray, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// JSONObjectArray is adapter for slice of JSONObject.
type JSONObjectArray []JSONObject

// Sort sorts slice of JSONObject.
func (s JSONObjectArray) Sort(less func(a, b JSONObject) bool) JSONObjectArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of JSONObject.
func (s JSONObjectArray) SortStable(less func(a, b JSONObject) bool) JSONObjectArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of JSONObject.
func (s JSONObjectArray) Retain(keep func(x JSONObject) bool) JSONObjectArray {
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
func (s JSONObjectArray) First() (v JSONObject, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s JSONObjectArray) Last() (v JSONObject, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *JSONObjectArray) PopFirst() (v JSONObject, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero JSONObject
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *JSONObjectArray) Pop() (v JSONObject, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
