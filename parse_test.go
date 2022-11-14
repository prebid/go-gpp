package gpp

import (
	"testing"

	"github.com/prebid/go-gpp/constants"
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
	}

	for _, test := range testData {
		result, err := Parse(test.gppString)

		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}
}
