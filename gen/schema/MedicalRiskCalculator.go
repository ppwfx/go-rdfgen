package schema

type MedicalRiskCalculatorTrait struct {

}

type MedicalRiskCalculator struct {
	MetaTrait
	MedicalRiskCalculatorTrait
	MedicalRiskEstimatorTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalRiskCalculator() (x MedicalRiskCalculator) {
	x.Type = "http://schema.org/MedicalRiskCalculator"

	return
}
