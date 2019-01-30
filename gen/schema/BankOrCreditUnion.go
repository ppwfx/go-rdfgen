package schema

type BankOrCreditUnionTrait struct {

}

type BankOrCreditUnion struct {
	MetaTrait
	BankOrCreditUnionTrait
	FinancialServiceTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewBankOrCreditUnion() (x BankOrCreditUnion) {
	x.Type = "http://schema.org/BankOrCreditUnion"

	return
}
