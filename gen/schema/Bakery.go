package schema

type BakeryTrait struct {

}

type Bakery struct {
	MetaTrait
	BakeryTrait
	FoodEstablishmentTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewBakery() (x Bakery) {
	x.Type = "http://schema.org/Bakery"

	return
}
