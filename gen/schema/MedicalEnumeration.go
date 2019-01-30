package schema

type MedicalEnumerationTrait struct {

}

type MedicalEnumeration struct {
	MetaTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalEnumeration() (x MedicalEnumeration) {
	x.Type = "http://schema.org/MedicalEnumeration"

	return
}
