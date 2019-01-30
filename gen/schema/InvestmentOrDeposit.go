package schema

type InvestmentOrDepositTrait struct {

	// The amount of money.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	// - http://schema.org/Number
	//
	Amount interface{} `json:"amount,omitempty" jsonld:"http://schema.org/amount"`

}

type InvestmentOrDeposit struct {
	MetaTrait
	InvestmentOrDepositTrait
	FinancialProductTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewInvestmentOrDeposit() (x InvestmentOrDeposit) {
	x.Type = "http://schema.org/InvestmentOrDeposit"

	return
}
