package schema

type OccupationalTherapyTrait struct {

}

type OccupationalTherapy struct {
	MetaTrait
	OccupationalTherapyTrait
	MedicalTherapyTrait
	TherapeuticProcedureTrait
	MedicalProcedureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewOccupationalTherapy() (x OccupationalTherapy) {
	x.Type = "http://schema.org/OccupationalTherapy"

	return
}
