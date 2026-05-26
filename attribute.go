package netlink

import (
	"encoding/binary"
	"errors"
)

// errInvalidAttribute specifies if an Attribute's length is incorrect.
var errInvalidAttribute = errors.New("invalid attribute; length too short or too large")

// An Attribute is a netlink attribute.  Attributes are packed and unpacked
// to and from the Data field of Message for some netlink families.
type Attribute struct {
	// Length of an Attribute, including this field and Type.
	Length uint16

	// The type of this Attribute, typically matched to a constant. Note that
	// flags such as Nested and NetByteOrder must be handled manually when
	// working with Attribute structures directly.
	Type uint16

	// An arbitrary payload which is specified by Type.
	Data []byte
}

// marshal marshals the contents of a into b and returns the number of bytes
// written to b, including attribute alignment padding.
func (a *Attribute) marshal(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// unmarshal unmarshals the contents of a byte slice into an Attribute.
func (a *Attribute) unmarshal(b []byte) error { _ = "STUB: not implemented"; return nil }

// No length, no data

// Not enough length for any data

// Data present

// MarshalAttributes packs a slice of Attributes into a single byte slice.
// In most cases, the Length field of each Attribute should be set to 0, so it
// can be calculated and populated automatically for each Attribute.
//
// It is recommend to use the AttributeEncoder type where possible instead of
// calling MarshalAttributes and using package nlenc functions directly.
func MarshalAttributes(attrs []Attribute) ([]byte, error) {
	_ = "STUB: not implemented"
	// Count how many bytes we should allocate to store each attribute's contents.
	return nil, nil
}

// Advance through b with idx to place attribute data at the correct offset.

// Infer the length of attribute if zero.

// Marshal a into b and advance idx to show many bytes are occupied.

// UnmarshalAttributes unpacks a slice of Attributes from a single byte slice.
//
// It is recommend to use the AttributeDecoder type where possible instead of calling
// UnmarshalAttributes and using package nlenc functions directly.
func UnmarshalAttributes(b []byte) ([]Attribute, error) { _ = "STUB: not implemented"; return nil, nil }

// Return a nil slice when there are no attributes to decode.

// An AttributeDecoder provides a safe, iterator-like, API around attribute
// decoding.
//
// It is recommend to use an AttributeDecoder where possible instead of calling
// UnmarshalAttributes and using package nlenc functions directly.
//
// The Err method must be called after the Next method returns false to determine
// if any errors occurred during iteration.
type AttributeDecoder struct {
	// ByteOrder defines a specific byte order to use when processing integer
	// attributes.  ByteOrder should be set immediately after creating the
	// AttributeDecoder: before any attributes are parsed.
	//
	// If not set, the native byte order will be used.
	ByteOrder binary.ByteOrder

	// The current attribute being worked on.
	a Attribute

	// The slice of input bytes and its iterator index.
	b []byte
	i int

	length int

	// Any error encountered while decoding attributes.
	err error
}

// NewAttributeDecoder creates an AttributeDecoder that unpacks Attributes
// from b and prepares the decoder for iteration.
func NewAttributeDecoder(b []byte) (*AttributeDecoder, error) {
	_ = "STUB: not implemented"
	return nil,

		// By default, use native byte order.
		nil
}

// Next advances the decoder to the next netlink attribute.  It returns false
// when no more attributes are present, or an error was encountered.
func (ad *AttributeDecoder) Next() bool {
	_ = "STUB: not implemented"

	// Hit an error, stop iteration.
	return false
}

// Exit if array pointer is at or beyond the end of the slice.

// Advance the pointer by at least one header's length.

// Type returns the Attribute.Type field of the current netlink attribute
// pointed to by the decoder.
//
// Type masks off the high bits of the netlink attribute type which may contain
// the Nested and NetByteOrder flags. These can be obtained by calling TypeFlags.
func (ad *AttributeDecoder) Type() uint16 {
	_ = "STUB: not implemented"
	// Mask off any flags stored in the high bits.
	return 0
}

// TypeFlags returns the two high bits of the Attribute.Type field of the current
// netlink attribute pointed to by the decoder.
//
// These bits of the netlink attribute type are used for the Nested and NetByteOrder
// flags, available as the Nested and NetByteOrder constants in this package.
func (ad *AttributeDecoder) TypeFlags() uint16 { _ = "STUB: not implemented"; return 0 }

// Len returns the number of netlink attributes pointed to by the decoder.
func (ad *AttributeDecoder) Len() int {
	_ = "STUB: not implemented"

	// count scans the input slice to count the number of netlink attributes
	// that could be decoded by Next().
	return 0
}

func (ad *AttributeDecoder) available() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Make sure there's at least a header's worth
// of data to read on each iteration.

// Extract the length of the attribute.

// Ignore zero-length attributes.

// Advance by at least a header's worth of bytes.

// data returns the Data field of the current Attribute pointed to by the decoder.
func (ad *AttributeDecoder) data() []byte {
	_ = "STUB: not implemented"

	// Err returns the first error encountered by the decoder.
	return nil
}

func (ad *AttributeDecoder) Err() error {
	_ = "STUB: not implemented"

	// Bytes returns the raw bytes of the current Attribute's data.
	return nil
}

func (ad *AttributeDecoder) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// String returns the string representation of the current Attribute's data.
func (ad *AttributeDecoder) String() string { _ = "STUB: not implemented"; return "" }

// Uint8 returns the uint8 representation of the current Attribute's data.
func (ad *AttributeDecoder) Uint8() uint8 { _ = "STUB: not implemented"; return 0 }

// Uint16 returns the uint16 representation of the current Attribute's data.
func (ad *AttributeDecoder) Uint16() uint16 { _ = "STUB: not implemented"; return 0 }

// Uint32 returns the uint32 representation of the current Attribute's data.
func (ad *AttributeDecoder) Uint32() uint32 { _ = "STUB: not implemented"; return 0 }

// Uint64 returns the uint64 representation of the current Attribute's data.
func (ad *AttributeDecoder) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

// Int8 returns the Int8 representation of the current Attribute's data.
func (ad *AttributeDecoder) Int8() int8 { _ = "STUB: not implemented"; return 0 }

// Int16 returns the Int16 representation of the current Attribute's data.
func (ad *AttributeDecoder) Int16() int16 { _ = "STUB: not implemented"; return 0 }

// Int32 returns the Int32 representation of the current Attribute's data.
func (ad *AttributeDecoder) Int32() int32 { _ = "STUB: not implemented"; return 0 }

// Int64 returns the Int64 representation of the current Attribute's data.
func (ad *AttributeDecoder) Int64() int64 { _ = "STUB: not implemented"; return 0 }

// Flag returns a boolean representing the Attribute.
func (ad *AttributeDecoder) Flag() bool { _ = "STUB: not implemented"; return false }

// Do is a general purpose function which allows access to the current data
// pointed to by the AttributeDecoder.
//
// Do can be used to allow parsing arbitrary data within the context of the
// decoder.  Do is most useful when dealing with nested attributes, attribute
// arrays, or decoding arbitrary types (such as C structures) which don't fit
// cleanly into a typical unsigned integer value.
//
// The function fn should not retain any reference to the data b outside of the
// scope of the function.
func (ad *AttributeDecoder) Do(fn func(b []byte) error) { _ = "STUB: not implemented"; return }

// Nested decodes data into a nested AttributeDecoder to handle nested netlink
// attributes. When calling Nested, the Err method does not need to be called on
// the nested AttributeDecoder.
//
// The nested AttributeDecoder nad inherits the same ByteOrder setting as the
// top-level AttributeDecoder ad.
func (ad *AttributeDecoder) Nested(fn func(nad *AttributeDecoder) error) {
	_ = "STUB: not implemented"
	// Because we are wrapping Do, there is no need to check ad.err immediately.
	return
}

// An AttributeEncoder provides a safe way to encode attributes.
//
// It is recommended to use an AttributeEncoder where possible instead of
// calling MarshalAttributes or using package nlenc directly.
//
// Errors from intermediate encoding steps are returned in the call to
// Encode.
type AttributeEncoder struct {
	// ByteOrder defines a specific byte order to use when processing integer
	// attributes.  ByteOrder should be set immediately after creating the
	// AttributeEncoder: before any attributes are encoded.
	//
	// If not set, the native byte order will be used.
	ByteOrder binary.ByteOrder

	attrs []Attribute
	err   error
}

// NewAttributeEncoder creates an AttributeEncoder that encodes Attributes.
func NewAttributeEncoder() *AttributeEncoder { _ = "STUB: not implemented"; return nil }

// Uint8 encodes uint8 data into an Attribute specified by typ.
func (ae *AttributeEncoder) Uint8(typ uint16, v uint8) { _ = "STUB: not implemented"; return }

// Uint16 encodes uint16 data into an Attribute specified by typ.
func (ae *AttributeEncoder) Uint16(typ uint16, v uint16) { _ = "STUB: not implemented"; return }

// Uint32 encodes uint32 data into an Attribute specified by typ.
func (ae *AttributeEncoder) Uint32(typ uint16, v uint32) { _ = "STUB: not implemented"; return }

// Uint64 encodes uint64 data into an Attribute specified by typ.
func (ae *AttributeEncoder) Uint64(typ uint16, v uint64) { _ = "STUB: not implemented"; return }

// Int8 encodes int8 data into an Attribute specified by typ.
func (ae *AttributeEncoder) Int8(typ uint16, v int8) { _ = "STUB: not implemented"; return }

// Int16 encodes int16 data into an Attribute specified by typ.
func (ae *AttributeEncoder) Int16(typ uint16, v int16) { _ = "STUB: not implemented"; return }

// Int32 encodes int32 data into an Attribute specified by typ.
func (ae *AttributeEncoder) Int32(typ uint16, v int32) { _ = "STUB: not implemented"; return }

// Int64 encodes int64 data into an Attribute specified by typ.
func (ae *AttributeEncoder) Int64(typ uint16, v int64) { _ = "STUB: not implemented"; return }

// Flag encodes a flag into an Attribute specified by typ.
func (ae *AttributeEncoder) Flag(typ uint16, v bool) {
	_ = "STUB: not implemented"
	// Only set flag on no previous error or v == true.
	return
}

// Flags have no length or data fields.

// String encodes string s as a null-terminated string into an Attribute
// specified by typ.
func (ae *AttributeEncoder) String(typ uint16, s string) { _ = "STUB: not implemented"; return }

// Length checking, thanks ubiquitousbyte on GitHub.

// Bytes embeds raw byte data into an Attribute specified by typ.
func (ae *AttributeEncoder) Bytes(typ uint16, b []byte) { _ = "STUB: not implemented"; return }

// Do is a general purpose function to encode arbitrary data into an attribute
// specified by typ.
//
// Do is especially helpful in encoding nested attributes, attribute arrays,
// or encoding arbitrary types (such as C structures) which don't fit cleanly
// into an unsigned integer value.
func (ae *AttributeEncoder) Do(typ uint16, fn func() ([]byte, error)) {
	_ = "STUB: not implemented"
	return
}

// Nested embeds data produced by a nested AttributeEncoder and flags that data
// with the Nested flag. When calling Nested, the Encode method should not be
// called on the nested AttributeEncoder.
//
// The nested AttributeEncoder nae inherits the same ByteOrder setting as the
// top-level AttributeEncoder ae.
func (ae *AttributeEncoder) Nested(typ uint16, fn func(nae *AttributeEncoder) error) {
	_ = "STUB: not implemented"
	// Because we are wrapping Do, there is no need to check ae.err immediately.
	return
}

// Encode returns the encoded bytes representing the attributes.
func (ae *AttributeEncoder) Encode() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
