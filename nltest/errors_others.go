//go:build plan9 || windows
// +build plan9 windows

package nltest

func isSyscallError(_ error) bool { _ = "STUB: not implemented"; return false }
