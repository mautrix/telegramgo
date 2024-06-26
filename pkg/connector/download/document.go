package download

import (
	"bytes"
	"context"

	"github.com/gotd/td/telegram/downloader"
	"github.com/gotd/td/tg"
)

func DownloadDocument(ctx context.Context, client downloader.Client, document *tg.Document) (data []byte, err error) {
	file := tg.InputDocumentFileLocation{
		ID:            document.GetID(),
		AccessHash:    document.GetAccessHash(),
		FileReference: document.GetFileReference(),
	}
	var buf bytes.Buffer
	_, err = downloader.NewDownloader().Download(client, &file).Stream(ctx, &buf)
	data = buf.Bytes()
	return
}
