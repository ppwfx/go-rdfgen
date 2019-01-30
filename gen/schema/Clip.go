package schema

type ClipTrait struct {

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

	// The series to which this episode or season belongs.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWorkSeries
	//
	PartOfSeries *CreativeWorkSeries `json:"partOfSeries,omitempty" jsonld:"http://schema.org/partOfSeries"`

	// The season to which this episode belongs.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWorkSeason
	//
	PartOfSeason *CreativeWorkSeason `json:"partOfSeason,omitempty" jsonld:"http://schema.org/partOfSeason"`

	// The episode to which this clip belongs.
	//
	// RangeIncludes:
	// - http://schema.org/Episode
	//
	PartOfEpisode *Episode `json:"partOfEpisode,omitempty" jsonld:"http://schema.org/partOfEpisode"`

	// Position of the clip within an ordered group of clips.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Integer
	//
	ClipNumber interface{} `json:"clipNumber,omitempty" jsonld:"http://schema.org/clipNumber"`

}

type Clip struct {
	MetaTrait
	ClipTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewClip() (x Clip) {
	x.Type = "http://schema.org/Clip"

	return
}
