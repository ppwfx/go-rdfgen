package schema

type BusinessFunctionTrait struct {

}

type BusinessFunction struct {
	MetaTrait
	BusinessFunctionTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBusinessFunction() (x BusinessFunction) {
	x.Type = "http://schema.org/BusinessFunction"

	return
}
