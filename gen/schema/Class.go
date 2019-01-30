package schema

type ClassTrait struct {

	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	//
	// RangeIncludes:
	// - http://schema.org/Class
	// - http://schema.org/Enumeration
	// - http://schema.org/Property
	//
	SupersededBy interface{} `json:"supersededBy,omitempty" jsonld:"http://schema.org/supersededBy"`

}

type Class struct {
	MetaTrait
	ClassTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewClass() (x Class) {
	x.Type = "http://schema.org/Class"

	return
}
