package schema

type PaymentStatusTypeTrait struct {

}

type PaymentStatusType struct {
	MetaTrait
	PaymentStatusTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewPaymentStatusType() (x PaymentStatusType) {
	x.Type = "http://schema.org/PaymentStatusType"

	return
}
