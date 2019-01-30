package schema

type QuantitativeValueTrait struct {

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

	// A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature.
	//
	// RangeIncludes:
	// - http://schema.org/StructuredValue
	// - http://schema.org/Enumeration
	// - http://schema.org/PropertyValue
	// - http://schema.org/QualitativeValue
	// - http://schema.org/QuantitativeValue
	//
	ValueReference interface{} `json:"valueReference,omitempty" jsonld:"http://schema.org/valueReference"`

	// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.<br/><br/>\n\nNote: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	//
	// RangeIncludes:
	// - http://schema.org/PropertyValue
	//
	AdditionalProperty *PropertyValue `json:"additionalProperty,omitempty" jsonld:"http://schema.org/additionalProperty"`

}

type QuantitativeValue struct {
	MetaTrait
	QuantitativeValueTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewQuantitativeValue() (x QuantitativeValue) {
	x.Type = "http://schema.org/QuantitativeValue"

	return
}
