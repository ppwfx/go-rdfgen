package schema

type PreschoolTrait struct {

}

type Preschool struct {
	MetaTrait
	PreschoolTrait
	EducationalOrganizationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewPreschool() (x Preschool) {
	x.Type = "http://schema.org/Preschool"

	return
}
