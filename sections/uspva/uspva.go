package uspva

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/util"
)

type USPVACoreSegment struct {
	Version                         byte
	SharingNotice                   byte
	SaleOptOutNotice                byte
	TargetedAdvertisingOptOutNotice byte
	SaleOptOut                      byte
	TargetedAdvertisingOptOut       byte
	SensitiveDataProcessing         []byte
	KnownChildSensitiveDataConsents byte
	MspaCoveredTransaction          byte
	MspaOptOutOptionMode            byte
	MspaServiceProviderMode         byte
}

type USPVA struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment USPVACoreSegment
}

func initUSPVACoreSegment(bs *util.BitStream) (USPVACoreSegment, error) {
	var result = USPVACoreSegment{}
	var err error

	result.Version, err = bs.ReadByteSize(6, err)
	result.SharingNotice, err = bs.ReadByteSize(2, err)
	result.SaleOptOutNotice, err = bs.ReadByteSize(2, err)
	result.TargetedAdvertisingOptOutNotice, err = bs.ReadByteSize(2, err)
	result.SaleOptOut, err = bs.ReadByteSize(2, err)
	result.TargetedAdvertisingOptOut, err = bs.ReadByteSize(2, err)
	result.SensitiveDataProcessing, err = bs.ReadTwoBitField(8, err)
	result.KnownChildSensitiveDataConsents, err = bs.ReadByteSize(2, err)
	result.MspaCoveredTransaction, err = bs.ReadByteSize(2, err)
	result.MspaOptOutOptionMode, err = bs.ReadByteSize(2, err)
	result.MspaServiceProviderMode, err = bs.ReadByteSize(2, err)

	return result, err
}

func NewUSPVA(encoded string) (USPVA, error) {
	uspva := USPVA{}

	bitStream, err := util.NewBitStreamFromBase64(encoded)
	if err != nil {
		return uspva, err
	}

	coreSegment, err := initUSPVACoreSegment(bitStream)
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
