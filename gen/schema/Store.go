package schema

type StoreTrait struct {

}

type Store struct {
	MetaTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewStore() (x Store) {
	x.Type = "http://schema.org/Store"

	return
}
