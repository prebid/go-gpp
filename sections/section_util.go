package sections

import (
	"errors"
	"fmt"

	"github.com/prebid/go-gpp/util"
)

type FieldAttributes struct {
	Size         int
	Position     uint16
	NumBitFields int
}

func SetBitFieldValue(attributes FieldAttributes, bs *util.BitStream, err error) ([]byte, error) {
	result := []byte{}

	if err != nil {
		return result, err
	}

	if attributes.NumBitFields == 0 {
		return result, errors.New("attribute NumBitFields is 0")
	}

	maxFields := attributes.NumBitFields * attributes.Size
	for i := 0; i < maxFields; i += attributes.Size {
		val, err := ReadByteValue(attributes.Size, bs)
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}

	return result, nil
}

func ReadByteValue(size int, bs *util.BitStream) (byte, error) {
	switch size {
	case 1:
		return bs.ReadByte1()
	case 2:
		return bs.ReadByte2()
	case 4:
		return bs.ReadByte4()
	case 6:
		return bs.ReadByte6()
	case 8:
		return bs.ReadByte8()
	default:
		return uint8(0), fmt.Errorf("unknown field size for reading bits: %d", size)
	}
}

func SetIntValue(attributes FieldAttributes, bs *util.BitStream) (byte, error) {
	bs.SetPosition(attributes.Position)

	result, err := ReadByteValue(attributes.Size, bs)

	return result, err
}
