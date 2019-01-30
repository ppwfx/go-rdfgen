package schema

type AudiobookTrait struct {

	// The duration of the item (movie, audio recording, event, etc.) in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	Duration *Duration `json:"duration,omitempty" jsonld:"http://schema.org/duration"`

	// A person who reads (performs) the audiobook.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	ReadBy *Person `json:"readBy,omitempty" jsonld:"http://schema.org/readBy"`

}

type Audiobook struct {
	MetaTrait
	AudiobookTrait
	AudioObjectTrait
	MediaObjectTrait
	CreativeWorkTrait
	ThingTrait
	BookTrait
	AdditionalTrait
}

func NewAudiobook() (x Audiobook) {
	x.Type = "http://schema.org/Audiobook"

	return
}
