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
	GPCSegment  sections.CommonUSGPCSegment
}

func NewUSPCO(encoded string) (USPCO, error) {
	uspco := USPCO{}

	coreBitStream, gpcBitStream, err := sections.CreateBitStreams(encoded, true)
	if err != nil {
		return uspco, err
	}

	coreSegment, err := sections.NewCommonUSCoreSegment(7, 1, coreBitStream)
	if err != nil {
		return uspco, err
	}

	gpcSegment := sections.CommonUSGPCSegment{
		SubsectionType: 1,
		Gpc:            false,
	}

	if gpcBitStream != nil {
		gpcSegment, err = sections.NewCommonUSGPCSegment(gpcBitStream)
		if err != nil {
			return uspco, err
		}
	}

	uspco = USPCO{
		SectionID:   constants.SectionUSPCO,
		Value:       encoded,
		CoreSegment: coreSegment,
		GPCSegment:  gpcSegment,
	}

	return uspco, nil
}

func (uspco USPCO) Encode(gpcIncluded bool) []byte {
	bs := util.NewBitStream(nil)
	uspco.CoreSegment.Encode(bs)
	res := bs.Base64Encode()
	if !gpcIncluded {
		return res
	}
	bs.Reset()
	res = append(res, '.')
	uspco.GPCSegment.Encode(bs)
	return append(res, bs.Base64Encode()...)
}

func (uspco USPCO) GetID() constants.SectionID {
	return uspco.SectionID
}

func (uspco USPCO) GetValue() string {
	return uspco.Value
}
