package schema

type AnatomicalStructureTrait struct {

	// An image containing a diagram that illustrates the structure and/or its component substructures and/or connections with other structures.
	//
	// RangeIncludes:
	// - http://schema.org/ImageObject
	//
	Diagram *ImageObject `json:"diagram,omitempty" jsonld:"http://schema.org/diagram"`

	// If applicable, a description of the pathophysiology associated with the anatomical system, including potential abnormal changes in the mechanical, physical, and biochemical functions of the system.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AssociatedPathophysiology string `json:"associatedPathophysiology,omitempty" jsonld:"http://schema.org/associatedPathophysiology"`

	// Function of the anatomical structure.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Function string `json:"function,omitempty" jsonld:"http://schema.org/function"`

	// Location in the body of the anatomical structure.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	BodyLocation string `json:"bodyLocation,omitempty" jsonld:"http://schema.org/bodyLocation"`

	// Other anatomical structures to which this structure is connected.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	//
	ConnectedTo *AnatomicalStructure `json:"connectedTo,omitempty" jsonld:"http://schema.org/connectedTo"`

	// The anatomical or organ system that this structure is part of.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalSystem
	//
	PartOfSystem *AnatomicalSystem `json:"partOfSystem,omitempty" jsonld:"http://schema.org/partOfSystem"`

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

	// Component (sub-)structure(s) that comprise this anatomical structure.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	//
	SubStructure *AnatomicalStructure `json:"subStructure,omitempty" jsonld:"http://schema.org/subStructure"`

}

type AnatomicalStructure struct {
	MetaTrait
	AnatomicalStructureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewAnatomicalStructure() (x AnatomicalStructure) {
	x.Type = "http://schema.org/AnatomicalStructure"

	return
}
