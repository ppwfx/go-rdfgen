package schema

type EpisodeTrait struct {

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

	// The series to which this episode or season belongs.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWorkSeries
	//
	PartOfSeries *CreativeWorkSeries `json:"partOfSeries,omitempty" jsonld:"http://schema.org/partOfSeries"`

	// Position of the episode within an ordered group of episodes.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Integer
	//
	EpisodeNumber interface{} `json:"episodeNumber,omitempty" jsonld:"http://schema.org/episodeNumber"`

	// The season to which this episode belongs.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWorkSeason
	//
	PartOfSeason *CreativeWorkSeason `json:"partOfSeason,omitempty" jsonld:"http://schema.org/partOfSeason"`

}

type Episode struct {
	MetaTrait
	EpisodeTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewEpisode() (x Episode) {
	x.Type = "http://schema.org/Episode"

	return
}
