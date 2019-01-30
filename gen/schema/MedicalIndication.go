package schema

type MedicalIndicationTrait struct {

}

type MedicalIndication struct {
	MetaTrait
	MedicalIndicationTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalIndication() (x MedicalIndication) {
	x.Type = "http://schema.org/MedicalIndication"

	return
}
