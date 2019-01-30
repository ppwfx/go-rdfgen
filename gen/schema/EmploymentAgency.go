package schema

type EmploymentAgencyTrait struct {

}

type EmploymentAgency struct {
	MetaTrait
	EmploymentAgencyTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewEmploymentAgency() (x EmploymentAgency) {
	x.Type = "http://schema.org/EmploymentAgency"

	return
}
