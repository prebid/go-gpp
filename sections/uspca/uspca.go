package uspca

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/util"
)

type USPCACoreSegment struct {
	Version                         byte
	SaleOptOutNotice                byte
	SharingOptOutNotice             byte
	SensitiveDataLimitUseNotice     byte
	SaleOptOut                      byte
	SharingOptOut                   byte
	SensitiveDataProcessing         []byte
	KnownChildSensitiveDataConsents []byte
	PersonalDataConsents            byte
	MspaCoveredTransaction          byte
	MspaOptOutOptionMode            byte
	MspaServiceProviderMode         byte
}

type USPCA struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment USPCACoreSegment
	GPCSegment  sections.CommonUSGPCSegment
}

func NewUSPCACoreSegment(bs *util.BitStream) (USPCACoreSegment, error) {
	var uspcaCore USPCACoreSegment
	var err error

	uspcaCore.Version, err = bs.ReadByte6()
	if err != nil {
		return uspcaCore, sections.ErrorHelper("CoreSegment.Version", err)
	}

	uspcaCore.SaleOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, sections.ErrorHelper("CoreSegment.SalesOptOutNotice", err)
	}

	uspcaCore.SharingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, sections.ErrorHelper("CoreSegment.SharingOptOutNotice", err)
	}

	uspcaCore.SensitiveDataLimitUseNotice, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, sections.ErrorHelper("CoreSegment.Version", err)
	}

	uspcaCore.SaleOptOut, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, sections.ErrorHelper("CoreSegment.SalesOptOut", err)
	}

	uspcaCore.SharingOptOut, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, sections.ErrorHelper("CoreSegment.SharingOptOut", err)
	}

	uspcaCore.SensitiveDataProcessing, err = bs.ReadTwoBitField(9)
	if err != nil {
		return uspcaCore, sections.ErrorHelper("CoreSegment.SensitiveDataProcessing", err)
	}

	uspcaCore.KnownChildSensitiveDataConsents, err = bs.ReadTwoBitField(2)
	if err != nil {
		return uspcaCore, sections.ErrorHelper("CoreSegment.KnownChildSensitiveDataConsents", err)
	}

	uspcaCore.PersonalDataConsents, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, sections.ErrorHelper("CoreSegment.PersonalDataConsents", err)
	}

	uspcaCore.MspaCoveredTransaction, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, sections.ErrorHelper("CoreSegment.MspaCoveredTransaction", err)
	}

	uspcaCore.MspaOptOutOptionMode, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, sections.ErrorHelper("CoreSegment.MspaOptOutOptionMode", err)
	}

	uspcaCore.MspaServiceProviderMode, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, sections.ErrorHelper("CoreSegment.MspaServiceProviderMode", err)
	}

	return uspcaCore, err
}

func NewUSPCA(encoded string) (USPCA, error) {
	uspca := USPCA{}

	bitStream, err := util.NewBitStreamFromBase64(encoded)
	if err != nil {
		return uspca, err
	}

	coreSegment, err := NewUSPCACoreSegment(bitStream)
	if err != nil {
		return uspca, err
	}

	gpcSegment, err := sections.NewCommonUSGPCSegment(bitStream)
	if err != nil {
		return uspca, err
	}

	uspca = USPCA{
		SectionID:   constants.SectionUSPCA,
		Value:       encoded,
		CoreSegment: coreSegment,
		GPCSegment:  gpcSegment,
	}

	return uspca, nil
}

func (uspca USPCA) GetID() constants.SectionID {
	return uspca.SectionID
}

func (uspca USPCA) GetValue() string {
	return uspca.Value
}
