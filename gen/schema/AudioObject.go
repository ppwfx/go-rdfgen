package schema

type AudioObjectTrait struct {

	// If this MediaObject is an AudioObject or VideoObject, the transcript of that object.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Transcript string `json:"transcript,omitempty" jsonld:"http://schema.org/transcript"`

}

type AudioObject struct {
	MetaTrait
	AudioObjectTrait
	MediaObjectTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewAudioObject() (x AudioObject) {
	x.Type = "http://schema.org/AudioObject"

	return
}
