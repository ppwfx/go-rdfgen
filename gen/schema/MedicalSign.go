package schema

type MedicalSignTrait struct {

	// A physical examination that can identify this sign.
	//
	// RangeIncludes:
	// - http://schema.org/PhysicalExam
	//
	IdentifyingExam *PhysicalExam `json:"identifyingExam,omitempty" jsonld:"http://schema.org/identifyingExam"`

	// A diagnostic test that can identify this sign.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalTest
	//
	IdentifyingTest *MedicalTest `json:"identifyingTest,omitempty" jsonld:"http://schema.org/identifyingTest"`

}

type MedicalSign struct {
	MetaTrait
	MedicalSignTrait
	MedicalSignOrSymptomTrait
	MedicalConditionTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalSign() (x MedicalSign) {
	x.Type = "http://schema.org/MedicalSign"

	return
}
