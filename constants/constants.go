package constants

type SectionID int

const (
	SectionTCFEU2 SectionID = 2
	SectionGPP    SectionID = 3
	SectionTCFCA  SectionID = 5
	SectionUSPV1  SectionID = 6
	SectionUSPNAT SectionID = 7
	SectionUSPCA  SectionID = 8
	SectionUSPVA  SectionID = 9
	SectionUSPCO  SectionID = 10
	SectionUSPUT  SectionID = 11
	SectionUSPCT  SectionID = 12
)

var SectionNamesByID = map[int]string{
	2:  "tcfeu2",
	3:  "gpp header",
	4:  "gpp signal integrity",
	5:  "tcfca",
	6:  "uspv1",
	7:  "uspnat",
	8:  "uspca",
	9:  "uspva",
	10: "uspco",
	11: "usput",
	12: "uspct",
}
