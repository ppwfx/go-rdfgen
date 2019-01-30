package schema

type WholesaleStoreTrait struct {

}

type WholesaleStore struct {
	MetaTrait
	WholesaleStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewWholesaleStore() (x WholesaleStore) {
	x.Type = "http://schema.org/WholesaleStore"

	return
}
