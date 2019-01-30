package schema

type QuantityTrait struct {

}

type Quantity struct {
	MetaTrait
	QuantityTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewQuantity() (x Quantity) {
	x.Type = "http://schema.org/Quantity"

	return
}
