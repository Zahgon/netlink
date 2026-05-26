// Package nltest provides utilities for netlink testing.
package nltest

import (
	"iter"

	"github.com/mdlayher/netlink"
)

// PID is the netlink header PID value assigned by nltest.
const PID = 1

// MustMarshalAttributes marshals a slice of netlink.Attributes to their binary
// format, but panics if any errors occur.
func MustMarshalAttributes(attrs []netlink.Attribute) []byte { _ = "STUB: not implemented"; return nil }

// Multipart sends a slice of netlink.Messages to the caller as a
// netlink multi-part message. If less than two messages are present,
// the messages are not altered.
func Multipart(msgs []netlink.Message) ([]netlink.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Last message has header type "done" in addition to multi-part flag.

// Error returns a netlink error to the caller with the specified error
// number, in the body of the specified request message.
func Error(number int, reqs []netlink.Message) ([]netlink.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A Func is a function that can be used to test netlink.Conn interactions.
// The function can choose to return zero or more netlink messages, or an
// error if needed.
//
// For a netlink request/response interaction, a request req is populated by
// netlink.Conn.Send and passed to the function.
//
// For multicast interactions, an empty request req is passed to the function
// when netlink.Conn.Receive is called.
//
// If a Func returns an error, the error will be returned as-is to the caller.
// If no messages and io.EOF are returned, no messages and no error will be
// returned to the caller, simulating a multi-part message with no data.
type Func func(req []netlink.Message) ([]netlink.Message, error)

// Dial sets up a netlink.Conn for testing using the specified Func. All requests
// sent from the connection will be passed to the Func.  The connection should be
// closed as usual when it is no longer needed.
func Dial(fn Func) *netlink.Conn { _ = "STUB: not implemented"; return nil }

// CheckRequest returns a Func that verifies that each message in an incoming
// request has the specified netlink header type and flags in the same slice
// position index, and then passes the request through to fn.
//
// The length of the types and flags slices must match the number of requests
// passed to the returned Func, or CheckRequest will panic.
//
// As an example:
//   - types[0] and flags[0] will be checked against reqs[0]
//   - types[1] and flags[1] will be checked against reqs[1]
//   - ... and so on
//
// If an element of types or flags is set to the zero value, that check will
// be skipped for the request message that occurs at the same index.
//
// As an example, if types[0] is 0 and reqs[0].Header.Type is 1, the check will
// succeed because types[0] was not specified.
func CheckRequest(types []netlink.HeaderType, flags []netlink.HeaderFlags, fn Func) Func {
	_ = "STUB: not implemented"
	return *new(Func)
}

// A socket is a netlink.Socket used for testing.
type socket struct {
	fn Func

	msgs []netlink.Message
	err  error
}

func (c *socket) Close() error { _ = "STUB: not implemented"; return nil }

func (c *socket) SendMessages(messages []netlink.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *socket) Send(m netlink.Message) error { _ = "STUB: not implemented"; return nil }

func (c *socket) Receive() ([]netlink.Message, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *socket) ReceiveIter() iter.Seq2[netlink.Message, error] {
	_ = "STUB: not implemented"
	return nil
}

// No messages set by Send means that we are emulating a
// multicast response or an error occurred.

// If the error is a system call error, wrap it in os.NewSyscallError
// to simulate what the Linux netlink.Conn does.

// Some generic error occurred and should be passed to the caller.

// Detect multi-part messages.

// When a multi-part message is detected, the messages are returned in
// batches of half the total messages, so that multiple calls to Receive or
// ReceiveIter from netlink.Conn are needed to drain all messages.

func panicf(format string, a ...any) { _ = "STUB: not implemented"; return }
