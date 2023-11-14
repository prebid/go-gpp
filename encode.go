package gpp

import (
	"errors"
	"fmt"
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/util"
	"sort"
	"strings"
)

const (
	// the first 6 bits of the header must always evaluate to the integer '3'.
	gppType    byte = 0x3
	gppVersion byte = 0x1
	// the range of SectionID must start with 1 and end with the maximum value represented by uint16.
	minSectionId constants.SectionID = 1
	maxSectionId constants.SectionID = 0xffff
)

var (
	sectionIdOutOfRangeErr = errors.New("section ID out of range")
	duplicatedSectionErr   = errors.New("duplicated sections")
)

func Encode(sections []Section) (string, error) {
	bs := util.NewBitStreamForWrite()
	builder := strings.Builder{}

	bs.WriteByte6(gppType)
	bs.WriteByte6(gppVersion)

	sort.Slice(sections, func(i, j int) bool {
		return sections[i].GetID() < sections[j].GetID()
	})

	if len(sections) > 0 && (sections[0].GetID() < minSectionId ||
		sections[len(sections)-1].GetID() > maxSectionId) {
		return "", sectionIdOutOfRangeErr
	}
	// Generate int range object.
	intRange := new(util.IntRange)
	// Since the minimum sectionID is 1, the previous one should start with -1, which makes it not continuous.
	var prevID constants.SectionID = -1
	for _, sec := range sections {
		id := sec.GetID()
		if id == prevID {
			return "", duplicatedSectionErr
		}
		if prevID+1 == id {
			intRange.Range[len(intRange.Range)-1].EndID = uint16(id)
		} else {
			intRange.Range = append(intRange.Range, util.IRange{StartID: uint16(id), EndID: uint16(id)})
		}
		prevID = id
	}
	intRange.Size = uint16(len(intRange.Range))

	err := bs.WriteIntRange(intRange)
	if err != nil {
		return "", fmt.Errorf("write int range error: %v", err)
	}

	builder.Write(bs.Base64Encode())

	for _, sec := range sections {
		builder.WriteByte('~')
		// By default, GPP is included.
		builder.Write(sec.Encode(true))
	}

	return builder.String(), nil
}
