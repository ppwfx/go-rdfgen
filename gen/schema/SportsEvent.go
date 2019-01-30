package schema

type SportsEventTrait struct {

	// The away team in a sports event.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/SportsTeam
	//
	AwayTeam interface{} `json:"awayTeam,omitempty" jsonld:"http://schema.org/awayTeam"`

	// A competitor in a sports event.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/SportsTeam
	//
	Competitor interface{} `json:"competitor,omitempty" jsonld:"http://schema.org/competitor"`

	// The home team in a sports event.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/SportsTeam
	//
	HomeTeam interface{} `json:"homeTeam,omitempty" jsonld:"http://schema.org/homeTeam"`

}

type SportsEvent struct {
	MetaTrait
	SportsEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewSportsEvent() (x SportsEvent) {
	x.Type = "http://schema.org/SportsEvent"

	return
}
