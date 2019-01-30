package schema

type PaymentMethodTrait struct {

}

type PaymentMethod struct {
	MetaTrait
	PaymentMethodTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewPaymentMethod() (x PaymentMethod) {
	x.Type = "http://schema.org/PaymentMethod"

	return
}
