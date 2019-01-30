package schema

type TreatmentIndicationTrait struct {

}

type TreatmentIndication struct {
	MetaTrait
	TreatmentIndicationTrait
	MedicalIndicationTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewTreatmentIndication() (x TreatmentIndication) {
	x.Type = "http://schema.org/TreatmentIndication"

	return
}
