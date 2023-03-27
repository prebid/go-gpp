package gpp

import (
	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/sections/uspca"
	"github.com/prebid/go-gpp/sections/uspco"
	"github.com/prebid/go-gpp/sections/uspct"
	"github.com/prebid/go-gpp/sections/uspnat"
	"github.com/prebid/go-gpp/sections/usput"
	"github.com/prebid/go-gpp/sections/uspva"
	"github.com/stretchr/testify/assert"
	"testing"
)

type gppEncodeTestData struct {
	description string
	sections    []Section
	expected    string
}

var testData = []gppEncodeTestData{
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
	{
		description: "USPCO GPP string encoding",
		expected:    "DBABJg~bSFgmJQ.YA",
		sections: []Section{uspco.USPCO{
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
				SubsectionType: 1,
				Gpc:            true,
			},
			SectionID: constants.SectionUSPCO,
			Value:     "bSFgmJQ.YA",
		},
		},
	},
	{
		description: "USPCT GPP string encoding",
		expected:    "DBABVg~bSFgmSZQ.YA",
		sections: []Section{uspct.USPCT{
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
			Value:     "bSFgmSZQ.YA",
		},
		},
	},
	{
		description: "USPNAT GPP string encoding",
		expected:    "DBABLA~DSJgmkoZJSA.YA",
		sections: []Section{uspnat.USPNAT{
			CoreSegment: uspnat.USPNATCoreSegment{
				Version:                             3,
				SharingNotice:                       1,
				SaleOptOutNotice:                    0,
				SharingOptOutNotice:                 2,
				TargetedAdvertisingOptOutNotice:     0,
				SensitiveDataProcessingOptOutNotice: 2,
				SensitiveDataLimitUseNotice:         1,
				SaleOptOut:                          2,
				SharingOptOut:                       0,
				TargetedAdvertisingOptOut:           0,
				SensitiveDataProcessing: []byte{
					2, 1, 2, 2, 1, 0, 2, 2, 0, 1, 2, 1,
				},
				KnownChildSensitiveDataConsents: []byte{
					0, 2,
				},
				PersonalDataConsents:    1,
				MspaCoveredTransaction:  1,
				MspaOptOutOptionMode:    0,
				MspaServiceProviderMode: 2,
			},
			GPCSegment: sections.CommonUSGPCSegment{
				SubsectionType: 1,
				Gpc:            true,
			},
			SectionID: constants.SectionUSPNAT,
			Value:     "DSJgmkoZJSA.YA",
		},
		},
	},
	{
		description: "USPUT GPP string encoding",
		expected:    "DBABFg~bSRYJllA",
		sections: []Section{usput.USPUT{
			CoreSegment: usput.USPUTCoreSegment{
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
	},
}

func TestEncode(t *testing.T) {
	for _, test := range testData {
		result, err := Encode(test.sections)

		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}
}

// go test -bench="^BenchmarkEncode$" -benchmem .
func BenchmarkEncode(b *testing.B) {
	var secs []Section
	for i := 0; i < len(testData); i++ {
		secs = append(secs, testData[i].sections[0])
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Encode(secs)
		if err != nil {
			b.Fatal(err)
		}
	}
}
