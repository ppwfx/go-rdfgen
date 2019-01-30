package schema

type LiquorStoreTrait struct {

}

type LiquorStore struct {
	MetaTrait
	LiquorStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewLiquorStore() (x LiquorStore) {
	x.Type = "http://schema.org/LiquorStore"

	return
}
