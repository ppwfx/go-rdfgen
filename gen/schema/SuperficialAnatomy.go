package schema

type SuperficialAnatomyTrait struct {

	// If applicable, a description of the pathophysiology associated with the anatomical system, including potential abnormal changes in the mechanical, physical, and biochemical functions of the system.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AssociatedPathophysiology string `json:"associatedPathophysiology,omitempty" jsonld:"http://schema.org/associatedPathophysiology"`

	// The significance associated with the superficial anatomy; as an example, how characteristics of the superficial anatomy can suggest underlying medical conditions or courses of treatment.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Significance string `json:"significance,omitempty" jsonld:"http://schema.org/significance"`

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

	// Anatomical systems or structures that relate to the superficial anatomy.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	// - http://schema.org/AnatomicalSystem
	//
	RelatedAnatomy interface{} `json:"relatedAnatomy,omitempty" jsonld:"http://schema.org/relatedAnatomy"`

}

type SuperficialAnatomy struct {
	MetaTrait
	SuperficialAnatomyTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewSuperficialAnatomy() (x SuperficialAnatomy) {
	x.Type = "http://schema.org/SuperficialAnatomy"

	return
}
