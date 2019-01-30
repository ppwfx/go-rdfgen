package schema

type ImageGalleryTrait struct {

}

type ImageGallery struct {
	MetaTrait
	ImageGalleryTrait
	CollectionPageTrait
	WebPageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewImageGallery() (x ImageGallery) {
	x.Type = "http://schema.org/ImageGallery"

	return
}
