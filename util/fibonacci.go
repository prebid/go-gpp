package util

import "fmt"

// Preload the smaller, more common, fibonacci values to speed up lookups.
var fibLookup = [fibLen]uint16{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181}

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
func (bs *BitStream) ReadFibonacciInt() (uint16, error) {

	lastBit, err := bs.ReadByte1()
	if err != nil {
		return 0, fmt.Errorf("error reading bit 1 of Integer(Fibonacci): %s", err)
	}
	nextBit, err := bs.ReadByte1()
	if err != nil {
		return 0, fmt.Errorf("error reading bit 2 of Integer(Fibonacci): %s", err)
	}

	// First bit encodes "1" if set. (fib(2)=1)
	result := uint16(lastBit)

	// At the start of each loop, lastBit has been processed, nextBit has not. if both lastBit and
	// nextBit are set, we discard nextBit and return the result
	for i := 3; (lastBit == 0) || (nextBit == 0); i++ {
		lastBit = nextBit
		nextBit, err = bs.ReadByte1()
		if err != nil {
			return 0, fmt.Errorf("error reading bit %d of Integer(Fibonacci): %s", i, err)
		}
		if lastBit == 1 {
			result = result + fibonacci(i)
		}

	}

	return result, nil
}

// ReadFibonacciRange
func (bs *BitStream) ReadFibonacciRange() (*IntRange, error) {
	numEntries, err := bs.ReadUInt12()
	if err != nil {
		return nil, fmt.Errorf("error reading size of Range(Fibonacci): %s", err)
	}
	var maxValue uint16
	var offset uint16

	ranges := make([]IRange, numEntries)
	for i := range ranges {
		bit, err := bs.ReadByte1()
		if err != nil {
			return nil, fmt.Errorf("error reading the boolean bit of a Range(Fibonacci) entry(%d): %s", i, err)
		}
		if bit == 0 {
			offset, err := bs.ReadFibonacciInt()
			if err != nil {
				return nil, fmt.Errorf("error reading an int offset value in a Range(Fibonacci) entry(%d): %s", i, err)
			}
			entry := offset + maxValue
			ranges[i].StartID = entry
			ranges[i].EndID = entry
			if entry > maxValue {
				maxValue = entry
			}
		} else {
			// first entry is an offset from the previous entry
			offset, err = bs.ReadFibonacciInt()
			ranges[i].StartID = maxValue + offset
			if err != nil {
				return nil, fmt.Errorf("error reading first int offset value in a Range(Fibonacci) entry(%d): %s", i, err)
			}
			// Second entry in a Fibonacci range is an offset from the first.
			offset, err = bs.ReadFibonacciInt()
			if err != nil {
				return nil, fmt.Errorf("error reading second int offset value in a Range(Fibonacci) entry(%d): %s", i, err)
			}
			ranges[i].EndID = ranges[i].StartID + offset
			if ranges[i].EndID > maxValue {
				maxValue = ranges[i].EndID
			}
		}
	}

	return &IntRange{Size: numEntries, Range: ranges, Max: maxValue}, nil
}
