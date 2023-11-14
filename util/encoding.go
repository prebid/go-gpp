package util

import (
	"encoding/base64"
	"errors"
)

var (
	fibEncodeNumOutOfRangeErr = errors.New("the number to be encoded is out of range")
	fibEncodeInvalidRange     = errors.New("the range is invalid")
)

func getByteSlice() []byte {
	// Most strings to be encoded are less than 8 bytes.
	return make([]byte, 0, 8)
}

func NewBitStreamForWrite() *BitStream {
	return NewBitStream(getByteSlice())
}

// enlarge the underlying byte slice.
// This function assumes bs.p always points to the end of the stream.
// Do NOT attempt to modify the bs.p while applying this method.
func (bs *BitStream) enlarge(n uint16) {
	final := int(bs.p+n+7) / 8
	// Under most circumstances, the value of final will hit this if condition.
	if final <= cap(bs.b) {
		bs.b = bs.b[:final]
		return
	}

	newCap := cap(bs.b)
	doubleCap := newCap + newCap
	if final > doubleCap {
		newCap = final
	} else {
		if newCap < 32 {
			newCap = doubleCap
		} else {
			for newCap < final {
				newCap += newCap / 4
			}
		}
	}
	temp := make([]byte, final, newCap)
	copy(temp, bs.b)
	bs.b = temp
}

// appendNBits appends n bits in b from left to right to the BitStream.
func (bs *BitStream) appendNBits(b []byte, n uint16) {
	bs.enlarge(n)
	var (
		i         uint16
		byteIndex uint16
		offset    uint16
	)
	for i = 0; i < n; i++ {
		byteIndex = bs.p / 8
		offset = bs.p % 8
		bs.b[byteIndex] |= (b[i/8] << (i % 8)) & 0x80 >> offset
		bs.p++
	}
}

// Base64Encode applies base64 encoding on the data in buffer.
func (bs *BitStream) Base64Encode() []byte {
	encoded := make([]byte, base64.RawURLEncoding.EncodedLen(len(bs.b)))
	base64.RawURLEncoding.Encode(encoded, bs.b)
	return encoded
}

// Reset clears all the data a BitStream holds.
func (bs *BitStream) Reset() {
	bs.b = getByteSlice()
	bs.p = 0
}

// WriteByte1 writes the rightmost bit in b into the BitStream.
func (bs *BitStream) WriteByte1(b byte) {
	bs.appendNBits([]byte{b << 7}, 1)
}

// WriteByte2 writes the rightmost 2 bits in b into the BitStream.
func (bs *BitStream) WriteByte2(b byte) {
	bs.appendNBits([]byte{b << 6}, 2)
}

// WriteByte4 writes the rightmost 4 bits in b into the BitStream.
func (bs *BitStream) WriteByte4(b byte) {
	bs.appendNBits([]byte{b << 4}, 4)
}

// WriteByte6 writes the rightmost 6 bits in b into the BitStream.
func (bs *BitStream) WriteByte6(b byte) {
	bs.appendNBits([]byte{b << 2}, 6)
}

// WriteByte8 writes a full byte b into the BitStream.
func (bs *BitStream) WriteByte8(b byte) {
	bs.appendNBits([]byte{b}, 8)
}

// WriteUInt12 writes the rightmost 12 bits in b into the BitStream.
// For instance, the input b is 0xff01, the effective part refers to the rightmost 12 bits,
// which should be 0xf01.
func (bs *BitStream) WriteUInt12(b uint16) {
	first := byte(b >> 4)
	second := byte(b << 4)
	bs.appendNBits([]byte{first, second}, 12)
}

// WriteUInt16 writes a full uint16 (two bytes) into the BitStream.
func (bs *BitStream) WriteUInt16(b uint16) {
	first := byte(b >> 8)
	second := byte(b)
	bs.appendNBits([]byte{first, second}, 16)
}

// WriteTwoBitField encapsulates WriteByte2 to get convenience to some extent, when encoding GPP strings.
func (bs *BitStream) WriteTwoBitField(bList []byte) {
	for _, b := range bList {
		bs.WriteByte2(b)
	}
}

