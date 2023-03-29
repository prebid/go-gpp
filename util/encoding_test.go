package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

type encodingTestDefinition struct {
	description string
	data        []byte
}

func writeNBits(bs *BitStream, data []byte, pointer, n int) (newPointer int) {
	var dataToWrite uint16
	offset := 0
	for i := pointer + n - 1; i >= pointer; i-- {
		dataToWrite |= uint16(data[i/8]>>(7-i%8)) & 0x0001 << offset
		offset++
	}

	switch n {
	case 1:
		bs.WriteByte1(byte(dataToWrite))
	case 2:
		bs.WriteByte2(byte(dataToWrite))
	case 4:
		bs.WriteByte4(byte(dataToWrite))
	case 6:
		bs.WriteByte6(byte(dataToWrite))
	case 8:
		bs.WriteByte8(byte(dataToWrite))
	case 12:
		bs.WriteUInt12(dataToWrite)
	case 16:
		bs.WriteUInt16(dataToWrite)
	}
	return pointer + n
}

var testCases = []*encodingTestDefinition{
	{
		"Test 1",
		[]byte("too young too simple sometimes naive"),
	},
	{
		"Test 2",
		[]byte("I shall dedicate myself to the interests of the country in life and death"),
	},
	{
		"Test 3",
		testData,
	},
}

func TestWriteByte1(t *testing.T) {
	for no, c := range testCases {
		bs := NewBitStream(nil)
		p := 0
		for p < len(c.data)*8 {
			p = writeNBits(bs, c.data, p, 1)
		}
		assert.Equal(t, c.data, bs.b, fmt.Sprintf("test case %d [%s] failed", no, c.description))
	}
}

func TestWriteByte2(t *testing.T) {
	for no, c := range testCases {
		bs := NewBitStream(nil)
		p := 0
		for p < len(c.data)*8 {
			p = writeNBits(bs, c.data, p, 2)
		}
		assert.Equal(t, c.data, bs.b, fmt.Sprintf("test case %d [%s] failed", no, c.description))
	}
}

func TestWriteByte4(t *testing.T) {
	for no, c := range testCases {
		bs := NewBitStream(nil)
		p := 0
		for p < len(c.data)*8 {
			p = writeNBits(bs, c.data, p, 4)
		}
		assert.Equal(t, c.data, bs.b, fmt.Sprintf("test case %d [%s] failed", no, c.description))
	}
}

func TestWriteByte6(t *testing.T) {
	for no, c := range testCases {
		bs := NewBitStream(nil)
		// Cut the data so that the number of bits of the data is a multiple of 6.
		// The least common multiple of 6 and 8 is 24, which is 3 bytes.
		c.data = c.data[:len(c.data)/3*3]
		p := 0
		for p < len(c.data)*8 {
			p = writeNBits(bs, c.data, p, 6)
		}
		assert.Equal(t, c.data, bs.b, fmt.Sprintf("test case %d [%s] failed", no, c.description))
	}
}

func TestWriteByte8(t *testing.T) {
	for no, c := range testCases {
		bs := NewBitStream(nil)
		p := 0
		for p < len(c.data)*8 {
			p = writeNBits(bs, c.data, p, 8)
		}
		assert.Equal(t, c.data, bs.b, fmt.Sprintf("test case %d [%s] failed", no, c.description))
	}
}

func TestWriteUInt12(t *testing.T) {
	for no, c := range testCases {
		bs := NewBitStream(nil)
		c.data = c.data[:len(c.data)/3*3] // the least common multiple of 12 and 8 is 24, which is 3 bytes.
		p := 0
		for p < len(c.data)*8 {
			p = writeNBits(bs, c.data, p, 12)
		}
		assert.Equal(t, c.data, bs.b, fmt.Sprintf("test case %d [%s] failed", no, c.description))
	}
}

func TestWriteUInt16(t *testing.T) {
	for no, c := range testCases {
		bs := NewBitStream(nil)
		c.data = c.data[:len(c.data)/2*2] // the least common multiple of 16 and 8 is 16, which is 2 bytes.
		p := 0
		for p < len(c.data)*8 {
			p = writeNBits(bs, c.data, p, 16)
		}
		assert.Equal(t, c.data, bs.b, fmt.Sprintf("test case %d [%s] failed", no, c.description))
	}
}

func TestArbitraryWrite(t *testing.T) {
	choices := []int{1, 2, 4, 6, 8, 12, 16}
	rand.Seed(time.Now().UnixNano())
	for no, c := range testCases {
		bs := NewBitStream(nil)
		p := 0
		N := len(c.data) * 8
		for p < N {
			rest := N - p
			for choices[len(choices)-1] > rest {
				choices = choices[:len(choices)-1]
			}
			p = writeNBits(bs, c.data, p, choices[rand.Intn(len(choices))])
		}

		assert.Equal(t, c.data, bs.b, fmt.Sprintf("test case %d [%s] failed", no, c.description))
	}
}

func TestWriteFibonacciInt(t *testing.T) {
	var fibWriteTestData = []*struct {
		values []uint16
		result []byte
		err    error
	}{
		// 1 = 11
		// 2 = 011
		// 4 = 1011
		// 5 = 00011
		// 7 = 01011
		// 12 = 101011
		// 6764 = 0101 0101 0101 0101 011
		{[]uint16{1}, []byte{0xc0}, nil},
		{[]uint16{2}, []byte{0x60}, nil},
		{[]uint16{4}, []byte{0xb0}, nil},
		{[]uint16{5}, []byte{0x18}, nil},
		{[]uint16{7, 12}, []byte{0x5d, 0x60}, nil},
		{[]uint16{0, 6765}, nil, fibEncodeNumOutOfRangeErr},
		{[]uint16{6764}, []byte{0x55, 0x55, 0x60}, nil},
	}

	for _, test := range fibWriteTestData {
		bs := NewBitStream(nil)
		var err error
		for _, v := range test.values {
			err = bs.WriteFibonacciInt(v)
			assert.Equal(t, test.err, err)
		}
		assert.Equal(t, test.result, bs.b)
	}
}

func TestWriteIntRange(t *testing.T) {
	cases := []*struct {
		r      []IRange
		result []byte
		err    error
	}{
		{
			[]IRange{{2, 2}, {4, 4}, {6, 9}},
			[]byte{0b00000000, 0b00110011, 0b00111011, 0b00110000},
			nil,
		},
		{
			[]IRange{{2, 5}, {6, 8}, {13, 16}, {28, 39}},
			[]byte{0b00000000, 0b01001011, 0b00111110, 0b11100011, 0b00111101, 0b01100101, 0b10000000},
			nil,
		},
		{
			[]IRange{{2, 1}, {4, 4}, {6, 9}},
			nil,
			fibEncodeInvalidRange,
		},
		{
			[]IRange{{2, 2}, {2, 4}, {6, 9}},
			nil,
			fibEncodeInvalidRange,
		},
		{
			[]IRange{{2, 5}, {6, 8}, {13, 16}, {28, 9999}},
			nil,
			fibEncodeNumOutOfRangeErr,
		},
	}

	var err error
	for _, c := range cases {
		bs := NewBitStream(nil)
		intRange := IntRange{Range: c.r, Size: uint16(len(c.r))}
		err = bs.WriteIntRange(&intRange)
		assert.Equal(t, c.err, err)
		if err == nil {
			assert.Equal(t, c.result, bs.b)
		}
	}
}
