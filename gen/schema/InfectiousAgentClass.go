package schema

type InfectiousAgentClassTrait struct {

}

type InfectiousAgentClass struct {
	MetaTrait
	InfectiousAgentClassTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewInfectiousAgentClass() (x InfectiousAgentClass) {
	x.Type = "http://schema.org/InfectiousAgentClass"

	return
}
