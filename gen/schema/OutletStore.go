package schema

type OutletStoreTrait struct {

}

type OutletStore struct {
	MetaTrait
	OutletStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewOutletStore() (x OutletStore) {
	x.Type = "http://schema.org/OutletStore"

	return
}
