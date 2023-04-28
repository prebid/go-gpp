package uspct

import (
	"testing"

	"github.com/revcontent-production/go-gpp/constants"
	"github.com/revcontent-production/go-gpp/sections"
	"github.com/stretchr/testify/assert"
)

type uspctTestData struct {
	description string
	gppString   string
	expected    USPCT
}

func TestUSPCT(t *testing.T) {
	testData := []uspctTestData{
		{
			description: "should populate USPCT segments correctly",
			gppString:   "bSFgmSZW.YAAA",
			/*
				011011 01 00 10 00 01 0110000010011001 001001 10 01 01 01 1 011
			*/
			expected: USPCT{
				CoreSegment: sections.CommonUSCoreSegment{
					Version:                         27,
					SharingNotice:                   1,
					SaleOptOutNotice:                0,
					TargetedAdvertisingOptOutNotice: 2,
					SaleOptOut:                      0,
					TargetedAdvertisingOptOut:       1,
					SensitiveDataProcessing: []byte{
						1, 2, 0, 0, 2, 1, 2, 1,
					},
					KnownChildSensitiveDataConsents: []byte{
						0, 2, 1,
					},
					MspaCoveredTransaction:  2,
					MspaOptOutOptionMode:    1,
					MspaServiceProviderMode: 1,
				},
				GPCSegment: sections.CommonUSGPCSegment{
					SubsectionType: 1,
					Gpc:            true,
				},
				SectionID: constants.SectionUSPCT,
				Value:     "bSFgmSZW.YAAA",
			},
		},
	}

	for _, test := range testData {
		result, err := NewUSPCT(test.gppString)

		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, constants.SectionUSPCT, result.GetID())
		assert.Equal(t, test.gppString, result.GetValue())
	}
}
