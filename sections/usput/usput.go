package usput

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/util"
)

type USPUT struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment sections.USPUTCoreSegment
}

func NewUSPUT(encoded string) (USPUT, error) {
	usput := USPUT{}

	bitStream, err := util.NewBitStreamFromBase64(encoded)
	if err != nil {
		return usput, err
	}

	coreSegment, err := sections.NewUPSUTCoreSegment(bitStream)
	if err != nil {
		return usput, err
	}

	usput = USPUT{
		SectionID:   constants.SectionUSPUT,
		Value:       encoded,
		CoreSegment: coreSegment,
	}

	return usput, nil
}

func (usput USPUT) GetID() constants.SectionID {
	return usput.SectionID
}

func (usput USPUT) GetValue() string {
	return usput.Value
}
