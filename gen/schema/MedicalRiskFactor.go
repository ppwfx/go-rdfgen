package schema

type MedicalRiskFactorTrait struct {

	// The condition, complication, etc. influenced by this factor.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalEntity
	//
	IncreasesRiskOf *MedicalEntity `json:"increasesRiskOf,omitempty" jsonld:"http://schema.org/increasesRiskOf"`

}

type MedicalRiskFactor struct {
	MetaTrait
	MedicalRiskFactorTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalRiskFactor() (x MedicalRiskFactor) {
	x.Type = "http://schema.org/MedicalRiskFactor"

	return
}
