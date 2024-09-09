package uspor

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/util"
)

type USPORCoreSegment struct {
	Version                         byte
	ProcessingNotice                byte
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

type USPOR struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment USPORCoreSegment
	GPCSegment  sections.CommonUSGPCSegment
}

func NewUPSORCoreSegment(bs *util.BitStream) (USPORCoreSegment, error) {
	var usporCore USPORCoreSegment
	var err error

	usporCore.Version, err = bs.ReadByte6()
	if err != nil {
		return usporCore, sections.ErrorHelper("CoreSegment.Version", err)
	}

	usporCore.ProcessingNotice, err = bs.ReadByte2()
	if err != nil {
		return usporCore, sections.ErrorHelper("CoreSegment.ProcessingNotice", err)
	}

	usporCore.SaleOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return usporCore, sections.ErrorHelper("CoreSegment.SaleOptOutNotice", err)
	}

	usporCore.TargetedAdvertisingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return usporCore, sections.ErrorHelper("CoreSegment.TargetedAdvertisingOptOutNotice", err)
	}

	usporCore.SaleOptOut, err = bs.ReadByte2()
	if err != nil {
		return usporCore, sections.ErrorHelper("CoreSegment.SaleOptOut", err)
	}

	usporCore.TargetedAdvertisingOptOut, err = bs.ReadByte2()
	if err != nil {
		return usporCore, sections.ErrorHelper("CoreSegment.TargetedAdvertisingOptOut", err)
	}

	usporCore.SensitiveDataProcessing, err = bs.ReadTwoBitField(11)
	if err != nil {
		return usporCore, sections.ErrorHelper("CoreSegment.SensitiveDataProcessing", err)
	}

	usporCore.KnownChildSensitiveDataConsents, err = bs.ReadTwoBitField(3)
	if err != nil {
		return usporCore, sections.ErrorHelper("CoreSegment.KnownChildSensitiveDataConsents", err)
	}

	usporCore.AdditionalDataProcessingConsent, err = bs.ReadByte2()
	if err != nil {
		return usporCore, sections.ErrorHelper("CoreSegment.AdditionalDataProcessingConsent", err)
	}

	usporCore.MspaCoveredTransaction, err = bs.ReadByte2()
	if err != nil {
		return usporCore, sections.ErrorHelper("CoreSegment.MspaCoveredTransaction", err)
	}

	usporCore.MspaOptOutOptionMode, err = bs.ReadByte2()
	if err != nil {
		return usporCore, sections.ErrorHelper("CoreSegment.MspaOptOutOptionMode", err)
	}

	usporCore.MspaServiceProviderMode, err = bs.ReadByte2()
	if err != nil {
		return usporCore, sections.ErrorHelper("CoreSegment.MspaServiceProviderMode", err)
	}

	return usporCore, nil
}

func (segment USPORCoreSegment) Encode(bs *util.BitStream) {
	bs.WriteByte6(segment.Version)
	bs.WriteByte2(segment.ProcessingNotice)
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

func NewUSPOR(encoded string) (USPOR, error) {
	uspor := USPOR{}

	bitStream, gpcBitStream, err := sections.CreateBitStreams(encoded, true)
	if err != nil {
		return uspor, err
	}

	coreSegment, err := NewUPSORCoreSegment(bitStream)
	if err != nil {
		return uspor, err
	}

	gpcSegment := sections.CommonUSGPCSegment{
		SubsectionType: 1,
		Gpc:            false,
	}

	if gpcBitStream != nil {
		gpcSegment, err = sections.NewCommonUSGPCSegment(gpcBitStream)
		if err != nil {
			return uspor, err
		}
	}

	uspor = USPOR{
		SectionID:   constants.SectionUSPOR,
		Value:       encoded,
		CoreSegment: coreSegment,
		GPCSegment:  gpcSegment,
	}

	return uspor, nil
}

func (uspor USPOR) Encode(gpcIncluded bool) []byte {
	bs := util.NewBitStreamForWrite()
	uspor.CoreSegment.Encode(bs)
	res := bs.Base64Encode()
	if !gpcIncluded {
		return res
	}
	bs.Reset()
	res = append(res, '.')
	uspor.GPCSegment.Encode(bs)
	return append(res, bs.Base64Encode()...)
}

func (uspor USPOR) GetID() constants.SectionID {
	return uspor.SectionID
}

func (uspor USPOR) GetValue() string {
	return uspor.Value
}
