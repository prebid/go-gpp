package gpp

import (
	"fmt"
	"strings"

	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/util"
)

type GppContainer struct {
	Version      int
	SectionTypes []constants.SectionID
	Sections     []Section
}

type Section interface {
	GetID() constants.SectionID
	GetValue() string // base64 encoding usually, but plaintext for ccpa
}

func Parse(v string) (GppContainer, error) {
	var gpp GppContainer

	sectionStrings := strings.Split(v, "~")

	bs, err := util.NewBitStreamFromBase64(sectionStrings[0])
	if err != nil {
		return gpp, err
	}
	if bs.Len() < 3 {
		return gpp, fmt.Errorf("GPP Parse: a GPP string should be at least 3 bytes long")
	}

	if err != nil {
		return gpp, err
	}
	if sectionStrings[0][0] != constants.SectionGPPByte {
		return gpp, fmt.Errorf("GPP Parse: a GPP string header must have type=%d", constants.SectionGPP)
	}
	bs.SetPosition(6)

	ver, err := bs.ReadByte6()
	if err != nil {
		return gpp, err
	}
	gpp.Version = int(ver)

	intRange, err := bs.ReadFibonacciRange()
	if err != nil {
		return gpp, err
	}

	// We do not count the GPP header as a section
	secCount := len(sectionStrings) - 1
	secIDs := make([]constants.SectionID, 0, secCount)

	for _, sec := range intRange.Range {
		for i := sec.StartID; i <= sec.EndID; i++ {
			secIDs = append(secIDs, constants.SectionID(i))
		}
	}
	if len(secIDs) != secCount {
		return gpp, fmt.Errorf("Section IDs do not match the number of sections: found %d IDs, have %d sections", len(secIDs), secCount)
	}
	gpp.SectionTypes = secIDs

	sections := make([]Section, secCount)
	for i, id := range secIDs {
		switch id {
		default:
			sections[i] = GenericSection{sectionID: id, value: sectionStrings[i+1]}
		}
	}

	gpp.Sections = sections

	return gpp, nil
}

type GenericSection struct {
	sectionID constants.SectionID
	value     string
}

func (gs GenericSection) GetID() constants.SectionID {
	return gs.sectionID
}

func (gs GenericSection) GetValue() string {
	return gs.value
}
