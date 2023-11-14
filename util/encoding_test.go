package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func prepareNBits(data []byte, pointer, n int) (dataToWrite uint16) {
	offset := 0
	for i := pointer + n - 1; i >= pointer; i-- {
		dataToWrite |= uint16(data[i/8]>>(7-i%8)) & 0x0001 << offset
		offset++
	}
	return dataToWrite
}

func TestEnlarge(t *testing.T) {
	testCases := []struct {
		name   string
		bytes  []byte
		n      uint16
		expect int
	}{
		{
			name:   "test_no_exceed",
			bytes:  make([]byte, 3, 8),
			n:      3,
			expect: 8,
		},
		{
			name:   "test_no_exceed_boundary",
			bytes:  make([]byte, 3, 8),
			n:      5,
			expect: 8,
		},
		{
			name:   "test_exceed_doubling",
			bytes:  make([]byte, 3, 8),
			n:      6,
			expect: 16,
		},
		{
			name:   "test_exceed_quartering_1",
			bytes:  make([]byte, 33, 33),
			n:      1,
			expect: 33 + 33/4,
		},
		{
			name:   "test_exceed_quartering_2",
			bytes:  make([]byte, 40, 40),
			n:      30,
			expect: 77,
		},
		{
			name:   "test_exceed_direct_set",
			bytes:  make([]byte, 0, 3),
			n:      9,
			expect: 9,
		},
		{
			name:   "test_expand_zero",
			bytes:  make([]byte, 0, 0),
			n:      0,
			expect: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bs := NewBitStream(tc.bytes)
			bs.p = uint16(len(tc.bytes)) * 8
			bs.enlarge(tc.n * 8)
			assert.Equal(t, tc.expect, cap(bs.b))
		})
	}
}

func TestWriteByte1(t *testing.T) {
	testCases := []struct {
		name string
		data []byte
	}{
		{
			name: "single_byte",
			data: []byte("a"),
		},
		{
			name: "bytes_string",
			data: []byte("too young too simple sometimes naive"),
		},
		{
			name: "bytes_unicode",
			data: []byte{0x04, 0xa2, 0x03, 0xb1, 0x00, 0x2b},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bs := NewBitStream(nil)
			p := 0
			for p < len(tc.data)*8 {
				dataToWrite := prepareNBits(tc.data, p, 1)
				bs.WriteByte1(byte(dataToWrite))
				p += 1
			}
			assert.Equal(t, tc.data, bs.b)
		})
	}
}

func TestWriteByte2(t *testing.T) {
	testCases := []struct {
		name string
		data []byte
	}{
		{
			name: "single_byte",
			data: []byte("a"),
		},
		{
			name: "bytes_string",
			data: []byte("too young too simple sometimes naive"),
		},
		{
			name: "bytes_unicode",
			data: []byte{0x04, 0xa2, 0x03, 0xb1, 0x00, 0x2b},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bs := NewBitStream(nil)
			p := 0
			for p < len(tc.data)*8 {
				dataToWrite := prepareNBits(tc.data, p, 2)
				bs.WriteByte2(byte(dataToWrite))
				p += 2
			}
			assert.Equal(t, tc.data, bs.b)
		})
	}
}

func TestWriteByte4(t *testing.T) {
	testCases := []struct {
		name string
		data []byte
	}{
		{
			name: "nil_bytes",
			data: nil,
		},
		{
			name: "single_byte",
			data: []byte{'a'},
		},
		{
			name: "bytes_string",
			data: []byte("too young too simple sometimes naive"),
		},
		{
			name: "bytes_unicode",
			data: []byte{0x04, 0xa2, 0x03, 0xb1, 0x00, 0x2b},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bs := NewBitStream(nil)
			p := 0
			for p < len(tc.data)*8 {
				dataToWrite := prepareNBits(tc.data, p, 4)
				bs.WriteByte4(byte(dataToWrite))
				p += 4
			}
			assert.Equal(t, tc.data, bs.b)
		})
	}
}

func TestWriteByte6(t *testing.T) {
	testCases := []struct {
		name string
		data []byte
	}{
		{
			name: "nil_bytes",
			data: nil,
		},
		{
			name: "triple_bytes",
			data: []byte{'a', 'a', 'a'},
		},
		{
			name: "bytes_string",
			data: []byte("too young too simple sometimes naive"),
		},
		{
			name: "bytes_unicode",
			data: []byte{0x04, 0xa2, 0x03, 0xb1, 0x00, 0x2b},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bs := NewBitStream(nil)
			p := 0
			for p < len(tc.data)*8 {
				dataToWrite := prepareNBits(tc.data, p, 6)
				bs.WriteByte6(byte(dataToWrite))
				p += 6
			}
			assert.Equal(t, tc.data, bs.b)
		})
	}
}

func TestWriteByte8(t *testing.T) {
	testCases := []struct {
		name string
		data []byte
	}{
		{
			name: "nil_bytes",
			data: nil,
		},
		{
			name: "single_byte",
			data: []byte{'a'},
		},
		{
			name: "bytes_string",
			data: []byte("too young too simple sometimes naive"),
		},
		{
			name: "bytes_unicode",
			data: []byte{0x04, 0xa2, 0x03, 0xb1, 0x00, 0x2b},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bs := NewBitStream(nil)
			p := 0
			for p < len(tc.data)*8 {
				dataToWrite := prepareNBits(tc.data, p, 8)
				bs.WriteByte8(byte(dataToWrite))
				p += 8
			}
			assert.Equal(t, tc.data, bs.b)
		})
	}
}

