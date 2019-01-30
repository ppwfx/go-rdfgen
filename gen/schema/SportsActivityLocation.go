package schema

type SportsActivityLocationTrait struct {

}

type SportsActivityLocation struct {
	MetaTrait
	SportsActivityLocationTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewSportsActivityLocation() (x SportsActivityLocation) {
	x.Type = "http://schema.org/SportsActivityLocation"

	return
}
