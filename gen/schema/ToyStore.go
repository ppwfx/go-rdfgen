package schema

type ToyStoreTrait struct {

}

type ToyStore struct {
	MetaTrait
	ToyStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewToyStore() (x ToyStore) {
	x.Type = "http://schema.org/ToyStore"

	return
}
