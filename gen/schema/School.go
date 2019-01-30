package schema

type SchoolTrait struct {

}

type School struct {
	MetaTrait
	SchoolTrait
	EducationalOrganizationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewSchool() (x School) {
	x.Type = "http://schema.org/School"

	return
}
