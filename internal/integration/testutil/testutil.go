//go:build linux
// +build linux

package testutil

import (
	"testing"

	"github.com/google/nftables"
	"github.com/vishvananda/netns"
)

// SkipUnprivileged skips the test if the current user is not root.
func SkipUnprivileged(t testing.TB) { _ = "STUB: not implemented"; return }

// NewNS creates a new network namespace and returns its handle, along with a
// cleanup function to close it.
func NewNS(t testing.TB) (netns.NsHandle, func()) {
	_ = "STUB: not implemented"
	return *new(netns.NsHandle), nil
}

// Locking the thread is necessary to ensure that the namespace created by
// `netns.New()` is correctly associated with the current goroutine.

// NewNftablesConn creates a new nftables connection within the specified
// network namespace.
func NewNftablesConn(t testing.TB, ns netns.NsHandle) *nftables.Conn {
	_ = "STUB: not implemented"
	return nil
}
