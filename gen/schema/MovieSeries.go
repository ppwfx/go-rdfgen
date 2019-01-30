package schema

type MovieSeriesTrait struct {

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

	// An actor, e.g. in tv, radio, movie, video games etc. Actors can be associated with individual items or with a series, episode, clip.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Actors *Person `json:"actors,omitempty" jsonld:"http://schema.org/actors"`

	// A director of e.g. tv, radio, movie, video games etc. content. Directors can be associated with individual items or with a series, episode, clip.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Directors *Person `json:"directors,omitempty" jsonld:"http://schema.org/directors"`

	// The composer of the soundtrack.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/MusicGroup
	//
	MusicBy interface{} `json:"musicBy,omitempty" jsonld:"http://schema.org/musicBy"`

	// The trailer of a movie or tv/radio series, season, episode, etc.
	//
	// RangeIncludes:
	// - http://schema.org/VideoObject
	//
	Trailer *VideoObject `json:"trailer,omitempty" jsonld:"http://schema.org/trailer"`

}

type MovieSeries struct {
	MetaTrait
	MovieSeriesTrait
	CreativeWorkSeriesTrait
	CreativeWorkTrait
	ThingTrait
	SeriesTrait
	IntangibleTrait
	AdditionalTrait
}

func NewMovieSeries() (x MovieSeries) {
	x.Type = "http://schema.org/MovieSeries"

	return
}
