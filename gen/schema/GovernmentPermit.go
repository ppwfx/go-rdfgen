package schema

type GovernmentPermitTrait struct {

}

type GovernmentPermit struct {
	MetaTrait
	GovernmentPermitTrait
	PermitTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewGovernmentPermit() (x GovernmentPermit) {
	x.Type = "http://schema.org/GovernmentPermit"

	return
}
