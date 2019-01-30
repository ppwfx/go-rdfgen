package schema

type DDxElementTrait struct {

	// One or more alternative conditions considered in the differential diagnosis process as output of a diagnosis process.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalCondition
	//
	Diagnosis *MedicalCondition `json:"diagnosis,omitempty" jsonld:"http://schema.org/diagnosis"`

	// One of a set of signs and symptoms that can be used to distinguish this diagnosis from others in the differential diagnosis.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalSignOrSymptom
	//
	DistinguishingSign *MedicalSignOrSymptom `json:"distinguishingSign,omitempty" jsonld:"http://schema.org/distinguishingSign"`

}

type DDxElement struct {
	MetaTrait
	DDxElementTrait
	MedicalIntangibleTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewDDxElement() (x DDxElement) {
	x.Type = "http://schema.org/DDxElement"

	return
}
