package schema

type LoanOrCreditTrait struct {

	// The currency in which the monetary amount is expressed (in 3-letter <a href=\"http://en.wikipedia.org/wiki/ISO_4217\">ISO 4217</a> format).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Currency string `json:"currency,omitempty" jsonld:"http://schema.org/currency"`

	// The amount of money.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	// - http://schema.org/Number
	//
	Amount interface{} `json:"amount,omitempty" jsonld:"http://schema.org/amount"`

	// Assets required to secure loan or credit repayments. It may take form of third party pledge, goods, financial instruments (cash, securities, etc.)
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Thing
	//
	RequiredCollateral interface{} `json:"requiredCollateral,omitempty" jsonld:"http://schema.org/requiredCollateral"`

	// The type of a loan or credit.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	LoanType interface{} `json:"loanType,omitempty" jsonld:"http://schema.org/loanType"`

	// The period of time after any due date that the borrower has to fulfil its obligations before a default (failure to pay) is deemed to have occurred.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	GracePeriod *Duration `json:"gracePeriod,omitempty" jsonld:"http://schema.org/gracePeriod"`

	// A form of paying back money previously borrowed from a lender. Repayment usually takes the form of periodic payments that normally include part principal plus interest in each payment.
	//
	// RangeIncludes:
	// - http://schema.org/RepaymentSpecification
	//
	LoanRepaymentForm *RepaymentSpecification `json:"loanRepaymentForm,omitempty" jsonld:"http://schema.org/loanRepaymentForm"`

	// The duration of the loan or credit agreement.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	LoanTerm *QuantitativeValue `json:"loanTerm,omitempty" jsonld:"http://schema.org/loanTerm"`

	// The only way you get the money back in the event of default is the security. Recourse is where you still have the opportunity to go back to the borrower for the rest of the money.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	RecourseLoan bool `json:"recourseLoan,omitempty" jsonld:"http://schema.org/recourseLoan"`

	// Whether the terms for payment of interest can be renegotiated during the life of the loan.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	RenegotiableLoan bool `json:"renegotiableLoan,omitempty" jsonld:"http://schema.org/renegotiableLoan"`

}

type LoanOrCredit struct {
	MetaTrait
	LoanOrCreditTrait
	FinancialProductTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewLoanOrCredit() (x LoanOrCredit) {
	x.Type = "http://schema.org/LoanOrCredit"

	return
}
