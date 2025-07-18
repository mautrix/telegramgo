package thumbnail_test

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"image/png"

	"go.mau.fi/mautrix-telegram/pkg/gotd/telegram/thumbnail"
	"go.mau.fi/mautrix-telegram/pkg/gotd/tg"
)

func ExampleExpand() {
	userPhoto := tg.UserProfilePhoto{
		StrippedThumb: []uint8{
			0x01, 0x28, 0x28, 0x29, 0x40, 0x1d, 0xff, 0x00, 0x2a, 0x41, 0x8e, 0xe3, 0x22, 0xac, 0x42, 0x51,
			0xa0, 0x28, 0x03, 0x19, 0x1b, 0x9e, 0x07, 0x3f, 0xad, 0x69, 0x51, 0xbd, 0x91, 0xcf, 0x42, 0x29,
			0xea, 0xc8, 0x9a, 0x36, 0x0a, 0x1b, 0x6e, 0x01, 0xe0, 0x73, 0xd6, 0x99, 0xdf, 0xbf, 0xe2, 0x2a,
			0xdc, 0x71, 0xc8, 0xca, 0x11, 0x9d, 0x0e, 0xc3, 0xd7, 0x6f, 0x4f, 0xa5, 0x47, 0x2d, 0xb3, 0x2c,
			0x84, 0x20, 0x38, 0xec, 0x71, 0xd6, 0xa6, 0x12, 0xb3, 0xd4, 0xd2, 0xac, 0x14, 0x96, 0x8b, 0x52,
			0x0a, 0x2a, 0xc4, 0xb6, 0xa6, 0x38, 0x8b, 0x97, 0x04, 0x8e, 0xd8, 0xa2, 0xb6, 0x4d, 0x33, 0x95,
			0xc5, 0xad, 0x1a, 0x1b, 0x04, 0xcb, 0x11, 0xf9, 0xd0, 0x11, 0xeb, 0x8e, 0x45, 0x58, 0x8e, 0x68,
			0x93, 0x73, 0x96, 0x03, 0x79, 0xce, 0xd1, 0x55, 0x63, 0x50, 0x54, 0xb3, 0x64, 0xe3, 0xf0, 0x1f,
			0x89, 0xa4, 0x75, 0x20, 0xee, 0xe3, 0x6f, 0x6e, 0xd9, 0xa9, 0x71, 0x4d, 0x9a, 0x46, 0x72, 0x8a,
			0x2e, 0x79, 0xf2, 0x4a, 0x71, 0x02, 0x60, 0x7f, 0x79, 0xaa, 0x54, 0x46, 0x51, 0x97, 0x72, 0xcc,
			0x7f, 0x2a, 0x20, 0x74, 0x78, 0xc1, 0x40, 0x00, 0xf4, 0xf4, 0xa7, 0xb3, 0x05, 0x52, 0xc7, 0xa0,
			0xac, 0x9f, 0x63, 0x78, 0xae, 0xad, 0x94, 0xaf, 0xa4, 0xcb, 0x08, 0xc7, 0x41, 0xc9, 0xa2, 0xab,
			0xc8, 0xe6, 0x47, 0x66, 0x3d, 0xcd, 0x15, 0xbc, 0x55, 0x91, 0xc9, 0x39, 0x73, 0x3b, 0x88, 0x09,
			0x18, 0xf6, 0xe7, 0x14, 0x8c, 0xc5, 0x8e, 0x58, 0xe4, 0xd1, 0x45, 0x32, 0x6e, 0x3e, 0x19, 0x5a,
			0x27, 0xdc, 0x3a, 0x77, 0x1e, 0xb5, 0x62, 0xe2, 0xe5, 0x1e, 0x1d, 0xa9, 0xd4, 0xf5, 0xe3, 0xa5,
			0x14, 0x52, 0x71, 0x4d, 0xdc, 0xa5, 0x36, 0x95, 0x8a, 0x74, 0x51, 0x45, 0x32, 0x4f,
		},
	}

	result, err := thumbnail.Expand(userPhoto.StrippedThumb)
	if err != nil {
		panic(err)
	}

	img, err := jpeg.Decode(bytes.NewReader(result))
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		panic(err)
	}
	fmt.Println("IMAGE:" + base64.StdEncoding.EncodeToString(buf.Bytes()))
}
