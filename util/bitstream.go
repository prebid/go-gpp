package util

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
)

type BitStream struct {
	p uint16 // position
	b []byte
}

// NewBitStream creates a new butstream object
func NewBitStream(b []byte) *BitStream {
	return &BitStream{p: 0, b: b}
}

// NewBitStreamFromBase64 creates a new bit stream object from a base64 encoded string
func NewBitStreamFromBase64(encoded string) (*BitStream, error) {
	buff := []byte(encoded)
	decoded := make([]byte, base64.RawURLEncoding.DecodedLen(len(buff)))
	n, err := base64.RawURLEncoding.Decode(decoded, buff)
	if err != nil {
		return nil, err
	}
	decoded = decoded[:n:n]
	return NewBitStream(decoded), nil
}

// GetPosition reads out the position of the bit pointer in the bit stream
func (bs *BitStream) GetPosition() uint16 {
	return bs.p
}

// SetPosition sets the position of the bit pointer in the bit stream
func (bs *BitStream) SetPosition(pos uint16) {
	bs.p = pos
}

// Len returns the number of bytes in the BitStream
func (bs *BitStream) Len() uint16 {
	return uint16(len(bs.b))
}

// ReadByte1 reads 1 bit fron the bitstream, advancing the pointer
func (bs *BitStream) ReadByte1() (byte, error) {
	b, err := ParseByte1(bs.b, bs.p)
	if err == nil {
		bs.p = bs.p + 1
		return b, nil
	}
	return 0, err
}

// ParseByte1 parses 1 bit of data from the data array, starting at the given index
func ParseByte1(data []byte, bitStartIndex uint16) (byte, error) {
	startByte := bitStartIndex / 8
	bitOffset := bitStartIndex % 8
	if uint16(len(data)) < (startByte + 1) {
		return 0, fmt.Errorf("ParseByte1 expected 1 bits to start at bit %d, but the byte array was only %d bytes long", bitStartIndex, len(data))
	}

	return (data[startByte] & (0x80 >> bitOffset)) >> (7 - bitOffset), nil
}

// ReadByte4 reads 4 bits fron the bitstream, advancing the pointer
func (bs *BitStream) ReadByte4() (byte, error) {
	b, err := ParseByte4(bs.b, bs.p)
	if err == nil {
		bs.p = bs.p + 4
		return b, nil
	}
	return 0, err
}

// ParseByte4 parses 4 bits of data from the data array, starting at the given index
func ParseByte4(data []byte, bitStartIndex uint16) (byte, error) {
	startByte := bitStartIndex / 8
	bitStartOffset := bitStartIndex % 8
	if bitStartOffset < 5 {
		if uint16(len(data)) < (startByte + 1) {
			return 0, fmt.Errorf("ParseByte4 expected 4 bits to start at bit %d, but the byte array was only %d bytes long", bitStartIndex, len(data))
		}
		return (data[startByte] & (0xf0 >> bitStartOffset)) >> (4 - bitStartOffset), nil
	}
	if uint16(len(data)) < (startByte+2) && bitStartOffset > 4 {
		return 0, fmt.Errorf("ParseByte4 expected 4 bits to start at bit %d, but the byte array was only %d bytes long (needs second byte)", bitStartIndex, len(data))
	}

	leftBits := (data[startByte] & (0xf0 >> bitStartOffset)) << (bitStartOffset - 4)
	bitsConsumed := 8 - bitStartOffset
	overflow := 4 - bitsConsumed
	rightBits := (data[startByte+1] & (0xf0 << (4 - overflow))) >> (8 - overflow)
	return leftBits | rightBits, nil
}

// ReadByte6 reads 6 bits fron the bitstream, advancing the pointer
func (bs *BitStream) ReadByte6() (byte, error) {
	b, err := ParseByte6(bs.b, bs.p)
	if err == nil {
		bs.p = bs.p + 6
		return b, nil
	}
	return 0, err
}

// ParseByte6 parses 6 bits of data from the data array, starting at the given index
func ParseByte6(data []byte, bitStartIndex uint16) (byte, error) {
	startByte := bitStartIndex / 8
	bitStartOffset := bitStartIndex % 8
	if bitStartOffset < 3 {
		if uint16(len(data)) < (startByte + 1) {
			return 0, fmt.Errorf("ParseByte6 expected 6 bits to start at bit %d, but the byte array was only %d bytes long", bitStartIndex, len(data))
		}
		return data[startByte] >> (2 - bitStartIndex), nil
	}
	if uint16(len(data)) < (startByte + 2) {
		return 0, fmt.Errorf("ParseByte6 expected 6 bits to start at bit %d, but the byte array was only %d bytes long", bitStartIndex, len(data))
	}

	leftBits := (data[startByte] & (0xfc >> bitStartOffset)) << (bitStartOffset - 2)
	bitsConsumed := 8 - bitStartOffset
	overflow := 6 - bitsConsumed
	// If the overflow is negative, rightBits get shifted out of existance.
	rightBits := (data[startByte+1] & (0xfc << (6 - overflow))) >> (8 - overflow)
	return leftBits | rightBits, nil
}

