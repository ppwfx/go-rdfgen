package schema

type MedicalTrialDesignTrait struct {

}

type MedicalTrialDesign struct {
	MetaTrait
	MedicalTrialDesignTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	MedicalEnumerationTrait
	AdditionalTrait
}

func NewMedicalTrialDesign() (x MedicalTrialDesign) {
	x.Type = "http://schema.org/MedicalTrialDesign"

	return
}
