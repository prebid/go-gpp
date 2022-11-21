package uspca

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/util"
)

type USPCA struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment sections.USPCACoreSegment
	GPCSegment  sections.CommonUSGPCSegment
}

func NewUSPCA(encoded string) (USPCA, error) {
	uspca := USPCA{}

	bitStream, err := util.NewBitStreamFromBase64(encoded)
	if err != nil {
		return uspca, err
	}

	coreSegment, err := sections.NewUSPCACoreSegment(bitStream)
	if err != nil {
		return uspca, err
	}

	gpcSegment, err := sections.NewCommonUSGPCSegment(bitStream)
	if err != nil {
		return uspca, err
	}

	uspca = USPCA{
		SectionID:   constants.SectionUSPCA,
		Value:       encoded,
		CoreSegment: coreSegment,
		GPCSegment:  gpcSegment,
	}

	return uspca, nil
}

func (uspca USPCA) GetID() constants.SectionID {
	return uspca.SectionID
}

func (uspca USPCA) GetValue() string {
	return uspca.Value
}
