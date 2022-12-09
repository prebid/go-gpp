package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fibonacci test data:
// 11 = 1
// 011 = 2
// 1011 = 4
// 01011 = 7
// 101011 = 12
// 1000011 = 14
// 1597 (17) + 233 (13) + 8 (6) = 1838  (00001000000100011)
// 10946 (21) + 1597 (17) + 233 (13) + 8 (6) = 12784 (000010000001000100011)

// pad the left bits with random data to ensure we test all sorts of byte alignments.

var fibtestData = []testDefinition{
	{[]byte{0x9b}, 6, 1, ""},                 // (100110)11
	{[]byte{0x2e, 0x2e}, 11, 2, ""},          // (00101110 001)011(10)
	{[]byte{0xb0}, 0, 4, ""},                 // 1011(0000)
	{[]byte{0x29, 0x62}, 6, 7, ""},           // (0010 10)01 011(00010)
	{[]byte{0x9d, 0x6c}, 5, 12, ""},          // (1001 1)101 011(0 1100)
	{[]byte{0x24, 0x87}, 8, 14, ""},          // (0010 0100) 1000 011(1)
	{[]byte{0x58, 0x40, 0x8c}, 5, 1838, ""},  // (0101 1)000 0100 0000 1000 11(00)
	{[]byte{0x08, 0x11, 0x1a}, 0, 12784, ""}, // 0000 1000 0001 0001 0001 1(010)
}

func TestReadFibonnaciInt(t *testing.T) {
	bs := BitStream{b: testData, p: 47}
	i, err := bs.ReadFibonacciInt()
	assert.Equal(t, "error reading bit 2 of Integer(Fibonacci): expected 1 bit at bit 48, but the byte array was only 6 bytes long", err.Error())

	for _, test := range fibtestData {
		bs = BitStream{b: test.data, p: test.offset}
		i, err = bs.ReadFibonacciInt()
		assert.Nil(t, err, "Unexpected error: %v", err)
		assert.Equal(t, int(test.value), int(i))
	}
}

// Test Fibonacci(Range)
// Test data: range section 1
// 0000 0000 0101 (5 range sections) 0 (1st range singular) 01011 (7) 1 (2nd range double) 100011 ((7)+9=16)
// 10011 ((16)+6=22) 1 (3rd range double) 011 ((22)+2=24) 11 ((24)+1=25) 0 (4th range singular)
// 0100000011 ((25)+57=82) 1 (5th range double) 001001001000101010100011 ((82)+57152=57234) 01000011 ((57234)+23=57257)
// 0000 0000 0101 0 010 11 1 1 0001 1 100 11 1 0 11 11 0 010 0000 011 1 0010 0100 1000 1010 1010 0011 0100 0011
// 0000 0000 0101  0010  1111  0001  1100  1110   1111  0010 0000  0111 0010 0100 1000 1010 1010 0011 0100 0011
//   0    0    5     2     F     1     C     E      F     2    0     7    2    4    8    A    A    3    4    3
// 00 52 F1 CE F2 07 24 8A A3 43
//

func TestReadFibonacciRange(t *testing.T) {
	bs := &BitStream{b: []byte{0x00, 0x52, 0xf1, 0xce, 0xf2, 0x07, 0x24, 0x8a, 0xa3, 0x43}}

	ir, err := bs.ReadFibonacciRange()
	assert.Nil(t, err, "Unexpected error: %v", err)

	expected := &IntRange{Size: 5, Max: 57257, Range: []IRange{
		{StartID: 7, EndID: 7},
		{StartID: 16, EndID: 22},
		{StartID: 24, EndID: 25},
		{StartID: 82, EndID: 82},
		{StartID: 57234, EndID: 57257},
	}}
	assert.Equal(t, expected, ir)
}
