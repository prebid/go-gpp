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

func (uspct USPCT) Encode(gpcIncluded bool) []byte {
	bs := util.NewBitStream(nil)
	uspct.CoreSegment.Encode(bs)
	res := bs.Base64Encode()
	if !gpcIncluded {
		return res
	}
	bs.Reset()
	res = append(res, '.')
	uspct.GPCSegment.Encode(bs)
	return append(res, bs.Base64Encode()...)
}

func (uspct USPCT) GetID() constants.SectionID {
	return uspct.SectionID
}

func (uspct USPCT) GetValue() string {
	return uspct.Value
}
