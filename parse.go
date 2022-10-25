package gpp

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/prebid/go-gpp/constants"
	"github.com/prebid/go-gpp/util"
)

type GppContainer struct {
	Version        int
	Sectiontypes   []constants.SectionID
	Sections       []Section
	SectionStrings []string
}

type Section interface {
	GetID() constants.SectionID
	GetValue() string // base64 encoding usually, but plaintext for ccpa
}

func Parse(v string) (GppContainer, error) {
	var gpp GppContainer

	gpp.SectionStrings = strings.Split(v, "~")

	buff := []byte(gpp.SectionStrings[0])
	decoded := make([]byte, base64.RawURLEncoding.DecodedLen(len(buff)))
	n, err := base64.RawURLEncoding.Decode(decoded, buff)
	if err != nil {
		return gpp, err
	}
	decoded = decoded[:n:n]

	bs := util.NewBitStream(decoded)

	gppType, err := util.ReadByte6(bs)
	if err != nil {
		return gpp, err
	}
	if gppType != 3 {
		return gpp, fmt.Errorf("GPP Parse: a GPP string header must have type=3, got %d", gppType)
	}

	ver, err := util.ReadByte6(bs)
	if err != nil {
		return gpp, err
	}
	fmt.Printf("Version is %d\n", int(ver))
	gpp.Version = int(ver)

	intRange, err := util.ReadFibonacciRange(bs)
	if err != nil {
		return gpp, err
	}

	secIDs := make([]constants.SectionID, len(intRange.Range)+1)
	secIDs[0] = 3 // GPP Header section has an "ID" of 3

	for i, sec := range intRange.Range {
		if sec.StartID != sec.EndID {
			return gpp, fmt.Errorf("GPP Parse: Sections Range(Int) contains a range per entry")
		}
		secIDs[i+1] = constants.SectionID(sec.StartID)
	}
	gpp.Sectiontypes = secIDs

	return gpp, nil
}
