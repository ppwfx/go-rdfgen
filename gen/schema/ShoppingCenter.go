package schema

type ShoppingCenterTrait struct {

}

type ShoppingCenter struct {
	MetaTrait
	ShoppingCenterTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewShoppingCenter() (x ShoppingCenter) {
	x.Type = "http://schema.org/ShoppingCenter"

	return
}
