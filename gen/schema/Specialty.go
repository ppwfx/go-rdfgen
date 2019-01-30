package schema

type SpecialtyTrait struct {

}

type Specialty struct {
	MetaTrait
	SpecialtyTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewSpecialty() (x Specialty) {
	x.Type = "http://schema.org/Specialty"

	return
}
