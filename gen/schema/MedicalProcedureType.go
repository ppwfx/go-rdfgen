package schema

type MedicalProcedureTypeTrait struct {

}

type MedicalProcedureType struct {
	MetaTrait
	MedicalProcedureTypeTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalProcedureType() (x MedicalProcedureType) {
	x.Type = "http://schema.org/MedicalProcedureType"

	return
}
