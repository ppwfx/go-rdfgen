package schema

type StructuredValueTrait struct {

}

type StructuredValue struct {
	MetaTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewStructuredValue() (x StructuredValue) {
	x.Type = "http://schema.org/StructuredValue"

	return
}
