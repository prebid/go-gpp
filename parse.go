package gpp

import (
	"fmt"
	"strings"

	"github.com/revcontent-production/go-gpp/constants"
	"github.com/revcontent-production/go-gpp/sections/uspca"
	"github.com/revcontent-production/go-gpp/sections/uspco"
	"github.com/revcontent-production/go-gpp/sections/uspct"
	"github.com/revcontent-production/go-gpp/sections/uspnat"
	"github.com/revcontent-production/go-gpp/sections/usput"
	"github.com/revcontent-production/go-gpp/sections/uspva"
	"github.com/revcontent-production/go-gpp/util"
)

const (
	SectionGPPByte  byte = 'D'
	MaxHeaderLength      = 3
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

func Parse(v string) (GppContainer, []error) {
	var gpp GppContainer

	sectionStrings := strings.Split(v, "~")

	// IAB's base64 conversion means a 6 bit grouped value can be converted to 8 bit bytes.
	// Any leftover bits <8 would be skipped in normal base64 decoding.
	// Therefore, pad with 6 '0's w/ `A` to ensure that all bits are decoded into bytes.
	headerString := sectionStrings[0]
	for {
		headerString += "A"
		if len(headerString) >= 6 {
			break
		}
	}
	bs, err := util.NewBitStreamFromBase64(headerString)
	if err != nil {
		return gpp, []error{fmt.Errorf("error parsing GPP header, base64 decoding: %s", err)}
	}
	if bs.Len() < MaxHeaderLength {
		return gpp, []error{fmt.Errorf("error parsing GPP header, should be at least %d bytes long", MaxHeaderLength)}
	}

	// base64 encoding codes just 6 bits into each byte. The first 6 bits of the header must always evaluate
	// to the integer '3' as the GPP header type. Short cut the processing of a 6 bit integer with a simple
	// byte comparison to shave off a few CPU cycles.
	if sectionStrings[0][0] != SectionGPPByte {
		return gpp, []error{fmt.Errorf("error parsing GPP header, header must have type=%d", constants.SectionGPP)}
	}
	// We checked the GPP header type above outside of the bitstream framework, so we advance the bit stream past the first 6 bits.
	bs.SetPosition(6)

	ver, err := bs.ReadByte6()
	if err != nil {
		return gpp, []error{fmt.Errorf("error parsing GPP header, unable to parse GPP version: %s", err)}
	}
	gpp.Version = int(ver)

	intRange, err := bs.ReadFibonacciRange()
	if err != nil {
		return gpp, []error{fmt.Errorf("error parsing GPP header, section identifiers: %s", err)}
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
		return gpp, []error{fmt.Errorf("error parsing GPP header, section IDs do not match the number of sections: found %d IDs, have %d sections", len(secIDs), secCount)}
	}
	gpp.SectionTypes = secIDs

	sections := make([]Section, secCount)
	var errs []error
	for i, id := range secIDs {
		switch id {
		case constants.SectionUSPNAT:
			sections[i], err = uspnat.NewUSPNAT(sectionStrings[i+1])
			if err != nil {
				errs = append(errs, fmt.Errorf("error parsing %s consent string: %s", constants.SectionNamesByID[int(id)], err))
			}
		case constants.SectionUSPCA:
			sections[i], err = uspca.NewUSPCA(sectionStrings[i+1])
			if err != nil {
				errs = append(errs, fmt.Errorf("error parsing %s consent string: %s", constants.SectionNamesByID[int(id)], err))
			}
		case constants.SectionUSPVA:
			sections[i], err = uspva.NewUSPVA(sectionStrings[i+1])
			if err != nil {
				errs = append(errs, fmt.Errorf("error parsing %s consent string: %s", constants.SectionNamesByID[int(id)], err))
			}
		case constants.SectionUSPCO:
			sections[i], err = uspco.NewUSPCO(sectionStrings[i+1])
			if err != nil {
				errs = append(errs, fmt.Errorf("error parsing %s consent string: %s", constants.SectionNamesByID[int(id)], err))
			}
		case constants.SectionUSPUT:
			sections[i], err = usput.NewUSPUT(sectionStrings[i+1])
			if err != nil {
				errs = append(errs, fmt.Errorf("error parsing %s consent string: %s", constants.SectionNamesByID[int(id)], err))
			}
		case constants.SectionUSPCT:
			sections[i], err = uspct.NewUSPCT(sectionStrings[i+1])
			if err != nil {
				errs = append(errs, fmt.Errorf("error parsing %s consent string: %s", constants.SectionNamesByID[int(id)], err))
			}
		default:
			sections[i] = GenericSection{sectionID: id, value: sectionStrings[i+1]}
			if err != nil {
				errs = append(errs, fmt.Errorf("error parsing unsupported (section %d) consent string: %s", int(id), err))
			}
		}
	}

	gpp.Sections = sections

	return gpp, errs
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
