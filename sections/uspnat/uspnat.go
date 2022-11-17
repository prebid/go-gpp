package uspnat

import (
	"fmt"

	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/util"
)

type USPNATCoreSegment struct {
	Version                             byte
	SharingNotice                       byte
	SaleOptOutNotice                    byte
	SharingOptOutNotice                 byte
	TargetedAdvertisingOptOutNotice     byte
	SensitiveDataProcessingOptOutNotice byte
	SensitiveDataLimitUseNotice         byte
	SaleOptOut                          byte
	SharingOptOut                       byte
	TargetedAdvertisingOptOut           byte
	SensitiveDataProcessing             []byte
	KnownChildSensitiveDataConsents     []byte
	PersonalDataConsents                byte
	MspaCoveredTransaction              byte
	MspaOptOutOptionMode                byte
	MspaServiceProviderMode             byte
}

type USPNATGPCSegment struct {
	Gpc byte
}

type USPNAT struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment USPNATCoreSegment
	GPCSegment  USPNATGPCSegment
}

func initUSPNATCoreSegment(bs *util.BitStream) (USPNATCoreSegment, error) {
	var result = USPNATCoreSegment{}
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

	result.SharingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set SharingOptOutNotice: %s", err.Error())
	}

	result.TargetedAdvertisingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set TargetedAdvertisingOptOutNotice: %s", err.Error())
	}

	result.SensitiveDataProcessingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set SensitiveDataProcessingOptOutNotice: %s", err.Error())
	}

	result.SensitiveDataLimitUseNotice, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set SensitiveDataLimitUseNotice: %s", err.Error())
	}

	result.SaleOptOut, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set SaleOptOut: %s", err.Error())
	}
	result.SharingOptOut, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set SharingOptOut: %s", err.Error())
	}

	result.TargetedAdvertisingOptOut, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set TargetedAdvertisingOptOut: %s", err.Error())
	}

	result.SensitiveDataProcessing, err = bs.ReadTwoBitField(12)
	if err != nil {
		return result, fmt.Errorf("unable to set SensitiveDataProcessing: %s", err.Error())
	}

	result.KnownChildSensitiveDataConsents, err = bs.ReadTwoBitField(2)
	if err != nil {
		return result, fmt.Errorf("unable to set KnownChildSensitiveDataConsents: %s", err.Error())
	}

	result.PersonalDataConsents, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set PersonalDataConsents: %s", err.Error())
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

func initUSPNATGPCSegment(bs *util.BitStream) (USPNATGPCSegment, error) {
	var result = USPNATGPCSegment{}
	var err error

	result.Gpc, err = bs.ReadByte1()
	if err != nil {
		return result, fmt.Errorf("unable to set GPC: %s", err.Error())
	}

	return result, nil
}

func NewUSPNAT(encoded string) (USPNAT, error) {
	uspnat := USPNAT{}

	bitStream, err := util.NewBitStreamFromBase64(encoded)
	if err != nil {
		return uspnat, err
	}

	coreSegment, err := initUSPNATCoreSegment(bitStream)
	if err != nil {
		return uspnat, err
	}

	gpcSegment, err := initUSPNATGPCSegment(bitStream)
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
