package schema

type CreativeWorkSeasonTrait struct {

	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	ProductionCompany *Organization `json:"productionCompany,omitempty" jsonld:"http://schema.org/productionCompany"`

	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Actor *Person `json:"actor,omitempty" jsonld:"http://schema.org/actor"`

	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Director *Person `json:"director,omitempty" jsonld:"http://schema.org/director"`

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

	// The trailer of a movie or tv/radio series, season, episode, etc.
	//
	// RangeIncludes:
	// - http://schema.org/VideoObject
	//
	Trailer *VideoObject `json:"trailer,omitempty" jsonld:"http://schema.org/trailer"`

	// An episode of a tv, radio or game media within a series or season.
	//
	// RangeIncludes:
	// - http://schema.org/Episode
	//
	Episode *Episode `json:"episode,omitempty" jsonld:"http://schema.org/episode"`

	// An episode of a TV/radio series or season.
	//
	// RangeIncludes:
	// - http://schema.org/Episode
	//
	Episodes *Episode `json:"episodes,omitempty" jsonld:"http://schema.org/episodes"`

	// The number of episodes in this season or series.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	NumberOfEpisodes *Integer `json:"numberOfEpisodes,omitempty" jsonld:"http://schema.org/numberOfEpisodes"`

	// The series to which this episode or season belongs.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWorkSeries
	//
	PartOfSeries *CreativeWorkSeries `json:"partOfSeries,omitempty" jsonld:"http://schema.org/partOfSeries"`

	// Position of the season within an ordered group of seasons.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Integer
	//
	SeasonNumber interface{} `json:"seasonNumber,omitempty" jsonld:"http://schema.org/seasonNumber"`

}

type CreativeWorkSeason struct {
	MetaTrait
	CreativeWorkSeasonTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewCreativeWorkSeason() (x CreativeWorkSeason) {
	x.Type = "http://schema.org/CreativeWorkSeason"

	return
}
