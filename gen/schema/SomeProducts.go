package schema

type SomeProductsTrait struct {

	// The current approximate inventory level for the item or items.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	InventoryLevel *QuantitativeValue `json:"inventoryLevel,omitempty" jsonld:"http://schema.org/inventoryLevel"`

}

type SomeProducts struct {
	MetaTrait
	SomeProductsTrait
	ProductTrait
	ThingTrait
	AdditionalTrait
}

func NewSomeProducts() (x SomeProducts) {
	x.Type = "http://schema.org/SomeProducts"

	return
}
