package schema

type ElementarySchoolTrait struct {

}

type ElementarySchool struct {
	MetaTrait
	ElementarySchoolTrait
	EducationalOrganizationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewElementarySchool() (x ElementarySchool) {
	x.Type = "http://schema.org/ElementarySchool"

	return
}
