package schema

type BusinessEntityTypeTrait struct {

}

type BusinessEntityType struct {
	MetaTrait
	BusinessEntityTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBusinessEntityType() (x BusinessEntityType) {
	x.Type = "http://schema.org/BusinessEntityType"

	return
}
