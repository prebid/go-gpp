package sections

import (
	"errors"
	"fmt"
	"strings"

	"github.com/revcontent-production/go-gpp/util"
)

// CommonUSCoreSegment replicates the structure of many US State consent strings, and as such is shared
// across those state implementations. Do note an array of size one is a valid construct
// if the state only has one subfield for that field.
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

func (segment CommonUSCoreSegment) Encode(bs *util.BitStream) {
	bs.WriteByte6(segment.Version)
	bs.WriteByte2(segment.SharingNotice)
	bs.WriteByte2(segment.SaleOptOutNotice)
	bs.WriteByte2(segment.TargetedAdvertisingOptOutNotice)
	bs.WriteByte2(segment.SaleOptOut)
	bs.WriteByte2(segment.TargetedAdvertisingOptOut)
	bs.WriteTwoBitField(segment.SensitiveDataProcessing)
	bs.WriteTwoBitField(segment.KnownChildSensitiveDataConsents)
	bs.WriteByte2(segment.MspaCoveredTransaction)
	bs.WriteByte2(segment.MspaOptOutOptionMode)
	bs.WriteByte2(segment.MspaServiceProviderMode)
}

func NewCommonUSGPCSegment(bs *util.BitStream) (CommonUSGPCSegment, error) {
	var commonUSGPC CommonUSGPCSegment
	var err error

	commonUSGPC.SubsectionType, err = bs.ReadByte2()
	if err != nil {
		return commonUSGPC, ErrorHelper("GPCSegment.SubsectionType", err)
	}

	if commonUSGPC.SubsectionType != 1 {
		return commonUSGPC, fmt.Errorf("invalid subsection type %d for GPC segment", commonUSGPC.SubsectionType)
	}

	gpc, err := bs.ReadByte1()
	if err != nil {
		return commonUSGPC, ErrorHelper("GPCSegment.Gpc", err)
	}
	commonUSGPC.Gpc = (gpc == 1)

	return commonUSGPC, nil
}

func (gpc CommonUSGPCSegment) Encode(bs *util.BitStream) {
	bs.WriteByte2(gpc.SubsectionType)
	if gpc.Gpc {
		bs.WriteByte1(1)
	} else {
		bs.WriteByte1(0)
	}
}

func CreateBitStreams(encoded string, gpcCheck bool) (*util.BitStream, *util.BitStream, error) {
	segments := strings.Split(encoded, ".")

	if len(segments) == 0 {
		return nil, nil, errors.New("no segments found in encoded string")
	}

	coreBitStream, err := util.NewBitStreamFromBase64(segments[0])
	if err != nil {
		return nil, nil, err
	}

	if gpcCheck && len(segments) > 1 {
		gpcBitStream, err := util.NewBitStreamFromBase64(segments[1])
		if err != nil {
			return nil, nil, err
		}

		return coreBitStream, gpcBitStream, nil
	}

	return coreBitStream, nil, nil
}
