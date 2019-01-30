package schema

type HomeGoodsStoreTrait struct {

}

type HomeGoodsStore struct {
	MetaTrait
	HomeGoodsStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewHomeGoodsStore() (x HomeGoodsStore) {
	x.Type = "http://schema.org/HomeGoodsStore"

	return
}
