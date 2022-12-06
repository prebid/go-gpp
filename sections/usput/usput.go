package usput

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/util"
)

type USPUTCoreSegment struct {
	Version                             byte
	SharingNotice                       byte
	SaleOptOutNotice                    byte
	TargetedAdvertisingOptOutNotice     byte
	SensitiveDataProcessingOptOutNotice byte
	SaleOptOut                          byte
	TargetedAdvertisingOptOut           byte
	SensitiveDataProcessing             []byte
	KnownChildSensitiveDataConsents     byte
	MspaCoveredTransaction              byte
	MspaOptOutOptionMode                byte
	MspaServiceProviderMode             byte
}

type USPUT struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment USPUTCoreSegment
}

func NewUPSUTCoreSegment(bs *util.BitStream) (USPUTCoreSegment, error) {
	var usputCore USPUTCoreSegment
	var err error

	usputCore.Version, err = bs.ReadByte6()
	if err != nil {
		return usputCore, sections.ErrorHelper("CoreSegment.Version", err)
	}

	usputCore.SharingNotice, err = bs.ReadByte2()
	if err != nil {
		return usputCore, sections.ErrorHelper("CoreSegment.SharingNotice", err)
	}

	usputCore.SaleOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return usputCore, sections.ErrorHelper("CoreSegment.SaleOptOutNotice", err)
	}

	usputCore.TargetedAdvertisingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return usputCore, sections.ErrorHelper("CoreSegment.TargetedAdvertisingOptOutNotice", err)
	}

	usputCore.SensitiveDataProcessingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return usputCore, sections.ErrorHelper("CoreSegment.SensitiveDataProcessingOptOutNotice", err)
	}

	usputCore.SaleOptOut, err = bs.ReadByte2()
	if err != nil {
		return usputCore, sections.ErrorHelper("CoreSegment.SaleOptOut", err)
	}

	usputCore.TargetedAdvertisingOptOut, err = bs.ReadByte2()
	if err != nil {
		return usputCore, sections.ErrorHelper("CoreSegment.TargetedAdvertisingOptOut", err)
	}

	usputCore.SensitiveDataProcessing, err = bs.ReadTwoBitField(8)
	if err != nil {
		return usputCore, sections.ErrorHelper("CoreSegment.SensitiveDataProcessing", err)
	}

	usputCore.KnownChildSensitiveDataConsents, err = bs.ReadByte2()
	if err != nil {
		return usputCore, sections.ErrorHelper("CoreSegment.KnownChildSensitiveDataConsents", err)
	}

	usputCore.MspaCoveredTransaction, err = bs.ReadByte2()
	if err != nil {
		return usputCore, sections.ErrorHelper("CoreSegment.MspaCoveredTransaction", err)
	}

	usputCore.MspaOptOutOptionMode, err = bs.ReadByte2()
	if err != nil {
		return usputCore, sections.ErrorHelper("CoreSegment.MspaOptOutOptionMode", err)
	}

	usputCore.MspaServiceProviderMode, err = bs.ReadByte2()
	if err != nil {
		return usputCore, sections.ErrorHelper("CoreSegment.MspaServiceProviderMode", err)
	}

	return usputCore, nil
}

func NewUSPUT(encoded string) (USPUT, error) {
	usput := USPUT{}

	bitStream, err := util.NewBitStreamFromBase64(encoded)
	if err != nil {
		return usput, err
	}

	coreSegment, err := NewUPSUTCoreSegment(bitStream)
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
