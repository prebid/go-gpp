package uspva

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type uspvaTestData struct {
	description string
	gppString   string
	expected    USPVA
}

func TestUSPVA(t *testing.T) {
	testData := []uspvaTestData{
		{
			description: "Test 1",
			gppString:   "bSFgmiU",
			/*
				011011 01 00 10 00 01 0110000010011010 00 10 01 01
			*/
			expected: USPVA{
				CoreSegment: USPVACoreSegment{
					Version:                         27,
					SharingNotice:                   1,
					SaleOptOutNotice:                0,
					TargetedAdvertisingOptOutNotice: 2,
					SaleOptOut:                      0,
					TargetedAdvertisingOptOut:       1,
					SensitiveDataProcessing: []byte{
						1, 2, 0, 0, 2, 1, 2, 2,
					},
					KnownChildSensitiveDataConsents: 0,
					MspaCoveredTransaction:          2,
					MspaOptOutOptionMode:            1,
					MspaServiceProviderMode:         1,
				},
				SectionID: 9,
				Value:     "bSFgmiU",
			},
		},
	}

	for _, test := range testData {
		result, err := NewUSPVA(test.gppString)

		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}
}
