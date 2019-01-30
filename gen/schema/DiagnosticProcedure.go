package schema

type DiagnosticProcedureTrait struct {

}

type DiagnosticProcedure struct {
	MetaTrait
	DiagnosticProcedureTrait
	MedicalProcedureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewDiagnosticProcedure() (x DiagnosticProcedure) {
	x.Type = "http://schema.org/DiagnosticProcedure"

	return
}
