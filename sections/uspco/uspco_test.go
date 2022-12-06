package uspco

import (
	"testing"

	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/stretchr/testify/assert"
)

type uspcoTestData struct {
	description string
	gppString   string
	expected    USPCO
}

func TestUSPCO(t *testing.T) {
	testData := []uspcoTestData{
		{
			description: "should populate USPCO segments correctly",
			gppString:   "bSFgmJQA",
			/*
				011011 01 00 10 00 01 01100000100110 00 10 01 01 00 0
			*/
			expected: USPCO{
				CoreSegment: sections.CommonUSCoreSegment{
					Version:                         27,
					SharingNotice:                   1,
					SaleOptOutNotice:                0,
					TargetedAdvertisingOptOutNotice: 2,
					SaleOptOut:                      0,
					TargetedAdvertisingOptOut:       1,
					SensitiveDataProcessing: []byte{
						1, 2, 0, 0, 2, 1, 2,
					},
					KnownChildSensitiveDataConsents: []byte{0},
					MspaCoveredTransaction:          2,
					MspaOptOutOptionMode:            1,
					MspaServiceProviderMode:         1,
				},
				GPCSegment: sections.CommonUSGPCSegment{
					SubsectionType: 0,
					Gpc:            false,
				},
				SectionID: constants.SectionUSPCO,
				Value:     "bSFgmJQA",
			},
		},
	}

	for _, test := range testData {
		result, err := NewUSPCO(test.gppString)

		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, constants.SectionUSPCO, result.GetID())
		assert.Equal(t, test.gppString, result.GetValue())
	}
}
