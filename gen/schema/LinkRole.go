package schema

type LinkRoleTrait struct {

	// The language of the content or performance or used in an action. Please use one of the language codes from the <a href=\"http://tools.ietf.org/html/bcp47\">IETF BCP 47 standard</a>. See also <a class=\"localLink\" href=\"http://schema.org/availableLanguage\">availableLanguage</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Language
	// - http://schema.org/Text
	//
	InLanguage interface{} `json:"inLanguage,omitempty" jsonld:"http://schema.org/inLanguage"`

	// Indicates the relationship type of a Web link.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	LinkRelationship string `json:"linkRelationship,omitempty" jsonld:"http://schema.org/linkRelationship"`

}

type LinkRole struct {
	MetaTrait
	LinkRoleTrait
	RoleTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewLinkRole() (x LinkRole) {
	x.Type = "http://schema.org/LinkRole"

	return
}
