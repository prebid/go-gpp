package uspca

import (
	"fmt"

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

	result.Version, err = bs.ReadByte6()
	if err != nil {
		return result, fmt.Errorf("unable to set Version: %s", err.Error())
	}

	result.SalesOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set SalesOptOutNotice: %s", err.Error())
	}

	result.SharingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set SharingOptOutNotice: %s", err.Error())
	}

	result.SensitiveDataLimitUseNotice, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set SensitiveDataLimitUseNotice: %s", err.Error())
	}

	result.SalesOptOut, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set SalesOptOut: %s", err.Error())
	}

	result.SharingOptOut, err = bs.ReadByte2()
	if err != nil {
		return result, fmt.Errorf("unable to set SharingOptOut: %s", err.Error())
	}

	result.SensitiveDataProcessing, err = bs.ReadTwoBitField(9)
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

func initUSPCAGPCSegment(bs *util.BitStream) (USPCAGPCSegment, error) {
	var result = USPCAGPCSegment{}
	var err error

	result.Gpc, err = bs.ReadByte1()
	if err != nil {
		return result, fmt.Errorf("unable to set GPC: %s", err.Error())
	}

	return result, nil
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
