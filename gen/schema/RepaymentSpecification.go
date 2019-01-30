package schema

type RepaymentSpecificationTrait struct {

	// a type of payment made in cash during the onset of the purchase of an expensive good/service. The payment typically represents only a percentage of the full purchase price.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	// - http://schema.org/Number
	//
	DownPayment interface{} `json:"downPayment,omitempty" jsonld:"http://schema.org/downPayment"`

	// The amount to be paid as a penalty in the event of early payment of the loan.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	//
	EarlyPrepaymentPenalty *MonetaryAmount `json:"earlyPrepaymentPenalty,omitempty" jsonld:"http://schema.org/earlyPrepaymentPenalty"`

	// The amount of money to pay in a single payment.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	//
	LoanPaymentAmount *MonetaryAmount `json:"loanPaymentAmount,omitempty" jsonld:"http://schema.org/loanPaymentAmount"`

	// Frequency of payments due, i.e. number of months between payments. This is defined as a frequency, i.e. the reciprocal of a period of time.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	LoanPaymentFrequency float64 `json:"loanPaymentFrequency,omitempty" jsonld:"http://schema.org/loanPaymentFrequency"`

	// The number of payments contractually required at origination to repay the loan. For monthly paying loans this is the number of months from the contractual first payment date to the maturity date.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	NumberOfLoanPayments float64 `json:"numberOfLoanPayments,omitempty" jsonld:"http://schema.org/numberOfLoanPayments"`

}

type RepaymentSpecification struct {
	MetaTrait
	RepaymentSpecificationTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewRepaymentSpecification() (x RepaymentSpecification) {
	x.Type = "http://schema.org/RepaymentSpecification"

	return
}
