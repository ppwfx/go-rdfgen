package schema

type VideoObjectTrait struct {

	// The caption for this object.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Caption string `json:"caption,omitempty" jsonld:"http://schema.org/caption"`

	// Thumbnail image for an image or video.
	//
	// RangeIncludes:
	// - http://schema.org/ImageObject
	//
	Thumbnail *ImageObject `json:"thumbnail,omitempty" jsonld:"http://schema.org/thumbnail"`

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

	// If this MediaObject is an AudioObject or VideoObject, the transcript of that object.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Transcript string `json:"transcript,omitempty" jsonld:"http://schema.org/transcript"`

	// The frame size of the video.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	VideoFrameSize string `json:"videoFrameSize,omitempty" jsonld:"http://schema.org/videoFrameSize"`

	// The quality of the video.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	VideoQuality string `json:"videoQuality,omitempty" jsonld:"http://schema.org/videoQuality"`

}

type VideoObject struct {
	MetaTrait
	VideoObjectTrait
	MediaObjectTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewVideoObject() (x VideoObject) {
	x.Type = "http://schema.org/VideoObject"

	return
}
