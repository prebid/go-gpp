package sections

import (
	"fmt"

	"github.com/prebid/go-gpp/util"
)

type CommonUSCoreSegment struct {
	Version                            byte
	SharingNotice                      byte
	SaleOptOutNotice                   byte
	TargetedAdvertisingOptOutNotice    byte
	SaleOptOut                         byte
	TargetedAdvertisingOptOut          byte
	SensitiveDataProcessing            []byte
	KnownChildSensitiveDataConsentsInt byte
	KnownChildSensitiveDataConsentsArr []byte
	MspaCoveredTransaction             byte
	MspaOptOutOptionMode               byte
	MspaServiceProviderMode            byte
}

type CommonUSGPCSegment struct {
	Gpc byte
}

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

func errorHelper(name string, err error) error {
	return fmt.Errorf("unable to set field %s due to parse error: %s", name, err.Error())
}

func NewCommonUSCoreSegment(sensitiveDataFields int, knownChildDataFields int, bs *util.BitStream) (CommonUSCoreSegment, error) {
	var commonUSCore CommonUSCoreSegment
	var err error

	commonUSCore.Version, err = bs.ReadByte6()
	if err != nil {
		return commonUSCore, errorHelper("CoreSegment.Version", err)
	}

	commonUSCore.SharingNotice, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, errorHelper("CoreSegment.SharingNotice", err)
	}

	commonUSCore.SaleOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, errorHelper("CoreSegment.SaleOptOutNotice", err)
	}

	commonUSCore.TargetedAdvertisingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, errorHelper("CoreSegment.TargetedAdvertisingOptOutNotice", err)
	}

	commonUSCore.SaleOptOut, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, errorHelper("CoreSegment.SaleOptOut", err)
	}

	commonUSCore.TargetedAdvertisingOptOut, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, errorHelper("CoreSegment.TargetedAdvertisingOptOut", err)
	}

	commonUSCore.SensitiveDataProcessing, err = bs.ReadTwoBitField(sensitiveDataFields)
	if err != nil {
		return commonUSCore, errorHelper("CoreSegment.SensitiveDataProcessing", err)
	}

	if knownChildDataFields > 0 {
		commonUSCore.KnownChildSensitiveDataConsentsArr, err = bs.ReadTwoBitField(knownChildDataFields)
		if err != nil {
			return commonUSCore, errorHelper("CoreSegment.KnownChildSensitiveDataConsentsArr", err)
		}
	} else {
		commonUSCore.KnownChildSensitiveDataConsentsInt, err = bs.ReadByte2()
		commonUSCore.KnownChildSensitiveDataConsentsArr = []byte{}
		if err != nil {
			return commonUSCore, errorHelper("CoreSegment.KnownChildSensitiveDataConsentsInt", err)
		}
	}

	commonUSCore.MspaCoveredTransaction, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, errorHelper("CoreSegment.MspaCoveredTransaction", err)
	}

	commonUSCore.MspaOptOutOptionMode, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, errorHelper("CoreSegment.MspaOptOutOptionMode", err)
	}

	commonUSCore.MspaServiceProviderMode, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, errorHelper("CoreSegment.MspaServiceProviderMode", err)
	}

	return commonUSCore, nil
}

func NewCommonUSGPCSegment(bs *util.BitStream) (CommonUSGPCSegment, error) {
	var commonUSGPC CommonUSGPCSegment
	var err error

	commonUSGPC.Gpc, err = bs.ReadByte1()
	if err != nil {
		return commonUSGPC, errorHelper("GPCSegment.Gpc", err)
	}

	return commonUSGPC, nil
}

func NewUSPCACoreSegment(bs *util.BitStream) (USPCACoreSegment, error) {
	var uspcaCore USPCACoreSegment
	var err error

	uspcaCore.Version, err = bs.ReadByte6()
	if err != nil {
		return uspcaCore, errorHelper("CoreSegment.Version", err)
	}

	uspcaCore.SalesOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, errorHelper("CoreSegment.SalesOptOutNotice", err)
	}

	uspcaCore.SharingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, errorHelper("CoreSegment.SharingOptOutNotice", err)
	}

	uspcaCore.SensitiveDataLimitUseNotice, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, errorHelper("CoreSegment.Version", err)
	}

	uspcaCore.SalesOptOut, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, errorHelper("CoreSegment.SalesOptOut", err)
	}

	uspcaCore.SharingOptOut, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, errorHelper("CoreSegment.SharingOptOut", err)
	}

	uspcaCore.SensitiveDataProcessing, err = bs.ReadTwoBitField(9)
	if err != nil {
		return uspcaCore, errorHelper("CoreSegment.SensitiveDataProcessing", err)
	}

	uspcaCore.KnownChildSensitiveDataConsents, err = bs.ReadTwoBitField(2)
	if err != nil {
		return uspcaCore, errorHelper("CoreSegment.KnownChildSensitiveDataConsents", err)
	}

	uspcaCore.PersonalDataConsents, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, errorHelper("CoreSegment.PersonalDataConsents", err)
	}

	uspcaCore.MspaCoveredTransaction, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, errorHelper("CoreSegment.MspaCoveredTransaction", err)
	}

	uspcaCore.MspaOptOutOptionMode, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, errorHelper("CoreSegment.MspaOptOutOptionMode", err)
	}

	uspcaCore.MspaServiceProviderMode, err = bs.ReadByte2()
	if err != nil {
		return uspcaCore, errorHelper("CoreSegment.MspaServiceProviderMode", err)
	}

	return uspcaCore, err
}

