package schema

type MedicalSpecialtyTrait struct {

}

type MedicalSpecialty struct {
	MetaTrait
	MedicalSpecialtyTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	SpecialtyTrait
	AdditionalTrait
}

func NewMedicalSpecialty() (x MedicalSpecialty) {
	x.Type = "http://schema.org/MedicalSpecialty"

	return
}
