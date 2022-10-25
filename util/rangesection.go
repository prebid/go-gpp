package util

// ReadIntRange parses a Range(Int) and returns an IntRange struct
func ReadIntRange(bs *BitStream) (*IntRange, error) {
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
			entry, err := ReadUInt16(bs)
			if err != nil {
				return nil, err
			}
			ranges[i].StartID = entry
			ranges[i].EndID = entry
			if entry > maxValue {
				maxValue = entry
			}
		} else {
			ranges[i].StartID, err = ReadUInt16(bs)
			if err != nil {
				return nil, err
			}
			ranges[i].EndID, err = ReadUInt16(bs)
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
	if id < 0 || id > ir.Max {
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
