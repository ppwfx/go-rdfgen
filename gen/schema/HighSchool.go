package schema

type HighSchoolTrait struct {

}

type HighSchool struct {
	MetaTrait
	HighSchoolTrait
	EducationalOrganizationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewHighSchool() (x HighSchool) {
	x.Type = "http://schema.org/HighSchool"

	return
}
