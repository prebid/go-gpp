package uspct

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/util"
)

type USPCT struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment sections.CommonUSCoreSegment
}

func NewUSPCT(encoded string) (USPCT, error) {
	uspct := USPCT{}

	bitStream, err := util.NewBitStreamFromBase64(encoded)
	if err != nil {
		return uspct, err
	}

	coreSegment, err := sections.NewCommonUSCoreSegment(8, 3, bitStream)
	if err != nil {
		return uspct, err
	}

	uspct = USPCT{
		SectionID:   constants.SectionUSPCT,
		Value:       encoded,
		CoreSegment: coreSegment,
	}

	return uspct, nil
}

func (uspct USPCT) GetID() constants.SectionID {
	return uspct.SectionID
}

func (uspct USPCT) GetValue() string {
	return uspct.Value
}
