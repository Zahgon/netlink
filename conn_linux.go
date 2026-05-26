//go:build linux

package netlink

import (
	"iter"
	"sync"
	"syscall"
	"time"

	"github.com/mdlayher/socket"
	"golang.org/x/net/bpf"
)

var _ Socket = &conn{}

// A conn is the Linux implementation of a netlink sockets connection.
type conn struct {
	s    *socket.Conn
	pool *sync.Pool
}

// dial is the entry point for Dial. dial opens a netlink socket using
// system calls, and returns its PID.
func dial(family int, config *Config) (*conn, uint32, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Prepare the netlink socket.

// newConn binds a connection to netlink using the input *socket.Conn.
func newConn(s *socket.Conn, config *Config) (*conn, uint32, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Socket must be closed in the event of any system call errors, to avoid
// leaking file descriptors.

// The caller has requested the strict option set. Historically we have
// recommended checking for ENOPROTOOPT if the kernel does not support
// the option in question, but that may result in a silent failure and
// unexpected behavior for the user.
//
// Treat any error here as a fatal error, and require the caller to deal
// with it.

// SendMessages serializes multiple Messages and sends them to netlink.
func (c *conn) SendMessages(messages []Message) error { _ = "STUB: not implemented"; return nil }

// Send sends a single Message to netlink.
func (c *conn) Send(m Message) error { _ = "STUB: not implemented"; return nil }

// Receive receives one or more Messages from netlink.
func (c *conn) Receive() ([]Message, error) { _ = "STUB: not implemented"; return nil, nil }

// getBuffer returns the buffer to use for receiving messages and a function to
// release it back to the pool if applicable. If the pool is not configured, a
// new buffer is allocated by peeking the size of the next message to be
// received. The buffer size is aligned to the next multiple of the alignment of a netlink
// message, to avoid panic when parsing messages from the buffer.
func (c *conn) getBuffer() ([]byte, func(), error) { _ = "STUB: not implemented"; return nil, nil, nil }

// ownedBuffer trims b to the aligned size of the received datagram and, when
// using pooled buffers, copies it to memory owned by the caller.
func (c *conn) ownedBuffer(b []byte, n int) []byte { _ = "STUB: not implemented"; return nil }

// A pooled buffer may not have an aligned length. Copy the received bytes
// to an aligned buffer and let the zero-value tail bytes act as padding.

// ReceiveIter returns an iterator over Messages received from netlink.
func (c *conn) ReceiveIter() iter.Seq2[Message, error] { _ = "STUB: not implemented"; return nil }

// Read out all available messages
// TODO(mdlayher): deal with OOB message data if available, such as
// when PacketInfo ConnOption is true.

// Our buffer was too small to read the entire message,
// this should not happen since we peeked above, but if it does,
// return an error.

// Close closes the connection.
func (c *conn) Close() error {
	_ = "STUB: not implemented"

	// JoinGroup joins a multicast group by ID.
	return nil
}

func (c *conn) JoinGroup(group uint32) error { _ = "STUB: not implemented"; return nil }

// LeaveGroup leaves a multicast group by ID.
func (c *conn) LeaveGroup(group uint32) error { _ = "STUB: not implemented"; return nil }

// SetBPF attaches an assembled BPF program to a conn.
func (c *conn) SetBPF(filter []bpf.RawInstruction) error { _ = "STUB: not implemented"; return nil }

// RemoveBPF removes a BPF filter from a conn.
func (c *conn) RemoveBPF() error { _ = "STUB: not implemented"; return nil }

// SetOption enables or disables a netlink socket option for the Conn.
func (c *conn) SetOption(option ConnOption, enable bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Return the typical Linux error for an unknown ConnOption.

func (c *conn) SetDeadline(t time.Time) error      { _ = "STUB: not implemented"; return nil }
func (c *conn) SetReadDeadline(t time.Time) error  { _ = "STUB: not implemented"; return nil }
func (c *conn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetReadBuffer sets the size of the operating system's receive buffer
// associated with the Conn.
func (c *conn) SetReadBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }

// SetReadBuffer sets the size of the operating system's transmit buffer
// associated with the Conn.
func (c *conn) SetWriteBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }

// ReadBuffer reads the size of the operating system's receive buffer
// associated with the Conn.
func (c *conn) ReadBuffer() (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// WriteBuffer reads the size of the operating system's transmit buffer
		// associated with the Conn.
		nil
}

func (c *conn) WriteBuffer() (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// SyscallConn returns a raw network connection.
		nil
}

func (c *conn) SyscallConn() (syscall.RawConn, error) {
	_ = "STUB: not implemented"
	return *

	// linuxOption converts a ConnOption to its Linux value.
	new(syscall.RawConn), nil
}

func linuxOption(o ConnOption) (int, bool) { _ = "STUB: not implemented"; return 0, false }

// sysToHeader converts a syscall.NlMsghdr to a Header.
func sysToHeader(r syscall.NlMsghdr) Header {
	_ = "STUB: not implemented"
	// NB: the memory layout of Header and syscall.NlMsgHdr must be
	// exactly the same for this unsafe cast to work
	return *new(Header)
}

// newError converts an error number from netlink into the appropriate
// system call error for Linux.
func newError(errno int) error { _ = "STUB: not implemented"; return nil }
