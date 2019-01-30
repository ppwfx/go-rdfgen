package schema

type DefinedTermSetTrait struct {

	// A Defined Term contained in this term set.
	//
	// RangeIncludes:
	// - http://schema.org/DefinedTerm
	//
	HasDefinedTerm *DefinedTerm `json:"hasDefinedTerm,omitempty" jsonld:"http://schema.org/hasDefinedTerm"`

}

type DefinedTermSet struct {
	MetaTrait
	DefinedTermSetTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewDefinedTermSet() (x DefinedTermSet) {
	x.Type = "http://schema.org/DefinedTermSet"

	return
}
