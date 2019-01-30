package schema

type RestrictedDietTrait struct {

}

type RestrictedDiet struct {
	MetaTrait
	RestrictedDietTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewRestrictedDiet() (x RestrictedDiet) {
	x.Type = "http://schema.org/RestrictedDiet"

	return
}