/*
	WriteFibonacciInt writes int based on Fibonacci Coding https://en.wikipedia.org/wiki/Fibonacci_coding.

By definition, Fibonacci numbers are numbers that are the sum of the two Fibonacci numbers that come before.
The numbers are [excluding 0, 1,] 1 (0+1), 2 (1+1), 3 (1+2), 5 (2+3), 8 (3+5), etc. All of the integers that
are larger and equal to 0 are able to split into a combination of at least one Fibonacci numbers. For instance,
6 refers to 1+5, while 7 refers to 2+5, where 1, 2, 5 are Fibonacci numbers.

To encode an integer using Fibonacci Coding, just follow these steps:
1.Find the largest Fibonacci number equal to or less than N; subtract this number from N, keeping track of the remainder.
2.If the number subtracted was the i-th Fibonacci number F(i), put a 1 in place i−2 in the code word (counting the left-most digit as place 0).
3.Repeat the previous steps, substituting the remainder for N, until a remainder of 0 is reached.
4.Place an additional 1 after the rightmost digit in the code word.

So a Fibonacci encoded bit sequence is all about marking the corresponding bit as 1 while leaving the others as 0.
Here are some examples:

	FibonacciInt(1) = 11			1-> (1*1)->1
	FibonacciInt(2) = 011 			2-> (1*1+2*1)->01
	FibonacciInt(3) = 0011			3-> (1*0+2*0+3*1)->001
	FibonacciInt(4) = 1011			4-> (1*1+2*0+3*1)->101
	FibonacciInt(12) = 101011		12->(1*1+2*0+3*1+5*0+8*1)->10101

According to the algorithm, no sequential ones can be found in the bit sequence, and the final bit is always a single 1,
which makes it easy to extract several continuous Fibonacci encoded integers out from a bit stream if we
place an additional 1 in the end, since 11 could be an end-of-sequence indicator. The advantage of this is,
that a sequence of numbers can be encoded into a sequence of bits without the need to know the length of the
bit sequences in advance.
*/
func (bs *BitStream) WriteFibonacciInt(num uint16) error {
	// The num should be [1,6765). Actually once the num is larger than or equal to 987,
	// the efficiency of Fibonacci Encoding would be no better than 'WriteUint16'.
	if num <= 0 || num >= fibLookup[fibLen-1]+fibLookup[fibLen-2] {
		return fibEncodeNumOutOfRangeErr
	}
	// Binary Search to find the largest fibonacci number less than or equal to num.
	lo, hi := 2, fibLen-1
	for lo < hi {
		mid := (lo + hi) / 2
		// Since lo < hi and hi is at most equal to fibLen-1,
		// the largest lo+hi is (fibLen-2 + fibLen-1),
		// the largest mid should be fibLen-2 ((2*fibLen-3)/2 => [2*(fibLen-1)-1]/2 => fibLen-2),
		// the mid+1 will never be larger than fibLen-1 and an index of mid+1 is safe here.
		if num >= fibLookup[mid] && num < fibLookup[mid+1] {
			lo = mid
			break
		}
		if num < fibLookup[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	// Calculate the fibonacci-encoded sequence.
	var fibEncoded uint32 = 1
	offset := 1
	for i := lo; i >= 2; i-- {
		if num >= fibLookup[i] {
			num -= fibLookup[i]
			fibEncoded |= 1 << offset
		}
		offset++
	}
	encodedLength := lo
	fibEncoded <<= 32 - encodedLength

	bs.appendNBits([]byte{
		byte(fibEncoded >> 24),
		byte(fibEncoded >> 16),
		byte(fibEncoded >> 8),
		byte(fibEncoded),
	}, uint16(encodedLength))
	return nil
}

/*
WriteIntRange writes an IntRange object into the BitStream.

The basic idea is 12 bits for size, 1 bit indicating if a range refers to a single number or not, followed by
one or two Fibonacci encoded numbers, and then followed by other ranges fell behind.

For example, an IntRange consisting of [1, 1], [3, 4] would be encoded into 000000000010 0 11 1 0011 11.
There are two ranges, so the size is encoded into 000000000010. The first range [1, 1] refers to a single integer 1,
so the indicator is 0, followed by Fibonacci encoded 1, while the next range [3, 4] refers to a range containing
not only one integer, so the indicator is 1, followed by Fibonacci encoded 3 and 1 (4-3).
*/
func (bs *BitStream) WriteIntRange(intRange *IntRange) error {
	var err error
	bs.WriteUInt12(intRange.Size)
	// Assume that the ranges are ordered.
	var prevID uint16
	for _, r := range intRange.Range {
		if r.EndID < r.StartID || prevID >= r.StartID {
			return fibEncodeInvalidRange
		}
		if r.StartID == r.EndID {
			bs.WriteByte1(0)
			err = bs.WriteFibonacciInt(r.StartID - prevID)
			if err != nil {
				return err
			}
		} else {
			bs.WriteByte1(1)
			err = bs.WriteFibonacciInt(r.StartID - prevID)
			if err != nil {
				return err
			}
			err = bs.WriteFibonacciInt(r.EndID - r.StartID)
			if err != nil {
				return err
			}
		}
		prevID = r.EndID
	}
	return nil
}
