package schema

type FurnitureStoreTrait struct {

}

type FurnitureStore struct {
	MetaTrait
	FurnitureStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewFurnitureStore() (x FurnitureStore) {
	x.Type = "http://schema.org/FurnitureStore"

	return
}
