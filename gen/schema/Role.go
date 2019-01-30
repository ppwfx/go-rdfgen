package schema

type RoleTrait struct {

	// The end date and time of the item (in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>).
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	// - http://schema.org/Date
	//
	EndDate interface{} `json:"endDate,omitempty" jsonld:"http://schema.org/endDate"`

	// The start date and time of the item (in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>).
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	// - http://schema.org/Date
	//
	StartDate interface{} `json:"startDate,omitempty" jsonld:"http://schema.org/startDate"`

	// A position played, performed or filled by a person or organization, as part of an organization. For example, an athlete in a SportsTeam might play in the position named 'Quarterback'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	NamedPosition interface{} `json:"namedPosition,omitempty" jsonld:"http://schema.org/namedPosition"`

	// A role played, performed or filled by a person or organization. For example, the team of creators for a comic book might fill the roles named 'inker', 'penciller', and 'letterer'; or an athlete in a SportsTeam might play in the position named 'Quarterback'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	RoleName interface{} `json:"roleName,omitempty" jsonld:"http://schema.org/roleName"`

}

type Role struct {
	MetaTrait
	RoleTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewRole() (x Role) {
	x.Type = "http://schema.org/Role"

	return
}
