package schema

type MonetaryAmountTrait struct {

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

	// The currency in which the monetary amount is expressed (in 3-letter <a href=\"http://en.wikipedia.org/wiki/ISO_4217\">ISO 4217</a> format).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Currency string `json:"currency,omitempty" jsonld:"http://schema.org/currency"`

	// The upper value of some characteristic or property.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	MaxValue float64 `json:"maxValue,omitempty" jsonld:"http://schema.org/maxValue"`

	// The lower value of some characteristic or property.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	MinValue float64 `json:"minValue,omitempty" jsonld:"http://schema.org/minValue"`

	// The value of the quantitative value or property value node.<br/><br/>\n\n<ul>\n<li>For <a class=\"localLink\" href=\"http://schema.org/QuantitativeValue\">QuantitativeValue</a> and <a class=\"localLink\" href=\"http://schema.org/MonetaryAmount\">MonetaryAmount</a>, the recommended type for values is 'Number'.</li>\n<li>For <a class=\"localLink\" href=\"http://schema.org/PropertyValue\">PropertyValue</a>, it can be 'Text;', 'Number', 'Boolean', or 'StructuredValue'.</li>\n</ul>\n
	//
	// RangeIncludes:
	// - http://schema.org/StructuredValue
	// - http://schema.org/Text
	// - http://schema.org/Boolean
	// - http://schema.org/Number
	//
	Value interface{} `json:"value,omitempty" jsonld:"http://schema.org/value"`

}

type MonetaryAmount struct {
	MetaTrait
	MonetaryAmountTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMonetaryAmount() (x MonetaryAmount) {
	x.Type = "http://schema.org/MonetaryAmount"

	return
}
