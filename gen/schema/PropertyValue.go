package schema

type PropertyValueTrait struct {

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

	// A technique or technology used in a <a class=\"localLink\" href=\"http://schema.org/Dataset\">Dataset</a> (or <a class=\"localLink\" href=\"http://schema.org/DataDownload\">DataDownload</a>, <a class=\"localLink\" href=\"http://schema.org/DataCatalog\">DataCatalog</a>),\ncorresponding to the method used for measuring the corresponding variable(s) (described using <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a>). This is oriented towards scientific and scholarly dataset publication but may have broader applicability; it is not intended as a full representation of measurement, but rather as a high level summary for dataset discovery.<br/><br/>\n\nFor example, if <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a> is: molecule concentration, <a class=\"localLink\" href=\"http://schema.org/measurementTechnique\">measurementTechnique</a> could be: \"mass spectrometry\" or \"nmr spectroscopy\" or \"colorimetry\" or \"immunofluorescence\".<br/><br/>\n\nIf the <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a> is \"depression rating\", the <a class=\"localLink\" href=\"http://schema.org/measurementTechnique\">measurementTechnique</a> could be \"Zung Scale\" or \"HAM-D\" or \"Beck Depression Inventory\".<br/><br/>\n\nIf there are several <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a> properties recorded for some given data object, use a <a class=\"localLink\" href=\"http://schema.org/PropertyValue\">PropertyValue</a> for each <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a> and attach the corresponding <a class=\"localLink\" href=\"http://schema.org/measurementTechnique\">measurementTechnique</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	MeasurementTechnique interface{} `json:"measurementTechnique,omitempty" jsonld:"http://schema.org/measurementTechnique"`

	// A commonly used identifier for the characteristic represented by the property, e.g. a manufacturer or a standard code for a property. propertyID can be\n(1) a prefixed string, mainly meant to be used with standards for product properties; (2) a site-specific, non-prefixed string (e.g. the primary key of the property or the vendor-specific id of the property), or (3)\na URL indicating the type of the property, either pointing to an external vocabulary, or a Web resource that describes the property (e.g. a glossary entry).\nStandards bodies should promote a standard prefix for the identifiers of properties from their standards.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	PropertyID interface{} `json:"propertyID,omitempty" jsonld:"http://schema.org/propertyID"`

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

}

type PropertyValue struct {
	MetaTrait
	PropertyValueTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewPropertyValue() (x PropertyValue) {
	x.Type = "http://schema.org/PropertyValue"

	return
}
