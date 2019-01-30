package schema

type BankAccountTrait struct {

	// The type of a bank account.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	BankAccountType interface{} `json:"bankAccountType,omitempty" jsonld:"http://schema.org/bankAccountType"`

	// An overdraft is an extension of credit from a lending institution when an account reaches zero. An overdraft allows the individual to continue withdrawing money even if the account has no funds in it. Basically the bank allows people to borrow a set amount of money.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	//
	AccountOverdraftLimit *MonetaryAmount `json:"accountOverdraftLimit,omitempty" jsonld:"http://schema.org/accountOverdraftLimit"`

	// A minimum amount that has to be paid in every month.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	//
	AccountMinimumInflow *MonetaryAmount `json:"accountMinimumInflow,omitempty" jsonld:"http://schema.org/accountMinimumInflow"`

}

type BankAccount struct {
	MetaTrait
	BankAccountTrait
	FinancialProductTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBankAccount() (x BankAccount) {
	x.Type = "http://schema.org/BankAccount"

	return
}
