package schema

type MiddleSchoolTrait struct {

}

type MiddleSchool struct {
	MetaTrait
	MiddleSchoolTrait
	EducationalOrganizationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewMiddleSchool() (x MiddleSchool) {
	x.Type = "http://schema.org/MiddleSchool"

	return
}
