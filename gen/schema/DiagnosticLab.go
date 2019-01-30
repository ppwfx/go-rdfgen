package schema

type DiagnosticLabTrait struct {

	// A diagnostic test or procedure offered by this lab.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalTest
	//
	AvailableTest *MedicalTest `json:"availableTest,omitempty" jsonld:"http://schema.org/availableTest"`

}

type DiagnosticLab struct {
	MetaTrait
	DiagnosticLabTrait
	MedicalOrganizationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewDiagnosticLab() (x DiagnosticLab) {
	x.Type = "http://schema.org/DiagnosticLab"

	return
}
