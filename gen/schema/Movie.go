package schema

type MovieTrait struct {

	// The duration of the item (movie, audio recording, event, etc.) in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	Duration *Duration `json:"duration,omitempty" jsonld:"http://schema.org/duration"`

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

	// The country of the principal offices of the production company or individual responsible for the movie or program.
	//
	// RangeIncludes:
	// - http://schema.org/Country
	//
	CountryOfOrigin *Country `json:"countryOfOrigin,omitempty" jsonld:"http://schema.org/countryOfOrigin"`

	// Languages in which subtitles/captions are available, in <a href=\"http://tools.ietf.org/html/bcp47\">IETF BCP 47 standard format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Language
	//
	SubtitleLanguage interface{} `json:"subtitleLanguage,omitempty" jsonld:"http://schema.org/subtitleLanguage"`

}

type Movie struct {
	MetaTrait
	MovieTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMovie() (x Movie) {
	x.Type = "http://schema.org/Movie"

	return
}
