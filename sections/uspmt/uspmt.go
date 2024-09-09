package uspmt

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/util"
)

type USPMTCoreSegment struct {
	Version                         byte
	SharingNotice                   byte
	SaleOptOutNotice                byte
	TargetedAdvertisingOptOutNotice byte
	SaleOptOut                      byte
	TargetedAdvertisingOptOut       byte
	SensitiveDataProcessing         []byte
	KnownChildSensitiveDataConsents []byte
	AdditionalDataProcessingConsent byte
	MspaCoveredTransaction          byte
	MspaOptOutOptionMode            byte
	MspaServiceProviderMode         byte
}

func NewUSMTCoreSegment(bs *util.BitStream) (USPMTCoreSegment, error) {
	var usmt USPMTCoreSegment
	var err error

	usmt.Version, err = bs.ReadByte6()
	if err != nil {
		return usmt, sections.ErrorHelper("USMTSegment.Version", err)
	}

	usmt.SharingNotice, err = bs.ReadByte2()
	if err != nil {
		return usmt, sections.ErrorHelper("USMTSegment.SharingNotice", err)
	}

	usmt.SaleOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return usmt, sections.ErrorHelper("USMTSegment.SaleOptOutNotice", err)
	}

	usmt.TargetedAdvertisingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return usmt, sections.ErrorHelper("USMTSegment.TargetedAdvertisingOptOutNotice", err)
	}

	usmt.SaleOptOut, err = bs.ReadByte2()
	if err != nil {
		return usmt, sections.ErrorHelper("USMTSegment.SaleOptOut", err)
	}

	usmt.TargetedAdvertisingOptOut, err = bs.ReadByte2()
	if err != nil {
		return usmt, sections.ErrorHelper("USMTSegment.TargetedAdvertisingOptOut", err)
	}

	usmt.SensitiveDataProcessing, err = bs.ReadTwoBitField(8)
	if err != nil {
		return usmt, sections.ErrorHelper("USMTSegment.SensitiveDataProcessing", err)
	}

	usmt.KnownChildSensitiveDataConsents, err = bs.ReadTwoBitField(3)
	if err != nil {
		return usmt, sections.ErrorHelper("USMTSegment.KnownChildSensitiveDataConsentsArr", err)
	}

	usmt.AdditionalDataProcessingConsent, err = bs.ReadByte2()
	if err != nil {
		return usmt, sections.ErrorHelper("USMTSegment.AdditionalDataProcessingConsent", err)
	}

	usmt.MspaCoveredTransaction, err = bs.ReadByte2()
	if err != nil {
		return usmt, sections.ErrorHelper("USMTSegment.MspaCoveredTransaction", err)
	}

	usmt.MspaOptOutOptionMode, err = bs.ReadByte2()
	if err != nil {
		return usmt, sections.ErrorHelper("USMTSegment.MspaOptOutOptionMode", err)
	}

	usmt.MspaServiceProviderMode, err = bs.ReadByte2()
	if err != nil {
		return usmt, sections.ErrorHelper("USMTSegment.MspaServiceProviderMode", err)
	}

	return usmt, nil
}

func (segment USPMTCoreSegment) Encode(bs *util.BitStream) {
	bs.WriteByte6(segment.Version)
	bs.WriteByte2(segment.SharingNotice)
	bs.WriteByte2(segment.SaleOptOutNotice)
	bs.WriteByte2(segment.TargetedAdvertisingOptOutNotice)
	bs.WriteByte2(segment.SaleOptOut)
	bs.WriteByte2(segment.TargetedAdvertisingOptOut)
	bs.WriteTwoBitField(segment.SensitiveDataProcessing)
	bs.WriteTwoBitField(segment.KnownChildSensitiveDataConsents)
	bs.WriteByte2(segment.AdditionalDataProcessingConsent)
	bs.WriteByte2(segment.MspaCoveredTransaction)
	bs.WriteByte2(segment.MspaOptOutOptionMode)
	bs.WriteByte2(segment.MspaServiceProviderMode)
}

type USPMT struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment USPMTCoreSegment
	GPCSegment  sections.CommonUSGPCSegment
}

func NewUSPMT(encoded string) (USPMT, error) {
	uspmt := USPMT{}

	coreBitStream, gpcBitStream, err := sections.CreateBitStreams(encoded, true)
	if err != nil {
		return uspmt, err
	}

	coreSegment, err := NewUSMTCoreSegment(coreBitStream)
	if err != nil {
		return uspmt, err
	}

	gpcSegment := sections.CommonUSGPCSegment{
		SubsectionType: 1,
		Gpc:            false,
	}

	if gpcBitStream != nil {
		gpcSegment, err = sections.NewCommonUSGPCSegment(gpcBitStream)
		if err != nil {
			return uspmt, err
		}
	}

	uspmt = USPMT{
		SectionID:   constants.SectionUSPMT,
		Value:       encoded,
		CoreSegment: coreSegment,
		GPCSegment:  gpcSegment,
	}

	return uspmt, nil
}

func (uspmt USPMT) Encode(gpcIncluded bool) []byte {
	bs := util.NewBitStreamForWrite()
	uspmt.CoreSegment.Encode(bs)
	res := bs.Base64Encode()
	if !gpcIncluded {
		return res
	}
	bs.Reset()
	res = append(res, '.')
	uspmt.GPCSegment.Encode(bs)
	return append(res, bs.Base64Encode()...)
}

func (uspmt USPMT) GetID() constants.SectionID {
	return uspmt.SectionID
}

func (uspmt USPMT) GetValue() string {
	return uspmt.Value
}
