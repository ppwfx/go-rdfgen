package schema

type DatedMoneySpecificationTrait struct {

	// The end date and time of the item (in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>).
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	// - http://schema.org/Date
	//
	EndDate interface{} `json:"endDate,omitempty" jsonld:"http://schema.org/endDate"`

	// The start date and time of the item (in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>).
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	// - http://schema.org/Date
	//
	StartDate interface{} `json:"startDate,omitempty" jsonld:"http://schema.org/startDate"`

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

}

type DatedMoneySpecification struct {
	MetaTrait
	DatedMoneySpecificationTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDatedMoneySpecification() (x DatedMoneySpecification) {
	x.Type = "http://schema.org/DatedMoneySpecification"

	return
}
