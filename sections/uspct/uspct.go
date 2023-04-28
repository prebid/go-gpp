package uspct

import (
	"github.com/revcontent-production/go-gpp/constants"
	"github.com/revcontent-production/go-gpp/sections"
)

type USPCT struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment sections.CommonUSCoreSegment
	GPCSegment  sections.CommonUSGPCSegment
}

func NewUSPCT(encoded string) (USPCT, error) {
	uspct := USPCT{}

	coreBitStream, gpcBitStream, err := sections.CreateBitStreams(encoded, true)
	if err != nil {
		return uspct, err
	}

	coreSegment, err := sections.NewCommonUSCoreSegment(8, 3, coreBitStream)
	if err != nil {
		return uspct, err
	}

	gpcSegment := sections.CommonUSGPCSegment{
		SubsectionType: 1,
		Gpc:            false,
	}

	if gpcBitStream != nil {
		gpcSegment, err = sections.NewCommonUSGPCSegment(gpcBitStream)
		if err != nil {
			return uspct, err
		}
	}

	uspct = USPCT{
		SectionID:   constants.SectionUSPCT,
		Value:       encoded,
		CoreSegment: coreSegment,
		GPCSegment:  gpcSegment,
	}

	return uspct, nil
}

func (uspct USPCT) GetID() constants.SectionID {
	return uspct.SectionID
}

func (uspct USPCT) GetValue() string {
	return uspct.Value
}
