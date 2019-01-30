package schema

type UnitPriceSpecificationTrait struct {

	// The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	UnitCode interface{} `json:"unitCode,omitempty" jsonld:"http://schema.org/unitCode"`

	// A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for\n<a href='unitCode'>unitCode</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	UnitText string `json:"unitText,omitempty" jsonld:"http://schema.org/unitText"`

	// A short text or acronym indicating multiple price specifications for the same offer, e.g. SRP for the suggested retail price or INVOICE for the invoice price, mostly used in the car industry.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PriceType string `json:"priceType,omitempty" jsonld:"http://schema.org/priceType"`

	// This property specifies the minimal quantity and rounding increment that will be the basis for the billing. The unit of measurement is specified by the unitCode property.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	BillingIncrement float64 `json:"billingIncrement,omitempty" jsonld:"http://schema.org/billingIncrement"`

	// The reference quantity for which a certain price applies, e.g. 1 EUR per 4 kWh of electricity. This property is a replacement for unitOfMeasurement for the advanced cases where the price does not relate to a standard unit.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	ReferenceQuantity *QuantitativeValue `json:"referenceQuantity,omitempty" jsonld:"http://schema.org/referenceQuantity"`

}

type UnitPriceSpecification struct {
	MetaTrait
	UnitPriceSpecificationTrait
	PriceSpecificationTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewUnitPriceSpecification() (x UnitPriceSpecification) {
	x.Type = "http://schema.org/UnitPriceSpecification"

	return
}
