package schema

type RecommendedDoseScheduleTrait struct {

}

type RecommendedDoseSchedule struct {
	MetaTrait
	RecommendedDoseScheduleTrait
	DoseScheduleTrait
	MedicalIntangibleTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewRecommendedDoseSchedule() (x RecommendedDoseSchedule) {
	x.Type = "http://schema.org/RecommendedDoseSchedule"

	return
}
