package schema

type ApprovedIndicationTrait struct {

}

type ApprovedIndication struct {
	MetaTrait
	ApprovedIndicationTrait
	MedicalIndicationTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewApprovedIndication() (x ApprovedIndication) {
	x.Type = "http://schema.org/ApprovedIndication"

	return
}
