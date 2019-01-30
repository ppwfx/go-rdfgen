package schema

type MedicalDevicePurposeTrait struct {

}

type MedicalDevicePurpose struct {
	MetaTrait
	MedicalDevicePurposeTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalDevicePurpose() (x MedicalDevicePurpose) {
	x.Type = "http://schema.org/MedicalDevicePurpose"

	return
}
