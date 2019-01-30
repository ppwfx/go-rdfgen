package schema

type InvestmentFundTrait struct {

}

type InvestmentFund struct {
	MetaTrait
	InvestmentFundTrait
	InvestmentOrDepositTrait
	FinancialProductTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewInvestmentFund() (x InvestmentFund) {
	x.Type = "http://schema.org/InvestmentFund"

	return
}
