package schema

type GovernmentServiceTrait struct {

	// The operating organization, if different from the provider.  This enables the representation of services that are provided by an organization, but operated by another organization like a subcontractor.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	ServiceOperator *Organization `json:"serviceOperator,omitempty" jsonld:"http://schema.org/serviceOperator"`

}

type GovernmentService struct {
	MetaTrait
	GovernmentServiceTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewGovernmentService() (x GovernmentService) {
	x.Type = "http://schema.org/GovernmentService"

	return
}
