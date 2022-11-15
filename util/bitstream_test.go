package util

import (
	"testing"
)

// Define some test data

// 0000 0100 1010 0010 0000 0011 1011 0001 0000 0000 0010 1011

var testdata = []byte{0x04, 0xa2, 0x03, 0xb1, 0x00, 0x2b}

type testDefinition struct {
	data   []byte // The data to feed the function
	offset uint16 // The bit offset in the byte slice to start
	value  uint64 // The value we expect the function to return (64 bit to allow for future functions that extract larger ints)
}

var test1Bit = []testDefinition{
	{testdata, 2, 0},  // testdata 0 in first byte
	{testdata, 5, 1},  // testdata 1 in first byte
	{testdata, 23, 1}, // testdata 1 in last bit of third byte
	{testdata, 47, 1}, // testdata 1 in last bit of last byte
}

func TestReadByte1(t *testing.T) {
	bs := BitStream{b: testdata, p: 80}

	b, err := bs.ReadByte1()
	assertStringsEqual(t, "Expected 1 bit at bit 80, but the byte array was only 6 bytes long", err.Error())

	for _, test := range test1Bit {
		bs = BitStream{b: test.data, p: test.offset}
		b, err = bs.ReadByte1()
		assertNilError(t, err)
		assertBytesEqual(t, byte(test.value), b)
	}

}

var test4Bits = []testDefinition{
	{testdata, 21, 7},           // testdata duplicate of Offset which involves flowing over to a second byte
	{testdata, 12, 2},           // testdata duplicate of Offset which aligns with a nibble and doesn't span over multiple bytes
	{testdata, 44, 11},          // testdata duplicate of Offset which aligns with a nibble and doesn't span over multiple bytes
	{testdata, 6, 2},            // testdata duplicate of Offset which involves flowing over to a second byte
	{[]byte{0x10}, 0, 1},        // No offset
	{[]byte{0x92}, 4, 2},        // Offset which aligns with a nibble and doesn't span over multiple bytes
	{[]byte{0x99}, 1, 3},        // Offset which doesn't align with a nibble.
	{[]byte{0x01, 0xe0}, 7, 15}, // Offset which involves flowing over to a second byte
}

func TestReadByte4(t *testing.T) {
	bs := BitStream{b: testdata, p: 46}

	b, err := bs.ReadByte4()
	assertStringsEqual(t, "Expected 4 bits to start at bit 46, but the byte array was only 6 bytes long (needs second byte)", err.Error())

	bs = BitStream{b: testdata, p: 80}
	b, err = bs.ReadByte4()
	assertStringsEqual(t, "Expected 4 bits to start at bit 80, but the byte array was only 6 bytes long", err.Error())

	for _, test := range test4Bits {
		bs = BitStream{b: test.data, p: test.offset}
		b, err = bs.ReadByte4()
		assertNilError(t, err)
		assertBytesEqual(t, byte(test.value), b)
	}
}

var test6Bits = []testDefinition{
	{testdata, 21, 29},          // testdata duplicate of Offset which involves flowing over to a second byte
	{testdata, 12, 8},           // testdata duplicate of Offset which aligns with a nibble
	{testdata, 6, 10},           // testdata duplicate of Offset which involves flowing over to a second byte
	{[]byte{0x10}, 0, 4},        // No offset
	{[]byte{0x92, 0x80}, 4, 10}, // Offset which aligns with a nibble
	{[]byte{0x99}, 1, 76},       // Offset which doesn't align with a nibble.
	{[]byte{0x01, 0xe0}, 7, 60}, // Offset which involves flowing over to a second byte
}

func TestReadByte6(t *testing.T) {
	bs := BitStream{b: testdata, p: 46}

	b, err := bs.ReadByte6()
	assertStringsEqual(t, "Expected 6 bits to start at bit 46, but the byte array was only 6 bytes long (needs second byte)", err.Error())

	bs = BitStream{b: testdata, p: 80}
	b, err = bs.ReadByte6()
	assertStringsEqual(t, "Expected 6 bits to start at bit 80, but the byte array was only 6 bytes long", err.Error())

	for _, test := range test6Bits {
		bs = BitStream{b: test.data, p: test.offset}
		b, err = bs.ReadByte6()
		assertNilError(t, err)
		assertBytesEqual(t, byte(test.value), b)
	}
}

// Used https://cryptii.com/ to convert 8 bit sequeces to integers
var test8Bits = []testDefinition{
	{testdata, 4, 0x4a}, // Offset that alligns to a nibble
	{testdata, 7, 81},   // Odd Offset
	{testdata, 26, 196}, // Even offset that does not align to a nibble
	{testdata, 6, 40},   // Second even offset that does not align to a nibble
	{testdata, 8, 162},  // Zero offset
}

