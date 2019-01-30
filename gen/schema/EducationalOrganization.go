package schema

type EducationalOrganizationTrait struct {

	// Alumni of an organization.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Alumni *Person `json:"alumni,omitempty" jsonld:"http://schema.org/alumni"`

}

type EducationalOrganization struct {
	MetaTrait
	EducationalOrganizationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewEducationalOrganization() (x EducationalOrganization) {
	x.Type = "http://schema.org/EducationalOrganization"

	return
}
