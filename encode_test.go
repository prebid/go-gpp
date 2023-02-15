package gpp

import (
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/sections/uspca"
	"github.com/prebid/go-gpp/sections/uspva"
	"github.com/stretchr/testify/assert"
	"testing"
)

type gppEncodeTestData struct {
	description string
	sections    []Section
	expected    string
}

func TestEncode(t *testing.T) {
	testData := []gppEncodeTestData{
		{
			description: "USPCA GPP string encoding",
			expected:    "DBABBg~xlgWEYCY.QA",
			sections: []Section{uspca.USPCA{
				CoreSegment: uspca.USPCACoreSegment{
					Version:                     49,
					SaleOptOutNotice:            2,
					SharingOptOutNotice:         1,
					SensitiveDataLimitUseNotice: 1,
					SaleOptOut:                  2,
					SharingOptOut:               0,
					SensitiveDataProcessing: []byte{
						0, 1, 1, 2, 0, 1, 0, 1, 2,
					},
					KnownChildSensitiveDataConsents: []byte{
						0, 0,
					},
					PersonalDataConsents:    0,
					MspaCoveredTransaction:  2,
					MspaOptOutOptionMode:    1,
					MspaServiceProviderMode: 2,
				},
				GPCSegment: sections.CommonUSGPCSegment{
					SubsectionType: 1,
					Gpc:            false,
				},
				SectionID: 8,
				Value:     "xlgWEYCZAA"},
			},
		},
		{
			description: "USPVA GPP string encoding",
			expected:    "DBABRg~bSFgmiU",
			sections: []Section{uspva.USPVA{
				CoreSegment: sections.CommonUSCoreSegment{
					Version:                         27,
					SharingNotice:                   1,
					SaleOptOutNotice:                0,
					TargetedAdvertisingOptOutNotice: 2,
					SaleOptOut:                      0,
					TargetedAdvertisingOptOut:       1,
					SensitiveDataProcessing: []byte{
						1, 2, 0, 0, 2, 1, 2, 2,
					},
					KnownChildSensitiveDataConsents: []byte{0},
					MspaCoveredTransaction:          2,
					MspaOptOutOptionMode:            1,
					MspaServiceProviderMode:         1,
				},
				SectionID: 9,
				Value:     "bSFgmiU"},
			},
		},
	}

	for _, test := range testData {
		result, err := Encode(test.sections)

		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}
}
