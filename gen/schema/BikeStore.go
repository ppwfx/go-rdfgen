package schema

type BikeStoreTrait struct {

}

type BikeStore struct {
	MetaTrait
	BikeStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewBikeStore() (x BikeStore) {
	x.Type = "http://schema.org/BikeStore"

	return
}
