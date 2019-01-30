package schema

type HobbyShopTrait struct {

}

type HobbyShop struct {
	MetaTrait
	HobbyShopTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewHobbyShop() (x HobbyShop) {
	x.Type = "http://schema.org/HobbyShop"

	return
}
