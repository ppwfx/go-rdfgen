package schema

type MedicalSignOrSymptomTrait struct {

	// Specifying a cause of something in general. e.g in medicine , one of the causative agent(s) that are most directly responsible for the pathophysiologic process that eventually results in the occurrence.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalCause
	//
	Cause *MedicalCause `json:"cause,omitempty" jsonld:"http://schema.org/cause"`

	// A possible treatment to address this condition, sign or symptom.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalTherapy
	//
	PossibleTreatment *MedicalTherapy `json:"possibleTreatment,omitempty" jsonld:"http://schema.org/possibleTreatment"`

}

type MedicalSignOrSymptom struct {
	MetaTrait
	MedicalSignOrSymptomTrait
	MedicalConditionTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalSignOrSymptom() (x MedicalSignOrSymptom) {
	x.Type = "http://schema.org/MedicalSignOrSymptom"

	return
}
