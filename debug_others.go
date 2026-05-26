//go:build !linux

package netlink

// newDebugger creates a debugger by parsing key=value arguments.
func newDebugger(args []string) *debugger { _ = "STUB: not implemented"; return nil }

// debugf prints debugging information at the specified level, if d.Level is high enough to print the message.
func (d *debugger) debugf(level int, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}
