package schema

type MedicalRiskScoreTrait struct {

	// The algorithm or rules to follow to compute the score.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Algorithm string `json:"algorithm,omitempty" jsonld:"http://schema.org/algorithm"`

}

type MedicalRiskScore struct {
	MetaTrait
	MedicalRiskScoreTrait
	MedicalRiskEstimatorTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalRiskScore() (x MedicalRiskScore) {
	x.Type = "http://schema.org/MedicalRiskScore"

	return
}
