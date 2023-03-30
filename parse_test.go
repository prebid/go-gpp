package gpp

import (
	"fmt"
	"testing"

	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections"
	"github.com/prebid/go-gpp/sections/uspca"
	"github.com/prebid/go-gpp/sections/uspva"
	"github.com/stretchr/testify/assert"
)

type gppTestData struct {
	description   string
	gppString     string
	expected      GppContainer
	expectedError []error
}

func TestParse(t *testing.T) {
	testData := map[string]gppTestData{
		"GPP-tcf": {
			description: "GPP string with EU TCF V2",
			gppString:   "DBABMA~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA",
			expected: GppContainer{
				Version:      1,
				SectionTypes: []constants.SectionID{2},
				Sections: []Section{GenericSection{sectionID: 2,
					value: "CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA"}},
			},
		},
		"GPP-tcv-usp": {
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
		"GPP-tcfca-usp": {
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
		"GPP-uspca": {
			description: "GPP string with USPCA",
			gppString:   "DBABBgA~xlgWEYCZAA",
			expected: GppContainer{
				Version:      1,
				SectionTypes: []constants.SectionID{8},
				Sections: []Section{uspca.USPCA{
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
		},
		"GPP-uspva": {
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
		"GPP-tcf-error": {
			description: "GPP string with EU TCF V2",
			gppString:   "DBGBMA~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA",
			expected: GppContainer{
				Version:      1,
				SectionTypes: []constants.SectionID{2},
				Sections: []Section{GenericSection{sectionID: 2,
					value: "CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA"}},
			},
			expectedError: []error{fmt.Errorf("error parsing GPP header, section identifiers: error reading an int offset value in a Range(Fibonacci) entry(1): error reading bit 4 of Integer(Fibonacci): expected 1 bit at bit 32, but the byte array was only 4 bytes long")},
		},
		"GPP-uspca-error": {
			description:   "GPP string with USPCA",
			gppString:     "DBABBgA~xlgWE",
			expectedError: []error{fmt.Errorf("error parsing uspca consent string: illegal base64 data at input byte 4")},
		},
	}

	for name, test := range testData {
		t.Run(name, func(t *testing.T) {
			result, err := Parse(test.gppString)

			if len(test.expectedError) == 0 {
				assert.Nil(t, err)
				assert.Equal(t, test.expected, result)
			} else {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}

// go test -bench="^BenchmarkParse$" -benchmem .
// BenchmarkParse-8          625084              1912 ns/op            1472 B/op         48 allocs/op (Apple M1 Pro)
func BenchmarkParse(b *testing.B) {
	const gppString = "DBABrGA~DSJgmkoZJSA.YA~BlgWEYCY.QA~BSFgmiU~bSFgmJQ.YA~BWJYJllA~bSFgmSZQ.YA"
	for i := 0; i < b.N; i++ {
		_, err := Parse(gppString)
		if err != nil {
			b.Fatal(err)
		}
	}
}
