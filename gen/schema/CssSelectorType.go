package schema

type CssSelectorTypeTrait struct {

}

type CssSelectorType struct {
	MetaTrait
	CssSelectorTypeTrait
	TextTrait
	AdditionalTrait
}

func NewCssSelectorType() (x CssSelectorType) {
	x.Type = "http://schema.org/CssSelectorType"

	return
}
