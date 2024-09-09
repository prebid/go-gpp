package usptx

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/util"
)

type USPTXCoreSegment struct {
	Version                         byte
	ProcessingNotice                byte
	SaleOptOutNotice                byte
	TargetedAdvertisingOptOutNotice byte
	SaleOptOut                      byte
	TargetedAdvertisingOptOut       byte
	SensitiveDataProcessing         []byte
	KnownChildSensitiveDataConsents byte
	AdditionalDataProcessingConsent byte
	MspaCoveredTransaction          byte
	MspaOptOutOptionMode            byte
	MspaServiceProviderMode         byte
}

type USPTX struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment USPTXCoreSegment
	GPCSegment  sections.CommonUSGPCSegment
}

func NewUPSTXCoreSegment(bs *util.BitStream) (USPTXCoreSegment, error) {
	var usptxCore USPTXCoreSegment
	var err error

	usptxCore.Version, err = bs.ReadByte6()
	if err != nil {
		return usptxCore, sections.ErrorHelper("CoreSegment.Version", err)
	}

	usptxCore.ProcessingNotice, err = bs.ReadByte2()
	if err != nil {
		return usptxCore, sections.ErrorHelper("CoreSegment.ProcessingNotice", err)
	}

	usptxCore.SaleOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return usptxCore, sections.ErrorHelper("CoreSegment.SaleOptOutNotice", err)
	}

	usptxCore.TargetedAdvertisingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return usptxCore, sections.ErrorHelper("CoreSegment.TargetedAdvertisingOptOutNotice", err)
	}

	usptxCore.SaleOptOut, err = bs.ReadByte2()
	if err != nil {
		return usptxCore, sections.ErrorHelper("CoreSegment.SaleOptOut", err)
	}

	usptxCore.TargetedAdvertisingOptOut, err = bs.ReadByte2()
	if err != nil {
		return usptxCore, sections.ErrorHelper("CoreSegment.TargetedAdvertisingOptOut", err)
	}

	usptxCore.SensitiveDataProcessing, err = bs.ReadTwoBitField(8)
	if err != nil {
		return usptxCore, sections.ErrorHelper("CoreSegment.SensitiveDataProcessing", err)
	}

	usptxCore.KnownChildSensitiveDataConsents, err = bs.ReadByte2()
	if err != nil {
		return usptxCore, sections.ErrorHelper("CoreSegment.KnownChildSensitiveDataConsents", err)
	}

	usptxCore.AdditionalDataProcessingConsent, err = bs.ReadByte2()
	if err != nil {
		return usptxCore, sections.ErrorHelper("CoreSegment.AdditionalDataProcessingConsent", err)
	}

	usptxCore.MspaCoveredTransaction, err = bs.ReadByte2()
	if err != nil {
		return usptxCore, sections.ErrorHelper("CoreSegment.MspaCoveredTransaction", err)
	}

	usptxCore.MspaOptOutOptionMode, err = bs.ReadByte2()
	if err != nil {
		return usptxCore, sections.ErrorHelper("CoreSegment.MspaOptOutOptionMode", err)
	}

	usptxCore.MspaServiceProviderMode, err = bs.ReadByte2()
	if err != nil {
		return usptxCore, sections.ErrorHelper("CoreSegment.MspaServiceProviderMode", err)
	}

	return usptxCore, nil
}

func (segment USPTXCoreSegment) Encode(bs *util.BitStream) {
	bs.WriteByte6(segment.Version)
	bs.WriteByte2(segment.ProcessingNotice)
	bs.WriteByte2(segment.SaleOptOutNotice)
	bs.WriteByte2(segment.TargetedAdvertisingOptOutNotice)
	bs.WriteByte2(segment.SaleOptOut)
	bs.WriteByte2(segment.TargetedAdvertisingOptOut)
	bs.WriteTwoBitField(segment.SensitiveDataProcessing)
	bs.WriteByte2(segment.KnownChildSensitiveDataConsents)
	bs.WriteByte2(segment.AdditionalDataProcessingConsent)
	bs.WriteByte2(segment.MspaCoveredTransaction)
	bs.WriteByte2(segment.MspaOptOutOptionMode)
	bs.WriteByte2(segment.MspaServiceProviderMode)
}

func NewUSPTX(encoded string) (USPTX, error) {
	usptx := USPTX{}

	bitStream, gpcBitStream, err := sections.CreateBitStreams(encoded, true)
	if err != nil {
		return usptx, err
	}

	coreSegment, err := NewUPSTXCoreSegment(bitStream)
	if err != nil {
		return usptx, err
	}

	gpcSegment := sections.CommonUSGPCSegment{
		SubsectionType: 1,
		Gpc:            false,
	}

	if gpcBitStream != nil {
		gpcSegment, err = sections.NewCommonUSGPCSegment(gpcBitStream)
		if err != nil {
			return usptx, err
		}
	}
	usptx = USPTX{
		SectionID:   constants.SectionUSPTX,
		Value:       encoded,
		CoreSegment: coreSegment,
		GPCSegment:  gpcSegment,
	}

	return usptx, nil
}

func (usptx USPTX) Encode(gpcIncluded bool) []byte {
	bs := util.NewBitStreamForWrite()
	usptx.CoreSegment.Encode(bs)
	res := bs.Base64Encode()
	if !gpcIncluded {
		return res
	}
	bs.Reset()
	res = append(res, '.')
	usptx.GPCSegment.Encode(bs)
	return append(res, bs.Base64Encode()...)
}

func (usput USPTX) GetID() constants.SectionID {
	return usput.SectionID
}

func (usput USPTX) GetValue() string {
	return usput.Value
}
