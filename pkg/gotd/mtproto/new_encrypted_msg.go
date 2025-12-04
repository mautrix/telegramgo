package mtproto

import (
	"github.com/go-faster/errors"

	"go.mau.fi/mautrix-telegram/pkg/gotd/bin"
	"go.mau.fi/mautrix-telegram/pkg/gotd/crypto"
	"go.mau.fi/mautrix-telegram/pkg/gotd/proto"
)

func (c *Conn) newEncryptedMessage(id int64, seq int32, payload bin.Encoder, b *bin.Buffer) error {
	s := c.session()

	// TODO(tdakkota): Smarter gzip.
	// 	1) Generate Length() method for every encoder, to count length without encoding.
	// 	2) Re-use buffer instead of using yet one.
	// 	3) Do not send proto.GZIP if gzipped size is equal or bigger.
	var d crypto.EncryptedMessageData
	if c.compressThreshold <= 0 {
		d = crypto.EncryptedMessageData{
			SessionID: s.ID,
			Salt:      s.Salt,
			MessageID: id,
			SeqNo:     seq,
			Message:   payload,
		}
	} else {
		payloadBuf := bufPool.Get()
		defer bufPool.Put(payloadBuf)
		if err := payload.Encode(payloadBuf); err != nil {
			return errors.Wrap(err, "encode payload")
		}

		if payloadBuf.Len() > c.compressThreshold {
			d = crypto.EncryptedMessageData{
				SessionID: s.ID,
				Salt:      s.Salt,
				MessageID: id,
				SeqNo:     seq,
				Message:   proto.GZIP{Data: payloadBuf.Raw()},
			}
		} else {
			d = crypto.EncryptedMessageData{
				SessionID:              s.ID,
				Salt:                   s.Salt,
				MessageID:              id,
				SeqNo:                  seq,
				MessageDataLen:         int32(payloadBuf.Len()),
				MessageDataWithPadding: payloadBuf.Buf,
			}
		}
	}

	if err := c.cipher.Encrypt(s.Key, d, b); err != nil {
		return errors.Wrap(err, "encrypt")
	}

	return nil
}
