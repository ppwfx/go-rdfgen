package schema

type PsychologicalTreatmentTrait struct {

}

type PsychologicalTreatment struct {
	MetaTrait
	PsychologicalTreatmentTrait
	TherapeuticProcedureTrait
	MedicalProcedureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewPsychologicalTreatment() (x PsychologicalTreatment) {
	x.Type = "http://schema.org/PsychologicalTreatment"

	return
}
