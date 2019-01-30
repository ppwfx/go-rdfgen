package schema

type SportsTeamTrait struct {

	// A person that acts in a coaching role for a sports team.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Coach *Person `json:"coach,omitempty" jsonld:"http://schema.org/coach"`

	// A person that acts as performing member of a sports team; a player as opposed to a coach.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Athlete *Person `json:"athlete,omitempty" jsonld:"http://schema.org/athlete"`

}

type SportsTeam struct {
	MetaTrait
	SportsTeamTrait
	SportsOrganizationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewSportsTeam() (x SportsTeam) {
	x.Type = "http://schema.org/SportsTeam"

	return
}
