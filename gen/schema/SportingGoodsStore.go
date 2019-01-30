package schema

type SportingGoodsStoreTrait struct {

}

type SportingGoodsStore struct {
	MetaTrait
	SportingGoodsStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewSportingGoodsStore() (x SportingGoodsStore) {
	x.Type = "http://schema.org/SportingGoodsStore"

	return
}
