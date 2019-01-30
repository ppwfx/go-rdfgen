package schema

type PhysicalTherapyTrait struct {

}

type PhysicalTherapy struct {
	MetaTrait
	PhysicalTherapyTrait
	MedicalTherapyTrait
	TherapeuticProcedureTrait
	MedicalProcedureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewPhysicalTherapy() (x PhysicalTherapy) {
	x.Type = "http://schema.org/PhysicalTherapy"

	return
}
