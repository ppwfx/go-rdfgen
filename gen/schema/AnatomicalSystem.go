package schema

type AnatomicalSystemTrait struct {

	// If applicable, a description of the pathophysiology associated with the anatomical system, including potential abnormal changes in the mechanical, physical, and biochemical functions of the system.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AssociatedPathophysiology string `json:"associatedPathophysiology,omitempty" jsonld:"http://schema.org/associatedPathophysiology"`

	// A medical condition associated with this anatomy.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalCondition
	//
	RelatedCondition *MedicalCondition `json:"relatedCondition,omitempty" jsonld:"http://schema.org/relatedCondition"`

	// A medical therapy related to this anatomy.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalTherapy
	//
	RelatedTherapy *MedicalTherapy `json:"relatedTherapy,omitempty" jsonld:"http://schema.org/relatedTherapy"`

	// Specifying something physically contained by something else. Typically used here for the underlying anatomical structures, such as organs, that comprise the anatomical system.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	// - http://schema.org/AnatomicalSystem
	//
	ComprisedOf interface{} `json:"comprisedOf,omitempty" jsonld:"http://schema.org/comprisedOf"`

	// Related anatomical structure(s) that are not part of the system but relate or connect to it, such as vascular bundles associated with an organ system.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	//
	RelatedStructure *AnatomicalStructure `json:"relatedStructure,omitempty" jsonld:"http://schema.org/relatedStructure"`

}

type AnatomicalSystem struct {
	MetaTrait
	AnatomicalSystemTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewAnatomicalSystem() (x AnatomicalSystem) {
	x.Type = "http://schema.org/AnatomicalSystem"

	return
}
