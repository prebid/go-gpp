package sections

import (
	"fmt"

	"github.com/prebid/go-gpp/util"
)

type CommonUSCoreSegment struct {
	Version                         byte
	SharingNotice                   byte
	SaleOptOutNotice                byte
	TargetedAdvertisingOptOutNotice byte
	SaleOptOut                      byte
	TargetedAdvertisingOptOut       byte
	SensitiveDataProcessing         []byte
	KnownChildSensitiveDataConsents []byte
	MspaCoveredTransaction          byte
	MspaOptOutOptionMode            byte
	MspaServiceProviderMode         byte
}

type CommonUSGPCSegment struct {
	SubsectionType byte
	Gpc            bool
}

func ErrorHelper(name string, err error) error {
	return fmt.Errorf("unable to set field %s due to parse error: %s", name, err.Error())
}

func NewCommonUSCoreSegment(sensitiveDataFields int, knownChildDataFields int, bs *util.BitStream) (CommonUSCoreSegment, error) {
	var commonUSCore CommonUSCoreSegment
	var err error

	commonUSCore.Version, err = bs.ReadByte6()
	if err != nil {
		return commonUSCore, ErrorHelper("CoreSegment.Version", err)
	}

	commonUSCore.SharingNotice, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, ErrorHelper("CoreSegment.SharingNotice", err)
	}

	commonUSCore.SaleOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, ErrorHelper("CoreSegment.SaleOptOutNotice", err)
	}

	commonUSCore.TargetedAdvertisingOptOutNotice, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, ErrorHelper("CoreSegment.TargetedAdvertisingOptOutNotice", err)
	}

	commonUSCore.SaleOptOut, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, ErrorHelper("CoreSegment.SaleOptOut", err)
	}

	commonUSCore.TargetedAdvertisingOptOut, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, ErrorHelper("CoreSegment.TargetedAdvertisingOptOut", err)
	}

	commonUSCore.SensitiveDataProcessing, err = bs.ReadTwoBitField(sensitiveDataFields)
	if err != nil {
		return commonUSCore, ErrorHelper("CoreSegment.SensitiveDataProcessing", err)
	}

	commonUSCore.KnownChildSensitiveDataConsents, err = bs.ReadTwoBitField(knownChildDataFields)
	if err != nil {
		return commonUSCore, ErrorHelper("CoreSegment.KnownChildSensitiveDataConsentsArr", err)
	}

	commonUSCore.MspaCoveredTransaction, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, ErrorHelper("CoreSegment.MspaCoveredTransaction", err)
	}

	commonUSCore.MspaOptOutOptionMode, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, ErrorHelper("CoreSegment.MspaOptOutOptionMode", err)
	}

	commonUSCore.MspaServiceProviderMode, err = bs.ReadByte2()
	if err != nil {
		return commonUSCore, ErrorHelper("CoreSegment.MspaServiceProviderMode", err)
	}

	return commonUSCore, nil
}

func NewCommonUSGPCSegment(bs *util.BitStream) (CommonUSGPCSegment, error) {
	var commonUSGPC CommonUSGPCSegment
	var err error

	commonUSGPC.SubsectionType, err = bs.ReadByte2()
	if err != nil {
		return commonUSGPC, ErrorHelper("GPCSegment.SubsectionType", err)
	}

	gpc, err := bs.ReadByte1()
	if err != nil {
		return commonUSGPC, ErrorHelper("GPCSegment.Gpc", err)
	}
	commonUSGPC.Gpc = (gpc == 1)

	return commonUSGPC, nil
}
