package usptx

import (
	"testing"

	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/stretchr/testify/assert"
)

type usptxTestData struct {
	description string
	gppString   string
	expected    USPTX
}

func TestUSPTX(t *testing.T) {
	testData := []usptxTestData{
		{
			description: "should populate USPTX segments correctly",
			gppString:   "bVVVVVVA",
			/*
				011011 01 01 01 01 01 01 0101010101010101 01 01 01 01 01 10000
			*/
			expected: USPTX{
				CoreSegment: USPTXCoreSegment{
					Version:                         27,
					ProcessingNotice:                1,
					SaleOptOutNotice:                1,
					TargetedAdvertisingOptOutNotice: 1,
					SaleOptOut:                      1,
					TargetedAdvertisingOptOut:       1,
					SensitiveDataProcessing:         []byte{1, 1, 1, 1, 1, 1, 1, 1},
					KnownChildSensitiveDataConsents: 1,
					AdditionalDataProcessingConsent: 1,
					MspaCoveredTransaction:          1,
					MspaOptOutOptionMode:            1,
					MspaServiceProviderMode:         1,
				},
				SectionID: constants.SectionUSPTX,
				Value:     "bVVVVVVA",
				GPCSegment: sections.CommonUSGPCSegment{
					SubsectionType: 1,
					Gpc:            false,
				},
			},
		},
	}

	for _, test := range testData {
		result, err := NewUSPTX(test.gppString)
		encodedString := string(test.expected.Encode(false))

		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, constants.SectionUSPTX, result.GetID())
		assert.Equal(t, test.gppString, result.GetValue())
		assert.Equal(t, test.gppString, encodedString)
	}
}
