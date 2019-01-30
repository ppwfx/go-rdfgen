package schema

type SpeakableSpecificationTrait struct {

	// A CSS selector, e.g. of a <a class=\"localLink\" href=\"http://schema.org/SpeakableSpecification\">SpeakableSpecification</a> or <a class=\"localLink\" href=\"http://schema.org/WebPageElement\">WebPageElement</a>. In the latter case, multiple matches within a page can constitute a single conceptual \"Web page element\".
	//
	// RangeIncludes:
	// - http://schema.org/CssSelectorType
	//
	CssSelector *CssSelectorType `json:"cssSelector,omitempty" jsonld:"http://schema.org/cssSelector"`

	// An XPath, e.g. of a <a class=\"localLink\" href=\"http://schema.org/SpeakableSpecification\">SpeakableSpecification</a> or <a class=\"localLink\" href=\"http://schema.org/WebPageElement\">WebPageElement</a>. In the latter case, multiple matches within a page can constitute a single conceptual \"Web page element\".
	//
	// RangeIncludes:
	// - http://schema.org/XPathType
	//
	Xpath *XPathType `json:"xpath,omitempty" jsonld:"http://schema.org/xpath"`

}

type SpeakableSpecification struct {
	MetaTrait
	SpeakableSpecificationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewSpeakableSpecification() (x SpeakableSpecification) {
	x.Type = "http://schema.org/SpeakableSpecification"

	return
}
