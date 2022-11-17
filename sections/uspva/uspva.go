package uspva

import (
	"fmt"

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

	result.Version, err = bs.ReadByte6()
	if err != nil {
		return result, fmt.Errorf("unable to set Version: %s", err.Error())
	}

	result.SharingNotice, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set SharingNotice: %s", err.Error())
	}

	result.SaleOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set SaleOptOutNotice: %s", err.Error())
	}

	result.TargetedAdvertisingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set TargetedAdvertisingOptOutNotice: %s", err.Error())
	}

	result.SaleOptOut, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set SaleOptOut: %s", err.Error())
	}

	result.TargetedAdvertisingOptOut, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set TargetedAdvertisingOptOut: %s", err.Error())
	}

	result.SensitiveDataProcessing, err = bs.ReadTwoBitField(8)
	if err != nil {
		return result, fmt.Errorf("unable to set SensitiveDataProcessing: %s", err.Error())
	}

	result.KnownChildSensitiveDataConsents, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set KnownChildSensitiveDataConsents: %s", err.Error())
	}
	result.MspaCoveredTransaction, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set MspaCoveredTransaction: %s", err.Error())
	}

	result.MspaOptOutOptionMode, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set MspaOptOutOptionMode: %s", err.Error())
	}

	result.MspaServiceProviderMode, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set MspaServiceProviderMode: %s", err.Error())
	}

	return result, nil
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
