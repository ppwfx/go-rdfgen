package schema

type CityTrait struct {

}

type City struct {
	MetaTrait
	CityTrait
	AdministrativeAreaTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewCity() (x City) {
	x.Type = "http://schema.org/City"

	return
}
