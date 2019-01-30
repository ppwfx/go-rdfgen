package schema

type IceCreamShopTrait struct {

}

type IceCreamShop struct {
	MetaTrait
	IceCreamShopTrait
	FoodEstablishmentTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewIceCreamShop() (x IceCreamShop) {
	x.Type = "http://schema.org/IceCreamShop"

	return
}
