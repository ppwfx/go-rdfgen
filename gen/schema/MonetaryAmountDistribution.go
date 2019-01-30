package schema

type MonetaryAmountDistributionTrait struct {

	// The currency in which the monetary amount is expressed (in 3-letter <a href=\"http://en.wikipedia.org/wiki/ISO_4217\">ISO 4217</a> format).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Currency string `json:"currency,omitempty" jsonld:"http://schema.org/currency"`

}

type MonetaryAmountDistribution struct {
	MetaTrait
	MonetaryAmountDistributionTrait
	QuantitativeValueTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMonetaryAmountDistribution() (x MonetaryAmountDistribution) {
	x.Type = "http://schema.org/MonetaryAmountDistribution"

	return
}
