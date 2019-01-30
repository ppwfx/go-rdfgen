package schema

type XPathTypeTrait struct {

}

type XPathType struct {
	MetaTrait
	XPathTypeTrait
	TextTrait
	AdditionalTrait
}

func NewXPathType() (x XPathType) {
	x.Type = "http://schema.org/XPathType"

	return
}
