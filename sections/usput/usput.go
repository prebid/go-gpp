package usput

import (
	"github.com/revcontent-production/go-gpp/constants"
	"github.com/revcontent-production/go-gpp/sections"
	"github.com/revcontent-production/go-gpp/util"
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

func (segment USPUTCoreSegment) Encode(bs *util.BitStream) {
	bs.WriteByte6(segment.Version)
	bs.WriteByte2(segment.SharingNotice)
	bs.WriteByte2(segment.SaleOptOutNotice)
	bs.WriteByte2(segment.TargetedAdvertisingOptOutNotice)
	bs.WriteByte2(segment.SensitiveDataProcessingOptOutNotice)
	bs.WriteByte2(segment.SaleOptOut)
	bs.WriteByte2(segment.TargetedAdvertisingOptOut)
	bs.WriteTwoBitField(segment.SensitiveDataProcessing)
	bs.WriteByte2(segment.KnownChildSensitiveDataConsents)
	bs.WriteByte2(segment.MspaCoveredTransaction)
	bs.WriteByte2(segment.MspaOptOutOptionMode)
	bs.WriteByte2(segment.MspaServiceProviderMode)
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

func (usput USPUT) Encode(bool) []byte {
	bs := util.NewBitStreamForWrite()
	usput.CoreSegment.Encode(bs)
	return bs.Base64Encode()
}

func (usput USPUT) GetID() constants.SectionID {
	return usput.SectionID
}

func (usput USPUT) GetValue() string {
	return usput.Value
}
