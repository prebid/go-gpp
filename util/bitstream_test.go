package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Define some test data

// 0000 0100 1010 0010 0000 0011 1011 0001 0000 0000 0010 1011

var testData = []byte{0x04, 0xa2, 0x03, 0xb1, 0x00, 0x2b}

type testDefinition struct {
	data   []byte // The data to feed the function
	offset uint16 // The bit offset in the byte slice to start
	value  uint64 // The value we expect the function to return (64 bit to allow for future functions that extract larger ints)
	err    string // Expected error value
}

func TestReadByte1(t *testing.T) {
	testSet := map[string]testDefinition{
		"Bit out of bounds":       {testData, 80, 0, "expected 1 bit at bit 80, but the byte array was only 6 bytes long"},
		"0 in first byte":         {testData, 2, 0, ""},  // testData 0 in first byte
		"1 in first byte":         {testData, 5, 1, ""},  // testData 1 in first byte
		"1 in last bit, 3rd byte": {testData, 23, 1, ""}, // testData 1 in last bit of third byte
		"Last bit last byte":      {testData, 47, 1, ""}, // testData 1 in last bit of last byte
	}

	for name, test := range testSet {
		t.Run(name, func(t *testing.T) {
			bs := BitStream{b: test.data, p: test.offset}
			b, err := bs.ReadByte1()
			if test.err == "" {
				assert.Nil(t, err, "Found unexpected error: %s", err)
				assert.Equal(t, byte(test.value), b)
				assert.Equal(t, test.offset+1, bs.p)
			} else {
				assert.Equal(t, test.err, err.Error())
			}
		})
	}

}

func TestReadByte4(t *testing.T) {
	testSet := map[string]testDefinition{
		"Bits overrun": {testData, 46, 0,
			"expected 4 bits to start at bit 46, but the byte array was only 6 bytes long (needs second byte)"},
		"Bits out of bounds": {testData, 80, 0,
			"expected 4 bits to start at bit 80, but the byte array was only 6 bytes long"},
		"Spans 2 bytes":     {testData, 21, 7, ""},           // testData duplicate of Offset which involves flowing over to a second byte
		"nibble aligned 1":  {testData, 12, 2, ""},           // testData duplicate of Offset which aligns with a nibble and doesn't span over multiple bytes
		"nibble aligned 2":  {testData, 44, 11, ""},          // testData duplicate of Offset which aligns with a nibble and doesn't span over multiple bytes
		"Spans 2 bytes 2":   {testData, 6, 2, ""},            // testData duplicate of Offset which involves flowing over to a second byte
		"No offset":         {[]byte{0x10}, 0, 1, ""},        // No offset
		"nibble aligned 3":  {[]byte{0x92}, 4, 2, ""},        // Offset which aligns with a nibble and doesn't span over multiple bytes
		"unaligned, 1 byte": {[]byte{0x99}, 1, 3, ""},        // Offset which doesn't align with a nibble.
		"Spans 2 bytes 3":   {[]byte{0x01, 0xe0}, 7, 15, ""}, // Offset which involves flowing over to a second byte
	}

	for name, test := range testSet {
		t.Run(name, func(t *testing.T) {
			bs := BitStream{b: test.data, p: test.offset}
			b, err := bs.ReadByte4()
			if test.err == "" {
				assert.Nil(t, err, "Found unexpected error: %s", err)
				assert.Equal(t, byte(test.value), b)
				assert.Equal(t, test.offset+4, bs.p)
			} else {
				assert.Equal(t, test.err, err.Error())
			}
		})
	}
}

func TestReadByte6(t *testing.T) {
	testSet := map[string]testDefinition{
		"Bits overrun": {testData, 46, 0,
			"expected 6 bits to start at bit 46, but the byte array was only 6 bytes long (needs second byte)"},
		"Bits out of bounds": {testData, 80, 0,
			"expected 6 bits to start at bit 80, but the byte array was only 6 bytes long"},
		"Spans 2 bytes":    {testData, 21, 29, ""},          // testData duplicate of Offset which involves flowing over to a second byte
		"Nibble aligned":   {testData, 12, 8, ""},           // testData duplicate of Offset which aligns with a nibble
		"Spans 2 bytes 2":  {testData, 6, 10, ""},           // testData duplicate of Offset which involves flowing over to a second byte
		"No offset":        {[]byte{0x10}, 0, 4, ""},        // No offset
		"nibble aligned 2": {[]byte{0x92, 0x80}, 4, 10, ""}, // Offset which aligns with a nibble
		"unaligned":        {[]byte{0x99}, 1, 76, ""},       // Offset which doesn't align with a nibble.
		"spans 2 bytes 3":  {[]byte{0x01, 0xe0}, 7, 60, ""}, // Offset which involves flowing over to a second byte
	}

	for name, test := range testSet {
		t.Run(name, func(t *testing.T) {
			bs := BitStream{b: test.data, p: test.offset}
			b, err := bs.ReadByte6()
			if test.err == "" {
				assert.Nil(t, err, "Found unexpected error: %s", err)
				assert.Equal(t, byte(test.value), b)
				assert.Equal(t, test.offset+6, bs.p)
			} else {
				assert.Equal(t, test.err, err.Error())
			}
		})
	}
}

