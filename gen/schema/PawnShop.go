package schema

type PawnShopTrait struct {

}

type PawnShop struct {
	MetaTrait
	PawnShopTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewPawnShop() (x PawnShop) {
	x.Type = "http://schema.org/PawnShop"

	return
}