func NewUSPNATCoreSegment(bs *util.BitStream) (USPNATCoreSegment, error) {
	var uspnatCore USPNATCoreSegment
	var err error

	uspnatCore.Version, err = bs.ReadByte6()
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.Version", err)
	}

	uspnatCore.SharingNotice, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.SharingNotice", err)
	}

	uspnatCore.SaleOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.SaleOptOutNotice", err)
	}

	uspnatCore.SharingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.SharingOptOutNotice", err)
	}

	uspnatCore.TargetedAdvertisingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.TargetedAdvertisingOptOutNotice", err)
	}

	uspnatCore.SensitiveDataProcessingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.SensitiveDataProcessingOptOutNotice", err)
	}

	uspnatCore.SensitiveDataLimitUseNotice, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.SensitiveDataLimitUseNotice", err)
	}

	uspnatCore.SaleOptOut, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.SaleOptOut", err)
	}

	uspnatCore.SharingOptOut, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.SharingOptOut", err)
	}

	uspnatCore.TargetedAdvertisingOptOut, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.TargetedAdvertisingOptOut", err)
	}

	uspnatCore.SensitiveDataProcessing, err = bs.ReadTwoBitField(12)
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.SensitiveDataProcessing", err)
	}

	uspnatCore.KnownChildSensitiveDataConsents, err = bs.ReadTwoBitField(2)
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.KnownChildSensitiveDataConsents", err)
	}

	uspnatCore.PersonalDataConsents, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.PersonalDataConsents", err)
	}

	uspnatCore.MspaCoveredTransaction, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.MspaCoveredTransaction", err)
	}

	uspnatCore.MspaOptOutOptionMode, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.MspaOptOutOptionMode", err)
	}

	uspnatCore.MspaServiceProviderMode, err = bs.ReadByte2()
	if err != nil {
		return uspnatCore, errorHelper("CoreSegment.MspaServiceProviderMode", err)
	}

	return uspnatCore, nil
}

func NewUPSUTCoreSegment(bs *util.BitStream) (USPUTCoreSegment, error) {
	var usputCore USPUTCoreSegment
	var err error

	usputCore.Version, err = bs.ReadByte6()
	if err != nil {
		return usputCore, errorHelper("CoreSegment.Version", err)
	}

	usputCore.SharingNotice, err = bs.ReadByte2()
	if err != nil {
		return usputCore, errorHelper("CoreSegment.SharingNotice", err)
	}

	usputCore.SaleOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return usputCore, errorHelper("CoreSegment.SaleOptOutNotice", err)
	}

	usputCore.TargetedAdvertisingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return usputCore, errorHelper("CoreSegment.TargetedAdvertisingOptOutNotice", err)
	}

	usputCore.SensitiveDataProcessingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return usputCore, errorHelper("CoreSegment.SensitiveDataProcessingOptOutNotice", err)
	}

	usputCore.SaleOptOut, err = bs.ReadByte2()
	if err != nil {
		return usputCore, errorHelper("CoreSegment.SaleOptOut", err)
	}

	usputCore.TargetedAdvertisingOptOut, err = bs.ReadByte2()
	if err != nil {
		return usputCore, errorHelper("CoreSegment.TargetedAdvertisingOptOut", err)
	}

	usputCore.SensitiveDataProcessing, err = bs.ReadTwoBitField(8)
	if err != nil {
		return usputCore, errorHelper("CoreSegment.SensitiveDataProcessing", err)
	}

	usputCore.KnownChildSensitiveDataConsents, err = bs.ReadByte2()
	if err != nil {
		return usputCore, errorHelper("CoreSegment.KnownChildSensitiveDataConsents", err)
	}

	usputCore.MspaCoveredTransaction, err = bs.ReadByte2()
	if err != nil {
		return usputCore, errorHelper("CoreSegment.MspaCoveredTransaction", err)
	}

	usputCore.MspaOptOutOptionMode, err = bs.ReadByte2()
	if err != nil {
		return usputCore, errorHelper("CoreSegment.MspaOptOutOptionMode", err)
	}

	usputCore.MspaServiceProviderMode, err = bs.ReadByte2()
	if err != nil {
		return usputCore, errorHelper("CoreSegment.MspaServiceProviderMode", err)
	}

	return usputCore, nil
}
