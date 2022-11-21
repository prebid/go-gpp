package usput

import (
	"testing"

	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/stretchr/testify/assert"
)

type usputTestData struct {
	description string
	gppString   string
	expected    USPUT
}

func TestUSPUT(t *testing.T) {
	testData := []usputTestData{
		{
			description: "should populate USPUT segments correctly",
			gppString:   "bSRYJllA",
			/*
				011011 01 00 10 01 00 01 0110000010011001 01 10 01 01
			*/
			expected: USPUT{
				CoreSegment: sections.USPUTCoreSegment{
					Version:                             27,
					SharingNotice:                       1,
					SaleOptOutNotice:                    0,
					TargetedAdvertisingOptOutNotice:     2,
					SensitiveDataProcessingOptOutNotice: 1,
					SaleOptOut:                          0,
					TargetedAdvertisingOptOut:           1,
					SensitiveDataProcessing: []byte{
						1, 2, 0, 0, 2, 1, 2, 1,
					},
					KnownChildSensitiveDataConsents: 1,
					MspaCoveredTransaction:          2,
					MspaOptOutOptionMode:            1,
					MspaServiceProviderMode:         1,
				},
				SectionID: constants.SectionUSPUT,
				Value:     "bSRYJllA",
			},
		},
	}

	for _, test := range testData {
		result, err := NewUSPUT(test.gppString)

		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, constants.SectionUSPUT, result.GetID())
		assert.Equal(t, test.gppString, result.GetValue())
	}
}
