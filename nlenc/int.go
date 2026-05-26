package nlenc

// PutUint8 encodes a uint8 into b.
// If b is not exactly 1 byte in length, PutUint8 will panic.
//
// Deprecated: Use inline "b[n] = v" instead.
func PutUint8(b []byte, v uint8) { _ = "STUB: not implemented"; return }

// PutUint16 encodes a uint16 into b using the host machine's native endianness.
// If b is not exactly 2 bytes in length, PutUint16 will panic.
//
// Deprecated: Use [binary.NativeEndian.PutUint16] instead.
func PutUint16(b []byte, v uint16) { _ = "STUB: not implemented"; return }

// PutUint32 encodes a uint32 into b using the host machine's native endianness.
// If b is not exactly 4 bytes in length, PutUint32 will panic.
//
// Deprecated: Use [binary.NativeEndian.PutUint32] instead.
func PutUint32(b []byte, v uint32) { _ = "STUB: not implemented"; return }

// PutUint64 encodes a uint64 into b using the host machine's native endianness.
// If b is not exactly 8 bytes in length, PutUint64 will panic.
//
// Deprecated: Use [binary.NativeEndian.PutUint64] instead.
func PutUint64(b []byte, v uint64) { _ = "STUB: not implemented"; return }

// PutInt32 encodes a int32 into b using the host machine's native endianness.
// If b is not exactly 4 bytes in length, PutInt32 will panic.
//
// Deprecated: Use [binary.NativeEndian.PutUint32] instead (with a typecast).
func PutInt32(b []byte, v int32) { _ = "STUB: not implemented"; return }

// Uint8 decodes a uint8 from b.
// If b is not exactly 1 byte in length, Uint8 will panic.
//
// Deprecated: Use inline "b[n]" instead.
func Uint8(b []byte) uint8 { _ = "STUB: not implemented"; return 0 }

// Uint16 decodes a uint16 from b using the host machine's native endianness.
// If b is not exactly 2 bytes in length, Uint16 will panic.
//
// Deprecated: Use [binary.NativeEndian.Uint16] instead.
func Uint16(b []byte) uint16 { _ = "STUB: not implemented"; return 0 }

// Uint32 decodes a uint32 from b using the host machine's native endianness.
// If b is not exactly 4 bytes in length, Uint32 will panic.
//
// Deprecated: Use [binary.NativeEndian.Uint32] instead.
func Uint32(b []byte) uint32 { _ = "STUB: not implemented"; return 0 }

// Uint64 decodes a uint64 from b using the host machine's native endianness.
// If b is not exactly 8 bytes in length, Uint64 will panic.
//
// Deprecated: Use [binary.NativeEndian.Uint64] instead.
func Uint64(b []byte) uint64 { _ = "STUB: not implemented"; return 0 }

// Int32 decodes an int32 from b using the host machine's native endianness.
// If b is not exactly 4 bytes in length, Int32 will panic.
//
// Deprecated: Use [binary.NativeEndian.Uint32] instead (with a typecast).
func Int32(b []byte) int32 { _ = "STUB: not implemented"; return 0 }

// Uint8Bytes encodes a uint8 into a newly-allocated byte slice. It is a
// shortcut for allocating a new byte slice and filling it using PutUint8.
//
// Deprecated: Use inline "[]byte{v}" instead.
func Uint8Bytes(v uint8) []byte {
	_ = "STUB: not implemented"

	// Uint16Bytes encodes a uint16 into a newly-allocated byte slice using the
	// host machine's native endianness.  It is a shortcut for allocating a new
	// byte slice and filling it using PutUint16.
	return nil
}

func Uint16Bytes(v uint16) []byte { _ = "STUB: not implemented"; return nil }

// Uint32Bytes encodes a uint32 into a newly-allocated byte slice using the
// host machine's native endianness.  It is a shortcut for allocating a new
// byte slice and filling it using PutUint32.
func Uint32Bytes(v uint32) []byte { _ = "STUB: not implemented"; return nil }

// Uint64Bytes encodes a uint64 into a newly-allocated byte slice using the
// host machine's native endianness.  It is a shortcut for allocating a new
// byte slice and filling it using PutUint64.
func Uint64Bytes(v uint64) []byte { _ = "STUB: not implemented"; return nil }

// Int32Bytes encodes a int32 into a newly-allocated byte slice using the
// host machine's native endianness.  It is a shortcut for allocating a new
// byte slice and filling it using PutInt32.
func Int32Bytes(v int32) []byte { _ = "STUB: not implemented"; return nil }
