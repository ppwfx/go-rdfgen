package schema

type AmusementParkTrait struct {

}

type AmusementPark struct {
	MetaTrait
	AmusementParkTrait
	EntertainmentBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewAmusementPark() (x AmusementPark) {
	x.Type = "http://schema.org/AmusementPark"

	return
}
