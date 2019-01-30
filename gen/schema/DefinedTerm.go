package schema

type DefinedTermTrait struct {

	// A <a class=\"localLink\" href=\"http://schema.org/DefinedTermSet\">DefinedTermSet</a> that contains this term.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	// - http://schema.org/DefinedTermSet
	//
	InDefinedTermSet interface{} `json:"inDefinedTermSet,omitempty" jsonld:"http://schema.org/inDefinedTermSet"`

	// A code that identifies this <a class=\"localLink\" href=\"http://schema.org/DefinedTerm\">DefinedTerm</a> within a <a class=\"localLink\" href=\"http://schema.org/DefinedTermSet\">DefinedTermSet</a>
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TermCode string `json:"termCode,omitempty" jsonld:"http://schema.org/termCode"`

}

type DefinedTerm struct {
	MetaTrait
	DefinedTermTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDefinedTerm() (x DefinedTerm) {
	x.Type = "http://schema.org/DefinedTerm"

	return
}
