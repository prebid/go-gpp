package uspnat

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/util"
)

type USPNAT struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment sections.USPNATCoreSegment
	GPCSegment  sections.CommonUSGPCSegment
}

func NewUSPNAT(encoded string) (USPNAT, error) {
	uspnat := USPNAT{}

	bitStream, err := util.NewBitStreamFromBase64(encoded)
	if err != nil {
		return uspnat, err
	}

	coreSegment, err := sections.NewUSPNATCoreSegment(bitStream)
	if err != nil {
		return uspnat, err
	}

	gpcSegment, err := sections.NewCommonUSGPCSegment(bitStream)
	if err != nil {
		return uspnat, err
	}

	uspnat = USPNAT{
		SectionID:   constants.SectionUSPNAT,
		Value:       encoded,
		CoreSegment: coreSegment,
		GPCSegment:  gpcSegment,
	}

	return uspnat, nil
}

func (uspnat USPNAT) GetID() constants.SectionID {
	return uspnat.SectionID
}

func (uspnat USPNAT) GetValue() string {
	return uspnat.Value
}
