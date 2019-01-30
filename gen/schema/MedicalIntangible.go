package schema

type MedicalIntangibleTrait struct {

}

type MedicalIntangible struct {
	MetaTrait
	MedicalIntangibleTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalIntangible() (x MedicalIntangible) {
	x.Type = "http://schema.org/MedicalIntangible"

	return
}
