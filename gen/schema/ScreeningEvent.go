package schema

type ScreeningEventTrait struct {

	// Languages in which subtitles/captions are available, in <a href=\"http://tools.ietf.org/html/bcp47\">IETF BCP 47 standard format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Language
	//
	SubtitleLanguage interface{} `json:"subtitleLanguage,omitempty" jsonld:"http://schema.org/subtitleLanguage"`

	// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	VideoFormat string `json:"videoFormat,omitempty" jsonld:"http://schema.org/videoFormat"`

	// The movie presented during this event.
	//
	// RangeIncludes:
	// - http://schema.org/Movie
	//
	WorkPresented *Movie `json:"workPresented,omitempty" jsonld:"http://schema.org/workPresented"`

}

type ScreeningEvent struct {
	MetaTrait
	ScreeningEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewScreeningEvent() (x ScreeningEvent) {
	x.Type = "http://schema.org/ScreeningEvent"

	return
}
