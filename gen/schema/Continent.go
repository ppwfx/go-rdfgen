package schema

type ContinentTrait struct {

}

type Continent struct {
	MetaTrait
	ContinentTrait
	LandformTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewContinent() (x Continent) {
	x.Type = "http://schema.org/Continent"

	return
}
