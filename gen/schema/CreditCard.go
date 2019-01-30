package schema

type CreditCardTrait struct {

	// The minimum payment is the lowest amount of money that one is required to pay on a credit card statement each month.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	// - http://schema.org/Number
	//
	MonthlyMinimumRepaymentAmount interface{} `json:"monthlyMinimumRepaymentAmount,omitempty" jsonld:"http://schema.org/monthlyMinimumRepaymentAmount"`

}

type CreditCard struct {
	MetaTrait
	CreditCardTrait
	LoanOrCreditTrait
	FinancialProductTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	PaymentCardTrait
	PaymentMethodTrait
	EnumerationTrait
	AdditionalTrait
}

func NewCreditCard() (x CreditCard) {
	x.Type = "http://schema.org/CreditCard"

	return
}
