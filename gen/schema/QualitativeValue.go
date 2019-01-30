package schema

type QualitativeValueTrait struct {

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

	// This ordering relation for qualitative values indicates that the subject is equal to the object.
	//
	// RangeIncludes:
	// - http://schema.org/QualitativeValue
	//
	Equal *QualitativeValue `json:"equal,omitempty" jsonld:"http://schema.org/equal"`

	// This ordering relation for qualitative values indicates that the subject is greater than the object.
	//
	// RangeIncludes:
	// - http://schema.org/QualitativeValue
	//
	Greater *QualitativeValue `json:"greater,omitempty" jsonld:"http://schema.org/greater"`

	// This ordering relation for qualitative values indicates that the subject is greater than or equal to the object.
	//
	// RangeIncludes:
	// - http://schema.org/QualitativeValue
	//
	GreaterOrEqual *QualitativeValue `json:"greaterOrEqual,omitempty" jsonld:"http://schema.org/greaterOrEqual"`

	// This ordering relation for qualitative values indicates that the subject is lesser than the object.
	//
	// RangeIncludes:
	// - http://schema.org/QualitativeValue
	//
	Lesser *QualitativeValue `json:"lesser,omitempty" jsonld:"http://schema.org/lesser"`

	// This ordering relation for qualitative values indicates that the subject is lesser than or equal to the object.
	//
	// RangeIncludes:
	// - http://schema.org/QualitativeValue
	//
	LesserOrEqual *QualitativeValue `json:"lesserOrEqual,omitempty" jsonld:"http://schema.org/lesserOrEqual"`

	// This ordering relation for qualitative values indicates that the subject is not equal to the object.
	//
	// RangeIncludes:
	// - http://schema.org/QualitativeValue
	//
	NonEqual *QualitativeValue `json:"nonEqual,omitempty" jsonld:"http://schema.org/nonEqual"`

}

type QualitativeValue struct {
	MetaTrait
	QualitativeValueTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewQualitativeValue() (x QualitativeValue) {
	x.Type = "http://schema.org/QualitativeValue"

	return
}
