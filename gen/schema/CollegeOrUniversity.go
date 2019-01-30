package schema

type CollegeOrUniversityTrait struct {

}

type CollegeOrUniversity struct {
	MetaTrait
	CollegeOrUniversityTrait
	EducationalOrganizationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewCollegeOrUniversity() (x CollegeOrUniversity) {
	x.Type = "http://schema.org/CollegeOrUniversity"

	return
}
