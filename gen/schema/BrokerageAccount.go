package schema

type BrokerageAccountTrait struct {

}

type BrokerageAccount struct {
	MetaTrait
	BrokerageAccountTrait
	InvestmentOrDepositTrait
	FinancialProductTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBrokerageAccount() (x BrokerageAccount) {
	x.Type = "http://schema.org/BrokerageAccount"

	return
}
