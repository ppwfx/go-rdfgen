package schema

type AutoWashTrait struct {

}

type AutoWash struct {
	MetaTrait
	AutoWashTrait
	AutomotiveBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewAutoWash() (x AutoWash) {
	x.Type = "http://schema.org/AutoWash"

	return
}
