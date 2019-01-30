package schema

type MortgageLoanTrait struct {

	// Amount of mortgage mandate that can be converted into a proper mortgage at a later stage.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	//
	LoanMortgageMandateAmount *MonetaryAmount `json:"loanMortgageMandateAmount,omitempty" jsonld:"http://schema.org/loanMortgageMandateAmount"`

	// Whether borrower is a resident of the jurisdiction where the property is located.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	DomiciledMortgage bool `json:"domiciledMortgage,omitempty" jsonld:"http://schema.org/domiciledMortgage"`

}

type MortgageLoan struct {
	MetaTrait
	MortgageLoanTrait
	LoanOrCreditTrait
	FinancialProductTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMortgageLoan() (x MortgageLoan) {
	x.Type = "http://schema.org/MortgageLoan"

	return
}
