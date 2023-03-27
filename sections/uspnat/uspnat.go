package uspnat

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
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

type USPNAT struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment USPNATCoreSegment
	GPCSegment  sections.CommonUSGPCSegment
}

func NewUSPNATCoreSegment(bs *util.BitStream) (USPNATCoreSegment, error) {
	var uspnatCore USPNATCoreSegment
	var err error

	uspnatCore.Version, err = bs.ReadByte6()
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.Version", err)
	}

	uspnatCore.SharingNotice, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.SharingNotice", err)
	}

	uspnatCore.SaleOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.SaleOptOutNotice", err)
	}

	uspnatCore.SharingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.SharingOptOutNotice", err)
	}

	uspnatCore.TargetedAdvertisingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.TargetedAdvertisingOptOutNotice", err)
	}

	uspnatCore.SensitiveDataProcessingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.SensitiveDataProcessingOptOutNotice", err)
	}

	uspnatCore.SensitiveDataLimitUseNotice, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.SensitiveDataLimitUseNotice", err)
	}

	uspnatCore.SaleOptOut, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.SaleOptOut", err)
	}

	uspnatCore.SharingOptOut, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.SharingOptOut", err)
	}

	uspnatCore.TargetedAdvertisingOptOut, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.TargetedAdvertisingOptOut", err)
	}

	uspnatCore.SensitiveDataProcessing, err = bs.ReadTwoBitField(12)
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.SensitiveDataProcessing", err)
	}

	uspnatCore.KnownChildSensitiveDataConsents, err = bs.ReadTwoBitField(2)
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.KnownChildSensitiveDataConsents", err)
	}

	uspnatCore.PersonalDataConsents, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.PersonalDataConsents", err)
	}

	uspnatCore.MspaCoveredTransaction, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.MspaCoveredTransaction", err)
	}

	uspnatCore.MspaOptOutOptionMode, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.MspaOptOutOptionMode", err)
	}

	uspnatCore.MspaServiceProviderMode, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, sections.ErrorHelper("CoreSegment.MspaServiceProviderMode", err)
	}

	return uspnatCore, nil
}

func (segment USPNATCoreSegment) Encode(bs *util.BitStream) {
	bs.WriteByte6(segment.Version)
	bs.WriteByte2(segment.SharingNotice)
	bs.WriteByte2(segment.SaleOptOutNotice)
	bs.WriteByte2(segment.SharingOptOutNotice)
	bs.WriteByte2(segment.TargetedAdvertisingOptOutNotice)
	bs.WriteByte2(segment.SensitiveDataProcessingOptOutNotice)
	bs.WriteByte2(segment.SensitiveDataLimitUseNotice)
	bs.WriteByte2(segment.SaleOptOut)
	bs.WriteByte2(segment.SharingOptOut)
	bs.WriteByte2(segment.TargetedAdvertisingOptOut)
	bs.WriteTwoBitField(segment.SensitiveDataProcessing)
	bs.WriteTwoBitField(segment.KnownChildSensitiveDataConsents)
	bs.WriteByte2(segment.PersonalDataConsents)
	bs.WriteByte2(segment.MspaCoveredTransaction)
	bs.WriteByte2(segment.MspaOptOutOptionMode)
	bs.WriteByte2(segment.MspaServiceProviderMode)
}

func NewUSPNAT(encoded string) (USPNAT, error) {
	uspnat := USPNAT{}

	coreBitStream, gpcBitStream, err := sections.CreateBitStreams(encoded, true)
	if err != nil {
		return uspnat, err
	}

	coreSegment, err := NewUSPNATCoreSegment(coreBitStream)
	if err != nil {
		return uspnat, err
	}

	gpcSegment := sections.CommonUSGPCSegment{
		SubsectionType: 1,
		Gpc:            false,
	}

	if gpcBitStream != nil {
		gpcSegment, err = sections.NewCommonUSGPCSegment(gpcBitStream)
		if err != nil {
			return uspnat, err
		}
	}

	uspnat = USPNAT{
		SectionID:   constants.SectionUSPNAT,
		Value:       encoded,
		CoreSegment: coreSegment,
		GPCSegment:  gpcSegment,
	}

	return uspnat, nil
}

func (uspnat USPNAT) Encode(gpcIncluded bool) []byte {
	bs := util.NewBitStreamForWrite()
	uspnat.CoreSegment.Encode(bs)
	res := bs.Base64Encode()
	if !gpcIncluded {
		return res
	}
	bs.Reset()
	res = append(res, '.')
	uspnat.GPCSegment.Encode(bs)
	return append(res, bs.Base64Encode()...)
}

func (uspnat USPNAT) GetID() constants.SectionID {
	return uspnat.SectionID
}

func (uspnat USPNAT) GetValue() string {
	return uspnat.Value
}
