package uspva

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/util"
)

type USPVA struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment sections.CommonUSCoreSegment
}

func NewUSPVA(encoded string) (USPVA, error) {
	uspva := USPVA{}

	bitStream, err := util.NewBitStreamFromBase64(encoded)
	if err != nil {
		return uspva, err
	}

	coreSegment, err := sections.NewCommonUSCoreSegment(8, 0, bitStream)
	if err != nil {
		return uspva, err
	}

	uspva = USPVA{
		SectionID:   constants.SectionUSPVA,
		Value:       encoded,
		CoreSegment: coreSegment,
	}

	return uspva, nil
}

func (uspva USPVA) GetID() constants.SectionID {
	return uspva.SectionID
}

func (uspva USPVA) GetValue() string {
	return uspva.Value
}