func TestReadByte8(t *testing.T) {
	bs := BitStream{b: []byte{0x44, 0x76}, p: 11}
	b, err := bs.ReadByte8()
	assertStringsEqual(t, "Expected 8 bits to start at bit 11, but the byte array was only 2 bytes long", err.Error())

	bs = BitStream{b: []byte{0x44, 0x76}, p: 18}
	b, err = bs.ReadByte8()
	assertStringsEqual(t, "Expected 8 bits to start at bit 18, but the byte array was only 2 bytes long", err.Error())

	for _, test := range test8Bits {
		bs = BitStream{b: test.data, p: test.offset}
		b, err = bs.ReadByte8()
		assertNilError(t, err)
		assertBytesEqual(t, byte(test.value), b)
	}
}

var test12Bits = []testDefinition{
	{testdata, 10, 2176}, // Even Offset that does not align to a nibble, but fits 2 bytes
	{testdata, 16, 59},   // Zero Offset
	{testdata, 19, 472},  // Odd Offset that overflows to 3rd byte
	{testdata, 1, 148},   // Odd offset that fits 2 bytes
	{testdata, 22, 3780}, // Another even unaligned offset that overflows to 3rd byte
	{testdata, 4, 1186},  // Offset that aligns to a nibble (these can never overflow)
	{testdata, 36, 0x2b}, // Corner Case
}

func TestReadUInt12(t *testing.T) {
	bs := BitStream{b: testdata, p: 44}
	i, err := bs.ReadUInt12()
	assertStringsEqual(t, "Expected a 12-bit int to start at bit 44, but the byte array was only 6 bytes long", err.Error())

	bs = BitStream{b: testdata, p: 40}
	i, err = bs.ReadUInt12()
	assertStringsEqual(t, "Expected a 12-bit int to start at bit 40, but the byte array was only 6 bytes long", err.Error())

	bs = BitStream{b: testdata, p: 37}
	i, err = bs.ReadUInt12()
	assertStringsEqual(t, "Expected a 12-bit int to start at bit 37, but the byte array was only 6 bytes long", err.Error())

	for _, test := range test12Bits {
		bs = BitStream{b: test.data, p: test.offset}
		i, err = bs.ReadUInt12()
		assertNilError(t, err)
		assertUInt16sEqual(t, uint16(test.value), i)
	}
}

var test16Bits = []testDefinition{
	{testdata, 10, 34830}, // Even offset that does not align to a nibble
	{testdata, 16, 945},   // Zero offset
	{testdata, 19, 7560},  // Odd offset
	{testdata, 1, 2372},   // Odd offset
	{testdata, 22, 60480}, // Second even offset that does not align to a nibble
	{testdata, 4, 18976},  // Nibble aligned offset
}

func TestReadUInt16(t *testing.T) {
	bs := BitStream{b: testdata, p: 44}
	i, err := bs.ReadUInt16()
	assertStringsEqual(t, "Expected a 16-bit int to start at bit 44, but the byte array was only 6 bytes long", err.Error())

	bs = BitStream{b: testdata, p: 40}
	i, err = bs.ReadUInt16()
	assertStringsEqual(t, "Expected a 16-bit int to start at bit 40, but the byte array was only 6 bytes long", err.Error())

	for _, test := range test16Bits {
		bs = BitStream{b: test.data, p: test.offset}
		i, err = bs.ReadUInt16()
		assertNilError(t, err)
		assertUInt16sEqual(t, uint16(test.value), i)
	}
}

func assertNilError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func assertStringsEqual(t *testing.T, expected string, actual string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Strings were not equal. Expected %s, actual %s", expected, actual)
	}
}

func assertBytesEqual(t *testing.T, expected byte, actual byte) {
	t.Helper()
	if actual != expected {
		t.Errorf("bytes were not equal. Expected %d, actual %d", expected, actual)
	}
}

func assertUInt8sEqual(t *testing.T, expected uint8, actual uint8) {
	t.Helper()
	if actual != expected {
		t.Errorf("Ints were not equal. Expected %d, actual %d", expected, actual)
	}
}

func assertUInt16sEqual(t *testing.T, expected uint16, actual uint16) {
	t.Helper()
	if actual != expected {
		t.Errorf("Ints were not equal. Expected %d, actual %d", expected, actual)
	}
}

func assertIntsEqual(t *testing.T, expected int, actual int) {
	t.Helper()
	if actual != expected {
		t.Errorf("Ints were not equal. Expected %d, actual %d", expected, actual)
	}
}

func assertBoolsEqual(t *testing.T, expected bool, actual bool) {
	t.Helper()
	if actual != expected {
		t.Errorf("Bools were not equal. Expected %t, actual %t", expected, actual)
	}
}