// ReadByte8 reads 8 bits fron the bitstream, advancing the pointer
func (bs *BitStream) ReadByte8() (byte, error) {
	b, err := ParseByte8(bs.b, bs.p)
	if err == nil {
		bs.p = bs.p + 8
		return b, nil
	}
	return 0, err
}

// ParseByte8 parses 8 bits of data from the data array, starting at the given index
func ParseByte8(data []byte, bitStartIndex uint16) (byte, error) {
	startByte := bitStartIndex / 8
	bitStartOffset := bitStartIndex % 8
	if bitStartOffset == 0 {
		if uint16(len(data)) < (startByte + 1) {
			return 0, fmt.Errorf("ParseByte8 expected 8 bits to start at bit %d, but the byte array was only %d bytes long", bitStartIndex, len(data))
		}
		return data[startByte], nil
	}
	if uint16(len(data)) < (startByte + 2) {
		return 0, fmt.Errorf("ParseByte8 expected 8 bits to start at bit %d, but the byte array was only %d bytes long", bitStartIndex, len(data))
	}

	leftBits := (data[startByte] & (0xff >> bitStartOffset)) << bitStartOffset
	shiftComplement := 8 - bitStartOffset
	rightBits := (data[startByte+1] & (0xff << shiftComplement)) >> shiftComplement
	return leftBits | rightBits, nil
}

// ReadUInt12 reads 12 bits fron the bitstream, advancing the pointer
func (bs *BitStream) ReadUInt12() (uint16, error) {
	i, err := ParseUInt12(bs.b, bs.p)
	if err == nil {
		bs.p = bs.p + 12
		return i, nil
	}
	return 0, err
}

// ParseUInt12 parses 12 bits of data fromt the data array, starting at the given index
func ParseUInt12(data []byte, bitStartIndex uint16) (uint16, error) {
	end := bitStartIndex + 12
	endByte := end / 8
	endOffset := end % 8

	if endOffset > 0 {
		endByte++
	}
	if uint16(len(data)) < endByte {
		return 0, fmt.Errorf("ParseUInt12 expected a 12-bit int to start at bit %d, but the byte array was only %d bytes long",
			bitStartIndex, len(data))
	}

	leftByte, err := ParseByte4(data, bitStartIndex)
	if err != nil {
		return 0, fmt.Errorf("ParseUInt12 error on left byte: %s", err)
	}
	rightByte, err := ParseByte8(data, bitStartIndex+4)
	if err != nil {
		return 0, fmt.Errorf("ParseUInt12 error on right byte: %s", err)
	}
	return binary.BigEndian.Uint16([]byte{leftByte, rightByte}), nil
}

// ReadUInt16 reads 16 bits fron the bitstream, advancing the pointer
func (bs *BitStream) ReadUInt16() (uint16, error) {
	i, err := ParseUInt16(bs.b, bs.p)
	if err == nil {
		bs.p = bs.p + 16
		return i, nil
	}
	return 0, err
}

// ParseUInt16  parses a 16-bit integer from the data array, starting at the given index
func ParseUInt16(data []byte, bitStartIndex uint16) (uint16, error) {
	startByte := bitStartIndex / 8
	bitStartOffset := bitStartIndex % 8
	if bitStartOffset == 0 {
		if uint16(len(data)) < (startByte + 2) {
			return 0, fmt.Errorf("ParseUInt16 expected a 16-bit int to start at bit %d, but the byte array was only %d bytes long", bitStartIndex, len(data))
		}
		return binary.BigEndian.Uint16(data[startByte : startByte+2]), nil
	}
	if uint16(len(data)) < (startByte + 3) {
		return 0, fmt.Errorf("ParseUInt16 expected a 16-bit int to start at bit %d, but the byte array was only %d bytes long", bitStartIndex, len(data))
	}

	leftByte, err := ParseByte8(data, bitStartIndex)
	if err != nil {
		return 0, fmt.Errorf("ParseUInt16 error on left byte: %s", err)
	}
	rightByte, err := ParseByte8(data, bitStartIndex+8)
	if err != nil {
		return 0, fmt.Errorf("ParseUInt16 error on right byte: %s", err)
	}
	return binary.BigEndian.Uint16([]byte{leftByte, rightByte}), nil
}
