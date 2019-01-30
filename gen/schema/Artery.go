package schema

type ArteryTrait struct {

	// The branches that comprise the arterial structure.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	//
	ArterialBranch *AnatomicalStructure `json:"arterialBranch,omitempty" jsonld:"http://schema.org/arterialBranch"`

	// The area to which the artery supplies blood.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	//
	SupplyTo *AnatomicalStructure `json:"supplyTo,omitempty" jsonld:"http://schema.org/supplyTo"`

	// The anatomical or organ system that the artery originates from.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	//
	Source *AnatomicalStructure `json:"source,omitempty" jsonld:"http://schema.org/source"`

}

type Artery struct {
	MetaTrait
	ArteryTrait
	VesselTrait
	AnatomicalStructureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewArtery() (x Artery) {
	x.Type = "http://schema.org/Artery"

	return
}
