package binaryin

import (
	"io"
	"strings"

	"github.com/pkg/errors"
)

// BinaryIn struct provides methods for reading in bits from io.Reader,
// either one bit at a time (as a bool), 8 bits at a time (as a byte), 16 bits at a time (as a int16),
// 32 bits at a time (as an int), or 64 bits at a time (as a int64).
// All primitive types are assumed to be represented using their golang standard representations,
// in big-endian (most significant byte first) order.
type BinaryIn struct {
	in            io.Reader // input stream
	buffer        int       // one byte buffer
	n             int       // number of bits left in buffer
	isInitialized bool      // has BinaryIn been called for first time?
}

// NewBinaryIn constructs the BinaryIn struct
func NewBinaryIn(r io.Reader) *BinaryIn {
	return &BinaryIn{in: r, buffer: 0, n: 0, isInitialized: false}
}

func (bs *BinaryIn) initialize() {
	bs.fillBuffer()
	bs.isInitialized = true
}

func (bs *BinaryIn) fillBuffer() {
	p := make([]byte, 1)
	if _, err := bs.in.Read(p); err != nil {
		if err == io.EOF {
			bs.buffer = -1 // -1 means EOF
			bs.n = -1
		} else {
			panic(err)
		}
	} else {
		bs.n = 8
		bs.buffer = int(p[0])
	}
}

// IsEmpty returns true if input stream is empty.
func (bs *BinaryIn) IsEmpty() bool {
	if !bs.isInitialized {
		bs.initialize()
	}
	return bs.buffer == -1 // -1 means EOF
}

// ReadBool reads the next bit of data from input stream and return as a bool.
func (bs *BinaryIn) ReadBool() (bool, error) {
	if bs.IsEmpty() {
		return false, errors.New("reading from empty input stream")

	}
	bs.n--
	bit := ((bs.buffer >> bs.n) & 1) == 1
	if bs.n == 0 {
		bs.fillBuffer()
	}
	return bit, nil
}

// ReadByte reads the next 8 bits from input stream and return as an 8-bit byte.
func (bs *BinaryIn) ReadByte() (byte, error) {
	if bs.IsEmpty() {
		return 0, errors.New("reading from empty input stream")
	}
	// special case when aligned byte
	if bs.n == 8 {
		x := byte(bs.buffer)
		bs.fillBuffer()
		return x, nil
	}
	// combine last n bits of current buffer with first 8-n bits of new buffer
	x := bs.buffer
	x <<= (8 - bs.n) // filled with (8-bs.n) zero bits in the right
	oldN := bs.n
	bs.fillBuffer()
	if bs.IsEmpty() {
		return 0, errors.New("reading from empty input stream")

	}
	bs.n = oldN
	// |= compound bitwise or operator  used with a variable and a constant
	// to "set" (set to 1) particular bits in a variable.
	//	x  x  x  x  x  x  x  x    variable
	//	0  0  0  0  0  0  1  1    mask
	//	----------------------
	//	x  x  x  x  x  x  1  1
	//  bits unchanged  |bits set
	x |= (bs.buffer >> bs.n) // here bs.buffer >= 0, so signed right shift equals unsigned right shift
	return byte(x), nil
}

// ReadInt reads the next 32 bits from input stream and return as a 32-bit int.
func (bs *BinaryIn) ReadInt() (int, error) {
	x := 0
	// 32 bit int equals 4 byte
	for i := 0; i < 4; i++ {
		b, err := bs.ReadByte()
		if err != nil {
			return -1, err
		}
		x <<= 8     // filled 8 zero bits in the right
		x |= int(b) // set the rightmost 8 bits to b
	}
	return x, nil
}

// ReadInt16 reads the next 16 bits from input stream and return as a 16-bit int16
func (bs *BinaryIn) ReadInt16() (int16, error) {
	var x int16
	for i := 0; i < 2; i++ {
		byteValue, err := bs.ReadByte()
		if err != nil {
			return -1, err
		}
		x <<= 8
		x |= int16(byteValue)
	}
	return x, nil
}

// ReadInt64 reads the next 64 bits from input stream and return as a 64-bit int64.
func (bs *BinaryIn) ReadInt64() (int64, error) {
	var x int64
	for i := 0; i < 2; i++ {
		intValue, err := bs.ReadInt()
		if err != nil {
			return -1, err
		}
		x <<= 32
		x |= int64(intValue)
	}
	return x, nil
}

// ReadIntR reads the next r bits from input stream and return as an r-bit int.
func (bs *BinaryIn) ReadIntR(r int) (int, error) {
	if r < 1 || r > 32 {
		return -1, errors.Errorf("illegal value of r = %d\n", r)
	}
	if r == 32 {
		return bs.ReadInt()
	}
	x := 0
	for i := 0; i < r; i++ {
		x <<= 1
		bit, err := bs.ReadBool()
		if err != nil {
			return -1, err
		}
		if bit {
			x |= 1
		}
	}
	return x, nil
}

// ReadString reads the remaining bytes of data from input stream and return as a string.
func (bs *BinaryIn) ReadString() (string, error) {
	if bs.IsEmpty() {
		return "", errors.New("reading from empty input stream")
	}
	var s strings.Builder
	for !bs.IsEmpty() {
		b, _ := bs.ReadByte()
		s.WriteByte(b)
	}
	return s.String(), nil
}
