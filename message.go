package netlink

import (
	"errors"
	"iter"
)

// Flags which may apply to netlink attribute types when communicating with
// certain netlink families.
const (
	Nested       uint16 = 0x8000
	NetByteOrder uint16 = 0x4000

	// attrTypeMask masks off Type bits used for the above flags.
	attrTypeMask uint16 = 0x3fff
)

// Various errors which may occur when attempting to marshal or unmarshal
// a Message to and from its binary form.
var (
	errIncorrectMessageLength = errors.New("netlink message header length incorrect")
	errShortMessage           = errors.New("not enough data to create a netlink message")
	errUnalignedMessage       = errors.New("input data is not properly aligned for netlink message")
)

// HeaderFlags specify flags which may be present in a Header.
type HeaderFlags uint16

const (
	// General netlink communication flags.

	// Request indicates a request to netlink.
	Request HeaderFlags = 1

	// Multi indicates a multi-part message, terminated by Done on the
	// last message.
	Multi HeaderFlags = 2

	// Acknowledge requests that netlink reply with an acknowledgement
	// using Error and, if needed, an error code.
	Acknowledge HeaderFlags = 4

	// Echo requests that netlink echo this request back to the sender.
	Echo HeaderFlags = 8

	// DumpInterrupted indicates that a dump was inconsistent due to a
	// sequence change.
	DumpInterrupted HeaderFlags = 16

	// DumpFiltered indicates that a dump was filtered as requested.
	DumpFiltered HeaderFlags = 32

	// Flags used to retrieve data from netlink.

	// Root requests that netlink return a complete table instead of a
	// single entry.
	Root HeaderFlags = 0x100

	// Match requests that netlink return a list of all matching entries.
	Match HeaderFlags = 0x200

	// Atomic requests that netlink send an atomic snapshot of its entries.
	// Requires CAP_NET_ADMIN or an effective UID of 0.
	Atomic HeaderFlags = 0x400

	// Dump requests that netlink return a complete list of all entries.
	Dump HeaderFlags = Root | Match

	// Flags used to create objects.

	// Replace indicates request replaces an existing matching object.
	Replace HeaderFlags = 0x100

	// Excl indicates request does not replace the object if it already exists.
	Excl HeaderFlags = 0x200

	// Create indicates request creates an object if it doesn't already exist.
	Create HeaderFlags = 0x400

	// Append indicates request adds to the end of the object list.
	Append HeaderFlags = 0x800

	// Flags for extended acknowledgements.

	// Capped indicates the size of a request was capped in an extended
	// acknowledgement.
	Capped HeaderFlags = 0x100

	// AcknowledgeTLVs indicates the presence of netlink extended
	// acknowledgement TLVs in a response.
	AcknowledgeTLVs HeaderFlags = 0x200
)

// String returns the string representation of a HeaderFlags.
func (f HeaderFlags) String() string { _ = "STUB: not implemented"; return "" }

// HeaderType specifies the type of a Header.
type HeaderType uint16

const (
	// Noop indicates that no action was taken.
	Noop HeaderType = 0x1

	// Error indicates an error code is present, which is also used to indicate
	// success when the code is 0.
	Error HeaderType = 0x2

	// Done indicates the end of a multi-part message.
	Done HeaderType = 0x3

	// Overrun indicates that data was lost from this message.
	Overrun HeaderType = 0x4
)

// String returns the string representation of a HeaderType.
func (t HeaderType) String() string { _ = "STUB: not implemented"; return "" }

// NB: the memory layout of Header and Linux's syscall.NlMsgHdr must be
// exactly the same.  Cannot reorder, change data type, add, or remove fields.
// Named types of the same size (e.g. HeaderFlags is a uint16) are okay.

// A Header is a netlink header.  A Header is sent and received with each
// Message to indicate metadata regarding a Message.
type Header struct {
	// Length of a Message, including this Header.
	Length uint32

	// Contents of a Message.
	Type HeaderType

	// Flags which may be used to modify a request or response.
	Flags HeaderFlags

	// The sequence number of a Message.
	Sequence uint32

	// The port ID of the sending process.
	PID uint32
}

// A Message is a netlink message.  It contains a Header and an arbitrary
// byte payload, which may be decoded using information from the Header.
//
// Data is often populated with netlink attributes. For easy encoding and
// decoding of attributes, see the AttributeDecoder and AttributeEncoder types.
type Message struct {
	Header Header
	Data   []byte
}

// MarshalBinary marshals a Message into a byte slice.
func (m Message) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary unmarshals the contents of a byte slice into a Message.
func (m *Message) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

// Don't allow misleading length

// checkMessage checks a single Message for netlink errors.
func checkMessage(m Message) error {
	_ = "STUB: not implemented"
	// NB: All non-nil errors returned from this function *must* be of type
	// OpError in order to maintain the appropriate contract with callers of
	// this package.
	return nil
}

// The libnl documentation indicates that type error can
// contain error codes:
// https://www.infradead.org/~tgr/libnl/doc/core.html#core_errmsg.
//
// However, rtnetlink at least seems to also allow errors to occur at the
// end of a multipart message with done/multi and an error number.

// Error code followed by nlmsghdr/ext ack attributes.

// If no data, there must be no error number so just  exit early. Some
// of the unit tests hard-coded this but I don't actually know if this
// case occurs in the wild.

// Done|Multi potentially followed by ext ack attributes.

// Neither, nothing to do.

// Errno occupies 4 bytes.

// 0 indicates no error.

// Error code is a negative integer, convert it into an OS-specific raw
// system call error, but do not wrap with os.NewSyscallError to signify
// that this error was produced by a netlink message; not a system call.

// TODO(mdlayher): investigate the Capped flag.

// No extended acknowledgement.

// Flags indicate an extended acknowledgement. The type/flags combination
// checked above determines the offset where the TLVs occur.

// There is an nlmsghdr preceding the TLVs.

// The TLVs should be at the offset indicated by the nlmsghdr.length,
// plus the offset where the header began. But make sure the calculated
// offset is still in-bounds.

// There is no nlmsghdr preceding the TLVs, parse them directly.

// Malformed TLVs, just return the OpError with the info we have.

// unix.NLMSGERR_ATTR_MSG

// unix.NLMSGERR_ATTR_OFFS

// Explicitly ignore ad.Err: malformed TLVs, just return the OpError with
// the info we have.

// parseMessagesIter returns an iterator over netlink messages in b.
// Each iteration yields a Message and any error encountered during parsing.
// The iterator stops on the first error or when all messages are parsed.
//
// If b is less that NLMSG_HDRLEN bytes, no error or messages will be returned.
// The same applies for any trailing bytes whose length is less than
// NLMSG_HDRLEN.
func parseMessagesIter(b []byte) iter.Seq2[Message, error] { _ = "STUB: not implemented"; return nil }

// Return if the consumer has stopped iterating.
