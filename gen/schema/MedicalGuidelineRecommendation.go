package schema

type MedicalGuidelineRecommendationTrait struct {

	// Strength of the guideline's recommendation (e.g. 'class I').
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	RecommendationStrength string `json:"recommendationStrength,omitempty" jsonld:"http://schema.org/recommendationStrength"`

}

type MedicalGuidelineRecommendation struct {
	MetaTrait
	MedicalGuidelineRecommendationTrait
	MedicalGuidelineTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalGuidelineRecommendation() (x MedicalGuidelineRecommendation) {
	x.Type = "http://schema.org/MedicalGuidelineRecommendation"

	return
}
