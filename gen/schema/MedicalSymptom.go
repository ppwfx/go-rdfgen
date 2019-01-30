package schema

type MedicalSymptomTrait struct {

}

type MedicalSymptom struct {
	MetaTrait
	MedicalSymptomTrait
	MedicalSignOrSymptomTrait
	MedicalConditionTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalSymptom() (x MedicalSymptom) {
	x.Type = "http://schema.org/MedicalSymptom"

	return
}
