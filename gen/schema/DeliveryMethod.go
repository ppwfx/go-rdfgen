package schema

type DeliveryMethodTrait struct {

}

type DeliveryMethod struct {
	MetaTrait
	DeliveryMethodTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDeliveryMethod() (x DeliveryMethod) {
	x.Type = "http://schema.org/DeliveryMethod"

	return
}
