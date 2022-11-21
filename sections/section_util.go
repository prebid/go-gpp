package sections

import (
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

func NewCommonUSCoreSegment(sensitiveDataFields int, knownChildDataFields int, bs *util.BitStream) (CommonUSCoreSegment, error) {
	var commonUSCore CommonUSCoreSegment
	var err error

	commonUSCore.Version, err = bs.ReadByteSize(6, err)
	commonUSCore.SharingNotice, err = bs.ReadByteSize(2, err)
	commonUSCore.SaleOptOutNotice, err = bs.ReadByteSize(2, err)
	commonUSCore.TargetedAdvertisingOptOutNotice, err = bs.ReadByteSize(2, err)
	commonUSCore.SaleOptOut, err = bs.ReadByteSize(2, err)
	commonUSCore.TargetedAdvertisingOptOut, err = bs.ReadByteSize(2, err)
	commonUSCore.SensitiveDataProcessing, err = bs.ReadTwoBitField(sensitiveDataFields, err)

	if knownChildDataFields > 0 {
		commonUSCore.KnownChildSensitiveDataConsentsArr, err = bs.ReadTwoBitField(knownChildDataFields, err)
	} else {
		commonUSCore.KnownChildSensitiveDataConsentsInt, err = bs.ReadByteSize(2, err)
		commonUSCore.KnownChildSensitiveDataConsentsArr = []byte{}
	}

	commonUSCore.MspaCoveredTransaction, err = bs.ReadByteSize(2, err)
	commonUSCore.MspaOptOutOptionMode, err = bs.ReadByteSize(2, err)
	commonUSCore.MspaServiceProviderMode, err = bs.ReadByteSize(2, err)

	return commonUSCore, err
}

func NewCommonUSGPCSegment(bs *util.BitStream) (CommonUSGPCSegment, error) {
	var commonUSGPC CommonUSGPCSegment
	var err error

	commonUSGPC.Gpc, err = bs.ReadByteSize(1, err)

	return commonUSGPC, err
}

func NewUSPCACoreSegment(bs *util.BitStream) (USPCACoreSegment, error) {
	var uspcaCore USPCACoreSegment
	var err error

	uspcaCore.Version, err = bs.ReadByteSize(6, err)
	uspcaCore.SalesOptOutNotice, err = bs.ReadByteSize(2, err)
	uspcaCore.SharingOptOutNotice, err = bs.ReadByteSize(2, err)
	uspcaCore.SensitiveDataLimitUseNotice, err = bs.ReadByteSize(2, err)
	uspcaCore.SalesOptOut, err = bs.ReadByteSize(2, err)
	uspcaCore.SharingOptOut, err = bs.ReadByteSize(2, err)
	uspcaCore.SensitiveDataProcessing, err = bs.ReadTwoBitField(9, err)
	uspcaCore.KnownChildSensitiveDataConsents, err = bs.ReadTwoBitField(2, err)
	uspcaCore.PersonalDataConsents, err = bs.ReadByteSize(2, err)
	uspcaCore.MspaCoveredTransaction, err = bs.ReadByteSize(2, err)
	uspcaCore.MspaOptOutOptionMode, err = bs.ReadByteSize(2, err)
	uspcaCore.MspaServiceProviderMode, err = bs.ReadByteSize(2, err)

	return uspcaCore, err
}

func NewUSPNATCoreSegment(bs *util.BitStream) (USPNATCoreSegment, error) {
	var uspnatCore USPNATCoreSegment
	var err error

	uspnatCore.Version, err = bs.ReadByteSize(6, err)
	uspnatCore.SharingNotice, err = bs.ReadByteSize(2, err)
	uspnatCore.SaleOptOutNotice, err = bs.ReadByteSize(2, err)
	uspnatCore.SharingOptOutNotice, err = bs.ReadByteSize(2, err)
	uspnatCore.TargetedAdvertisingOptOutNotice, err = bs.ReadByteSize(2, err)
	uspnatCore.SensitiveDataProcessingOptOutNotice, err = bs.ReadByteSize(2, err)
	uspnatCore.SensitiveDataLimitUseNotice, err = bs.ReadByteSize(2, err)
	uspnatCore.SaleOptOut, err = bs.ReadByteSize(2, err)
	uspnatCore.SharingOptOut, err = bs.ReadByteSize(2, err)
	uspnatCore.TargetedAdvertisingOptOut, err = bs.ReadByteSize(2, err)
	uspnatCore.SensitiveDataProcessing, err = bs.ReadTwoBitField(12, err)
	uspnatCore.KnownChildSensitiveDataConsents, err = bs.ReadTwoBitField(2, err)
	uspnatCore.PersonalDataConsents, err = bs.ReadByteSize(2, err)
	uspnatCore.MspaCoveredTransaction, err = bs.ReadByteSize(2, err)
	uspnatCore.MspaOptOutOptionMode, err = bs.ReadByteSize(2, err)
	uspnatCore.MspaServiceProviderMode, err = bs.ReadByteSize(2, err)

	return uspnatCore, err
}

func NewUPSUTCoreSegment(bs *util.BitStream) (USPUTCoreSegment, error) {
	var usputCore USPUTCoreSegment
	var err error

	usputCore.Version, err = bs.ReadByteSize(6, err)
	usputCore.SharingNotice, err = bs.ReadByteSize(2, err)
	usputCore.SaleOptOutNotice, err = bs.ReadByteSize(2, err)
	usputCore.TargetedAdvertisingOptOutNotice, err = bs.ReadByteSize(2, err)
	usputCore.SensitiveDataProcessingOptOutNotice, err = bs.ReadByteSize(2, err)
	usputCore.SaleOptOut, err = bs.ReadByteSize(2, err)
	usputCore.TargetedAdvertisingOptOut, err = bs.ReadByteSize(2, err)
	usputCore.SensitiveDataProcessing, err = bs.ReadTwoBitField(8, err)
	usputCore.KnownChildSensitiveDataConsents, err = bs.ReadByteSize(2, err)
	usputCore.MspaCoveredTransaction, err = bs.ReadByteSize(2, err)
	usputCore.MspaOptOutOptionMode, err = bs.ReadByteSize(2, err)
	usputCore.MspaServiceProviderMode, err = bs.ReadByteSize(2, err)

	return usputCore, err
}
