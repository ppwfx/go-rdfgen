package schema

type EnumerationTrait struct {

	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	//
	// RangeIncludes:
	// - http://schema.org/Class
	// - http://schema.org/Enumeration
	// - http://schema.org/Property
	//
	SupersededBy interface{} `json:"supersededBy,omitempty" jsonld:"http://schema.org/supersededBy"`

}

type Enumeration struct {
	MetaTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewEnumeration() (x Enumeration) {
	x.Type = "http://schema.org/Enumeration"

	return
}
