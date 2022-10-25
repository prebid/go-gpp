package util

// Preload the smaller, more common, fibonacci values to speed up lookups.
var fibLookup [20]uint16 = [20]uint16{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181}

const fibLen int = 20

// fibonacci returns the ith fibonacci number.
func fibonacci(i int) uint16 {
	if i < fibLen {
		return fibLookup[i]
	}
	return fibonacci(i-2) + fibonacci(i-1)
}

// ReadFibonacciInt parses 1 fibonacci encoded int from the data array, starting at the given index
// Note that 0 cannot be fibonacci encoded, and thus will only be seen if there is an error.
func ReadFibonacciInt(bs *BitStream) (uint16, error) {

	lastBit, err := ReadByte1(bs)
	if err != nil {
		return 0, err
	}
	nextBit, err := ReadByte1(bs)
	if err != nil {
		return 0, err
	}

	// First bit encodes "1" if set. (fib(2)=1)
	result := uint16(lastBit)

	// At the start of each loop, lastBit has been processed, nextBit has not. if both lastBit and
	// nextBit are set, we discard nextBit and return the result
	for i := 3; (lastBit == 0) || (nextBit == 0); i++ {
		lastBit = nextBit
		nextBit, err = ReadByte1(bs)
		if err != nil {
			return 0, err
		}
		if lastBit == 1 {
			result = result + fibonacci(i)
		}

	}

	return result, nil
}

// ReadFibonacciRange
func ReadFibonacciRange(bs *BitStream) (*IntRange, error) {
	numEntries, err := ReadUInt12(bs)
	if err != nil {
		return nil, err
	}
	var maxValue uint16

	ranges := make([]IRange, numEntries)
	for i := range ranges {
		bit, err := ReadByte1(bs)
		if err != nil {
			return nil, err
		}
		if bit == 0 {
			entry, err := ReadFibonacciInt(bs)
			if err != nil {
				return nil, err
			}
			ranges[i].StartID = entry
			ranges[i].EndID = entry
			if entry > maxValue {
				maxValue = entry
			}
		} else {
			ranges[i].StartID, err = ReadFibonacciInt(bs)
			if err != nil {
				return nil, err
			}
			ranges[i].EndID, err = ReadFibonacciInt(bs)
			if err != nil {
				return nil, err
			}
			if ranges[i].EndID > maxValue {
				maxValue = ranges[i].EndID
			}
		}
	}

	return &IntRange{Size: numEntries, Range: ranges, Max: maxValue}, nil
}
