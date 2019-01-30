package schema

type MedicalTrialTrait struct {

	// The phase of the clinical trial.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Phase string `json:"phase,omitempty" jsonld:"http://schema.org/phase"`

	// Specifics about the trial design (enumerated).
	//
	// RangeIncludes:
	// - http://schema.org/MedicalTrialDesign
	//
	TrialDesign *MedicalTrialDesign `json:"trialDesign,omitempty" jsonld:"http://schema.org/trialDesign"`

}

type MedicalTrial struct {
	MetaTrait
	MedicalTrialTrait
	MedicalStudyTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalTrial() (x MedicalTrial) {
	x.Type = "http://schema.org/MedicalTrial"

	return
}
