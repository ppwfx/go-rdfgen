package schema

type VeinTrait struct {

	// The anatomical or organ system that the vein flows into; a larger structure that the vein connects to.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	//
	Tributary *AnatomicalStructure `json:"tributary,omitempty" jsonld:"http://schema.org/tributary"`

	// The anatomical or organ system drained by this vessel; generally refers to a specific part of an organ.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	// - http://schema.org/AnatomicalSystem
	//
	RegionDrained interface{} `json:"regionDrained,omitempty" jsonld:"http://schema.org/regionDrained"`

	// The vasculature that the vein drains into.
	//
	// RangeIncludes:
	// - http://schema.org/Vessel
	//
	DrainsTo *Vessel `json:"drainsTo,omitempty" jsonld:"http://schema.org/drainsTo"`

}

type Vein struct {
	MetaTrait
	VeinTrait
	VesselTrait
	AnatomicalStructureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewVein() (x Vein) {
	x.Type = "http://schema.org/Vein"

	return
}
