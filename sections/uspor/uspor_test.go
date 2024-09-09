package uspor

import (
	"github.com/prebid/go-gpp/sections"
	"testing"

	"github.com/prebid/go-gpp/constants"
	"github.com/stretchr/testify/assert"
)

type usporTestData struct {
	description string
	gppString   string
	expected    USPOR
}

func TestUSPOR(t *testing.T) {
	testData := []usporTestData{
		{
			description: "should populate USPOR segments correctly",
			gppString:   "bSFgmGACUA.YA",
			/*
				011011 01 00 10 00 01 0110000010011000011000 000000 00 10 01 01 1100 01 1
			*/
			expected: USPOR{
				CoreSegment: USPORCoreSegment{
					Version:                         27,
					ProcessingNotice:                1,
					SaleOptOutNotice:                0,
					TargetedAdvertisingOptOutNotice: 2,
					SaleOptOut:                      0,
					TargetedAdvertisingOptOut:       1,
					SensitiveDataProcessing: []byte{
						1, 2, 0, 0, 2, 1, 2, 0, 1, 2, 0,
					},
					KnownChildSensitiveDataConsents: []byte{0, 0, 0},
					AdditionalDataProcessingConsent: 0,
					MspaCoveredTransaction:          2,
					MspaOptOutOptionMode:            1,
					MspaServiceProviderMode:         1,
				},
				SectionID: constants.SectionUSPOR,
				Value:     "bSFgmGACUA.YA",
				GPCSegment: sections.CommonUSGPCSegment{
					SubsectionType: 1,
					Gpc:            true,
				},
			},
		},
	}

	for _, test := range testData {
		result, err := NewUSPOR(test.gppString)
		encodedString := string(test.expected.Encode(true))

		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, constants.SectionUSPOR, result.GetID())
		assert.Equal(t, test.gppString, result.GetValue())
		assert.Equal(t, test.gppString, encodedString)
	}
}
