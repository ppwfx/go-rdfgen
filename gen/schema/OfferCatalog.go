package schema

type OfferCatalogTrait struct {

}

type OfferCatalog struct {
	MetaTrait
	OfferCatalogTrait
	ItemListTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewOfferCatalog() (x OfferCatalog) {
	x.Type = "http://schema.org/OfferCatalog"

	return
}
