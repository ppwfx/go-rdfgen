package schema

type AdministrativeAreaTrait struct {

}

type AdministrativeArea struct {
	MetaTrait
	AdministrativeAreaTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewAdministrativeArea() (x AdministrativeArea) {
	x.Type = "http://schema.org/AdministrativeArea"

	return
}
