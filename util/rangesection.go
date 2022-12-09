package util

import "fmt"

// ReadIntRange parses a Range(Int) and returns an IntRange struct
func (bs *BitStream) ReadIntRange() (*IntRange, error) {
	numEntries, err := bs.ReadUInt12()
	if err != nil {
		return nil, fmt.Errorf("error reading size of Range(Int): %s", err)
	}
	var maxValue uint16

	ranges := make([]IRange, numEntries)
	for i := range ranges {
		bit, err := bs.ReadByte1()
		if err != nil {
			return nil, fmt.Errorf("error reading the boolean bit of a Range(Int) entry: %s", err)
		}
		if bit == 0 {
			entry, err := bs.ReadUInt16()
			if err != nil {
				return nil, fmt.Errorf("error reading an int value in a Range(Int) entry: %s", err)
			}
			ranges[i].StartID = entry
			ranges[i].EndID = entry
			if entry > maxValue {
				maxValue = entry
			}
		} else {
			ranges[i].StartID, err = bs.ReadUInt16()
			if err != nil {
				return nil, fmt.Errorf("error reading first int value in a Range(Int) entry: %s", err)
			}
			ranges[i].EndID, err = bs.ReadUInt16()
			if err != nil {
				return nil, fmt.Errorf("error reading second int value in a Range(Int) entry: %s", err)
			}
			if ranges[i].EndID > maxValue {
				maxValue = ranges[i].EndID
			}
		}
	}

	return &IntRange{Size: numEntries, Range: ranges, Max: maxValue}, nil
}

type IntRange struct {
	Size  uint16
	Range []IRange
	Max   uint16
}

type IRange struct {
	StartID uint16
	EndID   uint16
}

// IsSet checks to see if an ID is contained within a range set
func (ir *IntRange) IsSet(id uint16) bool {
	if id > ir.Max {
		return false
	}

	for i := range ir.Range {
		if ir.Range[i].Contains(id) {
			return true
		}
	}
	return false
}

// Contains checks to see if an ID is contained within a range
func (r IRange) Contains(id uint16) bool {
	return r.StartID <= id && r.EndID >= id
}
