package schema

type TireShopTrait struct {

}

type TireShop struct {
	MetaTrait
	TireShopTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewTireShop() (x TireShop) {
	x.Type = "http://schema.org/TireShop"

	return
}
