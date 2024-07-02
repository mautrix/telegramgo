package connector

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/gotd/td/tg"
	"github.com/rs/zerolog"
	"maunium.net/go/mautrix/bridgev2"
	"maunium.net/go/mautrix/bridgev2/networkid"
	"maunium.net/go/mautrix/mediaproxy"

	"go.mau.fi/mautrix-telegram/pkg/connector/ids"
	"go.mau.fi/mautrix-telegram/pkg/connector/media"
)

var _ bridgev2.DirectMediableNetwork = (*TelegramConnector)(nil)

type getMessages interface {
	GetMessages() []tg.MessageClass
}

func (tc *TelegramConnector) Download(ctx context.Context, mediaID networkid.MediaID) (mediaproxy.GetMediaResponse, error) {
	info, err := ids.ParseDirectMediaInfo(mediaID)
	if err != nil {
		return nil, err
	}
	log := zerolog.Ctx(ctx).With().
		Str("component", "direct download").
		Any("info", info).
		Logger()
	ctx = log.WithContext(ctx)
	log.Info().Msg("handling direct download")

	logins, err := tc.Bridge.GetUserLoginsInPortal(ctx, info.PeerType.AsPortalKey(info.ChatID))
	if err != nil {
		return nil, err
	} else if len(logins) == 0 {
		return nil, fmt.Errorf("no user logins in the portal (%s %d)", info.PeerType, info.ChatID)
	}

	client := logins[0].Client.(*TelegramClient)
	var messages tg.MessagesMessagesClass
	switch info.PeerType {
	case ids.PeerTypeUser, ids.PeerTypeChat:
		messages, err = client.client.API().MessagesGetMessages(ctx, []tg.InputMessageClass{
			&tg.InputMessageID{ID: int(info.MessageID)},
		})
	case ids.PeerTypeChannel:
		// TODO test this
		messages, err = client.client.API().ChannelsGetMessages(ctx, &tg.ChannelsGetMessagesRequest{
			Channel: &tg.InputChannel{ChannelID: info.ChatID},
			ID: []tg.InputMessageClass{
				&tg.InputMessageID{ID: int(info.MessageID)},
			},
		})
	default:
		return nil, fmt.Errorf("unknown peer type %s", info.PeerType)
	}
	if err != nil {
		return nil, err
	}

	var msgMedia tg.MessageMediaClass
	if m, ok := messages.(getMessages); !ok {
		return nil, fmt.Errorf("unknown message type %T", messages)
	} else {
		var found bool
		for _, message := range m.GetMessages() {
			if msg, ok := message.(*tg.Message); ok && msg.ID == int(info.MessageID) {
				msgMedia = msg.Media
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("no media found with ID %d", info.MessageID)
		}
	}

	var data []byte
	var mimeType string
	switch msgMedia := msgMedia.(type) {
	case *tg.MessageMediaPhoto:
		data, _, _, mimeType, err = media.DownloadPhotoMedia(ctx, client.client.API(), msgMedia)
	case *tg.MessageMediaDocument:
		document, ok := msgMedia.Document.(*tg.Document)
		if !ok {
			return nil, fmt.Errorf("unrecognized document type %T", msgMedia.Document)
		}

		// Download the thumbnail for this media rather than the media itself.
		if info.Thumbnail {
			_, _, largestThumbnail := media.GetLargestPhotoSize(document.Thumbs)
			data, mimeType, err = media.DownloadFileLocation(ctx, client.client.API(), &tg.InputDocumentFileLocation{
				ID:            document.GetID(),
				AccessHash:    document.GetAccessHash(),
				FileReference: document.GetFileReference(),
				ThumbSize:     largestThumbnail.GetType(),
			})
		} else {
			mimeType = document.GetMimeType()
			data, err = media.DownloadDocument(ctx, client.client.API(), document)
		}
	default:
		return nil, fmt.Errorf("unhandled media type %T", msgMedia)
	}
	if err != nil {
		return nil, err
	}

	return &mediaproxy.GetMediaResponseData{
		Reader:        io.NopCloser(bytes.NewBuffer(data)),
		ContentType:   mimeType,
		ContentLength: int64(len(data)),
	}, nil
}

func (tg *TelegramConnector) SetUseDirectMedia() {
	tg.useDirectMedia = true
}
