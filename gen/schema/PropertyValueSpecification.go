package schema

type PropertyValueSpecificationTrait struct {

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

	// The default value of the input.  For properties that expect a literal, the default is a literal value, for properties that expect an object, it's an ID reference to one of the current values.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Thing
	//
	DefaultValue interface{} `json:"defaultValue,omitempty" jsonld:"http://schema.org/defaultValue"`

	// Whether multiple values are allowed for the property.  Default is false.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	MultipleValues bool `json:"multipleValues,omitempty" jsonld:"http://schema.org/multipleValues"`

	// Whether or not a property is mutable.  Default is false. Specifying this for a property that also has a value makes it act similar to a \"hidden\" input in an HTML form.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	ReadonlyValue bool `json:"readonlyValue,omitempty" jsonld:"http://schema.org/readonlyValue"`

	// The stepValue attribute indicates the granularity that is expected (and required) of the value in a PropertyValueSpecification.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	StepValue float64 `json:"stepValue,omitempty" jsonld:"http://schema.org/stepValue"`

	// Specifies the allowed range for number of characters in a literal value.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	ValueMaxLength float64 `json:"valueMaxLength,omitempty" jsonld:"http://schema.org/valueMaxLength"`

	// Specifies the minimum allowed range for number of characters in a literal value.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	ValueMinLength float64 `json:"valueMinLength,omitempty" jsonld:"http://schema.org/valueMinLength"`

	// Indicates the name of the PropertyValueSpecification to be used in URL templates and form encoding in a manner analogous to HTML's input@name.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ValueName string `json:"valueName,omitempty" jsonld:"http://schema.org/valueName"`

	// Specifies a regular expression for testing literal values according to the HTML spec.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ValuePattern string `json:"valuePattern,omitempty" jsonld:"http://schema.org/valuePattern"`

	// Whether the property must be filled in to complete the action.  Default is false.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	ValueRequired bool `json:"valueRequired,omitempty" jsonld:"http://schema.org/valueRequired"`

}

type PropertyValueSpecification struct {
	MetaTrait
	PropertyValueSpecificationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewPropertyValueSpecification() (x PropertyValueSpecification) {
	x.Type = "http://schema.org/PropertyValueSpecification"

	return
}
