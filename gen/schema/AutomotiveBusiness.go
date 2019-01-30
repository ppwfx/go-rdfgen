package schema

type AutomotiveBusinessTrait struct {

}

type AutomotiveBusiness struct {
	MetaTrait
	AutomotiveBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewAutomotiveBusiness() (x AutomotiveBusiness) {
	x.Type = "http://schema.org/AutomotiveBusiness"

	return
}
