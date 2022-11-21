package uspco

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/util"
)

type USPCO struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment sections.CommonUSCoreSegment
}

func NewUSPCO(encoded string) (USPCO, error) {
	uspco := USPCO{}

	bitStream, err := util.NewBitStreamFromBase64(encoded)
	if err != nil {
		return uspco, err
	}

	coreSegment, err := sections.NewCommonUSCoreSegment(7, 0, bitStream)
	if err != nil {
		return uspco, err
	}

	uspco = USPCO{
		SectionID:   constants.SectionUSPCO,
		Value:       encoded,
		CoreSegment: coreSegment,
	}

	return uspco, nil
}

func (uspco USPCO) GetID() constants.SectionID {
	return uspco.SectionID
}

func (uspco USPCO) GetValue() string {
	return uspco.Value
}
