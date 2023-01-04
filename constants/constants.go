package constants

type SectionID int

const (
	SectionTCFEU2 SectionID = 2
	SectionGPP    SectionID = 3
	SectionUSPV1  SectionID = 6
	SectionUSPNAT SectionID = 7
	SectionUSPCA  SectionID = 8
	SectionUSPVA  SectionID = 9
	SectionUSPCO  SectionID = 10
	SectionUSPUT  SectionID = 11
	SectionUSPCT  SectionID = 12
)

var SectionIDNames = []string{"ID0", "ID1", "tcfeu2", "gpp header", "ID4", "ID5", "uspv1", "uspnat",
	"uspca", "uspva", "uspco", "usput", "uspct"}
