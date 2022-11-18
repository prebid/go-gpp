package uspnat

import (
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

	result.Version, err = bs.ReadByteSize(6, err)
	result.SharingNotice, err = bs.ReadByteSize(2, err)
	result.SaleOptOutNotice, err = bs.ReadByteSize(2, err)
	result.SharingOptOutNotice, err = bs.ReadByteSize(2, err)
	result.TargetedAdvertisingOptOutNotice, err = bs.ReadByteSize(2, err)
	result.SensitiveDataProcessingOptOutNotice, err = bs.ReadByteSize(2, err)
	result.SensitiveDataLimitUseNotice, err = bs.ReadByteSize(2, err)
	result.SaleOptOut, err = bs.ReadByteSize(2, err)
	result.SharingOptOut, err = bs.ReadByteSize(2, err)
	result.TargetedAdvertisingOptOut, err = bs.ReadByteSize(2, err)
	result.SensitiveDataProcessing, err = bs.ReadTwoBitField(12, err)
	result.KnownChildSensitiveDataConsents, err = bs.ReadTwoBitField(2, err)
	result.PersonalDataConsents, err = bs.ReadByteSize(2, err)
	result.MspaCoveredTransaction, err = bs.ReadByteSize(2, err)
	result.MspaOptOutOptionMode, err = bs.ReadByteSize(2, err)
	result.MspaServiceProviderMode, err = bs.ReadByteSize(2, err)

	return result, err
}

func initUSPNATGPCSegment(bs *util.BitStream) (USPNATGPCSegment, error) {
	var result = USPNATGPCSegment{}
	var err error

	result.Gpc, err = bs.ReadByteSize(1, err)

	return result, err
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
