package schema

type InfectiousDiseaseTrait struct {

	// How the disease spreads, either as a route or vector, for example 'direct contact', 'Aedes aegypti', etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TransmissionMethod string `json:"transmissionMethod,omitempty" jsonld:"http://schema.org/transmissionMethod"`

	// The actual infectious agent, such as a specific bacterium.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	InfectiousAgent string `json:"infectiousAgent,omitempty" jsonld:"http://schema.org/infectiousAgent"`

	// The class of infectious agent (bacteria, prion, etc.) that causes the disease.
	//
	// RangeIncludes:
	// - http://schema.org/InfectiousAgentClass
	//
	InfectiousAgentClass *InfectiousAgentClass `json:"infectiousAgentClass,omitempty" jsonld:"http://schema.org/infectiousAgentClass"`

}

type InfectiousDisease struct {
	MetaTrait
	InfectiousDiseaseTrait
	MedicalConditionTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewInfectiousDisease() (x InfectiousDisease) {
	x.Type = "http://schema.org/InfectiousDisease"

	return
}
