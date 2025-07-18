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

// AutoDownloadSettings represents TL type `autoDownloadSettings#baa57628`.
// Autodownload settings
//
// See https://core.telegram.org/constructor/autoDownloadSettings for reference.
type AutoDownloadSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Disable automatic media downloads?
	Disabled bool
	// Whether to preload the first seconds of videos larger than the specified limit
	VideoPreloadLarge bool
	// Whether to preload the next audio track when you're listening to music
	AudioPreloadNext bool
	// Whether to enable data saving mode in phone calls
	PhonecallsLessData bool
	// Whether to preload stories¹; in particular, the first documentAttributeVideo²
	// preload_prefix_size bytes of story videos should be preloaded.
	//
	// Links:
	//  1) https://core.telegram.org/api/stories
	//  2) https://core.telegram.org/constructor/documentAttributeVideo
	StoriesPreload bool
	// Maximum size of photos to preload
	PhotoSizeMax int
	// Maximum size of videos to preload
	VideoSizeMax int64
	// Maximum size of other files to preload
	FileSizeMax int64
	// Maximum suggested bitrate for uploading videos
	VideoUploadMaxbitrate int
	// A limit, specifying the maximum number of files that should be downloaded in parallel
	// from the same DC, for files smaller than 20MB.
	SmallQueueActiveOperationsMax int
	// A limit, specifying the maximum number of files that should be downloaded in parallel
	// from the same DC, for files bigger than 20MB.
	LargeQueueActiveOperationsMax int
}

// AutoDownloadSettingsTypeID is TL type id of AutoDownloadSettings.
const AutoDownloadSettingsTypeID = 0xbaa57628

// Ensuring interfaces in compile-time for AutoDownloadSettings.
var (
	_ bin.Encoder     = &AutoDownloadSettings{}
	_ bin.Decoder     = &AutoDownloadSettings{}
	_ bin.BareEncoder = &AutoDownloadSettings{}
	_ bin.BareDecoder = &AutoDownloadSettings{}
)

func (a *AutoDownloadSettings) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.Disabled == false) {
		return false
	}
	if !(a.VideoPreloadLarge == false) {
		return false
	}
	if !(a.AudioPreloadNext == false) {
		return false
	}
	if !(a.PhonecallsLessData == false) {
		return false
	}
	if !(a.StoriesPreload == false) {
		return false
	}
	if !(a.PhotoSizeMax == 0) {
		return false
	}
	if !(a.VideoSizeMax == 0) {
		return false
	}
	if !(a.FileSizeMax == 0) {
		return false
	}
	if !(a.VideoUploadMaxbitrate == 0) {
		return false
	}
	if !(a.SmallQueueActiveOperationsMax == 0) {
		return false
	}
	if !(a.LargeQueueActiveOperationsMax == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AutoDownloadSettings) String() string {
	if a == nil {
		return "AutoDownloadSettings(nil)"
	}
	type Alias AutoDownloadSettings
	return fmt.Sprintf("AutoDownloadSettings%+v", Alias(*a))
}

// FillFrom fills AutoDownloadSettings from given interface.
func (a *AutoDownloadSettings) FillFrom(from interface {
	GetDisabled() (value bool)
	GetVideoPreloadLarge() (value bool)
	GetAudioPreloadNext() (value bool)
	GetPhonecallsLessData() (value bool)
	GetStoriesPreload() (value bool)
	GetPhotoSizeMax() (value int)
	GetVideoSizeMax() (value int64)
	GetFileSizeMax() (value int64)
	GetVideoUploadMaxbitrate() (value int)
	GetSmallQueueActiveOperationsMax() (value int)
	GetLargeQueueActiveOperationsMax() (value int)
}) {
	a.Disabled = from.GetDisabled()
	a.VideoPreloadLarge = from.GetVideoPreloadLarge()
	a.AudioPreloadNext = from.GetAudioPreloadNext()
	a.PhonecallsLessData = from.GetPhonecallsLessData()
	a.StoriesPreload = from.GetStoriesPreload()
	a.PhotoSizeMax = from.GetPhotoSizeMax()
	a.VideoSizeMax = from.GetVideoSizeMax()
	a.FileSizeMax = from.GetFileSizeMax()
	a.VideoUploadMaxbitrate = from.GetVideoUploadMaxbitrate()
	a.SmallQueueActiveOperationsMax = from.GetSmallQueueActiveOperationsMax()
	a.LargeQueueActiveOperationsMax = from.GetLargeQueueActiveOperationsMax()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AutoDownloadSettings) TypeID() uint32 {
	return AutoDownloadSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*AutoDownloadSettings) TypeName() string {
	return "autoDownloadSettings"
}

