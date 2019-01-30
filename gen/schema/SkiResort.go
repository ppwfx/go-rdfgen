package schema

type SkiResortTrait struct {

}

type SkiResort struct {
	MetaTrait
	SkiResortTrait
	SportsActivityLocationTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewSkiResort() (x SkiResort) {
	x.Type = "http://schema.org/SkiResort"

	return
}
