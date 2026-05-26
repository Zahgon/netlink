//go:build !plan9 && !windows

package nltest

func isSyscallError(err error) bool { _ = "STUB: not implemented"; return false }
