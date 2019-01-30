package schema

type JewelryStoreTrait struct {

}

type JewelryStore struct {
	MetaTrait
	JewelryStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewJewelryStore() (x JewelryStore) {
	x.Type = "http://schema.org/JewelryStore"

	return
}
