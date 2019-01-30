package schema

type ItemListOrderTypeTrait struct {

}

type ItemListOrderType struct {
	MetaTrait
	ItemListOrderTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewItemListOrderType() (x ItemListOrderType) {
	x.Type = "http://schema.org/ItemListOrderType"

	return
}
