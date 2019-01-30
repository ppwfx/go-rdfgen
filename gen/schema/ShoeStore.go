package schema

type ShoeStoreTrait struct {

}

type ShoeStore struct {
	MetaTrait
	ShoeStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewShoeStore() (x ShoeStore) {
	x.Type = "http://schema.org/ShoeStore"

	return
}
