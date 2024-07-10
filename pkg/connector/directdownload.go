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

	transferer := media.NewTransferer(client.client.API())
	var readyTransferer *media.ReadyTransferer
	switch msgMedia := msgMedia.(type) {
	case *tg.MessageMediaPhoto:
		readyTransferer = transferer.WithPhoto(msgMedia.Photo)
	case *tg.MessageMediaDocument:
		document, ok := msgMedia.Document.(*tg.Document)
		if !ok {
			return nil, fmt.Errorf("unknown document type %T", msgMedia.Document)
		}
		for _, attr := range document.GetAttributes() {
			if attr.TypeID() == tg.DocumentAttributeStickerTypeID {
				transferer = transferer.WithStickerConfig(tc.Config.AnimatedSticker)
			}
		}

		readyTransferer = transferer.WithDocument(msgMedia.Document, info.Thumbnail)
	default:
		return nil, fmt.Errorf("unhandled media type %T", msgMedia)
	}

	data, fileInfo, err := readyTransferer.Download(ctx)
	if err != nil {
		log.Err(err).Msg("failed to download media")
		return nil, err
	}

	return &mediaproxy.GetMediaResponseData{
		Reader:        io.NopCloser(bytes.NewBuffer(data)),
		ContentType:   fileInfo.MimeType,
		ContentLength: int64(fileInfo.Size),
	}, nil
}

func (tg *TelegramConnector) SetUseDirectMedia() {
	tg.useDirectMedia = true
}
