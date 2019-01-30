package schema

type SurgicalProcedureTrait struct {

}

type SurgicalProcedure struct {
	MetaTrait
	SurgicalProcedureTrait
	MedicalProcedureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewSurgicalProcedure() (x SurgicalProcedure) {
	x.Type = "http://schema.org/SurgicalProcedure"

	return
}
