package schema

type PriceSpecificationTrait struct {

	// The date when the item becomes valid.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	ValidFrom *DateTime `json:"validFrom,omitempty" jsonld:"http://schema.org/validFrom"`

	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	ValidThrough *DateTime `json:"validThrough,omitempty" jsonld:"http://schema.org/validThrough"`

	// The interval and unit of measurement of ordering quantities for which the offer or price specification is valid. This allows e.g. specifying that a certain freight charge is valid only for a certain quantity.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	EligibleQuantity *QuantitativeValue `json:"eligibleQuantity,omitempty" jsonld:"http://schema.org/eligibleQuantity"`

	// The transaction volume, in a monetary unit, for which the offer or price specification is valid, e.g. for indicating a minimal purchasing volume, to express free shipping above a certain order volume, or to limit the acceptance of credit cards to purchases to a certain minimal amount.
	//
	// RangeIncludes:
	// - http://schema.org/PriceSpecification
	//
	EligibleTransactionVolume *PriceSpecification `json:"eligibleTransactionVolume,omitempty" jsonld:"http://schema.org/eligibleTransactionVolume"`

	// The highest price if the price is a range.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	MaxPrice float64 `json:"maxPrice,omitempty" jsonld:"http://schema.org/maxPrice"`

	// The lowest price if the price is a range.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	MinPrice float64 `json:"minPrice,omitempty" jsonld:"http://schema.org/minPrice"`

	// The offer price of a product, or of a price component when attached to PriceSpecification and its subtypes.<br/><br/>\n\nUsage guidelines:<br/><br/>\n\n<ul>\n<li>Use the <a class=\"localLink\" href=\"http://schema.org/priceCurrency\">priceCurrency</a> property (with standard formats: <a href=\"http://en.wikipedia.org/wiki/ISO_4217\">ISO 4217 currency format</a> e.g. \"USD\"; <a href=\"https://en.wikipedia.org/wiki/List_of_cryptocurrencies\">Ticker symbol</a> for cryptocurrencies e.g. \"BTC\"; well known names for <a href=\"https://en.wikipedia.org/wiki/Local_exchange_trading_system\">Local Exchange Tradings Systems</a> (LETS) and other currency types e.g. \"Ithaca HOUR\") instead of including <a href=\"http://en.wikipedia.org/wiki/Dollar_sign#Currencies_that_use_the_dollar_or_peso_sign\">ambiguous symbols</a> such as '$' in the value.</li>\n<li>Use '.' (Unicode 'FULL STOP' (U+002E)) rather than ',' to indicate a decimal point. Avoid using these symbols as a readability separator.</li>\n<li>Note that both <a href=\"http://www.w3.org/TR/xhtml-rdfa-primer/#using-the-content-attribute\">RDFa</a> and Microdata syntax allow the use of a \"content=\" attribute for publishing simple machine-readable values alongside more human-friendly formatting.</li>\n<li>Use values from 0123456789 (Unicode 'DIGIT ZERO' (U+0030) to 'DIGIT NINE' (U+0039)) rather than superficially similiar Unicode symbols.</li>\n</ul>\n
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Number
	//
	Price interface{} `json:"price,omitempty" jsonld:"http://schema.org/price"`

	// The currency of the price, or a price component when attached to <a class=\"localLink\" href=\"http://schema.org/PriceSpecification\">PriceSpecification</a> and its subtypes.<br/><br/>\n\nUse standard formats: <a href=\"http://en.wikipedia.org/wiki/ISO_4217\">ISO 4217 currency format</a> e.g. \"USD\"; <a href=\"https://en.wikipedia.org/wiki/List_of_cryptocurrencies\">Ticker symbol</a> for cryptocurrencies e.g. \"BTC\"; well known names for <a href=\"https://en.wikipedia.org/wiki/Local_exchange_trading_system\">Local Exchange Tradings Systems</a> (LETS) and other currency types e.g. \"Ithaca HOUR\".
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PriceCurrency string `json:"priceCurrency,omitempty" jsonld:"http://schema.org/priceCurrency"`

	// Specifies whether the applicable value-added tax (VAT) is included in the price specification or not.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	ValueAddedTaxIncluded bool `json:"valueAddedTaxIncluded,omitempty" jsonld:"http://schema.org/valueAddedTaxIncluded"`

}

type PriceSpecification struct {
	MetaTrait
	PriceSpecificationTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewPriceSpecification() (x PriceSpecification) {
	x.Type = "http://schema.org/PriceSpecification"

	return
}
