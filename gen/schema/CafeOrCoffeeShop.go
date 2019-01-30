package schema

type CafeOrCoffeeShopTrait struct {

}

type CafeOrCoffeeShop struct {
	MetaTrait
	CafeOrCoffeeShopTrait
	FoodEstablishmentTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewCafeOrCoffeeShop() (x CafeOrCoffeeShop) {
	x.Type = "http://schema.org/CafeOrCoffeeShop"

	return
}
