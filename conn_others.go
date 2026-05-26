//go:build !linux
// +build !linux

package netlink

import (
	"fmt"
	"iter"
	"runtime"
)

// errUnimplemented is returned by all functions on platforms that
// cannot make use of netlink sockets.
var errUnimplemented = fmt.Errorf("netlink: not implemented on %s/%s",
	runtime.GOOS, runtime.GOARCH)

var _ Socket = &conn{}

// A conn is the no-op implementation of a netlink sockets connection.
type conn struct{}

// All cross-platform functions and Socket methods are unimplemented outside
// of Linux.

func dial(_ int, _ *Config) (*conn, uint32, error) { _ = "STUB: not implemented"; return nil, 0, nil }
func newError(_ int) error                         { _ = "STUB: not implemented"; return nil }

func (c *conn) Send(_ Message) error                   { _ = "STUB: not implemented"; return nil }
func (c *conn) SendMessages(_ []Message) error         { _ = "STUB: not implemented"; return nil }
func (c *conn) Receive() ([]Message, error)            { _ = "STUB: not implemented"; return nil, nil }
func (c *conn) Close() error                           { _ = "STUB: not implemented"; return nil }
func (c *conn) ReceiveIter() iter.Seq2[Message, error] { _ = "STUB: not implemented"; return nil }
