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
	err         error
}

var testData = []gppEncodeTestData{
	{
		description: "USPCA GPP string encoding",
		expected:    "DBABh4A~BlgWEYCY.QA~BSFgmiU",
		sections: []Section{
			uspca.USPCA{
				CoreSegment: uspca.USPCACoreSegment{
					Version:                     1,
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
				SectionID: constants.SectionUSPCA,
				Value:     "BlgWEYCY.QA"},
			uspva.USPVA{
				CoreSegment: sections.CommonUSCoreSegment{
					Version:                         1,
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
				SectionID: constants.SectionUSPVA,
				Value:     "BSFgmiU"},
		},
	},
	{
		description: "USPVA GPP string encoding",
		expected:    "DBABRg~bSFgmiU",
		sections: []Section{
			uspva.USPVA{
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
				SectionID: constants.SectionUSPVA,
				Value:     "bSFgmiU"},
		},
	},
	{
		description: "USPCO GPP string encoding",
		expected:    "DBABJg~bSFgmJQ.YA",
		sections: []Section{
			uspco.USPCO{
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
				Value:     "bSFgmJQ.YA"},
		},
	},
	{
		description: "USPCT GPP string encoding",
		expected:    "DBABVg~bSFgmSZQ.YA",
		sections: []Section{
			uspct.USPCT{
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
				Value:     "bSFgmSZQ.YA"},
		},
	},
	{
		description: "USPNAT GPP string encoding",
		expected:    "DBABLA~DSJgmkoZJSA.YA",
		sections: []Section{
			uspnat.USPNAT{
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
				Value:     "DSJgmkoZJSA.YA"},
		},
	},
	{
		description: "GPP string encoding for multiple sections",
		expected:    "DBADLO8~BSJgmkoZJSA.YA~BSFgmiU~BWJYJllA~BSFgmSZQ.YA",
		sections: []Section{
			usput.USPUT{
				CoreSegment: usput.USPUTCoreSegment{
					Version:                             1,
					SharingNotice:                       1,
					SaleOptOutNotice:                    1,
					TargetedAdvertisingOptOutNotice:     2,
					SensitiveDataProcessingOptOutNotice: 0,
					SaleOptOut:                          2,
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
				Value:     "BWJYJllA"},
			uspnat.USPNAT{
				CoreSegment: uspnat.USPNATCoreSegment{
					Version:                             1,
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
				Value:     "BSJgmkoZJSA.YA"},
			uspct.USPCT{
				CoreSegment: sections.CommonUSCoreSegment{
					Version:                         1,
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
				Value:     "BSFgmSZQ.YA"},
			uspva.USPVA{
				CoreSegment: sections.CommonUSCoreSegment{
					Version:                         1,
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
				SectionID: constants.SectionUSPVA,
				Value:     "BSFgmiU"},
		},
	},
	{
		description: "GPP string encoding minimum section ID",
		expected:    "DBABYA~test_minimum",
		sections:    []Section{GenericSection{sectionID: 1, value: "test_minimum"}},
	},
	{
		description: "GPP string encoding section ID zero",
		expected:    "",
		sections:    []Section{GenericSection{sectionID: 0, value: "section_with_id_zero"}},
		err:         sectionIdOutOfRangeErr,
	},
}

func TestEncode(t *testing.T) {
	for _, test := range testData {
		result, err := Encode(test.sections)

		assert.Equal(t, test.err, err)
		assert.Equal(t, test.expected, result)

		if err != nil {
			continue
		}
		// Parse result to see whether the GPP string can be translated back into the original sections.
		container, errs := Parse(result)
		if len(errs) != 0 {
			t.Fatal(errs)
		}
		secSet := make(map[constants.SectionID]Section, len(container.Sections))
		for _, secParsed := range container.Sections {
			secSet[secParsed.GetID()] = secParsed
		}
		for _, section := range test.sections {
			if secParsed, ok := secSet[section.GetID()]; !ok {
				t.Fatalf("Section %v was not parsed successfully", section.GetID())
			} else {
				assert.Equal(t, section, secParsed)
			}
		}
	}
}

// Decode given GPP strings and re-encode them to see if we can get the original ones back.
func TestEncode2(t *testing.T) {
	gppStrings := []string{
		"DBABh4A~BlgWEYCY.QA~BSFgmiU",
		"DBABRg~bSFgmiU",
		"DBABJg~bSFgmJQ.YA",
		"DBABVg~bSFgmSZQ.YA",
		"DBABLA~DSJgmkoZJSA.YA",
		"DBADLO8~BSJgmkoZJSA.YA~BSFgmiU~BWJYJllA~BSFgmSZQ.YA",
		"DBABrGA~DSJgmkoZJSA.YA~BlgWEYCY.QA~BSFgmiU~bSFgmJQ.YA~BWJYJllA~bSFgmSZQ.YA",
		"DBACLMA~BAAAAAAAAAA.QA~BaAAAAA",
		"DBACTjw~1YYN~BSZZYgkA~BaRlkCSA.QA",
		"DBACMYA~CPpcCoAPpcCoAPoABABGCyCUACAAACAAAAAAAVQAQAVABZABABYAAAAA.QADgIAAA.IABE~CPpcCoAPpcCoAPoABABGCyCQAEAAAEAAAAEFABAEEAN8AEAN4A.YAAAAAAAAAA",
	}
	for _, s := range gppStrings {
		container, errs := Parse(s)
		if len(errs) != 0 {
			t.Fatal(errs)
		}
		gpp, err := Encode(container.Sections)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, s, gpp)
	}
}

// go test -bench="^BenchmarkEncode$" -benchmem .
// BenchmarkEncode-8         845827              1301 ns/op             472 B/op         27 allocs/op (Apple M1 Pro)
func BenchmarkEncode(b *testing.B) {
	secSet := map[constants.SectionID]Section{}
	for i := 0; i < len(testData); i++ {
		for _, section := range testData[i].sections {
			if _, ok := secSet[section.GetID()]; ok {
				continue
			}
			secSet[section.GetID()] = section
		}
	}
	var secs []Section
	for _, val := range secSet {
		secs = append(secs, val)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Encode(secs)
		if err != nil {
			b.Fatal(err)
		}
	}
}
