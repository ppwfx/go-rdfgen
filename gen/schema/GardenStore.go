package schema

type GardenStoreTrait struct {

}

type GardenStore struct {
	MetaTrait
	GardenStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewGardenStore() (x GardenStore) {
	x.Type = "http://schema.org/GardenStore"

	return
}
