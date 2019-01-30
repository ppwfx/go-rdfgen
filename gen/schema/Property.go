package schema

type PropertyTrait struct {

	// Relates a property to a class that is (one of) the type(s) the property is expected to be used on.
	//
	// RangeIncludes:
	// - http://schema.org/Class
	//
	DomainIncludes *Class `json:"domainIncludes,omitempty" jsonld:"http://schema.org/domainIncludes"`

	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	//
	// RangeIncludes:
	// - http://schema.org/Class
	// - http://schema.org/Enumeration
	// - http://schema.org/Property
	//
	SupersededBy interface{} `json:"supersededBy,omitempty" jsonld:"http://schema.org/supersededBy"`

	// Relates a property to a property that is its inverse. Inverse properties relate the same pairs of items to each other, but in reversed direction. For example, the 'alumni' and 'alumniOf' properties are inverseOf each other. Some properties don't have explicit inverses; in these situations RDFa and JSON-LD syntax for reverse properties can be used.
	//
	// RangeIncludes:
	// - http://schema.org/Property
	//
	InverseOf *Property `json:"inverseOf,omitempty" jsonld:"http://schema.org/inverseOf"`

	// Relates a property to a class that constitutes (one of) the expected type(s) for values of the property.
	//
	// RangeIncludes:
	// - http://schema.org/Class
	//
	RangeIncludes *Class `json:"rangeIncludes,omitempty" jsonld:"http://schema.org/rangeIncludes"`

}

type Property struct {
	MetaTrait
	PropertyTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewProperty() (x Property) {
	x.Type = "http://schema.org/Property"

	return
}
