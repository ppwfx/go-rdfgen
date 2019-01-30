package schema

type AutoBodyShopTrait struct {

}

type AutoBodyShop struct {
	MetaTrait
	AutoBodyShopTrait
	AutomotiveBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewAutoBodyShop() (x AutoBodyShop) {
	x.Type = "http://schema.org/AutoBodyShop"

	return
}
