package schema

type ExchangeRateSpecificationTrait struct {

	// The currency in which the monetary amount is expressed (in 3-letter <a href=\"http://en.wikipedia.org/wiki/ISO_4217\">ISO 4217</a> format).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Currency string `json:"currency,omitempty" jsonld:"http://schema.org/currency"`

	// The current price of a currency.
	//
	// RangeIncludes:
	// - http://schema.org/UnitPriceSpecification
	//
	CurrentExchangeRate *UnitPriceSpecification `json:"currentExchangeRate,omitempty" jsonld:"http://schema.org/currentExchangeRate"`

	// The difference between the price at which a broker or other intermediary buys and sells foreign currency.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	// - http://schema.org/Number
	//
	ExchangeRateSpread interface{} `json:"exchangeRateSpread,omitempty" jsonld:"http://schema.org/exchangeRateSpread"`

}

type ExchangeRateSpecification struct {
	MetaTrait
	ExchangeRateSpecificationTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewExchangeRateSpecification() (x ExchangeRateSpecification) {
	x.Type = "http://schema.org/ExchangeRateSpecification"

	return
}
