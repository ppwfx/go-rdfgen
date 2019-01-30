package schema

type PaymentServiceTrait struct {

}

type PaymentService struct {
	MetaTrait
	PaymentServiceTrait
	FinancialProductTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewPaymentService() (x PaymentService) {
	x.Type = "http://schema.org/PaymentService"

	return
}
