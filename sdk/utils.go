package sdk

// https://go.googlesource.com/go/+/go1.9.2/src/encoding/hex/hex.go
const hextable = "0123456789abcdef"

// EncodedLen returns the length of an encoding of n source bytes.
// Specifically, it returns n * 2.
func EncodedLen(n int) int { return n * 2 }

// Encode encodes src into EncodedLen(len(src))
// bytes of dst. As a convenience, it returns the number
// of bytes written to dst, but this value is always EncodedLen(len(src)).
// Encode implements hexadecimal encoding.
func Encode(dst, src []byte) int {
	for i, v := range src {
		dst[i*2] = hextable[v>>4]
		dst[i*2+1] = hextable[v&0x0f]
	}
	return len(src) * 2
}

func ToSelector(in []byte) uint32 {
	ret := uint32(0)
	ret += uint32(in[3])
	ret += uint32(in[2]) << 8
	ret += uint32(in[1]) << 16
	ret += uint32(in[0]) << 24
	return ret
}

func uint64ToBytes(v uint64) [8]uint8 {
	b := [8]uint8{0, 0, 0, 0, 0, 0, 0, 0}
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// big-endian
func bytesToUint64(b []uint8) uint64 {
	ret := uint64(0)
	lsbIdx := min(len(b), 8) - 1
	for i := lsbIdx; i >= 0; i-- {
		ret = ret + (uint64(b[i]) << (8 * (lsbIdx - i)))
	}
	return ret
}
