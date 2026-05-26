//go:build gofuzz
// +build gofuzz

package netlink

func fuzz(b1 []byte) int {
	_ = "STUB: not implemented"
	// 1. unmarshal, marshal, unmarshal again to check m1 and m2 for equality
	// after a round trip. checkMessage is also used because there is a fair
	// amount of tricky logic around testing for presence of error headers and
	// extended acknowledgement attributes.
	return 0
}

// 2. marshal again and compare b2 and b3 (b1 may have reserved bytes set
// which we ignore and fill with zeros when marshaling) for equality.

// 3. unmarshal any possible attributes from m1's data and marshal them
// again for comparison.
