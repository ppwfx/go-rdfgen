package schema

type CountryTrait struct {

}

type Country struct {
	MetaTrait
	CountryTrait
	AdministrativeAreaTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewCountry() (x Country) {
	x.Type = "http://schema.org/Country"

	return
}
