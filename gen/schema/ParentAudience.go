package schema

type ParentAudienceTrait struct {

	// Maximal age of the child.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	ChildMaxAge float64 `json:"childMaxAge,omitempty" jsonld:"http://schema.org/childMaxAge"`

	// Minimal age of the child.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	ChildMinAge float64 `json:"childMinAge,omitempty" jsonld:"http://schema.org/childMinAge"`

}

type ParentAudience struct {
	MetaTrait
	ParentAudienceTrait
	PeopleAudienceTrait
	AudienceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewParentAudience() (x ParentAudience) {
	x.Type = "http://schema.org/ParentAudience"

	return
}
