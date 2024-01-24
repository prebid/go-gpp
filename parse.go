package gpp

import (
	"fmt"
	"strings"

	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/sections/uspca"
	"github.com/prebid/go-gpp/sections/uspco"
	"github.com/prebid/go-gpp/sections/uspct"
	"github.com/prebid/go-gpp/sections/uspnat"
	"github.com/prebid/go-gpp/sections/usput"
	"github.com/prebid/go-gpp/sections/uspva"
	"github.com/prebid/go-gpp/util"
)

const (
	SectionGPPByte      byte = 'D'
	MinHeaderCharacters      = 4
)

type GppContainer struct {
	Version      int
	SectionTypes []constants.SectionID
	Sections     []Section
}

type Section interface {
	GetID() constants.SectionID
	GetValue() string // base64 encoding usually, but plaintext for ccpa
	Encode(bool) []byte
}

func Parse(v string) (GppContainer, []error) {
	var gpp GppContainer

	sectionStrings := strings.Split(v, "~")

	header := sectionStrings[0]
	if err := fastFailHeaderValidate(header); err != nil {
		return gpp, []error{err}
	}

	bs, err := util.NewBitStreamFromBase64(header)
	if err != nil {
		return gpp, []error{fmt.Errorf("error parsing GPP header, base64 decoding: %s", err)}
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

// fastFailHeaderValidate performs quick validations of the header section before decoding
// the bit stream.
func fastFailHeaderValidate(h string) error {
	// the GPP header must be at least 24 bits to represent the type, version, and a fibonacci sequence
	// of at least 1 item. this requires at least 4 characters.
	if len(h) < MinHeaderCharacters {
		return fmt.Errorf("error parsing GPP header, should be at least %d bytes long", MinHeaderCharacters)
	}

	// base64-url encodes 6 bits into each character. the first 6 bits of GPP header must always
	// evaluate to the integer '3', so we can short cut by checking the first character directly.
	if h[0] != SectionGPPByte {
		return fmt.Errorf("error parsing GPP header, header must have type=%d", constants.SectionGPP)
	}

	return nil
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

func (gs GenericSection) Encode(bool) []byte {
	return []byte(gs.value)
}
