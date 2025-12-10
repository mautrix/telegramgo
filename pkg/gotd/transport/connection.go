package transport

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/go-faster/errors"

	"go.mau.fi/mautrix-telegram/pkg/gotd/bin"
)

// Conn is transport connection.
type Conn interface {
	Send(ctx context.Context, b *bin.Buffer) error
	Recv(ctx context.Context, b *bin.Buffer) error
	Close() error
}

var _ Conn = (*connection)(nil)

// connection is MTProto connection.
type connection struct {
	conn  net.Conn
	codec Codec

	readMux  sync.Mutex
	writeMux sync.Mutex
}

var (
	DefaultWriteDeadline = 1 * time.Minute
	DefaultReadDeadline  = 2 * time.Minute
)

// Send sends message from buffer using MTProto connection.
func (c *connection) Send(ctx context.Context, b *bin.Buffer) error {
	// Serializing access to deadlines.
	c.writeMux.Lock()
	defer c.writeMux.Unlock()

	deadline, ok := ctx.Deadline()
	if !ok {
		deadline = time.Now().Add(DefaultWriteDeadline)
	}
	if err := c.conn.SetWriteDeadline(deadline); err != nil {
		return errors.Wrap(err, "set write deadline")
	}

	if err := c.codec.Write(c.conn, b); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

// Recv reads message to buffer using MTProto connection.
func (c *connection) Recv(ctx context.Context, b *bin.Buffer) error {
	// Serializing access to deadlines.
	c.readMux.Lock()
	defer c.readMux.Unlock()

	deadline, ok := ctx.Deadline()
	if !ok {
		deadline = time.Now().Add(DefaultReadDeadline)
	}
	if err := c.conn.SetReadDeadline(deadline); err != nil {
		return errors.Wrap(err, "set read deadline")
	}

	if err := c.codec.Read(c.conn, b); err != nil {
		return errors.Wrap(err, "read")
	}

	return nil
}

// Close closes MTProto connection.
func (c *connection) Close() error {
	return c.conn.Close()
}
