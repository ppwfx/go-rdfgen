package schema

type JointTrait struct {

	// The biomechanical properties of the bone.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	BiomechnicalClass string `json:"biomechnicalClass,omitempty" jsonld:"http://schema.org/biomechnicalClass"`

	// The name given to how bone physically connects to each other.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	StructuralClass string `json:"structuralClass,omitempty" jsonld:"http://schema.org/structuralClass"`

	// The degree of mobility the joint allows.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/MedicalEntity
	//
	FunctionalClass interface{} `json:"functionalClass,omitempty" jsonld:"http://schema.org/functionalClass"`

}

type Joint struct {
	MetaTrait
	JointTrait
	AnatomicalStructureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewJoint() (x Joint) {
	x.Type = "http://schema.org/Joint"

	return
}
