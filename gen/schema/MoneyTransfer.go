package schema

type MoneyTransferTrait struct {

	// The amount of money.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	// - http://schema.org/Number
	//
	Amount interface{} `json:"amount,omitempty" jsonld:"http://schema.org/amount"`

	// A bank or bank’s branch, financial institution or international financial institution operating the beneficiary’s bank account or releasing funds for the beneficiary
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/BankOrCreditUnion
	//
	BeneficiaryBank interface{} `json:"beneficiaryBank,omitempty" jsonld:"http://schema.org/beneficiaryBank"`

}

type MoneyTransfer struct {
	MetaTrait
	MoneyTransferTrait
	TransferActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewMoneyTransfer() (x MoneyTransfer) {
	x.Type = "http://schema.org/MoneyTransfer"

	return
}
