package schema

type MedicalRiskEstimatorTrait struct {

	// The condition, complication, or symptom whose risk is being estimated.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalEntity
	//
	EstimatesRiskOf *MedicalEntity `json:"estimatesRiskOf,omitempty" jsonld:"http://schema.org/estimatesRiskOf"`

	// A modifiable or non-modifiable risk factor included in the calculation, e.g. age, coexisting condition.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalRiskFactor
	//
	IncludedRiskFactor *MedicalRiskFactor `json:"includedRiskFactor,omitempty" jsonld:"http://schema.org/includedRiskFactor"`

}

type MedicalRiskEstimator struct {
	MetaTrait
	MedicalRiskEstimatorTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalRiskEstimator() (x MedicalRiskEstimator) {
	x.Type = "http://schema.org/MedicalRiskEstimator"

	return
}
