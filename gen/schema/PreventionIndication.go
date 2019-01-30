package schema

type PreventionIndicationTrait struct {

}

type PreventionIndication struct {
	MetaTrait
	PreventionIndicationTrait
	MedicalIndicationTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewPreventionIndication() (x PreventionIndication) {
	x.Type = "http://schema.org/PreventionIndication"

	return
}
