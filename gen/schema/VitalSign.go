package schema

type VitalSignTrait struct {

}

type VitalSign struct {
	MetaTrait
	VitalSignTrait
	MedicalSignTrait
	MedicalSignOrSymptomTrait
	MedicalConditionTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewVitalSign() (x VitalSign) {
	x.Type = "http://schema.org/VitalSign"

	return
}
