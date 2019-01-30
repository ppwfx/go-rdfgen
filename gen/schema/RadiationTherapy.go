package schema

type RadiationTherapyTrait struct {

}

type RadiationTherapy struct {
	MetaTrait
	RadiationTherapyTrait
	MedicalTherapyTrait
	TherapeuticProcedureTrait
	MedicalProcedureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewRadiationTherapy() (x RadiationTherapy) {
	x.Type = "http://schema.org/RadiationTherapy"

	return
}