// TypeInfo returns info about TL type.
func (a *AutoDownloadSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "autoDownloadSettings",
		ID:   AutoDownloadSettingsTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Disabled",
			SchemaName: "disabled",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "VideoPreloadLarge",
			SchemaName: "video_preload_large",
			Null:       !a.Flags.Has(1),
		},
		{
			Name:       "AudioPreloadNext",
			SchemaName: "audio_preload_next",
			Null:       !a.Flags.Has(2),
		},
		{
			Name:       "PhonecallsLessData",
			SchemaName: "phonecalls_less_data",
			Null:       !a.Flags.Has(3),
		},
		{
			Name:       "StoriesPreload",
			SchemaName: "stories_preload",
			Null:       !a.Flags.Has(4),
		},
		{
			Name:       "PhotoSizeMax",
			SchemaName: "photo_size_max",
		},
		{
			Name:       "VideoSizeMax",
			SchemaName: "video_size_max",
		},
		{
			Name:       "FileSizeMax",
			SchemaName: "file_size_max",
		},
		{
			Name:       "VideoUploadMaxbitrate",
			SchemaName: "video_upload_maxbitrate",
		},
		{
			Name:       "SmallQueueActiveOperationsMax",
			SchemaName: "small_queue_active_operations_max",
		},
		{
			Name:       "LargeQueueActiveOperationsMax",
			SchemaName: "large_queue_active_operations_max",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *AutoDownloadSettings) SetFlags() {
	if !(a.Disabled == false) {
		a.Flags.Set(0)
	}
	if !(a.VideoPreloadLarge == false) {
		a.Flags.Set(1)
	}
	if !(a.AudioPreloadNext == false) {
		a.Flags.Set(2)
	}
	if !(a.PhonecallsLessData == false) {
		a.Flags.Set(3)
	}
	if !(a.StoriesPreload == false) {
		a.Flags.Set(4)
	}
}

// Encode implements bin.Encoder.
func (a *AutoDownloadSettings) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode autoDownloadSettings#baa57628 as nil")
	}
	b.PutID(AutoDownloadSettingsTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AutoDownloadSettings) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode autoDownloadSettings#baa57628 as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode autoDownloadSettings#baa57628: field flags: %w", err)
	}
	b.PutInt(a.PhotoSizeMax)
	b.PutLong(a.VideoSizeMax)
	b.PutLong(a.FileSizeMax)
	b.PutInt(a.VideoUploadMaxbitrate)
	b.PutInt(a.SmallQueueActiveOperationsMax)
	b.PutInt(a.LargeQueueActiveOperationsMax)
	return nil
}

// Decode implements bin.Decoder.
func (a *AutoDownloadSettings) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode autoDownloadSettings#baa57628 to nil")
	}
	if err := b.ConsumeID(AutoDownloadSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AutoDownloadSettings) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode autoDownloadSettings#baa57628 to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: field flags: %w", err)
		}
	}
	a.Disabled = a.Flags.Has(0)
	a.VideoPreloadLarge = a.Flags.Has(1)
	a.AudioPreloadNext = a.Flags.Has(2)
	a.PhonecallsLessData = a.Flags.Has(3)
	a.StoriesPreload = a.Flags.Has(4)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: field photo_size_max: %w", err)
		}
		a.PhotoSizeMax = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: field video_size_max: %w", err)
		}
		a.VideoSizeMax = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: field file_size_max: %w", err)
		}
		a.FileSizeMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: field video_upload_maxbitrate: %w", err)
		}
		a.VideoUploadMaxbitrate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: field small_queue_active_operations_max: %w", err)
		}
		a.SmallQueueActiveOperationsMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#baa57628: field large_queue_active_operations_max: %w", err)
		}
		a.LargeQueueActiveOperationsMax = value
	}
	return nil
}

