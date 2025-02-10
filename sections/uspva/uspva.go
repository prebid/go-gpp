package uspva

import (
	"github.com/revcontent-production/go-gpp/constants"
	"github.com/revcontent-production/go-gpp/sections"
	"github.com/revcontent-production/go-gpp/util"
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

	// NOTE: VA has only a single field in the KnownChildSensitiveDataConsents array. It otherwise
	// matches the common core segment fields, so is being generated as a one element slice to keep
	// it consistent with the majority of states. We would like to keep the code as common and
	// consistent as possible across the different privacy constructs.
	coreSegment, err := sections.NewCommonUSCoreSegment(8, 1, bitStream)
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

func (uspva USPVA) Encode(bool) []byte {
	bs := util.NewBitStreamForWrite()
	uspva.CoreSegment.Encode(bs)
	return bs.Base64Encode()
}

func (uspva USPVA) GetID() constants.SectionID {
	return uspva.SectionID
}

func (uspva USPVA) GetValue() string {
	return uspva.Value
}
