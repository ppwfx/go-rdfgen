package schema

type ImageObjectTrait struct {

	// The caption for this object.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Caption string `json:"caption,omitempty" jsonld:"http://schema.org/caption"`

	// exif data for this object.
	//
	// RangeIncludes:
	// - http://schema.org/PropertyValue
	// - http://schema.org/Text
	//
	ExifData interface{} `json:"exifData,omitempty" jsonld:"http://schema.org/exifData"`

	// Indicates whether this image is representative of the content of the page.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	RepresentativeOfPage bool `json:"representativeOfPage,omitempty" jsonld:"http://schema.org/representativeOfPage"`

	// Thumbnail image for an image or video.
	//
	// RangeIncludes:
	// - http://schema.org/ImageObject
	//
	Thumbnail *ImageObject `json:"thumbnail,omitempty" jsonld:"http://schema.org/thumbnail"`

}

type ImageObject struct {
	MetaTrait
	ImageObjectTrait
	MediaObjectTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewImageObject() (x ImageObject) {
	x.Type = "http://schema.org/ImageObject"

	return
}
