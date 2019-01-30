package schema

type GroceryStoreTrait struct {

}

type GroceryStore struct {
	MetaTrait
	GroceryStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewGroceryStore() (x GroceryStore) {
	x.Type = "http://schema.org/GroceryStore"

	return
}
