package schema

type GenderTypeTrait struct {

}

type GenderType struct {
	MetaTrait
	GenderTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewGenderType() (x GenderType) {
	x.Type = "http://schema.org/GenderType"

	return
}
