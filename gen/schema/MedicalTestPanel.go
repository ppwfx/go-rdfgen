package schema

type MedicalTestPanelTrait struct {

	// A component test of the panel.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalTest
	//
	SubTest *MedicalTest `json:"subTest,omitempty" jsonld:"http://schema.org/subTest"`

}

type MedicalTestPanel struct {
	MetaTrait
	MedicalTestPanelTrait
	MedicalTestTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalTestPanel() (x MedicalTestPanel) {
	x.Type = "http://schema.org/MedicalTestPanel"

	return
}
