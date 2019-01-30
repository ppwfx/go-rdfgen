package schema

type FinancialProductTrait struct {

	// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	FeesAndCommissionsSpecification interface{} `json:"feesAndCommissionsSpecification,omitempty" jsonld:"http://schema.org/feesAndCommissionsSpecification"`

	// The annual rate that is charged for borrowing (or made by investing), expressed as a single percentage number that represents the actual yearly cost of funds over the term of a loan. This includes any fees or additional costs associated with the transaction.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Number
	//
	AnnualPercentageRate interface{} `json:"annualPercentageRate,omitempty" jsonld:"http://schema.org/annualPercentageRate"`

	// The interest rate, charged or paid, applicable to the financial product. Note: This is different from the calculated annualPercentageRate.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Number
	//
	InterestRate interface{} `json:"interestRate,omitempty" jsonld:"http://schema.org/interestRate"`

}

type FinancialProduct struct {
	MetaTrait
	FinancialProductTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewFinancialProduct() (x FinancialProduct) {
	x.Type = "http://schema.org/FinancialProduct"

	return
}
