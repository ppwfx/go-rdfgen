package schema

type ClothingStoreTrait struct {

}

type ClothingStore struct {
	MetaTrait
	ClothingStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewClothingStore() (x ClothingStore) {
	x.Type = "http://schema.org/ClothingStore"

	return
}
