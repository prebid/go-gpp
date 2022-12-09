package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test data: range section 1
// 0000 0000 0101 (5 range sections) 0 (1st range singular) 0000 0000 0000 0111 (7) 1 (2nd range double) 0000 0000 0001 0000 (16)
// 0000 0000 0001 0110 (22) 1 (3rd range double) 0000 0000 0001 1000 (24) 0000 0000 0001 1001 (25) 0 (4th range singular)
// 0000 0000 0101 0010 (82) 1 (5th range double) 1101 1111 1001 0010 (57234) 1101 1111 1010 1001 (57257)

// 0000 0000 0101 0 0000 0000 0000 0111 1 0000 0000 0001 0000 0000 0000 0001 0110 1 0000 0000 0001 1000 0000 0000 0001 1001 0 0000 0000 0101 0010 1 1101 1111 1001 0010 1101 1111 1010 1001
// base64 encoded: AFAAPABAAFoAMAAyAFLvyW_UgA

// Note that ReadIntRange uses functions from bitstream to read the parts. Those functions are tested against random
// offsets, so we don't have to test this for various offsets as well.

func TestReadIntRange(t *testing.T) {
	encoded := "AFAAPABAAFoAMAAyAFLvyW_UgA"
	bs, err := NewBitStreamFromBase64(encoded)

	assert.Nil(t, err, "Unexpected error: %v", err)

	ir, err := bs.ReadIntRange()
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
