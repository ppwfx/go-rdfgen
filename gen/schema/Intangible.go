package schema

type IntangibleTrait struct {

}

type Intangible struct {
	MetaTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewIntangible() (x Intangible) {
	x.Type = "http://schema.org/Intangible"

	return
}
