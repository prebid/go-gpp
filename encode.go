package gpp

import (
	"fmt"
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/util"
	"sort"
	"strings"
)

const (
	// the first 6 bits of the header must always evaluate to the interger '3'.
	gppType    byte = 0x3
	gppVersion byte = 0x1
)

var (
	duplicatedSectionErr = fmt.Errorf("duplicated sections")
)

func Encode(sections []Section) (string, error) {
	bs := util.NewBitStreamForWrite()
	builder := strings.Builder{}

	bs.WriteByte6(gppType)
	bs.WriteByte6(gppVersion)

	sort.Slice(sections, func(i, j int) bool {
		return sections[i].GetID() < sections[j].GetID()
	})
	// Generate int range object.
	intRange := new(util.IntRange)
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
	intRange.Size = uint16(len(sections))

	err := bs.WriteIntRange(intRange)
	if err != nil {
		return "", fmt.Errorf("write int range error: %v", err)
	}

	builder.Write(bs.Base64Encode())

	for _, sec := range sections {
		// TODO: add a parameter to decide whether a GPC segment should be included.
		builder.WriteByte('~')
		builder.Write(sec.Encode(true))
	}

	return builder.String(), nil
}
