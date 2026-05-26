//go:build linux

package netlink

import (
	"io"
)

// newDebugger creates a debugger by parsing key=value arguments.
func newDebugger(args []string) *debugger { _ = "STUB: not implemented"; return nil }

// debugf prints debugging information at the specified level, if d.Level is high enough to print the message.
func (d *debugger) debugf(level int, format string, v ...any) { _ = "STUB: not implemented"; return }

// nlmsgFprintfHeader prints the netlink message header to fd.
func nlmsgFprintfHeader(fd io.Writer, nlh Header) { _ = "STUB: not implemented"; return }

// nlmsgFprintf prints a single Message for netlink errors and attributes.
func nlmsgFprintf(fd io.Writer, m Message, colorize bool) { _ = "STUB: not implemented"; return }

// Neither, nothing to do.

// Errno occupies 4 bytes.

// Flags indicate an extended acknowledgement. The type/flags combination
// checked above determines the offset where the TLVs occur.

// There is an nlmsghdr preceding the TLVs.

// The TLVs should be at the offset indicated by the nlmsghdr.length,
// plus the offset where the header began. But make sure the calculated
// offset is still in-bounds.

// There is no nlmsghdr preceding the TLVs, parse them directly.

// Make sure there's at least a header's worth of data to read on each iteration.

// Extract the length of the attribute.

// extract the type

// print attribute header

// Ignore zero-length attributes.

// If nested check the next attribute

// Print the remaining attributes bytes

// ternary returns iftrue if cond is true, else iffalse.
func ternary(cond bool, iftrue string, iffalse string) string { _ = "STUB: not implemented"; return "" }
