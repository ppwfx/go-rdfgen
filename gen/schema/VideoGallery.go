package schema

type VideoGalleryTrait struct {

}

type VideoGallery struct {
	MetaTrait
	VideoGalleryTrait
	CollectionPageTrait
	WebPageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewVideoGallery() (x VideoGallery) {
	x.Type = "http://schema.org/VideoGallery"

	return
}
