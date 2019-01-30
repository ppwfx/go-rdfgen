package schema

type MedicalConditionStageTrait struct {

	// The substage, e.g. 'a' for Stage IIIa.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	SubStageSuffix string `json:"subStageSuffix,omitempty" jsonld:"http://schema.org/subStageSuffix"`

	// The stage represented as a number, e.g. 3.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	StageAsNumber float64 `json:"stageAsNumber,omitempty" jsonld:"http://schema.org/stageAsNumber"`

}

type MedicalConditionStage struct {
	MetaTrait
	MedicalConditionStageTrait
	MedicalIntangibleTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalConditionStage() (x MedicalConditionStage) {
	x.Type = "http://schema.org/MedicalConditionStage"

	return
}
