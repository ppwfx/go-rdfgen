package schema

type LymphaticVesselTrait struct {

	// The anatomical or organ system drained by this vessel; generally refers to a specific part of an organ.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	// - http://schema.org/AnatomicalSystem
	//
	RegionDrained interface{} `json:"regionDrained,omitempty" jsonld:"http://schema.org/regionDrained"`

	// The vasculature the lymphatic structure originates, or afferents, from.
	//
	// RangeIncludes:
	// - http://schema.org/Vessel
	//
	OriginatesFrom *Vessel `json:"originatesFrom,omitempty" jsonld:"http://schema.org/originatesFrom"`

	// The vasculature the lymphatic structure runs, or efferents, to.
	//
	// RangeIncludes:
	// - http://schema.org/Vessel
	//
	RunsTo *Vessel `json:"runsTo,omitempty" jsonld:"http://schema.org/runsTo"`

}

type LymphaticVessel struct {
	MetaTrait
	LymphaticVesselTrait
	VesselTrait
	AnatomicalStructureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewLymphaticVessel() (x LymphaticVessel) {
	x.Type = "http://schema.org/LymphaticVessel"

	return
}
