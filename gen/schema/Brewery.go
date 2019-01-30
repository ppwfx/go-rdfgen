package schema

type BreweryTrait struct {

}

type Brewery struct {
	MetaTrait
	BreweryTrait
	FoodEstablishmentTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewBrewery() (x Brewery) {
	x.Type = "http://schema.org/Brewery"

	return
}