func TestReadByte8(t *testing.T) {
	// Used https://cryptii.com/ to convert 8 bit sequeces to integers
	testSet := map[string]testDefinition{
		"Bits overrun": {[]byte{0x44, 0x76}, 11, 0,
			"expected 8 bits to start at bit 11, but the byte array was only 2 bytes long"},
		"Bits out of bounds": {[]byte{0x44, 0x76}, 18, 0,
			"expected 8 bits to start at bit 18, but the byte array was only 2 bytes long"},
		"nibble aligned": {testData, 4, 0x4a, ""}, // Offset that alligns to a nibble
		"odd offset":     {testData, 7, 81, ""},   // Odd Offset
		"even offset":    {testData, 26, 196, ""}, // Even offset that does not align to a nibble
		"even offset 2":  {testData, 6, 40, ""},   // Second even offset that does not align to a nibble
		"zero offset":    {testData, 8, 162, ""},  // Zero offset
	}

	for name, test := range testSet {
		t.Run(name, func(t *testing.T) {
			bs := BitStream{b: test.data, p: test.offset}
			b, err := bs.ReadByte8()
			if test.err == "" {
				assert.Nil(t, err, "Found unexpected error: %s", err)
				assert.Equal(t, byte(test.value), b)
				assert.Equal(t, test.offset+8, bs.p)
			} else {
				assert.Equal(t, test.err, err.Error())
			}
		})
	}
}

func TestReadUInt12(t *testing.T) {
	testSet := map[string]testDefinition{
		"Bytes overrun": {testData, 44, 0,
			"expected a 12-bit int to start at bit 44, but the byte array was only 6 bytes long"},
		"Bits overrun": {testData, 40, 0,
			"expected a 12-bit int to start at bit 40, but the byte array was only 6 bytes long"},
		"Bits out of bounds": {testData, 80, 0,
			"expected a 12-bit int to start at bit 80, but the byte array was only 6 bytes long"},
		"Even offset, 2 bytes": {testData, 10, 2176, ""}, // Even Offset that does not align to a nibble, but fits 2 bytes
		"Zero offset":          {testData, 16, 59, ""},   // Zero Offset
		"Odd offset, 3 bytes":  {testData, 19, 472, ""},  // Odd Offset that overflows to 3rd byte
		"Odd offset, 2 bytes":  {testData, 1, 148, ""},   // Odd offset that fits 2 bytes
		"Even offset, 3 bytes": {testData, 22, 3780, ""}, // Another even unaligned offset that overflows to 3rd byte
		"Nibble aligned":       {testData, 4, 1186, ""},  // Offset that aligns to a nibble (these can never overflow)
		"Corner case":          {testData, 36, 0x2b, ""}, // Corner Case
	}

	for name, test := range testSet {
		t.Run(name, func(t *testing.T) {
			bs := BitStream{b: test.data, p: test.offset}
			i, err := bs.ReadUInt12()
			if test.err == "" {
				assert.Nil(t, err, "Found unexpected error: %s", err)
				assert.Equal(t, uint16(test.value), i)
				assert.Equal(t, test.offset+12, bs.p)
			} else {
				assert.Equal(t, test.err, err.Error())
			}
		})
	}
}

func TestReadUInt16(t *testing.T) {
	testSet := map[string]testDefinition{
		"Bytes overrun": {testData, 44, 0,
			"expected a 16-bit int to start at bit 44, but the byte array was only 6 bytes long"},
		"Bits overrun": {testData, 40, 0,
			"expected a 16-bit int to start at bit 40, but the byte array was only 6 bytes long"},
		"Bits out of bounds": {testData, 80, 0,
			"expected a 16-bit int to start at bit 80, but the byte array was only 6 bytes long"},
		"Even offset":    {testData, 10, 34830, ""}, // Even offset that does not align to a nibble
		"Zero offset":    {testData, 16, 945, ""},   // Zero offset
		"Odd offset":     {testData, 19, 7560, ""},  // Odd offset
		"Odd offset 2":   {testData, 1, 2372, ""},   // Odd offset
		"Even offset 2":  {testData, 22, 60480, ""}, // Second even offset that does not align to a nibble
		"Nibble aligned": {testData, 4, 18976, ""},  // Nibble aligned offset
	}

	for name, test := range testSet {
		t.Run(name, func(t *testing.T) {
			bs := BitStream{b: test.data, p: test.offset}
			i, err := bs.ReadUInt16()
			if test.err == "" {
				assert.Nil(t, err, "Found unexpected error: %s", err)
				assert.Equal(t, uint16(test.value), i)
				assert.Equal(t, test.offset+16, bs.p)
			} else {
				assert.Equal(t, test.err, err.Error())
			}
		})
	}
}
