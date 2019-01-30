package schema

type TVEpisodeTrait struct {

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

	// The TV series to which this episode or season belongs.
	//
	// RangeIncludes:
	// - http://schema.org/TVSeries
	//
	PartOfTVSeries *TVSeries `json:"partOfTVSeries,omitempty" jsonld:"http://schema.org/partOfTVSeries"`

}

type TVEpisode struct {
	MetaTrait
	TVEpisodeTrait
	EpisodeTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewTVEpisode() (x TVEpisode) {
	x.Type = "http://schema.org/TVEpisode"

	return
}
