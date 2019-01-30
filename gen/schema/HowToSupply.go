package schema

type HowToSupplyTrait struct {

	// The estimated cost of the supply or supplies consumed when performing instructions.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/MonetaryAmount
	//
	EstimatedCost interface{} `json:"estimatedCost,omitempty" jsonld:"http://schema.org/estimatedCost"`

}

type HowToSupply struct {
	MetaTrait
	HowToSupplyTrait
	HowToItemTrait
	ListItemTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewHowToSupply() (x HowToSupply) {
	x.Type = "http://schema.org/HowToSupply"

	return
}
