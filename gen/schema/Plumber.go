package schema

type PlumberTrait struct {

}

type Plumber struct {
	MetaTrait
	PlumberTrait
	HomeAndConstructionBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewPlumber() (x Plumber) {
	x.Type = "http://schema.org/Plumber"

	return
}
