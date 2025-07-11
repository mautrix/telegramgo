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

// InputPeerClassArray is adapter for slice of InputPeerClass.
type InputPeerClassArray []InputPeerClass

// Sort sorts slice of InputPeerClass.
func (s InputPeerClassArray) Sort(less func(a, b InputPeerClass) bool) InputPeerClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPeerClass.
func (s InputPeerClassArray) SortStable(less func(a, b InputPeerClass) bool) InputPeerClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPeerClass.
func (s InputPeerClassArray) Retain(keep func(x InputPeerClass) bool) InputPeerClassArray {
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
func (s InputPeerClassArray) First() (v InputPeerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPeerClassArray) Last() (v InputPeerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPeerClassArray) PopFirst() (v InputPeerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPeerClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPeerClassArray) Pop() (v InputPeerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputPeerChat returns copy with only InputPeerChat constructors.
func (s InputPeerClassArray) AsInputPeerChat() (to InputPeerChatArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPeerChat)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputPeerUser returns copy with only InputPeerUser constructors.
func (s InputPeerClassArray) AsInputPeerUser() (to InputPeerUserArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPeerUser)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputPeerChannel returns copy with only InputPeerChannel constructors.
func (s InputPeerClassArray) AsInputPeerChannel() (to InputPeerChannelArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPeerChannel)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputPeerUserFromMessage returns copy with only InputPeerUserFromMessage constructors.
func (s InputPeerClassArray) AsInputPeerUserFromMessage() (to InputPeerUserFromMessageArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPeerUserFromMessage)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputPeerChannelFromMessage returns copy with only InputPeerChannelFromMessage constructors.
func (s InputPeerClassArray) AsInputPeerChannelFromMessage() (to InputPeerChannelFromMessageArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPeerChannelFromMessage)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputPeerChatArray is adapter for slice of InputPeerChat.
type InputPeerChatArray []InputPeerChat

// Sort sorts slice of InputPeerChat.
func (s InputPeerChatArray) Sort(less func(a, b InputPeerChat) bool) InputPeerChatArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPeerChat.
func (s InputPeerChatArray) SortStable(less func(a, b InputPeerChat) bool) InputPeerChatArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPeerChat.
func (s InputPeerChatArray) Retain(keep func(x InputPeerChat) bool) InputPeerChatArray {
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
func (s InputPeerChatArray) First() (v InputPeerChat, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPeerChatArray) Last() (v InputPeerChat, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPeerChatArray) PopFirst() (v InputPeerChat, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPeerChat
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPeerChatArray) Pop() (v InputPeerChat, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputPeerUserArray is adapter for slice of InputPeerUser.
type InputPeerUserArray []InputPeerUser

// Sort sorts slice of InputPeerUser.
func (s InputPeerUserArray) Sort(less func(a, b InputPeerUser) bool) InputPeerUserArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPeerUser.
func (s InputPeerUserArray) SortStable(less func(a, b InputPeerUser) bool) InputPeerUserArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPeerUser.
func (s InputPeerUserArray) Retain(keep func(x InputPeerUser) bool) InputPeerUserArray {
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
func (s InputPeerUserArray) First() (v InputPeerUser, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPeerUserArray) Last() (v InputPeerUser, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPeerUserArray) PopFirst() (v InputPeerUser, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPeerUser
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPeerUserArray) Pop() (v InputPeerUser, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputPeerChannelArray is adapter for slice of InputPeerChannel.
type InputPeerChannelArray []InputPeerChannel

// Sort sorts slice of InputPeerChannel.
func (s InputPeerChannelArray) Sort(less func(a, b InputPeerChannel) bool) InputPeerChannelArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPeerChannel.
func (s InputPeerChannelArray) SortStable(less func(a, b InputPeerChannel) bool) InputPeerChannelArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPeerChannel.
func (s InputPeerChannelArray) Retain(keep func(x InputPeerChannel) bool) InputPeerChannelArray {
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
func (s InputPeerChannelArray) First() (v InputPeerChannel, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPeerChannelArray) Last() (v InputPeerChannel, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPeerChannelArray) PopFirst() (v InputPeerChannel, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPeerChannel
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPeerChannelArray) Pop() (v InputPeerChannel, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputPeerUserFromMessageArray is adapter for slice of InputPeerUserFromMessage.
type InputPeerUserFromMessageArray []InputPeerUserFromMessage

// Sort sorts slice of InputPeerUserFromMessage.
func (s InputPeerUserFromMessageArray) Sort(less func(a, b InputPeerUserFromMessage) bool) InputPeerUserFromMessageArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPeerUserFromMessage.
func (s InputPeerUserFromMessageArray) SortStable(less func(a, b InputPeerUserFromMessage) bool) InputPeerUserFromMessageArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPeerUserFromMessage.
func (s InputPeerUserFromMessageArray) Retain(keep func(x InputPeerUserFromMessage) bool) InputPeerUserFromMessageArray {
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
func (s InputPeerUserFromMessageArray) First() (v InputPeerUserFromMessage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPeerUserFromMessageArray) Last() (v InputPeerUserFromMessage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPeerUserFromMessageArray) PopFirst() (v InputPeerUserFromMessage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPeerUserFromMessage
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPeerUserFromMessageArray) Pop() (v InputPeerUserFromMessage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputPeerChannelFromMessageArray is adapter for slice of InputPeerChannelFromMessage.
type InputPeerChannelFromMessageArray []InputPeerChannelFromMessage

// Sort sorts slice of InputPeerChannelFromMessage.
func (s InputPeerChannelFromMessageArray) Sort(less func(a, b InputPeerChannelFromMessage) bool) InputPeerChannelFromMessageArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPeerChannelFromMessage.
func (s InputPeerChannelFromMessageArray) SortStable(less func(a, b InputPeerChannelFromMessage) bool) InputPeerChannelFromMessageArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPeerChannelFromMessage.
func (s InputPeerChannelFromMessageArray) Retain(keep func(x InputPeerChannelFromMessage) bool) InputPeerChannelFromMessageArray {
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
func (s InputPeerChannelFromMessageArray) First() (v InputPeerChannelFromMessage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPeerChannelFromMessageArray) Last() (v InputPeerChannelFromMessage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPeerChannelFromMessageArray) PopFirst() (v InputPeerChannelFromMessage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPeerChannelFromMessage
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPeerChannelFromMessageArray) Pop() (v InputPeerChannelFromMessage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