func TestWriteUInt12(t *testing.T) {
	testCases := []struct {
		name string
		data []byte
	}{
		{
			name: "triple_bytes",
			data: []byte{'a', 'a', 'a'},
		},
		{
			name: "bytes_string",
			data: []byte("too young too simple sometimes naive"),
		},
		{
			name: "bytes_unicode",
			data: []byte{0x04, 0xa2, 0x03, 0xb1, 0x00, 0x2b},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bs := NewBitStream(nil)
			p := 0
			for p < len(tc.data)*8 {
				dataToWrite := prepareNBits(tc.data, p, 12)
				bs.WriteUInt12(dataToWrite)
				p += 12
			}
			assert.Equal(t, tc.data, bs.b)
		})
	}
}

func TestWriteUInt16(t *testing.T) {
	testCases := []struct {
		name string
		data []byte
	}{
		{
			name: "double_bytes",
			data: []byte{'a', 'a'},
		},
		{
			name: "bytes_string",
			data: []byte("too young too simple sometimes naive"),
		},
		{
			name: "bytes_unicode",
			data: []byte{0x04, 0xa2, 0x03, 0xb1, 0x00, 0x2b},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bs := NewBitStream(nil)
			p := 0
			for p < len(tc.data)*8 {
				dataToWrite := prepareNBits(tc.data, p, 16)
				bs.WriteUInt16(dataToWrite)
				p += 16
			}
			assert.Equal(t, tc.data, bs.b)
		})
	}
}

func TestArbitraryWrite(t *testing.T) {
	testCases := []*struct {
		name    string
		data    []byte
		offsets []int
	}{
		{
			name:    "arbitrary_write_1",
			data:    []byte{0xde, 0xad, 0xbe, 0xef},
			offsets: []int{16, 12, 4},
		},
		{
			name:    "arbitrary_write_2",
			data:    []byte{0xde, 0xad, 0xbe, 0xef},
			offsets: []int{12, 12, 8},
		},
		{
			name:    "arbitrary_write_3",
			data:    []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0b01000000},
			offsets: []int{16, 12, 4, 8, 6, 6, 6, 1, 4, 1, 2},
		},
		{
			name:    "arbitrary_write_4",
			data:    []byte{0x81, 0x09, 0x75},
			offsets: []int{1, 4, 2, 4, 6, 2, 4, 1},
		},
		{
			name:    "arbitrary_write_string",
			data:    []byte("GPP framework"),
			offsets: []int{16, 16, 8, 8, 8, 4, 4, 4, 4, 6, 6, 6, 6, 4, 2, 1, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bs := NewBitStream(nil)
			p := 0
			for _, offset := range tc.offsets {
				dataToWrite := prepareNBits(tc.data, p, offset)
				switch offset {
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
				p += offset
			}

			assert.Equal(t, tc.data, bs.b)
		})
	}
}

func TestWriteFibonacciInt(t *testing.T) {
	var fibWriteTestData = []*struct {
		name   string
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
		{"write_int_1", []uint16{1}, []byte{0xc0}, nil},
		{"write_int_2", []uint16{2}, []byte{0x60}, nil},
		{"write_int_4", []uint16{4}, []byte{0xb0}, nil},
		{"write_int_5", []uint16{5}, []byte{0x18}, nil},
		{"write_int_7_12", []uint16{7, 12}, []byte{0x5d, 0x60}, nil},
		{"out_of_range", []uint16{0, 6765}, nil, fibEncodeNumOutOfRangeErr},
		{"write_int_6764", []uint16{6764}, []byte{0x55, 0x55, 0x60}, nil},
	}

	for _, test := range fibWriteTestData {
		t.Run(test.name, func(t *testing.T) {
			bs := NewBitStream(nil)
			var err error
			for _, v := range test.values {
				err = bs.WriteFibonacciInt(v)
				assert.Equal(t, test.err, err)
			}
			assert.Equal(t, test.result, bs.b)
		})
	}
}

func TestWriteIntRange(t *testing.T) {
	testCases := []*struct {
		name   string
		r      []IRange
		result []byte
		err    error
	}{
		{
			"write_int_range_1",
			[]IRange{{2, 2}, {4, 4}, {6, 9}},
			[]byte{0b00000000, 0b00110011, 0b00111011, 0b00110000},
			nil,
		},
		{
			"write_int_range_2",
			[]IRange{{2, 5}, {6, 8}, {13, 16}, {28, 39}},
			[]byte{0b00000000, 0b01001011, 0b00111110, 0b11100011, 0b00111101, 0b01100101, 0b10000000},
			nil,
		},
		{
			"write_invalid_range_1",
			[]IRange{{2, 1}, {4, 4}, {6, 9}},
			[]byte{0b00000000, 0b00110000}, // Only the size will be successfully written
			fibEncodeInvalidRange,
		},
		{
			"write_invalid_range_2",
			[]IRange{{2, 2}, {2, 4}, {6, 9}},
			// 0000 0000 0011 for size, 0 for single-int-range indicator and 011 for fibonacci encoded 2
			[]byte{0b00000000, 0b00110011},
			fibEncodeInvalidRange,
		},
		{
			"write_out_of_range",
			[]IRange{{2, 5}, {6, 8}, {13, 16}, {28, 9999}},
			// error encountered when try writing the last range
			[]byte{0b00000000, 0b01001011, 0b00111110, 0b11100011, 0b00111101, 0b01100000},
			fibEncodeNumOutOfRangeErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bs := NewBitStream(nil)
			intRange := IntRange{Range: tc.r, Size: uint16(len(tc.r))}
			err := bs.WriteIntRange(&intRange)
			assert.Equal(t, tc.result, bs.b)
			assert.Equal(t, tc.err, err)
		})
	}
}
