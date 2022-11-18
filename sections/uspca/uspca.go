package uspca

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/util"
)

type USPCACoreSegment struct {
	Version                         byte
	SalesOptOutNotice               byte
	SharingOptOutNotice             byte
	SensitiveDataLimitUseNotice     byte
	SalesOptOut                     byte
	SharingOptOut                   byte
	SensitiveDataProcessing         []byte
	KnownChildSensitiveDataConsents []byte
	PersonalDataConsents            byte
	MspaCoveredTransaction          byte
	MspaOptOutOptionMode            byte
	MspaServiceProviderMode         byte
}

type USPCAGPCSegment struct {
	Gpc byte
}

type USPCA struct {
	SectionID   constants.SectionID
	Value       string
	CoreSegment USPCACoreSegment
	GPCSegment  USPCAGPCSegment
}

func initUSPCACoreSegment(bs *util.BitStream) (USPCACoreSegment, error) {
	var result = USPCACoreSegment{}
	var err error

	result.Version, err = bs.ReadByteSize(6, err)
	result.SalesOptOutNotice, err = bs.ReadByteSize(2, err)
	result.SharingOptOutNotice, err = bs.ReadByteSize(2, err)
	result.SensitiveDataLimitUseNotice, err = bs.ReadByteSize(2, err)
	result.SalesOptOut, err = bs.ReadByteSize(2, err)
	result.SharingOptOut, err = bs.ReadByteSize(2, err)
	result.SensitiveDataProcessing, err = bs.ReadTwoBitField(9, err)
	result.KnownChildSensitiveDataConsents, err = bs.ReadTwoBitField(2, err)
	result.PersonalDataConsents, err = bs.ReadByteSize(2, err)
	result.MspaCoveredTransaction, err = bs.ReadByteSize(2, err)
	result.MspaOptOutOptionMode, err = bs.ReadByteSize(2, err)
	result.MspaServiceProviderMode, err = bs.ReadByteSize(2, err)

	return result, err
}

func initUSPCAGPCSegment(bs *util.BitStream) (USPCAGPCSegment, error) {
	var result = USPCAGPCSegment{}
	var err error

	result.Gpc, err = bs.ReadByteSize(1, err)

	return result, err
}

func NewUSPCA(encoded string) (USPCA, error) {
	uspca := USPCA{}

	bitStream, err := util.NewBitStreamFromBase64(encoded)
	if err != nil {
		return uspca, err
	}

	coreSegment, err := initUSPCACoreSegment(bitStream)
	if err != nil {
		return uspca, err
	}

	gpcSegment, err := initUSPCAGPCSegment(bitStream)
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
