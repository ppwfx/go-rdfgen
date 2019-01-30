package schema

type GovernmentOrganizationTrait struct {

}

type GovernmentOrganization struct {
	MetaTrait
	GovernmentOrganizationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewGovernmentOrganization() (x GovernmentOrganization) {
	x.Type = "http://schema.org/GovernmentOrganization"

	return
}
