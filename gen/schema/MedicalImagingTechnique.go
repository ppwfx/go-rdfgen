package schema

type MedicalImagingTechniqueTrait struct {

}

type MedicalImagingTechnique struct {
	MetaTrait
	MedicalImagingTechniqueTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalImagingTechnique() (x MedicalImagingTechnique) {
	x.Type = "http://schema.org/MedicalImagingTechnique"

	return
}
