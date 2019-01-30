package schema

type HowToItemTrait struct {

	// The required quantity of the item(s).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Number
	//
	RequiredQuantity interface{} `json:"requiredQuantity,omitempty" jsonld:"http://schema.org/requiredQuantity"`

}

type HowToItem struct {
	MetaTrait
	HowToItemTrait
	ListItemTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewHowToItem() (x HowToItem) {
	x.Type = "http://schema.org/HowToItem"

	return
}