// SetDisabled sets value of Disabled conditional field.
func (a *AutoDownloadSettings) SetDisabled(value bool) {
	if value {
		a.Flags.Set(0)
		a.Disabled = true
	} else {
		a.Flags.Unset(0)
		a.Disabled = false
	}
}

// GetDisabled returns value of Disabled conditional field.
func (a *AutoDownloadSettings) GetDisabled() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(0)
}

// SetVideoPreloadLarge sets value of VideoPreloadLarge conditional field.
func (a *AutoDownloadSettings) SetVideoPreloadLarge(value bool) {
	if value {
		a.Flags.Set(1)
		a.VideoPreloadLarge = true
	} else {
		a.Flags.Unset(1)
		a.VideoPreloadLarge = false
	}
}

// GetVideoPreloadLarge returns value of VideoPreloadLarge conditional field.
func (a *AutoDownloadSettings) GetVideoPreloadLarge() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(1)
}

// SetAudioPreloadNext sets value of AudioPreloadNext conditional field.
func (a *AutoDownloadSettings) SetAudioPreloadNext(value bool) {
	if value {
		a.Flags.Set(2)
		a.AudioPreloadNext = true
	} else {
		a.Flags.Unset(2)
		a.AudioPreloadNext = false
	}
}

// GetAudioPreloadNext returns value of AudioPreloadNext conditional field.
func (a *AutoDownloadSettings) GetAudioPreloadNext() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(2)
}

// SetPhonecallsLessData sets value of PhonecallsLessData conditional field.
func (a *AutoDownloadSettings) SetPhonecallsLessData(value bool) {
	if value {
		a.Flags.Set(3)
		a.PhonecallsLessData = true
	} else {
		a.Flags.Unset(3)
		a.PhonecallsLessData = false
	}
}

// GetPhonecallsLessData returns value of PhonecallsLessData conditional field.
func (a *AutoDownloadSettings) GetPhonecallsLessData() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(3)
}

// SetStoriesPreload sets value of StoriesPreload conditional field.
func (a *AutoDownloadSettings) SetStoriesPreload(value bool) {
	if value {
		a.Flags.Set(4)
		a.StoriesPreload = true
	} else {
		a.Flags.Unset(4)
		a.StoriesPreload = false
	}
}

// GetStoriesPreload returns value of StoriesPreload conditional field.
func (a *AutoDownloadSettings) GetStoriesPreload() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(4)
}

// GetPhotoSizeMax returns value of PhotoSizeMax field.
func (a *AutoDownloadSettings) GetPhotoSizeMax() (value int) {
	if a == nil {
		return
	}
	return a.PhotoSizeMax
}

// GetVideoSizeMax returns value of VideoSizeMax field.
func (a *AutoDownloadSettings) GetVideoSizeMax() (value int64) {
	if a == nil {
		return
	}
	return a.VideoSizeMax
}

// GetFileSizeMax returns value of FileSizeMax field.
func (a *AutoDownloadSettings) GetFileSizeMax() (value int64) {
	if a == nil {
		return
	}
	return a.FileSizeMax
}

// GetVideoUploadMaxbitrate returns value of VideoUploadMaxbitrate field.
func (a *AutoDownloadSettings) GetVideoUploadMaxbitrate() (value int) {
	if a == nil {
		return
	}
	return a.VideoUploadMaxbitrate
}

// GetSmallQueueActiveOperationsMax returns value of SmallQueueActiveOperationsMax field.
func (a *AutoDownloadSettings) GetSmallQueueActiveOperationsMax() (value int) {
	if a == nil {
		return
	}
	return a.SmallQueueActiveOperationsMax
}

// GetLargeQueueActiveOperationsMax returns value of LargeQueueActiveOperationsMax field.
func (a *AutoDownloadSettings) GetLargeQueueActiveOperationsMax() (value int) {
	if a == nil {
		return
	}
	return a.LargeQueueActiveOperationsMax
}
