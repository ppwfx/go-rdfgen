package schema

type PaymentChargeSpecificationTrait struct {

	// The delivery method(s) to which the delivery charge or payment charge specification applies.
	//
	// RangeIncludes:
	// - http://schema.org/DeliveryMethod
	//
	AppliesToDeliveryMethod *DeliveryMethod `json:"appliesToDeliveryMethod,omitempty" jsonld:"http://schema.org/appliesToDeliveryMethod"`

	// The payment method(s) to which the payment charge specification applies.
	//
	// RangeIncludes:
	// - http://schema.org/PaymentMethod
	//
	AppliesToPaymentMethod *PaymentMethod `json:"appliesToPaymentMethod,omitempty" jsonld:"http://schema.org/appliesToPaymentMethod"`

}

type PaymentChargeSpecification struct {
	MetaTrait
	PaymentChargeSpecificationTrait
	PriceSpecificationTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewPaymentChargeSpecification() (x PaymentChargeSpecification) {
	x.Type = "http://schema.org/PaymentChargeSpecification"

	return
}
