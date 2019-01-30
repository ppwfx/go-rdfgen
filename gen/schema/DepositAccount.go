package schema

type DepositAccountTrait struct {

}

type DepositAccount struct {
	MetaTrait
	DepositAccountTrait
	InvestmentOrDepositTrait
	FinancialProductTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	BankAccountTrait
	AdditionalTrait
}

func NewDepositAccount() (x DepositAccount) {
	x.Type = "http://schema.org/DepositAccount"

	return
}
