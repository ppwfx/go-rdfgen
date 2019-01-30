package schema

type PalliativeProcedureTrait struct {

}

type PalliativeProcedure struct {
	MetaTrait
	PalliativeProcedureTrait
	MedicalProcedureTrait
	MedicalEntityTrait
	ThingTrait
	MedicalTherapyTrait
	TherapeuticProcedureTrait
	AdditionalTrait
}

func NewPalliativeProcedure() (x PalliativeProcedure) {
	x.Type = "http://schema.org/PalliativeProcedure"

	return
}
