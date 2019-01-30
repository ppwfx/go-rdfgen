package schema

type OrganizationRoleTrait struct {

	// A number associated with a role in an organization, for example, the number on an athlete's jersey.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	NumberedPosition float64 `json:"numberedPosition,omitempty" jsonld:"http://schema.org/numberedPosition"`

}

type OrganizationRole struct {
	MetaTrait
	OrganizationRoleTrait
	RoleTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewOrganizationRole() (x OrganizationRole) {
	x.Type = "http://schema.org/OrganizationRole"

	return
}
