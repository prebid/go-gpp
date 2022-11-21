package gpp

import (
	"testing"

	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/sections/uspca"
	"github.com/prebid/go-gpp/sections/uspva"
	"github.com/stretchr/testify/assert"
)

type gppTestData struct {
	description string
	gppString   string
	expected    GppContainer
}

func TestParse(t *testing.T) {
	testData := []gppTestData{
		{
			description: "GPP string with EU TCF V2",
			gppString:   "DBABMA~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA",
			expected: GppContainer{
				Version:      1,
				SectionTypes: []constants.SectionID{2},
				Sections: []Section{GenericSection{sectionID: 2,
					value: "CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA"}},
			},
		},
		{
			description: "GPP string with EU TCF v2 and US Privacy",
			gppString:   "DBACNYA~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA~1YNN",
			expected: GppContainer{
				Version:      1,
				SectionTypes: []constants.SectionID{2, 6},
				Sections: []Section{GenericSection{sectionID: 2,
					value: "CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA"},
					GenericSection{sectionID: 6,
						value: "1YNN"}},
			},
		},
		{
			description: "GPP string with Canadian TCF and US Privacy",
			gppString:   "DBABjw~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA~1YNN",
			expected: GppContainer{
				Version:      1,
				SectionTypes: []constants.SectionID{5, 6},
				Sections: []Section{GenericSection{sectionID: 5,
					value: "CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA"},
					GenericSection{sectionID: 6,
						value: "1YNN"}},
			},
		},
		{
			description: "GPP string with USPCA",
			gppString:   "DBABBgA~xlgWEYCZAA",
			expected: GppContainer{
				Version:      1,
				SectionTypes: []constants.SectionID{8},
				Sections: []Section{uspca.USPCA{
					CoreSegment: sections.USPCACoreSegment{
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
		},
		{
			description: "GPP string with USPVA",
			gppString:   "DBABRgA~bSFgmiU",
			expected: GppContainer{
				Version:      1,
				SectionTypes: []constants.SectionID{9},
				Sections: []Section{uspva.USPVA{
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
		},
	}

	for _, test := range testData {
		result, err := Parse(test.gppString)

		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}
}
