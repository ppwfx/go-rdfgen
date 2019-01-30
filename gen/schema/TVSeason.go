package schema

type TVSeasonTrait struct {

	// The country of the principal offices of the production company or individual responsible for the movie or program.
	//
	// RangeIncludes:
	// - http://schema.org/Country
	//
	CountryOfOrigin *Country `json:"countryOfOrigin,omitempty" jsonld:"http://schema.org/countryOfOrigin"`

	// The TV series to which this episode or season belongs.
	//
	// RangeIncludes:
	// - http://schema.org/TVSeries
	//
	PartOfTVSeries *TVSeries `json:"partOfTVSeries,omitempty" jsonld:"http://schema.org/partOfTVSeries"`

}

type TVSeason struct {
	MetaTrait
	TVSeasonTrait
	CreativeWorkTrait
	ThingTrait
	CreativeWorkSeasonTrait
	AdditionalTrait
}

func NewTVSeason() (x TVSeason) {
	x.Type = "http://schema.org/TVSeason"

	return
}
