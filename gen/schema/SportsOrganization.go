package schema

type SportsOrganizationTrait struct {

	// A type of sport (e.g. Baseball).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	Sport interface{} `json:"sport,omitempty" jsonld:"http://schema.org/sport"`

}

type SportsOrganization struct {
	MetaTrait
	SportsOrganizationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewSportsOrganization() (x SportsOrganization) {
	x.Type = "http://schema.org/SportsOrganization"

	return
}
