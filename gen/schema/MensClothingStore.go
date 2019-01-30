package schema

type MensClothingStoreTrait struct {

}

type MensClothingStore struct {
	MetaTrait
	MensClothingStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewMensClothingStore() (x MensClothingStore) {
	x.Type = "http://schema.org/MensClothingStore"

	return
}
