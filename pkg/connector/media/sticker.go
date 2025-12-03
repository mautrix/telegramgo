// mautrix-telegram - A Matrix-Telegram puppeting bridge.
// Copyright (C) 2025 Sumner Evans
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package media

import (
	"context"
	"os"
	"strconv"

	"github.com/rs/zerolog"
	"go.mau.fi/util/ffmpeg"
	"go.mau.fi/util/lottie"
)

type AnimatedStickerConfig struct {
	Target          string `yaml:"target"`
	ConvertFromWebm bool   `yaml:"convert_from_webm"`
	Args            struct {
		Width  int `yaml:"width"`
		Height int `yaml:"height"`
		FPS    int `yaml:"fps"`
	} `yaml:"args"`
}

type ConvertedSticker struct {
	Success           bool
	NewPath           string
	MIMEType          string
	ThumbnailData     []byte
	ThumbnailMIMEType string
	Width             int
	Height            int
	Size              int
}

func (c *AnimatedStickerConfig) convertWebm(ctx context.Context, src *os.File) *ConvertedSticker {
	if !c.ConvertFromWebm || c.Target == "webm" {
		return nil
	}
	log := zerolog.Ctx(ctx).With().Str("animated_sticker_target", c.Target).Logger()
	if !ffmpeg.Supported() {
		log.Warn().Msg("Not converting webm sticker as ffmpeg is not installed")
		return nil
	}
	var newPath string
	var err error
	switch c.Target {
	case "png":
		newPath, err = ffmpeg.ConvertPath(
			ctx, src.Name(), ".png",
			[]string{"-ss", "0", "-c:v", "libvpx-vp9"},
			[]string{"-frames:v", "1"},
			false,
		)
	case "gif":
		newPath, err = ffmpeg.ConvertPath(
			ctx, src.Name(), ".gif",
			[]string{"-c:v", "libvpx-vp9"},
			[]string{"-vf", "split[s0][s1];[s0]palettegen[p];[s1][p]paletteuse"},
			false,
		)
	case "webp":
		newPath, err = ffmpeg.ConvertPath(
			ctx, src.Name(), ".webp",
			[]string{"-c:v", "libvpx-vp9"},
			[]string{"-loop", "0"},
			false,
		)
	default:
		log.Error().Msg("Unknown target format for webm conversion")
		return nil
	}
	if err != nil {
		log.Err(err).Msg("Failed to convert webm sticker")
		return nil
	}
	var outputSize int64
	stat, err := os.Stat(newPath)
	if err != nil {
		log.Err(err).Msg("Failed to stat converted sticker")
	} else {
		outputSize = stat.Size()
	}

	_ = src.Close()
	return &ConvertedSticker{
		Success:  true,
		NewPath:  newPath,
		MIMEType: "image/" + c.Target,
		Width:    c.Args.Width,
		Height:   c.Args.Height,
		Size:     int(outputSize),
	}
}

func (c *AnimatedStickerConfig) convert(ctx context.Context, src *os.File) *ConvertedSticker {
	if c.Target == "disable" {
		return nil
	}

	log := zerolog.Ctx(ctx).With().Str("animated_sticker_target", c.Target).Logger()

	if !lottie.Supported() {
		log.Warn().Msg("Not converting lottie sticker as lottieconverter is not installed")
		return nil
	} else if (c.Target == "webp" || c.Target == "webm") && !ffmpeg.Supported() {
		log.Warn().Msg("Not converting lottie sticker as target is webp/webm, but ffmpeg is not installed")
		return nil
	}
	outputFilename := src.Name() + "." + c.Target

	var thumbnailData []byte
	var mimeType, thumbnailMIMEType string

	var err error
	switch c.Target {
	case "png":
		mimeType = "image/png"
		err = lottie.Convert(ctx, src, outputFilename, nil, c.Target, c.Args.Width, c.Args.Height, "1")
	case "gif":
		mimeType = "image/gif"
		err = lottie.Convert(ctx, src, outputFilename, nil, c.Target, c.Args.Width, c.Args.Height, strconv.Itoa(c.Args.FPS))
	case "webm", "webp":
		thumbnailMIMEType = "image/png"
		if c.Target == "webm" {
			mimeType = "video/webm"
		} else {
			mimeType = "image/webp"
		}
		thumbnailData, err = lottie.FFmpegConvert(ctx, src, outputFilename, c.Args.Width, c.Args.Height, c.Args.FPS)
		if err != nil {
			break
		}
	default:
		log.Error().Msg("Unknown target format")
		return nil
	}
	if err != nil {
		_ = os.Remove(outputFilename)
		log.Err(err).Msg("Failed to convert animated sticker")
		return nil
	}
	var outputSize int64
	stat, err := os.Stat(outputFilename)
	if err != nil {
		log.Err(err).Msg("Failed to stat converted sticker")
	} else {
		outputSize = stat.Size()
	}

	_ = src.Close()
	return &ConvertedSticker{
		Success:           true,
		NewPath:           outputFilename,
		MIMEType:          mimeType,
		ThumbnailData:     thumbnailData,
		ThumbnailMIMEType: thumbnailMIMEType,
		Width:             c.Args.Width,
		Height:            c.Args.Height,
		Size:              int(outputSize),
	}
}
