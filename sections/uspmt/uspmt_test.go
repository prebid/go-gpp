package uspmt

import (
	"testing"

	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/stretchr/testify/assert"
)

type uspmtTestData struct {
	description string
	gppString   string
	expected    USPMT
}

func TestUSPCO(t *testing.T) {
	testData := []uspmtTestData{
		{
			description: "should populate USPMT segments correctly",
			gppString:   "bSFgmAGU.YA",
			/*
				011011 01 00 10 00 01 0110000010011000 000000 01 10 01 01 01 1 1110
			*/
			expected: USPMT{
				CoreSegment: USPMTCoreSegment{
					Version:                         27,
					SharingNotice:                   1,
					SaleOptOutNotice:                0,
					TargetedAdvertisingOptOutNotice: 2,
					SaleOptOut:                      0,
					TargetedAdvertisingOptOut:       1,
					SensitiveDataProcessing: []byte{
						1, 2, 0, 0, 2, 1, 2, 0,
					},
					KnownChildSensitiveDataConsents: []byte{0, 0, 0},
					AdditionalDataProcessingConsent: 1,
					MspaCoveredTransaction:          2,
					MspaOptOutOptionMode:            1,
					MspaServiceProviderMode:         1,
				},
				GPCSegment: sections.CommonUSGPCSegment{
					SubsectionType: 1,
					Gpc:            true,
				},
				SectionID: constants.SectionUSPMT,
				Value:     "bSFgmAGU.YA",
			},
		},
	}

	for _, test := range testData {
		result, err := NewUSPMT(test.gppString)
		encodedString := string(test.expected.Encode(true))

		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, constants.SectionUSPMT, result.GetID())
		assert.Equal(t, test.gppString, result.GetValue())
		assert.Equal(t, test.gppString, encodedString)
	}
}
